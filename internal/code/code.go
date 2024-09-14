package code

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s %s", time.Now(), os.Getenv("PORT"))
}

func Run() {
	PORT := ":" + os.Getenv("PORT")
	http.HandleFunc("/", greet)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
