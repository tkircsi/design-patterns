package isp

// Interface Segregation Principle
// Don't put too much into an interface. Split into separate interfaces
// YAGNI - You Ain't Going to Need It

func Main() {

}

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Sacn(d Document) {

}

type OldFashionedPrinter struct{}

func (o OldFashionedPrinter) Print(d Document) {
	// OK
}

// Depricated
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Depricated
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// ISP
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {

}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {

}

func (p Photocopier) Scan(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
