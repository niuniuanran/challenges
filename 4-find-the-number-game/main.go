package main

import ("fmt" )

func main() {
	secret:= 401
	var guess int
	for true{
	fmt.Print("Enter a number: ")
	fmt.Scan(&guess)
	if secret < guess {
		fmt.Println("It's less!")
	} else if secret > guess {
		fmt.Println("It's more!")
	} else {
		fmt.Println("✨ You win! ✨")
		break
	}

	}





}