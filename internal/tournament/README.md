Duel, Trial, and Arena are three launch modes for the Simulation Config, sharing a single Run method.

The interaction type is defined by the configuration settings and the number of pairs.

```
type SimConfig struct {
	Rounds
	Noise
	Pairs (agent indices)
	RNG (determinism for a single simulation)
	Logger
	Planned: evolutionary parameters
}
```

The main difference between a Round-Robin tournament and the Arena lies in the parallelization of moves. In a Round-Robin tournament, agents compete sequentially, pair by pair, across many independent simulations. In the Arena, decisions are made in parallel for all agents within a single, unified simulation.

This parallel structure better reflects overall and evolutionary dynamics, enables the formation of reputations, and can lead to the emergence of coalitions.

Round-Robin is a cyclic scheduling algorithm where elements or tasks are processed one after another, in turn.

ON THE VICTORY OF EVIL STRATEGIES

It is impossible to determine an exact ratio of good to evil strategies that guarantees when evil ones will dominate. This is because strategies vary widely — that's the core point of running a simulation at all.

A rough estimate, however, suggests that for evil strategies to win, the ratio needs to be roughly 80/20 in favor of evil.

There are three logger types: Full (all information), Aggregated (with a configurable interval), and Silent.

___

Дуэль, Испытание и Арена - способы запуска Конфигурации Симуляции, один метод Run.

Тип взаимодействия определяется из настройки конфигурации и количестве пар.

```
type SimConfig struct {
	Кол-во раундов
	Шум
	Пары агентов по индексам 
	Детерминирование случайности для одной симуляции
	Логгер
	ПЛАНИРУЕТСЯ эволюционные параметры
}
```

Разница ТУРНИРА ДУЭЛЕЙ и АРЕНЫ в параллельности ходов
В турнире дуэлей агенты соревнуются друг с другом по очереди, последовательно по парам, много отдельных симуляций,
когда в Арене принимаются решения параллельно для всех агентов, в единой симуляции,
следствие - лучшее отображение общей и эволюционной динамики, создание репутаций и, возможно, коалиций

Турнир дуэлей - это Раунд-робин, алгоритм циклического распределения, при котором элементы или задачи обрабатываются по очереди, один за другим

ПОБЕДА ЗЛЫХ СТРАТЕГИЙ

НЕЛЬЗЯ точно определить соотношение добрые к злым чтобы понять в какой момент злые будут лидировать
Потому что стратегии разные, в этом прикол симуляции вообще

НО примерный расчет, чтобы злые стратегии поебдили - это 80/20 злых к добрым

Три типа логгера: полный(вся информация), агреггированный(можно настроить шаг) и тихий.