// 首页菜单
package models

type Menu struct {
	Id   int
	Name string
	Url  string
}

// MenuList 菜单列表
func MenuList() map[int]Menu {
	return map[int]Menu{
		1: {Id: 1, Name: "编程", Url: "/1"},
		2: {Id: 2, Name: "读书", Url: "/2"},
		3: {Id: 3, Name: "旅行", Url: "/3"},
	}
}

// Menus 菜单字典
func Menus() map[int]string {
	menu := make(map[int]string)
	for _, m := range MenuList() {
		menu[m.Id] = m.Name
	}
	return menu
}

// MenuDesc 菜单描述
func MenuDesc(id int) Menu {
	if desc, ok := MenuList()[id]; ok {
		return desc
	}
	return Menu{}
}
