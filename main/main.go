package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "xhgfq4ix2kix:pscale_pw_u2xdcZYq5DLSHHBFwwk89g83NmsqeWsiRcQzjFQ9kRU@tcp(av51do8uu7vi.us-east-1.psdb.cloud)/dbtest?tls=true")
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
