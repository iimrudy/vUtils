package optional

import "github.com/iimrudy/vUtils/utils"

type Optional struct {
	item interface{}
}

func (o *Optional) IsPresent() bool {
	return o.item != nil
}

func (o *Optional) Get() interface{} {
	return o.item
}

func (o *Optional) IfIsPresent(f func(item interface{})) *Optional {
	if o.IsPresent() {
		f(o.item)
	}
	return o
}

func (o *Optional) IfNotPresent(f func()) *Optional {
	if !o.IsPresent() {
		f()
	}
	return o
}

func OfNullable(item interface{}) *Optional {
	x := new(Optional)
	x.item = item
	return x
}

func Of(item interface{}) *Optional {
	utils.PanicIfNull(item)
	x := new(Optional)
	x.item = item
	return x
}
