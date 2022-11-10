package main

import (
	"fmt"
	"time"
)

var (
	De1 bool
	De2 bool
)

func main() {

	var num int
	go Timer()  //设置本来就有的两个定时器
	go Remind() //设置在一进入程序就循环输出的“芜湖，起飞！”
	for {
		fmt.Println("你想要实现什么功能？\n1.增加定时器\n2.删除原有定时器\n3.取消定时器的下一次提醒")
		fmt.Println("输入其他数字视为取消设定")
		fmt.Scan(&num)
		switch num {
		case 1:
			AddTimer()
			break
		case 2:
			Delete()
			break
		case 3:
			Delay()
		default:
			return
		}
	}
}
func Delay() {
	var num int
	fmt.Println("你想要取消哪一个呢？\n1.2点的\n2.6点的")
	fmt.Scan(&num)
	switch num {
	case 1:
		De1 = true
		fmt.Println("成功取消下一次提醒")
		break
	case 2:
		De2 = true
		fmt.Println("成功取消下一次提醒")
		break
	default:
		fmt.Println("格式错误")
	}
}
func Remind() {
	ticker := time.NewTimer(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("芜湖！起飞！")
		}
	}

}

func Delete() {
	ticker := time.NewTimer(30 * time.Second)
	timer1 := time.NewTimer(time.Now().Sub(time.Now()))
	timer2 := time.NewTimer(time.Now().Sub(time.Now()))
	var (
		num int
	)
	fmt.Println("想要删除哪个计时器？\n1.每天早上2点的\n2.每天早上6点的\n3.循环的那个")
	fmt.Scan(&num)
	switch num {
	case 1:
		timer2.Stop()
		fmt.Println("删除成功")
		break
	case 2:
		timer1.Stop()
		fmt.Println("删除成功")
		break
	case 3:
		fmt.Println("删除成功")
		ticker.Stop()
	default:
		fmt.Println("你输出的格式不正确哦")
		return
	}
}
func AddTimer() {
	var num int
	fmt.Println("添加一个定时器吧：\n1.一次性\n2.重复性")
	fmt.Scan(&num)
	switch num {
	case 1:
		Onetime()
		break
	case 2:
		Repeat()
		break
	default:
		fmt.Println("输入格式错误！")
		return
	}
}
func Onetime() {
	var (
		things string
		t1     int
	)
	fmt.Println("输入你想过多长时间提醒你干什么事情？格式：时间(秒) 事件")
	fmt.Scan(&t1, &things)
	ticker2 := time.NewTicker(time.Second)
	select {
	case <-ticker2.C:
		fmt.Println(things)
	}
}
func Repeat() {
	var (
		things string
		t2     int
	)
	fmt.Println("输入你想每过多长时间提醒你干什么事情？格式：时间(秒) 事件")
	fmt.Scan(&t2, &things)
	ticker3 := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker3.C:
			fmt.Println(things)
		}
	}
}
func Timer() {

	//设置本来就有的两个定时器
	m := map[int]string{
		1: "我还能再战4小时!",
		2: "我要去图书馆开卷！",
	}
	for {
		now := time.Now()
		var next1, next2 time.Time
		hour, minute, second := now.Clock()
		judge1 := hour > 2 && minute > 0 && second > 0
		judge2 := hour > 6 && minute > 0 && second > 0
		//先判断现在的时间在2点或者6点之前、之后还是之间
		//如果在之前，就直接设定定时器；如果在之后，就从明天的时间来设定
		if judge1 == true {
			if judge2 == true {
				next1 = now.Add(time.Hour * 24)
				next2 = now.Add(time.Hour * 24)
			} else {
				next1 = now.Add(time.Hour * 24)
				next2 = now
			}
		} else {
			next1 = now
			next2 = now
		}
		next1 = time.Date(next1.Year(), next1.Month(), next1.Day(), 2, 0, 0, 0, next1.Location())
		next2 = time.Date(next1.Year(), next1.Month(), next1.Day(), 6, 0, 0, 0, next1.Location())
		timer1 := time.NewTimer(next1.Sub(now))
		timer2 := time.NewTimer(next2.Sub(now))
		select {
		case <-timer1.C:
			if De1 == true {
				time.Sleep(24 * time.Hour)//我希望取消下一次提醒可以采用睡24小时的办法
			}
			fmt.Println(m[1])
		case <-timer2.C:
			if De2 == true {
				time.Sleep(24 * time.Hour)
			}
			fmt.Println(m[2])
		}
	}
}
