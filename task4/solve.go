package main

import "unicode"

func RemoveEven(arr[] int) []int{
	result := make([]int, 0)
        for _, element := range arr {
		if element % 2 > 0 {
			result = append(result, element)
		}
        }
        return result
}

func PowerGenerator(y int) func() int {
	x := 1    
	return func() int {        
		x = x*y       
		return x    
	}
}

func DifferentWordsCount(str string) int{
	wrds := 0
	subst := ""
	extr := make(map[string]bool)
	str2 := str + " ";
	for _, symb := range str2{
		if unicode.IsLetter(symb) {
            		subst = subst + string(unicode.ToLower(symb)) //Make all the letters little ones
		}
		if subst != "" {
			if !extr[subst] {
                		wrds = wrds + 1
			}
		extr[subst] = true
		subst = ""
		}
	}
	return wrds
}
