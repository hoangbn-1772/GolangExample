package main

import "fmt"

func main() {
	employees := []Employee{
		Employee{empId: 1, basicPay: 5000, reward: 2000},
		Employee{empId: 2, basicPay: 4000, reward: 7000},
		Employee{empId: 3, basicPay: 5000, reward: 1000},
	}
	emp := Employees{emps: employees}
	emp.display()

	var emp2 *Employee
	var cal Calculator
	cal = emp2
	fmt.Printf("%d", emp2.Calculate())

	// emp := &Employee{empId: 1, basicPay: 2, reward: 4}
	// a := &SalaryCalculator{salary: emp} //! Bacause Employee implement Calculator
	// fmt.Printf("Salary: %d", a.salary.Calculate())
}

/**
* * Define interface
 */
type Calculator interface {
	Calculate() int
}

type SalaryCalculator struct {
	salary Calculator
}

type Employee struct {
	empId    int
	basicPay int
	reward   int
}

/**
* * Total salary (Implement Interface)
 */
func (emp *Employee) Calculate() int {
	if emp == nil {
		return -1
	}

	return (emp.basicPay + emp.reward)
}

type Employees struct {
	emps []Employee
}

func (employees *Employees) display() {
	for _, v := range employees.emps {
		fmt.Printf("Id: %d - Salary: %d\n", v.empId, v.Calculate())
	}
}
