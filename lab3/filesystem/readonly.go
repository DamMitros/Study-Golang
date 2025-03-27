package filesystem

type ReadOnlyFile struct {
	File
}

func NewReadOnlyFile(name string, path string) *ReadOnlyFile {
	return &ReadOnlyFile{*NewFile(name, path)}
}

func (f *ReadOnlyFile) Write(p []byte) (n int, err error) {
	return 0, ErrPermissionDenied
}
