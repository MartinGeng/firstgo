package main

import (
	"fmt"
)
/*
	map 是一种特殊的数据结构：一种元素对（pair）的无序集合，pair 的一个元素是 key，
	对应的另一个元素是 value，所以这个结构也称为关联数组或字典。
	这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。

	map 是引用类型，可以使用如下声明：
	var map1 map[keytype]valuetype
	var map1 map[string]int
	key 可以是任意可以用 == 或者 != 操作符比较的类型，
	比如 string、int、float。所以数组、切片和结构体不能作为 key 但是指针和接口类型可以

	map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，
	无论实际上存储了多少数据。通过 key 在 map 中寻找值是很快的，比线性查找快得多，
	但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果你很在乎性能的话还是建议用切片来解决问题。

	不要使用 new，永远用 make 来构造 map.why?
*/
func main() {
	var mapList map[string]int
	var mapAssigened map[string]int

	mapList = map[string]int{"one":1, "two":2, "three":3}
	mapCreated := make(map[string]float32)
	mapAssigened = mapList

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigened["two"] = 3

	fmt.Println("map literal at maplist is ", mapList["one"])
	fmt.Println("map literal at maplist is ", mapCreated["key2"])
	fmt.Println("map literal at maplist is ", mapList["two"])
	fmt.Println("map literal at maplist is ", mapList["ten"])
}