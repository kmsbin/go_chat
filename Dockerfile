FROM golang:1.16.5

RUN mkdir /app
ADD . /app
WORKDIR /app

# COPY the source code as the last step
COPY go.mod . 
COPY go.sum .
COPY . .
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/hello

RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]