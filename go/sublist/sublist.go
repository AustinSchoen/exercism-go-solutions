package sublist

type List []int
type Relation string

func (l List) isEqualTo(b []int) bool {
	if len(l) != len(b) {
		return false
	}

	for i, v := range l {
		if v != b[i] {
			return false
		}
	}

	return true
}

func (l List) isSuperlistOf(b []int) bool {
	if len(l) < len(b) {
		return false
	}

	// How many times to shift the sublist over
	shifts := len(l) - len(b)

	for i := 0; i <= shifts; i++ {
		lSubList := l[i : len(b)+i]
		if lSubList.isEqualTo(b) {
			return true
		}
	}

	return false
}

func (l List) isSubListOf(b []int) bool {
	if len(l) > len(b) {
		return false
	}

	// How many times to shift the sublist over
	shifts := len(b) - len(l)

	for i := 0; i <= shifts; i++ {
		bSubList := List(b[i : len(l)+i])
		if bSubList.isEqualTo(l) {
			return true
		}
	}

	return false
}

func Sublist(listOne []int, listTwo []int) Relation {
	list := List(listOne)

	if list.isEqualTo(listTwo) {
		return Relation("equal")
	}

	if list.isSuperlistOf(listTwo) {
		return Relation("superlist")
	}

	if list.isSubListOf(listTwo) {
		return Relation("sublist")
	}

	return Relation("unequal")
}
