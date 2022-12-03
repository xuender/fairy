package meta

import (
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/xuender/fairy/pb"
)

// nolint: gochecknoglobals
var _formats = [...][2]string{
	{"\\$[yY]{4}", "2006"},
	{"\\$[yY]{2}", "06"},
	{"\\$[mM]{2}", "01"},
	{"\\$[mM]{1}", "1"},
	{"\\$[dD]{2}", "02"},
	{"\\$[dD]{1}", "2"},
}

// Info 文件信息.
type Info struct {
	Path  string
	Meta  pb.Meta
	Date  time.Time
	Error error
}

func NewInfoError(path string, err error) *Info {
	return &Info{
		Path:  path,
		Error: err,
	}
}

// Target 根据格式生成目的目录.
func (p Info) Target(format string) string {
	for _, str := range _formats {
		format = regexp.MustCompile(str[0]).ReplaceAllString(format, str[1])
	}

	return p.Date.Format(format)
}

func (p Info) String() string {
	return fmt.Sprintf("%s: %v %s %v", p.Path, p.Meta, p.Date.Format("2006-01-02 15:04:05"), p.Error)
}

// Output 输出.
func (p Info) Output(writer io.Writer) {
	fmt.Fprintln(writer, p.Path)

	if p.Error == nil {
		fmt.Fprintf(writer, "类型: %v\n", p.Meta)
		fmt.Fprintf(writer, "时间: %s\n", p.Date.Format("2006-01-02 15:04:05"))

		return
	}

	fmt.Fprintf(writer, "错误: %v\n", p.Error)
}
