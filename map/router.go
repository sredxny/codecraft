package mymap

import "time"

type Route struct {
	path       string
	travelTime time.Duration
}

type RouteProvider interface {
	GetRoutes(pointA, pointB string) []Route
}

func selectShortestRoute(routes []Route) Route {
	shortestTravelTime := time.Duration(9999999) * time.Minute
	var shortestRoute Route

	for index := range routes {
		if routes[index].travelTime < shortestTravelTime {
			shortestRoute = routes[index]
			shortestTravelTime = shortestRoute.travelTime
		}
	}
	return shortestRoute
}

func (r *Route) GetPath() string {
	return r.path
}
