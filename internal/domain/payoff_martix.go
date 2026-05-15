package domain

var PayoffMatrix = [3][3]int{
	//		Share  Hold  Take
	/* Share*/ {+4, 0, -3},
	/* Hold */ {+1, 0, -1},
	/* Take */ {+7, 0, -2},
}

func Payoff(my, their Act) (int, int) {
	myPayoff := PayoffMatrix[my][their]
	theirPayoff := PayoffMatrix[their][my]
	return myPayoff, theirPayoff
}
