package list

// WIP

/*
type SafeArrayList struct {
	sync.Once
	items []interface{}
	size  int
}

func (al *SafeArrayList) Add(item interface{}) {
	al.Do(func() {
		al.fixCapacity(al.size + 1)
		al.items[al.size] = item
		al.size += 1
	})

}

func (al *SafeArrayList) IndexOf(item interface{}) int {
	if item == nil {
		for i, v := range al.items {
			if v == nil {
				return i
			}
		}
	} else {
		for i, v := range al.items {
			if v == item {
				return i
			}
		}
	}
	return -1
}

func (al *SafeArrayList) fixCapacity(minCapacity int) {
	oldCapacity := len(al.items)
	if minCapacity > oldCapacity {
		oldData := al.items
		newCapacity := (oldCapacity*3)/2 + 1
		if newCapacity < minCapacity {
			newCapacity = minCapacity
		}
		newData := make([]interface{}, newCapacity)
		copy(newData, oldData)
		al.items = newData
	}
}

func (al *SafeArrayList) Size() int {
	return al.size
}

func (al *SafeArrayList) Contains(item interface{}) bool {
	return al.IndexOf(item) >= 0
}

func (al *SafeArrayList) Clear() {
	for i := 0; i < al.size; i++ {
		al.items[i] = nil
	}
	al.size = 0
}

func (al *SafeArrayList) Get(index int) interface{} {
	al.RangeCheck(index)
	return al.items[index]
}

func (al *SafeArrayList) Iterate(fun func(item interface{}, index int)) {
	al.Do(func() {
		for index, v := range al.items {
			fun(v, index)
		}
	})

}

func (al *SafeArrayList) RangeCheck(index int) {
	if index >= al.size {
		panic(fmt.Sprintf("SafeArrayList # IndexOutOfBound - Given index (%d) is >= of SafeArrayList size (%d)", index, al.size))
	}
}

func (al *SafeArrayList) Set(index int, item interface{}) {
	al.RangeCheck(index)
	al.items[index] = item
}

func (al *SafeArrayList) RemoveIndex(index int) (item interface{}) {
	al.RangeCheck(index)
	item = al.items[index]
	al.items[index] = nil
	return
}

func (al *SafeArrayList) Remove(item interface{}) {
	a := al.IndexOf(item)
	if a >= 0 {
		al.items[a] = nil
	}
}

func NewSafeArrayList() *SafeArrayList {
	arr := new(SafeArrayList)
	arr.size = 0
	arr.items = make([]interface{}, 10)
	return arr
}
*/
