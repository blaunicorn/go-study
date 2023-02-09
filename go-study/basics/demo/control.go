package demo

import "fmt"

// 流程控制语句
func Control() {
	// 递增递减语句
	// ++
	// --  no front end
	d := 0
	d++
	fmt.Println(d)

	// 条件语句
	// go语言没有括号
	a := 2
	if a == 1 {
		fmt.Println(a)
	} else if a == 2 {
		fmt.Println(a)
	} else {
		fmt.Println(a)
	}

	// 选择语句
	a1 := 1
	switch a1 {
	case 0:
		fmt.Println("a1=0")
		fallthrough // 穿过，可以再走下一个
	case 1:
		fmt.Println("a1=1")
	default:
		fmt.Print(("都不是"))
	}

	// 循环语句
	count := 10
	//方式一
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
	// 方式二
	for {
		count++
		fmt.Println(count)
		if count > 20 {
			break
		}
	}

	// 跳转语句
	// break 与 continue配合标签可永不多层循环嵌套的跳出
	// goto 是调整执行位置
A:
	for i := 0; i < 10; i++ {
		if i == 1 {
			continue // 跳出本次循环
		}
		fmt.Println("我在A循环", i)
		if i > 3 {
			// 跳出A
			break A
			// 跳到B
			goto B
		}
	}
B:
	fmt.Println("我是B循环")
	fmt.Println("结束")

}
