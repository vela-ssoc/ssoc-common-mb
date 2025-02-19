package gridfs2

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

type File interface {
	io.ReadCloser
	GridFile() *model.GridFile
}

type gridFile struct {
	file *model.GridFile
	qry  *query.Query

	serial int
	chunk  []byte
	err    error
}

func (gf *gridFile) GridFile() *model.GridFile {
	return gf.file
}

func (gf *gridFile) Read(b []byte) (int, error) {
	if err := gf.err; err != nil {
		return 0, err
	}

	except, n := len(b), 0
	for except > n && gf.err == nil {
		if len(gf.chunk) == 0 {
			if err := gf.loadChunk(); err != nil {
				gf.err = err
				break
			}
		}
		i := copy(b[n:], gf.chunk)
		n += i
		gf.chunk = gf.chunk[i:]
	}
	if n > 0 {
		return n, nil
	}

	return 0, gf.err
}

func (gf *gridFile) Close() error {
	if gf.err == nil {
		gf.err = os.ErrClosed
	}
	return nil
}

func (gf *gridFile) loadChunk() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	tbl := gf.qry.GridChunk
	chk, err := tbl.WithContext(ctx).
		Where(tbl.FileID.Eq(gf.file.ID), tbl.Serial.Eq(gf.serial)).
		First()
	if err != nil {
		return err
	}

	gf.serial++
	gf.chunk = chk.Data

	return nil
}
