go run main.go db -env prod migrate
go run main.go runserver -env prod -dbdebug true
