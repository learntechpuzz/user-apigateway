FROM alpine:3.7

MAINTAINER hakkim.badhusha@wipro.com

WORKDIR /

WORKDIR /app

COPY . .

WORKDIR /app/cmd

EXPOSE 3000

ENTRYPOINT ["/app/user-apigatway", "--port", "3000"]