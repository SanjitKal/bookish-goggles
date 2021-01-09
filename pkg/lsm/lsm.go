package lsm

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"time"
)

// implements log structured merge tree interface

type SSTableInfo struct {
	fileName     string         // sstable filename
	index        map[string]int // map from key to it's offset in the sstable file
	creationTime int64          // time of creation for sstable file
}

type LogStructuredMergeTree struct {
	maxIndexSize int            // max number of keys in in memory index for an sstable
	ssTableList  []*SSTableInfo // sorted in increasing order of creation time
}

func (lsm *LogStructuredMergeTree) Init(maxIndexSize int) (err error) {
	err = lsm.UpdateMaxIndexSize(maxIndexSize)
	if err != nil {
		return err
	}
	lsm.ssTableList = make([]*SSTableInfo, 0)
	err := lsm.Load()
	if err != nil {
		return err
	}
	return nil
}

func (lsm *LogStructuredMergeTree) UpdateMaxIndexSize(maxIndexSize int) (err error) {
	if maxIndexSize < 0 && maxIndexSize != -1 {
		return errors.New("Max index size must either be a positive integer, or -1 to represent an unbounded max index size.")
	}
	lsm.maxIndexSize = maxIndexSize
	return nil
}

func (lsm *LogStructuredMergeTree) Lookup(key string) (val string, err error) {
	// Iterate through each sstable info struct in most recently written to least recently written order.
	// For each sstable, read the associated in memory index to determine the
	// range of values to read from the sstable file. Scan this range for the kv pair.
	return "", nil
}

func (lsm *LogStructuredMergeTree) WriteNewSSTable(sortedKeys []string, vals []string) (err error) {
	fileName := fmt.Sprintf("sst_%d.txt", len(lsm.ssTableList))
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0666)
	index := make(map[string]int)
	if err != nil {
		return errors.New(fmt.Sprintf("Issue creating sstable file: %s", err))
	}
	addIndexEntryFreq := 1
	if lsm.maxIndexSize != -1 && len(sortedKeys) > lsm.maxIndexSize {
		addIndexEntryFreq = int(len(sortedKeys) / lsm.maxIndexSize)
	}
	w := bufio.NewWriter(f)
	currentByteOffset := 0
	for i := 0; i < len(sortedKeys); i++ {
		entry := fmt.Sprintf("%d,%d,%s%s", len(sortedKeys[i]), len(vals[i]), sortedKeys[i], vals[i])
		entry = fmt.Sprintf("%d,%s", len(entry), entry)
		bytesWritten, err := w.WriteString(entry)
		if bytesWritten != len(entry) || err != nil {
			return errors.New(fmt.Sprintf("Issue writing entry to sstable file: %s", err))
		}
		if math.Mod(float64(i+1), float64(addIndexEntryFreq)) == 0 {
			index[sortedKeys[i]] = currentByteOffset
		}
		currentByteOffset = currentByteOffset + bytesWritten
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	newSSTableInfo := &SSTableInfo{fileName: fileName, creationTime: time.Now().Unix(), index: index}
	lsm.ssTableList = append(lsm.ssTableList, newSSTableInfo)
	return nil
}

func (lsm *LogStructuredMergeTree) Compact() (err error) {
	return nil
}

func (lsm *LogStructuredMergeTree) Merge() (err error) {
	return nil
}

func (lsm *LogStructuredMergeTree) Load() (err error) {
	// Populate the SSTable info list by reading the metadata file upon
	// initialization.
	return nil
}
