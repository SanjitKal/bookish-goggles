package memtable

// implementes memtable interface

type Memtable struct {
	capacity int
	size int
	// store tree
}

func (memt *Memtable) Init(capacity string) (err error) {
	return nil
}

func (memt *Memtable) Read(key string) (val string, err error) {
	return "", nil
}

func (memt *Memtable) Insert(key string, val string) (err error) {
	return nil
}

func (memt *Memtable) PopMin() (key string, val string, err error) {
	return "", "", nil
}

func (memt *Memtable) PopAll() (key []string, val []string, err error) {
	return nil, nil, nil
}

func (memt *Memtable) GetSize() (size int, err error) {
	return 0, nil
}

func(memt *Memtable) GetCapacity() (cap int, err error) {
	return 0, nil
} 
