package prime_number

// func FindPrimeNumber(num int, out chan<- int) {
// 	for j := 2; j <= num; j++ {
// 		var flag = 0
// 		for i := 2; i <= j/2; i++ {
// 			if j%i == 0 {
// 				flag = 1
// 				break
// 			}
// 		}
// 		if flag == 0 {
// 			fmt.Println(j)
// 			out <- j
// 		}
// 	}
// }

//
func Generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
