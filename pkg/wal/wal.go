package wal

// implements write ahead log interface

type WriteAheadLog struct {
	filename string
	// store tree
}

func (wal *WriteAheadLog) Init(filename string) (err error) {
	return nil
}

func (wal *WriteAheadLog) ReadAll() (ops []string, err error) {
	return nil, nil
}

func (wal *WriteAheadLog) Append(op string) (err error) {
	return nil
}

func (wal *WriteAheadLog) Clear() (err error) {
	return nil
}
