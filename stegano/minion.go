package stegano

import (
	"bytes"
	"io"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/gridfs"
	"github.com/vela-ssoc/vela-common-mba/ciphertext"
)

//func AppendStream(file gridfs.File, v any) (gridfs.File, error) {
//	enc, err := encipher.EncryptJSON(v)
//	if err != nil {
//		return nil, err
//	}
//	w := &warpFile{
//		file:  file,
//		read:  io.MultiReader(file, bytes.NewReader(enc)),
//		size:  file.Size() + int64(len(enc)),
//		mtime: time.Now(),
//	}
//	return w, nil
//}

func AppendStream(file gridfs.File, v any) (gridfs.File, error) {
	enc, err := ciphertext.EncryptPayload(v)
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
