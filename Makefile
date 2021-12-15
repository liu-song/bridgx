vet:
	@echo "go vet ."
	@go vet $$(go list ./...) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

check: vet

format:
	#go get golang.org/x/tools/cmd/goimports
	find . -name '*.go' | grep -Ev 'vendor|thrift_gen' | xargs goimports -w

build:
	sh ./scripts/build_api.sh && sh ./scripts/build_scheduler.sh

run:
	sh ./output/run_api.sh

clean:
	rm -rf output

server: format clean build run

docker-build-scheduler:
	docker build -t 172.16.16.172:12380/bridgx/bridgx-scheduler:v0.2 -f ./SCHEDULER.Dockerfile ./

docker-build-api:
	docker build -t 172.16.16.172:12380/bridgx/bridgx-api:v0.2 -f ./API.Dockerfile ./

docker-push-scheduler:
	docker push 172.16.16.172:12380/bridgx/bridgx-scheduler:v0.2

docker-push-api:
	docker push 172.16.16.172:12380/bridgx/bridgx-api:v0.2

docker-all: clean docker-build-scheduler docker-build-api docker-push-scheduler docker-push-api

# Quick start
# Pull images from dockerhub and run
docker-run-linux:
	#@echo "install mysql and etcd, yes or no? "${install_mysql_etcd}
	@echo -n "install mysql? [y/N] " && read ans && if [ $${ans:-'N'} = 'y' ]; then make docker-run-mysql; fi
	@echo -n "install etcd? [y/N] " && read ans && if [ $${ans:-'N'} = 'y' ]; then make docker-run-etcd; fi
    #sh ./run-for-linux.sh

docker-run-mac:
	sh ./run-for-mac.sh

docker-run-mysql:
	docker run -d --name bridgx_db -e MYSQL_ROOT_PASSWORD=mtQ8chN2 -e MYSQL_DATABASE=bridgx -e MYSQL_USER=gf -e MYSQL_PASSWORD=db@galaxy-future.com -p 3306:3306 -v $(pwd)/init/mysql:/docker-entrypoint-initdb.d yobasystems/alpine-mariadb:10.5.11

docker-run-etcd:
	docker run -d --name bridgx_etcd -e ALLOW_NONE_AUTHENTICATION=yes -e ETCD_ADVERTISE_CLIENT_URLS=http://etcd:2379 -p 2379:2379 -p 2380:2380 bitnami/etcd:3

docker-container-stop:
	docker ps -aq | xargs docker stop
	docker ps -aq | xargs docker rm

docker-image-rm:
	docker image prune --force --all

# Immersive experience
# Compile and run by docker-compose
docker-compose-start:
	docker-compose up -d

docker-compose-stop:
	docker-compose down

docker-compose-build:
	docker-compose build

docker-tag:
	docker tag bridgx_api galaxyfuture/bridgx-api:latest
	docker tag bridgx_scheduler galaxyfuture/bridgx-scheduler:latest

docker-push-hub:
	docker push galaxyfuture/bridgx-api:latest
	docker push galaxyfuture/bridgx-scheduler:latest

docker-hub-all: docker-compose-build docker-tag docker-push-hub