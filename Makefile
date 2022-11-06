build:
	go build -ldflags="-s -w" myctl.go
	$(if $(shell command -v upx), upx goctl)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o myctl-darwin goctl.go
	$(if $(shell command -v upx), upx goctl-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o myctl.exe goctl.go
	$(if $(shell command -v upx), upx goctl.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o myctl-linux goctl.go
	$(if $(shell command -v upx), upx goctl-linux)
