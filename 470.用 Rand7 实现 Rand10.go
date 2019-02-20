package code

import (
	"math/rand"
	"time"
)

// Rand10 : 生成 1 到 10 范围内的均匀随机整数
func Rand10() int {
	num := 100
	for num > 10 {
		num = (Rand7()*7+Rand7())/4 - 1
	}
	return num
}

// Rand7 : 生成 1 到 7 范围内的均匀随机整数
func Rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}
