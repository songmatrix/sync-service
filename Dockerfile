FROM golang:1.19 as build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o /bin/app ./cmd/server/main.go

FROM alpine:3.16
WORKDIR /go/app
RUN adduser --disabled-password --gecos "" --uid 1000 default
USER default
COPY --from=build --chown=default /bin/app /bin/app
ENTRYPOINT [ "/bin/app" ]
