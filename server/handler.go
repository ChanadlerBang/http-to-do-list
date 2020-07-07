package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	ToDo        ToDoList
)

func init()  {
	ToDo = ToDoList{make(map[string]string), 0}
}

func Add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/add.html")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}
		addEvent(&ToDo, r.Form["event"][0])
		http.Redirect(w, r, "/view", http.StatusFound)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/edit.html")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}
		if strings.Contains(r.URL.RawQuery, "event") {
			key := r.URL.RawQuery[strings.Index(r.URL.RawQuery, "event") + len("event="):]
			key = getTime(key)
			fmt.Println(strings.Join(r.Form["event"], " "))
			err := EditEvent(&ToDo, key, r.Form["event"][0])
			if err != nil {
				log.Fatal("Edit Event: ", err)
			}
			http.Redirect(w, r, "/view", http.StatusFound)
		}
	}
}

func View(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	err := r.ParseForm()
	if err != nil {
		log.Fatal("ParseForm: ", err)
	}
	if r.Method == "GET" {
		if strings.Contains(r.URL.RawQuery, "event") {
			str := r.URL.RawQuery[strings.Index(r.URL.RawQuery, "event") + len("event="):]
			str = getTime(str)
			fmt.Println(str)
			err := DeleteEvent(&ToDo, str)
			if err != nil {
				log.Fatal("Delete Event: ", err)
			} else {
				http.Redirect(w, r, "/view", http.StatusFound)
			}
		} else {
			t, err := template.ParseFiles("views/view.html")
			if err != nil {
				log.Fatal("ParseFiles: ", err)
			}
			t.Execute(w, ToDo.events)
		}
	} else {
		// TO Do HERE
	}
}

func getTime(str string) string {
	return strings.Replace(str, "%20", " ", -1)
}