set GOOS=linux&& set GOARCH=amd64&& go build -o taskservice server/taskservice/cmd/main.go
docker build . -t taskservice:latest
docker tag taskservice:latest registry.heroku.com/go-server/taskservice:2
docker push registry.heroku.com/go-server/taskservice:2