# go-toodledo

Go Client and library for Toodledo.

Status: WIP

## Usage

```sh
$ toodledo
Usage:
  toodledo [command]

Available Commands:
  auth        Manage authentication
  completion  Generate completion script
  config      Manage config
  context     Manage toodledo contexts
  folder      Manage toodledo folders
  goal        Manage toodledo goals
  help        Help about any command
  task        Manage toodledo tasks

Flags:
      --access_token string
      --config string         config file (default is $HOME/.cobra.yaml)
  -h, --help                  help for toodledo
  -v, --version               version for toodledo

Use "toodledo [command] --help" for more information about a command.
```

### Auth

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