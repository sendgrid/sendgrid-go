You can use Docker to easily try out or test sendgrid-go.

# Quickstart
1. Install Docker and docker-compose in your machine
2. Change into the docker directory, `cd docker`
3. run `docker-compose build`
4. run `docker-compose up`
5. Put the script you want to test in `./src`
6. Run `docker-compose run mock-server`, this will open a terminal on the container
7. Run your code with `go run your-code.go` once inside container

# Examples
You can use this as an example to show how this works. 

1. Copy and save in `./src/test.go`, e.g. `test.go`.
```
package main

import (
  "os"
  "fmt"
  "github.com/sendgrid/sendgrid-go"
  "log"
)

func main() {
  apiKey := os.Getenv("SENDGRID_API_KEY")
  host := "http://localhost:4010"
  request := sendgrid.GetRequest(apiKey, "/v3/api_keys", host)
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
2. Run `docker-compose run mock-server`.
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
