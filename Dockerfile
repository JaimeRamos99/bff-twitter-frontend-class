FROM golang:alpine

ENV GIN_MODE="test"
ENV PORT=8080

WORKDIR /bff-twitter-frontend-class

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY mock-data ./mock-data
COPY /tweets ./tweets
COPY main.go ./

RUN go build -o ./executable .

EXPOSE $PORT

ENTRYPOINT ["./executable"]