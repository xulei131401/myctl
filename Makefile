isSupportUpxCommand := $(shell command -v upx)
mainGo := "myctl"

.PHONY: all tidy run clean help

build:
	@go build -ldflags="-s -w" myctl.go
	$(if ($(isSupportUpxCommand)), @upx myctl)

.PHONY: mac
mac:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/myctl-darwin-amd64 $(mainGo) .go
	$(if ($(isSupportUpxCommand)), @upx myctl-darwin)

.PHONY: win
win:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/myctl-windows-amd64.exe $(mainGo) .go
	$(if ($(isSupportUpxCommand)), @upx myctl.exe)


.PHONY: linux
linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/myctl-linux-amd64 $(mainGo) .go
	$(if ($(isSupportUpxCommand)), @upx myctl-linux)

all:win linux mac

.PHONY: tidy
tidy:
	@go mod tidy

run:
	@go run ./

clean:
	@if [ -f ./bin/${mainGo}-linux64 ] ; then rm ./bin/myctl-linux-amd64; fi
	@if [ -f ./bin/${mainGo}-win64.exe ] ; then rm ./bin/myctl-windows-amd64.exe; fi
	@if [ -f ./bin/${mainGo}-darwin64 ] ; then rm ./bin/myctl-darwin-amd64; fi

help:
	@echo "make (make build) - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make mac - 编译 Go 代码, 生成mac的二进制文件"
	@echo "make linux - 编译 Go 代码, 生成linux二进制文件"
	@echo "make win - 编译 Go 代码, 生成windows二进制文件"
	@echo "make tidy - 执行go mod tidy"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除编译的二进制文件"
	@echo "make all - 编译多平台的二进制文件"