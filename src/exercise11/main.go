package main

import(

	"sync"

	"fmt"

	"time"

)

type Fetcher interface{

	//Fetch returns the body of URL and

	//a slice of URLs found on that page.

	Fetch(url string) (body string, urls[] string, err error)

}

//Crawl uses fetcher to recursively crawl

//pages starting with url, to a maximum of depth.

//由于爬取的过程是多个线程并发进行的，因此我们对该方法进行了改造

//传入了MutexSignal变量，用来控制Map的并发

func Crawl(signal *MutexSignal, url string, depth int, fetch Fetcher){

	//TODO: Fetch URLs in parallel

	//TODO: Don't fetch the same URL twice.

	//This implementation doesn't do either:

	//达到最大深度则退出

	if depth <= 0{

		return

	}

	//获取Fetcher的实际内容

	body, urls, err := fetcher.Fetch(url)

	if err != nil{

		fmt.Println(err)

		return

	}

	fmt.Printf("found: %s %q\n", url, body)

	for _,u := range urls{

		//获取Signal变量的锁，获取对signal.Map的控制权

		signal.m.Lock()

		//判断当前地址是否在map中存在

		if _,ok := signal.v[u];!ok{

			//在map中不存在该键时，对map进行写操作

			signal.v[u]=true

			//继续递归并发爬取地址，同时深度减1

			go Crawl(signal,u,depth-1,fetcher)

		}

		//在完成写操作后，释放锁资源

		signal.m.Unlock()

	}

	return

}

type MutexSignal struct{

	m sync.Mutex

	v map[string]bool

}

func main(){

	signal := &MutexSignal{v:make(map[string]bool)}

	//将初始的爬取地址设为起始点

	signal.v["https://golang.org"]=true

	Crawl(signal,"https://golang.org",4,fetcher)

	//需要设定时间的延迟，否则main函数结束了，其他线程的爬取还未结束

	time.Sleep(100)

}

type fakeResult struct{

	body string

	urls []string

}

//fakeFetcher is Fetcher that returns canned results.

type fakeFetcher map[string] *fakeResult


func(f fakeFetcher) Fetch(url string) (string, []string, error){

	if res, ok := f[url]; ok{

		return res.body, res.urls, nil

	}

	return "",nil, fmt.Errorf("not found: %s",url)

}

// fetcher is a populated fakeFetcher.


var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
