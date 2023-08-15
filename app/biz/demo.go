package biz

import (
	"framework/ay"
	"framework/dal/model"
	"framework/dal/query"
)

type demo struct {
	*base
}

func (d demo) Get(id int64) (res *model.Demo, err error) {
	qe := query.Use(ay.Db).Demo
	if res, err = qe.WithContext(d.ctx).Where(qe.Id.Eq(id)).Take(); err != nil {
		return
	}
	return
}
