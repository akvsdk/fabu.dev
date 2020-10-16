package service

import (
	"fabu.dev/api/model"
)

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

func (s *Member) GetMemberInfo(memberId uint64) (*model.Member, error) {
	member := model.NewMember()

	err := member.Db().Where("id = ?", memberId).First(member).Error
	if err != nil {
		return nil, err
	}

	return member, nil
}