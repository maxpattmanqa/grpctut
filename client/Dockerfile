FROM golang:latest

# Install git.
# Git is required for fetching the dependencies.
RUN apt-get install git

WORKDIR $GOPATH/src/github.com/maxpa_qzrio/grpctut/client
COPY . .

# Fetch dependencies
RUN go mod tidy

# Build the binary. for grpc gateway
RUN go build . 
EXPOSE 8080
#EXPOSE 9090
# Run the hello binary.
ENTRYPOINT ["./client"]
