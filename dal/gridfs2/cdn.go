package gridfs

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func NewCDN(fs FS, dir string) FS {
	// 防止自我嵌套
	if _, ok := fs.(*cdnFS); ok {
		return fs
	}

	return &cdnFS{
		gfs:   fs,
		dir:   dir,
		tasks: make(map[string]struct{}, 16),
	}
}

type cdnFS struct {
	gfs   FS
	dir   string
	mutex sync.Mutex
	tasks map[string]struct{}
}

func (cf *cdnFS) Open(ctx context.Context, id int64) (File, error) {
	fd, err := cf.gfs.Open(ctx, id)
	if err != nil || fd.Size() <= int64(fd.Shard()) {
		return fd, err
	}

	// TODO implement me
	panic("implement me")
}

func (cf *cdnFS) Delete(ctx context.Context, id int64) error {
	if err := cf.gfs.Delete(ctx, id); err != nil {
		return err
	}
	name := strconv.FormatInt(id, 10)
	_ = os.Remove(filepath.Join(cf.dir, name))

	return nil
}

func (cf *cdnFS) Create(ctx context.Context, name string, r io.Reader) (File, error) {
	return cf.gfs.Create(ctx, name, r)
}

func (cf *cdnFS) openDisk(ctx context.Context, fl File) (File, error) {
	name := strconv.FormatInt(fl.ID(), 10)
	if fd, _ := os.Open(filepath.Join(cf.dir, name)); fd != nil {
		df := &diskFile{fl: fl, fd: fd}
		return df, nil
	}

	cf.mutex.Lock()
	defer cf.mutex.Unlock()
	cf.tasks[name] = struct{}{}

	return nil, nil
}
