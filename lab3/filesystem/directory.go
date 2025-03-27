package filesystem

import (
	"time"
)

type Directory struct {
	name string
	path string
	items []FileSystemItem
	createdAt time.Time
	modifiedAt time.Time
}

func NewDirectory(name string, path string) *Directory {
	return &Directory{
		name: name,
		path: path,
		items: []FileSystemItem{},
		createdAt: time.Now(),
		modifiedAt: time.Now(),
	}
}

func (d *Directory) Name() string {return d.name}
func (d *Directory) Path() string {return d.path}
func (d *Directory) CreatedAt() time.Time {return d.createdAt}
func (d *Directory) ModifiedAt() time.Time {return d.modifiedAt}

func (d *Directory) Size() int64 {
	var size int64
	for _, item := range d.items {
		size += item.Size()
	}
	return size
}

func (d *Directory) AddItem(item FileSystemItem) error {
	for _, i := range d.items {
		if i.Name() == item.Name() {
			return ErrItemExists
		}
	}
	d.items = append(d.items, item)
	d.modifiedAt = time.Now()
	return nil
}

func (d *Directory) RemoveItem(name string) error {
	for i, item := range d.items {
		if item.Name() == name {
			d.items = append(d.items[:i], d.items[i+1:]...)
			d.modifiedAt = time.Now()
			return nil
		}
	}
	return ErrItemNotFound
}

func (d *Directory) Items() []FileSystemItem {
	return d.items
}