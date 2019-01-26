package util_test

import (
	"os"
	"testing"
	"github.com/BobbyShaftoe/go-shell-replace/util"
)

func TestArchiveDir(t *testing.T) {
	// Test archiving into destination directory
	var err error
	src, dst, dirs := "src", "dst", "test/dir"
	os.MkdirAll(src+"/"+dirs, 0755)
	os.MkdirAll(dst, 0755)

	a := &util.TarGzArgs{Name: "test.tar.gz", Src: src, DstDir: dst}
	if err = a.CreateArchive(); err != nil {
		t.Errorf("Error creating archive %v\n", err)
	}

	os.RemoveAll(src)
	os.RemoveAll(dst)
}
