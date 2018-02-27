package main

import (
	"flag"
	"log"

	dumbno "github.com/ncsa/dumbno-client-go"
)

var (
	src, dst     string
	sport, dport int
	proto        string

	addr string
)

func init() {
	flag.StringVar(&src, "src", "", "Source ip")
	flag.StringVar(&dst, "dst", "", "Destination ip")

	flag.IntVar(&sport, "sport", 0, "Source port")
	flag.IntVar(&dport, "dport", 0, "Destination port")

	flag.StringVar(&proto, "proto", "ip", "protocol")

	flag.StringVar(&addr, "addr", "127.0.0.1:9000", "dumbno endpoint")
}

func main() {
	flag.Parse()
	c, err := dumbno.NewClient(addr)
	if err != nil {
		log.Fatal(err)
	}
	err = c.AddACL(dumbno.FilterRequest{
		Src:   src,
		Dst:   dst,
		Sport: sport,
		Dport: dport,
		Proto: proto,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("ok")
}
