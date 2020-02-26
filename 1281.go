package solution

import "strconv"

func subtractProductAndSum(n int) int {
     nstr := strconv.Itoa(n)
     mul := 1
     sum := 0
     for index, _ := range nstr {
     	 nint, _ := strconv.Atoi(string(nstr[index]))
         mul *= nint
         sum += nint
	 }
	 return mul - sum
}