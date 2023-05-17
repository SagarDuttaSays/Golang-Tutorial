package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in Golang")

	// var studentsList = []string{}
	// studentsList = append(studentsList, "Arpit", "Ashish", "Arjun", "Balaji")

	// fmt.Println("Student List:", studentsList)

	// fmt.Println("Enter Teachers' name")
	// reader := bufio.NewReader(os.Stdin)

	// var teachersList = make([]string, 3)
	// for i := 1; i < 5; i++ {
	// 	name, _ := reader.ReadString('\n')
	// 	teachersList = append(teachersList, name)
	// }

	// fmt.Println("Teacher List:", teachersList)

	// //sort package

	// fmt.Println(sort.SearchStrings(studentsList, "Arpit"))
	// sort.Strings(teachersList)
	// fmt.Println("After sorting:", teachersList)

	// var num = []int{}
	// num = append(num, 234, 456, 123, 345)
	// fmt.Println("numerical Slice is sorted:", sort.IntsAreSorted(num))
	// sort.Ints(num)
	// fmt.Println("After sorting")
	// fmt.Println("numerical Slice is sorted:", sort.IntsAreSorted(num))

	//removing values from a slice
	var courses = []string{"React js", "Node js", "MongoDB", "Express js", "MySQL", "Swift", "Java"}
	sort.Strings(courses)
	fmt.Println("courses:", courses)
	index := sort.SearchStrings(courses, "MySQL")
	if index < len(courses) {
		courses = append(courses[:index], courses[index+1:]...)
	}
	fmt.Println("courses:", courses)
}
