package day09

import (
	"slices"
	"strconv"
)

type block struct {
	id   int
	pos  int
	size int
}

type fileSystem struct {
	files []block
	free  []block
}

func (fs *fileSystem) storeFile(id int, pos int, size int) {
	fs.files = append(fs.files, block{id: id, pos: pos, size: size})
}

func (fs *fileSystem) indicateFree(pos int, size int) {
	fs.free = append(fs.free, block{id: -1, pos: pos, size: size})
}

func (fs *fileSystem) compact() {
	// Start from last file, move each file to first available free slot
	var compactedFiles []block
	for fileIndex := len(fs.files) - 1; fileIndex > 0; fileIndex-- {
		file := fs.files[fileIndex]
		for i := range fs.free {
			free := &fs.free[i]
			if free.pos >= file.pos {
				continue
			}
			if free.size >= file.size {
				fs.files = slices.Delete(fs.files, fileIndex, fileIndex+1)
				compactedFiles = append(compactedFiles, block{id: file.id, pos: free.pos, size: file.size})
				free.size -= file.size
				free.pos += file.size
				if free.size == 0 {
					fs.free = slices.Delete(fs.free, i, i+1)
				}
				break
			}
		}
	}
	fs.files = append(fs.files, compactedFiles...)
}

func (fs fileSystem) checksum() int {
	checksum := 0
	for _, file := range fs.files {
		for i := 0; i < file.size; i++ {
			checksum += file.id * (file.pos + i)
		}
	}
	return checksum
}

func SolvePart1(input string) int {
	fileSystem := parseFileSystem(input, true)
	fileSystem.compact()
	return fileSystem.checksum()
}

func SolvePart2(input string) int {
	fileSystem := parseFileSystem(input, false)
	fileSystem.compact()
	return fileSystem.checksum()
}

func parseFileSystem(input string, parseIntoSingleBlocks bool) fileSystem {
	var fs fileSystem
	isBlockCount := true
	fileId := 0
	blockIndex := 0
	for _, rune := range input {
		num, _ := strconv.Atoi(string(rune))
		if isBlockCount {
			if parseIntoSingleBlocks {
				// For part 1, split files into single blocks
				for i := 0; i < num; i++ {
					fs.storeFile(fileId, blockIndex+i, 1)
				}
			} else {
				fs.storeFile(fileId, blockIndex, num)
			}
			fileId++
		} else {
			fs.indicateFree(blockIndex, num)
		}

		isBlockCount = !isBlockCount
		blockIndex += num
	}
	return fs
}
