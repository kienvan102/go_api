package studentstorage

import (
	"context"
	studentmodel "first-apis/module/student/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *studentmodel.StudentUpdate)  error{
	if err := s.db.Table(studentmodel.Student{}.TableName()).
	Where("id = ?", id).
	Updates(&data).Error; err != nil {
		return err
    }	
	return  nil
}