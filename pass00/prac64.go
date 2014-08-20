// golang practice 64 - Exit

package main

import "fmt"
import "os"

func main() {

	defer fmt.Println("!")

	os.Exit(3)
}
