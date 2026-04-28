package main

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFind(t *testing.T) {
	// sqlmock gives us a fake *sql.DB for testing.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening mock database: %s", err)
	}
	defer db.Close()

	createdAt := time.Date(2026, 4, 28, 10, 0, 0, 0, time.UTC)
	startDate := createdAt
	dueDate := createdAt.AddDate(0, 1, 0)

	// Build the row that the repository should receive from the database.
	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"start_date",
		"due_date",
		"status",
		"priority",
		"description",
		"created_at",
	}).AddRow(
		1,
		"Go ile Test Yazmayı Öğren",
		startDate,
		dueDate,
		1,
		1,
		"Test yazmak gerçekten çok önemlidir. Kodu kalitesini artırır. Hataları giderir.",
		createdAt,
	)

	// Expect the exact query used by Find.
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, title, start_date, due_date, status, priority, description, created_at
		 FROM tasks
		 WHERE id = ?`,
	)).
		WithArgs(1).
		WillReturnRows(rows)

	repo := NewRepository(db)

	task, err := repo.Find(1)
	if err != nil {
		t.Fatalf("failed to find task: %s", err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}

	if task.Title != "Go ile Test Yazmayı Öğren" {
		t.Errorf("unexpected title: %s", task.Title)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %s", err)
	}
}

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening mock database: %s", err)
	}
	defer db.Close()

	now := time.Date(2026, 4, 28, 10, 0, 0, 0, time.UTC)

	task := Task{
		Title:       "Learn React Native",
		StartDate:   now,
		DueDate:     now.AddDate(0, 1, 0),
		Status:      1,
		Priority:    1,
		Description: "Mobile development is now essential in modern software development.",
		CreatedAt:   now,
	}

	// Expect the INSERT query and the values passed to Exec.
	mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO tasks (title, start_date, due_date, status, priority, description, created_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
	)).
		WithArgs(
			task.Title,
			task.StartDate,
			task.DueDate,
			task.Status,
			task.Priority,
			task.Description,
			task.CreatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	lastID, err := repo.Add(task)
	if err != nil {
		t.Fatalf("failed to add task: %s", err)
	}

	if lastID != 1 {
		t.Errorf("expected last inserted id 1, got %d", lastID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %s", err)
	}
}
