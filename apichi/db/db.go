package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/filmons/chiapi/models"

)

var DB *sql.DB

func InitDB(dataSourceName string) error {
    var err error
    DB, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        return err
    }
    return DB.Ping()
}


func GetUserByID(id string) (models.Tasks, error) {
    var task models.Tasks
    err := DB.QueryRow("SELECT id, title, created_at FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.CreatedAt)
    if err != nil {
        return models.Tasks{}, err
    }
    return task, nil
}
