FROM golang:1.10.2 as builder

WORKDIR /go/src/github.com/alextanhongpin/go-github-scraper

COPY main.go .

# Additionally copy the internal folder, since this is not a package
COPY internal/ internal/

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/github.com/alextanhongpin/go-github-scraper/app .

# Metadata params
ARG VERSION
ARG BUILD_DATE
ARG VCS_URL
ARG VCS_REF
ARG NAME
ARG VENDOR

# Metadata
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name=$NAME \
      org.label-schema.description="Github Scraper with Golang" \
      org.label-schema.url="https://example.com" \
      org.label-schema.vcs-url=https://github.com/alextanhongpin/$VCS_URL \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vendor=$VENDOR \
      org.label-schema.version=$VERSION \
      org.label-schema.docker.schema-version="1.0" \
      org.label-schema.docker.cmd="docker run -d alextanhongpin/go-github-scraper"

CMD ["./app"]