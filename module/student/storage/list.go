package studentstorage


import (
	"context"
	studentmodel "first-apis/module/student/model"
	
)

func (s *sqlStore) List(ctx context.Context, paging studentmodel.Paging)  ([]studentmodel.Student, error) {
	var data []studentmodel.Student
	if paging.Page < 0{
		paging.Page = 1
	}
	if paging.Limit < 0{
        paging.Limit = 5
    }
	s.db.
	Offset((paging.Page - 1)*paging.Limit).
	Order("id desc").
	Limit(paging.Limit).
	Find(&data)
	
	return data,nil
}