package list

import (
	"fmt"
)

type ArrayList struct {
	items []interface{}
	size  int
}

func (al *ArrayList) Add(item interface{}) {
	al.fixCapacity(al.size + 1)
	al.items[al.size] = item
	al.size += 1
}

func (al *ArrayList) IndexOf(item interface{}) int {
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

func (al *ArrayList) fixCapacity(minCapacity int) {
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

func (al *ArrayList) Size() int {
	return al.size
}

func (al *ArrayList) Contains(item interface{}) bool {
	return al.IndexOf(item) >= 0
}

func (al *ArrayList) Clear() {
	for i := 0; i < al.size; i++ {
		al.items[i] = nil
	}
	al.size = 0
}

func (al *ArrayList) Get(index int) interface{} {
	al.RangeCheck(index)
	return al.items[index]
}

func (al *ArrayList) Iterate(fun func(item interface{}, index int)) {
	size := al.Size()
	for index, v := range al.items {
		if size != al.Size() {
			panic("ArrayList # IndexOutOfBound - Concurrent array access.")
		}
		fun(v, index)
	}
}

func (al *ArrayList) RangeCheck(index int) {
	if index >= al.size {
		panic(fmt.Sprintf("ArrayList # IndexOutOfBound - Given index (%d) is >= of arraylist size (%d)", index, al.size))
	}
}

func (al *ArrayList) Set(index int, item interface{}) {
	al.RangeCheck(index)
	al.items[index] = item
}

func (al *ArrayList) RemoveIndex(index int) (item interface{}) {
	al.RangeCheck(index)
	item = al.items[index]
	al.items[index] = nil
	return
}

func (al *ArrayList) Remove(item interface{}) {
	a := al.IndexOf(item)
	if a >= 0 {
		al.items[a] = nil
	}
}

func NewArrayList() *ArrayList {
	arr := new(ArrayList)
	arr.size = 0
	arr.items = make([]interface{}, 10)
	return arr
}

func NewArrayListSize(size int) *ArrayList {
	arr := new(ArrayList)
	arr.size = 0
	arr.items = make([]interface{}, size)
	return arr
}
