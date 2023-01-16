package listops

type IntList []int

type binFunc = func(int, int) int
type predFunc = func(int) bool
type unaryFunc = func(int) int

//Determines size using nil pointer exception
func (i IntList) Length() (l int) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	for c := 0; ; c++ {
		_ = i[c]
		l++
	}
}

// Apply fold operation starting from right
func (i IntList) Foldr(fn binFunc, init int) int {
	current := init

	for x := i.Length(); x > 0; x-- {
		current = fn(i[x-1], current)
	}

	return current
}

// Apply fold operation starting from left
func (i IntList) Foldl(fn binFunc, init int) int {
	current := init

	for x := 0; x < i.Length(); x++ {
		current = fn(current, i[x])
	}

	return current
}

// Add IntList to IntList
func (i IntList) Append(newItem IntList) IntList {
	newList := make(IntList, i.Length())

	for x := 0; x < i.Length(); x++ {
		newList[x] = i[x]
	}

	for x := 0; x < newItem.Length(); x++ {
		newList = newList.Push(newItem[x])
	}

	return newList
}

//Concatenates multiple IntList's into one
func (i IntList) Concat(lists []IntList) IntList {
	for x := range lists {
		i = i.Append(lists[x])
	}
	return i
}

//Filters list based on given function
func (i IntList) Filter(fn predFunc) IntList {
	newList := make(IntList, 0)

	for x := 0; x < i.Length(); x++ {
		if fn(i[x]) {
			newList = newList.Push(i[x])
		}
	}

	return newList
}

//Reverses list
func (i IntList) Reverse() IntList {
	newList := make(IntList, 0)

	for x := i.Length(); x > 0; x-- {
		newList = newList.Push(i[x-1])
	}

	return newList
}

// Applies function to each list item & returns new list
func (i IntList) Map(fn unaryFunc) IntList {
	newList := make(IntList, 0)

	for x := 0; x < i.Length(); x++ {
		newList = newList.Push(fn(i[x]))
	}

	return newList
}

// The tests didn't ask for this function but for some reason Append()
// is for adding IntLists (like Concat() should be) instead of adding
// new integers to the current IntList. We need the below to facilitate
// adding new items, I've named it Push after the Ruby Array function
func (i IntList) Push(newItem int) IntList {
	newSize := i.Length() + 1
	n := make(IntList, newSize)

	for x := 0; x < i.Length(); x++ {
		n[x] = i[x]
	}

	n[newSize-1] = newItem

	return n
}
