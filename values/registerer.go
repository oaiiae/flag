package values

import (
	"flag"
	"net/mail"
	"net/netip"
	"net/url"
	"time"
)

type Registerer struct {
	Varer interface {
		Var(value flag.Value, name string, usage string)
	}
}

func FlagSetRegisterer(fs *flag.FlagSet) Registerer { return Registerer{fs} }

func (r Registerer) IPAddr(name string, value netip.Addr, usage string) *netip.Addr {
	r.Varer.Var(StringerVar(&value, netip.ParseAddr), name, usage)
	return &value
}

func (r Registerer) IPAddrVar(p *netip.Addr, name string, value netip.Addr, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, netip.ParseAddr), name, usage)
}

func (r Registerer) IPAddrList(name string, value []netip.Addr, usage string) *[]netip.Addr {
	r.Varer.Var(StringerListVar(&value, netip.ParseAddr), name, usage)
	return &value
}

func (r Registerer) IPAddrListVar(p *[]netip.Addr, name string, value []netip.Addr, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, netip.ParseAddr), name, usage)
}

func (r Registerer) IPAddrSlice(name string, value []netip.Addr, sep string, usage string) *[]netip.Addr {
	r.Varer.Var(StringerSliceVar(&value, sep, netip.ParseAddr), name, usage)
	return &value
}

func (r Registerer) IPAddrSliceVar(p *[]netip.Addr, name string, value []netip.Addr, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, netip.ParseAddr), name, usage)
}

func (r Registerer) IPAddrPort(name string, value netip.AddrPort, usage string) *netip.AddrPort {
	r.Varer.Var(StringerVar(&value, netip.ParseAddrPort), name, usage)
	return &value
}

func (r Registerer) IPAddrPortVar(p *netip.AddrPort, name string, value netip.AddrPort, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, netip.ParseAddrPort), name, usage)
}

func (r Registerer) IPAddrPortList(name string, value []netip.AddrPort, usage string) *[]netip.AddrPort {
	r.Varer.Var(StringerListVar(&value, netip.ParseAddrPort), name, usage)
	return &value
}

func (r Registerer) IPAddrPortListVar(p *[]netip.AddrPort, name string, value []netip.AddrPort, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, netip.ParseAddrPort), name, usage)
}

func (r Registerer) IPAddrPortSlice(name string, value []netip.AddrPort, sep string, usage string) *[]netip.AddrPort {
	r.Varer.Var(StringerSliceVar(&value, sep, netip.ParseAddrPort), name, usage)
	return &value
}

func (r Registerer) IPAddrPortSliceVar(p *[]netip.AddrPort, name string, value []netip.AddrPort, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, netip.ParseAddrPort), name, usage)
}

func (r Registerer) IPPrefix(name string, value netip.Prefix, usage string) *netip.Prefix {
	r.Varer.Var(StringerVar(&value, netip.ParsePrefix), name, usage)
	return &value
}

func (r Registerer) IPPrefixVar(p *netip.Prefix, name string, value netip.Prefix, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, netip.ParsePrefix), name, usage)
}

func (r Registerer) IPPrefixList(name string, value []netip.Prefix, usage string) *[]netip.Prefix {
	r.Varer.Var(StringerListVar(&value, netip.ParsePrefix), name, usage)
	return &value
}

func (r Registerer) IPPrefixListVar(p *[]netip.Prefix, name string, value []netip.Prefix, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, netip.ParsePrefix), name, usage)
}

func (r Registerer) IPPrefixSlice(name string, value []netip.Prefix, sep string, usage string) *[]netip.Prefix {
	r.Varer.Var(StringerSliceVar(&value, sep, netip.ParsePrefix), name, usage)
	return &value
}

func (r Registerer) IPPrefixSliceVar(p *[]netip.Prefix, name string, value []netip.Prefix, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, netip.ParsePrefix), name, usage)
}

func (r Registerer) MailAddr(name string, value *mail.Address, usage string) **mail.Address {
	r.Varer.Var(StringerVar(&value, mail.ParseAddress), name, usage)
	return &value
}

func (r Registerer) MailAddrVar(p **mail.Address, name string, value *mail.Address, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, mail.ParseAddress), name, usage)
}

func (r Registerer) MailAddrList(name string, value []*mail.Address, usage string) *[]*mail.Address {
	r.Varer.Var(StringerListVar(&value, mail.ParseAddress), name, usage)
	return &value
}

func (r Registerer) MailAddrListVar(p *[]*mail.Address, name string, value []*mail.Address, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, mail.ParseAddress), name, usage)
}

func (r Registerer) MailAddrSlice(name string, value []*mail.Address, sep string, usage string) *[]*mail.Address {
	r.Varer.Var(StringerSliceVar(&value, sep, mail.ParseAddress), name, usage)
	return &value
}

func (r Registerer) MailAddrSliceVar(p *[]*mail.Address, name string, value []*mail.Address, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, mail.ParseAddress), name, usage)
}

func (r Registerer) Time(name string, value time.Time, layout string, usage string) *time.Time {
	r.Varer.Var(TimeVar(&value, layout), name, usage)
	return &value
}

func (r Registerer) TimeVar(p *time.Time, name string, value time.Time, layout string, usage string) {
	*p = value
	r.Varer.Var(TimeVar(p, layout), name, usage)
}

func (r Registerer) TimeList(name string, value []time.Time, layout string, usage string) *[]time.Time {
	r.Varer.Var(TimeListVar(&value, layout), name, usage)
	return &value
}

func (r Registerer) TimeListVar(p *[]time.Time, name string, value []time.Time, layout string, usage string) {
	*p = value
	r.Varer.Var(TimeListVar(p, layout), name, usage)
}

func (r Registerer) TimeSliceVar(p *[]time.Time, name string, value []time.Time, sep string, layout string, usage string) {
	*p = value
	r.Varer.Var(TimeSliceVar(p, sep, layout), name, usage)
}

func (r Registerer) TimeSlice(name string, value []time.Time, sep string, layout string, usage string) *[]time.Time {
	r.Varer.Var(TimeSliceVar(&value, sep, layout), name, usage)
	return &value
}

func (r Registerer) URL(name string, value *url.URL, usage string) **url.URL {
	r.Varer.Var(StringerVar(&value, url.Parse), name, usage)
	return &value
}

func (r Registerer) URLVar(p **url.URL, name string, value *url.URL, usage string) {
	*p = value
	r.Varer.Var(StringerVar(p, url.Parse), name, usage)
}

func (r Registerer) URLList(name string, value []*url.URL, usage string) *[]*url.URL {
	r.Varer.Var(StringerListVar(&value, url.Parse), name, usage)
	return &value
}

func (r Registerer) URLListVar(p *[]*url.URL, name string, value []*url.URL, usage string) {
	*p = value
	r.Varer.Var(StringerListVar(p, url.Parse), name, usage)
}

func (r Registerer) URLSlice(name string, value []*url.URL, sep string, usage string) *[]*url.URL {
	r.Varer.Var(StringerSliceVar(&value, sep, url.Parse), name, usage)
	return &value
}

func (r Registerer) URLSliceVar(p *[]*url.URL, name string, value []*url.URL, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, url.Parse), name, usage)
}
