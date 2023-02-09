package studentstorage

import (
	"context"
	studentmodel "first-apis/module/student/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int)  error {
	if err := s.db.Table(studentmodel.Student{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
    }	
	return nil
}