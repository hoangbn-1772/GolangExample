package main

import "fmt"

func main() {
	st := Student{
		student_name:    "Kutsuzawa",
		student_age:     27,
		student_address: "Japanese",
	}
	st.display()

	phone_number := PhoneNumber("abc-12345")
	phone_number.displayPhoneNumber()

	st2 := Student{
		student_name:    "Hoang",
		student_age:     23,
		student_address: "VietNamese",
	}
	st2.display2()
	st2.changeAge(2)
	st2.display2()
}

type Student struct {
	student_name    string
	student_age     int
	student_address string
}

/**
* * Method with structure and receiver is value
 */
func (st Student) display() {
	fmt.Println()
	fmt.Printf("Name: %s - Age: %d - Address: %s", st.student_name, st.student_age, st.student_address)
}

/**
* * Method with non-structure
 */
type PhoneNumber string

func (p PhoneNumber) displayPhoneNumber() {
	fmt.Println()
	fmt.Printf("Phone Number: %s", p)
}

/**
* *Method with receiver is pointer
 */
func (st *Student) display2() {
	fmt.Println()
	fmt.Printf("Name: %s - Age: %d - Address: %s", st.student_name, st.student_age, st.student_address)
}

//! Can't change age
func (st Student) changeAge(newAge int) {
	st.student_age = newAge
}

//! Can change age
func (st *Student) changeAge(newAge int) {
	st.student_age = newAge
}
