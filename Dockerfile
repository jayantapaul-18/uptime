FROM golang:1.17.6-alpine3.15
# FROM arm32v7/golang:1.17
# RUN mkdir /user && \
#     echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
#     echo 'nobody:x:65534:' > /user/group

ENV GO111MODULE=on
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init && \
    rm -rf /var/cache/apk/* /tmp/*

WORKDIR /app/agent
COPY ./go.mod ./go.sum ./
RUN go mod download && rm go.mod go.sum
COPY . .
RUN go build -o ./out/agent .
RUN chmod +x ./out/agent 
EXPOSE 3088
CMD ["./out/agent"]