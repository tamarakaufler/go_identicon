package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/tamarakaufler/go_identicon"
    "github.com/julienschmidt/httprouter"
)

// Routing handler
// params: http.ResponseWriter
// params: *http.Request
// params: httprouter.Params

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    // Retrieve the string that will be used to create a unique identicon
    s := ps.ByName("email")
    icon := identicon.New(s)
    icon_data := icon.Create()

    w.Header().Set("Content-Type", "text/html")

    log.Printf("Creating identicon for '%s'\n", s)
    fmt.Fprintf(w, "Creating identicon for [%s]<br /><br />", s)
    fmt.Fprintf(w, "<html><body><img src='data:image/png;base64,%s'></body></html>", icon_data)
}

func main() {

    // Routing
    router := httprouter.New()
    router.GET("/identicon/icon/:email", handler)

    // HTTP server
    log.Fatal(http.ListenAndServe(":3333", router))
}
