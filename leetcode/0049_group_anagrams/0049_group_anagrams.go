package main

import (
	"fmt"
)

// 心得体会
// 1.数组可以做 map 的 key
// 2.range 字符串因为 UTF-8 解码成本会更久（range 时是以字符为单位，而不是字节）

func groupAnagrams(strs []string) [][]string {

	//1.排序
	//m := make(map[string][]string)
	//for _, v := range strs {
	//	b := []byte(v)
	//	sort.Slice(b, func(i, j int) bool {
	//		return b[i] < b[j]
	//	})
	//	key := string(b)
	//	m[key] = append(m[key], v)
	//}
	//var res [][]string
	//for _, v := range m {
	//	res = append(res, v)
	//}
	//return res

	//2.字符计数（for range）
	//m := make(map[[26]int][]string)
	//for _, v := range strs {
	//	cnt := [26]int{}
	//	for _, vv := range v {
	//		cnt[vv-'a']++
	//	}
	//	m[cnt] = append(m[cnt], v)
	//}
	//var res [][]string
	//for _, v := range m {
	//	res = append(res, v)
	//}
	//return res

	//3.字符计数（直接for）
	m := make(map[[26]int][]string)

	for _, s := range strs {
		cnt := [26]int{}
		for i := 0; i < len(s); i++ {
			cnt[s[i]-'a']++
		}
		m[cnt] = append(m[cnt], s)
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	res := groupAnagrams(strs)

	fmt.Println(res)
}
