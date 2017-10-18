package heap

func (hp *Heap) Peek() interface{} {
	if hp.IsEmpty() {
		return nil
	}
	return hp.content[0]
}
