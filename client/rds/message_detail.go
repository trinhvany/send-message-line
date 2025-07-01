package rds

import (
	"small_demo_go/system"
	"small_demo_go/model"
)

func (cli *client) GetMessageDetail(MessageID int64)([]*model.MessageDetail, error) {
	query := cli.db
	query = query.Select("message_detail.*")
	query = query.Where("delete_flg = ?", system.DeleteFlagNotDeleted)
	query = query.Where("message_id = ?", MessageID)

	res := []*model.MessageDetail{}
	err := query.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}