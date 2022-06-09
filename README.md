# go-toodledo

Go library and Cli for Toodledo.

Status: Under Development

## Usage

```sh
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
      --config string         config file (default is $HOME/.config/toodledo/conf.yaml)
  -h, --help                  help for toodledo

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
  editor      
  list        
  uncomplete  
  view        

Flags:
  -h, --help   help for task

Global Flags:
      --access_token string   
      --config string         config file (default is $HOME/config/toodledo/conf.yaml)

Use "toodledo task [command] --help" for more information about a command.

> toodledo task list --context home --status nextaction
         # │ [X] │ TITLE            │     STATUS │ CONTEXT │ PRIORITY │ FOLDER  │ GOAL │ DUE        │ REPEAT      │ LENGTH │ TIMER
───────────┼─────┼──────────────────┼────────────┼─────────┼──────────┼─────────┼──────┼────────────┼─────────────┼────────┼───────
 327077755 │ [ ] │ next-action item │ NextAction │ home    │     High │ to-0128 │ b    │ 2022-02-24 │ FREQ=WEEKLY │ 1m0s   │ 20m0s
 327078471 │ [ ] │ abc2x            │ NextAction │ home    │     High │         │ c    │            │             │        │
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

### Tasks

```shell
# list
> toodledo task list --context home --goal goal-b
INFO[0002] Syncing tasks
         # │ [X] │ TITLE    │ STATUS │ CONTEXT │ PRIORITY │ FOLDER │ GOAL   │ DUE │ REPEAT │ LENGTH │ TIMER
───────────┼─────┼──────────┼────────┼─────────┼──────────┼────────┼────────┼─────┼────────┼────────┼───────
 334313679 │ [ ] │ cooking3 │   None │ home    │   Medium │        │ goal-b │     │        │        │

# complete
> toodledo task complete 323245685
         # │ [X] │ TITLE  │ STATUS │ CONTEXT │ PRIORITY │ FOLDER  │ GOAL    │ DUE │ REPEAT │ LENGTH │ TIMER
───────────┼─────┼────────┼────────┼─────────┼──────────┼─────────┼─────────┼─────┼────────┼────────┼───────
 323245685 │ [X] │ test-c │   None │ Not Set │      Low │ Not Set │ Not Set │   0 │        │      0 │     0

# edit
> toodledo task edit --title cooking4-8 --context a --status nextaction 334313701
         # │ [X] │ TITLE      │     STATUS │ CONTEXT │ PRIORITY │ FOLDER │ GOAL   │ DUE │ REPEAT │ LENGTH │ TIMER
───────────┼─────┼────────────┼────────────┼─────────┼──────────┼────────┼────────┼─────┼────────┼────────┼───────
 334313701 │ [ ] │ cooking4-8 │ NextAction │ a       │   Medium │ a      │ goal-b │     │        │        │
```

## Build

```sh
> git clone https://github.com/alswl/go-toodledo.git
> make
> ./bin/toodledo --help
```