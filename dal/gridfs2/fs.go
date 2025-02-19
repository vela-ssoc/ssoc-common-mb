package gridfs2

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

type FS interface {
	// Open 通过文件 ID 打开文件。
	Open(ctx context.Context, id int64) (File, error)

	// Remove 通过文件 ID 删除文件。
	Remove(ctx context.Context, id int64) error

	// Create 创建文件。
	Create(ctx context.Context, name string, r io.Reader) (File, error)
}

func New(qry *query.Query) FS {
	return &gridFS{
		qry: qry,
	}
}

type gridFS struct {
	qry *query.Query
}

func (gfs *gridFS) Open(ctx context.Context, id int64) (File, error) {
	tbl := gfs.qry.GridFile
	file, err := tbl.WithContext(ctx).Where(tbl.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	gf := &gridFile{file: file, qry: gfs.qry}

	return gf, nil
}

func (gfs *gridFS) Remove(ctx context.Context, id int64) error {
	return gfs.qry.Transaction(func(tx *query.Query) error {
		file, chunk := tx.GridFile, tx.GridChunk
		if _, err := file.WithContext(ctx).
			Where(file.ID.Eq(id)).
			Delete(); err != nil {
			return err
		}

		_, err := chunk.WithContext(ctx).
			Where(chunk.FileID.Eq(id)).
			Delete()

		return err
	})
}

func (gfs *gridFS) Create(ctx context.Context, name string, r io.Reader) (File, error) {
	const chunkSize = 60 * 1024

	ext := strings.ToLower(filepath.Ext(name))
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	now := time.Now()
	file := &model.GridFile{
		Name:      name,
		Extension: ext,
		MIME:      mimeType,
		ChunkSize: chunkSize,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := gfs.qry.Transaction(func(tx *query.Query) error {
		fileTbl, chunkTbl := gfs.qry.GridFile, gfs.qry.GridChunk
		fileDo, chunkDo := fileTbl.WithContext(ctx), chunkTbl.WithContext(ctx)
		if err := fileDo.Create(file); err != nil {
			return err
		}

		md5hash, sha1hash, sha256hash := md5.New(), sha1.New(), sha256.New()
		rd := io.TeeReader(r, io.MultiWriter(md5hash, sha1hash, sha256hash))

		var size, serial int
		fileID := file.ID
		for {
			buf := make([]byte, chunkSize)
			n, err := rd.Read(buf)
			if n > 0 {
				chk := &model.GridChunk{FileID: fileID, Serial: serial, Data: buf[:n]}
				if err = chunkDo.Create(chk); err != nil {
					return err
				}
				serial++
				size += n
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
		}

		md5sum := md5hash.Sum(nil)
		sha1sum := sha1hash.Sum(nil)
		sha256sum := sha256hash.Sum(nil)
		file.MD5 = hex.EncodeToString(md5sum)
		file.SHA1 = hex.EncodeToString(sha1sum)
		file.SHA256 = hex.EncodeToString(sha256sum)
		file.Size = int64(size)
		file.UpdatedAt = time.Now()

		return fileDo.Save(file)
	}); err != nil {
		return nil, err
	}

	gf := &gridFile{file: file, qry: gfs.qry}

	return gf, nil
}
