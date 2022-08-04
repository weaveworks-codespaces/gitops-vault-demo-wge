package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/magiconair/properties"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from flux-vault-demo!!!")
}

func readCredsFile(w http.ResponseWriter, req *http.Request) {
	credsFile := getEnv("CREDS_FILE", "./dummy_creds")
	p, err := properties.LoadFile(credsFile, properties.UTF8)
	if err != nil {
		fmt.Fprintf(w, "Oops!! Cannot find file \"%v\"", credsFile)
		return
	}

	username := p.GetString("username", "unknown")
	password := p.GetString("password", "myNotSoSecretPassword")

	fmt.Fprintf(w, "Hello %v!! Your password is %v", username, password)
}

func readEnvVars(w http.ResponseWriter, req *http.Request) {
	username := getEnv("DEMO_USERNAME", "unknown")
	password := getEnv("DEMO_PASSWORD", "myNotSoSecretPassword")

	fmt.Fprintf(w, "Hello %v!! Your password is %v", username, password)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/file", readCredsFile)
	http.HandleFunc("/env", readEnvVars)

	fmt.Println("starting server")
	http.ListenAndServe(":8080", nil)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
