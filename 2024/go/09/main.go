package main

import (
	"fmt"
	"os"
)

func ParseInput(filepath string) []SpaceDisk {
	data, _ := os.ReadFile(filepath)

	result := []SpaceDisk{}

	id := 0
	start := 0
	for i := 0; i < len(data); i += 2 {
		store := int(data[i] - 48)
		free := 0
		if i+1 >= len(data) {
			free = 0
		} else {
			free = int(data[i+1] - 48)
		}

		result = append(result, SpaceDisk{id, start, store, free})
		id += 1
		start += store + free
	}

	return result
}

type SpaceDisk struct {
	Id    int
	Start int
	Store int
	Free  int
}

type SpaceSlot struct {
	Id *int
}

func computeChecksum(input []SpaceSlot) int {
	result := 0

	for i := 0; i < len(input); i++ {
		if input[i].Id != nil {
			result += i * *input[i].Id
		}
	}

	return result
}

func SpaceSlotsToString(input []SpaceSlot) string {
	result := ""
	for _, slot := range input {
		if slot.Id == nil {
			result += "."
		} else {
			result += fmt.Sprintf("|%d|", *slot.Id)
		}
		result += "\n"
	}
	return result
}

func ParseSpaceSlot(input []SpaceDisk) []SpaceSlot {
	result := []SpaceSlot{}
	for _, disk := range input {
		for i := 0; i < disk.Store; i++ {
			result = append(result, SpaceSlot{&disk.Id})
		}
		for i := 0; i < disk.Free; i++ {
			result = append(result, SpaceSlot{nil})
		}
	}
	return result
}

func findFreeIndex(input []SpaceSlot, startFrom int) int {
	for i := startFrom; i < len(input); i++ {
		if input[i].Id == nil {
			return i
		}
	}
	return -1
}

func findLastIndex(input []SpaceSlot, startFrom int) int {
	for i := startFrom; i >= 0; i-- {
		if input[i].Id != nil {
			return i
		}
	}
	return -1
}

func SolutionPartOne(input []SpaceDisk) int {
	disks := ParseSpaceSlot(input)

	freeIndex := findFreeIndex(disks, 0)
	lastIndex := findLastIndex(disks, len(disks)-1)

	for {
		if freeIndex >= len(disks) {
			break
		}
		if lastIndex < 0 {
			break
		}

		// swap the freeIndex with the lastIndex
		disks[freeIndex], disks[lastIndex] = disks[lastIndex], disks[freeIndex]
		freeIndex += 1
		if disks[freeIndex].Id != nil {
			freeIndex = findFreeIndex(disks, freeIndex)
		}
		lastIndex -= 1
		if disks[lastIndex].Id == nil {
			lastIndex = findLastIndex(disks, lastIndex)
		}

		if freeIndex >= lastIndex {
			break
		}
	}

	return computeChecksum(disks)
}

func PrintOutput(input []SpaceDisk, output []SpaceDisk) {
	fmt.Println("Input:")
	for _, disk := range input {
		fmt.Printf("%d(start:%d, store:%d, free:%d)\n", disk.Id, disk.Start, disk.Store, disk.Free)
	}

	fmt.Println("Output:")
	for _, disk := range output {
		fmt.Printf("%d(start:%d, store:%d, free:%d)\n", disk.Id, disk.Start, disk.Store, disk.Free)
	}
	fmt.Println()
}

func SolutionPartTwo(input []SpaceDisk) int {
	output := make([]SpaceDisk, len(input))
	copy(output, input)

	i := len(input) - 1
	for i > 0 {
		diskI := output[i]

		for j := 0; j < i; j++ {
			diskJ := output[j]
			if diskJ.Free > 0 && diskJ.Free >= diskI.Store {
				if j == i-1 {
					diskI.Free = diskJ.Free + diskI.Free
					diskI.Start = diskJ.Start + diskJ.Store
					diskJ.Free = 0

					output[j] = diskJ
					output[i] = diskI

					break
				} else {
					prevDiskI := output[i-1]

					prevDiskI.Free += diskI.Store + diskI.Free

					diskI.Start = diskJ.Start + diskJ.Store
					diskI.Free = diskJ.Free - diskI.Store

					diskJ.Free = 0

					output = append(output[:i], output[i+1:]...)
					output = append(output[:j+1], append([]SpaceDisk{diskI}, output[j+1:]...)...)
					output[j] = diskJ
					output[i] = prevDiskI

					i += 1
					break
				}
			}
		}

		i -= 1
	}

	return computeChecksum(ParseSpaceSlot(output))
}
