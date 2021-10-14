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
<br>

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
<br>

## Sha512 Encoding
<br>

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
<br>

## Base64 Encoding
<br>

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
<br>

## Base64 Decoding
<br>

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

//

## ArrayList

//

## Optimizations

//

# Optional

//

# Random

Used chars for random strings (safe and unsage):
- 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-
  
<br>

## Safe Randoms
<br>

### Safe Random String

```golang
import "github.com/iimrudy/vUtils/random"

func main() {
    length := 10 // string length
    var result string = random.GenerateSafeString(length)
	print(result)
}
```
<br>

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
<br>

### Unsafe Random String

```golang
import "github.com/iimrudy/vUtils/random"

func main() {
    length := 10 // string length
    var result string = random.GenerateString(length)
	print(result)
}
```
<br>

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


