.PHONY: help
help:
	@echo "help! you know I need some... help!"

.PHONY: all
all: build

.PHONY: run
run: build
	@$(shell basename `pwd`)

.PHONY: exec
exec:
	@$(shell basename `pwd`)

.PHONY: build
build:
	@go build .

.PHONY: osx
osx:
	@GOOS=darwin GOARCH=amd64 go build .

.PHONY: linux
linux:
	@GOOS=linux GOARCH=amd64 go build .

.PHONY: install
install:
	@go install .

.PHONY: regen
regen: backup clean gen redux restore

.PHONY: backup
backup:
	@echo "backing up"
	@find . -maxdepth 1 -name "*.go" -exec cp -f {} {}.backup \;

.PHONY: clean
clean:
	@echo "cleaning"
	@rm -rf app client swagger tool
	@find . -maxdepth 1 -name "*.go" -exec rm -f {} \;

.PHONY: gen
gen:
	@echo "generating"
	@goagen bootstrap -d $(shell basename `pwd`)/design

.PHONY: redux
redux:
	@echo "gen'n redux"
	@goagen -d $(shell basename `pwd`)/design gen --pkg-path=github.com/verdverm/goa-redux

.PHONY: restore
restore:
	@echo "restoring"
	@./restore.py

.PHONY: clean_backups
clean_backups:
	@find . -maxdepth 1 -name "*.backup" -exec rm -f {} \;

.PHONY: clean_all
clean_all: clean clean_backups

.PHONY: swagger
swagger:
	@goagen swagger -d $(shell basename `pwd`)/design

.PHONY: cli
cli:
	cd tool/$(shell basename `pwd`)-cli && go install .
