package handlers


import (
	"fmt"
	"net/http"
	// "encoding/json"
	// "io"
	// "io/ioutil"
	// "strconv"
	// "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Registration(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Registration\n")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login\n")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logout\n")
}

func PasswordRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Password request\n")
}
