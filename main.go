package main

import (
	"log"
	"os"

	"github.com/ubpat16/task-tracker/cmd"
	"github.com/ubpat16/task-tracker/internal/fs"
)

type RouteAction string

const (
	ADD        RouteAction = "add"
	DELETE     RouteAction = "delete"
	INPROGRESS RouteAction = "mark-in-progress"
	TASKDONE   RouteAction = "mark-done"
	UPDATE     RouteAction = "update"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Println("No argument provided")
		return
	}
	route := args[0]

	var fileName = "task.json"
	file, err := fs.New(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	storage := fs.JSONStorage(file)

	switch route {
	case "add":
		taskString := args[1]
		if err := cmd.AddTask(&storage, taskString); err != nil {
			log.Fatal(err)
		}
	case "delete":
		taskID := args[1]
		if err := cmd.RemoveTask(&storage, taskID); err != nil {
			log.Fatal(err)
		}
	case "mark-in-progress":
		taskID := args[1]
		status := fs.PROGRESS
		if err := cmd.UpdateStatus(&storage, taskID, status); err != nil {
			log.Fatal(err)
		}
	case "mark-done":
		taskID := args[1]
		status := fs.DONE
		if err := cmd.UpdateStatus(&storage, taskID, status); err != nil {
			log.Fatal(err)
		}
	case "update":
		taskID := args[1]
		title := args[2]
		if err := cmd.Update(&storage, taskID, title); err != nil {
			log.Fatal(err)
		}
	default:
		if len(args) == 1 {
			cmd.ListTask(&storage, fs.ALL)
		} else {
			listType := args[1]
			switch listType {
			case "done":
				cmd.ListTask(&storage, fs.ListDone)
			case "todo":
				cmd.ListTask(&storage, fs.ListTodo)
			case "in-progress":
				cmd.ListTask(&storage, fs.ListInProgress)
			}
		}

	}
	defer file.Close()
}
