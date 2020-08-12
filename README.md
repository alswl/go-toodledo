# go-toodledo

## Usage

Grant: `open 'https://api.toodledo.com/3/account/authorize.php?response_type=code&client_id=donetoday&state=YourState&scope=basic%20tasks'`


## Install

```
go mod download
```


## Build


```
go build

export TOODLEDO_APP_CLIENT_ID=
export TOODLEDO_APP_SECRET=
export TOODLEDO_AUTH_TOKEN=
./go-toodledo

```


