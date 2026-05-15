package domain

type MirrorMode int

const (
	MirrorDirect MirrorMode = iota

	MirrorDefense // Take → Hold
	MirrorBad     // Hold → Take

	MirrorCold // Share → Hold
	MirrorNice // Hold → Share

	MirrorOpp // Share ↔ Take

	MirrorSelf // копирование своего хода
)

type RuleValue struct {
	Fix    *Act
	Prob   map[Act]float64
	Mirror *MirrorMode
}

type Strategy struct {
}
