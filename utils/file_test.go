package utils

import (
	"os"
	"os/user"
	"path"
	"testing"
)

func TestIsFileExists(t *testing.T) {
	// SetUp
	cUser, err := user.Current()
	if err != nil {
		t.Fatalf("%v", err)
	}
	testDir := path.Join(cUser.HomeDir, "/test_gouno_file_exists/")
	err = os.MkdirAll(testDir, os.ModePerm)
	if err != nil {
		t.Fatalf("%v", err)
	}
	for _, i := range []string{"123", "中文.docx"} {
		_, err = os.Create(path.Join(testDir, i))
		if err != nil {
			t.Fatalf("%v", err)
		}
	}
	os.Mkdir(path.Join(testDir, "notFile"), os.ModePerm)

	var tests = []struct {
		input  string
		output bool
	}{
		{path.Join(testDir, "123"), true},
		{path.Join(testDir, "中文.docx"), true},
		{path.Join(testDir, "notFile"), false},
		{path.Join(testDir, "notExists"), false},
	}
	for _, test := range tests {
		r := IsFileExists(test.input)
		if r != test.output {
			t.Errorf("IsFileExists([%v])->[%v], But get:[%v]", test.input, test.output, r)
		}
	}
	os.RemoveAll(testDir)
}

func TestJoinCacheFile(t *testing.T) {
	dir := "/home/xsy/gouno/cache"
	p, _ := JoinCacheFile(dir, "123.doc")
	r := "/home/xsy/gouno/cache/123.doc"
	if p != r {
		t.Errorf("JoinCacheFile([%s]->[%s], But get:[%s]", dir, r, p)
	}
}
