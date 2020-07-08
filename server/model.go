/* Since we began to use MySql, what's in this go file should be no longer necessary */

package server

import (
	"errors"
	"time"
)

type Option struct {
	Title string
}

type ToDoList struct {
	events map[string]string
	count  int
}

func addEvent(toDo *ToDoList, event string) {
	toDo.events[time.Now().Format("2006-01-02 15:04:05")] = event
	toDo.count++
}

func GetEvent(a, b string) string {
	return a + " - " + b
}

func EditEvent(toDo *ToDoList, key, value string) error {
	_, ok := toDo.events[key]
	if !ok {
		return errors.New("error: event does not exist")
	}
	toDo.events[key] = value
	return nil
}

func DeleteEvent(toDo *ToDoList, key string) error {
	_, ok := toDo.events[key]
	if !ok {
		return errors.New("event does not exist")
	}
	delete(toDo.events, key)
	toDo.count--
	return nil
}
