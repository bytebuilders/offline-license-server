// Package templates Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// eula.html
// index.html
// offerletter.html
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

var _eulaHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5f\x6f\x23\x35\x10\x7f\xe7\x53\xcc\x99\x57\x36\xbe\x3f\x2f\x80\xbc\x11\xa1\x9c\xfa\x52\x41\xd1\x15\xa1\x7b\x42\x13\x7b\x92\x35\xf8\x1f\xb6\x37\xbd\x7c\x7b\xe4\xf5\x2e\x4d\xd2\xed\xa9\x40\xae\x27\xa4\xee\x4b\x3c\x9e\x9f\x67\x66\xe7\x37\xf6\x7a\x22\x5e\xfc\xf0\xd3\xc5\xcd\xfb\xeb\xb7\xd0\x65\x6b\x96\x5f\x88\xfa\x03\x20\x3a\x42\x55\x06\x00\xc2\x52\x46\x90\x1d\xc6\x44\xb9\x65\x7d\xde\x34\x5f\x33\xe0\x87\x4a\x87\x96\x5a\xb6\xd3\x74\x1b\x7c\xcc\x0c\xa4\x77\x99\x5c\x6e\xd9\xad\x56\xb9\x6b\x15\xed\xb4\xa4\x66\x10\xbe\x02\xed\x74\xd6\x68\x9a\x24\xd1\x50\xfb\xea\xce\x54\xd6\xd9\xd0\x72\x15\x42\xba\xf0\x8a\xe0\x9d\xdf\xe4\x5b\x8c\x04\x57\x5a\x92\x4b\x04\xab\x6d\x24\xb2\xe4\xb2\xe0\x15\x5a\x97\x19\xed\xfe\x80\x48\xa6\x65\xa9\xf3\x31\xcb\x3e\x83\x96\xde\x31\xc8\xfb\x40\x2d\xd3\x16\xb7\xc4\x83\xdb\x32\xe8\x22\x6d\x5a\xd6\xe5\x1c\xd2\xb7\x9c\x4b\xe5\x16\x18\x42\x92\x5e\xd1\x42\x7a\xcb\x07\x64\xe2\x21\x7a\xd5\xcb\x9c\xf8\xa4\xe4\xc5\x5c\xe2\x1b\xdc\x95\x41\xf3\xe6\xf5\x87\x37\xaf\x17\xc5\xe0\x41\x00\xc3\x08\xc6\x30\xf2\xde\x50\xea\x88\x32\x1b\xa7\xef\x3b\xfe\x3d\x29\x32\x7a\x17\x17\x8e\x32\x77\xc1\xf2\x75\x6f\x2c\x7e\xf7\x72\xf1\xcd\xe2\x25\x97\x29\x55\x79\x61\xb5\x5b\xc8\x94\xaa\x9d\x21\x4f\x82\x4f\xcc\x88\xb5\x57\xfb\x31\x84\x44\x32\x6b\xef\x40\x1a\x4c\xa9\x65\xa3\x38\x06\x08\x20\x94\xde\x4d\xba\xc2\x0d\x6a\x47\xf1\x6f\xed\xa9\xde\xf4\xd6\x25\xd0\xa9\xb1\x7e\xad\x0d\x95\x91\x24\x97\x29\x92\x3a\x58\x33\xb7\xaa\x40\x3b\x34\x9b\x23\x58\x29\xa6\x57\x13\x6e\x20\x8e\xdd\x91\xfc\xf6\x97\xab\x15\x5c\x92\xa3\x88\xd9\x47\xc1\xbb\x57\x27\x4b\x37\x3e\x5a\xc0\xe1\x75\x5a\xc6\x7f\xe3\xd4\x1b\x64\x60\x29\x77\x5e\xb5\x2c\xf8\x94\x4f\x9c\x1d\xc7\xb5\xd1\x64\xd4\x3d\x44\x21\x0d\xd7\x64\x26\xd4\x20\xb0\xe5\x85\xb7\x01\xdd\x1e\xae\x68\x8b\x06\x7e\x44\x4b\x82\x0f\xaa\x99\xf5\x27\x19\x8d\xde\xcc\x78\x01\x10\xda\x85\x3e\xcf\x28\x60\xdc\x34\xb2\xfa\x64\xb3\x90\xd1\xc1\x60\x63\x1e\x51\x6b\x3c\xd3\x87\x07\xf4\xc1\xa0\xa4\xce\x1b\x45\xb1\x65\x73\x10\x3e\xf3\x6e\x5c\xe9\xdd\xbd\xa4\xd6\xc9\x73\xa7\xfa\x57\x5a\x27\x9d\x9f\x20\xcf\xca\x5b\xd4\xee\xb3\xa5\x19\x60\x87\xa6\xa7\xff\xc6\x01\x80\x78\xd1\x34\x20\xc2\x14\x71\x47\x26\x94\x3d\xa7\xd0\x6d\xcb\x86\xbe\xe9\x74\x02\xb2\xa8\x0d\xe8\x04\xda\xed\xd0\x68\x25\x78\x58\x42\xd3\x3c\x11\xa1\x75\xef\xac\x94\x8a\x94\xd2\xa7\xa7\x15\xab\xa3\x67\x5e\x4f\x3d\x9c\x8b\xd7\x9f\x7b\x9f\x71\xf8\xb6\x7c\xf9\xe9\xd9\xfc\x73\x72\xf6\xcc\xe7\xa9\x87\x73\xf1\x79\x5d\x6f\x37\x4f\xf4\x75\x1b\xef\x52\xcf\x74\x9e\x7a\x38\x1b\x9d\xb8\x2f\x37\x62\xb8\xa1\x68\xcf\x70\xdc\x1e\xe0\x12\x19\x92\xf7\x2f\x57\x23\xb0\x6a\x41\x97\x3b\x58\x0d\xa1\xc9\x14\x2d\x9b\x78\x3f\x9c\x9b\x37\x01\x20\x7c\x18\x0e\x96\x91\x22\x74\xae\xc7\xf9\xb8\xea\xb3\x1a\x00\x0f\x19\xe3\xd5\xda\x23\x9d\x05\xdc\x6f\x3f\xe2\xea\x7a\xf5\xfe\xf2\x5f\x39\x12\xbc\x66\x66\x36\xbb\x0f\x14\xd7\x13\x5c\xb7\xde\xf5\xa1\x74\x66\x70\x6d\xd0\x7d\x9e\x2a\x49\x35\x82\x26\x18\x74\x53\x95\x1c\xcd\x3d\x92\xb8\x35\x26\x2d\x3f\xc2\xdc\xf7\x45\x7f\x9e\x1a\xd9\xfa\xd9\x04\x4f\xcf\xa5\x37\xea\x4c\xc5\x68\x30\x6b\xd7\x3f\xbc\x53\xa0\x10\x37\x40\xfe\x0f\x45\x59\x8e\xc6\x6d\xf4\x7d\xa0\xd9\xfa\x7c\x6c\x95\xad\xfb\x9c\xef\x5a\xda\x51\xd2\xa9\x29\xed\x36\x9b\x52\x97\xfa\xb5\xd5\x99\x2d\xc7\x16\x92\xe0\xa2\x98\x44\x99\x05\xaf\x2b\xfe\xd9\xab\x1d\x4f\x95\xde\xf3\xa8\xe1\x3d\x06\x1d\x89\x07\x42\x49\xb7\x9c\xd8\x10\xbc\xb6\xe9\x82\xd7\xbf\x56\xfe\x0a\x00\x00\xff\xff\x41\xb6\x52\x6a\x72\x11\x00\x00")

