package cmd

import (
	"log"

	"github.com/ubpat16/task-tracker/internal/fs"
)

func RemoveTask(f *fs.FileSystem, taskID string) error {
	err := f.Tasks.Remove(taskID)
	if err != nil {
		log.Println("Error removing task")
		return err
	}
	log.Println("Task Removed")
	return nil
}
