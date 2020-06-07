package utils

import (
	"fmt"
)

type familyAccount struct {
	key string
	loop bool
	//定义余额
	balance float64
	//收支金额
	money float64
	//收支说明
	note string
	//收支详情
	details string
}

func NewObject() *familyAccount{
	return &familyAccount{
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		details : "收支\t账号金额\t收支金额\t说  明",
	}
}

func (f *familyAccount) showDetail(){
	fmt.Println("----------当前收支明细记录----------")
	fmt.Println(f.details)
}

func (f *familyAccount) income(){
	fmt.Println("收入金额:")
	fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Println("收入说明:")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", f.balance, f.money, f.note)
}

func (f *familyAccount) pay(){
	fmt.Println("支出金额:")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("余额不足")
		return 
	}
	f.balance -= f.money
	fmt.Println("支出说明:")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n支出\t%v\t-%v\t%v", f.balance, f.money, f.note)
}

func (f *familyAccount) exits(){
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
		f.loop = false
	}
}

func (f *familyAccount) MainMenu(){
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")

		fmt.Println("请选择...<1-4>")
		fmt.Scanln(&f.key)

		switch f.key {
		case "1":
			f.showDetail()
		case "2":
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exits()
		default:
			fmt.Println("请输入正确的选择..")
		}

		if !f.loop {
			break
		}
	}

	fmt.Println("您退出家庭记账软件使用...")
}