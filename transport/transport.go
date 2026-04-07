package transport

import "fmt"

type Transport struct {
	Id    uint64
	Speed float64
}

type Auto struct {
	Transport
	Wheels uint64
	Weight float64
}

type Special struct {
	Auto
	Kind string
}

func NewSpecial(id uint64, speed float64, wheels uint64, weight float64, kind string) *Special {
	return &Special{Auto{Transport{id, speed}, wheels, weight}, kind}
}

type Action interface {
	Voice()
	Gas()
	Autopilot()
	StartUP()
	TurnOFF()
	GetSpeed() float64
	GetKind() string
}

func (p *Special) Voice()            { fmt.Printf("%s: Бип бип!\n", p.Kind) }
func (p *Special) Gas()              { fmt.Printf("%s: Врум врум!\n", p.Kind) }
func (p *Special) Autopilot()        { fmt.Printf("%s: Автопилот включен!\n", p.Kind) }
func (p *Special) StartUP()          { fmt.Printf("%s: Двигатель заведен!\n", p.Kind) }
func (p *Special) TurnOFF()          { fmt.Printf("%s: Двигатель заглушен!\n", p.Kind) }
func (p *Special) GetSpeed() float64 { return p.Speed }
func (p *Special) GetKind() string   { return p.Kind }
