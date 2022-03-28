GOPATH=${HOME}/go

%:
	@true

.PHONY: run
run:
	./scripts/fmt.sh $(filter-out $@,$(MAKECMDGOALS))
	./scripts/run.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: install
install:
	./scripts/install.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: build
build:
	./scripts/build.sh $(filter-out $@,$(MAKECMDGOALS))