package bar

import(
	log "github.com/sirupsen/logrus"
)

/**
	INTERFACE COMPOSITION
 */
// Writer contract definition
type Writer interface {
	Write([]byte) (int, error)
}

// Reader contract definition
type Reader interface {
	Read() ([]byte, error)
}

// ReadWriter composition example
type ReadWriter interface {
	Reader
	Writer
}

// defaultReadWriter implements ReadWriter OOP
type defaultReadWriter struct {}

func NewDefaultReadWriter() ReadWriter{
	return &defaultReadWriter{}
}

func (d defaultReadWriter) Read() ([]byte, error) {
	return nil, nil
}

func (d defaultReadWriter) Write(r []byte) (int, error) {

	return len(r), nil
}

/**
ANONYMOUS INHERITANCE
*/
type Foo struct {
	Index int
	bar string
}

func (f *Foo) Fooo() error{
	log.Info("Foo")
	return nil
}

func (f *Foo) FooBar() error{
	log.Info("Foo")
	return nil
}

// Bar embeds Foo Struct, what give us access to the background type
type Bar struct {
	Foo
	foo string
}

// Note on initialitation how we need to instantiate "Foo"
func NewBar() *Bar {
	return &Bar{
		Foo: Foo{
			Index: 10,
			bar: "asdasd",
		},
		foo: "",
	}
}

// Bar has access to all embedded Foo Attributes and methods
func (f *Bar) handle() {
	f.Index++
	f.foo = ""
}

// Uncommenting this will override Foo::Fooo method
/*
func (f *Bar) Fooo() error{
	log.Info("Bar")
	return nil
}*/
