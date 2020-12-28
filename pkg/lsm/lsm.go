package lsm

// implements log structured merge tree interface

type LogStructuredMergeTree struct {
	capacity int
	size int
	// store levels of files
}

func (lsm *LogStructuredMergeTree) Init() (err error) {
	return nil
}

func (lsm *LogStructuredMergeTree) Insert(keys []string, vals []string) (err error) {
	return nil
}

func (lsm *LogStructuredMergeTree) Compact() (err error) {
	return nil
}