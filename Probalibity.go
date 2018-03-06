package main

import (
	"math/rand"
	"fmt"
)

func main() {
	test := []float64{8.9, 1.2, 3.7}  //测试元素
	result := make([]float64, 0)      //根据概率随机出现的元素
	countMap := make(map[float64]int) //用map统计每个元素出现的次数
	for i := 0; i < 100; i++ { //循环执行100次看每个元素出现的概率
		index := cala(test)
		result = append(result, test[index])
		fmt.Print(test[index], " ")
	}
	for _, v := range result { //统计次数
		if countMap[v] != 0 {
			countMap[v] ++
		} else {
			countMap[v] = 1
		}
	}
	fmt.Println()
	for k, v := range countMap {
		fmt.Printf("%f%s%d%s\n", k, "出现", v, "次")
	}
	//fmt.Println(countMap)

}

/**
计算元素出现的次数
生成一个0-sum(data)之间的分布的随机数，然后遍历当前的list
每次循环中累加list的当前值，并且与改随机数进行比较，最终找到该数(也即是随机出现的这个数)
return 返回的是找到该数的索引
 */
func cala(data []float64) int {
	//生成一个0-sum(data)之间的分布的随机数
	randNum := rand.Float64() * sum(data)
	sumTmp := 0.00
	for i := 0; i < len(data); i++ {
		sumTmp += float64(data[i])
		if sumTmp > randNum {
			return i
		}
	}
	return -1
}

/**
求和
 */
func sum(data []float64) (sum float64) {
	for _, v := range data {
		sum += v
	}
	return sum
}
