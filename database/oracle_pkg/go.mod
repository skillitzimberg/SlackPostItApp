module oracle_pkg

go 1.12

require (
	github.com/apptreesoftware/go-workflow v0.0.0-20190325161958-22837a95fe8a
	github.com/apptreesoftware/step_library/database/db_common v0.0.0
	gopkg.in/goracle.v2 v2.12.3
)

replace github.com/apptreesoftware/step_library/database/db_common => ../db_common
