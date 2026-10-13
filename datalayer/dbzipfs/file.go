package dbzipfs

import (
	"context"
	"io/fs"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"github.com/vela-ssoc/ssoc-common-mb/datalayer/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type File struct {
	db     *repository.BaseDB
	info   *model.ZipFile
	parent context.Context

	once func() error
	stm  *mongo.GridFSDownloadStream
}

func (f *File) Stat() (fs.FileInfo, error) { return f.info, nil }

func (f *File) Read(p []byte) (int, error) {
	if err := f.once(); err != nil {
		return 0, err
	}

	return f.stm.Read(p)
}

func (f *File) Close() error {
	if f.stm != nil {
		return f.stm.Close()
	}

	return nil
}

func (f *File) openDownloadStream() error {
	coll := f.db.ZipFile().GridFSBucket()
	stm, err := coll.OpenDownloadStream(f.parent, f.info.GridID)
	if err != nil {
		return err
	}
	f.stm = stm

	return nil
}
