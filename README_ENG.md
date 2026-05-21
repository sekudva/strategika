WARNING: Commentaries in my project often written in Russian, please use translator if you need clarity

# STRATEGIKA
## Legacy of Axelrod's Tournament

Author: Selena Kudinova

### History of the Strategy Tournament

Robert Axelrod, an American political scientist, published "The Evolution of Cooperation" in 1984.
To naturally observe interactions between different strategies, Axelrod created a tournament in 1979 based on repeated iterations of the "Prisoner's Dilemma" and invited fellow researchers from game theory to create their own strategies, based on 200 duel rounds.
Later, in 1980, hobbyists joined the strategies created by professionals in their fields, and the number of rounds became pseudo-random within the range of approximately 200 rounds.

### Strategika — Heir to the Idea

The "Strategika" project is built upon Axelrod's tournaments with several modifications and extensions, written in Go using composition and interfaces.
The goal of the project is not so much to run tournaments as to observe different behavioral dynamics under different conditions.
In addition to classic one-on-one duels, there are "trial" modes where one strategy is tested against a group of strategies, and "arena" modes where strategies coexist simultaneously. The existence of such arenas allows for the introduction of reputation for each strategy-agent, unspoken coalitions, and an extension of the original "Prisoner's Dilemma" rules by introducing a neutral action, "Abstain", alongside "Cooperate" and "Defect".
It is possible to select specific strategies participating in the competition, since a universal winning strategy cannot exist and everything depends on external conditions (that is, who the opponent strategies are).

### Mechanics

**Three actions instead of two.** In the original: "Cooperate" and "Defect".
Strategika introduces a third action, "Abstain" — a neutral action
that changes the dynamics: you can defend without attacking in response.

**Three tournament formats:**
- Duel (1 × 1) — Axelrod's classic
- Trial (1 × Group) — one strategy against a group
- Arena (Many × Many) — all against all with reputation and unspoken coalitions

**Reputation.** In the Arena, agents observe others' interactions and form
reputations of each other: who is a predator, who is prey, who is a defender.
Inside each agent, a numerical ratio of certain action types is stored.

**Strategies from three rules.** Each strategy is described by a neutral
state, a trigger threshold, and a reaction. Values can be fixed or probabilistic.
Complex strategies are assembled by adding modifiers.

### Planned Extensions

- Ecosystem with mutations and survival
- Web-based strategy builder