package studentstorage


import (
	"context"
	studentmodel "first-apis/module/student/model"
	
)

func (s *sqlStore) List(ctx context.Context, paging studentmodel.Paging)  ([]studentmodel.Student, error) {
	var data []studentmodel.Student
	
	if err := s.db.
	Offset((paging.Page - 1)*paging.Limit).
	Order("id desc").
	Limit(paging.Limit).
	Find(&data).Error; err != nil {
		return nil,err
	}
	
	return data,nil
}