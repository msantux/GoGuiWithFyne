BINARY_NAME=GoldWatcher.app
APP_NAME=GoldWatcher
VERSION=1.0.1
BUILD_NO=2

## build: build binary and package it
build:
	rm -rf ${BINARY_NAME}
	/Users/santux/go/bin/fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release
	rm -f go-for-gold

## run: builds and runs the application
run:
	env DB_PATH="./sql.db" go run .

## clean runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleanned!"

## test: runs all tests
test:
	go test -v ./.