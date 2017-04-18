knight
======

One HTTP web server with reloader for Go, knight detects file changes and restart the server automatically.

## install
    
    $ go get github.com/fengsp/knight

## usage
Basically you just need to set one watching path.

```Go
package main
    
import (
    "fmt"
    "net/http"
    "github.com/fengsp/knight"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "It works!")
}

func main() {
    http.HandleFunc("/", handler)
    // pass your root path in
    // pass "" to use current working directory
    knight := knight.NewKnight("/private/tmp/test")
    knight.ListenAndServe(":8000", nil)
}
```

You can get it working like this:

    $ go run test.go
     * Knight serving on :8000
 	 * Restarting with reloader
 	 * Detected change, reloading
 	 * Restarting with reloader
 	 * Detected change, reloading
 	 * Restarting with reloader
