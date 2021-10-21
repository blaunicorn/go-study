package main

import "fmt"

/*
程序说明：
分50个金币
实现dispatchCoin函数
*/
var (
	coins        = 50
	users        = []string{"Mathew", "Sarah", "Augusetsu", "heidi", "elilimd", "Giana", "Adriano", "Aaron"}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("left corns:", left)
	for k, v := range distribution {
		fmt.Printf("%s:%d\n", k, v)
	}
}

func dispatchCoin() (left int) {
	// 1、依次拿到每个人的名字
	// 2、拿到一个人的 名字根据分金币的规则分金币；
	// 2.1 每个人分的金币数应该保存在distribution中；const
	// 2.2 记录下剩余的金币数
	// 3. 整个第二步执行完就能得到最终每个人分的金币数和剩余金币数量

	for _, name := range users {
		for _, c := range name {
			// ''字符 ""字符串
			switch c {
			case 'e', 'E':
				distribution[name]++

			case 'i', 'I':
				distribution[name] += 2
			case 'o', 'O':
				distribution[name] += 3
			case 'u', 'U':
				distribution[name] += 4
			}
		}
	}
	left = 0
	for _, value := range distribution {
		left += value
	}
	left = coins - left

	return
}
