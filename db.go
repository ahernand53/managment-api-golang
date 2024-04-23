package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MySQL!")

	return &MySQLStorage{
		db: db,
	}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	// init db schema

	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createProjectsTable(); err != nil {
		return nil, err
	}

	if err := s.createTasksTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MySQLStorage) createProjectsTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	if err != nil {
		return err
	}

	return nil
}

func (s *MySQLStorage) createTasksTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			status ENUM('todo', 'in_progress', 'done') NOT NULL DEFAULT 'todo',
			projectId INT UNSIGNED NOT NULL,
			assignedTo INT UNSIGNED NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			FOREIGN KEY (projectID) REFERENCES projects(id),
			FOREIGN KEY (assignedTo) REFERENCES users(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	if err != nil {
		return err
	}

	return nil
}

func (s *MySQLStorage) createUsersTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	if err != nil {
		return err
	}

	return nil
}
