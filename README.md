# Gitlab Utility

## Configuration

You can create a GitLab personal access token here: https://git.psu.edu/-/profile/personal_access_tokens

~/.gitlabctl.yaml
```
gitlab:
  url: https://git.psu.edu/api/v4
  token: xxxxxxxxxxxxxxxxxxxxxxx
```

## Usage

```
$ gitlabctl 
Using config file: /home/crh5255/.gitlabctl.yaml
          _ __  __      __         __  __
   ____ _(_) /_/ /___ _/ /_  _____/ /_/ /
  / __ `/ / __/ / __ `/ __ \/ ___/ __/ / 
 / /_/ / / /_/ / /_/ / /_/ / /__/ /_/ /  
 \__, /_/\__/_/\__,_/_.___/\___/\__/_/   
/____/                                   
A tool used by software engineering to manage gitlab.

Usage:
  gitlabctl [flags]
  gitlabctl [command]

Available Commands:
  archive     archive gitlab projects
  clone       helper for cloning projects
  completion  Generates bash completion scripts
  help        Help about any command
  projects    list all gitlab projects
  tag         tag gitlab projects

Flags:
      --config string   config file (default is $HOME/.gitlabctl.yaml)
      --dry-run         log actions, don't perform gitlab api calls (for testing)
  -h, --help            help for gitlabctl

Use "gitlabctl [command] --help" for more information about a command.
```
