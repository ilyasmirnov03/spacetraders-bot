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
	Data _type.NavResponse
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
		cargo = body.Data.Cargo
		navigate_to_waypoint(sell_waypoint, sell_cargo, body.Data.Cooldown.RemainingSeconds)
		return
	}
	time.Sleep(time.Duration(body.Data.Cooldown.RemainingSeconds) * time.Second)
	extract_resources()
}

func navigate_to_waypoint(waypoint string, on_arrival void_func, cooldown int) {
	body, err := CallApi[NavResponse]("/my/ships/"+mining_ship+"/navigate", "POST", []byte(`{"waypointSymbol": "`+waypoint+`"}`))
	if err != nil {
		return
	}
	time_to_arrival, _ := helpers.TimeDiffInSeconds(body.Data.Nav.Route.Arrival)
	// Determine timeout based on cooldown and time to arrival
	timeout := int(time_to_arrival)
	if time_to_arrival < int64(cooldown) {
		timeout += cooldown - int(time_to_arrival)
	}
	fmt.Println("Navigating to " + waypoint + ", arrives in " + strconv.Itoa(int(time_to_arrival)) + "s")
	time.Sleep(time.Duration(timeout) * time.Second)
	on_arrival()
}

func refuel() {
	CallApi[any]("/my/ships/"+mining_ship+"/refuel", "POST", nil)
	fmt.Println("Ship is fully refueled")
}

func sell_cargo() {
	CallApi[any]("/my/ships/"+mining_ship+"/dock", "POST", nil)
	for _, v := range cargo.Inventory {
		body := []byte(`{
			"symbol": "` + v.Symbol + `",
			"units": "` + strconv.Itoa(v.Units) + `"
		}`)
		_, err := CallApi[any]("/my/ships/"+mining_ship+"/sell", "POST", body)
		if err != nil {
			return
		}
		fmt.Println("Sold " + strconv.Itoa(v.Units) + " " + v.Symbol)
	}
	refuel()
	CallApi[any]("/my/ships/"+mining_ship+"/orbit", "POST", nil)
	navigate_to_waypoint(mine_waypoint, extract_resources, 0)
}
