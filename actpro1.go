package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type user struct {
	name    string
	phoneno uint64
	address string
}

type users struct {
	us []user
}

type product struct {
	protype string
}

type order struct {
	orderid     string
	productname string
	quantity    int
}

type orders struct {
	od []order
}

/*type userver interface {
	adduser(u []user, phoneno uint64) bool
}*/

func (u *users) adduser(pno uint64) bool {
	var flag bool = false
	for _, elem := range u.us {
		fmt.Printf("%v\n", elem)
		if elem.phoneno == pno {
			flag = true
			break
		}
	}

	return flag
}

func getproducts(pro []product) {
	fmt.Println("list of products : ")
	for _, ele := range pro {
		fmt.Printf("%s\n", ele.protype)
	}
	fmt.Printf("\n\n")
}

func (o *orders) palceorder(item string, quan int) {
	var ordid string = "oid" + strconv.Itoa(rand.Intn(1000))
	var ornew order = order{orderid: ordid, productname: item, quantity: quan}
	o.od = append(o.od, ornew)
}

func (o *orders) changeorder(odid string, quan int) {

	for i, elem := range o.od {
		if elem.orderid == odid {
			o.od[i].quantity = quan
			break
		}
	}

}

func (o *orders) getorders() {
	for i := 0; i < len(o.od); i++ {
		fmt.Printf("orders :%v\n", o.od[i])
	}

	fmt.Printf("\n\n")
}

func (o *orders) deleteorder(doid string) {
	for i := 0; i < len(o.od); i++ {
		if o.od[i].orderid == doid {
			o.od = append(o.od[:i], o.od[i+1:]...)
			break
		}
	}
}

func main() {
	var u users
	var ods orders
	var first user = user{name: "murali", phoneno: 9856325675, address: "pune"}
	u.us = append(u.us, first)
	//fmt.Printf("%v\n", u[0])
	var uname string
	var phno uint64
	var addr string
	//var uv userver
	var flag bool

	var choice int

	var prot []product = []product{
		{protype: "tv"},
		{protype: "mobile"},
		{protype: "washing machine"},
		{protype: "pan"}}

	for {
		fmt.Printf("enter :\n1-add user\n2-To get products list\n3-place order\n4-changing order\n5-delete order \n6-get orders\n")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			{
				fmt.Println("enter user name, phoneno, place")
				fmt.Scanln(&uname)
				fmt.Scanln(&phno)
				fmt.Scanln(&addr)

				first = user{name: uname, phoneno: phno, address: addr}

				flag = u.adduser(phno)

				if !flag {
					u.us = append(u.us, first)
					fmt.Printf("user registered\n")
				} else {
					fmt.Printf("phone no already registered\n")
				}
				fmt.Printf("\n\n users present\n\n")
				for _, elem := range u.us {
					fmt.Printf("%v\n", elem)
				}
			}
		case 2:
			{
				getproducts(prot)
			}

		case 3:
			{

				fmt.Println("place order")
				var proname string
				var quan int
				fmt.Println("enter product name and quantity for order")
				fmt.Scanln(&proname)
				fmt.Scanln(&quan)

				ods.palceorder(proname, quan)

				fmt.Println("orders paced are:")
				ods.getorders()

			}

		case 4:
			{
				var ordid string
				var quan int

				fmt.Println("enter  orderid for changing the order")
				fmt.Scanln(&ordid)

				fmt.Println("enter  quantity to change for order")
				fmt.Scanln(&quan)

				ods.changeorder(ordid, quan)

				fmt.Println("orders after changing are:")
				ods.getorders()

			}
		case 5:
			{
				var delid string
				fmt.Printf("enter the orderid to delete order:\n")
				fmt.Scanln(&delid)

				ods.deleteorder(delid)

				fmt.Println("orders after deleteing are:")
				ods.getorders()

			}
		case 6:
			{
				ods.getorders()
			}

		default:
			os.Exit(1)

		}
	}

}
