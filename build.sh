version=`git describe --tags HEAD`
env GOOS=linux GOARCH=amd64 go build \
      -ldflags "-X github.com/broodbear/tools/cmd.Version=$version" \
      -o bin/bear-linux-amd64 \
      ./main.go
