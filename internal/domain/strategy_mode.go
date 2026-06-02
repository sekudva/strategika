package domain

// Правило внутри стратегии, которое предполагает повторение ходов противника
type MirrorMode int

const (
	MirrorDirect MirrorMode = iota // полное копирование

	MirrorDefense // Take → Hold
	MirrorBad     // Hold → Take

	MirrorCold // Share → Hold
	MirrorNice // Hold → Share

	MirrorOpp // Share ↔ Take

	MirrorSelf // копирование своего хода

	// Можно расширить опционал, но чисто теоретически это бессмысленно
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

// тип как часто срабатывает триггер
type TriggerMode int

// почти бессмысленно, по умолчанию самый логичный вариант
// можно придумать альтернативу или использовать в комбинации со сложными модификторами
const (
	TriggerStreakAfter TriggerMode = iota // После N раз (пока не закончится серия)
	TriggerEvery                          // каждый N-й подряд в серии (n=1 = на каждый)
	TriggerExactly                        // ровно на N-й, потом пропускать ответ на триггер
)

type TriggerTest func(history Rounds, act Act, n int) bool

var triggerTable = map[TriggerMode]TriggerTest{

	TriggerStreakAfter: func(h Rounds, act Act, n int) bool {
		return h.streak(act) >= n
	},

	TriggerEvery: func(h Rounds, act Act, n int) bool {
		s := h.streak(act)
		return s >= n && s%n == 0
	},

	TriggerExactly: func(h Rounds, act Act, n int) bool {
		return h.streak(act) == n
	},
}
