package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	BookTitle  string `json:"bookTitle"`
	BookAuthor string `json:"bookAuthor"`
	BookGenre  string `json:"bookGenre"`
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
	// type

	// insert, err := db.Query("INSERT INTO `BookSugesstion` (`bookTitle`, `bookAuthor`, `bookGenre`) VALUES('BOOOB', 'MisterBOOOB', 'comedy')")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer insert.Close()

	res, err := db.Query("SELECT `BookTitle`,`bookAuthor`, `bookGenre`  FROM `BookSugesstion`")
	if err != nil {
		fmt.Println(err)
		return
	}

	for res.Next() {
		err = res.Scan(&user.BookAuthor, &user.BookTitle, &user.BookGenre)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(fmt.Sprintf("Author: %s with ganre %d with%s", user.BookAuthor, user.BookGenre, user.BookTitle))

	}

	fmt.Println("Successfully connected to PlanetScale!")
	//page
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":3000", nil)
}

var user User

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintf("Author: %s with ganre %d with%s", user.BookAuthor, user.BookGenre, user.BookTitle)
}

func GetDatabase() (*sql.DB, error) {
	//read
	err := godotenv.Load() // читать .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	return db, err
}
