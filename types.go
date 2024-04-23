package main

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectID    int64     `json:"project_id"`
	AssignedToId int64     `json:"assigned_to_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
