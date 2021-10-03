package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
	"unicode/utf8"
)

func main() {
	puzzles := map[string]func(){
		"puzzle1": func() {
			var π = 22 / 7.0
			fmt.Println(π)
		},

		"puzzle2": func() {
			var m map[string]int
			fmt.Println(m["errors"])
		},

		"puzzle3": func() {
			city := "Kraków"
			fmt.Println(len(city))
		},

		"puzzle3 - cityu": func() {
			city := "Kraków"
			fmt.Println(utf8.RuneCountInString(city))
		},

		// "puzzle4 - nil.go": func() {
		// 	n := nil
		// 	fmt.Println(n) // -> "0"
		// },

		"puzzle5 - raw.go": func() {
			s := `a\tb`
			p := "a\tb"
			fmt.Println(s)
			fmt.Println(p)
		},

		"puzzle6 - time.go": func() {
			// Invalid version
			// timeout := 3
			// fmt.Println("before ")
			// time.Sleep(timeout * time.Millisecond)
			// fmt.Println("after")

			// Must cast timeout explicitly to multiply
			const timeout time.Duration = 3
			fmt.Println("before ")
			time.Sleep(timeout * time.Millisecond)
			fmt.Println("after")
		},

		"puzzle7 - float.go": func() {
			n := 1.1
			fmt.Println(n * n)
		},

		"puzzle8 - sleep_sort.go": func() {
			var wg sync.WaitGroup
			for _, n := range []int{3, 2, 1} {
				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(time.Duration(n) * time.Millisecond)
					fmt.Printf("%d ", n)
				}()
			}
			wg.Wait()
			fmt.Println()
			// Prints 1, 1, 1 because n is the closure
		},

		"puzzle8 - sleep_sort_param.go": func() {
			var wg sync.WaitGroup
			for _, n := range []int{3, 2, 1} {
				wg.Add(1)
				go func(arg int) {
					defer wg.Done()
					time.Sleep(time.Duration(arg) * time.Millisecond)
					fmt.Printf("%d ", arg)
				}(n)
			}
			wg.Wait()
			fmt.Println()
		},

		"puzzle8 - sleep_sort_scope.go": func() {
			var wg sync.WaitGroup
			for _, n := range []int{3, 2, 1} {
				n := n
				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(time.Duration(n) * time.Millisecond)
					fmt.Printf("%d ", n)
				}()
			}
			wg.Wait()
			fmt.Println()
		},

		"puzzle9 - time_eq.go": func() {
			t1 := time.Now()

			// serialize the time into json
			data, err := json.Marshal(t1)
			if err != nil {
				log.Fatal(err)
			}

			// unmarshall the serialized time
			var t2 time.Time
			if err := json.Unmarshal(data, &t2); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Does t1 == t2? -- ", t1 == t2)
		},

		"puzzle9 - time_eq_comparison.go": func() {
			t1 := time.Now()

			// serialize the time into json
			data, err := json.Marshal(t1)
			if err != nil {
				log.Fatal(err)
			}

			// unmarshall the serialized time
			var t2 time.Time
			if err := json.Unmarshal(data, &t2); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Does t1 == t2? -- ", t1.Equal(t2))
		},

		"puzzle10 - append.go": func() {
			a := []int{1, 2, 3}
			b := append(a[:1], 10)
			fmt.Printf("a=%v, b=%v\n", a, b)
		},
	}

	keys := make([]string, 0, len(puzzles))
	for key := range puzzles {
		keys = append(keys, key)
	}

	// sort.Strings(keys)

	for i := range keys {
		key := keys[i]
		element := puzzles[key]
		fmt.Println("Puzzle: ", key)
		fmt.Println("------------------------------")
		element()
		fmt.Println("")
	}
}
