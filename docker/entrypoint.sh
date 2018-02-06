#!/bin/bash 
clear

# check for + link mounted libraries:
if [ -d /mnt/sendgrid-go ]; then
  rm -rf /go/src/github.com/sendgrid/sendgrid-go
  `cd /go/src/github.com/sendgrid/ && ln -s /mnt/sendgrid-go && cd sendgrid-go && go get -v -d`
  echo "Linked mounted sendgrid-go's code to /go/src/github.com/sendgrid/sendgrid-go"
fi

SENDGRID_GO_VERSION=$(go run get-version.go)
echo "Welcome to sendgrid-go docker v${SENDGRID_GO_VERSION}."
echo 

if [ "$1" != "--no-mock" ]
then
  echo "Starting Prism in mock mode. Calls made to Prism will not actually send emails."
  echo "Disable this by running this container with --no-mock."
  prism run --mock --spec $OAI_SPEC_URL 2> /dev/null &
else
  echo "Starting Prism in live (--no-mock) mode. Calls made to Prism will send emails."
  prism run --spec $OAI_SPEC_URL 2> /dev/null  &
fi
echo "To use Prism, make API calls to localhost:4010."
echo "To stop Prism, run \"kill $!\" from the shell."
 
echo
echo "To test with sendgrid-go."
echo "Put your file in ./src, which mounted to /data in container, and run with \"go run [file].go\"."
echo 

cd /data
bash
