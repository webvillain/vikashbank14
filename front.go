package touchup

import "fmt"

func WelcomeIntro() {
	fmt.Println("                                                                   ")
	fmt.Println("\t\t\t\t\t\t\t|-------                                                 -------|")
	fmt.Println("\t\t\t\t\t\t\t                Welcome To V-Banking Application               ")
	fmt.Println("\t\t\t\t\t\t\t  Application Will Start In Your Local-Machine On Port : 8000  ")
	fmt.Println("\t\t\t\t\t\t\t|_____________________________________________________________|")
	fmt.Println("")
	fmt.Println("\t\t\t\t\t _______________________________________________________________________________________________________________________________________________")
	fmt.Println("\t\t\t\t\t  You Can Perform Below Listed 5 Operations With Our Application Database By Making An Http Requests.")
	fmt.Println("\t\t\t\t\t======================================================================================================")
	fmt.Println("\t\t\t\t\t\t\t  1. Get List Of All User Existing In Our Banking Application Database.")
	fmt.Println("\t\t\t\t\t\t\t             2. Get A Single User Information By Its Id.")
	fmt.Println("\t\t\t\t\t\t\t            3. Create A New User In Application Database.")
	fmt.Println("\t\t\t\t\t\t\t         4. Delete A User From Application Database By Its Id.")
	fmt.Println("\t\t\t\t\t\t\t  5. Update A User Information Like Name , Email , UserName By Its Id.")
	fmt.Println("\t\t\t\t\t ______________________________________________________________________________________________________")
	fmt.Println("")
}

func IntroAfterDatabaseConnection() {
	fmt.Println("\t\t\t\t\t__________________________You Are Now Connected To V-Banking Application Database ______________________")
	fmt.Println("\t\t\t\t\t________________________________________________________________________________________________________")
}
