package main

//Exercise: Equivalent Binary Trees
//1. Implement the Walk function.

//2. Test the Walk function.

//The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

//Create a new channel ch and kick off the walker:

//go Walk(tree.New(1), ch)
//Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

//3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

//4. Test the Same function.

//Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

import(

	"golang.org/x/tour/tree"

	"fmt"

)

//Walk walks the tree t sending all values

//from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int){

	//需要在程序退出前关闭channel通道，否则会出现通道一直在等待数据导致死锁

	defer close(ch)

	//声明一个函数变量，方便进行递归调用

	var walk func(t *tree.Tree)

	//使用先序遍历的方式读取二叉树的每一个叶子节点

	//由于树中叶子节点的数值是有序的，但结构是随机的，也只能通过 先序遍历 得到有序的数值结果

	walk = func(t *tree.Tree){

		if t== nil{

			return

		}

		walk(t.Left)

		//将叶子节点的值放入channel

		ch <- t.Value

		walk(t.Right)

	}

	//实际调用该递归函数

	walk(t)

}

//Same determines whether the trees

//t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool{

	ch1:=make(chan int)

	ch2:=make(chan int)

	//并发遍历两棵树的内容

	go Walk(t1,ch1)

	go Walk(t2,ch2)

	for{

		//依次取出channel中的数值

		v1,ok1 := <-ch1

		v2,ok2 := <-ch2

		//如果两个通道中的数值数量不一致，那么就一定不会是相等的树

		if !ok1 || !ok2{

			return ok1==ok2

		}

		//如果数值不相等也不是相等的树

		if v1!=v2{

			return false

		}

	}

	return true

}

func main(){

	fmt.Println(Same(tree.New(1), tree.New(1))) //true

	fmt.Println(Same(tree.New(1), tree.New(2))) //false

}