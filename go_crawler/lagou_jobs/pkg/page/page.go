package page

import "math"

func CalculateTotalPage(totalCount, PageSize float64) int {
	totalPage := float64(totalCount)/float64(PageSize)
	return int(math.Ceil(totalPage))
}
