# go-toodledo

Go library and Cli for Toodledo.

Status: WIP

## Usage

```sh
> toodledo help
Usage:
  toodledo [command]

Available Commands:
  auth         Manage authentication
  completion   Generate completion script
  config       Manage config
  context      Manage toodledo contexts
  folder       Manage toodledo folders
  goal         Manage toodledo goals
  help         Help about any command
  saved-search Manage toodledo saved search
  task         Manage toodledo tasks

Flags:
      --access_token string
      --config string         config file (default is $HOME/.toodledo.yaml)
  -h, --help                  help for toodledo
  -v, --version               version for toodledo

Use "toodledo [command] --help" for more information about a command.

> toodledo task --help
Manage toodledo tasks

Usage:
  toodledo task [command]

Available Commands:
  complete
  create      Create a task
  delete
  edit
  list
  uncomplete
  view

Flags:
  -h, --help   help for task

Global Flags:
      --access_token string
      --config string         config file (default is $HOME/.toodledo.yaml)

Use "toodledo task [command] --help" for more information about a command.

> toodledo task list --context home --status nextaction
         # │ [X] │ TITLE            │     STATUS │ CONTEXT │ PRIORITY │ FOLDER  │ GOAL │        DUE │ REPEAT      │ LENGTH │ TIMER
───────────┼─────┼──────────────────┼────────────┼─────────┼──────────┼─────────┼──────┼────────────┼─────────────┼────────┼───────
 327077755 │ [ ] │ next-action item │ NextAction │ home    │     High │ to-0128 │ b    │ 1645704000 │ FREQ=WEEKLY │     60 │  1200
 327078471 │ [ ] │ abc2x            │ NextAction │ home    │     High │         │ c    │          0 │             │      0 │     0
```

### Login

```sh
# login
toodledo auth login
# follow steps, open link your browser
toodledo auth login YOUR-CODE
# verify
toodledo auth me
```

## Build

```sh
git clone https://github.com/alswl/go-toodledo.git
make
./bin/toodledo --help
```