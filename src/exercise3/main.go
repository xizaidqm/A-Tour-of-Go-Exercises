package main

//Exercise: Maps
//Implement WordCount. It should return a map of the counts of each “word” in the string s.
//The wc.Test function runs a test suite against the provided function and prints success or failure.

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	//将字符串以空格格外，形成单词数组

	words := strings.Fields(s)

	ans := make(map[string]int)

	//遍历单词数组

	for _,w := range words{

		//如果该单词出现过，则对应的次数+1

		if ans[w]!=0{

			ans[w]+=1

		}else{

			ans[w]=1

		}

	}

	return ans
}


func main() {

	wc.Test(WordCount)

}