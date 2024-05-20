package models

type CustomErr struct {
	Err        error
	Message    string
	StatusCode int
}
