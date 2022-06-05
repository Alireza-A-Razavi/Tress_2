package main

import "fmt"

type SalaryCalculator interface {
  DisplaySalary()
}

type LeavesCalculator interface {
  CalculateLeavesLeft() int
}

// embedding interfaces
type EmployeeOperations interface {
  SalaryCalculator
  LeaveCalculator
}

type Employee struct {
  firstName string
  lastName string
  basicPay int
  pf int
  totalLeaves int
  leavesTaken int
}

func (e Employee) DisplaySalary() {
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
    return e.totalLeaves - e.leavesTaken
}

func main() {
  e := Employee {
    firstName: "Naveen",
    lastName: "Ramanathan",
    basicPay: 5000,
    pf: 200,
    totalLeaves: 30,
    leavesTaken: 5,
  }
  var s SalaryCalculator = e
  s.DisplaySalary()
  var l LeaveCalculator = e
  fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}
