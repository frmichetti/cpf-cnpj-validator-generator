# cpf-cnpj-validator-generator

## Start project
```
go run main/main.go
```

## Compile for Windows
```
GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe main/main.go
```

## Compile for Linux
### 64-bit
```
$ GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux main/main.go
```

### 32-bit
```
$ GOOS=linux GOARCH=386 go build -o bin/app-386-linux main/main.go
```