package pb

import (
	"io"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
	"github.com/xuender/oils/logs"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	_headSize = 265
)

// nolint: gochecknoglobals
var caser = cases.Title(language.English)

func GetMetaByReader(readCloser io.ReadCloser) (Meta, error) {
	defer readCloser.Close()

	meta := Meta_Unknown
	head := make([]byte, _headSize)

	if _, err := readCloser.Read(head); err != nil {
		return meta, err
	}

	kind, err := filetype.Match(head)
	if err != nil {
		return meta, err
	}

	logs.Debugw("GetMeta", "kind", kind)

	if filetype.IsArchive(head) {
		return Meta_Archive, nil
	}

	if filetype.IsImage(head) {
		return Meta_Image, nil
	}

	if value, has := Meta_value[caser.String(kind.Extension)]; has {
		meta = Meta(value)
	}

	return meta, nil
}

func GetMetaByExt(path string) Meta {
	kind := filetype.GetType(filepath.Ext(path))

	if value, has := Meta_value[caser.String(kind.MIME.Type)]; has {
		return Meta(value)
	}

	return Meta_Unknown
}

func GetMeta(path string) (Meta, error) {
	file, err := os.Open(path)
	if err != nil {
		return Meta_Unknown, err
	}

	meta, err := GetMetaByReader(file)
	if err != nil {
		return meta, err
	}

	if meta == Meta_Unknown {
		meta = GetMetaByExt(path)
	}

	return meta, nil
}

// nolint: gochecknoinits
func init() {
	filetype.AddType(".go", "Golang")
	filetype.AddType(".java", "Java")
	filetype.AddType(".js", "JavaScript")
}