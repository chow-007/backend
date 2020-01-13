package utils

import (
	"fmt"
	"strings"
)

// 切片去重
func SliceDuplicateRemoval(elements []string) (result []string) {
	tempMap := make(map[string]string)
	for _, v := range elements {
		tempMap[v] = v
	}
	for _, v := range tempMap {
		result = append(result, v)
	}
	return
}

func GetSafetySqlOr(fields []string) string {
	var safetyFields []string
	for _, f := range fields {
		safetyFields = append(safetyFields, fmt.Sprintf("code = '%s'", f))
	}
	return fmt.Sprintf("( %s )", strings.Join(safetyFields, " OR "))
}


