// Package fs
package fs

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskStatus string

const (
	TODO     TaskStatus = "todo"
	PROGRESS TaskStatus = "in progress"
	DONE     TaskStatus = "done"
)

type UpdateAction string

const (
	Status      UpdateAction = "status"
	Description UpdateAction = "title"
)

type ListType string

const (
	ListDone       ListType = "done"
	ListTodo       ListType = "todo"
	ListInProgress ListType = "in-progress"
	ALL            ListType = "all"
)

type TaskStorage struct {
	file *os.File
}

func (t *TaskStorage) Add(task *Task) (taskID int64, error error) {

	payloadTasks := []Task{}
	if err := json.NewDecoder(t.file).Decode(&payloadTasks); err != nil && !errors.Is(err, io.EOF) {
		log.Println("error decoding file")
		return 0, err
	}

	newTaskID := 1

	if len(payloadTasks) > 0 {
		newTaskID = len(payloadTasks) + 1
	}

	task.Status = TODO
	task.ID = newTaskID
	task.CreatedAt = time.Now()

	payloadTasks = append(payloadTasks, *task)

	if _, err := t.file.Seek(0, io.SeekStart); err != nil {
		return 0, err
	}

	if err := t.file.Truncate(0); err != nil {
		return 0, err
	}

	if err := json.NewEncoder(t.file).Encode(&payloadTasks); err != nil {
		log.Println("error encoding task")
		return 0, err
	}

	return int64(newTaskID), nil
}

func (t *TaskStorage) List(listType ListType) []Task {

	payloadTasks := []Task{}
	if err := json.NewDecoder(t.file).Decode(&payloadTasks); err != nil {
		log.Println("error decoding file")
		return payloadTasks
	}

	filteredTasks := []Task{}
	switch listType {
	case ALL:
		return payloadTasks
	default:
		for id := range payloadTasks {
			switch listType {
			case ListTodo:
				if payloadTasks[id].Status == TODO {
					filteredTasks = append(filteredTasks, payloadTasks[id])
				}
			case ListInProgress:
				if payloadTasks[id].Status == PROGRESS {
					filteredTasks = append(filteredTasks, payloadTasks[id])
				}
			case ListDone:
				if payloadTasks[id].Status == DONE {
					filteredTasks = append(filteredTasks, payloadTasks[id])
				}
			}
		}
	}

	return filteredTasks
}

func (t *TaskStorage) Remove(taskID string) error {
	TaskID, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		log.Println("error parsing int")
		return err
	}

	payloadTasks := []Task{}
	if err := json.NewDecoder(t.file).Decode(&payloadTasks); err != nil && !errors.Is(err, io.EOF) {
		log.Println("error decoding file")
		return err
	}

	for index := range payloadTasks {
		if payloadTasks[index].ID == int(TaskID) {
			payloadTasks = append(payloadTasks[:index], payloadTasks[index+1:]...)
		}
	}

	if _, err := t.file.Seek(0, io.SeekStart); err != nil {
		return err
	}

	if err := t.file.Truncate(0); err != nil {
		return err
	}

	if err := json.NewEncoder(t.file).Encode(&payloadTasks); err != nil {
		log.Println("error encoding file")
		return err
	}

	return nil
}

func (t *TaskStorage) Update(taskID string, updateAction UpdateAction, task Task) (action UpdateAction, error error) {
	TaskID, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		log.Println("error parsing taskID")
		return action, err
	}
	payloadTasks := []Task{}
	if err := json.NewDecoder(t.file).Decode(&payloadTasks); err != nil && !errors.Is(err, io.EOF) {
		log.Println("error decoding file")
		return updateAction, err
	}

	for id := range payloadTasks {
		if payloadTasks[id].ID == int(TaskID) {
			if updateAction == Status {
				payloadTasks[id].Status = task.Status
			}
			if updateAction == Description {
				payloadTasks[id].Description = task.Description
			}
			payloadTasks[id].UpdatedAt = time.Now()
		}

	}

	if _, err := t.file.Seek(0, io.SeekStart); err != nil {
		return updateAction, err
	}

	if err := t.file.Truncate(0); err != nil {
		return updateAction, err
	}

	if err := json.NewEncoder(t.file).Encode(&payloadTasks); err != nil {
		log.Println("error encoding file")
		return updateAction, err
	}
	return updateAction, nil
}
