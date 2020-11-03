package main

import (
  "log"
  "net/http"
  "os"
  "path/filepath"
)

func main() {
  
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
          log.Fatal(err)
  }

  static := filepath.Join(dir, "/main.runfiles/__main__/src/server/static")

  fs := http.FileServer(http.Dir(static))
  http.Handle("/", fs)  

  log.Println("Listening on :8080...")
  errServer := http.ListenAndServe(":8080", nil)
  if errServer != nil {
    log.Fatal(errServer)
  }
}