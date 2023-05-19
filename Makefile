GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
		   --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
		   --validate_out=paths=source_relative,lang=go:./api \
	       $(API_PROTO_FILES)

	sh scripts/http-omitempty.sh

.PHONY: compose
compose:
	docker compose -f ./deploy/docker-compose.yml down
	docker compose -f ./deploy/docker-compose.yml up -d --force-recreate

.PHONY: openapi
# generate openapi
openapi:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --openapi_out=. \
           $(API_PROTO_FILES)

.PHONY: doc
# generate html or markdown doc
doc:
	protoc --proto_path=. \
           --proto_path=./third_party \
	   	   --doc_out=. \
	   	   --doc_opt=html,index.html \
	   	   $(API_PROTO_FILES)

.PHONY: dockerBuild
# generate dockerBuild
dockerBuild:
	sh ./scripts/docker-build.sh

.PHONY: build
# generate build
build:
	go build -o server cmd/main.go cmd/wire_gen.go

.PHONY: sql
# generate sql
sql:
	gentool -dsn "root:123456@tcp(172.21.0.2:3306)/go-layout?charset=utf8mb4&parseTime=True&loc=Local" --modelPkgName="./internal/data/model" -outPath="./internal/data/gen"

.PHONY: configs
# generate configs
configs:
	kratos proto client configs/conf/conf.proto

.PHONY: all
# generate all
all:
	make init;
	make api;
	make compose;
	make sql;
	make configs;
	make build;
	make dockerBuild;
	make run;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
