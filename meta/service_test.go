package meta_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/fairy/meta"
	"github.com/xuender/fairy/pb"
)

func TestService_Info(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	service := meta.NewService()
	info := service.Info("service.go")

	ass.Nil(info.Error)
	ass.Equal(pb.Meta_Golang, info.Meta)
}
