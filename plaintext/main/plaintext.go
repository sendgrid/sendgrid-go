package main

import (
    "os"
    "github.com/sendgrid/sendgrid-go/plaintext"
)

func main() {
    plaintext.FromHTML(os.Stdin, os.Stdout)
}
