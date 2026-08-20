package cmd

import (
	"github.com/ubpat16/task-tracker/internal/fs"
	"log"
)

func Update(f *fs.FileSystem, taskID string, title string) error {
	task := fs.Task{Description: title}
	action, err := f.Tasks.Update(taskID, fs.Description, task)
	if err != nil {
		return err
	}

	log.Printf("Updated %s Successfully", action)
	return nil
}

func UpdateStatus(f *fs.FileSystem, taskID string, status fs.TaskStatus) error {
	task := fs.Task{Status: status}
	action, err := f.Tasks.Update(taskID, fs.Status, task)
	if err != nil {
		return err
	}

	log.Printf("Updated task ID: %s's %s Successfully", taskID, action)
	return nil
}
