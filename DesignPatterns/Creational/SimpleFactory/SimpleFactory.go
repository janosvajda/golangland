package main

import (
	"fmt"
)

type Engine interface {
	Assamble()
	Start()
}

type CarEngine struct{}
type AirplaneEngine struct{}
type SteamEngine struct{}

func (c CarEngine) Assamble() {
	fmt.Println("Assembling components for car engine", c)
}

func (a AirplaneEngine) Assamble() {
	fmt.Println("Assembling components for train engine", a)
}

func (s SteamEngine) Assamble() {
	fmt.Println("Assembling components for Steam engine", s)
}

func (c CarEngine) Start() {
	fmt.Println("Starting the car engine", c)
}

func (a AirplaneEngine) Start() {
	fmt.Println("Starting the train engine", a)
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
