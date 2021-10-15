# vUtils

A collection of utils for golang

## Implementations

<ol>

<li><a href="#delta-of-binary-data">Delta Of Binary Data</a></li>

<li><a href="#encoding"> Encoding </li>


<li><a href="#arraylist">Files Utilities</a></li>

<li><a href="#lists">Lists</a></li>
<ul>
    <li> <a href="#arraylist">Array List</a>
        <ul>
            <li> <a href="#arraylist"> How To Use </a> </li>
            <li> <a href="#optimizations"> Optimization </a> </li>
        </ul>
    </li>
    <!--<li> GoRoutineSafe ArrayList
        <ul>
            <li> How To Use </li>
            <li> Optimization </li>
        </ul>
    </li>-->
</ul>


<li><a href="#optional"> Optional </li>

<li><a href="#random"> Random </li>
<ul>
    <li> <a href="#safe-randoms">Safe Randoms</a>
        <ul>
            <li> <a href="#safe-random-string"> String </a> </li>
            <li> <a href="#safe-random-integer"> Integer </a> </li>
        </ul>
    </li>
    <li> <a href="#unsafe-randoms">Unsafe Randoms</a>
        <ul>
            <li> <a href="#unsafe-random-string"> String </a> </li>
            <li> <a href="#unsafe-random-integer"> Integer </a> </li>
        </ul>
    </li>
</ul>

</ol>

# Delta Of Binary Data

//

# Encoding

## Sha256 Encoding

### Bytes Encoding

```golang
import "github.com/iimrudy/vUtils/encoding"

func main() {
    var result []byte = encoding.Sha256EncodeBytes([]byte("Encoding This String to Bytes using sha256"))
	print(result)
}
```

### String Encoding

```golang
import "github.com/iimrudy/vUtils/encoding"

func main() {
    var result string = encoding.Sha256EncodeString("Encoding This String to String using sha256")
	print(result)
}
```

---

## Sha512 Encoding

### Bytes Encoding

```golang
import "github.com/iimrudy/vUtils/encoding"

func main() {
    var result []byte = encoding.Sha512EncodeBytes([]byte("Encoding This String to Bytes using sha521"))
	print(result)
}
```

### String Encoding

```golang
import "github.com/iimrudy/vUtils/encoding"

func main() {
    var result string = encoding.Sha512EncodeString("Encoding This String to String using sha512")
	print(result)
}
```

---

## Base64 Encoding

### Bytes Encoding

```golang
import "github.com/iimrudy/vUtils/encoding"

func main() {
    var result []byte = encoding.Base64EncodeBytes([]byte("Encoding This String to Bytes using Base64"))
	print(result)
}
```

### String Encoding

```golang
import "github.com/iimrudy/vUtils/encoding"

func main() {
    var result string = encoding.Base64EncodeBytes("Encoding This String to String using Base64")
	print(result)
}
```

---

## Base64 Decoding

### Bytes Decoding

```golang
import "github.com/iimrudy/vUtils/encoding"

func main() {
    var result []byte = encoding.Base64DecodeBytes([]byte("RGVjb2RpbmcgVGhpcyBTdHJpbmcgdG8gQnl0ZXMgdXNpbmcgQmFzZTY0"))
	print(result)
}
```

### String Decoding

```golang
import "github.com/iimrudy/vUtils/encoding"

func main() {
    var result string = encoding.Base64DecodeString("RGVjb2RpbmcgVGhpcyBTdHJpbmcgdG8gU3RyaW5nIHVzaW5nIEJhc2U2NA")
	print(result)
}
```


# Files Utilities

//

# Lists

## ArrayList

Add an item to the arraylist
```golang
func (al *ArrayList) Add(item interface{}) {}
```

Get the index of an item
```golang
func (al *ArrayList) IndexOf(item interface{}) int {}
```

Get the size of the arraylist
```golang
func (al *ArrayList) Size() int {}
```

Check if a list contains an item
```golang
func (al *ArrayList) Contains(item interface{}) bool {}
```

Clear the list
```golang
func (al *ArrayList) Clear() {}
```

Get an item given the index (Can panic if index is out of bound)
```golang
func (al *ArrayList) Get(index int) interface{} {}
```

