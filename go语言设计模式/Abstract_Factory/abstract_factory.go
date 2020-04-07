package    Abstract_Factory

import "fmt"

type Launch interface {
	Cook()
}

type Rice struct {
	
}

func (r *Rice)Cook(){
	fmt.Println("it  is  rice")
}

type Tomato struct {

}
func (t  *Tomato)Cook() {
	fmt.Println("it is tomato")
}

type LauchFactory interface {
	CreateFood()  Launch
	CreateVegetable() Launch
}

type SimpleLunchFactory struct {

}

func NewSimpleLunchFactory() LauchFactory {
	return &SimpleLunchFactory{}
}

func (s  *SimpleLunchFactory)CreateFood()Launch {
	return  &Rice{}
}

func (s *SimpleLunchFactory)CreateVegetable()Launch {
	return &Tomato{}
}