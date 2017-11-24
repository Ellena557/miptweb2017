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

func PowerGenerator(a int) func() int {
	x := 1    
	return func() int {        
		x = x*y       
		return x    
	}
}

func DifferentWordsCount(str string) int{
	wrds = 0
	subst = ""
	dict := make(map[string]bool)
	for _, symb := range str{
		if unicode.IsLetter(symb) {
            		subst = subst + string(unicode.ToUpper(symb)) //Make all the letters big ones
		}
		if subst != "" {
			if !set[subst] {
                		wrds = wrds + 1
			}
		set[subst] = true
		subst = ""
		}
	}
	return wrds
}