Iterate over the array given an Iterator Function
```golang
func (al *ArrayList) Iterate(fun func(item interface{}, index int)) {}
```

Panic if the given index is OutOfBound
```golang
func (al *ArrayList) RangeCheck(index int) {}
```

Set an item into the arrayList in the given index (Can panic if index is out of bound)
```golang
func (al *ArrayList) Set(index int, item interface{}) {}
```

Remove an item in the given index (Can panic if index is out of bound)
```golang
func (al *ArrayList) RemoveIndex(index int) (item interface{}) {}
```

Remove an item from the arraylist
```golang
func (al *ArrayList) Remove(item interface{}) {}
```

Example
```golang
import "github.com/iimrudy/vUtils/list"

func main() {
    // create a new arraylist
    var theList *list.ArrayList = list.NewArrayList()
    element1 := "An Item"
    element2 := "Another Item"
    // Add an item
    theList.Add(element1)
    theList.Add(element2)

    // Remove an item
    theList.Remove(element1)

    print(theList.Contains(element1))

    // Clear The List
    theList.Clear()
}
```

## Optimizations

If you know how many items you can fit into the arrayList you can create it with an initial size.

```golang
import "github.com/iimrudy/vUtils/list"

func main() {
    // create a new arraylist
    var initialSize int = 100
    var theList *list.ArrayList = list.NewArrayListSize(initialSize)
    
    // do stuff with the arrayList
}
```

This improve a lot the timings when adding new elements.

Example: Adding 10.000.000 Items on the arrayList

- List created with list.NewArrayList() timings: 783071 microSeconds

- List created with list.NewArrayListSize(10000000) timings: 315598 microSeconds

You can see that creating the list with a known starting size is more efficient

# Optional

As java documentation say 'A container object which may or may not contain a non-null value.'

Check if there is an item in the OptionalContainer
```golang
func (o *Optional) IsPresent() bool {}
```

Get the item in the OptionalContainer
```golang
func (o *Optional) Get() interface{} {}
```

Call the function if the item is presente into the OptionalContainer (Chainable)
```golang
func (o *Optional) IfIsPresent(f func(item interface{})) *Optional {}
```
Call the function if the item is NOT presente into the OptionalContainer (Chainable)
```golang
func (o *Optional) IfNotPresent(f func()) *Optional {}
```

```golang
import "github.com/iimrudy/vUtils/optional"

func main() {
    // create a new arraylist
    var optional1 *optional.Optional = optional.Of(nil) // this will panic, cannot provide null items
    var optional2 *optional.Optional = optional.OfNullable(nil) // this is fine, null-Able, can contain a null var.

    optional1.IfIsPresent(func(item interface{}) {
        print("The Item is present ", item)
	})

    optional1.IfNotPresent(func() {
        print("The Item is NOT present")
	})

    // Or you can chain the functions
    optional1.IfIsPresent(func(item interface{}) {
		print("The Item is present ", item)
	}).IfNotPresent(func() { // here is the chained function.
		print("The Item is NOT present")
	})
}
```


# Random

Used chars for random strings (safe and unsage):
- 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-
  
<br>

## Safe Randoms

### Safe Random String

```golang
import "github.com/iimrudy/vUtils/random"

func main() {
    length := 10 // string length
    var result string = random.GenerateSafeString(length)
	print(result)
}
```

### Safe Random Integer
```golang
import "github.com/iimrudy/vUtils/random"

func main() {
    var result1 int = random.RandomIntSafe(0, 100) // generate an int beetwen 0 and 100
	print(result1)

    var result2 int = random.RandomInt64Safe(0, 100) // generate an int64 number beetwen 0 and 100
	print(result2)
}
```
---

## Unsafe Randoms

### Unsafe Random String

```golang
import "github.com/iimrudy/vUtils/random"

func main() {
    length := 10 // string length
    var result string = random.GenerateString(length)
	print(result)
}
```

### Unsafe Random Integer
```golang
import "github.com/iimrudy/vUtils/random"

func main() {
    var result1 int = random.RandomInt(0, 100) // generate an int beetwen 0 and 100
	print(result1)

    var result2 int = random.RandomInt64(0, 100) // generate an int64 number beetwen 0 and 100
	print(result2)
}
```


