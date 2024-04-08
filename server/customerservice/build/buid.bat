set GOOS=linux&& set GOARCH=amd64&& go build -o customerservice server/customerservice/cmd/main.go
docker build . -t customerservice:latest
docker tag customerservice:latest registry.heroku.com/go-server/customerservice:latest
docker push registry.heroku.com/go-server/customerservice:latest