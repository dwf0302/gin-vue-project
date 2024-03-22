package request

import (
	"go-vue-project/model/common/request"
	"go-vue-project/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
