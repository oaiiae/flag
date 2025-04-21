package values

import (
	"flag"
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

func (r Registerer) IPAddrSlice(name string, value []netip.Addr, sep string, usage string) *[]netip.Addr {
	r.Varer.Var(StringerSliceVar(&value, sep, netip.ParseAddr), name, usage)
	return &value
}

func (r Registerer) IPAddrSliceVar(p *[]netip.Addr, name string, value []netip.Addr, sep string, usage string) {
	*p = value
	r.Varer.Var(StringerSliceVar(p, sep, netip.ParseAddr), name, usage)
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
