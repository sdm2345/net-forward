package main

import (
	"flag"
	"fmt"
	"github.com/sdm2345/net-forward/app"
	"log"
	"net/url"
)

type ArrayParam []string

func (i *ArrayParam) String() string {
	return fmt.Sprint(*i)
}

func (i *ArrayParam) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var rules ArrayParam
	var listens ArrayParam
	flag.Var(&listens, "l", "-l tcp://0.0.0.0:1234 local listen address ")
	flag.Var(&rules, "f", "-f tcp://remote_ip:80 remote address")

	flag.Parse()
	if len(rules) == 0 || len(rules) != len(listens) {
		flag.Usage()
		return
	}
	//check
	for i := range listens {
		info, err := url.Parse(listens[i])
		if err != nil {
			log.Fatal("err", err)
		}
		if info.Scheme != "tcp" {
			log.Fatal("unsupported schema", info.Scheme)
		}
		listens[i] = info.Host
		info, err = url.Parse(rules[i])
		if err != nil {
			log.Fatal("err", err)
		}
		if info.Scheme != "tcp" {
			log.Fatal("unsupported schema", info.Scheme)
		}
		rules[i] = info.Host
	}
	app.StartForward(listens, rules)
}
