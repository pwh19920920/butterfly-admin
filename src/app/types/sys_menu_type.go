package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pwh19920920/butterfly-admin/src/app/domain/entity"
	"github.com/pwh19920920/butterfly/response"
)

type SysMenuQueryRequest struct {
	response.RequestPaging
}

type SysMenuTreeResponse struct {
	entity.SysMenu
	Options  []entity.SysMenuOption `json:"options"`
	Children []SysMenuTreeResponse  `json:"children"`
}

type SysMenuCreateRequest struct {
	entity.SysMenu
	Options []entity.SysMenuOption `json:"options"`
}

func (req SysMenuCreateRequest) ValidateForCreate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Code, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Name, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Path, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Icon, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Sort, validation.Required, validation.Max(9999)),
		validation.Field(&req.Parent, validation.Required, validation.Min(0)),
		validation.Field(&req.Route, validation.Required, validation.Length(0, 255)),
	)
}

func (req SysMenuCreateRequest) ValidateForModify() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Code, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Name, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Path, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Icon, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Sort, validation.Required, validation.Max(9999)),
		validation.Field(&req.Parent, validation.Required, validation.Min(0)),
		validation.Field(&req.Route, validation.Required, validation.Length(0, 255)),
		validation.Field(&req.Id, validation.Required),
	)
}
