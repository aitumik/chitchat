package main

import (
  "net/http"
  "html/template"
  "github.com/aitumik/chitchat"
  "fmt"
)


func main() {
  mux := http.NewServeMux()

  files := http.FileServer(http.Dir("/public"))

  //handle the static files
  mux.Handle("/static/",http.StripPrefix("/static",files))

  mux.HandleFunc("/",index)

  server := &http.Server{
    Addr: "0.0.0.0:5000",
    Handler: mux,
  }

  server.ListenAndServe()
}


func index(writer http.ResponseWriter, request *http.Request) {
  threads,err := data.Threads()
  generateHTML(writer,nil,"layout","index")
  /**
  threads,err := data.Threads(){
    _,err := session(w,r)
    if err == nil {
      generateHTML(writer,threads,"layout" "public.navbar.html","index")
    } else {
      generateHTML(writer,threads, "layout","private.navbar.html","index")
    }
  }
  */
}

func generateHTML(writer http.ResponseWriter, data interface{}, fn ...string) {
  var files []string

  for _, file := range fn {
    files = append(files, fmt.Sprintf("templates/%s.html",file))
  }

  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer,"layout", data)
}
