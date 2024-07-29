package main


// factors are numbers you can multiply together to get another number 
// 2 and 3 are factors of 6 because 2 * 3 = 6
// if you divide numbers and the remainder is 0 then it is a factor 

// function that checks if the factor is a factor of base 
func IsFactor(base int, factor int) bool {
    return base % factor == 0
}