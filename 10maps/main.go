package main

import "fmt"

func main()  {
	langs := make(map[string]string)
	langs["js"] = "javascript"
	langs["rb"] = "ruby"
	langs["py"] = "python"

	fmt.Println(langs) // not comma separated
	fmt.Println(langs["js"])

	delete(langs, "rb")
	fmt.Println(langs)

	// looping in maps
	for key, value := range langs { // key, value is also comma ok syntax
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}