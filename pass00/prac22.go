// golang practice 22 - Channels

package main

import "fmt"

func main() {

    messages := make(chan string)


    go func() { messages <- "ping" }()



    msg := <-messages
    fmt.Println(msg)
}