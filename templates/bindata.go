// Package templates Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// index.html
// pricing.html
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xcd\x72\x1b\x37\x0c\xbe\xf7\x29\x10\x5e\x72\xe9\x8a\xf1\xad\xcd\x70\x35\x6d\x6d\x1f\x32\xed\xb4\x99\x3a\x93\x4e\x8f\x14\x09\x6b\x19\x73\xc9\x1d\x02\xbb\x89\xa6\xd3\x77\xef\x70\x7f\x64\x4b\xab\x54\x8a\x2d\xf9\x60\x09\x04\x08\xe0\xfb\x40\x88\x84\x7a\x75\xf3\xc7\xf5\x87\xbf\xdf\xdf\x42\xc5\xb5\x5f\x7e\xa7\x86\x0f\x00\x55\xa1\xb6\xf9\x0b\x80\xaa\x91\x35\x98\x4a\x27\x42\x2e\x45\xcb\xf7\xc5\x0f\x02\xe4\x53\x65\xd0\x35\x96\xa2\x73\xf8\xb9\x89\x89\x05\x98\x18\x18\x03\x97\xe2\xb3\xb3\x5c\x95\x16\x3b\x67\xb0\xe8\x85\xef\xc1\x05\xc7\x4e\xfb\x82\x8c\xf6\x58\x5e\x3d\xba\x62\xc7\x1e\x97\x3f\x37\x0d\x5d\x47\x8b\xf0\x9b\x33\x18\x08\xe1\x0e\x53\x87\x49\xc9\x41\x3d\x98\x7a\x17\x1e\xfa\x6f\x00\x09\x7d\x29\x88\x37\x1e\xa9\x42\x64\x31\x2e\x57\x09\xef\x4b\x51\x31\x37\xf4\x56\x4a\x63\xc3\xe2\x13\x59\xf4\xae\x4b\x8b\x80\x2c\x43\x53\xcb\x55\xeb\x6b\xfd\xd3\x9b\xc5\x8f\x8b\x37\xd2\x10\x0d\xf2\xa2\x76\x61\x61\x88\x06\x3f\x7d\x6a\x4a\x4e\x64\xa8\x55\xb4\x9b\x31\x05\x42\xc3\x2e\x06\x30\x5e\x13\x95\x62\x14\xc5\x72\x8c\xaf\xac\xeb\x26\x5d\xa6\x43\xbb\x80\x69\xab\xcd\x04\x5f\x4d\xea\x1e\x98\xf8\x3a\xf0\xea\xea\xc9\xb6\xfb\x98\x6a\xd0\x7d\xac\x52\x48\x47\xd4\x62\xe1\x87\x0d\x02\x6a\xe4\x2a\xda\x52\x34\x91\xf8\x49\xac\xdd\x6c\xee\x1d\x7a\xbb\xa3\xcd\x84\xea\x15\xfa\xc9\xa2\x17\xc4\xf2\x77\x5d\xa3\x92\xbd\xb0\x67\xbd\x87\x2d\x45\xbf\xe7\x0f\x40\xb9\xd0\xb4\x3c\x9e\x8c\xfc\x5f\x4c\x3b\x7a\x85\x00\xde\x34\x58\x0a\xc6\x2f\x2c\xa0\xf1\xda\x60\x15\xbd\xc5\x54\x8a\xed\x81\xd8\xba\x92\xd6\x75\x3b\x68\x66\x0b\xcf\x81\xf7\x57\x4c\x0f\x70\x5b\x6b\xe7\x5f\x08\x72\x6f\x11\x46\xd0\x98\x5d\x8b\x99\x72\x87\x85\x99\x76\x60\xe5\x2b\x5b\x77\x69\x9a\xa9\x3b\xed\x5b\x9c\x2b\x8e\xb1\x09\xa0\x5e\x15\x05\xa8\x66\x4a\xad\x42\xdf\x80\xa3\xc2\xea\xb0\xce\x67\xf6\x43\xe5\x08\xfa\x94\xc0\x11\xb8\xd0\x69\xef\xac\x92\xcd\x12\x8a\xe2\xfc\x65\x79\x9f\xa2\x6d\x0d\x3f\xbf\x26\x4f\x6c\x08\x3d\x1a\x9e\x99\xf4\xdd\x9b\x35\xe0\x72\xb3\x0c\x01\xc5\x58\xb6\x49\x9c\x6f\x02\x50\xb1\xe9\x7b\x7e\x64\x9a\x58\x53\x55\x60\x60\x4c\x4d\x72\x84\x07\xf7\x00\xdc\x65\x33\xb8\xdd\x9a\xc1\xad\x75\xd9\xcd\xa1\x00\x72\x88\x70\x72\x70\x13\xeb\xba\x0d\x8e\x37\xff\x1b\xfb\x7a\xb2\x3a\x4f\xe8\x87\x76\x85\x76\x75\x1c\xf8\xaf\xed\x0a\x6f\x7e\x39\x3b\xf2\x31\xfc\x31\xe8\x63\xf4\x0b\x60\xef\x74\xeb\xf9\x34\xf8\x1f\xb3\xe9\x45\x18\x18\x92\x38\x85\x84\x21\x87\x0b\xf0\x90\xef\xa2\xd3\x68\xe8\x6f\xad\x4b\xb0\xd0\xa7\x70\x0a\x09\x7d\x06\x67\xe6\xa0\x8b\x1b\xbd\xc6\x74\x9c\x82\x8f\x83\xe1\xd9\x19\x98\x12\x38\x46\xc0\x14\xff\xcc\xf8\x4d\x0c\x14\x3d\x1e\xc7\x7f\x3d\x18\x9e\x1d\xff\x94\xc0\x31\xfc\x53\xfc\x97\xe1\x57\x72\xb8\x35\x66\x37\xce\x81\x1b\xf5\x32\x4f\x96\x7c\x90\x53\x40\x46\x82\x6b\xdf\x12\x63\x82\x77\x37\xe7\x79\xa2\x99\xc1\xdf\x19\x5f\x69\x70\xe0\x49\xe1\xc2\x7d\x9c\xe5\xf1\x67\x1b\x20\x57\x50\x07\xfb\x76\x3f\x45\x13\x2d\xce\x0a\xb1\xcc\x9d\x6f\xd8\xc3\x1a\x19\x02\x41\x16\x0b\xda\x10\x63\x0d\x45\x2c\x3f\x51\x0c\x8d\xe6\xaa\x7c\xfd\xcf\x22\x8f\x27\x56\xb3\x5e\xb4\xce\xfe\xfb\x5a\xc9\x03\xfe\xf6\x61\x34\x2f\xab\xdb\x29\xc4\xef\xd4\xd6\x54\x68\x1e\x56\xf1\xcb\xa1\x67\xca\xd3\x0a\x71\xa4\xa9\x1e\xdb\x2d\x53\x23\x70\x6a\x71\x56\x91\xfc\xf7\x0e\xf4\x3a\x21\x02\x47\xe0\x6a\xce\xa4\xd2\x7b\x83\x92\x6e\x1a\xca\x1c\x2d\x4c\xac\xa5\xc7\xb5\xf6\x92\x23\xc9\xf9\x63\x13\x60\xc9\x98\x6a\x02\x1d\x6c\x9e\xf7\x86\x7e\x22\x25\xf5\xbc\x5c\xb3\x7e\x39\x74\x60\xbf\xbd\x63\xf2\x81\x5a\xa7\xd8\x36\xf8\xac\x22\xac\x5a\xe6\xc7\xf9\x6d\x94\x1c\x15\x79\xb6\x14\xcb\xbb\x76\x55\x3b\x56\x72\x58\xff\xb6\x5c\x95\xcc\xb7\xcd\x76\x12\x7c\x54\xe6\x5f\x10\x33\xfd\xb8\x28\x39\x4c\x93\x4a\x0e\x43\xf7\x7f\x01\x00\x00\xff\xff\x4a\xbb\x22\xe0\x8c\x0f\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 3980, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pricingHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4d\x6f\x23\x37\x0f\xbe\xbf\xbf\x82\xab\x4b\xf6\x05\x3a\xa3\xcd\xa9\x1f\xd0\xb8\x4d\xb3\x41\xd0\x0f\xb4\x6e\xbb\x45\x91\xd3\x42\x96\x18\x8f\x1a\x8d\xa4\x4a\x1a\xef\xba\x41\xfe\x7b\xa1\xd1\x8c\x63\x7b\x1c\xe7\xc3\x49\xb1\xe8\x25\xb1\x44\x8a\x7c\xf8\x88\x26\x45\xb3\x57\x6f\x7f\x3e\x7d\x77\x31\x3d\x83\x3a\x36\x7a\xf2\x3f\x96\xff\x01\xb0\x1a\xb9\x4c\x1f\x00\x58\x83\x91\x83\xa8\xb9\x0f\x18\x2b\xd2\xc6\xcb\xe2\x0b\x02\x74\x5d\x68\x78\x83\x15\x59\x28\xfc\xe0\xac\x8f\x04\x84\x35\x11\x4d\xac\xc8\x07\x25\x63\x5d\x49\x5c\x28\x81\x45\xb7\xf8\x0c\x94\x51\x51\x71\x5d\x04\xc1\x35\x56\xc7\xb7\xa6\xa2\x8a\x1a\x27\x27\xce\x85\x53\x2b\x11\xa6\xde\xca\x56\x44\x98\x7a\x25\x94\x99\xc3\x8f\x2a\x44\x46\xb3\x52\x3e\xa0\x95\xb9\xea\x3e\x01\x78\xd4\x15\x09\x71\xa9\x31\xd4\x88\x91\xf4\xdb\xb5\xc7\xcb\x8a\xd4\x31\xba\xf0\x15\xa5\x42\x9a\xf2\xcf\x20\x51\xab\x85\x2f\x0d\x46\x6a\x5c\x43\x67\xad\x6e\xf8\x37\x6f\xca\x2f\xcb\x37\x54\x84\x90\xd7\x65\xa3\x4c\x29\x42\xc8\x76\x3a\x80\x8c\x0e\x94\xb0\x99\x95\xcb\x1e\x42\x40\x11\x95\x35\x20\x34\x0f\xa1\x22\xfd\x92\x4c\x7a\xff\x4c\xaa\xc5\x20\x4b\xa4\x70\x65\xd0\xaf\xa4\x89\xe6\xe3\x41\xdc\x05\x46\xee\x0b\xbf\x3e\x5e\x3b\x7c\x69\x7d\x03\xbc\xf3\x58\x11\x02\x0d\xc6\xda\xca\x8a\x38\x1b\xe2\x9a\x8f\x4d\x14\x97\x0a\xb5\xdc\x90\x26\x22\xf9\x0c\xf5\xa0\xd1\x2d\xc8\xe4\x27\xde\x20\xa3\xdd\x62\x4b\x7b\x2b\x26\x6f\xf5\x96\x3d\x00\xa6\x8c\x6b\xe3\xd6\x26\xf4\x79\x92\xfe\x92\x91\xac\xb7\xd8\x1d\x1c\x4b\xe3\xd2\x61\x45\x22\x7e\xdc\x21\x73\x9a\x0b\xac\xad\x96\xe8\x2b\xb2\x2d\xa6\x5b\xe0\xa9\x54\x8b\x0d\x6e\xf2\xc6\xa1\x6c\xfd\x61\xfd\x15\x9c\x35\x5c\xe9\x97\xe1\x0c\x93\xe9\xa7\x91\x76\xc7\xd1\xbd\xac\x01\x2c\xb8\x6e\xf1\xf1\x74\x02\xb0\x57\x45\x01\xcc\x0d\xd0\x6a\xd4\x0e\x54\x28\x24\x37\xf3\x94\xfa\xef\x6a\x15\xa0\x83\x04\x2a\x80\x32\x0b\xae\x95\x64\xd4\x4d\xa0\x28\x5e\xe0\x5e\x4e\x4f\x0f\xbc\x8f\x9e\x7f\x21\xc8\x26\xd9\xeb\x19\xb9\xc5\xe5\xbf\x95\x73\xdf\xdb\x19\x74\x55\xe3\x79\x42\xcc\x05\xe8\x93\x8b\x72\x5a\x5b\xf3\x5c\x11\xa2\x46\x97\xcc\x7d\x7a\x51\x9e\xda\xc6\x71\xb3\xec\x80\xbe\x4c\x05\x11\xd9\xc3\x7f\xad\xf0\xf6\x6d\xf2\xe9\x9c\xad\xe9\x04\xd4\x28\x62\xaa\x56\x4d\xab\xa3\x72\xa9\x1f\x8f\x82\x66\x83\x56\xea\xb4\xd9\x39\xe9\x29\x5e\x2d\x87\xe3\x10\xd4\xdf\x58\x91\xcf\x77\x98\x01\x60\xd6\x75\x4f\x87\xbe\xd2\x5e\xb5\x33\x94\xb3\x02\x4d\x44\xef\xbc\x0a\x48\xe0\xfa\x1a\xd4\x25\xe0\x5f\x50\x0e\x8f\x81\x5d\x5a\x37\x37\x19\x11\xca\xeb\x6b\x40\x23\xe1\xe6\x66\xf2\x43\x3b\xc3\xb7\xdf\xc2\xd9\x4a\x0d\x5e\x9f\x18\xd3\x72\xfd\x7f\x46\xb3\xdb\x87\x23\x72\x7c\x39\xdf\x87\x25\xcb\xf7\xa0\x98\xf2\x65\x71\x12\x8a\x0b\xdb\x16\xe7\x16\x5e\x4f\x4f\x2e\xce\x9f\x00\xc3\x63\x40\xad\xd1\xef\x83\x72\xab\xb3\x07\xce\xaf\xbd\xd2\xe3\x21\xb4\x46\xab\x46\x45\x94\xfb\x30\xac\x29\xed\x01\xf1\xfb\xa0\xf5\x94\x8b\x09\x91\x87\xfa\xbe\x4c\x19\x2b\xed\x82\xf3\x5b\xd2\x3a\x30\x4f\xb2\xa7\x3b\xd3\x64\x5d\x7c\x37\x84\x03\x93\x24\x3b\xd9\x7f\x41\x23\x9d\xbb\xd1\x3c\xe6\x7a\x18\xcd\x56\x46\x95\x65\xc7\xdb\xe8\x19\x8a\xe0\x43\x6a\xda\x46\xa1\x14\x35\x8a\xab\x99\xfd\xb8\xab\x96\x6d\x74\x47\x1b\x86\x4e\xb8\x3a\x32\xd0\x1b\x7d\x8b\x84\x8e\x0d\x7c\x07\x7c\xee\x11\x21\x5a\x88\x35\x8e\xed\xf3\xad\x09\x8c\x3b\x17\x84\x95\x58\x0a\xdb\x50\x8d\x73\xae\x69\xb4\x81\x8e\x5b\x0b\xc0\x24\xa2\x6f\x02\x70\x23\xd3\x38\x29\x55\x22\x3f\x30\xca\x47\xaa\x63\xde\x77\xb5\x82\x3b\x98\xdf\x43\x7c\xea\x03\x73\x6f\x5b\x87\x4f\xba\x83\x59\x1b\xe3\xed\x60\xd8\xaf\x54\x28\xd2\xd0\xba\xe2\x75\x8e\x06\x3d\x8f\x48\xc0\x1a\xa1\x95\xb8\xaa\x8e\x62\xad\x42\x99\x26\xbb\x72\x98\xec\xe8\x7b\xea\xf2\x10\x48\x8e\x26\xe7\xfd\x09\xf8\xa5\xb5\x91\x27\x05\x46\xb3\xf1\xfb\x5f\xe6\x87\xc0\xee\xdf\xef\x03\xf0\x3c\x54\x3c\x0c\xf5\xd7\x01\x8d\x7c\xdf\x9d\xa8\xba\x44\x3a\x9a\x74\x93\xd2\xa3\x43\xd8\xda\x60\x34\x39\x5c\x8d\xda\xb7\xc2\xf4\x8d\x14\xc3\x97\x95\xd1\x3c\xae\x33\x9a\x7f\xdb\xf8\x27\x00\x00\xff\xff\x68\xe8\xa3\x2d\xf3\x10\x00\x00")

func pricingHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pricingHtml,
		"pricing.html",
	)
}

func pricingHtml() (*asset, error) {
	bytes, err := pricingHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pricing.html", size: 4339, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"index.html":   indexHtml,
	"pricing.html": pricingHtml,
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
	"index.html":   {indexHtml, map[string]*bintree{}},
	"pricing.html": {pricingHtml, map[string]*bintree{}},
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
