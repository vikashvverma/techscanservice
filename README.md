# techscanservice
Preview: http://techscan.programminggeek.in

## Prerequisite
- Install [Golang](https://golang.org/), [Glide](https://github.com/Masterminds/glide) and [PostgreSQL](https://www.postgresql.org/)

## Build & development
- Run `go env` and make sure your `GOPATH` is set. If not then set GOPATH, e.g. ` export GOPATH="/foo/bar/baz"`
- Run `mkdir -p $GOPATH/src/github.com/vikashvverma` or `go get github.com/vikashvverma/techscanservice`
- If you created directory in previous step instead of `go get` then move `techscanservice`(extracted from submitted zip) to `$GOPATH/src/github.com/vikashvverma` directory.
- Change working directory to `$GOPATH/src/github.com/vikashvverma/techscanservice` and follow subsequent steps.
- Run `glide install` to install dependencies
- Run `go build -o out/build/techscanservice ./apps/techscan/main.go` to build binary
- Update the config present at `./config/dev.env`. This is important and app won't start unless configs are valid.
- Leave the config as it is if you want to connect to default cloud instance.
- Run `./dev.sh` to start techscanservice.

