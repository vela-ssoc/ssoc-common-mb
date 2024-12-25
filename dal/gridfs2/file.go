package gridfs

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"gorm.io/gorm"
)

type Digester interface {
	MD5() string
	SHA1() string
	SHA256() string
}

type File interface {
	fs.File
	fs.FileInfo
	Digester
	ID() int64
	Shard() int
	MIME() string
}

type file struct {
	qry    *query.Query
	file   *model.GridFile
	serial int
	buff   []byte
	err    error
}

func (f *file) ID() int64 {
	return f.file.ID
}

func (f *file) MIME() string {
	return f.file.MIME
}

func (f *file) Shard() int {
	return f.file.Shard
}

func (f *file) MD5() string {
	return f.file.MD5
}

func (f *file) SHA1() string {
	return f.file.SHA1
}

func (f *file) SHA256() string {
	return f.file.SHA256
}

func (f *file) Name() string {
	return f.file.Name
}

func (f *file) Size() int64 {
	return f.file.Size
}

func (f *file) Mode() fs.FileMode {
	return 0o444
}

func (f *file) ModTime() time.Time {
	return f.file.UpdatedAt
}

func (f *file) IsDir() bool {
	return false
}

func (f *file) Sys() any {
	return nil
}

func (f *file) Stat() (fs.FileInfo, error) {
	return f, nil
}

func (f *file) Read(b []byte) (int, error) {
	if f.err != nil {
		return 0, f.err
	}

	var err error
	var n int
	bsz := len(b)
	for f.err == nil && bsz > n {
		if len(f.buff) == 0 {
			if err = f.nextSerial(); err != nil {
				f.err = err
				break
			}
		}

		i := copy(b[n:], f.buff)
		n += i
		f.buff = f.buff[i:]
	}
	if n == 0 {
		if f.err == nil {
			return 0, io.EOF
		}
		return 0, f.err
	}

	return n, nil
}

func (f *file) Close() error {
	if f.err == nil {
		f.err = fs.ErrClosed
	}
	return nil
}

func (f *file) nextSerial() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	tbl := f.qry.GridChunk
	chk, err := tbl.WithContext(ctx).
		Where(tbl.FileID.Eq(f.file.ID), tbl.Serial.Eq(f.serial)).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return io.EOF
		}
		return err
	}

	f.serial++
	f.buff = chk.Data

	return nil
}
