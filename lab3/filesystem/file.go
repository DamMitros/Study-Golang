package filesystem

import (
	"time"
)

type File struct {
	name string
	path string
	content []byte
	createdAt time.Time
	modifiedAt time.Time
}

func NewFile(name string, path string) *File {
	return &File{
		name: name,
		path: path,
		content: []byte{},
		createdAt: time.Now(),
		modifiedAt: time.Now(),
	}
}

func (f *File) Name() string {return f.name}
func (f *File) Path() string {return f.path}
func (f *File) Size() int64 {return int64(len(f.content))}
func (f *File) CreatedAt() time.Time {return f.createdAt}
func (f *File) ModifiedAt() time.Time {return f.modifiedAt}

func (f *File) Read(p []byte) (n int, err error) {
	return copy(p, f.content), nil
}

func (f *File) Write(p []byte) (n int, err error) {
	f.content = append(f.content, p...)
	f.modifiedAt = time.Now()
	return len(p), nil
}