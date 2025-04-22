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

func (r Registerer) Bool(name string, value bool, usage string) *bool {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) BoolVar(p *bool, name string, value bool, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) BoolList(name string, value []bool, usage string) *[]bool {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) BoolListVar(p *[]bool, name string, value []bool, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) BoolSlice(name string, value []bool, sep string, usage string) *[]bool {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) BoolSliceVar(p *[]bool, name string, value []bool, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Complex64(name string, value complex64, usage string) *complex64 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Complex64Var(p *complex64, name string, value complex64, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Complex64List(name string, value []complex64, usage string) *[]complex64 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Complex64ListVar(p *[]complex64, name string, value []complex64, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Complex64Slice(name string, value []complex64, sep string, usage string) *[]complex64 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Complex64SliceVar(p *[]complex64, name string, value []complex64, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Complex128(name string, value complex128, usage string) *complex128 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Complex128Var(p *complex128, name string, value complex128, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Complex128List(name string, value []complex128, usage string) *[]complex128 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Complex128ListVar(p *[]complex128, name string, value []complex128, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Complex128Slice(name string, value []complex128, sep string, usage string) *[]complex128 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Complex128SliceVar(p *[]complex128, name string, value []complex128, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Int(name string, value int, usage string) *int {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) IntVar(p *int, name string, value int, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) IntList(name string, value []int, usage string) *[]int {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) IntListVar(p *[]int, name string, value []int, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) IntSlice(name string, value []int, sep string, usage string) *[]int {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) IntSliceVar(p *[]int, name string, value []int, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Int8(name string, value int8, usage string) *int8 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Int8Var(p *int8, name string, value int8, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Int8List(name string, value []int8, usage string) *[]int8 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Int8ListVar(p *[]int8, name string, value []int8, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Int8Slice(name string, value []int8, sep string, usage string) *[]int8 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Int8SliceVar(p *[]int8, name string, value []int8, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}
func (r Registerer) Int16(name string, value int16, usage string) *int16 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Int16Var(p *int16, name string, value int16, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Int16List(name string, value []int16, usage string) *[]int16 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Int16ListVar(p *[]int16, name string, value []int16, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Int16Slice(name string, value []int16, sep string, usage string) *[]int16 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Int16SliceVar(p *[]int16, name string, value []int16, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Int32(name string, value int32, usage string) *int32 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Int32Var(p *int32, name string, value int32, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Int32List(name string, value []int32, usage string) *[]int32 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Int32ListVar(p *[]int32, name string, value []int32, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Int32Slice(name string, value []int32, sep string, usage string) *[]int32 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Int32SliceVar(p *[]int32, name string, value []int32, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Int64(name string, value int64, usage string) *int64 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Int64Var(p *int64, name string, value int64, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Int64List(name string, value []int64, usage string) *[]int64 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Int64ListVar(p *[]int64, name string, value []int64, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Int64Slice(name string, value []int64, sep string, usage string) *[]int64 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Int64SliceVar(p *[]int64, name string, value []int64, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Uint(name string, value uint, usage string) *uint {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) UintVar(p *uint, name string, value uint, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) UintList(name string, value []uint, usage string) *[]uint {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) UintListVar(p *[]uint, name string, value []uint, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) UintSlice(name string, value []uint, sep string, usage string) *[]uint {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) UintSliceVar(p *[]uint, name string, value []uint, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Uint8(name string, value uint8, usage string) *uint8 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Uint8Var(p *uint8, name string, value uint8, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Uint8List(name string, value []uint8, usage string) *[]uint8 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Uint8ListVar(p *[]uint8, name string, value []uint8, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Uint8Slice(name string, value []uint8, sep string, usage string) *[]uint8 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Uint8SliceVar(p *[]uint8, name string, value []uint8, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Uint16(name string, value uint16, usage string) *uint16 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Uint16Var(p *uint16, name string, value uint16, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Uint16List(name string, value []uint16, usage string) *[]uint16 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Uint16ListVar(p *[]uint16, name string, value []uint16, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Uint16Slice(name string, value []uint16, sep string, usage string) *[]uint16 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Uint16SliceVar(p *[]uint16, name string, value []uint16, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Uint32(name string, value uint32, usage string) *uint32 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Uint32Var(p *uint32, name string, value uint32, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Uint32List(name string, value []uint32, usage string) *[]uint32 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Uint32ListVar(p *[]uint32, name string, value []uint32, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Uint32Slice(name string, value []uint32, sep string, usage string) *[]uint32 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Uint32SliceVar(p *[]uint32, name string, value []uint32, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Uint64(name string, value uint64, usage string) *uint64 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Uint64Var(p *uint64, name string, value uint64, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Uint64List(name string, value []uint64, usage string) *[]uint64 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Uint64ListVar(p *[]uint64, name string, value []uint64, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Uint64Slice(name string, value []uint64, sep string, usage string) *[]uint64 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Uint64SliceVar(p *[]uint64, name string, value []uint64, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Float32(name string, value float32, usage string) *float32 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Float32Var(p *float32, name string, value float32, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Float32List(name string, value []float32, usage string) *[]float32 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Float32ListVar(p *[]float32, name string, value []float32, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Float32Slice(name string, value []float32, sep string, usage string) *[]float32 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Float32SliceVar(p *[]float32, name string, value []float32, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) Float64(name string, value float64, usage string) *float64 {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) Float64Var(p *float64, name string, value float64, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) Float64List(name string, value []float64, usage string) *[]float64 {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) Float64ListVar(p *[]float64, name string, value []float64, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) Float64Slice(name string, value []float64, sep string, usage string) *[]float64 {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) Float64SliceVar(p *[]float64, name string, value []float64, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

func (r Registerer) String(name string, value string, usage string) *string {
	r.Varer.Var(BasicVar(&value), name, usage)
	return &value
}

func (r Registerer) StringVar(p *string, name string, value string, usage string) {
	*p = value
	r.Varer.Var(BasicVar(p), name, usage)
}

func (r Registerer) StringList(name string, value []string, usage string) *[]string {
	r.Varer.Var(BasicListVar(&value), name, usage)
	return &value
}

func (r Registerer) StringListVar(p *[]string, name string, value []string, usage string) {
	*p = value
	r.Varer.Var(BasicListVar(p), name, usage)
}

func (r Registerer) StringSlice(name string, value []string, sep string, usage string) *[]string {
	r.Varer.Var(BasicSliceVar(&value, sep), name, usage)
	return &value
}

func (r Registerer) StringSliceVar(p *[]string, name string, value []string, sep string, usage string) {
	*p = value
	r.Varer.Var(BasicSliceVar(p, sep), name, usage)
}

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
