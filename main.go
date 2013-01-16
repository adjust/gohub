package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
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

func startWebserver() {
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
	out, err := exec.Command(shell).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Shell output was: %s\n", out)
}

var (
	port       = flag.String("port", "7654", "port to listen on")
	configFile = flag.String("config", "./config.json", "config")
)

func init() {
	flag.Parse()
}

func main() {
	loadConfig(configFile)
	startWebserver()
}
