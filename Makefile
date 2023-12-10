.PHONY: all test

# Search all directories to find ones containing a Makefile
DIRS := $(shell find . -type f -name Makefile -not -path "./Makefile" -exec dirname {} \;)

# Set the default version to "latest". Specify externally with make GOA_VERSION=<desired_version>
GOA_VERSION ?= latest

# Execute make in all identified directories
all:
	@for dir in $(DIRS); do \
		echo "Making in $$dir..."; \
		(cd $$dir && go get -u goa.design/goa/v3@$(GOA_VERSION) && go mod tidy && make) || exit 1; \
	done

# Execute "make test" in all identified directories
test:
	@for dir in $(DIRS); do \
		echo "Running tests in $$dir..."; \
		(cd $$dir && make test) || exit 1; \
	done
