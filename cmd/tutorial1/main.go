package main //Serve per indicare al compilatore che dentro questo package (cartella) c'è una funzione main da cui partire
import (
	"fmt"
)

type gasEngine struct {
    mpg uint8
    gallons uint8
}

type electricEngine struct{
    mpkwh uint8
    kwh uint8
}

type car [T gasEngine | electricEngine] struct{
	carMake string
	carModel string
	engine T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake: "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons:12.4,
			mpg:40
		}
	}
	var electricCar = car[electricEngine]{
		carMake: "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:57.5,
			mpg:4.17
		}
	}
}