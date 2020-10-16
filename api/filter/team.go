package filter

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
)

type Team struct {
	service *service.Team
}

func NewTeam() *Team {
	return &Team{
		service: service.NewTeam(),
	}
}

// 创建团队
func (f *Team) Create(c *gin.Context) (*model.TeamInfo, *api.Error) {
	params := &request.TeamCreateParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	operator := &model.Operator{
		Id:      c.GetInt64("member_id"),
		Account: c.GetString("account"),
	}

	// 调用service对应的方法
	teamInfo, err := f.service.CreateTeam(params, operator)

	return teamInfo, err
}
