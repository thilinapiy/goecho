# goecho
Go echo service with json responce

## Run the goecho service
> `$ go get`
> `$ go run server.go`

> `$ curl http://localhost:8080/echo/HelloWorld`
 
> {"Version":"v1","Msg":"HelloWorld","Time":"2017-01-20T11:09:46+05:30"} 


## Build the project
> `$ CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo .`


## Create the docker image
> `$ docker build -t goecho:latest .`


## Run the docker container
> `$ docker run -p 8080:8080 goecho`


