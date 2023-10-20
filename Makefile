ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
APP=oxford-bot
APP_EXECUTABLE="./out/$(APP)"
APP_VERSION:="1.0"


fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

tidy:
	go mod tidy

serve: fmt vet
	env $(cat test.env | xargs) CGO_ENABLED=0 go run cmd/*.go

compile:
	mkdir out
	CGO_ENABLED=0 go build -o $(APP_EXECUTABLE) -ldflags "-X main.version=$(APP_VERSION)" cmd/*.go
