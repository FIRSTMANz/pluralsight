package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

// func main() {
// 	fmt.Println("hello, FIRSTMAN!")

// }

// func main() {
// 	var i int
// 	i = 42
// 	fmt.Println(i)

// 	var f float32 = 3.14
// 	fmt.Println(f)

// 	firstName := "Sathianpong"
// 	fmt.Println(firstName)

// 	b := true
// 	fmt.Println(b)

// 	c := complex(3, 4)
// 	fmt.Println(c)

// 	r, im := real(c), imag(c)
// 	fmt.Println(r, im)

// }

//การใช้งาน pointer
// func main() {
// var firstName *string = new(string)

// *firstName = "Sathianpong"

// fmt.Println(*firstName)

// 	firstName := "Sathianpong"
// 	fmt.Println(firstName)

// 	ptr := &firstName
// 	fmt.Println(ptr , *ptr)

// }
// func main() {
// const pi = 3.1415
// println(pi)

// 	const c int = 3
// 	fmt.Println(c + 3)
// 	fmt.Println(float32(c) + 1.2)

// }

// const pi = 3.1415
// const (
// 	first = iota + 2
// 	second
// 	third
// )

// func main() {
// 	fmt.Println(first, second, third)
// }
// func main() {
// var arr [3]int
// arr[0] = 1
// arr[1] = 2
// arr[2] = 3
// fmt.Println(arr)

// arr2 := [3]int{1, 2, 3}

// slice := arr2[:]

// arr2[1] = 42
// slice[2] = 27

// fmt.Println(arr2, slice)
// slice := []int{1, 2, 3}
// fmt.Print(slice)

// slice = append(slice, 4, 26)
// fmt.Println(slice)

// s2 := slice[1:]
// s3 := slice[:2]
// s4 := slice[1:2]

// fmt.Println(s2, s3, s4)
// }

// func main() {
// 	m := map[string]int{"foo": 42}
// 	fmt.Println(m)
// 	fmt.Println(m["foo"])

// 	m["foo"] = 26
// 	fmt.Println(m)

// 	delete(m, "foo")
// 	fmt.Println(m)
// }

// func main() {
// 	type user struct {
// 		ID        int
// 		FirstName string
// 		LastName  string
// 	}
// 	var u user
// 	u.ID = 1
// 	u.FirstName = "Sathianpong"
// 	u.LastName = "Sangseema"
// 	fmt.Println(u)

// 	u2 := user{ID: 1,
// 		FirstName: "Sathianpong",
// 		LastName:  "Sangseema",
// 	}
// 	fmt.Println(u2)
// }
// func main() {
// 	u := models.User{
// 		ID:        2,
// 		FirstName: "Sathianpong",
// 		LastName:  "Sangseema",
// 	}
// 	fmt.Println(u)
// port := 3000
// err, bbb, _ := startWebServer(port)
// fmt.Println(err, bbb)
// }

// func startWebServer(port int) (error, bool, int) {
// 	fmt.Println("Starting server...")
//do important things
// fmt.Println("Server started on port", port)
// fmt.Println("Number of retries", numberOfRetries)
// return nil
// return errors.New("Something went wrong"), true, 1
// }

// func main() {
// 	controllers.RegisterControllers()
// 	http.ListenAndServe(":3000", nil)
// }
// type User struct {
// 	ID        int
// 	FirstName string
// 	LastName  string
// }

// func main() {
// 	u1 := User{
// 		ID:        1,
// 		FirstName: "Sathianpong",
// 		LastName:  "Sangseema",
// 	}
// 	u2 := User{
// 		ID:        2,
// 		FirstName: "test",
// 		LastName:  "test",
// 	}
// 	if u1 == u2 {
// 		println("Same User")
// 	} else if u1.FirstName == u2.FirstName {
// 		println("Similar user")
// 	}
// }
