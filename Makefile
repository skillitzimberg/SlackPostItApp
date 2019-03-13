all: build
build: build-famis |
		cd filesystem_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
		cd database/postgres_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
		cd google_sheets_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
		cd convert_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
		cd common_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
		cd NetCoreSteps/Accruent.Famis.Steps && dotnet publish -o publish -c Release

build-famis: |
	cd NetCoreSteps/Accruent.Famis.Steps && dotnet publish -o publish -c Release

build-filesystem: |
	cd filesystem_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
buildoracle: |
	cd database/oracle_pkg && env CC=x86_64-w64-mingw32-gcc && gox -osarch="windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
clean: |
	./scripts/clean_steps.sh
updatesdk: |
	cd filesystem_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/db_common && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/oracle_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/postgres_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd google_sheets_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd convert_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd common_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
publish: build |
	apptree workflow package publish -d common_pkg
	apptree workflow package publish -d google_sheets_pkg
	apptree workflow package publish -d database/postgres_pkg
	apptree workflow package publish -d convert_pkg
	apptree workflow package publish -d common_pkg
	apptree workflow package publish -d filesystem_pkg
	apptree workflow package publish -d NetCoreSteps/Accruent.Famis.Steps/publish

publish-famis: build-famis |
	apptree workflow package publish -d NetCoreSteps/Accruent.Famis.Steps/publish