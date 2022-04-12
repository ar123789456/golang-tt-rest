package usecase

type SubstrUseCase struct {
}

func NewSubstrUseCase() *SubstrUseCase {
	return &SubstrUseCase{}
}

func (s *SubstrUseCase) FindLongestSubstring(text string) string {
	posMap := map[byte]int{}
	pointer1, pointer2 := 1, 1

	result := ""
	for pointer2 <= len(text) {
		if posMap[text[pointer2-1]] == 0 {
			posMap[text[pointer2-1]] = pointer2
			if len(result) < len(text[pointer1-1:pointer2]) {
				result = text[pointer1-1 : pointer2]
			}
			pointer2++
		} else {
			delete(posMap, text[pointer1-1])
			pointer1++
		}
	}
	return result
}
