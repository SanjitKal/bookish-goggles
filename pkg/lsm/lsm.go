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

type IndexEntry struct {
	key string
	offset uint64
}

type SSTableInfo struct {
	fileName     string         // sstable filename
	index []*IndexEntry  // Array of sorted keys and their offsets with density 
	creationTime int64          // time of creation for sstable file
}

type LogStructuredMergeTree struct {
	// indexDensity==1 -> all keys stored in the index
	// 0<=indexDensity<= 1 -> indexDensity*num_keys keys stored in index
	indexDensity float64
	ssTableList  []*SSTableInfo // sorted in increasing order of creation time
}

func (lsm *LogStructuredMergeTree) Init(indexDensity float64) (err error) {
	err = lsm.UpdateIndexDensity(indexDensity)
	if err != nil {
		return err
	}
	lsm.ssTableList = make([]*SSTableInfo, 0)
	err = lsm.Load()
	if err != nil {
		return err
	}
	return nil
}

func (lsm *LogStructuredMergeTree) UpdateIndexDensity(indexDensity float64) (err error) {
	if indexDensity < 0 || indexDensity > 1 {
		return errors.New("Index density must be a float between 0 and 1")
	}
	lsm.indexDensity = indexDensity
	return nil
}

func (lsm *LogStructuredMergeTree) Lookup(key string) (val string, err error) {
	// Attempt to look for the key in each of the sstable in most recently written
	// to least recently written order
	for _, ssTable := range lsm.ssTableList {
		present, deleted, val, err := lsm.LookupInSSTable(key, ssTable)
		if err != nil {
			return "", fmt.Errorf("Error reading from sstable: %s", err)
		} else if !present {
			continue
		} else if deleted {
			return "", fmt.Errorf("Key was deleted from kv store")
		} else {
			return val, nil
		}
	}
	return "", fmt.Errorf("Key %s not found in sstable", key)
}

func (lsm *LogStructuredMergeTree) LookupInSSTable(key string, ssTable *SSTableInfo) (present bool, deleted bool, val string, err error) {
	// Use in memory index to determine what indexed keys to search between on disk
	inRange, leftOffset, rightOffset := lsm.GetKeyToRangeToRead(key, ssTable.index)
	if !inRange {
		return false, false, "", nil
	}
	// Search the range provided by getKeyToRangeToRead on disk for the key
	return lsm.ScanRangeForKey(ssTable, leftOffset, rightOffset)
}

func (lsm *LogStructuredMergeTree) GetKeyToRangeToRead(key string, index []*IndexEntry) (inRange bool, leftOffset int64, rightOffset int64) {
	return false, 0, -1
}

func (lsm *LogStructuredMergeTree) ScanRangeForKey(ssTable *SSTableInfo, leftOffset int64, rightOffset int64) (present bool, deleted bool, val string, err error) {
	return false, false, "", nil
}

func (lsm *LogStructuredMergeTree) WriteNewSSTable(sortedKeys []string, vals []string) (err error) {
	fileName := fmt.Sprintf("sst_%d.txt", len(lsm.ssTableList))
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0666)
	index := make([]*IndexEntry, 0)
	if err != nil {
		return errors.New(fmt.Sprintf("Issue creating sstable file: %s", err))
	}

	addIndexEntryFreq := int(1 / lsm.indexDensity)

	w := bufio.NewWriter(f)

	var currentByteOffset uint64 = 0

	for i := 0; i < len(sortedKeys); i++ {
		entry := fmt.Sprintf("%d%d%s%s", len(sortedKeys[i]), sortedKeys[i], len(vals[i]), vals[i])
		bytesWritten, err := w.WriteString(entry)
		if bytesWritten != len(entry) || err != nil {
			return errors.New(fmt.Sprintf("Issue writing entry to sstable file: %s", err))
		}
		if math.Mod(float64(i+1), float64(addIndexEntryFreq)) == 0 {
			index = append(index, &IndexEntry{key: sortedKeys[i], offset: currentByteOffset})
		}
		currentByteOffset = currentByteOffset + uint64(bytesWritten)
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
