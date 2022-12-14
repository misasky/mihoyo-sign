// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name   string `path:"name,optional" validate:"oneof=你 我"`
	Age    int    `form:"age,optional" validate:"omitempty,isValidAge"`
	Hobby  string `form:"hobby,optional" validate:"isValidHobby,oneof=ball swim"`
	Height int    `form:"height,optional" validate:"max=100,min=10"`
}

type Response struct {
	Message string `json:"message"`
}
