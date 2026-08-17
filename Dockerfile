### Base ###

FROM golang:1.26-alpine AS base

RUN mkdir -p /opt/app
WORKDIR /opt/app
RUN apk add build-base

COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

COPY . .

### Build ###

FROM base AS build

RUN GOOS=linux go build -tags musl -ldflags "-w -s" \
    -o worker-service cmd/worker/main.go

### TEST-EXEC ###

FROM base AS test-exec

ARG _outputdir="/tmp/coverage"
ARG COVERAGE_EXCLUDE

RUN mkdir -p ${_outputdir} && \
    CGO_ENABLED=1 go test ./... -coverprofile=coverage.tmp -coverpkg=./... -covermode=atomic -p 1 && \
    grep -vE "${COVERAGE_EXCLUDE}" coverage.tmp > ${_outputdir}/coverage.out && \
    go tool cover -html=${_outputdir}/coverage.out -o ${_outputdir}/coverage.html

### Test ###

FROM scratch AS test

ARG _outputdir="/tmp/coverage"
COPY --from=test-exec ${_outputdir}/coverage.out /
COPY --from=test-exec ${_outputdir}/coverage.html /

### Final ###

FROM alpine:3.24.1 AS final

WORKDIR /app

COPY --from=build /opt/app/worker-service /app/worker-service

CMD ["/app/worker-service"]