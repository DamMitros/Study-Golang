package models

type Stop struct {
	StopId       int    `json:"stopId"`
	StopCode     string `json:"stopCode"`
	StopName     string `json:"stopName"`
	StopShortName string `json:"stopShortName"`
	StopDesc     string `json:"stopDesc"`
	StopLat      float64 `json:"stopLat"`
	StopLon      float64 `json:"stopLon"`
	ZoneId       int `json:"zoneId"`
}

type StopsResponse struct {
	LastUpdate string `json:"lastUpdate"`
	Stops      []Stop `json:"stops"`
}

type Departure struct {
	Id                    string    `json:"id"`
	Delay                 int    `json:"delayInSeconds"`
	EstimatedTime         string `json:"estimatedTime"`
	HeadsignText          string `json:"headsign"`
	RouteId               int    `json:"routeId"`
	RouteShortName        string `json:"routeShortName"`
	ScheduledTripStartTime string `json:"scheduledTripStartTime"`
	TripId                int    `json:"tripId"`
	StatusMessage         string `json:"status"`
	Theoretically         string `json:"theoreticalTime"`
	Timestamp             string  `json:"timestamp"`
	Trip 									int `json:"trip"`	
	VehicleCode           int `json:"vehicleCode"`
	VehicleId             int `json:"vehicleId"`
	VehicleService        string `json:"vehicleService"`	
}

type DeparturesResponse struct {
	LastUpdate  string      `json:"lastUpdate"`
	Departures  []Departure `json:"departures"`
}

type StopTime struct {
	RouteId       int    `json:"routeId"`
	TripId 				int    `json:"tripId"`
	ArrivalTime 	string `json:"arrivalTime"`
	DepartureTime string `json:"departureTime"`
	StopId        int    `json:"stopId"`
	StopSequence  int    `json:"stopSequence"`
}

type StopTimesResponse struct {
	LastUpdate string     `json:"lastUpdate"`
	StopTimes  []StopTime `json:"stopTimes"`
}

type Route struct {
	RouteId int `json:"routeId"`
	RouteShortName string `json:"routeShortName"`
	RouteLongName string `json:"routeLongName"`
	RouteType string `json:"routeType"`
	AgencyId int `json:"agencyId"`
}

type RoutesResponse struct {
	LastUpdate string `json:"lastUpdate"`
	Routes []Route `json:"routes"`
}