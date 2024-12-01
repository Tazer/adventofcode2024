package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {

	list1, list2 := GetFileData(true)

	assert.NotEmpty(t, list1)
	assert.NotEmpty(t, list2)

	dist := GetDistances(list1, list2)

	assert.Equal(t, 11, dist)

}

func TestDay1Part2(t *testing.T) {

	list1, list2 := GetFileData(true)

	assert.NotEmpty(t, list1)
	assert.NotEmpty(t, list2)

	match := GetMatchings(list1, list2)

	assert.Equal(t, 31, match)

}
