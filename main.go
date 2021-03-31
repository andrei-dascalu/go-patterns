package main

import (
	"fmt"

	calcnew "github.com/austerus/golang-play/src/calc-new"
	calcnewer "github.com/austerus/golang-play/src/calc-newer"
	calcold "github.com/austerus/golang-play/src/calc-old"
	"github.com/austerus/golang-play/src/foptions"
)

func main() {
	fmt.Println("Test")

	calcOld := calcold.Calculator{}

	calcOld.Do(calcold.OP_ADD, 5.00)
	calcOld.Do(calcold.OP_SUB, 1.00)
	final := calcOld.Do(calcold.OP_MUL, 2.00)

	fmt.Printf("Value: %v\n", final)

	calcNew := calcnew.Calculator{}

	calcNew.Do(calcnew.Add, 5.00)
	calcNew.Do(calcnew.Sub, 1.00)
	calcNew.Do(calcnew.Mul, 2.00)
	last := calcNew.Do(calcnew.Sqrt, 0)

	fmt.Printf("Value: %v\n", last)

	calcNewer := calcnewer.Calculator{}

	calcNewer.Do(calcnewer.Add(5.00))
	calcNewer.Do(calcnewer.Sub(1.00))
	calcNewer.Do(calcnewer.Mul(2.00))
	lastest := calcNewer.Do(calcnewer.Sqrt())

	fmt.Printf("Value: %v\n", lastest)

	myHouse := foptions.NewHouse(foptions.WithFireplace(), foptions.WithoutFireplace(), foptions.SetFloors(5))
	myHouse.Present()
}
