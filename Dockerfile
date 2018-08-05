# build stage
FROM golang:alpine AS build-env
RUN apk add git
COPY . /go/src/test-coverage-plugin/
RUN cd /go/src/test-coverage-plugin/ && go get && go build -o app

# final stage
FROM alpine
RUN apk add ca-certificates
COPY --from=build-env /go/src/test-coverage-plugin/app /bin/
RUN chmod +x /bin/app
ENTRYPOINT /bin/app