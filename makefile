.DEFAULT_GOAL := help

PROJECTNAME=$(shell basename "$(PWD)")

install: go get 
build:
	@mkdir -p /opt/svm
	@cp ./configs/default.json /opt/svm/
	@go build -o $(SVM)/svm ./cmd/svm/main.go
build-local:
	@go build -o bin/svm ./cmd/svm/main.go
help: makefile 
	@echo "Choose a command run in "$(PROJECTNAME)":"
	@echo "  install\tInstall missing dependencies. Runs 'go get' internally. "
	@echo "  build\t\tBuild binary and place in 'GOBIN'"
	@echo "  build-local\tBuild binary and place in ./bin"

