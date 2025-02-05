// статический
package main

import "fmt"

type MapState interface {
	Play(navi *Navi, spirit *TeamSpirit)
}

type Navi struct {
	bit       string
	Aleksib   string
	jL        string
	iM        string
	wonderful string
}

type TeamSpirit struct {
	magixx  string
	chopper string
	zontix  string
	donk    string
	shiro   string
}

type Ancient struct{}

func (a *Ancient) Play(navi *Navi, spirit *TeamSpirit) {
	fmt.Println("Teams Navi and Team Spirit decided to play on Ancient!")
	fmt.Printf("Navi players lineup: %s, %s, %s, %s, %s\n", navi.bit, navi.Aleksib, navi.jL, navi.iM, navi.wonderful)
	fmt.Printf("Team Spirit players: %s, %s, %s, %s, %s\n", spirit.magixx, spirit.chopper, spirit.zontix, spirit.donk, spirit.shiro)
	fmt.Println("We are starting the long-awaited CS2 Major!")
}

type Game struct {
	currentMap MapState
	navi       *Navi
	spirit     *TeamSpirit
}

func (g *Game) SetMap(state MapState) {
	g.currentMap = state
}

func (g *Game) Play() {
	g.currentMap.Play(g.navi, g.spirit)
}

func main() {
	navi := &Navi{
		bit:       "bit",
		Aleksib:   "Aleksib",
		jL:        "jL",
		iM:        "iM",
		wonderful: "wonderful",
	}

	spirit := &TeamSpirit{
		magixx:  "magixx",
		chopper: "chopper",
		zontix:  "zontix",
		donk:    "donk",
		shiro:   "shiro",
	}

	game := &Game{
		navi:   navi,
		spirit: spirit,
	}

	game.SetMap(&Ancient{})
	game.Play()
}
