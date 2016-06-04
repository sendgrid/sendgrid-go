[![Build Status](https://travis-ci.org/sendgrid/rest.svg?branch=master)](https://travis-ci.org/sendgrid/rest) [![GoDoc](https://godoc.org/github.com/sendgrid/rest?status.png)](http://godoc.org/github.com/sendgrid/rest)

**HTTP REST client, simplified for Go**

# Announcements

All updates to this library is documented in our [CHANGELOG](https://github.com/sendgrid/rest/blob/master/CHANGELOG.md).

# Installation

```bash
go get github.com/sendgrid/rest
```

# Quick Start

`GET /your/api/{param}/call`

```go
package main

import "github.com/sendgrid/rest"
import "fmt"

func main() {
	const host = "https://api.example.com"
	param := "myparam"
	endpoint := "/your/api/" + param + "/call"
	baseURL := host + endpoint
	method := rest.Get
	request := rest.Request{
		Method:  method,
		BaseURL: baseURL,
	}
	response, err := rest.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```

`POST /your/api/{param}/call` with headers, query parameters and a request body.

```go
package main

import "github.com/sendgrid/rest"
import "fmt"

func main() {
	const host = "https://api.example.com"
	param := "myparam"
	endpoint := "/your/api/" + param + "/call"
	baseURL := host + endpoint
	Headers := make(map[string]string)
	key := os.Getenv("API_KEY")
	Headers["Authorization"] = "Bearer " + key
	Headers["X-Test"] = "Test"
	var Body = []byte(`{"some": 0, "awesome": 1, "data": 3}`)
	queryParams := make(map[string]string)
	queryParams["hello"] = "0"
	queryParams["world"] = "1"
	method := rest.Post
	request = rest.Request{
		Method:      method,
		BaseURL:     baseURL,
		Headers:     Headers,
		QueryParams: queryParams,
		Body:        Body,
	}
	response, err := rest.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```

# Usage

- [Example Code](https://github.com/sendgrid/rest/tree/master/examples)

## Roadmap

If you are interested in the future direction of this project, please take a look at our [milestones](https://github.com/sendgrid/rest/milestones). We would love to hear your feedback.

## How to Contribute

We encourage contribution to our projects, please see our [CONTRIBUTING](https://github.com/sendgrid/rest/blob/master/CONTRIBUTING.md) guide for details.

Quick links:

- [Feature Request](https://github.com/sendgrid/rest/blob/master/CONTRIBUTING.md#feature_request)
- [Bug Reports](https://github.com/sendgrid/rest/blob/master/CONTRIBUTING.md#submit_a_bug_report)
- [Sign the CLA to Create a Pull Request](https://github.com/sendgrid/rest/blob/master/CONTRIBUTING.md#cla)
- [Improvements to the Codebase](https://github.com/sendgrid/rest/blob/master/CONTRIBUTING.md#improvements_to_the_codebase)

# About

rest is guided and supported by the SendGrid [Developer Experience Team](mailto:dx@sendgrid.com).

rest is maintained and funded by SendGrid, Inc. The names and logos for rest are trademarks of SendGrid, Inc.

![SendGrid Logo]
(https://assets3.sendgrid.com/mkt/assets/logos_brands/small/sglogo_2015_blue-9c87423c2ff2ff393ebce1ab3bd018a4.png)