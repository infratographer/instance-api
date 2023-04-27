FROM golang:1.20 as builder

# Create and change to the app directory.
WORKDIR /go/src/app

# Retrieve application dependencies using Go modules.
COPY go.* ./
RUN go mod download && go mod verify

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=1 GOOS=linux go build -mod=readonly -v -o /go/bin/instance-api main.go

FROM gcr.io/distroless/base

COPY --from=builder /go/bin/instance-api /

# Run the web service on container startup.
ENTRYPOINT ["/instance-api"]
CMD ["serve"]
EXPOSE 7608
