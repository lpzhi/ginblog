package models

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTagTotal() (count int)  {
	db.Model(&Tag{}).Count(&count)
	return
}
