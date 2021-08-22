package main
//Exercise: Slices
//Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program,
//it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

//The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
import (
	"golang.org/x/tour/pic"
)
//uint8是指无符号类型的8位整型的基本数值类型，其范围是0到255，是为了和图片的像素范围相贴合
func Pic(dx, dy int) [][]uint8 {
	//创建二维数组，指定第二维的长度为dy
	a := make([][] uint8, dy)
	for i := range a{
		//创建一维数组，用于给定二维数组a中的每一行一维数组
		b := make([]uint8, dx)
		for j := range b{
			//设定一维数组b中的具体数值，即最终图形的像素值
			//在这里设置的每一点的像素值为i*j，读者可以自行按照题目提示，
			//设置为(i+j)/2，i^j等多个不同数值，并观察到多种图片的差别
			b[j] = uint8(i*j)
		}
		a[i]=b
	}
	return a
}
func main() {
	pic.Show(Pic)
}