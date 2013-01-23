package term

import "fmt"
import "strconv"

type Float float64

func NewFloat(text string) *Float {
    f, err := strconv.ParseFloat(text, 64)
    maybePanic(err)
    return (*Float)(&f)
}

func (self *Float) Value() float64 {
    return float64(*self)
}

func (self *Float) Functor() string {
    panic("Floats have no Functor()")
}
func (self *Float) Arity() int {
    panic("Floats have no Arity()")
}
func (self *Float) Arguments() []Term {
    panic("Floats have no Arguments()")
}
func (self *Float) Body() Term {
    panic("Floats have no Body()")
}
func (self *Float) Head() Term {
    panic("Floats have no Head()")
}
func (self *Float) IsClause() bool {
    return false
}
func (self *Float) String() string {
    return fmt.Sprintf("%g", self.Value())
}
func (self *Float) Indicator() string {
    return self.String()
}
func (self *Float) Error() error {
    panic("Can't call Error() on a Float")
}

func (self *Float) ReplaceVariables(env Bindings) Term {
    return self
}
