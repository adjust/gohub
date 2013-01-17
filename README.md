gohub
=====

## What is gohub?

gohub is a little webserver written in go. He waits for webhook calls by github to run little shell commands.

## What is it good for?

Imagine you have your repo spread over several instances. You can use gohub to automate updating all your cloned repos.

## How to use

Just edit the config.json to your needs. A short example:
You want to track the status of your Repository "repo" and the branch master. If there is an update to this branch you want to execute your shell script "niftyscript.sh".

```json
{
    "Hooks":[
        "Repo":"repo",
        "Branch":"master",
        "Shell":"niftyscript.sh"
    ]
}
```

Now start the server with
  
    go run main.go --port 6578

and add a git-webhook for your.domain.com:6578/repo_master. Everytime you push to master, your script gets executed.

## What about safety?

Git webhooks use only 4 different ips for their webhooks. (207.97.227.253, 50.57.128.197, 108.171.174.178, 50.57.231.61) You can easily restrict access to your gohup server by using either a firewall or an equivalent nginx configuration.
