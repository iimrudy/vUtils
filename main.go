package main

import (
	"log"
	"time"

	"github.com/iimrudy/vUtils/list"
)

func main() {

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

	log.Println("Init 10Size Timings :", TimerThisFunction(func() {
		lista := list.NewArrayList()
		doThings(lista)
	}))

	log.Println("FixedSize Timings :", TimerThisFunction(func() {
		lista := list.NewArrayListSize(100)
		doThings(lista)
	}))

	/*optional.Of(lista).IfIsPresent(func(item interface{}) {

	})*/

}

func TimerThisFunction(f func()) int64 {
	start := time.Now()
	f()
	end := time.Now()
	return end.UnixMicro() - start.UnixMicro()
}

func doThings(lista *list.ArrayList) {
	//lista.Clear()
	for x := 0; x < 90; x++ {
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
