FROM golang:1.18-alpine as buildStage

WORKDIR /app 

#copy from current working directory to working directory in docer image 
COPY . .

#bulding our app to a single binary executable file  specify directory where main entry point is. in this case, ./ or main.go 
RUN GOOS=linux CGO_ENABLED=0 go build -o goTrackit .

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
         


##alpine makes the docker image smaller , the smaller the better 😀
FROM alpine:latest 
WORKDIR /app

COPY --from=buildStage /app/goTrackit .

COPY --from=buildStage /app/migrate.linux-amd64 /app/migrate
# COPY app.env .
COPY .env .
COPY db/migration ./migration
COPY start.sh .
COPY wait-for.sh .
# COPY --from=buildStage /app/goTrackit /app

EXPOSE  8080
#command to executable that was built earlier
CMD ["/app/goTrackit"]
ENTRYPOINT [ "/app/start.sh" ]