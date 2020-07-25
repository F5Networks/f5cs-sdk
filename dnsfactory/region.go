package factory

type Region struct {
	// Field 0
	DisplayName *string `draft_validate:"required,max=64" json:"display_name,omitempty"`

	// Field 1
	Remark *string `draft_validate:"omitempty,max=128" json:"remark,omitempty"`

	// Field 2
	Sectors []Sector `draft_validate:"required,gte=1,lte=100,unique,dive" json:"sectors,omitempty"`
}

type Sector struct {
	// Field 0
	Code string `draft_validate:"required,min=1,max=50" json:"code,omitempty"`

	// Field 1
	Scale string `draft_validate:"required,oneof=continent country state" json:"scale,omitempty"`
}
