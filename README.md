# gover 
Simple package with public version variables.

## Usage
```
go get github.com/datainq/gover
```

## Compilation
When compiling the code one should specify the version:

```
VERSION=v1.0.0
GIT_COMMIT=$(git rev-list -1 HEAD)
BUILD_DATE=$(date +"%s")

go build -ldflags "-X github.com/datainq/gover.GitCommit=$GIT_COMMIT -X github.com/datainq/gover.BuildTime=$BUILD_DATE -X github.com/datainq/gover.Version=$VERSION" main.go
```

You can import `github.com/datainq/gover` and use it. The most common use case
is to print version at binary startup, help and exit.