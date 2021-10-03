package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"
	"unicode/utf8"
)

func main() {
	puzzles := map[string]func(){
		"puzzle01": func() {
			var π = 22 / 7.0
			fmt.Println(π)
		},

		"puzzle02": func() {
			var m map[string]int
			fmt.Println(m["errors"])
		},

		"puzzle03": func() {
			city := "Kraków"
			fmt.Println(len(city))
		},

		"puzzle03 - cityu": func() {
			city := "Kraków"
			fmt.Println(utf8.RuneCountInString(city))
		},

		// "puzzle004 - nil.go": func() {
		// 	n := nil
		// 	fmt.Println(n) // -> "0"
		// },

		"puzzle05 - raw.go": func() {
			s := `a\tb`
			p := "a\tb"
			fmt.Println(s)
			fmt.Println(p)
		},

		"puzzle06 - time.go": func() {
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

		"puzzle07 - float.go": func() {
			n := 1.1
			fmt.Println(n * n)
		},

		"puzzle08 - sleep_sort.go": func() {
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

		"puzzle08 - sleep_sort_param.go": func() {
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

		"puzzle08 - sleep_sort_scope.go": func() {
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

		"puzzle09 - time_eq.go": func() {
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

		"puzzle09 - time_eq_comparison.go": func() {
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
			c := []int{1, 2, 3}
			d := append(c[:2], 10)
			fmt.Printf("a=%v, b=%v\n", a, b)
			fmt.Printf("c=%v, d=%v\n", c, d)
		},

		"puzzle11 - struct.go": func() {
			type Log struct {
				Message  string
				LoggedAt time.Time
			}

			ts := time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC)
			log := Log{"Hello", ts}
			fmt.Printf("%v\n", log)
		},

		"puzzle12 - a_funky_number.go": func() {
			// hexadecimal floating point
			fmt.Println(0x1p-2)
		},

		"puzzle13 - range.go": func() {
			// fibs := func(n int) chan int {
			// 	ch := make(chan int)

			// 	go func() {
			// 		a, b := 1, 1
			// 		for i := 0; i < n; i++ {
			// 			ch <- a
			// 			a, b = b, a+b
			// 		}

			// 	}()
			// 	return ch
			// }

			// for i := range fibs(5) {
			// 	fmt.Printf("%d", i)
			// }

			// fmt.Println()
			// Code deadlocks

			// Good Version
			fibs := func(n int) chan int {
				ch := make(chan int)

				go func() {
					// Defer so we make sure
					defer close(ch)

					a, b := 1, 1
					for i := 0; i < n; i++ {
						ch <- a
						a, b = b, a+b
					}

				}()
				return ch
			}

			for i := range fibs(5) {
				fmt.Printf("%d", i)
			}

			fmt.Println()
			// Code deadlocks
		},

		"puzzle13 - range_rcv.go": func() {
			fibs := func(ctx context.Context, n int) chan int {
				ch := make(chan int)

				go func() {
					// Defer so we make sure
					defer close(ch)
					a, b := 1, 1
					for i := 0; i < n; i++ {
						ch <- a
						a, b = b, a+b
					}

				}()
				return ch
			}

			ctx, cancel := context.WithCancel(context.Background())
			ch := fibs(ctx, 5)
			for i := 0; i < 5; i++ {
				val := <-ch
				fmt.Printf("%d", val)
			}

			fmt.Println()
			cancel()
		},
	}

	keys := make([]string, 0, len(puzzles))
	for key := range puzzles {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for i := range keys {
		key := keys[i]
		element := puzzles[key]
		fmt.Println("Puzzle: ", key)
		fmt.Println("------------------------------")
		element()
		fmt.Println("")
	}
}
