FROM golang:alpine as dev
WORKDIR /
RUN apk add --update make
EXPOSE 5050
COPY . /
ENV CGO_ENABLED 0
RUN make build

FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
WORKDIR /run/
COPY --from=dev /build/application-service .
CMD ["./application-service"]
