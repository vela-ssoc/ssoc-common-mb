package stegano

import (
	"io"
	"io/fs"
	"strconv"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/gridfs"
)

type warpFile struct {
	file  gridfs.File
	read  io.Reader
	size  int64
	mtime time.Time
}

func (w *warpFile) Stat() (fs.FileInfo, error) {
	return w, nil
}

func (w *warpFile) Read(p []byte) (int, error) {
	return w.read.Read(p)
}

func (w *warpFile) Close() error {
	return w.file.Close()
}

func (w *warpFile) Name() string {
	return w.file.Name()
}

func (w *warpFile) Size() int64 {
	return w.size
}

func (w *warpFile) Mode() fs.FileMode {
	return w.file.Mode()
}

func (w *warpFile) ModTime() time.Time {
	return w.mtime
}

func (w *warpFile) IsDir() bool {
	return false
}

func (w *warpFile) Sys() any {
	return nil
}

func (w *warpFile) ID() int64 {
	return w.file.ID()
}

func (w *warpFile) MD5() string {
	return ""
}

func (w *warpFile) ContentType() string {
	return w.file.ContentType()
}

func (w *warpFile) ContentLength() string {
	return strconv.FormatInt(w.size, 10)
}

func (w *warpFile) Disposition() string {
	return w.file.Disposition()
}
