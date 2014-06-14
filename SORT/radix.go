package main

import (
	//"container/list"
	"fmt"
	"math/rand"
	"os"
)

type List struct {
	data int
	next int
}

//
func maxbit(data []int) int {
	var nosignval int
	for _, v := range data {
		if v < 0 {
			v = -v
		}
		if nosignval < v {
			nosignval = v
		}
	}
	c := 1
	for nosignval/10 > 0 {
		c++
		nosignval = nosignval / 10
	}
	return c
}

func RadixSort(data []int) {
	n := len(data)
	p := maxbit(data)
	r := 1
	var a, b, c, k, g int
	var bucket, d []List
	for j := 0; j < n; j++ {
		bucket = append(bucket, List{next: -1, data: j})
		d = append(d, List{next: -1, data: data[j]})
	}
	fmt.Println("maxbit", p)
	for i := 0; i < p; i++ {
		if i != 0 {
			//rest bucket
			for j := 0; j < n; j++ {
				bucket[j].next = -1
				bucket[j].data = j
			}
			for j := 0; j < n; j++ {
				d[j].data = data[j]
				d[j].next = -1
			}
			for k = n - 1; k >= 0; k-- {
				a = d[k].data / r
				b = a % 10
				d[k].next = bucket[b].next
				bucket[b].next = k
			}
		} else {
			for k = 0; k < n; k++ {
				a = d[k].data / r
				b = a % 10
				d[k].next = bucket[b].next
				bucket[b].next = k
			}
		}
		c = 0
		for k = 0; k < n; k++ {
			if bucket[k].next != -1 {
				g = bucket[k].next
				data[c] = d[g].data
				c++
				for d[g].next != -1 {
					data[c] = d[d[g].next].data
					c++
					g = d[g].next
				}
			}
		}
		r = r * 10
	}
}
func main() {
	var list []int
	r := rand.New(rand.NewSource(1111))
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		//if r.Intn(2) == 1 {
		//	list = append(list, -r.Intn(10000))
		//} else {	
		list = append(list, r.Intn(120))
		//}
	}
	for _, v := range list {
		fmt.Println(v)
	}
	RadixSort(list)
	for _, v := range list {
		fmt.Println(v)
	}
	os.Exit(1)
}
