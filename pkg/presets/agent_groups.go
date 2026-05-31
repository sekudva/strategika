// pkg/presets/agent_groups.go

package presets

import (
	"github.com/sekudva/strategika/internal/domain"
)

// ClassicStrategies возвращает набор классических стратегий(с новой логикой) из турнира Аксельрода
// которые не используют модификаторы
func ClassicStrategies() []*domain.Agent {
	return []*domain.Agent{
		TitForTat(),
		ForgivingTFT(),
		AlwaysShare(),
		AlwaysTake(),
		Joss(),
		TitFor2Tats(),
		Grudger(),
		EvilTFT(),
		Random(),
	}
}

// AggressiveGroup возвращает стратегии, склонные к атаке
func NonClassicGroup() []*domain.Agent {
	return []*domain.Agent{
		AlwaysHold(),
		Masochist(),
		Sadist(),
		Pacifist(),
		Bully(),
		Melting(),
		EvilCyclist(),
		GoodCyclist(),
		ColdCyclist(),
		WarmCyclist(),
		ColdGhost(),
		Crazy(),
		Unfriendly(),
		Irreconcilable(),
		StrictTeacher(),
		Patient(),
	}
}

// Evil Group -> Evil Win преположительно в такой подборке победят злые стратегии
// примерно 20/10/70 добрые/нейтральные/злые
func EvilDominates() []*domain.Agent {
	return []*domain.Agent{
		// ДОБРЫЕ (5 шт.) — оставляем самых беззащитных
		AlwaysShare(),
		Pacifist(), //
		//Patient(),        //
		//WarmCyclist(), //
		//Irreconcilable(), //
		//ForgivingTFT(), //

		// НЕЙТРАЛЬНЫЕ (2-3 шт.) — оставляем тех, кто не мешает злым
		ColdCyclist(), // почти всегда Hold
		Unfriendly(),  // почти всегда Hold

		// ЗЛЫЕ (17-18 шт.) — повторяем агрессоров
		AlwaysTake(),

		EvilTFT(), // начинаем с Take
		//EvilTFT(), // повтор
		//EvilTFT(), // повтор

		Sadist(), // Hold → Take
		Sadist(), // повтор
		Sadist(), // повтор

		Bully(), // Hold → Take
		Bully(), // повтор
		Bully(), // повтор

		Crazy(), // 50/50 Share/Take

		EvilCyclist(), // предсказуемый
	}
}
