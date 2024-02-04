package main

import (
  "fmt"
  "errors"
)

type GrassHopper struct {
  Place int
  Target int
}

type Jumper interface {
    WhereAmI() int
    Jump() (int, error)
}

func (g *GrassHopper) WhereAmI() int {
  return g.Place
}

func (g *GrassHopper) Jump() (int, error) {
  if g.Place == g.Target {
    return 0, errors.New("Error")
  }

  delta := g.Target - g.Place
  right := true
  if delta < 0 {
    delta = -delta
    right = false
  }

  n := 5
  if (delta < n) {
    n = delta
  }

  if right {
    g.Place += n
  } else {
    g.Place -= n
  }

  return g.Place, nil
}

func PlaceJumper(place, target int) Jumper {
  g := new(GrassHopper)
  g.Place = place
  g.Target = target
  return g
}

const (
	place  = 0
	target = 3
)

func main() {
	g := PlaceJumper(place, target)
	fmt.Println(g.WhereAmI())
	for {
		currPlace, err := g.Jump()
		if err != nil {
			break
		}
		fmt.Println(currPlace)
	}
}
