package heap

//SinkDown - swaps element at index with child downward the heap
//param  index {Number} - integer index
func (hp *Heap) SinkDown(index int) {

	var child1 interface{}
	// Look up the target element and its score.
	var length = len(hp.content)
	var element = hp.content[index]
	var found = false

	for !found {
		// Compute the indices of the child elements.
		var index_child2 = (index + 1) * 2
		var index_child1 = index_child2 - 1
		var swap_index = nilIndex
		// if first child exists (is inside the array)
		if index_child1 < length {
			// Look it up and compute its score.
			child1 = hp.content[index_child1]
			// If the score is less than our element"s, we need to swap_index.
			if hp.htype.IsMin() {
				if hp.cmp(child1,element) < 0 {
					swap_index = index_child1
				}
			} else {
				if hp.cmp(child1, element) > 0 {
					swap_index = index_child1
				}
			}
		}
		// Do the same checks for the other child.
		if index_child2 < length {
			var child2 = hp.content[index_child2]

			var temp_element interface{}
			//update swap_index candidate, with elemscore || ch1score
			if isnil(swap_index) {
				temp_element = element
			} else {
				temp_element = child1
			}

			if hp.htype.IsMin() {
				if hp.cmp(child2, temp_element) < 0 {
					swap_index = index_child2
				}
			} else {
				//max heap
				if hp.cmp(child2, temp_element) > 0 {
					swap_index = index_child2
				}
			}
		}
		// No need to swap_index further
		if isnil(swap_index) {
			found = true
		} else {
			// swap_index and continue.
			hp.content[index] = hp.content[swap_index]
			hp.content[swap_index] = element
			index = swap_index
		}
	}
}
