# FROM golang:latest as builder

# COPY . ./server


# ENV GO111MODULE=on
# RUN go mod tidy

# RUN CGO_ENABLED=0 GOOS=linux go build -o serverexec ./server.go

# From alpine:latest 
# WORKDIR /root/
# COPY --from=builder /server .
# CMD ["./serverexec"]

FROM golang:latest

# Install git.
# Git is required for fetching the dependencies.
RUN apt-get install git

WORKDIR $GOPATH/src/github.com/maxpa_qzrio/grpctut/server
COPY . .

# Fetch dependencies
RUN go mod tidy

# Build the binary. for grpc gateway
RUN go build .
EXPOSE 4040
#EXPOSE 9090
# Run the hello binary.
ENTRYPOINT ["./server"]
