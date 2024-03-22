package request

import (
	"go-vue-project/model/common/request"
	"go-vue-project/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
