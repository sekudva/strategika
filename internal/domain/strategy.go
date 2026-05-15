package domain

type MirrorMode int

const (
	MirrorDirect MirrorMode = iota // полное копирование

	MirrorDefense // Take → Hold
	MirrorBad     // Hold → Take

	MirrorCold // Share → Hold
	MirrorNice // Hold → Share

	MirrorOpp // Share ↔ Take

	MirrorSelf // копирование своего хода
)

var mirrorTable = map[MirrorMode]map[Act]Act{
	MirrorDirect: {
		Share: Share,
		Hold:  Hold,
		Take:  Take,
	},

	MirrorDefense: {
		Share: Share,
		Hold:  Hold,
		Take:  Hold, // Take → Hold RULE
	},

	MirrorBad: {
		Share: Share,
		Hold:  Take, // Hold → Take RULE
		Take:  Take,
	},

	MirrorCold: {
		Share: Hold, // Share → Hold RULE
		Hold:  Hold,
		Take:  Take,
	},

	MirrorNice: {
		Share: Share,
		Hold:  Share, // Hold → Share RULE
		Take:  Take,
	},

	MirrorOpp: {
		Share: Take, // Share ↔ Take RULE
		Hold:  Hold,
		Take:  Share, // Take ↔ Share RULE
	},
}

type RuleValue struct {
	Fix    *Act
	Prob   map[Act]float64
	Mirror *MirrorMode // nil == no MirrorMode
}

type Trigger struct {
	Act      Act
	Count    int
	Reaction RuleValue // ответ на триггер
}

type Strategy struct {
	Neutral RuleValue // нейтральное состояние
	Trigger *Trigger
	State   map[string]int // счетчик стратегии
}
