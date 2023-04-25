package validator

import (
	"api-go/model"
	validator "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidator(task model.Task) error
}

type taskValidator struct {}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidator(task model.Task) error {
	return validator.ValidateStruct(&task,
		validator.Field(&task.Title,validator.Required.Error("title is required"),
		validator.RuneLength(1,10).Error("limitted 1-10 characters")),
	)
}