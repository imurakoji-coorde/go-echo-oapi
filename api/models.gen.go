// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

// NewUser defines model for NewUser.
type NewUser struct {
	Age  float32 `form:"age" json:"age"`
	Name string  `form:"name" json:"name"`
}

// User defines model for user.
type User struct {
	Age  float32 `form:"age" json:"age"`
	Id   int64   `json:"id"`
	Name string  `form:"name" json:"name"`
}

// ListUsersParams defines parameters for ListUsers.
type ListUsersParams struct {
	// limit
	Limit int32 `form:"limit" json:"limit"`

	// page
	Page int32 `form:"page" json:"page"`
}

// CreateUsersJSONBody defines parameters for CreateUsers.
type CreateUsersJSONBody = NewUser

// CreateUsersJSONRequestBody defines body for CreateUsers for application/json ContentType.
type CreateUsersJSONRequestBody = CreateUsersJSONBody
