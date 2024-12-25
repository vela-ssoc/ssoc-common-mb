package gridfs

import (
	"io/fs"
	"os"
	"time"
)

type diskFile struct {
	fl File
	fd *os.File
}

func (d *diskFile) Stat() (fs.FileInfo, error) {
	return d.fl.Stat()
}

func (d *diskFile) Read(b []byte) (int, error) {
	return d.fd.Read(b)
}

func (d *diskFile) Close() error {
	err := d.fd.Close()
	_ = d.fl.Close()
	return err
}

func (d *diskFile) Name() string {
	return d.fl.Name()
}

func (d *diskFile) Size() int64 {
	return d.fl.Size()
}

func (d *diskFile) Mode() fs.FileMode {
	return d.fl.Mode()
}

func (d *diskFile) ModTime() time.Time {
	return d.fl.ModTime()
}

func (d *diskFile) IsDir() bool {
	return d.fl.IsDir()
}

func (d *diskFile) Sys() any {
	return d.fl.Sys()
}

func (d *diskFile) MD5() string {
	return d.fl.MD5()
}

func (d *diskFile) SHA1() string {
	return d.fl.SHA1()
}

func (d *diskFile) SHA256() string {
	return d.fl.SHA256()
}

func (d *diskFile) ID() int64 {
	return d.fl.ID()
}

func (d *diskFile) Shard() int {
	return d.fl.Shard()
}

func (d *diskFile) MIME() string {
	return d.fl.MIME()
}
