FROM golang:latest
WORKDIR /build
RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v github.com/spf13/viper
RUN go get -d -v github.com/spf13/pflag
COPY k8sconfig.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o k8sconfig .

FROM alpine:latest
LABEL MAINTAINER=jcicilio@educert.com
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=0 /build/k8sconfig /app/
EXPOSE 80
CMD ["/app/k8sconfig"]