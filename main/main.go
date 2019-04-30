package main

import "fmt"

func main()  {
	cache := CreateNewCache(5)
	cache.SaveValue("hehehe", "呵呵呵")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.SaveValue("hahaha", "哈哈哈")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.SaveValue("heihei", "嘿嘿")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.SaveValue("hahaha", "哼哼哼")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.SaveValue("hehehe", "哟哟哟")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.SaveValue("haihaihai", "嗨嗨嗨")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.SaveValue("huahuahua", "哗哗哗")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.SaveValue("xixixi", "嘻嘻嘻")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.GetValue("huahuahua")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.GetValue("haihaihai")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	iterator := cache.GetIterator()
	for iterator.HasNext() {
		key,value := iterator.Next()
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
	fmt.Println("===================分割线===================")

	cache.SaveValue("lalala", "啦啦啦")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	cache.GetValue("huahuahua")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	iterator2 := cache.GetIterator()
	for iterator2.HasNext() {
		key,value := iterator2.Next()
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
	fmt.Println("===================分割线===================")

	cache.Remove("xixixi")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	iterator3 := cache.GetIterator()
	for iterator3.HasNext() {
		key,value := iterator3.Next()
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
	fmt.Println("===================分割线===================")

	cache.Remove("heihei")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	iterator4 := cache.GetIterator()
	for iterator4.HasNext() {
		key,value := iterator4.Next()
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
	fmt.Println("===================分割线===================")

	cache.Remove("huahuahua")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	iterator5 := cache.GetIterator()
	for iterator5.HasNext() {
		key,value := iterator5.Next()
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
	fmt.Println("===================分割线===================")

	cache.SaveValue("lelele", "乐乐乐")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	iterator6 := cache.GetIterator()
	for iterator6.HasNext() {
		key,value := iterator6.Next()
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
	fmt.Println("===================分割线===================")

	cache.SaveValue("hehehe", "呵呵呵")
	fmt.Println("head:", cache.head.value)
	fmt.Println("tail:", cache.tail.value)
	fmt.Println("size:", cache.size)
	fmt.Println("===================分割线===================")

	iterator7 := cache.GetIterator()
	for iterator7.HasNext() {
		key,value := iterator7.Next()
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
	fmt.Println("===================分割线===================")
}
