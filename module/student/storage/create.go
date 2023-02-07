package studentstorage

import (
	"context"
	studentmodel "first-apis/module/student/model"
)

func (s *sqlStore) Create(ctx context.Context, data *studentmodel.Student) error{
	if err := s.db.Create(&data).Error; err != nil{
		return err
	}
	return nil
}