package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"wep_app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	switch err {
	case sql.ErrNoRows:
		zap.L().Warn("this is no community in db")
		err = nil
	case nil:
		return
	}
	return
}