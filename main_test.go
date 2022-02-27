package main

import (
	"testing"
)

func TestIsTmpFile(t *testing.T) {
	partterns := []struct {
		filename string
		expected bool
	}{
		{"hoge.tmp", true},
		{"hoge.fuga.tmp", true},
		{"hoge.jpg", false},
	}
	for idx, pattern := range partterns {
		actual := IsTmpFile(pattern.filename)
		if pattern.expected != actual {
			t.Errorf("pattern %d: want %v, actual %v", idx, pattern.expected, actual)
		}
	}
}

func TestListFiles(t *testing.T) {
	dir := "image/"

	patterns := []struct {
		expected string
	}{
		{dir + "0.tmp"},
		{dir + "1.tmp"},
		{dir + "2.tmp"},
		{dir + "3.tmp"},
		{dir + "4.tmp"},
		{dir + "5.tmp"},
		{dir + "6.tmp"},
		{dir + "7.tmp"},
		{dir + "8.tmp"},
	}

	// fmt.Println(ListFiles(dir))

	actual, err := ListFiles(dir)
	if err != nil {
		t.Errorf("Not found:%s", dir)
	}

	for idx, pattern := range patterns {
		if pattern.expected != actual[idx] {
			t.Errorf("pattern %d: want %s, actual %s", idx, pattern.expected, actual[idx])

		}
	}
}
