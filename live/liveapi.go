package live

import (
	"context"

	"github.com/emenwin/danghongyun-api/auth"
	"github.com/emenwin/danghongyun-api/client"
)

const (
	//LiveRestAPIURL rest url
	LiveRestAPIURL = "http://api.danghongyun.com/rest"
	//Version version
	Version = "2.0"
)

// DHLiveManager 提供了对直播进行管理的操作
type DHLiveManager struct {
	Client      *client.Client
	Credentials *auth.Credentials
}

// NewLiveManager 用来构建一个新的直播管理对象
func NewLiveManager(cred *auth.Credentials) *DHLiveManager {

	return NewLiveManagerEx(cred, nil)
}

// NewLiveManagerEx 用来构建一个新的直播管理对象
func NewLiveManagerEx(cred *auth.Credentials, clt *client.Client) *DHLiveManager {

	if clt == nil {
		clt = &client.DefaultClient
	}

	return &DHLiveManager{
		Client:      &client.DefaultClient,
		Credentials: cred,
	}
}

// TemplateCreate 用来修改文件状态, 禁用和启用文件的可访问性
// http://api.danghongyun.com/rest
func (m *DHLiveManager) TemplateCreate(parma Template) (*TemplateRespParam, error) {

	parma.Type = "1"
	action := "liveTemplateCreate"
	queryparam, _, _ := m.Credentials.Sign2(action, Version)

	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, parma)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

// GetTemplates 查询直播模板
// name: 模板名称  可选 （”“）
// ttype : string 直播模版类型。1：用户模版，2：系统模版	可选（""）
// transcodeType	所属直播类型。 0：普通直播、2：VR直播、3：进阶直播 可选(-)
// http://api.danghongyun.com/rest
func (m *DHLiveManager) GetTemplates(name string, ttype string, transcodeType int) (*TemplateListRespParam, error) {

	action := "liveGetTemplates"
	queryparam, _, _ := m.Credentials.Sign2(action, Version)

	params := map[string]interface{}{}
	if name != "" {
		params["name"] = name
	}

	if ttype != "" {
		params["type"] = ttype
	}

	if transcodeType != -1 {
		params["transcodeType"] = transcodeType
	}

	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateListRespParam
	err := m.Client.Call(context.Background(),
		&respTempalte, "GET", url, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}
