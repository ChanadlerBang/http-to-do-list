package server

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ToDo ToDoList
	db   *sql.DB
	err  error
)

//-> create table userinfo(
//-> id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
//-> event TEXT NOT NULL,
//-> created DATETIME NOT NULL DEFAULT NOW()
//-> )

func init() {
	db, err = sql.Open("mysql", "root:123456@/to_do_list?charset=utf8")
	checkErr(err, "Open database: ")
}

func Add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/add.html")
		log.Println(t.Execute(w, nil))
	} else {
		err = r.ParseForm()
		checkErr(err, "ParseForm: ")

		stmt, err := db.Prepare("INSERT INTO userinfo (event) VALUES (?)")
		checkErr(err, "Database Prepare: ")

		_, err = stmt.Exec(strings.Join(r.Form["event"], ""))
		checkErr(err, "Database Exec: ")

		http.Redirect(w, r, "/view", http.StatusFound)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/edit.html")
		log.Println(t.Execute(w, nil))
	} else {
		err = r.ParseForm()
		checkErr(err, "ParseForm: ")

		if strings.Contains(r.URL.RawQuery, "id") {
			id := r.URL.RawQuery[strings.Index(r.URL.RawQuery, "id")+len("id="):]

			stmt, err := db.Prepare("UPDATE userinfo SET event=? WHERE id=?")
			checkErr(err, "Database Prepare: ")

			_, err = stmt.Exec(strings.Join(r.Form["event"], ""), id)
			checkErr(err, "Database Exec: ")

			http.Redirect(w, r, "/view", http.StatusFound)
		}
	}
}

func View(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	err = r.ParseForm()
	checkErr(err, "ParseForm: ")
	if r.Method == "GET" {
		if strings.Contains(r.URL.RawQuery, "id") {
			id := r.URL.RawQuery[strings.Index(r.URL.RawQuery, "id")+len("id="):]
			fmt.Println(id)

			stmt, err := db.Prepare("DELETE FROM userinfo WHERE id=?")
			checkErr(err, "Database Prepare: ")

			_, err = stmt.Exec(id)
			checkErr(err, "Database Exec: ")

			http.Redirect(w, r, "/view", http.StatusFound)

		} else {
			t, err := template.ParseFiles("views/view.html")
			checkErr(err, "ParseFiles: ")

			rows, err := db.Query("SELECT id, event FROM userinfo")
			checkErr(err, "Database Query: ")

			events := make(map[int]string)
			for rows.Next() {
				var id int
				var event string
				err = rows.Scan(&id, &event)
				checkErr(err, "Database Scan: ")
				events[id] = event
				fmt.Println(strconv.Itoa(id) + event)
			}

			t.Execute(w, events)
		}
	} else {
		// TO Do HERE
	}
}

func getTime(str string) string {
	return strings.Replace(str, "%20", " ", -1)
}

func checkErr(er error, info string) {
	if er != nil {
		log.Fatal(info, er)
	}
}
