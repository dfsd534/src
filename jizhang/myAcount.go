package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true

	//定义余额
	balance := 10000.0
	//收支金额
	money := 0.0
	//收支说明
	note := 0.0
	//收支详情
	details := "收支\t账号金额\t收支金额\t说  明"

	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")

		fmt.Println("请选择...<1-4>")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("----------当前收支明细记录----------")
			fmt.Println(details)
		case "2":
			fmt.Println("收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("收入说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("支出金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("支出说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t-%v\t%v", balance, money, note)
		case "4":
			fmt.Println("你确定要退出么？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入y/n")
			}

			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选择..")
		}

		if !loop {
			break
		}
	}

	fmt.Println("您退出家庭记账软件使用...")
}
