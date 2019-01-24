FROM golang:1.9.1

ENV OAI_SPEC_URL="https://raw.githubusercontent.com/sendgrid/sendgrid-oai/master/oai_stoplight.json"
WORKDIR /root

# Install Prism
ADD https://raw.githubusercontent.com/stoplightio/prism/master/install.sh install.sh
RUN sh ./install.sh && \
    rm ./install.sh

# Get sendgrid-go
RUN git clone https://www.github.com/sendgrid/sendgrid-go /go/src/github.com/sendgrid-go && \
	cd /go/src/github.com/sendgrid-go && \
	go get

COPY entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh

COPY get-version.go get-version.go

ENTRYPOINT ["./entrypoint.sh"]
