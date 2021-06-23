package interface_alloc

type A int

func (a A) getVal() int {
	return int(a)
}

type B struct {
	A int
}

func (b B) getVal() int {
	return b.A
}

type Valer interface {
	getVal() int
}

func doubleVal(v Valer) int {
	return v.getVal() * 2
}

const count = 1000

func ALoop() int {
	total := 0
	for i := 0; i < count; i++ {
		var a = A(i)
		total += doubleVal(a)
	}
	return total
}

func BLoop() int {
	total := 0
	for i := 0; i < count; i++ {
		b := B{
			A: i,
		}
		total += doubleVal(b)
	}
	return total
}
