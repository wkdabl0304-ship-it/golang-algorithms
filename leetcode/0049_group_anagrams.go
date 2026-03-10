package main

func groupAnagrams(strs []string) [][]string {

	// 1.分类
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var cnt [26]int
		for _, char := range str {
			cnt[char-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}

	// 2.数据结构转换
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

// 解题思路：通过字母个数计数来找到异位词的共同点进行分类
