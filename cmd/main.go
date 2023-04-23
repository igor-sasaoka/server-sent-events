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

    // non-blocking with goroutine
    go handleInput() 
    
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-type", "text/event-stream")
        // message value is updated inside handleInput() func
        w.Write(message)
    })

    log.Fatal(http.ListenAndServe(":8080", nil)) 

}

// updates the message via user input
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

// format must be "data: {message}\n\n"
func parseMessage(i string) []byte {
    return []byte("data: " + i + "\n") 
}
