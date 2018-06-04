FROM golang:1.10.2 as builder

WORKDIR /go/src

# Copy vendor folder to avoid re-downloading the whole packages again
COPY vendor/ .

WORKDIR /go/src/github.com/alextanhongpin/go-github-scraper

COPY main.go go.mod .

# Additionally copy the internal folder, since this is not a package
# and will not be fetched from Github
COPY internal/ internal/

# COPY vendor/ vendor/

# RUN go version
# RUN go version && go get -u -v golang.org/x/vgo

# RUN vgo get ./...

# Uncomment to pull packages every time
# RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

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