package studentbiz 

import (
	"context"
	studentmodel "first-apis/module/student/model"
)

type GetStudentStore interface {
	FindDataWithCondition(ctx context.Context, 
		condition map[string]interface{},
		moreKeys ...string,
		) (*studentmodel.Student, error)
}

type getStudentBiz struct{
	store GetStudentStore
}

func NewGetStudentBiz(store GetStudentStore) *getStudentBiz {
	return &getStudentBiz{store: store}
}

func (biz *getStudentBiz) GetStudent(ctx context.Context, id int) (*studentmodel.Student, error) {

	Student, err := biz.store.FindDataWithCondition(ctx,map[string]interface{}{"id": id})
	
	if err!= nil {
        return nil, err
    }


	return Student,nil
}