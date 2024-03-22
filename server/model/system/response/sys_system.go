package response

import "go-vue-project/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
