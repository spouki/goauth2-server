package main

//import "fmt"
import "net/http"
import "log"
//import "github.com/gorilla/mux"
import "github.com/spouki/goauth2/router"

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080",router))
}

// EOF
