package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

/*
 * Tag... - a very simple struct
 */
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// print
	fmt.Fprintf(w, "Go MySQL Tutorial\n")

	// Open up our database connection.
//	db, err := sql.Open("mysql", "root:Qwer1234@tcp(mariadb01.mynet:3306)/teszt")
	db, err := sql.Open("mysql", "root:Qwer1234@tcp(mysqldb.mygoapp_network:3306)/teszt")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, name FROM persons")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
		fmt.Fprintf(w, tag.Name+"\n")
	}

}

/* ****************************
 *             MAIN
 * ****************************
 */
func main() {

	// set the route for /
	http.HandleFunc("/", helloHandler)

	// run server
	fmt.Println("Starting server on :5050")
	if err := http.ListenAndServe(":5050", nil); err != nil {
		log.Fatal(err)
	}

}
