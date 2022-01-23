Simple API sever and web application using golang. The focus of this project is to build uptime check service for any given domain .
Also, this project helps to understand how diffrent Golang configure work using diffrent pakage.

**Command to start the Server with hot reload option**
nodemon --exec go run server.go --signal SIGTERM

**Server Endpoint**
http://localhost:3088/app/v1/healthz

**Golang commands**
go mod tidy

**Docker Command**
This will “tag” the image uptime-server and build it. After it is built, we can run the image as a container.
docker build --tag uptime .
docker run --name uptime -p 3088:3088 uptime
docker run -d --name uptime-server -p 3088:3088 uptime-server
docker run -it -p 3088:3088 --rm --name uptime uptime
