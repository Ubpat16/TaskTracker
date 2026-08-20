// Package fs
package fs

import (
	"os"
)

type FileSystem struct {
	Tasks interface {
		Add(task *Task) (TaskID int64, error error)
		List(listType ListType) (task []Task)
		Remove(TaskID string) (error error)
		Update(TaskID string, updateAction UpdateAction, task Task) (action UpdateAction, error error)
	}
}

func JSONStorage(file *os.File) FileSystem {
	return FileSystem{
		Tasks: &TaskStorage{file},
	}
}

func New(fileName string) (file *os.File, error error) {
	return os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
}
