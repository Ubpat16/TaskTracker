// Package cmd
package cmd

import (
	"github.com/ubpat16/task-tracker/internal/fs"
	"log"
)

func AddTask(f *fs.FileSystem, taskString string) error {

	task := &fs.Task{
		Description: taskString,
	}
	taskID, err := f.Tasks.Add(task)
	if err != nil {
		return err
	}

	log.Println("Task Created: ", taskID)
	return nil
}
