FROM golang:1.11.2-alpine

# Update packages and install git.
RUN apk update && apk add git

# Move to project root
WORKDIR /go/src/github.com/george-kj/krsc

# Install dependencies
COPY Gopkg.* ./
RUN go get -u github.com/golang/dep/...
RUN dep ensure -vendor-only

# Other non-vendored files
COPY main.go config.toml ./
COPY app app
COPY conf conf
COPY domain domain
COPY handler handler
COPY repo repo

RUN go install .
CMD ["krsc"]