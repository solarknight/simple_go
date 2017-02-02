package main

type demoA struct {
}

type demoB struct {
}

type demoC struct {
	a demoA
	b demoB
}

func (a *demoA) toStr() string {
	return "this is demoA"
}

func (b *demoB) toStr() string {
	return "this is demoB"
}

func conflictMethodName() {
	c := new(demoC)
	// c.toStr()
	c.a.toStr()
}
