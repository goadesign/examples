design_path="$(shell pwd | sed "s|^${GOPATH}/src/||")/design"

.PHONY: regen
regen: backup clean gen restore

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
	@echo "generating $(design_path)"
	@goagen bootstrap -d $(design_path)

.PHONY: restore
restore:
	@echo "restoring"
	@./restore.py

.PHONY: clean_backups
clean_backups:
	@find . -maxdepth 1 -name "*.backup" -exec rm -f {} \;

.PHONY: clean_all
clean_all: clean clean_backups

