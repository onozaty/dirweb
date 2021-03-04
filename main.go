package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	flag "github.com/spf13/pflag"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {

	if len(commit) > 7 {
		commit = commit[:7]
	}
	fmt.Printf("dirweb v%s (%s)\n", version, commit)

	var dir string
	var port int
	var help bool

	flag.StringVarP(&dir, "dir", "d", "./", "dir")
	flag.IntVarP(&port, "port", "p", 3000, "port")
	flag.BoolVarP(&help, "help", "h", false, "help")
	flag.Parse()
	flag.CommandLine.SortFlags = false

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
