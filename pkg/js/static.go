// Code generated by "esc "; DO NOT EDIT.

package js

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/helpers.js": {
		local:   "pkg/js/helpers.js",
		size:    19177,
		modtime: 0,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "pkg/js",
	},
}
