all: build
build: |
	./scripts/build_steps.sh
clean: |
	./scripts/clean_steps.sh
publish: build |
	apptree workflow package publish -d common_pkg
	apptree workflow package publish -d google_sheets_pkg