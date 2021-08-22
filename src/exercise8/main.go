package main

//Exercise: rot13Reader
//A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

//For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.
//Reader (a stream of the decompressed data).

//Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher
//to all alphabetical characters.

//The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

import (

	"io"

	"os"

	"strings"

)

type rot13Reader struct {

	r io.Reader

}

// rot13字母替换

func rot13(b byte) byte{

	switch {

	//字符在字母表的前13个数中，则加13变成新的字母输出

	case 'A'<=b && b<='M':

		b += 13

	//字符在字母表的后13个字母中，则输出当前位置向前推13个次序的字母

	case 'M'<b && b<='Z':

		b -= 13

	//小写字母也同理

	case 'a'<=b && b<='m':

		b += 13

	case 'm'<b && b<='z':

		b -= 13

	}

	return b

}

//实现装饰者，用方法接收器包装io.Reader实现方法的加强

func (rot rot13Reader) Read(b []byte) (int,error){

	num,err := rot.r.Read(b)

	for i, value := range b {

		//调用rot13加密

		b[i] = rot13(value)

	}

	return num, err

}

func main() {

	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)

}