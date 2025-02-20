// testing stderr etc
package main
import (
	"fmt"
	err "errors"
)

// exec 2>>( while read X; do print "\e[91m${X}\e[0m" > /dev/tty; done & )

func main() {
// test stdout
fmt.Println("Hello, Standard World!")
	fmt.Println(err.New("testing stderr etc"))
fmt.Println("Hello, standard World!")
}