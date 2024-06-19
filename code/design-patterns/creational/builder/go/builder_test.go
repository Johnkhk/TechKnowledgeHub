package builder

import "fmt"

func main() {
	woodenHouseBuilder := &WoodenHouseBuilder{}
	brickHouseBuilder := &BrickHouseBuilder{}

	director := Director{builder: woodenHouseBuilder}
	simpleWoodenHouse := director.ConstructSimpleHouse()
	fmt.Println(simpleWoodenHouse) // Output: House with foundation: Wooden Foundation, structure: Wooden Structure, roof: Wooden Roof, rooms: 2 Wooden Rooms, backyard:

	director = Director{builder: brickHouseBuilder}
	luxuryBrickHouse := director.ConstructLuxuryHouse()
	fmt.Println(luxuryBrickHouse) // Output: House with foundation: Concrete Foundation, structure: Brick Structure, roof: Concrete Roof, rooms: 5 Brick Rooms, backyard: Brick Luxury Backyard
}
