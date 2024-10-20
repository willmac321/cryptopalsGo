package utils

func GetEnd(offset int, index int, arrSize int) int {
	end := index + offset
	if end > arrSize {
		end = arrSize
	}
	return end
}
