package utils

func Abs(num1, num2 int) int{
	if num1 > num2 {
		return num1-num2
	}
	return num2-num1
}