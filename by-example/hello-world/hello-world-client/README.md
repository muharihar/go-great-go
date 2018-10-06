# 1.4 Hello World (HTTP) Client

Let's make a simple "Hello World (HTTP) Client" using Go. Here's the full source code.


#### main.go

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// using hello-world-server, we need to run hello-world-server application first.
	resp, err := http.Get("http://127.0.0.1:8134/hello-world")

	// using mocky.io, we need an internet connection
	// resp, err := http.Get("http://www.mocky.io/v2/5bb8df303100005d003ed8c4")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

```

#### Project Structure

```bash
hello-world-client
    ├── README.md
    ├── build_and_run.sh
    ├── build_and_run_with_output.sh
    ├── dist
    │   ├── default
    │   │   ├── hello-world-client
    │   │   ├── output.hello-world-client.sample.txt
    │   │   └── output.hello-world-client.txt
    │   ├── linux
    │   │   └── hello-world-client
    │   ├── mac
    │   │   └── hello-world-client
    │   └── win
    │       └── hello-world-client.exe
    └── main.go

```

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
Hello World!

```

#### Golang Playground

N/A