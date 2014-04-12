knight
======

One HTTP web server with reloader for Go, knight detects file changes and restart the server automatically.

##install
    
    $ go get github.com/fengsp/knight

##usage
Basically you just need to set one watching path.
    
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
    	knight := knight.NewKnight(root="/private/tmp/test")
    	knight.ListenAndServe(":8000", nil)
	}
