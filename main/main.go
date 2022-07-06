package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.ExpandEnv("DTS") )
	return db, err
}

func main() {

	//db
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PlanetScale!")
	//page
	fmt.Println("start")
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":3000", nil)
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "start")
}
