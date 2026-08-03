package main

import "fmt"

//func main() {
//	temps := map[string]map[string]float64{
//		"Moscow": {
//			"Jan": -5.1,
//			"Feb": -7,
//			"Mar": 9.7,
//			///
//		},
//		"Spb": {
//			"Jan": -5.1,
//			"Feb": -7,
//			"Mar": 9.7,
//			///
//		},
//	}
//
//	fmt.Println(temps)
//}

//func main() {
//	temps := map[string]map[string]map[int]float64{
//		"Moscow": {
//			"Jan": {
//				1: -5.1,
//				2: -5.1,
//			},
//		},
//		"Spb": {
//			"Jan": {
//				1: -5.1,
//				2: -5.1,
//			},
//		},
//	}
//
//	fmt.Println(temps)
//}

func main() {
	temps := map[string]map[string]map[string]map[int]float64{
		"Russia": {
			"Moscow": {
				"Jan": {
					1: -5.1,
					2: -5.1,
				},
			},
			"Spb": {
				"Jan": {
					1: -5.1,
					2: -5.1,
				},
			},
		},
	}

	temps["Spain"] = map[string]map[string]map[int]float64{
		"Barca": {
			"Jan": {
				1: 5.1,
				2: 5.1,
			},
		},
	}

	moscowJan1 := temps["Russia"]["Moscow"]["Jan"][1]

	fmt.Println(temps)
	fmt.Println(moscowJan1)
}
