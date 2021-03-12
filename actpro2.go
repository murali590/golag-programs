package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Representative struct {
	Name   string
	Mobile string
	Place  string
}

type deliveryorg struct {
	reps []Representative
}

type order struct {
	orderid  string
	devplace string
}

type assignedrep struct {
	repassigned string
	ordid       string
	delplace    string
}

type assigns struct {
	ass []assignedrep
}

func (d *deliveryorg) AddDeliveryrep(rep Representative) bool {
	if d.reps == nil {
		d.reps = make([]Representative, 0)
	}

	for i, _ := range d.reps {
		if d.reps[i].Mobile == rep.Mobile {
			return false
		}
	}

	return true
}

func (as *assigns) assigndelv(d deliveryorg, od order) bool {
	for i, _ := range d.reps {
		if d.reps[i].Place == od.devplace {
			arep := assignedrep{repassigned: d.reps[i].Name, ordid: od.orderid, delplace: od.devplace}
			as.ass = append(as.ass, arep)
			return true
		}
	}

	return false

}

func (as *assigns) viewassg() {
	fmt.Println("assignments list:\n")
	for i, _ := range as.ass {
		fmt.Printf("%v\n", as.ass[i])
	}
}

func main() {
	var dev deliveryorg
	dev.reps = append(dev.reps, Representative{Name: "ghj", Mobile: "987456321", Place: "bang"})
	fmt.Printf("%v\n", dev.reps[0])
	//dev.AddDeliveryrep(Representative{Name: "ghj", Mobile: "987456321", Place: "bang"})

	var flag bool
	var rname string
	var phno string
	var addr string

	var assg assigns
	var odplace string
	ord := make([]order, 0)
	//var uv userver
	for {
		var userchoice int
		fmt.Printf("enter\n 1.add representative \n 2.Assign order to representative \n 3.display orders and its representative\n")
		fmt.Scanln(&userchoice)

		switch userchoice {
		case 1:
			{
				fmt.Println("enter representative name, phoneno, place")
				fmt.Scanln(&rname)
				fmt.Scanln(&phno)
				fmt.Scanln(&addr)

				addrep := Representative{Name: rname, Mobile: phno, Place: addr}

				flag = dev.AddDeliveryrep(addrep)

				if flag {
					dev.reps = append(dev.reps, addrep)
					for i := 0; i < len(dev.reps); i++ {
						fmt.Printf("%v\n", dev.reps[i])
					}
				} else {
					fmt.Println("mobile no already regristered")
				}
			}
		case 2:
			{
				fmt.Println("enter order place")
				fmt.Scanln(&odplace)

				var ordid string = "oid" + strconv.Itoa(rand.Intn(1000))
				var ornew order = order{orderid: ordid, devplace: odplace}

				flag = assg.assigndelv(dev, ornew)

				if flag {
					fmt.Println("order assigned to representative\n")
					ord = append(ord, ornew)
				} else {
					fmt.Println("place of order do not have any representative\n")
				}
			}

		case 3:
			{
				assg.viewassg()
			}

		default:
			{
				os.Exit(1)
			}

		}

	}
}
