package main

import (
	"fmt"
)

//This interface guarantee that each engine type must implement every method.
type Engine interface {
	Assamble()
	Start()
}

type CarEngine struct{}
type AirplaneEngine struct{}
type SteamEngine struct{}

//Car
func (c CarEngine) Assamble() {
	fmt.Println("Assembling components for car engine", c)
}

func (c CarEngine) Start() {
	fmt.Println("Starting the car engine", c)
}

//Airplane
func (a AirplaneEngine) Assamble() {
	fmt.Println("Assembling components for train engine", a)
}

func (a AirplaneEngine) Start() {
	fmt.Println("Starting the train engine", a)
}

//Steam Engine
func (s SteamEngine) Assamble() {
	fmt.Println("Assembling components for Steam engine", s)
}

func (s SteamEngine) Start() {
	fmt.Println("Starting the steam engine", s)
}

func Assamble(e Engine) {
	if e != nil {
		e.Assamble()
	}
}

func Start(e Engine) {
	if e != nil {
		e.Start()
	}
}

//This is the factory where engines are creating
func GetEngine(engineType string) Engine {
	switch engineType {
	case "car":
		return &CarEngine{}
	case "airplane":
		return &AirplaneEngine{}
	case "steam":
		return &SteamEngine{}
	default:
		fmt.Printf("Undefined engine type")
		return nil
	}
}

func main() {

	//Creates an Airplane engine
	engine := GetEngine("airplane")
	Assamble(engine)
	Start(engine)

	//Creates a Steam engine
	engine2 := GetEngine("steam")
	Assamble(engine2)
	Start(engine2)

	//Creates a Car engine
	engine3 := GetEngine("car")
	Assamble(engine3)
	Start(engine3)

}
