# Build the kubediag binary
FROM golang:1.21

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY parent_workflow.go parent_workflow.go
COPY child_workflow.go child_workflow.go
COPY starter-1/ starter-1/
COPY starter-5/ starter-5/
COPY worker/ worker/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod vendor -o start-1 starter-1/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod vendor -o start-5 starter-5/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod vendor -o work worker/main.go

CMD ["start-1"]
