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
        {
          "Repo":"repo",
          "Branch":"master",
          "Shell":"niftyscript.sh"
        }
    ]
}
```

Now start the server with
  
    go run main.go --port 6578

and add a git-webhook for your.domain.com:6578/repo_master. Everytime you push to master, your script gets executed.

## What about safety?

Git webhooks use only 4 different ips for their webhooks. (207.97.227.253, 50.57.128.197, 108.171.174.178, 50.57.231.61) You can easily restrict access to your gohup server by using either a firewall or an equivalent nginx configuration.


## License

This Software is licensed under the MIT License.

Copyright (c) 2012 adeven GmbH, 
http://www.adeven.com

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
