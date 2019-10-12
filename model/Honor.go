package model

import "time"

type Honor struct {
	Id string `json:"id"`
	StudentId string `json:"studentId"`
	StudentName string `json:"studentName"`
	HonorName string `json:"honorName"`
	AcquisitionTime  time.Time `json:"acquisitionTime"`
	DefaultModel
}

