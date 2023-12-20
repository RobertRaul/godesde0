package modelos

import (
	"time"
)

type User struct {
	Id       int
	Name     string
	CreateAt time.Time
	status   bool
}

func (usuario *User) AddUser(id int, name string, CreateAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreateAt = CreateAt
	usuario.status = status
}
