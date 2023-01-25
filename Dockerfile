FROM golang:1.19 as builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o initSetupScript main.go


FROM alpine:latest
WORKDIR /
COPY --from=builder /workspace/initSetupScript .
EXPOSE 8082
ENTRYPOINT ["/initSetupScript"]