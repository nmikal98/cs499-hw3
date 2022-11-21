FROM golang:1.18.4

# set a build-time vriable 
ARG DB=memdb

# set a working directory for the app
WORKDIR /go/src/github.com/nmikal98/cs499-hw3/hotelapp

# copy all the files to the container
COPY . .

# build and install app
RUN go mod init github.com/nmikal98/cs499-hw3/hotelapp
RUN go mod tidy
RUN go install -tags ${DB} -ldflags="-s -w" ./cmd/...
