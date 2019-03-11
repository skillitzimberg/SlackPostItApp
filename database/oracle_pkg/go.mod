module oracle_pkg

go 1.12

require (
	github.com/apptreesoftware/go-workflow v0.0.0-20190311174159-e547a30e43cd
	github.com/apptreesoftware/step_library_go/database/db_common v0.0.0-20190311183454-3ead5ec1df38
	gopkg.in/goracle.v2 v2.12.3
)

replace github.com/apptreesoftware/step_library_go/database/db_common v0.0.0-20190311183454-3ead5ec1df38 => ../db_common
