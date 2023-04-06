package main

import "fmt"

func MapLearning(){

// map "ok paradiam"
	{
		fmt.Println("------------------map \"ok paradiam\"")
		
		var map1 = map[string]int{
			"hossain": 19,
		}

		if value, ok := map1["hossain"]; !ok{
			fmt.Println("key is not exist")
		} else {
			fmt.Println("key exist and value is:", value)
		}

		if value, ok := map1["alireza"]; !ok{
			fmt.Println("key is not exist")
		} else {
			fmt.Println("key exist and value is:", value)
		}
	}

// loop in map
	{
		fmt.Println("\n\n------------------loop in map ")

		var map1 = map[string]int{
			"hossain": 19,
			"alireza": 20,
			"mohseni": 18,
		}

		for key, value := range map1{
			fmt.Println("key is:", key, ",value is:", value)
		}
	}

// map and struct (struct as value  &&  struct as key)
	{
		fmt.Println("\n\n------------------map and struct ")

		fmt.Println("struct as value:::")
		type student struct{
			ID uint
			Name string
			avg int
		}

		var sampleStudent = student{
			3,
			"mohseni",
			18,
		}

		var map1 = map[string]student{
			"alireza": {1, "alireza", 19,},
			"hossain": {2, "hossain", 20,},
			"mohseni": sampleStudent,
		}
		
		fmt.Printf("%+v\n", map1)


		fmt.Println("\nstruct as key:::")
		var map2 = map[student]string{
			sampleStudent: "mohseni",
			{2, "hossain", 20}: "hossain",
		}
		
		if _, ok := map2[student{ID: 2,}]; !ok{ 
			fmt.Println("Key is not Exist")}

		if _, ok := map2[student{ID: 2, Name: "hossain"}]; !ok{ 
			fmt.Println("Key is not Exist")}
		
		if value, ok := map2[student{ID: 2, Name: "hossain", avg: 20}]; !ok{
			fmt.Println("Key is not Exist")} else {
				fmt.Println("key exist and value is:", value)
			}				
	}

// declaraition map
	{
		fmt.Println("\n\n------------------map declaration")

	// declarition map with zero value (with make) **
	// --------------------------------
		map1 := make(map[string]int)
	// --------------------------------
		{
			fmt.Println("map1:::::")
			if map1 == nil {fmt.Println("map1 is nil")}
			fmt.Println(map1, len(map1))
			map1["hasan"] = 2
			fmt.Println(map1, len(map1))
		}

	// declarition map with zero value (shortest declare)
	// --------------------------------
		map2_0 := map[string]int{}
	// --------------------------------
		{
			fmt.Println("\nmap2_0:::::")
			if map2_0 == nil {fmt.Println("map2_0 is nil")}
			fmt.Println(map2_0, len(map2_0))
			map2_0["hasan"] = 2
			fmt.Println(map2_0, len(map2_0))
		}

	// declarition map with zero value (longest declare) **
	//--------------------------------
		var map2_1 = map[string]int{}
	// --------------------------------
		{
			fmt.Println("\nmap2_1:::::")
			if map2_1 == nil {fmt.Println("map2_1 is nil")}
			fmt.Println(map2_1, len(map2_1))
			map2_1["hasan"] = 2
			fmt.Println(map2_1, len(map2_1))
		}

	// declartion map and initilize value 
	//--------------------------------
		var map2_2 = map[string]int{
			"hossain": 19,
			"alireza": 20,
		}
	//--------------------------------
		{
			fmt.Println("\nmap2_2:::::")
			fmt.Println(map2_2, len(map2_2))
			map2_2["hasan"] = 2
			fmt.Println(map2_2, len(map2_2))
		}

	// declration nil map and initilize with make
	//--------------------------------
		var map3 map[string]int
	//--------------------------------
		{
			fmt.Println("\nmap3:::::")
			if map3 == nil {fmt.Println("map3 is nil")}
			map3 = make(map[string]int) // dar asl taze tarif shode
			fmt.Println(map3, len(map3))
			map3["hasan"] = 2
			fmt.Println(map3, len(map3))
		}

	// just declration nil map
	//--------------------------------
		var map4 map[string]int
	//--------------------------------
		{
			fmt.Println("\nmap4:::::")
			if map4 == nil {fmt.Println("map4 is nil")}
			fmt.Println(map4, len(map4))
			map4["hasan"] = 2
		}
	}
}

func main() {

	//panic handling
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	fmt.Println("\n\n============== map Learning ==============")
	MapLearning()
}
