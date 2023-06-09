package stegano

import (
	"bytes"
	"io"
	"time"

	"github.com/vela-ssoc/vela-common-mb/dal/gridfs"
	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mba/encipher"
	"github.com/vela-ssoc/vela-common-mba/steganography"
)

func MinionAppend(file gridfs.File, brk *model.Broker, bin *model.MinionBin, tags []string) (gridfs.File, error) {
	hide := &steganography.MinionHide{
		Servername: brk.Servername,
		LAN:        brk.LAN,
		VIP:        brk.VIP,
		Edition:    string(bin.Semver),
		Hash:       file.MD5(),
		Size:       file.Size(),
		DownloadAt: time.Now(),
		Tags:       tags,
	}
	enc, err := encipher.EncryptJSON(hide)
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
