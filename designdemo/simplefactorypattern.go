package main

import "fmt"

// 简单工厂模式

type User interface {
	GetRole() string
}

type Member struct {}

func (this *Member) GetRole() string  {
	return "会员用户"
}

type Admin struct {}

func (this *Admin) GetRole() string {
	return "后台管理用户"
}

const (
	Mem = iota
	Adm
)

func CreateUser(t int) User {
	switch t {
	case Mem:
		return new(Member)
	case Adm:
		return new(Admin)
	default:
		return new(Member)
	}
}

func main() {
	fmt.Println(CreateUser(1).GetRole())
}
