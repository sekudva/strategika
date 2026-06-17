Agent Reputation: This is a table of action-to-action transitions. It looks like this:

| Partner → Me   | Cooperation | Defense |
| -------------- | ----------- | ------- |
| Share → Share  | ++          | ·       |
| Share → Hold   | −           | ·       |
| Share → Take   | −−          | ·       |
| Hold → Share   | ++          | −       |
| Hold → Hold    | −           | +       |
| Hold → Take    | −−          | +       |
| Take → Share   | ·           | −−      |
| Take → Hold    | ·           | −       |
| Take → Take    | ·           | ++      |

Reputation is calculated by the agent within its own memory to prevent memory from being cluttered with millions of opponent reputations.
However, the logic for actions based on reputation varies — agents simply request another agent's reputation as needed.

Modifier convention: it uses ModContext and only operates on what is explicitly allowed. This acts as a conditional restriction to keep experiments fair.

```
type ModContext struct {
	History  Rounds
	Round    int
	Strategy *Strategy
	OpRep    Reputation
}

STRATEGIKA PAYOFF MATRIX
var PayoffMatrix = [3][3]int{
	//		Share  Hold  Take
	/* Share*/ {+4, 0, -3},
	/* Hold */ {+1, 0, -1},
	/* Take */ {+7, 0, -2},
}
```

ORIGINAL (AXELROD)
Dilemma conditions: T > R > P > L and 2R > T + L

T (Temptation) = 5 — temptation to defect
R (Reward) = 3 — reward for mutual cooperation
P (Punishment) = 1 — punishment for mutual defection
L (Looser) = 0 — payoff for being naive

With a shift applied:
{+1, 0, -2}
{+0, 0, -0}
{+3, 0, -1}

Alternative matrix:
+4	0	-3
+1	0	-1
+6	+2	-2

T=6, R=4, P=-2, S=-3: 6>4>-2>-3 ✅
2R=8 > T+S=3 ✅
Hold/Hold=0 between 4 and -2 ✅
Hold vs Take=-1 between -3 and 0 ✅
Take vs Hold=+2 between 0 and 6 ✅

___

Репутации агентов: Это таблица соотношений действий к предыдущим действиям
Она выглядит так:

| Партнёр → Я   | Cooperation | Defense |
| ------------- | ----------- | ------- |
| Share → Share | ++          | ·       |
| Share → Hold  | −           | ·       |
| Share → Take  | −−          | ·       |
| Hold → Share  | ++          | −       |
| Hold → Hold   | −           | +       |
| Hold → Take   | −−          | +       |
| Take → Share  | ·           | −−      |
| Take → Hold   | ·           | −       |
| Take → Take   | ·           | ++      |

Просчитывается агентом в своей памяти, чтобы памяти агентов не засорялись миллионами репутаций противников
Но логика действий на репутацию разная, агенты просто получают репутацию другого агента по требованию

Условность модификатора - он использует ModContext и работает только с тем, с чем разрешено, это условное ограничение для честности экспериментов

```
type ModContext struct {
	History  Rounds
	Round    int
	Strategy *Strategy
	OpRep    Reputation
}

ВЕСОВАЯ МАТРИЦА Стратегики
var PayoffMatrix = 3 на 3 int{
	//		Share  Hold  Take
	/* Share*/ {+4, 0, -3},
	/* Hold */ {+1, 0, -1},
	/* Take */ {+7, 0, -2},
}
```


ОРИГИНАЛ
Условия дилеммы: T > R > P > L и 2R > T + L

T (Temptation) = 5 — соблазн предать
R (Reward) = 3 — награда за взаимную кооперацию
P (Punishment) = 1 — наказание за взаимное предательство
L (Looser) = 0 — плата за наивность

со сдвигом будет
{+1, 0, -2}
{+0, 0, -0}
{+3, 0, -1}

Альтернативная матрица 
+4	0	-3
+1	0	-1
+6	+2	-2
T=6, R=4, P=-2, S=-3: 6>4>-2>-3 ✅
2R=8 > T+S=3 ✅
Hold/Hold=0 между 4 и -2 ✅
Hold vs Take=-1 между -3 и 0 ✅
Take vs Hold=+2 между 0 и 6 ✅
