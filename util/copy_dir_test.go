package util_test

import (
	"github.com/BobbyShaftoe/go-shell-replace/util"
	"os"
	"testing"
)

var dirs = []struct {
	src      string
	dst      string
	dirs     string
	expected string
}{
	{"src", "dst", "dir/dir1", "dst/dir/dir1"},
	{"src", "dst", "dir/dir2/dir3", "dst/dir/dir2/dir3"},
}

func TestCopyDir(t *testing.T) {
	// Test copy directories
	var err error
	for _, r := range dirs {
		src, dst, dirs, expect := r.src, r.dst, r.dirs, r.expected

		os.MkdirAll(src+"/"+dirs, 0755)
		os.MkdirAll(dst, 0755)

		c := &util.CopyDirArgs{Src: src, Dst: dst, Mode: 0755, IgnoreDot: true}
		if err = c.CopyDir(); err != nil {
			t.Errorf("CopyDir function returned error: %v\n", err)
		}
		if _, err = os.Stat(expect); err != nil {
			t.Errorf("Error copying directory: %v\n", err)
		}

		os.RemoveAll(src)
		os.RemoveAll(dst)
	}
}
