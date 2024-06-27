package composite

type Component interface {
	calculateSize() int
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) calculateSize() int {
	size := 0
	for _, component := range f.components {
		size += component.calculateSize()
	}
	return size
}

func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}

type File struct {
	size int
	name string
}

func (f *File) calculateSize() int {
	return f.size
}
