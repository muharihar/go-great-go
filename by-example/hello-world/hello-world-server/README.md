# Hello World (HTTP) Server

Let's make a simple "Hello World (HTTP) Server" using Go. Here's the full source code.


#### main.go

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\n")
	})

	fmt.Println("Press ctrl+c to shutdown the application.")
	log.Fatal(http.ListenAndServe("127.0.0.1:8134", nil))
}

```

#### Project Structure

```bash
hello-world-server
	├── README.md
	├── build_and_run.sh
	├── call_server_with_curl.sh
	├── dist
	│   ├── default
	│   │   └── hello-world-server
	│   ├── linux
	│   │   └── hello-world-server
	│   ├── mac
	│   │   └── hello-world-server
	│   └── win
	│       └── hello-world-server.exe
	└── main.go

```

## Build and Run

#### Running with "go run"

```bash
$ go run main.go
Press ctrl+c to shutdown the application.

```

Call hello world server with curl

```bash
$ sh call_server_with_curl.sh
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8134 (#0)
> GET /hello-world HTTP/1.1
> Host: 127.0.0.1:8134
> User-Agent: curl/7.54.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Sat, 06 Oct 2018 15:55:16 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
< 
Hello World!
* Connection #0 to host 127.0.0.1 left intact

```

#### Golang Playground

N/A