package main

import "fmt"

type person struct {
	fname       string
	lname       string
	iceflavours []string
}

func main() {
	//firstsecond()
	thirdfourth()
}

func firstsecond() {

	persArr := []person{}

	pers1 := person{
		fname:       "John",
		lname:       "Lennon",
		iceflavours: []string{"vanilla", "chocolate"},
	}

	persArr = append(persArr, pers1)

	pers2 := person{
		fname:       "Bill",
		lname:       "Gates",
		iceflavours: []string{"banana", "strawberry"},
	}

	persArr = append(persArr, pers2)

	for _, item := range persArr {
		fmt.Printf("\n%v %v \nlikes flavours:\n", item.fname, item.lname)
		for _, flavour := range item.iceflavours {
			fmt.Printf("\t%v\n", flavour)
		}
	}

	mappers := map[string]person{pers1.lname: pers1, pers2.lname: pers2}

	for key, value := range mappers {
		fmt.Println("Item", key, "Value", value)

	}

}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func thirdfourth() {
	veh1 := sedan{vehicle: vehicle{doors: 4, color: "black"}, luxury: true}
	fmt.Println(veh1)
	veh2 := truck{vehicle: vehicle{doors: 4, color: "green"}, fourWheel: true}
	fmt.Println(veh2)

	veh3 := struct {
		vehicle      vehicle
		hasExcavator bool
	}{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		hasExcavator: true,
	}

	fmt.Println(veh3)
}
