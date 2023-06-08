.PHONY: volumes
volumes:
	chmod u+x volumes.sh
	sh volumes.sh

.PHONY: compose
compose:
	docker compose -f ./docker-compose.yml down
	docker compose -f ./docker-compose.yml up -d --force-recreate

