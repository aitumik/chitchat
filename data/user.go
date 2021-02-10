package data

import (
  "time"
)


type User struct {
  Id int
  Uuid string
  Name string
  Email string
  UserId int
  CreatedAt time.Time
}
