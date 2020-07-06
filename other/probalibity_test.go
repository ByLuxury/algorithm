package other

import (
	"fmt"
	"testing"
)

// TestProbability ...
func TestProbability(t *testing.T) {
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
}
