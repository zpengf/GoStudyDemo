package Duotai

import "fmt"

type Bio struct {
	name string
}

func (b *Bio) SetName(name string) {
	b.name = name
}

type Live interface {
	run()
	eat()
}

type Person struct {
	Bio
}
type Animal struct {
	Bio
}

func (p *Person) run()  {
	fmt.Println(p.Bio.name+"跑步")
}
func (p *Person) eat()  {
	fmt.Println(p.Bio.name+"吃东西")
}

func (a *Animal) run()  {
	fmt.Println(a.Bio.name+"跑步")
}
func (a *Animal) eat()  {
	fmt.Println(a.Bio.name+"吃东西")
}
func test(live Live)  {
	live.run()
	live.eat()
}

func main() {
	p1:=new(Person)
	p1.SetName("人类")
	test(p1)

	p:=Person{Bio{name: "aaa"}}
	test(&p)

	a1:=new(Animal)
	a1.SetName("laohu")
	test(a1)

}