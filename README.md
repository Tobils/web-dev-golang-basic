# web-dev-golang-basic
development web menggunakan golang

```bash
mkdir golangweb
cd golangweb
go mod init golangweb
touch main.go
go run main.go

## auto reload
npm install -g nodemon
nodemon --exec go run main.go --signal SIGTERM
```