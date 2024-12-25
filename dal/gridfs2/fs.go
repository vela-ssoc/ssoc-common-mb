package gridfs

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
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

	// Delete 删除文件。
	Delete(ctx context.Context, id int64) error

	// Create 创建文件。
	Create(ctx context.Context, name string, r io.Reader) (File, error)
}

func New(qry *query.Query) FS {
	return &gridFS{qry: qry, shard: 60 * 1024}
}

type gridFS struct {
	qry   *query.Query
	shard int
}

func (gfs *gridFS) Open(ctx context.Context, id int64) (File, error) {
	tbl := gfs.qry.GridFile
	fd, err := tbl.WithContext(ctx).
		Where(tbl.ID.Eq(id)).
		First()
	if err != nil {
		return nil, err
	}
	fl := &file{file: fd, qry: gfs.qry}

	return fl, nil
}

func (gfs *gridFS) Delete(ctx context.Context, id int64) error {
	tbl := gfs.qry.GridFile
	_, err := tbl.WithContext(ctx).
		Where(tbl.ID.Eq(id)).
		Delete()

	return err
}

func (gfs *gridFS) Create(ctx context.Context, name string, r io.Reader) (File, error) {
	now := time.Now()
	ext := strings.ToLower(filepath.Ext(name))
	mtyp := mime.TypeByExtension(ext)
	if mtyp == "" {
		mtyp = "application/octet-stream"
	}
	fd := &model.GridFile{
		Name:      name,
		MIME:      mtyp,
		Extension: ext,
		Shard:     gfs.shard,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := gfs.qry.Transaction(func(tx *query.Query) error {
		ftbl, ctbl := tx.GridFile, tx.GridChunk
		fdao, cdao := ftbl.WithContext(ctx), ctbl.WithContext(ctx)
		if err := fdao.Create(fd); err != nil {
			return err
		}

		md5h, sha1h, sha256h := md5.New(), sha1.New(), sha256.New()
		mulsum := io.MultiWriter(md5h, sha1h, sha256h)
		input := io.TeeReader(r, mulsum)

		var size int64
		var serial int
		data := make([]byte, gfs.shard)
		for {
			// [io.ReadFull] 读出了数据，但是又没读满 data 就 [io.EOF]，此时会返回 [io.ErrUnexpectedEOF]，
			// 再次读取才会发生 [io.EOF]。
			n, err := io.ReadFull(input, data)
			if n == 0 {
				if err == io.EOF || errors.Is(err, io.ErrUnexpectedEOF) {
					break
				}
				return err
			}

			chk := &model.GridChunk{
				FileID: fd.ID,
				Serial: serial,
				Data:   data[:n],
			}
			if err = cdao.Create(chk); err != nil {
				return err
			}
			serial++
			size += int64(n)
		}

		md5sum := md5h.Sum(nil)
		sha1sum := sha1h.Sum(nil)
		sha256sum := sha256h.Sum(nil)
		fd.MD5 = hex.EncodeToString(md5sum)
		fd.SHA1 = hex.EncodeToString(sha1sum)
		fd.SHA256 = hex.EncodeToString(sha256sum)
		fd.Size = size
		fd.UpdatedAt = time.Now()
		_, err := fdao.Updates(fd)

		return err
	})
	if err != nil {
		return nil, err
	}

	fl := &file{qry: gfs.qry, file: fd}

	return fl, nil
}
