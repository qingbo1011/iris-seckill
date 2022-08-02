package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var sum int64 = 0

var productNum int64 = 1000000 // 预存商品数量

var mutex sync.Mutex // 互斥锁

var count int64 = 0 // 计数

func main() {
	http.HandleFunc("/getOne", GetProduct)
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		log.Fatal("Err:", err)
	}
}

// GetProduct 获取秒杀商品
func GetProduct(w http.ResponseWriter, req *http.Request) {
	if GetOneProduct() {
		w.Write([]byte("true"))
		return
	}
	w.Write([]byte("false"))
	return
}

// GetOneProduct 获取秒杀商品
func GetOneProduct() bool {
	mutex.Lock()         // 加锁
	defer mutex.Unlock() // 释放锁
	count += 1
	if count%100 == 0 { // 判断数据是否超限
		if sum < productNum {
			sum += 1
			fmt.Println(sum)
			return true
		}
	}
	return false
}
