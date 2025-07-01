package rds

import (
	"small_demo_go/system"
	"small_demo_go/model"
)

func (cli *client) GetMessage()([]*model.Message, error) {
	query := cli.db
	query = query.Select("messages.*")
	query = query.Where("status = ?", system.StatusProcessing)
	query = query.Where("delete_flg = ?", system.DeleteFlagNotDeleted)

	res := []*model.Message{}
	err := query.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}