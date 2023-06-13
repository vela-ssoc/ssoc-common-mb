package stegano

import (
	"bytes"
	"io"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/gridfs"
	"github.com/vela-ssoc/vela-common-mba/encipher"
)

func AppendStream(file gridfs.File, v any) (gridfs.File, error) {
	enc, err := encipher.EncryptJSON(v)
	if err != nil {
		return nil, err
	}
	w := &warpFile{
		file:  file,
		read:  io.MultiReader(file, bytes.NewReader(enc)),
		size:  file.Size() + int64(len(enc)),
		mtime: time.Now(),
	}
	return w, nil
}
