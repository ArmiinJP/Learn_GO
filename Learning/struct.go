package main

import "fmt"

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

}
