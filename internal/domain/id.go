package domain

type AgID uint32

type IDGenerator struct {
	lastID AgID
	used   map[AgID]bool
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{
		lastID: 1,
		used:   make(map[AgID]bool),
	}
}

func (g *IDGenerator) Next() AgID {
	for g.used[g.lastID] {
		g.lastID++
	}
	id := g.lastID
	g.used[id] = true
	g.lastID++
	return id
}

func (g *IDGenerator) Request(id AgID) AgID {
	if g.used[id] {
		return g.Next()
	}
	g.used[id] = true
	return id
}

func (g *IDGenerator) Free(id AgID) {
	delete(g.used, id)
}
