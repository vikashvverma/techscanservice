# techscanservice
Preview: http://techscan.programminggeek.in

## Prerequisite
- Install [Golang](https://golang.org/), [Glide](https://github.com/Masterminds/glide) and [PostgreSQL](https://www.postgresql.org/)

## Build & development
- Run `glide install` to install dependencies
- Run `go build -o out/build/techscanservice ./apps/techscan/main.go` to build binary
- Update the config present at `./config/dev.env`. This is important and app won't start unless configs are valid.
- Run `./dev.sh` to start techscanservice.

