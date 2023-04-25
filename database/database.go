package database

import "fmt"

var connection string

func init() {
	connection = "MySQL"
	fmt.Println(connection)
}

func GetDatabase() string {
	return connection
}
