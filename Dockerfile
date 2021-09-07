ARG version=latest
FROM golang:$version

ENV GODEBUG 'x509ignoreCN=0'
ENV GO111MODULE 'auto'

COPY prism/prism/nginx/cert.crt /usr/local/share/ca-certificates/cert.crt
RUN update-ca-certificates

WORKDIR /go/src/github.com/sendgrid/sendgrid-go
COPY . .

RUN make install
