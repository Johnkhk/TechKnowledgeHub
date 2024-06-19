package builder

import (
	"fmt"
)

type House struct {
	Foundation string
	Structure  string
	Roof       string
	Rooms      string
	Backyard   string
}

func (h House) String() string {
	return fmt.Sprintf("House with foundation: %s, structure: %s, roof: %s, rooms: %s, backyard: %s",
		h.Foundation, h.Structure, h.Roof, h.Rooms, h.Backyard)
}

type HouseBuilder interface {
	BuildFoundation() HouseBuilder
	BuildStructure() HouseBuilder
	BuildRoof() HouseBuilder
	BuildRooms(number int) HouseBuilder
	BuildBackyard(typ string) HouseBuilder
	Build() House
}

type WoodenHouseBuilder struct {
	house House
}

func (b *WoodenHouseBuilder) BuildFoundation() HouseBuilder {
	b.house.Foundation = "Wooden Foundation"
	return b
}

func (b *WoodenHouseBuilder) BuildStructure() HouseBuilder {
	b.house.Structure = "Wooden Structure"
	return b
}

func (b *WoodenHouseBuilder) BuildRoof() HouseBuilder {
	b.house.Roof = "Wooden Roof"
	return b
}

func (b *WoodenHouseBuilder) BuildRooms(number int) HouseBuilder {
	b.house.Rooms = fmt.Sprintf("%d Wooden Rooms", number)
	return b
}

func (b *WoodenHouseBuilder) BuildBackyard(typ string) HouseBuilder {
	b.house.Backyard = fmt.Sprintf("Wooden %s Backyard", typ)
	return b
}

func (b *WoodenHouseBuilder) Build() House {
	return b.house
}

type BrickHouseBuilder struct {
	house House
}

func (b *BrickHouseBuilder) BuildFoundation() HouseBuilder {
	b.house.Foundation = "Concrete Foundation"
	return b
}

func (b *BrickHouseBuilder) BuildStructure() HouseBuilder {
	b.house.Structure = "Brick Structure"
	return b
}

func (b *BrickHouseBuilder) BuildRoof() HouseBuilder {
	b.house.Roof = "Concrete Roof"
	return b
}

func (b *BrickHouseBuilder) BuildRooms(number int) HouseBuilder {
	b.house.Rooms = fmt.Sprintf("%d Brick Rooms", number)
	return b
}

func (b *BrickHouseBuilder) BuildBackyard(typ string) HouseBuilder {
	b.house.Backyard = fmt.Sprintf("Brick %s Backyard", typ)
	return b
}

func (b *BrickHouseBuilder) Build() House {
	return b.house
}

type Director struct {
	builder HouseBuilder
}

func (d *Director) ConstructSimpleHouse() House {
	return d.builder.BuildFoundation().BuildStructure().BuildRoof().BuildRooms(2).Build()
}

func (d *Director) ConstructLuxuryHouse() House {
	return d.builder.BuildFoundation().BuildStructure().BuildRoof().BuildRooms(5).BuildBackyard("Luxury").Build()
}

func (d *Director) ConstructHouseWith3Bedrooms() House {
	return d.builder.BuildFoundation().BuildStructure().BuildRoof().BuildRooms(3).BuildBackyard("Standard").Build()
}
