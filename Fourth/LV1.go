package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		skillNames       string
		releaseSkillFunc func(string)
	)
	ReleaseSkill(skillNames, releaseSkillFunc) //使用ReleaseSkill函数
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	var num int
	DefineNames := make(map[int]string)
	DefineTemplate := make(map[int]string)

	DefineSkillNames(DefineNames, DefineTemplate) //如果要创建技能才调用函数
	func(string) {
		fmt.Println("输入你想要输出的技能对应的编号吧")
		fmt.Scan(&skillNames)
	}(skillNames)

	releaseSkillFunc(skillNames)
	num, _ = strconv.Atoi(skillNames)
	fmt.Println(DefineNames[num], DefineTemplate[num]) //输出的最终形式
}
func DefineSkillNames(DefineNames, DefineTemplate map[int]string) {
	var i int
	var tem, name, words string
	var jud bool
	words = "傻逼"

	for i = 1; ; i++ {
		fmt.Println("你是否想要创建一个技能？true or false")
		fmt.Scan(&jud)
		if jud == true {
			fmt.Println("自定义一个技能吧")             //创建技能
			fmt.Scan(&name)                     //输入你想要的技能名字
			index := strings.Index(name, words) //查看关键字是否出现在了技能名字的字符串中
			if index != -1 {
				fmt.Println("你的技能中包含敏感词汇") //如果出现了，则不添加进map
				break
			} else {
				DefineNames[i] = name //如果没有出现，则添加进map
				fmt.Println("自定义释放技能前输出的模板吧")
				fmt.Scan(&tem) //添加释放技能前的模板
				DefineTemplate[i] = tem
			}
		} else {
			break
		}
	}
}
