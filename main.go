package main

import (
	"log"
	"time"

	"github.com/iimrudy/vUtils/encoding"
	"github.com/iimrudy/vUtils/list"
	"github.com/iimrudy/vUtils/optional"
)

func main() {

	var result []byte = encoding.Base64EncodeBytes([]byte("Encoding This String to Bytes using Base64"))
	print(result)

	/*os.Remove("./file1.txt")
	os.Remove("./file2.txt")
	os.Remove("./file3.delta")

	f1, err := os.OpenFile("./file1.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	f2, err := os.OpenFile("./file2.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	f3, err := os.OpenFile("./file3.delta", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	s1 := random.GenerateString(900)
	s2 := s1 + "----------" + random.GenerateString(900000)
	f1.WriteString(s1)
	f2.WriteString(s2)
	fmt.Println(f1.Close())
	fmt.Println(f2.Close())

	dlt, err := delta.GenerateDiffFormFilesByPath("./file1.txt", "./file2.txt")

	if err != nil {
		panic(err)
	}

	f3.Write(dlt.Bytes())
	fmt.Println(f3.Close())

	dltt, err := ddelta.Load(dlt.Bytes())
	if err != nil {
		panic(err)
	}

	applied, err := dltt.Apply([]byte(s1))
	if err != nil {
		panic(err)
	}

	if s2 == string(applied) {
		log.Println("UGUALEEE")
	} else {
		log.Println(string(applied))
	}*/
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

	log.Println("\n\n\nInit 10Size Timings :", TimerThisFunction(func() {
		lista := list.NewArrayList()
		doThings(lista)
	}), "\n\n\n ")

	log.Println("\n\n\nFixedSize Timings :", TimerThisFunction(func() {
		lista := list.NewArrayListSize(10000000)
		doThings(lista)
	}), "\n\n\n ")

	optional1 := optional.Of("")
	optional1.IfIsPresent(func(item interface{}) {
		print("The Item is present ", item)
	}).IfNotPresent(func() {
		print("The Item is NOT present ")
	})
}

func TimerThisFunction(f func()) int64 {
	start := time.Now()
	f()
	end := time.Now()
	return end.UnixMicro() - start.UnixMicro()
}

func doThings(lista *list.ArrayList) {
	//lista.Clear()
	for x := 0; x < 10000000; x++ {
		lista.Add(x)
	}
	log.Println("Size ", lista.Size())
	log.Println("Contains 9 ", lista.Contains(9))
	log.Println("Contains 99 ", lista.Contains(99))
	log.Println("IndexOF 9 - ", lista.IndexOf(9))
	log.Println("IndexOF 99 - ", lista.IndexOf(99))

	log.Println("REMOVING 9")
	lista.Remove(9)
	log.Println("Contains 9 ", lista.Contains(9))
	log.Println("IndexOF 9 - ", lista.IndexOf(9))

	log.Println("Clearing")
	lista.Clear()
	log.Println("Size ", lista.Size())
}
