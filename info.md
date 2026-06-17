
╔════════════════════════════════════════════════════════════════════════╗
║                    STRATEGIKA — QUICK REFERENCE GUIDE                  ║
╚════════════════════════════════════════════════════════════════════════╝

📌 MAIN MODES
────────────────────────────────────────────────────────────────────────
1) DUEL — one-on-one match. Full round-by-round logging.
2) ROUND ROBIN — round-robin tournament. Each plays each other (incl. self).
3) TRIAL — challenge: one leader against a selected group.
4) CIRCULAIRE — trial series: each leader against one group.
5) ARENA — all agents interact in a shared environment. Reputation forms.
6) ECOSYSTEM ARENA — arena with death threshold (elimination at low score).

🧠 AGENT ACTIONS
────────────────────────────────────────────────────────────────────────
Agent IS unit that carries strategy and has unique ID.

Actions:
SHARE (cooperate) — +4 vs Share, 0 vs Hold, -3 vs Take.
HOLD (neutral)    — +1 vs Share, 0 vs Hold, -1 vs Take.
TAKE (defect)     — +7 vs Share, 0 vs Hold, -2 vs Take.

⚙️ CONFIGURATION (on launch)
────────────────────────────────────────────────────────────────────────
• Rounds — how many turns each agent takes per simulation.
• Noise — probability of action distortion (0.0 – 1.0).

📊 LOGGING
────────────────────────────────────────────────────────────────────────
• DUEL and TRIAL: full round-by-round log (duel_log.txt).
• ROUND ROBIN, ARENA, ECOSYSTEM: aggregate stats ALL ROUNDS / 2. 
• ECOSYSTEM: additionally marks agent deaths.

🆘 NEED HELP?
────────────────────────────────────────────────────────────────────────
• README.md — project overview, mechanics, roadmap.
• Folder pkg/presets/ — all implemented strategies.
• Folder pkg/presets/mod/ — modifiers for advanced behaviors.
• sekudva@gmail.com — my mail address.

💡 STRATEGY EXAMPLES
────────────────────────────────────────────────────────────────────────
AlwaysShare   — always cooperates.
AlwaysTake    — always defects.
TitForTat     — copies opponent's last move.
Grudger       — defects forever after first betrayal.
Journalist    — doubles action (ignores Hold), sometimes provokes.
Patient       — patiently forgives up to 3 Take actions.
GroFman       — uses 2/7 probability to retaliate after a Take.
Pacifist      — never takes, only holds and shares.
EvilTFT       — TitForTat but starts with Take.
Random        — random action (33% Share / 34% Hold / 33% Take).

⁉️ I WANT SOMETHING THAT YOUR MENU DONT HAVE!
────────────────────────────────────────────────────────────────────────
Play with code as you like, it's open-source. I carefully try write 
as many comments as possible for other people to use my idea.
Feel free to experiment, that's how the best ideas are born.
If you want request something — email me at sekudva@gmail.com.