HOST=https://io.apptreesoftware.com

test: |
	echo ${HOST}
all: publish
build: build-dotnet build-go |
build-go: build-filesystem build-postgres build-googlesheets build-convert build-common build-logger build-webhook build-cache build-facility360 build-script
build-dotnet: build-famis
build-postgres: |
			cd database/postgres_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-postgres: build-postgres |
	apptree workflow package publish -d database/postgres_pkg --host ${HOST}
build-script: |
	cd script_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-script: build-script |
	apptree workflow package publish -d script_pkg --host ${HOST}
build-googlesheets: |
	cd google_sheets_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-googlesheets: build-googlesheets |
	apptree workflow package publish -d google_sheets_pkg --host ${HOST}
build-convert: |
	cd convert_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-convert: build-convert |
	apptree workflow package publish -d convert_pkg --host ${HOST}
build-common: |
	cd common_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-common: build-common |
	apptree workflow package publish -d common_pkg --host ${HOST}
build-cache: |
	cd cache_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-cache: build-cache |
	apptree workflow package publish -d cache_pkg --host ${HOST}
build-logger: |
	cd logger_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-logger: build-logger |
	apptree workflow package publish -d logger_pkg --host ${HOST}
build-famis: |
	cd NetCoreSteps/Accruent.Famis.Steps && dotnet publish -o publish -c Release
publish-famis: build-famis |
	apptree workflow package publish -d NetCoreSteps/Accruent.Famis.Steps/publish --host ${HOST}
build-filesystem: |
	cd filesystem_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-filesystem: build-filesystem |
	apptree workflow package publish -d filesystem_pkg --host ${HOST}
build-webhook: |
	cd webhook_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-webhook: build-webhook |
	apptree workflow package publish -d webhook_pkg --host ${HOST}
build-oracle: |
	cd database/oracle_pkg && env CC=x86_64-w64-mingw32-gcc gox -osarch="windows/amd64" -ldflags="-s -w" -output "main_windows_amd64"
publish-oracle: build-oracle |
	apptree workflow package publish -d database/oracle_pkg --host ${HOST}
build-facility360:
	cd facility360_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-facility360: build-facility360 |
	apptree workflow package publish -d facility360_pkg --host ${HOST}
updatesdk: |
	cd filesystem_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/db_common && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/oracle_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/postgres_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd google_sheets_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd convert_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd common_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd logger_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd webhook_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd cache_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd facility360_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd script_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
publish-go: publish-common publish-convert publish-postgres publish-googlesheets publish-filesystem publish-logger publish-cache publish-facility360 publish-script publish-webhook

publish-dotnet: publish-famis

publish: publish-go publish-dotnet

# To add a new step package:
# 1. add "build-<PACKAGE>: |" command
# 2. add "publish-<PACKAGE>: build-PACKAGE |" command
# 3. add new build command to "build-go" command
# 4. add new publish command to "publish-go" command
# 5. Add a new line to the updatesdk command with your package name
