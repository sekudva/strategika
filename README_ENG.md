# STRATEGIKA
## Legacy of Axelrod's Tournament

Author: Selena Kudinova
WARNING: Commentaries in that project often written in Russian(my native language), please use translator if you need clarity!

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

___
### Personal. From the Author
This project is deeply personal. Game theory is something I would love to pursue seriously, but I only understood that after the idea for this project was born. Before that, life felt like a fog.

I have ADHD, and socialization has always been difficult. Since childhood, I made decisions based on other people's perceptions — it seemed to me that a person couldn't be wrong about me. In my head, I would build a hypothetical trajectory of the conversation and choose the words I thought would satisfy the other person. If they were assertive — I was assertive. If they were gentle — I was gentler still. But beyond that, I also read the environment: if someone forbade themselves from showing emotion, I wanted to help them with that, and I did it deliberately.

Once I was told not to be sad, and I stopped showing sadness entirely — until people started getting angry at me for my constant cheerfulness. Since childhood I was told never to lie, and I lied for the first time at 17 — because everyone around me was lying and embellishing, and I didn't understand why that was rewarded.

The diversity of people's characters and intentions both inspires and exhausts me. I've always wanted a tool that would give me abstract instructions for behavior in different situations. I studied animal behavior, replayed my favorite games choosing the "evil" path, argued aimlessly, and lied to friends. All of it came from a desire to experiment on my own imaginary proving ground. Because no one could truly understand my inner tossing. No one could experience this complete absence of a coherent identity — and this temptation to understand all intentions in the world without the lenses of my own perception.

The nature of strategy tournament simulation turned out to be incredibly close to me. A tournament actually reflects the way I think. Thoughts and decisions that flash by in seconds during a conversation, and then linger for hours and days during reflection, go through similar trials. It's strange to find my own reflection not in one of the strategies, but in the tournament itself — and I want to cherish this feeling of unity with the concept, continuing and developing it as a tool for decision-making simulation.