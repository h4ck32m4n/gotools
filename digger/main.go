package digger

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

type File struct {
	Name string
}

type Folder struct {
	Name    string
	Files   []*File
	Folders map[string]*Folder
}

func Dig(dir string) *Folder {
	dir = path.Clean(dir)
	var tree *Folder
	var nodes = map[string]interface{}{}
	var walkFun filepath.WalkFunc = func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			nodes[p] = &Folder{path.Base(p), []*File{}, map[string]*Folder{}}
		} else {
			nodes[p] = &File{path.Base(p)}
		}
		return nil
	}
	err := filepath.Walk(dir, walkFun)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range nodes {
		var parentFolder *Folder
		if key == dir {
			tree = value.(*Folder)
			continue
		} else {
			parentFolder = nodes[path.Dir(key)].(*Folder)
		}

		switch v := value.(type) {
		case *File:
			parentFolder.Files = append(parentFolder.Files, v)
		case *Folder:
			parentFolder.Folders[v.Name] = v
		}
	}

	return tree
}

func (f *Folder) Build(path string, index int) {

	newPath := path
	if index == 0 {
		newPath += "/" + f.Name
	}
	os.Mkdir(newPath, 0770)

	for _, file := range f.Files {
		newFile := newPath + "/" + file.Name
		Touch(newFile)
	}

	for folder := range f.Folders {
		newFolder := newPath + "/" + folder
		f.Folders[folder].Build(newFolder, index+1)
	}
}

func (f *Folder) String() string {
	j, _ := json.Marshal(f)
	return string(j)
}

func (f *Folder) NodeTree(index int, output string) string {
	if index == 0 {
		output += f.Name + "\n"
	}
	for _, file := range f.Files {
		for i := 0; i < index; i++ {
			output += "\t"
		}
		output += "└── " + file.Name + "\n"
	}
	for folder := range f.Folders {
		for i := 0; i < index; i++ {
			output += "\t"
		}
		output += "└── " + f.Folders[folder].Name + "\n"
		output = f.Folders[folder].NodeTree(index+1, output)
	}
	return output
}

func (f *Folder) Tree() string {
	return f.NodeTree(0, "")
}
