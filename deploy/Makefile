.PHONY: volumes
volumes:
	chmod u+x volumes.sh
	sh volumes.sh

.PHONY: compose
# 失败查看 ~/.docker/confing.json, 去除s,使其找不到win11的docker，使用默认linux的docker
compose:
	docker compose -f ./docker-compose.yml down
	docker compose -f ./docker-compose.yml up -d --force-recreate

