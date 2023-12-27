package src

import (
	"fmt"
	"ilyasmirnov03/spacetraders-bot/helpers"
	_type "ilyasmirnov03/spacetraders-bot/type"
	"strconv"
	"time"
)

type Inventory struct {
	Symbol      string
	Name        string
	Description string
	Units       int
}

type Extract struct {
	Cargo    Cargo
	Cooldown Cooldown
}

type Cargo struct {
	Capacity  int
	Units     int
	Inventory []Inventory
}

type NavResponse struct {
	Data _type.Nav
}

type ExtractResponse struct {
	Data Extract
}

type Cooldown struct {
	ShipSymbol       string
	TotalSeconds     int
	RemainingSeconds int
	Expiration       string
}

type void_func func()

var mine_waypoint string
var sell_waypoint string
var mining_ship string
var cargo Cargo

func StartMining() {
	fmt.Println("Mining ship: ")
	fmt.Scanln(&mining_ship)
	fmt.Println("Mine waypoint symbol: ")
	fmt.Scanln(&mine_waypoint)
	fmt.Println("Sell waypoint symbol: ")
	fmt.Scanln(&sell_waypoint)
	extract_resources()
}

func extract_resources() {
	body, err := CallApi[ExtractResponse]("/my/ships/"+mining_ship+"/extract", "POST", nil)
	if err != nil {
		return
	}
	left := body.Data.Cargo.Capacity - body.Data.Cargo.Units
	fmt.Println("Ship just mined", strconv.Itoa(left)+" spaces left to mine")
	if left == 0 {
		fmt.Println("Ship's cargo is full, navigating to selling point.")
		navigate_to_waypoint(sell_waypoint, sell_cargo)
		return
	}
	time.Sleep(time.Duration(body.Data.Cooldown.RemainingSeconds) * time.Second)
	extract_resources()
}

func navigate_to_waypoint(waypoint string, on_arrival void_func) {
	body, err := CallApi[NavResponse]("/my/ships/"+mining_ship+"/navigate", "POST", []byte(`{"waypointSymbol": "`+waypoint+`"}`))
	if err != nil {
		return
	}
	time_to_arrival, _ := helpers.TimeDiffInSeconds(body.Data.Route.Arrival)
	fmt.Printf("Navigating to " + waypoint + ", arrives in " + strconv.Itoa(int(time_to_arrival)) + "s")
	time.Sleep(time.Duration(time_to_arrival) * time.Second)
	on_arrival()
}

func sell_cargo() {

}