func eulaHtmlBytes() ([]byte, error) {
	return bindataRead(
		_eulaHtml,
		"eula.html",
	)
}

func eulaHtml() (*asset, error) {
	bytes, err := eulaHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "eula.html", size: 4466, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x4d\x73\xdb\x36\x13\xbe\xbf\xbf\x62\x83\x4b\x2e\x2f\x85\x7c\x5c\xda\x0c\xa9\x69\xea\xf8\x90\x69\xa7\xcd\x34\x99\x74\x7a\x84\x80\x95\x88\x18\x04\x58\x60\xa9\x58\xe3\xf1\x7f\xef\x80\x20\x1d\x7d\x40\x96\x25\x4b\x3e\xd8\x5c\xee\xb3\x8b\x67\xf1\x60\x61\x02\xe5\x8b\x0f\x7f\x5e\x7d\xf9\xe7\xd3\x35\xd4\xd4\x98\xe9\xff\xca\xf4\x07\xa0\xac\x51\xa8\xf8\x00\x50\x36\x48\x02\x64\x2d\x7c\x40\xaa\x58\x47\xf3\xe2\x27\x06\x7c\xdd\x69\x45\x83\x15\x5b\x6a\xfc\xde\x3a\x4f\x0c\xa4\xb3\x84\x96\x2a\xf6\x5d\x2b\xaa\x2b\x85\x4b\x2d\xb1\xe8\x8d\xff\x83\xb6\x9a\xb4\x30\x45\x90\xc2\x60\xf5\xfa\x47\x2a\xd2\x64\x70\xfa\xbe\x6d\xc3\x95\x53\x08\xbf\x6b\x89\x36\x20\x7c\x46\xbf\x44\x5f\xf2\xe4\x4e\x50\xa3\xed\x0d\x78\x34\x15\x0b\xb5\xf3\x24\x3b\x02\x2d\x9d\x65\x40\xab\x16\x2b\xa6\x1b\xb1\x40\xde\xda\x05\x83\xda\xe3\xbc\x62\x35\x51\x1b\xde\x71\x2e\x95\x9d\x88\xb6\x0d\xd2\x29\x9c\x48\xd7\xf0\x1e\x19\x78\xeb\x9d\xea\x24\x05\x3e\x3a\x79\x4c\x17\xf8\x5c\x2c\xe3\x43\xf1\xf6\xcd\xed\xdb\x37\x93\x98\x70\x8d\x40\xff\x04\x03\x0d\x5a\x19\x0c\x35\x22\xb1\xe1\xf5\xee\xc0\xdf\x82\x42\xa3\x97\x7e\x62\x91\xb8\x6d\x1b\x3e\xeb\x4c\x23\x7e\x79\x35\xf9\x79\xf2\x8a\xcb\x10\x92\x3d\x69\xb4\x9d\xc8\x10\x52\x9e\x7e\x6e\x4a\x3e\xaa\x51\xce\x9c\x5a\x0d\x14\x02\x4a\xd2\xce\x82\x34\x22\x84\x8a\x0d\xe6\x40\x10\xa0\x54\x7a\x39\xfa\xa2\x1e\x42\x5b\xf4\x0f\xde\x6d\xbf\xe9\x1a\x1b\x40\x87\xa2\x71\x33\x6d\x30\x3e\x49\xb4\x84\x1e\xd5\x5a\x4c\x2e\x2a\x42\x6b\x61\xe6\x1b\xb0\xb8\x80\x5e\x8f\xb8\x5e\x38\xb6\x5f\xd8\xfa\xf5\x56\xe8\xdc\xf9\x06\x44\x5f\x4e\xc5\xb8\x0e\xa1\xc3\xc2\xa4\x20\x06\x0d\x52\xed\x54\xc5\x5a\x17\x68\x6b\xcc\x4d\x7a\x73\x8d\x46\xed\x20\xa2\x76\x62\x86\x66\x44\xf5\x06\x9b\xfe\x21\x1a\x2c\x79\x6f\x64\x22\xb6\xa6\xd2\x3b\x93\xc9\x0b\x50\x6a\xdb\x76\x34\x74\x43\xfc\xcd\xc6\xa8\xde\x31\xae\x4e\xc2\x5b\x62\xd0\x1a\x21\xb1\x76\x46\xa1\xaf\xd8\x43\x13\x6c\xa4\xe3\x4a\x2f\x77\x2a\xcc\xbe\x3c\xb5\xec\xbf\x9d\xbf\x81\xeb\x46\x68\x73\xa6\xe2\x33\x0e\x18\x26\x04\xe3\x30\x2c\x0b\xd8\x98\xa5\x2c\x22\xcd\xdc\x23\x29\x36\xa7\x33\x0b\x59\x0a\xd3\x61\xde\xf9\xd4\xd9\x07\x28\x5f\x14\x05\x94\xed\x48\xb9\x46\xd3\xc6\x0e\x50\xc2\x2e\x62\x7b\x7d\xa9\x75\x80\x9e\x26\xe8\x00\xda\x2e\x85\xd1\xaa\xe4\xed\x14\x8a\xe2\xb2\x52\x7e\x4a\x5b\xd8\xf3\x75\x5c\xc3\x05\x34\x28\x77\xbb\x6c\x00\x26\x2f\xe8\xd8\x8c\x69\x70\x36\x48\x3d\x9a\xf9\x40\x80\xd2\xb5\xfd\xd6\x35\x28\x72\xd3\xcd\x50\xcd\x8a\x7e\xbf\x69\xbd\x8e\x5d\x7e\x77\x07\x7a\x0e\xf8\x2f\x4c\x86\xba\x20\x87\xba\xbf\x4f\x1c\x50\xdd\xdd\x01\x5a\x05\xf7\xf7\xfb\x86\x04\xf8\xad\x9b\xe1\x87\x5f\xe1\xfa\x21\x1e\xae\x95\x8e\x34\xf6\x91\xe4\x89\xe5\x71\x45\x48\xd7\x34\x9d\xd5\xb4\x7a\xac\x86\x35\xd0\x29\x25\x5c\x8d\xe1\xe7\xad\x20\x90\x08\xf5\x21\x15\x76\x41\x47\x55\xf0\x39\x86\x5f\x4c\x83\x44\xee\x71\x09\x76\x30\x27\xf0\xbf\x90\x00\x71\x75\x2c\x45\x67\xe8\x29\xad\xb0\x0b\x3c\x7a\x29\x7d\x8d\x29\x2e\xda\x10\x89\xe4\xe1\x9e\xd8\xc1\x9d\x58\xcb\x05\x85\x89\x1f\x25\x4f\xd1\x65\x07\x77\x74\x29\xfd\xe7\xcf\x25\x55\xe9\x29\x1e\x16\x65\x1b\x76\x5a\x21\x17\x92\x64\xe9\x56\x62\x81\xfe\x90\x22\x39\xd8\x51\x75\x7c\x4d\x09\x2e\xa6\xc7\x48\xf0\x71\x39\x32\xa8\x93\xaa\xb8\x90\x18\xf1\xac\xe4\x0c\x1e\x12\x23\x07\x3b\xaa\x8c\xab\x94\x60\x3c\x41\x9c\x87\xbc\xe8\x94\x26\x77\x70\x25\xe5\x60\x47\x91\x7f\x9f\x12\x9c\x97\x7c\x2b\x6c\x7c\x11\xcf\xa8\x07\xf8\xef\x41\x1e\x55\xc2\xa7\x87\x1c\xcf\xaa\xa2\xe4\x69\xcc\xec\xe7\xe7\x9e\xcf\xee\xcb\x9f\x85\xe2\x8e\xe5\x2d\x12\x06\xb8\x32\x5d\x20\xf4\xf0\xf1\xc3\x79\xcf\x84\x32\xe5\xbd\xc0\xb1\x10\x32\x67\x12\x6d\xe7\x2e\xcb\xe9\xaf\xce\x42\xdc\x49\x84\x55\xef\x72\x94\xa5\x53\x79\x61\xa7\xf1\xdf\x82\x24\x03\x0b\x24\xb0\x01\xa2\x59\x84\x55\x20\x6c\xa0\x70\xd5\xb7\xe0\x6c\x2b\xa8\xae\x5e\xde\x4d\x1a\x24\xa1\x04\x89\x49\xa7\xd5\xfd\xcb\x92\xef\xc9\x99\x2b\xaf\x3d\x9f\xce\x4f\x15\x69\x63\x3d\xc8\x1a\xe5\xcd\xcc\xdd\xee\x3b\xf7\xac\x2b\x4a\x2e\x8c\xfa\x3d\x84\x8d\xbd\x49\xbe\xc3\xac\x82\xf1\xe7\x23\x88\x85\x47\x04\x72\x40\x75\x7e\xb6\x4b\xb1\x75\x89\xb4\x71\x73\x65\x70\x21\x0c\x27\x17\x78\xfe\xb4\x0b\x30\x25\xf4\x4d\x00\x61\x15\x48\x67\xd3\x7e\x1f\x4a\x2e\xf2\xd2\x66\x7b\x71\xdf\xe2\x3f\xbd\x1b\xe3\xc2\x5c\x78\xd7\xb5\xf8\x2c\xc1\x66\x1d\xd1\x8f\xbb\xaf\xc1\xd2\xa1\x30\xda\xde\xb0\xe9\xe7\x6e\xd6\x68\x2a\x79\x7a\x7f\x3a\xff\x92\xc7\x4f\x98\x8d\xfb\xaf\x4d\xd0\x86\xb9\x66\xc4\x1d\x4e\x8e\x1b\x60\xc9\xd3\xad\x5d\xc9\xd3\xed\xea\x7f\x01\x00\x00\xff\xff\x7c\x0a\xd5\x00\x75\x15\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 5493, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _offerletterHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x4f\x23\x37\x10\x7f\xbf\x4f\x31\xf8\xe5\x40\x62\xe3\x04\xd4\xaa\xa5\xbb\xd1\x9d\x80\x6b\x55\x95\x5e\x25\xa2\x4a\xb4\xea\x83\xb3\x9e\xcd\x1a\xbc\xf6\xca\x9e\x0d\x87\xaa\x7e\xf7\xca\xfb\x27\x64\xc3\xb6\x70\x25\x69\xee\x05\x6c\xcf\xf8\xe7\x99\x9f\xc7\x99\x99\x8d\x0f\x2e\x3e\x9e\xcf\x6e\x7e\xb9\x84\x9c\x0a\x3d\x7d\x13\x37\xff\x00\xe2\x1c\x85\x0c\x03\x80\xb8\x40\x12\x90\xe6\xc2\x79\xa4\x84\x55\x94\x45\xdf\x30\xe0\xeb\x42\x23\x0a\x4c\xd8\x52\xe1\x7d\x69\x1d\x31\x48\xad\x21\x34\x94\xb0\x7b\x25\x29\x4f\x24\x2e\x55\x8a\x51\x3d\x39\x06\x65\x14\x29\xa1\x23\x9f\x0a\x8d\xc9\xe4\x11\x8a\x14\x69\x9c\xbe\x2f\x4b\x7f\x6e\x25\xc2\xc7\x2c\x43\x07\x3f\x21\x11\x3a\xf8\x1e\x0d\x3a\x41\xd6\xc5\xbc\x51\x6b\xb6\x68\x65\xee\xc0\xa1\x4e\x98\xcf\xad\xa3\xb4\x22\x50\xa9\x35\x0c\xe8\xa1\xc4\x84\xa9\x42\x2c\x90\x97\x66\xc1\x20\x77\x98\x25\x2c\x27\x2a\xfd\x19\xe7\xa9\x34\x23\x51\x96\x3e\xb5\x12\x47\xa9\x2d\x78\xad\xe9\x79\xe9\xac\xac\x52\xf2\xbc\x13\xf2\x00\xe7\x79\x26\x96\x61\x10\x9d\x9e\x7c\x3a\x3d\x19\x05\xc0\x35\x03\xea\x11\xb4\x66\xd0\x83\x46\x9f\x23\x12\x6b\x97\x9f\x1e\x7c\xeb\x25\x6a\xb5\x74\x23\x83\xc4\x4d\x59\xf0\x79\xa5\x0b\xf1\x6e\x3c\xfa\x76\x34\xe6\xa9\xf7\xcd\x7c\x54\x28\x33\x4a\xbd\x6f\x70\xf8\xba\xc3\x7d\xc8\xca\x94\x77\x8b\xda\x89\x7a\x5f\x14\x68\x35\x52\xb8\x77\x5f\x8f\x26\xa3\x31\x97\xca\xd3\x23\xea\x4a\xba\x82\x7f\x62\x77\x7d\xfd\xbc\xbb\xff\x78\x6e\xe5\x43\x7b\xb8\xc7\x94\x94\x35\x90\x6a\xe1\x7d\xc2\xda\x69\x4b\x05\x40\x2c\xd5\xb2\x93\x85\x08\x10\xca\xa0\x5b\x49\x37\xe5\xba\x2a\x8c\x07\xe5\xa3\xc2\xce\x95\xc6\x30\x4a\xd1\x10\x3a\x94\x6b\x7b\x86\x76\x05\xd5\x5c\xe8\xac\xa7\x16\x42\x76\xd2\xe9\xd5\x21\xc2\x9e\x0f\xa5\x7c\xb2\x01\x91\x59\x57\x80\xa8\xdd\x4a\x18\x83\x02\x29\xb7\x32\x61\xa5\xf5\xb4\x71\x5a\xdf\xb0\x4c\xa1\x96\x4f\x34\xc2\x7d\x89\x39\xea\x4e\xab\x9e\xb0\xe9\xcf\xa2\xc0\x98\xd7\x93\x81\x1d\x1b\x24\x3a\xab\x07\x70\x01\x62\x65\xca\x8a\x06\x04\xd0\xbe\xc6\xf0\x97\x0d\xca\x5b\xf4\x1a\x60\x58\xa3\x79\x3d\x84\x9f\xfe\x41\x5e\x6a\x91\x62\x6e\xb5\x44\x97\xb0\x21\x15\x3e\xe0\x18\x97\x6a\xf9\x84\xc3\x66\x71\x5b\xcc\x5e\x16\x42\xe9\xdd\x53\x8b\xe1\x98\xd7\x71\xfb\x2f\x10\xcf\x92\x0b\xb0\x14\xba\xc2\xd7\x31\x0f\x10\x1f\x44\x11\xc4\x65\x67\x72\x8e\xba\x0c\x2f\x4b\x0a\xb3\x08\xcf\x76\x96\x2b\x0f\xb5\x99\xa0\x3c\x28\xb3\x14\x5a\xc9\x98\x97\x53\x88\xa2\x1d\x5f\xe3\x0c\x35\x96\xb9\x35\xdb\x7a\x25\xed\xb5\x11\x6a\xd6\xbf\xa0\xf5\x40\xdf\xe0\x7d\x5f\x21\xfc\x5e\x4a\x87\xde\x6f\xd7\x73\xd1\x80\x46\x5a\x19\x8c\x26\x2f\x26\xe1\x9a\x1c\x22\x41\x6b\xd2\xe7\x53\xb2\x65\x4e\xe0\x30\xd8\x0f\x27\x47\x3b\x24\xe7\xe4\xc5\xe4\xcc\xec\xbd\xd9\x77\x94\xb4\x8c\x9c\xee\x92\x91\xd3\x17\x33\x72\xae\xe8\xe1\x18\x7e\x53\x25\x84\xa4\xbb\x37\x6e\x7e\xb4\x73\x98\x85\x02\x60\xf7\x89\xa0\xa9\x33\xfe\x8f\x24\x7b\x6d\x33\xba\x17\x0e\xe1\xd2\x2c\x94\x41\x74\x5f\x50\xd6\xbd\x16\x5a\xb8\x87\xed\x86\xa0\xaf\x31\x87\x43\xcf\x54\xc5\x1c\xdd\x46\xf0\x7d\x35\x1e\x8f\xc7\x0c\x0a\x65\x12\x36\xd9\x6b\xec\x5d\x93\x70\x04\x17\x82\xb6\x9c\xbe\x7c\xc0\x8d\xa4\x20\x1c\xa6\xa5\x91\x6c\x37\x8b\x1d\x44\xd1\x36\xa9\x69\xea\xf0\x5f\x43\x25\x01\x1f\x9c\x2d\xb6\x4b\x90\x0d\xe8\xd1\x3e\x68\x7a\x0d\x27\x76\x8d\x93\x99\xdd\x05\x23\x68\xe4\x8e\xf9\xe8\x2f\x86\xfa\xf0\x39\x86\x42\xa5\xb9\x70\xb6\x2a\x71\x90\xac\x97\x3a\x3b\xaf\x88\x1e\x1b\xd2\x76\xa6\xea\xcc\x75\xc7\xba\x32\xd9\x57\xf3\x42\x11\x9b\xb6\x8d\x5f\xbf\x1d\x8c\x79\xb3\xeb\xbf\x3b\x1c\xf3\xd0\x35\xf6\x5a\xd6\xbe\x52\x6f\xba\x36\x89\x79\xdb\x41\xb7\x74\xc5\x3e\x75\xaa\x24\xf0\x2e\xfd\xbc\xee\xfe\x76\xb0\xb9\xbf\xf5\x6c\x1a\xf3\x06\xb3\x3d\xf0\x20\x8a\x9a\x5f\x27\x9b\x41\x7f\x07\x5c\xd7\x7a\x5d\x69\xdf\x99\xf2\x98\xa3\xf8\xad\x58\x8a\x66\x75\x75\x15\x87\x59\x65\x9a\x2f\x02\x87\x47\xf0\xe7\xca\xdf\x1a\xf9\xbc\x33\x45\x10\x89\x34\x3f\x7c\xfb\xfb\x5a\xb8\xfd\xf1\xf6\x78\x4d\xbf\x49\x85\x67\xd0\xc8\x8e\xd7\xd6\xc3\xc2\x07\xeb\x0a\x41\x67\xc0\xae\xae\xae\xe0\xe2\x18\x6e\x6e\x6e\x6e\x7a\x4a\x3e\xb7\xf7\x3f\xa0\x90\xe8\xce\x20\x13\xda\xe3\x4a\xf6\xd7\xd1\x77\x6f\xba\xd1\x61\x3b\x7e\xca\xc8\xa5\x91\xcf\xf0\x11\xf3\xe6\x23\x48\xcc\x9b\xcf\x63\x7f\x07\x00\x00\xff\xff\xf1\xb0\xd0\xbe\x36\x13\x00\x00")

func offerletterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_offerletterHtml,
		"offerletter.html",
	)
}

