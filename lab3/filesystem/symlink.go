package filesystem

import (
	"time"
)

type SymLink struct {
	name string
	path string
	createdAt time.Time
	ref FileSystemItem
}

func NewSymLink(name string, path string, ref FileSystemItem) *SymLink {
	return &SymLink{
		name: name,
		path: path,
		createdAt: time.Now(),
		ref: ref,
	}
}

func (s *SymLink) Name() string {return s.name}
func (s *SymLink) Path() string {return s.path}
func (s *SymLink) Size() int64 {return 0}
func (s *SymLink) CreatedAt() time.Time {return s.createdAt}
func (s *SymLink) ModifiedAt() time.Time {return s.ref.ModifiedAt()}