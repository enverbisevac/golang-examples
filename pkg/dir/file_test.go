package dir

import (
	"os"
	"testing"
	"testing/fstest"
)

func TestFilesInMemory(t *testing.T) {

	t.Parallel()

	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}

	files := Files(fsys, "")
	want := 4
	got := len(files)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFilesOnDisk(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("./test_data")
	want := 2
	got := len(Files(fsys, ""))
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("./test_data")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Files(fsys, "")
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Files(fsys, "")
	}
}

