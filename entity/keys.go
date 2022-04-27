package entity

type Key struct {
	KeyValue int
}

func NewKey(keyValue int) *Key {
	return &Key{
		KeyValue: keyValue,
	}
}

func GetKeyPack(lastVal, amount int) ([]*Key, int, error) {
	keys := make([]*Key, 0, amount)
	lastKey := lastVal
	for i := lastVal + 1; i <= amount; i++ {
		currKey := NewKey(i)
		lastKey = lastKey + 1
		keys = append(keys, currKey)
	}

	return keys, lastKey, nil

}
