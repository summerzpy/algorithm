package problem046

import "strconv"

/**
 * 解码
 * @param nums string字符串 数字串
 * @return int整型
举例：10->
 * 动态规划，类似跳台阶的题目 状态 状态转移方程
 * f(i)= f(i+1) + f(i+2) s[0]s[1] 属于[10,25]
         f(i+1) f否则
       = f(i+1) + g(i,i+1)f(i+2) -> if s[0]s[1] 属于[10,25],g(i,i+1) = 1;else 0
*/

func Solve(num int) int {
	if num<10{
		return 1
	}
	if num>=10 &&num<=25{
		return 2
	}
	str:=strconv.Itoa(num)
	dp:=make([]int,len(str))
	dp[0]=1    //边界
	if str[:2]>="10"&&str[:2]<="25"{    //边界
		dp[1]=2
	}else{
		dp[1]=1
	}
	for i:=2;i<len(str);i++{
		newnum:=str[i-1:i+1]
		if newnum>="10" && newnum<="25"{
			dp[i]=dp[i-1]+dp[i-2]     //递归关系式
		}else{
			dp[i]=dp[i-1]
		}
	}
	return dp[len(str)-1]
}
