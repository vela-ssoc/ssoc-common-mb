package gridfs

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/fs"
	"strconv"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

type FS interface {
	fs.FS
	OpenID(id int64) (File, error)
	Remove(id int64) error
	Write(io.Reader, string) (File, error)
}

func NewFS(qry *query.Query) FS {
	return &gridFS{qry: qry, burst: 60 * 1024}
}

type gridFS struct {
	qry   *query.Query // 数据库连接
	burst int          // 60K
}

// Open 通过文件名打开文件
func (f *gridFS) Open(name string) (fs.File, error) {
	id, err := strconv.ParseInt(name, 10, 64)
	if err != nil {
		return nil, err
	}
	fl, err := f.OpenID(id)
	if err != nil {
		return nil, err
	}

	return fl, nil
}

func (f *gridFS) OpenID(id int64) (File, error) {
	tbl := f.qry.GridFile
	gf, err := tbl.WithContext(context.Background()).
		Where(tbl.ID.Eq(id)).
		First()
	if err != nil {
		return nil, err
	}

	fl := &file{
		id:        gf.ID,
		name:      gf.Name,
		size:      gf.Size,
		checksum:  gf.Checksum,
		createdAt: gf.CreatedAt,
		qry:       f.qry,
	}

	return fl, nil
}

func (f *gridFS) Remove(id int64) error {
	return f.qry.Transaction(func(tx *query.Query) error {
		fileTbl, chunkTbl := tx.GridFile, tx.GridChunk
		if _, err := chunkTbl.WithContext(context.Background()).
			Where(chunkTbl.FileID.Eq(id)).
			Delete(); err != nil {
			return err
		}

		_, err := fileTbl.WithContext(context.Background()).
			Where(fileTbl.ID.Eq(id)).
			Delete()

		return err
	})
}

func (f *gridFS) Write(r io.Reader, name string) (File, error) {
	createdAt := time.Now()
	fl := &model.GridFile{Name: name, CreatedAt: createdAt}
	digest := md5.New()
	tr := io.TeeReader(r, digest)

	if err := f.qry.Transaction(func(tx *query.Query) error {
		fileTbl, chunkTbl := tx.GridFile, tx.GridChunk
		fileDo := fileTbl.WithContext(context.Background())
		chunkDo := chunkTbl.WithContext(context.Background())

		if err := fileDo.Create(fl); err != nil {
			return err
		}
		fileID := fl.ID

		var err error
		var n, serial int
		var filesize int64
		for {
			buf := make([]byte, f.burst)
			if n, err = tr.Read(buf); err != nil {
				if err == io.EOF {
					err = nil
				}
				break
			}

			chk := &model.GridChunk{
				FileID: fileID,
				Serial: serial,
				Data:   buf[:n],
			}
			if err = chunkDo.Create(chk); err != nil {
				return err
			}

			serial++
			filesize += int64(n)
		}

		fl.Size = filesize
		fl.Checksum = hex.EncodeToString(digest.Sum(nil))
		_, err = fileDo.Updates(fl)

		return err
	}); err != nil {
		return nil, err
	}

	cf := &file{
		id:        fl.ID,
		name:      fl.Name,
		size:      fl.Size,
		checksum:  fl.Checksum,
		createdAt: fl.CreatedAt,
		qry:       f.qry,
	}

	return cf, nil

	//var n, serial int
	//var filesize int64
	//for {
	//	if n, err = tr.Read(buf); err != nil {
	//		if err == io.EOF {
	//			err = nil
	//		}
	//		break
	//	}
	//	if _, err = tx.Exec(insertPart, fileID, serial, buf[:n]); err != nil {
	//		break
	//	}
	//	serial++
	//	filesize += int64(n)
	//}
	//
	//if err == nil {
	//	sum := hex.EncodeToString(digest.Sum(nil))
	//	updateFile := "UPDATE gridfs_file SET size = ?, checksum = ? WHERE id = ?"
	//	if _, err = tx.Exec(updateFile, filesize, sum, fileID); err == nil {
	//		if err = tx.Commit(); err == nil {
	//			fl := &file{
	//				id:        fileID,
	//				name:      name,
	//				size:      filesize,
	//				checksum:  sum,
	//				createdAt: createdAt,
	//				db:        f.db,
	//			}
	//
	//			return fl, nil
	//		}
	//	}
	//}
}
