### Pointer & Value Receiver 
example 1
```
package main

import "fmt"

type T struct {
	ID int
}

func (t *T) PrintID() { // pointer
// func (t T) PrintID() { // value
	fmt.Println(t.ID)
}

func F1() {
	ts := []T{{1}, {2}, {3}}
	for _, t := range ts {
		defer t.PrintID()
	}
}

func F2() {
	ts := []*T{&T{1}, &T{2}, &T{3}}
	for _, t := range ts {
		defer t.PrintID()
	}
}

func main() {
	fmt.Println("F1()")
	F1()
	fmt.Println()
	fmt.Println("F2()")
	F2()
}
```
example 2
```
type Car struct {
  model string
}
func (c *Car) PrintModel() {
// func (c Car) PrintModel() {
  fmt.Println(c.model)
}
func main() {
  c := Car{model: "DeLorean DMC-12"}
  defer c.PrintModel()
  c.model = "Chevrolet Impala"
}
```
https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01  

What’s going on?  
> func(receiver Type) f(input) result  
> =>  
> func f(receiver Type, input) result  

Remember that the passed params to a deferred func are saved aside immediately without waiting for the deferred func to be run.

So, when a method with a value-receiver is used with defer, the receiver will be copied (in this case, Car) at the time of registering, and the changes to it wouldn’t be visible (Car.model). Because the receiver is also an input param and evaluated immediately to “DeLorean DMC-12” when it’s registered with the defer.

On the other hand, when the receiver is a pointer when it’s called with defer, a new pointer is created, but the address it points to would be the same as the “c” pointer above. So, any changes to it would be reflected flawlessly.