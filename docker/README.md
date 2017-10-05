You can use Docker to easily try out or test sendgrid-go.

# Quickstart
1. Install Docker in your machine.
2. Run `docker run -it -v /path/to/your-code:/data dhoeric/sendgrid-go`.
3. Run your code with `go run your-code.go`

# Examples
1. Copy and save in you local machine, e.g. `test.go`.
```
package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
)

func main() {
	api_key := "JUST_A_TEST_KEY"
	host := "http://localhost:4010"
	request := sendgrid.GetRequest(api_key, "/v3/api_keys", host)
	request.Method = "GET"

	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```
2. Run `docker run -it -v $(pwd):/data dhoeric/sendgrid-go`.
3. In container, run `go run test.go`

More sample codes on [Docker Examples](/docker/examples).


# Options
## Specifying specific versions
To use different version of sendgrid-go, mount them with `-v <host_dir>:<container_dir>` options. When you put either repository under `/mnt`, the container will automatically detect it and make the proper symlinks. You can edit these files from the host machine while the container is running.

For instance, to use sendgrid-go v3.0.6:
```
$ git clone https://github.com/sendgrid/sendgrid-go --branch v3.0.6
$ realpath sendgrid-go
/path/to/sendgrid-go
$ docker run -it -v /path/to/sendgrid-go:/mnt/sendgrid-go \
>                -v /path/to/your-code:/data \
>                dhoeric/sendgrid-go
$ (in container) go run your-code.go
```
