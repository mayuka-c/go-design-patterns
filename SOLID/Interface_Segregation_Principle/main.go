package main

// ISP
// Interface Segregation Principle
// A client should never be forced to implement an interface that it doesn't use
// Instead of one fat interface, many small interfaces are preferred based on groups of methods
// This way, a client can never be forced to implement methods it does not use

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type Document struct {
}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (p *MyPrinter) Print(d Document) {}

type Photocopier struct {
}

func (p *Photocopier) Print(d Document) {}

func (p *Photocopier) Scan(d Document) {}

// Combining interfaces
type MultiFunctionDevice interface {
	Printer
	Scanner
}
