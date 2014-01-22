package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Repository struct {
	Name string
}

type GithubJson struct {
	Repository Repository
	Ref        string
}

type Config struct {
	Hooks []Hook
}

type Hook struct {
	Repo   string
	Branch string
	Shell  string
}

func loadConfig(configFile *string) {
	var config Config
	configData, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(config.Hooks); i++ {
		addHandler(config.Hooks[i].Repo, config.Hooks[i].Branch, config.Hooks[i].Shell)
	}
}

func setLog(logFile *string) {
	log_handler, err := os.OpenFile(*logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("cannot write log")
	}
	log.SetOutput(log_handler)
	log.SetFlags(5)
}

func startWebserver() {
	log.Println("starting webserver")
	http.ListenAndServe(":"+*port, nil)
}

func addHandler(repo, branch, shell string) {
	uri := branch
	branch = "refs/heads/" + branch
	http.HandleFunc("/"+repo+"_"+uri, func(w http.ResponseWriter, r *http.Request) {
		payload := r.FormValue("payload")
		var data GithubJson
		err := json.Unmarshal([]byte(payload), &data)
		if err != nil {
			log.Println(err)
		}
		if data.Repository.Name == repo && data.Ref == branch {
			executeShell(shell)
		}
	})
}

func executeShell(shell string) {
	out, err := exec.Command(shell).CombinedOutput()
	if err != nil {
		log.Printf("An error occured during command execution:\n"+
			"Output: %s\n"+
			"Error: %s\n", out, err)
	} else {
		log.Printf("Shell output was: %s\n", out)
	}
}

var (
	port       = flag.String("port", "7654", "port to listen on")
	configFile = flag.String("config", "./config.json", "config")
	logFile    = flag.String("log", "./log", "log file")
)

func init() {
	flag.Parse()
}

func main() {
	setLog(logFile)
	loadConfig(configFile)
	startWebserver()
}
