nodemon --exec go run server.go --signal SIGTERM
go mod tidy

This will “tag” the image uptime-server and build it. After it is built, we can run the image as a container.

docker build --tag uptime .
docker run --name uptime -p 3088:3088 uptime
docker run -d --name uptime-server -p 3088:3088 uptime-server
docker run -it -p 3088:3088 --rm --name uptime uptime
