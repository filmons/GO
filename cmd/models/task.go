package models
import "fmt"
import (
    "encoding/json"
    "io/ioutil"
    "os"
)


import (
    "time"
    "github.com/google/uuid"
)

// Task structure
type Task struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Status    Status    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

// Status type
type Status string

// Status constants
const (
    StatusDone    Status = "done"
    StatusPending Status = "pending"
)

// tasks slice to store tasks
var tasks []Task

// AddTask adds a new task to the tasks slice
func AddTask(title string, status Status) error {
    newTask := Task{
        ID:        generateID(),
        Title:     title,
        Status:    status,
        CreatedAt: time.Now(),
    }
    tasks = append(tasks, newTask)
    return nil
}

// GetAllTasks returns all tasks
func GetAllTasks() []Task {
    return tasks
}
// SaveTasks saves tasks to a JSON file
func SaveTasks(tasks []Task) error {
    data, err := json.Marshal(tasks)
    if err != nil {
        return err
    }
    return ioutil.WriteFile("tasks.json", data, 0644)
}

// LoadTasks loads tasks from a JSON file into the `tasks` slice.
func LoadTasks() error {
    data, err := ioutil.ReadFile("tasks.json")
    if err != nil {
        if os.IsNotExist(err) {
            tasks = []Task{} // Initialize to an empty slice if no file exists
            return nil
        }
        return err
    }

    err = json.Unmarshal(data, &tasks)
    if err != nil {
        return err
    }

    return nil
}


func GetTaskByID(id string) (Task, error) {
    for _, task := range tasks {
        if task.ID == id {
            return task, nil
        }
    }
    return Task{}, fmt.Errorf("task with ID %s not found", id)
}



func generateID() string {
    return uuid.New().String()
}


func UpdateTask(id, newTitle string, newStatus Status) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Title = newTitle
            tasks[i].Status = newStatus
            return SaveTasks(tasks) 
        }
    }
    return fmt.Errorf("task not found")
}


func DeleteTask(id string) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return SaveTasks(tasks) 
        }
    }
    return fmt.Errorf("task not found")
}
