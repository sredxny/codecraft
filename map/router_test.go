package mymap

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var routes []Route

type RouteStub struct {
	possibleRoutes []Route
}

func (routeStub *RouteStub) GetRoutes(pointA, pointB string) []Route {
	return routeStub.possibleRoutes
}

func init() {
	routes = []Route{
		{   path: "Path A",
			travelTime: time.Duration(25) * time.Minute,
		},
		{
			path:       "Path B",
			travelTime: time.Duration(30) * time.Minute,
		},
		{
			path:       "Path C",
			travelTime: time.Duration(45) * time.Minute,
		},
	}
}

func Test_Route(t *testing.T) {

	routesStub := RouteStub{possibleRoutes: routes}
	pointA := "Waterloo Station, Waterloo Rd, London SE1 8SW, United Kingdom"
	pointB := "87A Worship St, London EC2A 2BE, UK"
	possibleRoutes := routesStub.GetRoutes(pointA, pointB)
	chosenRoute := selectShortestRoute(possibleRoutes)

	t.Run("chose a route", func(t *testing.T) {
		assert.Contains(t, possibleRoutes, chosenRoute)
	})

	t.Run("choose the shortest route", func(t *testing.T) {
		assert.Equal(t, time.Duration(25)*time.Minute, chosenRoute.travelTime)
	})
}
