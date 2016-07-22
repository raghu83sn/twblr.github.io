package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	var result []int32
	for _, element := range vals {
		result = append(result, op(element))
	}
	return result

}

func filterInts(op filterOperation, vals []int32) []int32 {
	var result []int32
	for _, element := range vals {

		if op(element) {
			result = append(result, element)
		}

	}
	return result
}

func concatenate(dest []string, newValues ...string) []string {
	return nil
}

func equals(list1 []string, list2 []string) bool {
	return false
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
