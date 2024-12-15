package main

import "strings"

func addStrings(num1 string, num2 string) string {
	sum := strings.Builder{}
	carry := '0'

	for len(num1) > 0 || len(num2) > 0 {
		number1 := '0'
		if len(num1) != 0 {
			number1 = rune(num1[len(num1)-1])
			num1 = num1[:len(num1)-1]
		}

		number2 := '0'
		if len(num2) != 0 {
			number2 = rune(num2[len(num2)-1])
			num2 = num2[:len(num2)-1]
		}

		total := (number1 - '0') + (number2 - '0') + (carry - '0')
		if total > '9'-'0' {
			carry = '1'
			total += -'9' - 1 + '0'
		} else {
			carry = '0'
		}

		sum.WriteRune(total + '0')
	}
	if carry != '0' {
		sum.WriteRune(carry)
	}

	return reverse(sum.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}

// func addStrings(num1 string, num2 string) string {
//     res := ""

//     n1 := len(num1)
//     n2 := len(num2)
//     // 11
//     // 123
//     // n1 = 2
//     // n2 = 3

//     i := 1
//     // 3
//     hasAdditional := false

//     for i <= n1 && i <= n2 {
//         tmp1, _ := strconv.Atoi(string(num1[n1-i]))
//         tmp2, _ := strconv.Atoi(string(num2[n2-i]))
//         total := tmp1 + tmp2

//         if hasAdditional {
//             total++
//         }
//         hasAdditional = false

//         res = strconv.Itoa(total % 10) + res

//         if total/10 == 1 {
//             hasAdditional = true
//         }
//         i++
//     }
//     if i <= n1 {
//         if !hasAdditional {
//             res = num1[0:n1-i + 1] + res
//         } else {
//             for i <= n1 {
//                 tmp, _ := strconv.Atoi(string(num1[n1-i]))
//                 if hasAdditional {
//                     tmp++
//                 }
//                 hasAdditional = false

//                 res = strconv.Itoa(tmp % 10) + res

//                 if tmp/10 == 1 {
//                     hasAdditional = true
//                 }
//                 i++
//             }
//         }
//     }

//     if i <= n2 {
//         if !hasAdditional {
//             res = num2[0:n2-i + 1] + res
//         } else {
//             for i <= n2 {
//                 tmp, _ := strconv.Atoi(string(num2[n2-i]))
//                 if hasAdditional {
//                     tmp++
//                 }
//                 hasAdditional = false

//                 res = strconv.Itoa(tmp % 10) + res

//                 if tmp/10 == 1 {
//                     hasAdditional = true
//                 }
//                 i++
//             }
//         }
//     }
//     if hasAdditional {
//         res = "1" + res
//     }
//     return res
// }
