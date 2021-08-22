package main

//Exercise: Loops and Functions
//As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.
import "fmt"
func Sqrt(x float64) float64 {
	z := 1.0
	for {
		//判断z*z与x的大小关系
		if z*z > x {
			//当z*z-x在既定的容忍区间内，那么就可以认为z可以代替我们所求的平方根的真实值
			if z*z-x <= 1e-5 {
				return z
			} else {
				//否则，将z减半
				z /= 2
			}
		} else if z*z == x {
			return x
		} else {
			if x-z*z <= 1e-5 {
				return z
			} else {
				z = z + z/2
			}
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
