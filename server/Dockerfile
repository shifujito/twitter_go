# go version
FROM golang:1.16.3-alpine
# update git install
RUN apk add --update && apk add git && apk add curl && apk add postgresql-client
# app dir
RUN mkdir /go/src/app
# workgin dir
WORKDIR /go/src/app/server
# move work dir
COPY . /go/src/app
# install go module
# RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy && go install github.com/cosmtrek/air@v1.27.3
# run command
# ENTRYPOINT ["./bin/wait.sh"]
CMD ["air", "-c", ".air.toml"]
