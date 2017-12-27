package main

import (
	"log"

	dumbno "github.com/ncsa/dumbno-client-go"
)

func main() {
	c, err := dumbno.NewClient("127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	err = c.AddACL(dumbno.FilterRequest{
		Src: "1.2.3.4",
		Dst: "5.6.7.8",
	})
	if err != nil {
		log.Fatal(err)
	}
}
