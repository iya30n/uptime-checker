package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Status    string `json:"status" gorm:"not null"`
	Payload   string `json:"payload" gorm:"not null"`
	QueueName string `json:"queue_name" gorm:"not null"`
}

func (j *Job) Save() error {
	return db.Create(&j).Error
}
