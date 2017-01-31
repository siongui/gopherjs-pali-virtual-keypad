export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

PKG="github.com/siongui/gopherjs-pali-virtual-keypad"
GO_VERSION=1.7.5
DEV_DIR=../

devserver: fmt local js
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run devserver/server.go -dir=example

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@gopherjs build example/demo.go -o example/demo.js

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
	@go fmt example/*.go
	@go fmt devserver/*.go

local:
	@[ -d src/${PKG}/ ] || mkdir -p src/${PKG}/
	@cp keypad.go src/${PKG}/

install:
	@echo "\033[92mInstalling GopherJS and necessary packages ...\033[0m"
	go get -u github.com/gopherjs/gopherjs
	@#go get -u ${PKG}

deploy:
	@echo "\033[92mDeploy to GitHub Pages (Project) ...\033[0m"
	@rm example/*.go
	@ghp-import example/
	@git push origin gh-pages
	@git checkout example/

update_ubuntu:
	@echo "\033[92mUpdating Ubuntu ...\033[0m"
	@sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@cd $(DEV_DIR) ; wget https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
	@cd $(DEV_DIR) ; tar xvzf go$(GO_VERSION).linux-amd64.tar.gz
