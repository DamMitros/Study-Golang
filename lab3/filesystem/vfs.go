package filesystem

import (
	"path/filepath"
	"strings"
)

type VirtualFileSystem struct {
	root DirectorySystem
}

func NewVirtualFileSystem() *VirtualFileSystem {
	return &VirtualFileSystem{NewDirectory("/", "/")}
}

func (vfs *VirtualFileSystem) FindItem(path string) (FileSystemItem, error) {
	if path == "/" {return vfs.root, nil}

	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := vfs.root

	for i, part := range parts {
		found := false
		for _, item := range current.Items() {
			if item.Name() == part {
				if i == len(parts)-1 { return item, nil}

				if dir, ok := item.(DirectorySystem); ok {
					current = dir
					found = true
					break
				} else {
					return nil, ErrNotDirectory
				}
			}
		}
		if !found { return nil, ErrItemNotFound}
	}
	return nil, ErrItemNotFound
}

func (vfs *VirtualFileSystem) AddItem(item FileSystemItem) error {
	itemPath := filepath.ToSlash(filepath.Dir(item.Path()))
	if itemPath == "." { return ErrPermissionDenied}
	if itemPath == "/" { return vfs.root.AddItem(item)}

	parent, err := vfs.FindItem(itemPath)
	if err != nil {return err}

	if dir, ok := parent.(DirectorySystem); ok {
		return dir.AddItem(item)
	} else {
		return ErrNotDirectory
	}
}

func (vfs *VirtualFileSystem) RemoveItem(name string) error {
	item, err := vfs.FindItem(name)
	if err != nil { return err}

	parentPath := filepath.ToSlash(filepath.Dir(item.Path()))
	if parentPath == "." { return ErrPermissionDenied}
	if parentPath == "/" { return vfs.root.RemoveItem(name)}

	parent, err := vfs.FindItem(parentPath)
	if err != nil {return err}

	if dir, ok := parent.(DirectorySystem); ok {
		return dir.RemoveItem(filepath.Base(name))
	} else {
		return ErrNotDirectory
	}
}

func (vfs *VirtualFileSystem) Items() []FileSystemItem {
	return vfs.root.Items()
}