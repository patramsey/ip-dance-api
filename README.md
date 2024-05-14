# ip.dance api golang library #

<img src="https://img.ip.dance/ipdance.png" width="300" height="300">

## What is ip.dance?

ip.dance is a free service that returns your current public IP address. It's available at https://ip.dance/. It's API is available at https://api.ip.dance or by sending the content type of `application/json` to the URL [at ip.dance.](https://ip.dance/.)

## What is this Go package?

This Go package, is a simple library that fetches the current external IP address using the API [at ip.dance.](https://ip.dance/.)

## To use this library: 

```
package main

import (
    "fmt"
    "log"

    "github.com/patramsey/ip-dance-api"
)

func main() {
    os.Setenv("IP_DANCE_URL", "https://api.ip.dance")
    os.Setenv("IP_DANCE_USERAGENT", "golang-ip-dance-package-0.1")

    ip, err := ipdance.GetIp()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Current external IP address:", ip)
}
```

## To generate the ipdance-api.gen.go 

oapi-codegen --config=oapi-gen-config.yaml api.ip.dance-spec.json

This uses the oapi-codegen tool to generate the ipdance-api.gen.go file from the open API specification at [api.ip.dance-spec.json](https://github.com/patramsey/ip-dance-api/blob/master/api.ip.dance-spec.json). The configuration file is located at [oapi-gen-config.yaml](https://github.com/patramsey/ip-dance-api/blob/master/oapi-gen-config.yaml)