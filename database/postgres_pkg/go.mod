module postgres_pkg

go 1.12

require (
	github.com/apptreesoftware/go-workflow v0.0.0-20190313215837-b2bdb87898e5
	github.com/apptreesoftware/step_library/database/db_common v0.0.0
	github.com/lib/pq v1.0.0
)

replace github.com/apptreesoftware/step_library/database/db_common => ../db_common
