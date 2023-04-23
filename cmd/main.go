package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "bufio"
)
 
var message = []byte("data: \n\n")

func main() {
    fmt.Println("server running...")

    // goroutine call so this is non blocking 
    go handleInput() 
    
    // routing for static index.html file with front-end logic
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    // SSE endpoint that will send and event stream with the current message value
    http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-type", "text/event-stream")

        w.Write(message)
    })

    log.Fatal(http.ListenAndServe(":8080", nil)) 

}

// creates a input listener that updates the message that will be sent
// as event stream
func handleInput() {
    r := bufio.NewReader(os.Stdin)
    fmt.Println("Server Sent Events shell")

    for {
        fmt.Print("-> ")
        s, err := r.ReadString('\n')        
        if err != nil {
            panic(err)
        }    

        message = parseMessage(s)
    }
}

// SSE messages MUST have the format "data: {message}\n\n"
func parseMessage(i string) []byte {
    return []byte("data: " + i + "\n") 
}
