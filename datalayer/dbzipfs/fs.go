package dbzipfs

import (
	"context"
	"errors"
	"io/fs"
	"path"
	"sync"
	"time"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/model"
	"github.com/vela-ssoc/ssoc-common-mb/datalayer/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Zip struct {
	id     bson.ObjectID      // ArchiveID
	db     *repository.BaseDB // 数据库
	parent context.Context    // 读取数据库需要用到
}

func NewZip(parent context.Context, db *repository.BaseDB, id bson.ObjectID) *Zip {
	return &Zip{
		id:     id,
		db:     db,
		parent: parent,
	}
}

// Open 实现 fs.FS 接口。
func (z *Zip) Open(name string) (fs.File, error) {
	name = path.Clean(name)
	if err := z.validPath(name, "open"); err != nil {
		return nil, err
	}

	if name == "." {
		return z.newVirtualRoot(), nil
	}

	ctx, cancel := z.perContext()
	defer cancel()

	coll := z.db.ZipFile()
	filter := bson.M{"archive_id": z.id, "full_name": name}
	info, err := coll.FindOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fs.ErrNotExist
		}
		return nil, err
	}
	file := z.newFile(info)

	return file, nil
}

// ReadDir 实现 ReadDirFS 接口，可以被 fs.WalkDir 调用。
func (z *Zip) ReadDir(name string) ([]fs.DirEntry, error) {
	if err := z.validPath(name, "readdir"); err != nil {
		return nil, err
	}

	ctx, cancel := z.perContext()
	defer cancel()

	opts := options.Find().SetSort(bson.M{"file_name": 1})
	filter := bson.D{{Key: "archive_id", Value: z.id}, {Key: "directory", Value: name}}

	coll := z.db.ZipFile()
	files, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	dirs := make([]fs.DirEntry, 0, len(files))
	for _, file := range files {
		dirs = append(dirs, file)
	}

	return dirs, nil
}

func (z *Zip) perContext() (context.Context, context.CancelFunc) {
	if z.parent != nil {
		return z.parent, func() {}
	}

	return context.WithTimeout(context.Background(), time.Minute)
}

func (z *Zip) newVirtualRoot() *File {
	root := &model.ZipFile{
		ArchiveID:   z.id,
		FileName:    ".",
		FullName:    ".",
		IsDirectory: true,
		UpdatedAt:   time.Now(),
	}

	return z.newFile(root)
}

func (z *Zip) newFile(info *model.ZipFile) *File {
	file := &File{db: z.db, info: info, parent: z.parent}
	file.once = sync.OnceValue(file.openDownloadStream)

	return file
}

func (*Zip) validPath(name, op string) error {
	if fs.ValidPath(name) {
		return nil
	}

	return &fs.PathError{Op: op, Path: name, Err: fs.ErrInvalid}
}
