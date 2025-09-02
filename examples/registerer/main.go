package main

import (
	"flag"
	"fmt"
	"net/netip"
	"net/url"
	"time"

	"github.com/oaiiae/flag/values"
)

func main() {
	var (
		reg = values.FlagSetRegisterer(flag.CommandLine).WithEnv("FOO_")
		t   = reg.Time("ts", time.Now(), time.RFC3339, "set a `RFC3339 date`")
		u   = reg.URL("url", &url.URL{}, "set `URL`")
		as  = reg.IPAddrSlice("ipaddrs", nil, ",", "set comma-separated IP `addresses`")
		p   = reg.IPPrefix("prefix", netip.MustParsePrefix("0.0.0.0/0"), "set `CIDR`")
		ms  = reg.MailAddrList("mail", nil, "add a mail `address`")
		d   = reg.Duration("duration", 1234*time.Second, "set duration")
	)
	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		v := f.Value.(flag.Getter)
		fmt.Printf("%s\t%T\t%v\n", f.Name, v.Get(), v.Get())
	})
	fmt.Println()
	fmt.Println("t=", *t)
	fmt.Println("u=", *u)
	fmt.Println("as=", *as)
	fmt.Println("p=", *p)
	fmt.Println("ms=", *ms)
	fmt.Println("d=", *d)
}
