package main

// ============工厂模块============
// 抽象工厂接口

type AbstractFactory interface {
}

//具体工厂实现

type AppleFactory struct {
	AbstractFactory
}

func (appleFactory *AppleFactory) createFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (bananaFactory *BananaFactory) createFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)
	return fruit
}

func main() {

	appleFactory := new(AppleFactory)
	apple := appleFactory.createFruit()
	apple.Show()

	bananaFactory := new(BananaFactory)
	banana := bananaFactory.createFruit()
	banana.Show()
}
