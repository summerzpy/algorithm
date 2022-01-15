package ac038

func Permutation( str string ) []string {
	res := make([]string, 0)
	if str  == "" {
		return res
	}
	PermutationHandler([]byte(str),0,&res)
	return res
}
//如果字符串中有重复字母，考虑将字符串先排序

func PermutationHandler(strByte []byte,index int, result *[]string) {
	length := len(strByte)
	if index<length {
		strMap := make(map[byte]bool)
		for j := index;j<length;j++ {
			if !strMap[strByte[j]] {
				//用来记录字符是否已经在第x位固定过
				strMap[strByte[j]] = true
				if j != index {
					strByte[index], strByte[j] =  strByte[j], strByte[index]
				}
				PermutationHandler(strByte, index+1, result)
				if j != index {
					strByte[index], strByte[j] =  strByte[j], strByte[index]
				}
			}
		}
	} else {
		*result = append(*result, string(strByte))
	}
}
