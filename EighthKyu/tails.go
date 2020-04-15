package EighthKyu

// TODO Some new animals have arrived at the zoo. The zoo keeper is concerned that perhaps the animals do not have
//  the right tails. To help her, you must correct the broken function to make sure that the second argument (tail),
//  is the same as the last letter of the first argument (body) - otherwise the tail wouldn't fit!
func isThisMyTale(body string, tail rune) bool {
	if rune(body[len(body)-1]) == tail {
		return true
	}
	return false
}
