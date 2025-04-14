# Get Migration Library
### Import
 go get -u -d github.com/golang-migrate/migrate/v4

 go mod tidy

 ### Create File 

 migrate create -ext sql -dir migrations -seq create_example_table

 ## Add extra line to migration file
 ### For up
 sql-migrate up
 ### For down
 sql-migrate down

