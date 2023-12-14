package model

import "time"

type Book struct {
	Id        int
	Title     string
	Pages     int
	CreatedAt time.Time
}
