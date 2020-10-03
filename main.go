package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {

	fmt.Printf("dirweb %s (%s)\n", version, commit[:7])

	var dir string
	var port int
	var help bool

	flag.StringVar(&dir, "d", "./", "dir")
	flag.IntVar(&port, "p", 3000, "port")
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	fmt.Printf("Listening...  dir=%s port=%d\n", dir, port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe error: %v.", err)
		os.Exit(1)
	}
}
