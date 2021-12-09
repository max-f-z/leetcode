package main

import (
	"sort"
	"strings"
)

type fd struct {
	isFile  bool
	prefix  string
	name    string
	content string
	fds     []*fd
}

func (f *fd) getChild(name string) *fd {
	for _, fileDescriptor := range f.fds {
		if name == fileDescriptor.name {
			return fileDescriptor
		}
	}
	return nil
}

type FileSystem struct {
	root *fd
}

func Constructor588() FileSystem {
	fs := FileSystem{}
	fs.root = &fd{
		isFile: false,
		prefix: "/",
		fds:    []*fd{},
	}

	return fs
}

func (this *FileSystem) Ls(path string) []string {
	current := this.root
	paths := strings.Split(path, "/")

	for i := 1; i < len(paths); i++ {
		child := current.getChild(paths[i])
		if child == nil {
			break
		}
		current = child
	}

	ans := []string{}
	if current.isFile {
		ans = append(ans, current.name)
	} else {
		for _, fileDescriptor := range current.fds {
			ans = append(ans, fileDescriptor.name)
		}
		sort.Strings(ans)
	}

	return ans
}

func (this *FileSystem) Mkdir(path string) {
	paths := strings.Split(path, "/")
	current := this.root

	for i := 1; i < len(paths); i++ {
		child := current.getChild(paths[i])
		if child == nil {
			child = &fd{
				isFile: false,
				prefix: current.prefix + current.name + "/",
				name:   paths[i],
				fds:    []*fd{},
			}
			if current == this.root {
				child.prefix = "/"
			}
			current.fds = append(current.fds, child)
		}
		current = child
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	paths := strings.Split(filePath, "/")
	current := this.root

	for i := 1; i < len(paths); i++ {
		child := current.getChild(paths[i])
		if child == nil {
			child = &fd{
				isFile: false,
				prefix: current.prefix + paths[i],
				name:   paths[i],
				fds:    []*fd{},
			}
			current.fds = append(current.fds, child)
		}
		current = child
	}

	current.isFile = true
	current.content = current.content + content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	paths := strings.Split(filePath, "/")
	current := this.root

	for i := 1; i < len(paths); i++ {
		current = current.getChild(paths[i])
	}

	return current.content
}
