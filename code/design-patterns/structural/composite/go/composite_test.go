package composite

import "testing"

func TestCompositePattern(t *testing.T) {
	file1 := &File{name: "File1.txt", size: 100}
	file2 := &File{name: "File2.txt", size: 150}

	dir1 := &Folder{name: "Dir1"}
	dir2 := &Folder{name: "Dir2"}

	// Adding files to directories
	dir1.add(file1)
	dir1.add(file2)

	// Adding directory to another directory
	dir2.add(dir1)

	// Asserting sizes
	if file1.calculateSize() != 100 {
		t.Errorf("Expected size of file1: 100, but got %d", file1.calculateSize())
	}

	if file2.calculateSize() != 150 {
		t.Errorf("Expected size of file2: 150, but got %d", file2.calculateSize())
	}

	if dir1.calculateSize() != 250 {
		t.Errorf("Expected size of dir1: 250, but got %d", dir1.calculateSize())
	}

	if dir2.calculateSize() != 250 {
		t.Errorf("Expected size of dir2: 250, but got %d", dir2.calculateSize())
	}
}
