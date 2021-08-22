package main

//Exercise: Errors
//Copy your Sqrt function from the earlier exercise and modify it to return an error value.

//Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

import (

	"fmt"

	"math"

)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{

	//根据提示，使用float64(e)来防止无限循环，原因下面再解释

	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))

}

func Sqrt(x float64)(float64, error){

	if x>=0 {

		return math.Sqrt(x),nil

	}else{

		//如果求平方根的是负数，那么抛出Error

		return 0, ErrNegativeSqrt(x)

	}

}

func main(){

	fmt.Println(Sqrt(2))

	fmt.Println(Sqrt(-2))

}