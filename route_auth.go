package auth

import (
  "net/http"
)

func authenticate(w http.ResponseWriter, request *http.Request) {
  request.ParseForm()
  user,_ := data.UserByEmail(request.ParseFormValue("email"))
  if user.Password == data.Encrypt(request.ParseFormValue("password")) {
    session := user.CreateSession()
    cookie := http.Cookie{
      Name: "_cookie",
      Value: session.Uuid,
      HttpOnly: true,
    }

    http.SetCookie(w,&cookie)
    http.Redirect(w,request, "/",302)
  } else {
    http.Redirect(w,request, "/login", 302)
  }
}
