package utils

import (
	"reflect"
)

func UpdateModelValuesFromDTO(model interface{}, dto interface{}) {
	reflectedDTO := reflect.ValueOf(dto).Elem()
	reflectedModel := reflect.ValueOf(model).Elem()

	for i := 0; i < reflectedDTO.NumField(); i++ {
		if !reflectedDTO.Field(i).IsNil() {
			fieldName := reflectedDTO.Type().Field(i).Name
			fieldValue := reflectedDTO.Field(i).Elem()
			modelField := reflectedModel.FieldByName(fieldName)

			if modelField.IsValid() {
				if modelField.CanSet() {
					modelField.Set(fieldValue)
				}
			}

		}
	}
}
