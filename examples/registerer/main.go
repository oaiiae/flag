package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"

	"gitlab.com/oaiiae/flag/values"
)

var (
	reg = values.FlagSetRegisterer(flag.CommandLine)
	t   = reg.Time("ts", time.Now(), time.RFC3339, "set a `RFC3339 date`")
	u   = reg.URL("url", &url.URL{}, "set `URL`")
	as  = reg.IPAddrSlice("ipaddrs", nil, ",", "set comma-separated IP `addresses`")
)

func init() { flag.Parse() }

func main() {
	flag.VisitAll(func(f *flag.Flag) {
		v := f.Value.(flag.Getter)
		fmt.Printf("%s\t%T\t%v\n", f.Name, v.Get(), v.String())
	})
	fmt.Println()
	fmt.Println("t=", *t)
	fmt.Println("u=", *u)
	fmt.Println("as=", *as)
}
