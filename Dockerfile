FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
ADD ./app /app/
# ADD ./zoneinfo.zip /usr/local/go/lib/time/
#ADD ./demo.html /app/
ENTRYPOINT ["./app"]

#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app