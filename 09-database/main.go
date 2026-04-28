package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      int
	Priority    int
	Description string
	CreatedAt   time.Time
}

func dbConn() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345"
	dbName := "gotest"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	return db
}

type Repository interface {
	Find(id int) (Task, error)
	Add(task Task) (int64, error)
}

type repository struct {
	db *sql.DB
}

// NewRepository receives a database connection and stores it.
// This is important for testing because sqlmock passes a fake database here.
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Find searches one task by id.
// QueryRow is simpler than Query here because we expect only one row.
func (r repository) Find(id int) (Task, error) {
	var task Task

	err := r.db.QueryRow(
		`SELECT id, title, start_date, due_date, status, priority, description, created_at
		 FROM tasks
		 WHERE id = ?`,
		id,
	).Scan(
		&task.ID,
		&task.Title,
		&task.StartDate,
		&task.DueDate,
		&task.Status,
		&task.Priority,
		&task.Description,
		&task.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, errors.New("record not found")
		}
		return task, err
	}

	return task, nil
}

// Add inserts a new task and returns the generated id.
func (r repository) Add(task Task) (int64, error) {
	res, err := r.db.Exec(
		`INSERT INTO tasks (title, start_date, due_date, status, priority, description, created_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		task.Title,
		task.StartDate,
		task.DueDate,
		task.Status,
		task.Priority,
		task.Description,
		task.CreatedAt,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func main() {
	db := dbConn()
	myDB := NewRepository(db)

	task := Task{
		Title:       "React Native Öğren",
		StartDate:   time.Now(),
		DueDate:     time.Now().AddDate(0, 1, 0),
		Status:      1,
		Priority:    1,
		Description: "Mobil uygulama geliştirme artık günümüzün olmazsa olmazı.",
		CreatedAt:   time.Now(),
	}

	lastID, err := myDB.Add(task)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(lastID)
}
