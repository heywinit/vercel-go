package handler

import (
	"fmt"
	"net/http"
	"os"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	programName := os.Getenv("PROGRAM_NAME")
	fmt.Fprintf(w, "<h1>Hello from Go! %s</h1>", programName)
}
