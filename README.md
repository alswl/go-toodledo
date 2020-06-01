# go-toodledo

## Usage

Grant: `open 'https://api.toodledo.com/3/account/authorize.php?response_type=code&client_id=donetoday&state=YourState&scope=basic%20tasks'`


## Install

```
go get github.com/leodido/go-urn
go get github.com/go-playground/universal-translator
go get github.com/go-playground/validator
go get github.com/Sirupsen/logrus
```


## Build


```
go build

export TOODLEDO_APP_CLIENT_ID=
export TOODLEDO_APP_SECRET=
export TOODLEDO_AUTH_TOKEN=
./go-toodledo

```


