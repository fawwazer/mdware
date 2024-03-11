package task

import (
	"errors"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Kegiatan string
	Pemilik  string `gorm:"type:varchar(13);"`
}

type TaskModel struct {
	Connection *gorm.DB
}

func (tm *TaskModel) Insert(kegiatanBaru Task) (Task, error) {
	if err := tm.Connection.Create(&kegiatanBaru).Error; err != nil {
		return Task{}, err
	}

	return kegiatanBaru, nil
}

func (tm *TaskModel) ListKegiatan(pemilik string) ([]Task, error) {
	var result []Task
	if err := tm.Connection.Where("pemilik = ?", pemilik).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (tm *TaskModel) UpdateKegiatan(pemilik string, todoID uint, data Task) (Task, error) {
	var qry = tm.Connection.Where("pemilik = ? AND id = ?", pemilik, todoID).Updates(data)
	if err := qry.Error; err != nil {
		return Task{}, err
	}

	if qry.RowsAffected < 1 {
		return Task{}, errors.New("no data affected")
	}

	return data, nil
}
