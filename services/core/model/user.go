package model

import (
	"time"

	"github.com/gocql/gocql"
)

type User struct {
	Id         gocql.UUID
	Username   string
	Password   string
	Email      string
	CreateDate time.Time
}
