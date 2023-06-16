package main

func (factory *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	}
	return fruit

}

func main() {

	factory := new(Factory)
	apple := factory.CreateFruit("apple")
	apple.Show()
	banana := factory.CreateFruit("banana")
	banana.Show()
}
