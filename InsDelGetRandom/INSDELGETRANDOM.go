package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	randomMap  map[int]int
	randomList []int
}

func Constructor() RandomizedSet {

	return RandomizedSet{make(map[int]int), []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {

	if _, ok := this.randomMap[val]; ok {
		return false
	}
	this.randomList = append(this.randomList, val)
	this.randomMap[val] = len(this.randomList) - 1

	return true

}

func (this *RandomizedSet) Remove(val int) bool {

	if _, ok := this.randomMap[val]; !ok {
		return false
	}
	index := this.randomMap[val]
	last := len(this.randomList) - 1
	this.randomList[index] = this.randomList[last]
	this.randomMap[this.randomList[index]] = index
	this.randomList = this.randomList[0:last]
	delete(this.randomMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {

	random := rand.Intn(len(this.randomList))
	return this.randomList[random]
}

func main() {
	//val := 3
	obj := Constructor()
	fmt.Println("Object is", obj)
	param_1 := obj.Insert(1)
	fmt.Println("Object is", obj, " result is", param_1)
	param_1 = obj.Insert(2)
	fmt.Println("Object is", obj, " result is", param_1)
	param_1 = obj.Insert(3)
	fmt.Println("Object is", obj, " result is", param_1)
	param_1 = obj.Insert(2)
	fmt.Println("Object is", obj, " result is", param_1)

	fmt.Println("Time for random")
	num := obj.GetRandom()
	fmt.Println("Object is", obj, " result is", num)
	num = obj.GetRandom()
	fmt.Println("Object is", obj, " result is", num)
	num = obj.GetRandom()
	fmt.Println("Object is", obj, " result is", num)
	num = obj.GetRandom()
	fmt.Println("Object is", obj, " Random is", num)

	fmt.Println("removal\n\n\n")
	param_1 = obj.Remove(1)
	fmt.Println("Object is", obj, " result is", param_1)
	param_1 = obj.Remove(3)
	fmt.Println("Object is", obj, " result is", param_1)
	num = obj.GetRandom()
	fmt.Println("Object is", obj, " Random is", num)
	param_1 = obj.Remove(2)
	fmt.Println("Object is", obj, " result is", param_1)

	param_1 = obj.Remove(2)
	fmt.Println("Object is", obj, " result is", param_1)
	//  param_1 := obj.Insert(val);
	//  param_2 := obj.Remove(val);
	//  param_3 := obj.GetRandom();
}
