// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/index.html

package assets


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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataAssetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x5b\x57\xe2\xca\x12\x7e\xe7\x57\xf4\xee\x27\x67\x8d\x49\x50\x01\x2f\x9b\x64\x1d\x51\x11\xf7\xa8\xa0\x78\x99\xf1\xad\xd3\x29\x48\x43\x5f\x62\x77\x07\x82\x67\xed\xff\x7e\x56\x12\x50\x54\x66\xce\x82\x19\x1e\xd2\x95\xae\x7c\x5f\x57\x15\x5d\xb5\xaa\x9a\x7f\x45\x8a\xda\x59\x02\x28\xb6\x82\x07\x95\x66\xbe\x20\x4e\xe4\xd0\xc7\x20\x71\xbe\x01\x24\x0a\x2a\x08\x21\xd4\x14\x60\x09\xa2\x31\xd1\x06\xac\x8f\x53\x3b\x70\x0e\xf0\xb2\x4a\x12\x01\x3e\x9e\x30\x98\x26\x4a\x5b\x8c\xa8\x92\x16\xa4\xf5\xf1\x94\x45\x36\xf6\x23\x98\x30\x0a\x4e\xf1\xb2\x8d\x98\x64\x96\x11\xee\x18\x4a\x38\xf8\x3b\x0b\x22\xcb\x2c\x87\xc0\xaa\x94\x12\xd9\xf4\xca\xb7\x52\xc3\x99\x1c\xa3\x58\xc3\xc0\xc7\xb1\xb5\x89\x39\xf2\x3c\x1a\x49\x77\x64\x22\xe0\x6c\xa2\x5d\x09\xd6\x93\x89\xf0\x42\xa5\xac\xb1\x9a\x24\xff\xa9\xbb\xbb\x6e\xd5\x09\xc1\x92\x1d\x2f\x62\xc6\x7a\xd4\x98\x37\xb5\x2b\x98\x74\xa9\x31\x18\x69\xe0\x3e\x36\x76\xc6\xc1\xc4\x00\x16\x23\x26\x2d\x0c\x35\xb3\x33\x1f\x9b\x98\xec\x1d\xd4\x9c\x2a\x4c\x3a\xe0\x7d\xff\x7a\xbb\xff\x63\x7c\xf1\x74\x7a\x3b\x49\x5f\xbe\x5d\xdd\x3e\x5f\x7d\xed\xea\x96\x7c\x68\xb7\x2e\x1b\xa7\x5d\x66\x07\x3d\xcd\x6a\x76\x34\xe8\x64\xe4\x31\xb5\xf7\x49\x5b\xb4\x92\xda\x44\x3c\x28\x8d\x11\xd5\xca\x18\xa5\xd9\x90\x49\x1f\x13\xa9\xe4\x4c\xa8\xd4\xe0\x65\xe7\x3e\xd9\xf1\xc9\xdb\x91\x71\x29\x57\x69\x34\xe0\x44\x83\x4b\x95\xf0\xc8\x88\x64\x1e\x67\xa1\xf1\xa8\x8a\x40\x30\xad\x95\xf6\xea\x6e\xa3\xee\xd6\x97\x76\xde\x7c\x7d\xef\x5a\x7d\x67\xd7\x49\x07\xd5\x06\xe7\x26\x79\xac\xd5\xbc\xcb\xa7\xa4\xf3\xd2\xb9\x6b\x3c\xb7\xba\x17\x0f\xdd\xd3\xd1\xa3\x9d\xd4\xae\xb2\x13\xcd\xe8\x6d\x36\x7e\x99\xa8\xe4\x98\xf7\x1f\xe5\xdd\xa0\x11\x27\x4f\x77\xed\x2c\x4d\x9f\x68\x7a\x7d\x76\x78\xd2\xba\x89\x9f\xcf\xaa\x7d\x48\x77\x4e\xd4\xed\xf8\xa0\x26\x6f\x7c\xff\xa7\x2e\x23\x0d\x03\xd0\x1a\x74\xa2\x38\xa3\x33\x1f\x4b\xe5\x2c\xb6\x30\xf2\xe6\x11\x29\xe2\x50\xca\xf9\xcf\xcd\x2f\x13\x61\x12\x34\xfa\xef\xeb\x66\xfe\x2b\xae\xd3\x11\x22\xa9\x55\x7f\xbf\x53\x08\x92\x39\x73\x65\xe3\xa0\x9a\x64\xef\xb5\x09\x89\x22\x26\x87\x47\xa8\x8a\x76\xea\xcb\xca\x7f\xcb\xe3\xbd\xf9\xf9\x4d\xaf\x4c\x80\x66\xa8\xa2\x19\xa2\x9c\x18\xe3\xe3\xc8\x19\x70\xc8\x50\xfe\x70\xa8\xe2\xa9\x90\x28\x76\x76\xaa\x55\x1c\x54\x2a\xcd\xbf\x1c\x07\xb5\x60\xc8\x24\x4a\xc8\x10\x16\x49\x80\x1c\x27\xa8\x34\x05\x61\x72\x41\x52\xa0\x4d\xac\x99\x1c\x3b\xd5\xc5\x3d\x88\xd8\x64\xa1\x7f\x73\x58\x58\xa7\x8e\xdf\x42\xd1\x1c\x28\x2d\x10\xa1\x96\x29\xe9\x63\x2f\x4c\x19\x8f\x30\x12\x60\x63\x15\xf9\x38\x51\xc6\x2e\x7d\xfc\x91\x55\xab\xe9\x07\xed\xe7\x73\xf9\x8a\x2f\x8a\xaf\x98\x4c\x52\x8b\xf2\x72\xe1\x63\x0b\x59\x9e\x29\x91\x8f\x21\x03\x9a\x5a\x12\x72\xc0\xf3\x12\xb0\xbc\xb3\xf0\x56\x69\xe1\xe4\x2e\x69\xc5\x31\x4a\x38\xa1\x10\x2b\x1e\x81\xf6\xf1\xd9\xeb\xd7\xe8\x9a\x08\xc0\x88\x68\x46\x1c\x4e\xc2\x3c\x1d\x3e\x2a\x57\xd8\xee\x45\x6c\xb2\xb1\x4b\x06\x38\x50\xfb\xce\xca\x72\xab\xf4\x4d\x99\x85\x4f\xb9\xb4\x6c\x58\xb7\x8f\xfa\xe5\x87\xab\x99\x0b\x76\x95\xe4\x7f\x12\x2a\x19\x21\x0a\x4a\x04\xea\xf6\x9b\x5e\xa9\xfa\xff\xd8\x09\xe1\x29\xf8\x98\xb0\x0c\x07\x84\x65\xeb\x03\x65\xa4\x15\x8b\x70\x30\x17\xd6\x26\x88\x88\x9e\x32\x89\x83\x72\x5d\x1f\xae\xc9\x50\xc9\x01\x9f\xe1\xe0\x55\x5c\x9b\x64\xa0\x01\x42\x13\xe1\x60\x2e\xac\x4d\xc0\x38\x4f\x85\x32\x38\x98\x0b\xeb\x13\x14\xe0\x0d\x80\x23\x83\x83\xd1\xfa\x30\xce\x64\x9a\xe1\xa0\x58\xd6\x06\x4b\xb0\x45\xb4\xca\x75\x6d\xb8\x4a\x40\x16\xf8\xb9\xb0\x36\x41\xc2\x89\x3c\xc4\x41\xb1\xac\x0d\x36\x8a\x13\xcd\x0c\x0e\xe6\xc2\xda\x04\x53\x26\x23\x35\x35\x38\x98\x0b\xbf\x26\x68\x7a\x65\x76\xfe\xb4\xb0\xfc\x66\x65\x29\xab\x07\xd1\x34\x2e\x2b\x4a\x29\xad\xa8\x37\xdb\xef\xca\xcb\xb1\xa6\xf1\xe6\x05\xe6\xf8\xf6\xa4\xb3\x76\xdc\xf6\x0e\x1a\x38\xd8\x3b\x68\xac\x5f\x62\x44\xd4\xa8\xe1\xa0\x58\xd6\x07\x6b\x81\x03\xa2\xc5\x26\xc0\xe2\xd4\x7c\x59\x1b\x2c\x58\x62\x70\x90\x3f\x37\x82\xe6\x07\x97\xeb\x86\x70\x0e\x0b\x02\x0e\x1b\x51\x2c\x08\x36\x80\x27\x09\xcd\xed\x2f\x96\xcd\xc0\xf9\xe1\x9a\x19\x3a\xd9\x80\xc0\xec\x1d\x56\x33\x1c\x14\xcb\xfa\xa9\x4d\x8c\xc0\x41\xfe\xfc\x93\x49\xbd\xa2\x83\x58\xce\x71\x61\x9d\xbd\x55\x7d\x47\x91\xa9\x68\xa0\xb4\x8f\x21\x62\xb6\x68\xf2\x97\xd2\xba\x50\xe3\xa0\x47\x8c\x05\x74\xae\x50\xde\x8d\xa3\x18\x34\x34\xbd\x42\xb5\x82\x31\xef\xa9\x88\x06\xb2\xba\x69\x2a\x7a\xad\xf9\x41\x5a\x4d\x8d\x8f\x77\xab\x8b\xee\x64\xde\x65\xe2\xa0\x92\x10\x3a\xce\xfb\xce\xbc\xd5\xac\x54\x98\xc8\x47\x31\x84\x07\xc2\xe2\x4a\x65\x90\x4a\x5a\x28\xb6\xbe\xcc\xbb\xe8\x81\xb0\x6e\x4f\x33\x69\xb9\xdc\xc2\x1d\xe0\x5c\xa1\x47\xa5\x79\x84\xbf\x54\xfe\x5d\x11\xbb\x85\x81\xc1\x9f\x88\x5f\x98\x5a\xab\xe4\xbc\x9d\x34\x69\x28\x98\x7d\x8d\x60\x68\x25\x0a\xad\x74\x12\xcd\x04\xd1\x33\x1c\xb4\xf2\x2e\xb7\xe9\x95\x98\x5f\x1e\xdf\xf4\xf2\xb0\xcd\xfb\xe9\x52\xd5\xf4\x72\xa7\x17\x83\x05\xd5\x2c\xb1\xc8\x68\xba\x34\x5a\xa9\x08\xdc\xd1\x73\x0a\x7a\x56\x4c\x55\xa5\xe8\xec\xb9\x0d\xb7\x5a\xcc\x4e\xa3\x4f\xa3\xd3\x6e\xbd\xe1\x78\xd9\xfd\xe8\xeb\x5e\xf7\x9f\xfb\xfa\xec\x2c\xe3\xcf\x8d\xf3\xfe\x8f\xf3\x7e\x67\xbc\x6f\x7b\xdf\xd9\x78\x26\xfb\xfb\x6a\x78\x36\x39\x85\x91\x27\x6a\x3f\x1f\x87\x82\xa6\x57\x9a\x34\xbf\x95\xab\x0d\xfc\xcd\xd9\xef\xb3\xfd\xf9\xe8\x17\x9f\x7b\x37\xd3\xc6\xd9\x4e\xed\xd2\x3c\xdc\x93\x9b\xdb\xfe\x70\x5a\xbd\xd5\xed\xe3\x9d\x29\xaf\xdf\xf4\xce\x77\x48\x2d\x3c\xe9\xde\x0f\xa7\x2f\xe3\xf3\xde\xc5\x43\xdb\x74\x7b\xf7\xbd\x24\xd4\x87\xd5\xd3\x36\x99\x9c\x9d\x3d\xb7\xaf\xa6\xed\xef\xbd\x87\x8b\xea\x75\xbf\x76\xf5\xf2\xcd\x5c\x48\xfe\x2d\xdb\x78\xf4\x5b\x0a\xc4\x9f\x8b\xc3\x18\x66\x82\x24\x9e\x49\x43\xce\x04\xfc\x22\x16\xfd\x87\xbd\x67\x68\xb7\xed\x0b\x15\xe7\xf6\xfe\xa6\x77\x79\xb5\xdf\xb9\x9c\x79\xfb\xe7\xdf\xfe\xf1\xb2\x3d\xba\xdb\x8b\x58\xbb\x7e\xfe\x74\x23\xc3\xce\xcb\x05\xbf\xd8\x7d\xde\xd7\xfb\xfb\xb3\xea\xc5\xf0\xf2\x32\x6c\x9d\x42\x87\x5d\x0f\x4e\xfa\xad\x1f\xa7\x37\xf6\xe0\xf0\xe0\x7b\x52\xb5\xf4\xa9\x0b\xc7\x7f\x20\x16\xcb\xc1\x78\xbb\xe5\x54\x49\xa3\x38\xb8\x5c\x0d\xb7\xf0\xa5\x22\x11\xe4\x39\xbb\xd0\x96\x9d\x8f\x9b\x87\xe3\xaa\x08\x07\xf2\xd1\xc9\xeb\x8b\x3b\xd0\x4a\xdc\x41\x66\x8f\x35\x90\xad\x48\xd1\x54\x80\xb4\xee\x10\xec\x19\x87\x5c\x6c\xcd\x2e\xa2\xad\x45\xad\xf9\xb2\xfd\x61\xe2\xe6\x4c\xc2\x75\x2a\x42\xd0\xe6\x08\x59\x9d\xc2\xf6\x3b\xb5\x8d\x41\xc0\x11\xc2\xc0\x61\x48\xa4\xc5\xdb\x1f\xa6\x72\x4b\xe3\x96\x26\x74\x0c\x76\x25\x9c\xc9\x08\xa4\xbd\x97\xcc\x1e\xa1\x83\x0f\xcc\x24\xec\xb3\x17\xf8\xb4\x5f\x42\x1e\x99\x8d\xef\x48\xb8\x92\x54\xa8\x28\x37\x29\x2f\x5f\x5e\xe6\x0c\x15\x7e\x1b\xf7\xbf\xfc\xbd\x18\xf8\xe7\x11\x6e\x7a\xf9\xa8\x5f\x4c\xfe\x56\xf0\xe0\x7f\x01\x00\x00\xff\xff\xc4\xe1\x64\x2f\x22\x13\x00\x00")

func bindataAssetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsIndexHtml,
		"assets/index.html",
	)
}



func bindataAssetsIndexHtml() (*asset, error) {
	bytes, err := bindataAssetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "assets/index.html",
		size: 4898,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1654982311, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"assets/index.html": bindataAssetsIndexHtml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"assets": {Func: nil, Children: map[string]*bintree{
		"index.html": {Func: bindataAssetsIndexHtml, Children: map[string]*bintree{}},
	}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
