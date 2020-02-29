APP = torimo-article-api

.PHONY: log
log:
	docker-compose logs $(APP)

.PHONY: rm
rm:
	-docker rm $(shell docker ps -a --filter 'status=exited' -q)

.PHONY: run
run: up log

.PHONY: up
up:
	docker-compose up -d

.PHONY: reload
reload: rm build restart log

.PHONY: restart
restart:
	docker-compose restart $(APP)

.PHONY: build
build:
	docker-compose build $(APP)

.PHONY: clean
clean: down rm rmi

.PHONY: down
down:
	docker-compose down

.PHONY: rmi
rmi:
	docker image prune -f
