package main

import (
	"fmt"
	"reflect"
)

type Aa struct {
	p int
}

func (a *Aa) pp() {
	v := reflect.ValueOf(&a)
	v = v.Elem()
	newa := reflect.ValueOf(new(Aa))
	v.Set(newa)
	fmt.Printf("pp %p\n", a)
}

func main() {
	a := new(Aa)
	fmt.Printf("%p\n", a)
	a.pp()
	fmt.Printf("%p\n", a)

}
