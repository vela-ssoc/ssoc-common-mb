package model

import (
	"io/fs"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ZipFile struct {
	ID          bson.ObjectID `bson:"_id,omitempty"        json:"id"`                // 数据库 ID
	ArchiveID   bson.ObjectID `bson:"archive_id"           json:"archive_id"`        // 关联的 ZipArchive._id
	GridID      bson.ObjectID `bson:"grid_id,omitempty"    json:"grid_id"`           // 关联 GridFS 文件 ID
	Directory   string        `bson:"directory"            json:"directory"`         // 目录，根目录为空，例如：webapp/static
	FileName    string        `bson:"file_name"            json:"file_name"`         // 文件名，例如：index.html
	FullName    string        `bson:"full_name"            json:"full_name"`         // 文件全路径，例如：webapp/static/index.html
	FileSize    int64         `bson:"file_size"            json:"file_size"`         // 文件大小
	FileMode    uint32        `bson:"file_mode"            json:"file_mode"`         // 文件 Mode
	MIMEType    string        `bson:"mime_type,omitempty"  json:"mime_type"`         // MIME 例如：text/html
	IsDirectory bool          `bson:"is_directory"         json:"is_directory"`      // 是否目录节点
	Version     int64         `bson:"version,omitempty"    json:"version"`           // 版本计数器：每次子文件增删改时自增
	Checksum    Checksum      `bson:"checksum,omitempty"   json:"checksum,omitzero"` // 文件摘要码
	CreatedAt   time.Time     `bson:"created_at,omitempty" json:"created_at"`        // 创建时间
	UpdatedAt   time.Time     `bson:"updated_at,omitempty" json:"updated_at"`        // 最近一次修改时间
}

func (f ZipFile) Type() fs.FileMode {
	return f.Mode()
}

func (f ZipFile) Info() (fs.FileInfo, error) {
	return f, nil
}

func (f ZipFile) Name() string {
	return f.FileName
}

func (f ZipFile) Size() int64 {
	return f.FileSize
}

func (f ZipFile) Mode() fs.FileMode {
	if f.FileMode != 0 {
		return fs.FileMode(f.FileMode)
	}

	if f.IsDir() {
		return fs.ModeDir | 0755 // 目录：带有 ModeDir 标志和 755 权限
	}

	return 0644 // 普通文件：644 权限
}

func (f ZipFile) ModTime() time.Time {
	return f.UpdatedAt
}

func (f ZipFile) IsDir() bool {
	return f.IsDirectory
}

func (f ZipFile) Sys() any {
	return nil
}

func (ZipFile) CollectionInfo() (string, string) { return "zip_file", "ZIP 文件信息" }

type ZipFileMetadata struct {
	ArchiveID bson.ObjectID `bson:"archive_id" json:"archive_id"`
	Version   int64         `bson:"version"    json:"version"`
	Checksum  Checksum      `bson:"checksum"   json:"checksum"`
}
