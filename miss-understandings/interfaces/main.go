package main

import "fmt"

type SalaryCalculator  interface {
  CalculateSalary() int
}

type Permanent struct {
    empId     int
    basicPay  int
    pf        int
}

type Contract struct {
    empId     int
    basicPay  int
}

type Freelancer struct {
    empId     int
    ratePerHour int
    totalHours int
}


// salary of permanent employee is the sum of basic pay and pf
func (r Permanent) CalculateSalary() int {
  total := r.basicPay + r.pf
  return total
}

// salary of contract employee is the basic pay alone
func (r Contract) CalculateSalary() int {
  return r.basicPay
}

// salary of freelancers
func (r Freelancer) CalculateSalary() int {
  return r.ratePerHour * r.totalHours
}


/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
  expense := 0
  for _, v := range s {
    expense += v.CalculateSalary()
  }
  fmt.Printf("Total expense per month $%d", expense)
}


func findType(i SalaryCalculator) {
    switch i.(type) {
    case Freelancer:
        fmt.Printf("I am a freelancer\n")
    case Permanent:
        fmt.Printf("I am a permanent employee\n")
    default:
        fmt.Printf("I am a contract employee and my id is %d\n", i.empId)
    }
}
func main() {
  pemp1 := Permanent{
        empId:    1,
        basicPay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicPay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicPay: 3000,
    }
    freelancer1 := Freelancer{
        empId:       4,
        ratePerHour: 70,
        totalHours:  120,
    }
    freelancer2 := Freelancer{
        empId:       5,
        ratePerHour: 100,
        totalHours:  100,
    }
    employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
    totalExpense(employees)
    findType(pemp1)
    findType(pemp2)
    findType(cemp1)
    findType(freelancer1)
    findType(freelancer2)
}
