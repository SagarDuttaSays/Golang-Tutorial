package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")
	std1 := student{"Sagar Dutta", "BTech - 3rd year", 10037, 22, true}
	fmt.Println("Attendance status:", std1.GetAttendance())
	fmt.Println(std1.UpdateStandard("BTech - 4th year"))
	fmt.Println("Standard:", std1.Standard)
}

type student struct {
	Name       string
	Standard   string
	ID         int
	Age        int
	Attendance bool
}

func (s student) GetAttendance() bool {
	return s.Attendance
}

func (s student) UpdateStandard(c string) string {
	s.Standard = c
	fmt.Println("Standard:", s.Standard)
	return "Updated successfully"
}
