package database

var db string

func init() {
	db = "mongodb"
}

func GetNameDatabase() string {
	return "your database is "  + db
}