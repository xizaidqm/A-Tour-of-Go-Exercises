package main

//Exercise: Fibonacci closure
//Let's have some fun with functions.
//Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

import "fmt"

func fibonacci() func() int {

	//a,b分别代表前后两个加数

	//index用来判断是否是fibonacci数列中的前两个数

	a,b,index := 0,1,1

	return func() int {

		//如果是第一个数，直接返回0

		if index==1 {

			index+=1

			return 0

		}else if index==2{

			//如果是第二个数，直接返回1

			index+=1

			return 1

		}else{

			//从第三个数开始，为前两个数之和

			sum := a+b

			//交换数值，用于下一次调用时候的两个数相加求和，相当于数字的后移

			//注意！Go的数值交换可以不需要再开辟新的临时变量

			a,b = b,sum

			//返回本次调用的求和结果

			return sum

		}

	}

}

func main(){

	f := fibonacci()

	for i := 0; i<10; i++{

		fmt.Println(f()) // 0,1,1,2,3,5,8,13,21,34

	}
}
