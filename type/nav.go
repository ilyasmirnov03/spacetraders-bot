package _type

type Consumed struct {
	Amount    int    `json:"amount"`
	Timestamp string `json:"timestamp"`
}

type Fuel struct {
	Current  int      `json:"current"`
	Capacity int      `json:"capacity"`
	Consumed Consumed `json:"consumed"`
}

type Destination struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type Route struct {
	Destination   Destination `json:"destination"`
	Origin        Destination `json:"origin"`
	DepartureTime string      `json:"departureTime"`
	Arrival       string      `json:"arrival"`
}

type Nav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          Route  `json:"route"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type NavResponse struct {
	Fuel Fuel
	Nav  Nav
}
