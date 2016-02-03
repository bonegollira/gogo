package main

import (
  "log"
  "net/http"
  "text/template"
  "path/filepath"
  "sync"
  "flag"
  "os"
  "github.com/stretchr/gomniauth"
  //"github.com/stretchr/gomniauth/providers/facebook"
  "github.com/stretchr/gomniauth/providers/google"
  //"github.com/stretchr/gomniauth/providers/github"
  "github.com/stretchr/objx"
  "./trace"
)

func main () {
  addr := flag.String("addr", ":8080", "アプリケーションのアドレス")
  flag.Parse()

  gomniauth.SetSecurityKey("GoChat")
  gomniauth.WithProviders(
    google.New(googleClientID, googleSecret, "http://localhost:8080/auth/callback/google"),
  )

  r := newRoom()
  r.tracer = trace.New(os.Stdout)

  // http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {w.Write([]byte())})
  // http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("/assetsへのパス"))))
  // *templateHandler型（templateHandlerのポインタ型）にServeHTTPが実装されている
  // のでtemplateHandlerのアドレスを渡してポインタである*templateHandler型を渡す
  http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
  http.Handle("/login", &templateHandler{filename: "login.html"})
  http.HandleFunc("/auth/", loginHandler)
  http.Handle("/room", r)

  go r.run()

  log.Println("Webサーバを開始します。ポート: ", *addr)

  // start web server
  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}

type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

func (t *templateHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  t.once.Do(func () {
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  data := map[string]interface{}{
    "Host": r.Host,
  }
  if authCookie, err := r.Cookie("auth"); err == nil {
    data["UserData"] = objx.MustFromBase64(authCookie.Value)
  }
  t.templ.Execute(w, data)
}
