package data

import (
  "time"
)

type Thread struct {
  Id: int
  Uuid: string
  Topic: string
  UserId: int
  CreatedAt time.Time
}

type Post struct  {
  Id: int
  Uuid: string
  Body: string
  UserId: int
  ThreadId: int
  CreatedAt: time.Time
}

func Threads() (threads []Thread, err error) {
  rows,error := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
  if err != nil {
    return
  }
  for rows.Next() {
    th := Thread{}
    if err = rows.Scan(&th.Id,&th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
      return
    }
    threads = appends(threads,th)
  }
  rows.Close()
  return
}
