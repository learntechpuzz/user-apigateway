# User API Gateway

user-apigateway is a Golang RESTful services project, using Echo framework. 

## Getting started

1. Run the app using:

```bash

go run app/cmd/main.go

```

The application runs as an HTTP server at port 3000 (config.dev.yaml server.port). 

It provides the following RESTful endpoints:

* `POST v1/api/users`: create new user
* `GET v1/api/users`: get all users

