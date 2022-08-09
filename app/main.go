package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/magiconair/properties"
)

func readCredsDirectory(dir string) Creds {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		creds := Creds{"unableToReadDir", "¯\\_(ツ)_/¯"}
		return creds
	}

	creds := Creds{}
	for _, f := range files {
		content, err := ioutil.ReadFile(f.Name())
		if err != nil {
			creds := Creds{"unableToReadFile", "¯\\_(ツ)_/¯"}
			return creds
		}
		if f.Name() == "username" {
			creds.Username = string(content)
		} else if f.Name() == "password" {
			creds.Password = string(content)
		}
	}
	return creds
}

func readCredsFile() Creds {
	credsFile := getEnv("CREDS_FILE", "")
	fileInfo, err := os.Stat(credsFile)
	if err != nil {
		creds := Creds{"fileNotFound", "¯\\_(ツ)_/¯"}
		return creds
	}

	if fileInfo.IsDir() {
		return readCredsDirectory(credsFile)
	}

	p, err := properties.LoadFile(credsFile, properties.UTF8)
	if err != nil {
		creds := Creds{"fileNotFound", "¯\\_(ツ)_/¯"}
		return creds
	}

	username := p.GetString("username", "unknown")
	password := p.GetString("password", "¯\\_(ツ)_/¯")
	creds := Creds{username, password}
	return creds
}

func readEnvVars() Creds {
	username := getEnv("DEMO_USERNAME", "unknown")
	password := getEnv("DEMO_PASSWORD", "¯\\_(ツ)_/¯")
	creds := Creds{username, password}
	return creds
}

func main() {
	http.HandleFunc("/", serveTemplate)

	fmt.Println("starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func serveTemplate(w http.ResponseWriter, req *http.Request) {
	index := filepath.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(index)
	fc := readCredsFile()
	ec := readEnvVars()
	p := getEnv("POD_NAME", "Flux Vault Demo")
	tmpl.Execute(w, Data{p, fc, ec})
}

type Data struct {
	Pod       string `json:"pod"`
	FileCreds Creds  `json:"fileCreds"`
	EnvCreds  Creds  `json:"envCreds"`
}

type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
