package models

import "gorm.io/gorm"

type FailedJob struct {
	gorm.Model
	Status    string `json:"status" gorm:"not null"`
	Payload   string `json:"payload" gorm:"not null"`
	QueueName string `json:"queue_name" gorm:"not null"`
}

func (j *FailedJob) Save() error {
	return db.Create(&j).Error
}

func (j *FailedJob) Failed() ([]*FailedJob, error) {
	var jobs []*FailedJob
	res := db.Group("queue_name").Find(&jobs)
	return jobs, res.Error
}
