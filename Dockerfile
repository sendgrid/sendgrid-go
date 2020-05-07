ARG version=latest
FROM golang:$version

COPY prism/prism/nginx/cert.crt /usr/local/share/ca-certificates/cert.crt
RUN update-ca-certificates

WORKDIR /go/src/github.com/sendgrid/sendgrid-go
COPY . .

RUN make install
