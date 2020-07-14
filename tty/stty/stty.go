/*package main

import (
	"os/exec"
	"fmt"
	"log"
)

func main() {
	out, err := exec.Command("stty", "size").Output()
	fmt.Printf("out: %#v\n", out)
	fmt.Printf("err: %#v\n", err)
	if err != nil {
		log.Fatal(err)
	}
}*/

package main

import (
"os/exec"
"fmt"
"log"
"os"
)

func main() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	fmt.Printf("out: %#v\n", string(out))
	fmt.Printf("err: %#v\n", err)
	if err != nil {
		log.Fatal(err)
	}
}