package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

// START OMIT
func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadFile(usr.HomeDir + "/.bash_profile")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dat))
}
// END OMIT
