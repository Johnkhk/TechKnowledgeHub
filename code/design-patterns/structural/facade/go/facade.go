package facade

import "fmt"

// TV struct
type TV struct{}

func (t *TV) TurnOn() {
	fmt.Println("TV is ON")
}

func (t *TV) TurnOff() {
	fmt.Println("TV is OFF")
}

// SoundSystem struct
type SoundSystem struct{}

func (s *SoundSystem) TurnOn() {
	fmt.Println("Sound System is ON")
}

func (s *SoundSystem) TurnOff() {
	fmt.Println("Sound System is OFF")
}

// GamingConsole struct
type GamingConsole struct{}

func (g *GamingConsole) TurnOn() {
	fmt.Println("Gaming Console is ON")
}

func (g *GamingConsole) TurnOff() {
	fmt.Println("Gaming Console is OFF")
}

// HomeEntertainmentFacade struct
type HomeEntertainmentFacade struct {
	tv            *TV
	soundSystem   *SoundSystem
	gamingConsole *GamingConsole
}

func NewHomeEntertainmentFacade() *HomeEntertainmentFacade {
	return &HomeEntertainmentFacade{
		tv:            &TV{},
		soundSystem:   &SoundSystem{},
		gamingConsole: &GamingConsole{},
	}
}

func (h *HomeEntertainmentFacade) PlayGame() {
	h.gamingConsole.TurnOn()
	h.tv.TurnOn()
	h.soundSystem.TurnOn()
}

func (h *HomeEntertainmentFacade) WatchMovie() {
	h.tv.TurnOn()
	h.soundSystem.TurnOn()
}
