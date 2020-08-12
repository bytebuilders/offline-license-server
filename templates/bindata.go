// Package templates Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// index.html
package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xcd\x72\xdc\x36\x0c\xbe\xf7\x29\x10\x5e\x72\xa9\x96\xc9\xad\xcd\x50\x9a\xb6\xb6\x0f\x99\x76\xda\x4c\x9d\x4b\x8f\x5c\x12\xbb\x62\x4c\x91\x1c\x02\xda\x78\xa7\xd3\x77\xef\x50\xd2\xae\xd7\x2b\xa5\xfe\x89\x33\x53\x1f\x6c\x41\xf8\xf9\x88\x8f\x00\x04\xab\x57\x97\x7f\x5c\x7c\xfc\xeb\xc3\x15\xb4\xdc\xf9\xe6\x3b\x35\xfe\x01\x50\x2d\x6a\x5b\x1e\x00\x54\x87\xac\xc1\xb4\x3a\x13\x72\x2d\x7a\xde\x54\x3f\x08\x90\xa7\xca\xa0\x3b\xac\xc5\xce\xe1\xe7\x14\x33\x0b\x30\x31\x30\x06\xae\xc5\x67\x67\xb9\xad\x2d\xee\x9c\xc1\x6a\x10\xbe\x07\x17\x1c\x3b\xed\x2b\x32\xda\x63\xfd\xf6\x2e\x14\x3b\xf6\xd8\xfc\x9c\x12\x5d\x44\x8b\xf0\x9b\x33\x18\x08\xe1\x1a\xf3\x0e\xb3\x92\xa3\x7a\x34\xf5\x2e\xdc\x0c\x4f\x00\x19\x7d\x2d\x88\xf7\x1e\xa9\x45\x64\x31\xbd\x6e\x33\x6e\x6a\xd1\x32\x27\x7a\x27\xa5\xb1\x61\xf5\x89\x2c\x7a\xb7\xcb\xab\x80\x2c\x43\xea\xe4\xba\xf7\x9d\xfe\xe9\xcd\xea\xc7\xd5\x1b\x69\x88\x46\x79\xd5\xb9\xb0\x32\x44\x63\x9c\xe1\x68\x4a\x1e\xc8\x50\xeb\x68\xf7\xd3\x11\x08\x0d\xbb\x18\xc0\x78\x4d\x54\x8b\x49\x14\xcd\x84\xaf\xac\xdb\x1d\x74\x85\x0e\xed\x02\xe6\xa3\xb6\x10\xfc\xf6\xa0\x1e\x12\x3b\x51\x01\x7c\x81\x83\x3b\x67\xd9\xbe\x3d\x09\xb5\x89\xb9\x03\x3d\xe0\xd7\x42\x3a\xa2\x1e\x2b\x3f\x7a\x0a\xe8\x90\xdb\x68\x6b\x91\x22\xf1\x3d\x90\xd3\x13\x6e\x1c\x7a\x7b\x4f\x5b\x48\xd6\x6b\xf4\x07\x8b\x41\x10\xcd\xef\xba\x43\x25\x07\xe1\xcc\xfa\x2c\xdf\x1c\xfd\x59\x3c\x00\xe5\x42\xea\xf9\xec\x25\x4c\xd5\x53\x7e\x8b\x99\x6e\x8a\x38\x38\xce\xb5\xbc\x4f\x58\x0b\xc6\xdb\x05\x5d\xf2\xda\x60\x1b\xbd\xc5\x5c\x8b\x73\xb5\x3c\x3b\xbc\xb4\x6e\x77\x8f\x9b\xd9\x8b\xe7\x90\x75\xd5\x69\xe7\x1f\xcb\x16\xb4\x9a\x2a\x67\x62\xa0\xca\xe3\x86\x4f\xc4\xec\xb6\x2d\x3f\x8d\x4c\x2c\xc8\xcf\x63\xf3\x0b\xae\xff\x49\x27\xc0\x4e\xfb\x1e\x1f\xe4\xb9\xb4\x4d\xd2\xc7\x9e\x29\xe9\x81\xa3\x8a\x3a\xed\x7d\x79\x28\x89\xcf\x12\x2d\xa9\x1e\xa9\xd7\x04\x1b\x5d\x61\xd8\xa1\x8f\x09\x45\xa3\xa4\x9b\x41\xc8\x82\xf1\x34\xe0\x65\x8a\x97\x90\x6f\x8d\xd7\x9d\x2e\xad\x56\x71\x76\x3a\x6c\xfd\xe3\x4f\x31\xab\x2a\x00\xf5\xaa\xaa\x40\xa5\x03\x4a\x8b\x3e\x95\xf3\x58\x1d\xb6\x65\x5c\x7c\x6c\x1d\xc1\x70\x23\xe0\x08\x5c\xd8\x69\xef\xac\x92\xa9\x81\xaa\x7a\xf9\x7a\xfd\x90\xa3\xed\x0d\x3f\xbf\xbf\x4f\x6c\x08\x3d\x9a\x45\x4a\x47\x0d\xb8\x32\x93\x46\x40\x31\x55\xed\x41\x9c\x3b\x01\xa8\x98\x86\x71\x3b\x15\xda\x4d\xbf\x46\xbb\xae\x4c\xec\xba\x3e\x38\xde\xcf\x2b\xb2\xfc\x34\xbf\xf6\x6b\xbc\xfc\x05\x2e\x0e\x66\x70\x65\x5d\x09\xa3\xe4\x18\x6e\xc1\xeb\xf1\xe0\x18\x18\x73\xca\x8e\x16\xe6\xd6\x29\xfa\xd5\xd1\xee\x65\xe0\x89\x35\xb5\x0f\xa6\x7e\x5d\xac\x5e\x3a\xf3\x11\xfa\xc1\xc4\x47\xec\xaf\xcc\x5b\xc9\xb1\x52\xe6\x9d\x35\xef\xa2\x6f\x33\xbf\xcb\xfd\xe5\x80\x8c\x04\x17\xbe\x27\xc6\x0c\xef\x2f\xbf\xcd\xd7\xcf\x8c\xf1\xab\xde\xd9\xff\xcd\x47\x10\x16\x06\x93\x0b\x9b\x28\x9a\x3f\xfb\x00\xa5\xfe\x74\xb0\xef\x40\x99\x68\xb1\x29\x3d\x61\xd8\xc3\x16\x19\x02\x41\x11\x2b\xda\x13\x63\x07\x55\xac\x3f\x51\x0c\x49\x73\x5b\xbf\xfe\x7b\x55\x36\x46\xab\x59\xaf\x7a\x67\xff\x79\xad\xe4\xe0\x5e\x66\xda\x57\xdd\xdf\x63\x2e\xe0\xde\x1d\x9b\x16\xcd\xcd\x3a\xde\x2e\x4e\xfd\xc2\xf4\x74\x33\x1c\x49\x4c\x1c\x1f\x5d\x0e\xed\xc0\xb9\x47\x31\xff\xc8\x01\xbc\x07\xbd\xcd\x88\xc0\x11\xb8\xc5\x39\x80\x3e\xdb\x4f\x75\x4a\x54\x78\x58\x99\xd8\x49\x8f\x5b\xed\x25\x47\x92\x4b\xed\xd5\x30\xe6\x8e\x40\x07\x5b\xd6\xec\xb1\xaf\x48\x49\x3d\x33\x9d\xf7\xcd\x52\xe1\x3e\xbd\x73\x4a\x19\x6c\x73\xec\x13\x3e\xeb\x12\xd6\x3d\xf3\xdd\xda\x3c\x49\xe5\xdb\xef\xc2\x8d\x68\xae\xfb\x75\xe7\x58\xc9\xf1\xfd\xd3\xce\xaa\x64\x59\x84\x8f\x0b\xf8\x9d\xb2\x4c\x92\x61\x3b\x1e\xb7\xf9\x71\x89\x57\x72\xfc\x5f\xe7\xdf\x00\x00\x00\xff\xff\x87\x5e\xe1\x70\x03\x0d\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 3331, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.html": indexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": {indexHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}