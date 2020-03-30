package main

import (
	"flag"
	"fmt"
	"net-forward/app"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprint(*i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var rules arrayFlags
	flag.Var(&rules, "f", "-f tcp/local_ip/local_port/tcp/remote_ip/remote_port")

	flag.Parse()
	if len(rules) == 0 {
		flag.Usage()
		return
	}

	app.StartForward(rules)
}
