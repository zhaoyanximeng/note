package main

import "fmt"

// 抽象工厂模式，create也抽象出来

//type User interface {
//	GetRole() string
//}
//
//type Member struct {}
//
//func (this *Member) GetRole() string  {
//	return "会员用户"
//}
//
//type Admin struct {}
//
//func (this *Admin) GetRole() string {
//	return "后台管理用户"
//}

type AbstractFactory interface {
	CreateUser() User
}

type MemberFactory struct {}
func (this *MemberFactory) CreateUser() User {
	return &Member{}
}

type AdminFactory struct {}
func (this *AdminFactory) CreateUser() User{
	return &Admin{}
}

func main() {
	var fact AbstractFactory = new(AdminFactory)
	fmt.Println(fact.CreateUser().GetRole())
}