func offerletterHtml() (*asset, error) {
	bytes, err := offerletterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "offerletter.html", size: 4918, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pricingHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x4d\x8f\xdb\x36\x10\xbd\xf7\x57\x4c\x78\x49\x0a\x54\x66\x36\xb9\xb4\x85\xe4\x76\xbb\x09\x82\x7e\xa0\x75\xdb\x34\x45\x4e\x01\x4d\x8e\x2d\xd6\x14\xc9\x92\x94\x13\x77\xb1\xff\xbd\xa0\x28\x39\xfe\x90\xb5\x4e\x56\x8b\x06\xe8\x1e\xd6\xa2\xf8\x38\xf3\xf8\x66\x44\x8d\x26\x7f\xf0\xec\x97\xab\x97\xaf\x67\xcf\xa1\x0c\x95\x9a\x7e\x96\xa7\x1f\x80\xbc\x44\x26\xe2\x05\x40\x5e\x61\x60\xc0\x4b\xe6\x3c\x86\x82\xd4\x61\x91\x7d\x49\x80\xee\x4e\x6a\x56\x61\x41\xd6\x12\xdf\x5a\xe3\x02\x01\x6e\x74\x40\x1d\x0a\xf2\x56\x8a\x50\x16\x02\xd7\x92\x63\xd6\x0c\xbe\x00\xa9\x65\x90\x4c\x65\x9e\x33\x85\xc5\xc5\x7b\x53\x41\x06\x85\xd3\x4b\x6b\xfd\x95\x11\x08\x33\x67\x44\xcd\x03\xcc\x9c\xe4\x52\x2f\xe1\x27\xe9\x43\x4e\x13\x28\x2d\x50\x52\xaf\xc0\xa1\x2a\x88\x2f\x8d\x0b\xbc\x0e\x20\xb9\xd1\x04\xc2\xc6\x62\x41\x64\xc5\x96\x48\xad\x5e\x12\x28\x1d\x2e\x0a\x52\x86\x60\xfd\xd7\x94\x72\xa1\x27\xcc\x5a\xcf\x8d\xc0\x09\x37\x15\x6d\x90\x9e\xda\xe4\xd1\xd3\x6e\x92\x46\x73\x9e\x2e\xd8\x3a\x5e\x64\x4f\x9f\xbc\x7b\xfa\x64\x12\x0d\xee\x10\x68\xae\xa0\xa5\x11\x36\x0a\x7d\x89\x18\x48\x7b\xfb\xd8\xf1\x5f\x5e\xa0\x92\x6b\x37\xd1\x18\xa8\xb6\x15\x9d\xd7\xaa\x62\xdf\x3e\x9e\x7c\x35\x79\x4c\xb9\xf7\x69\x3c\xa9\xa4\x9e\x70\xef\x93\x9d\x46\xa1\x9c\x76\x31\xc9\xe7\x46\x6c\x5a\x0a\x1e\x79\x90\x46\x03\x57\xcc\xfb\x82\xb4\xc3\x96\x20\x40\x2e\xe4\xba\x9b\x8b\x51\x61\x52\xa3\xdb\xce\x1e\xce\xab\xba\xd2\x1e\xa4\xcf\x2a\x33\x97\x0a\xe3\x15\x47\x1d\xd0\xa1\xd8\x59\xd3\xb7\x2a\x42\x4b\xa6\x16\x7b\xb0\x98\x46\x17\x1d\xae\x09\x1c\xb9\x2d\xbc\xe5\xc5\x81\x81\x85\x71\x15\xb0\x66\x53\x05\x21\x50\x61\x28\x8d\x28\x88\x35\x3e\x1c\xf8\xda\xa7\xb5\x90\xa8\xc4\x11\x22\xc6\x8c\xcd\x51\x75\xa8\x66\x40\xa6\x3f\xb3\x0a\x73\xda\x0c\x7a\x56\x1c\x48\xe8\x8c\xea\xb1\x0b\x90\x4b\x6d\xeb\xd0\x33\x01\xed\xf3\x11\xff\x93\xde\xf9\xd6\x7a\x63\xa0\x1f\x91\x32\x3a\xe0\xbb\x13\xf3\x56\x31\x8e\xa5\x51\x02\x5d\x41\xfa\x20\xb4\x67\x63\x54\xc8\xf5\x91\x86\xe9\xe6\x58\xca\xfe\x69\xdc\x0a\x9e\x57\x4c\xaa\xfb\xd7\x17\xa3\x9b\xbb\x09\x3c\x60\xe2\x56\x85\x01\xd6\x4c\xd5\x78\x37\xf9\x01\xf2\x07\x59\x06\xb9\xed\x28\x97\xa8\x6c\x7c\xb8\x04\xd3\xcb\xf8\xe4\xbe\x2c\xa5\x87\x86\x26\x48\x0f\x52\xaf\x99\x92\x22\xa7\x76\x0a\x59\x76\xcf\xb1\xbc\xba\x1a\x29\x86\x6d\xbc\x38\x27\xfb\x81\xd9\xcd\xf2\x03\xbd\xff\xab\xfc\xfd\xc1\xcc\xa1\x39\xb9\xc6\xdd\x7a\x3a\x0c\x3f\xf9\xdd\xcf\x4a\xa3\xc7\xde\x39\x2a\xb4\xd1\xec\xa7\xbf\xfb\x2b\x53\x59\xa6\x37\x0d\xf1\xfb\x3f\xbd\x78\xf2\xf6\x7f\x7b\x41\xb4\x65\xc0\xdd\xf5\xdd\xc1\x79\x54\xc8\x43\x53\xc7\xd4\x2a\x48\x1b\xeb\x8e\x5e\x51\xf2\x0e\x19\x2b\x8a\x44\x84\xb4\xe1\xd8\x0e\x3b\x13\xe0\xe5\x3f\x58\x90\x8b\xc7\x27\x6c\x01\xe4\xc6\x36\xa5\x58\xfb\x1a\x58\xd5\x73\x14\xf3\xac\xa9\x9f\xac\x93\x1e\x09\x5c\x5f\x83\x5c\x00\xfe\x0d\x93\xae\xfa\xe9\x43\xdd\xdc\x24\x5a\x28\xae\xaf\x01\xb5\x80\x9b\x9b\xe9\x8f\xf5\x1c\x9f\x7d\x07\xcf\xb7\x30\x78\x74\xa9\x75\xcd\xd4\xe7\x39\x4d\x6e\x3f\x8c\x95\x65\x9b\xe5\x10\x9f\x34\x3f\xc0\x64\xc6\x36\xd9\xa5\xcf\x5e\x9b\x3a\x7b\x61\xe0\xd1\xec\xf2\xf5\x8b\xdb\xa9\x34\xaf\xb6\x5e\x3e\x0e\x3d\x2a\x85\x6e\x88\xd3\x7b\xcc\x00\xaf\xdf\x5a\xd0\x96\x4b\xcf\x7b\xf1\x0c\x3e\xb5\x56\xb2\x92\x01\xc5\x10\xa1\x1d\xd0\x00\xa3\x3f\x3a\xd4\x71\xc8\x86\xb8\xed\xd3\xf2\x81\xf9\xf2\xb6\x5c\x3a\x06\xf5\xd1\xfa\x3d\xa2\x46\xc8\xa4\xe4\xed\x64\x22\xed\x4e\x9f\xa6\x31\x56\x1a\x25\x6f\xc3\x51\x3b\xc2\x9c\xa6\x35\x4a\xcc\x62\x96\xc4\xcf\x97\x73\x8e\x80\x23\xdc\xa9\x84\x6a\xbe\x87\xc6\x39\x06\x1a\x9f\x83\x07\xc1\x0e\x62\x90\xcf\x47\x45\xf1\x98\xd0\x9a\xd5\x2a\x9c\xa3\xd6\x31\xf0\x14\xbd\x57\x11\x39\x92\x5e\xc9\xeb\xa0\x60\xbb\x90\x61\x4a\x23\x48\xb6\x36\x1b\xb6\x44\x77\x9b\x60\x7d\xb0\x3e\x6e\xaf\x12\x6e\x04\xb1\x3a\x8f\x27\xa5\xda\x07\x0c\x91\xf9\x70\x99\x72\x9a\xac\xf5\x16\x09\x27\x3e\xb9\x46\xae\x73\xce\x2d\x59\xf6\xea\x21\x5e\x22\x5f\xcd\xcd\xbb\x53\xa5\xca\x5e\x11\x6d\x7c\x57\x30\x6f\x97\x75\xe2\x07\x57\x23\xe9\x29\xe6\xe2\xdf\xf7\xc0\x96\x0e\x11\x82\x81\x50\x62\xbf\x1f\x76\xd0\xb3\xda\x6b\x94\x29\x5c\x32\x45\x83\xf1\xb4\xbf\xca\x04\x98\x06\x74\x95\x07\xa6\x05\x70\xa3\x85\x8c\x81\xf2\x39\x65\xbd\xf0\xfe\x18\x9d\xaa\x04\x07\xa2\x74\x70\xf3\x30\x48\xb1\x14\x5c\x3a\x53\x5b\xbc\x53\xbc\xe6\x75\x08\xef\x5b\x6d\xed\x48\xfa\x4c\x49\xbd\xda\xea\xbf\x44\x8d\x8e\x05\x24\x60\x34\x57\x92\xaf\x8a\x87\xa1\x94\x7e\x12\x0f\xca\x49\xd7\xc8\xa2\x6f\xa8\x4d\x7d\x2f\xf2\x70\xfa\xa2\x5d\x01\xbf\xd6\x26\xb0\x08\xc8\x69\x32\x7e\x7e\xd3\xe0\xae\x5b\x68\xdb\x0b\xdd\x26\x52\x1f\xe4\xbc\x1d\x7c\xe3\x51\x8b\x37\xcd\x8a\xa2\x49\xbe\x87\xd3\xa6\xe1\xf3\xd1\xdb\xe9\xb9\x99\xd3\xe8\x7c\xaf\x05\xb9\x0f\xda\x1b\xee\x0c\xe2\x69\xc0\xbb\xc3\x22\xa7\xa9\x71\x9a\xd3\xd4\xe6\xfe\x37\x00\x00\xff\xff\x96\x98\x49\xff\xfe\x16\x00\x00")

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

	info := bindataFileInfo{name: "pricing.html", size: 5886, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"eula.html":        eulaHtml,
	"index.html":       indexHtml,
	"offerletter.html": offerletterHtml,
	"pricing.html":     pricingHtml,
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
	"eula.html":        {eulaHtml, map[string]*bintree{}},
	"index.html":       {indexHtml, map[string]*bintree{}},
	"offerletter.html": {offerletterHtml, map[string]*bintree{}},
	"pricing.html":     {pricingHtml, map[string]*bintree{}},
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
