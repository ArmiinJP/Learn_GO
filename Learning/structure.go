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

func StructLearning(){
	type User struct{
		ID uint
		Name string
		Email string
	}

	//var declaration from struct
	{
		fmt.Println("-------------------var declareation")
		// define var from struct with zero value
		//----------------
			var u1 User
		//----------------
		
		
		// define var from struct with zero value
		//----------------
			var u2 = User{}
		//----------------

		
		// define var from struct with initilize value
		// when field name specify, you can use both of initlize and zero 
		// see u5
		//----------------
			var u3 = User{
				ID: 1,
				Name: "hasan",
				Email: "h@h",
			}
		//----------------


		// define var from struct with zero and initialize value
		// when field name specify, you can use both of initlize and zero 
		// ID with zero value and Name & Email Initialize value
		//----------------
			u4 := User{
				Name: "hasan",
				Email: "h@h",
			}
		//----------------
		
		
		// define var from struct with initialize value
		// when field name doesn't specify, you should initialize all field 
		//----------------
			var u5 = User{
				1,
				"hasan",
				"h@h",
			}
		//----------------

		fmt.Println("u1 is:", u1.ID, u1.Name, u1.Email)
		fmt.Println("u2 is:", u2.ID, u2.Name, u2.Email)
		fmt.Println("u3 is:", u3.ID, u3.Name, u3.Email)
		fmt.Println("u4 is:", u4.ID, u4.Name, u4.Email)
		fmt.Println("u5 is:", u5.ID, u5.Name, u5.Email)
	}

	//Pointer declaration from struct
	{
		fmt.Println("-------------------pointer declareation")
		
		//sample struct
		var u1 = User{
			ID: 3,
			Name: "hsaan",
			Email: "h@h",
		}


		// pointer to struct with nil value
		//---------------------------------------
			var u2 *User
		//---------------------------------------
		if u2 == nil{
			fmt.Println("u2 is nil")
			u2 = &u1
			fmt.Println("after assignment u2 is:", u2.ID, u2.Name, u2.Email)
		}
		//---------------------------------------


		// pointer to struct with zero value
		//---------------------------------------
			var u3 = &User{}
		//---------------------------------------
		if u3 == nil{
			fmt.Println("u3 is nil")
		} else {
			fmt.Println("u3 is:", u3.ID, u3.Name, u3.Email)
		}


		// pointer to struct with initialize value using assignment
		//---------------------------------------
			var u4 = &u1
		//---------------------------------------
		if u4 == nil{
			fmt.Println("u4 is nil")
		} else {
			fmt.Println("u4 is:", u4.ID, u4.Name, u4.Email)
		}


		// pointer to struct with initizalize value
		//---------------------------------------		
			var u5 = &User{
				ID: 2,
				Name: "hossain",
				Email: "h@h",
			}
		//---------------------------------------
		if u5 == nil{
			fmt.Println("u5 is nil")
		} else {
			fmt.Println("u5 is:", u5.ID, u5.Name, u5.Email)
		}


		//pointer to struct with zero value
		//---------------------------------------
			var u6 = new(User)
		//---------------------------------------
		if u6 == nil{
			fmt.Println("u6 is nil")
		} else {
			fmt.Println("u6 is:", u6.ID, u6.Name, u6.Email)
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

	fmt.Println("\n\n============== struct Learning ==============")
	StructLearning()
	
	fmt.Println("\n\n============== map Learning ==============")
	MapLearning()
}
