package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	fmt.Println(WhatIsThe)
	WhatIsThe = 0
}

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}
