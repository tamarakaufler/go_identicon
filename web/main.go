package main

import (
    "fmt"
    "log"

    "net/http"

    "github.com/tamarakaufler/go_identicon"
    "github.com/julienschmidt/httprouter"
)

// Routing handler
//  params: http.ResponseWriter
//  params: *http.Request
//  params: httprouter.Params

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    // Retrieve the string that will be used to create a unique identicon
    s := ps.ByName("email")

    // Create a unique identicon based on the string
    icon := identicon.New(s)
    icon_data := icon.Create()
    log.Printf("Created identicon for '%s'\n", s)

    // Display the identicon
    //  The image is embedded on the page using the Data URI scheme (as a base64 encoded string)
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "Creating identicon for [%s]<br /><br />", s)
    fmt.Fprintf(w, "<html><body><img src='data:image/png;base64,%s'></body></html>", icon_data)
}

func main() {

    // Routing
    //  Using 3rd party routing module because net/http allows setting up only static routes
    //  We want to provide the string as part of the url
    router := httprouter.New()
    router.GET("/identicon/:email", handler)

    // HTTP server
    log.Fatal(http.ListenAndServe(":3333", router))
}
