package main

import (
	// "context"
	"net/http"
	// "os"
	"github.com/a-h/templ"
)

func main() {

  http.Handle("/", templ.Handler(hello()))

  http.ListenAndServe(":8090", nil)

  

  // component := hello("world")
  // component.Render(context.Background(), os.Stdout)
}
