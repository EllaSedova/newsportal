// Code generated by zenrpc v2.2.11; DO NOT EDIT.

package vt

import (
	"context"
	"encoding/json"

	"github.com/vmkteam/zenrpc/v2"
	"github.com/vmkteam/zenrpc/v2/smd"
)

var RPC = struct {
	AuthService struct{ Login, Logout, Profile, ChangePassword, VfsAuthToken string }
	UserService struct{ Count, Get, GetByID, Add, Update, Delete, Validate string }
}{
	AuthService: struct{ Login, Logout, Profile, ChangePassword, VfsAuthToken string }{
		Login:          "login",
		Logout:         "logout",
		Profile:        "profile",
		ChangePassword: "changepassword",
		VfsAuthToken:   "vfsauthtoken",
	},
	UserService: struct{ Count, Get, GetByID, Add, Update, Delete, Validate string }{
		Count:    "count",
		Get:      "get",
		GetByID:  "getbyid",
		Add:      "add",
		Update:   "update",
		Delete:   "delete",
		Validate: "validate",
	},
}

func (AuthService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"Login": {
				Description: `Login authenticates user.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "login",
						Description: `User login`,
						Type:        smd.String,
					},
					{
						Name:        "password",
						Description: `User password`,
						Type:        smd.String,
					},
					{
						Name:        "remember",
						Description: `Remember for week`,
						Type:        smd.Boolean,
					},
				},
				Returns: smd.JSONSchema{
					Description: `User authentication key`,
					Type:        smd.String,
				},
				Errors: map[int]string{
					400: "Invalid login or password",
					500: "Internal Error",
				},
			},
			"Logout": {
				Description: `Logout current user from every session`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `Successful logout`,
					Type:        smd.Boolean,
				},
				Errors: map[int]string{
					401: "Invalid authentication credentials",
					500: "Internal Error",
				},
			},
			"Profile": {
				Description: `Profile is a function that returns current user profile`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `UserProfile`,
					Optional:    true,
					Type:        smd.Object,
					TypeName:    "UserProfile",
					Properties: smd.PropertyList{
						{
							Name: "id",
							Type: smd.Integer,
						},
						{
							Name: "createdAt",
							Type: smd.String,
						},
						{
							Name: "login",
							Type: smd.String,
						},
						{
							Name:     "lastActivityAt",
							Optional: true,
							Type:     smd.String,
						},
						{
							Name: "statusId",
							Type: smd.Integer,
						},
					},
				},
				Errors: map[int]string{
					401: "Invalid authentication credentials",
				},
			},
			"ChangePassword": {
				Description: `ChangePassword changes current user password.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "password",
						Description: `New user password`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `New user authentication key`,
					Type:        smd.String,
				},
				Errors: map[int]string{
					401: "Invalid authentication credentials",
					500: "Internal Error",
				},
			},
			"VfsAuthToken": {
				Description: `VfsAuthToken get auth token for VFS requests`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Type: smd.String,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s AuthService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.AuthService.Login:
		var args = struct {
			Login    string `json:"login"`
			Password string `json:"password"`
			Remember bool   `json:"remember"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"login", "password", "remember"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Login(ctx, args.Login, args.Password, args.Remember))

	case RPC.AuthService.Logout:
		resp.Set(s.Logout(ctx))

	case RPC.AuthService.Profile:
		resp.Set(s.Profile(ctx))

	case RPC.AuthService.ChangePassword:
		var args = struct {
			Password string `json:"password"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"password"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.ChangePassword(ctx, args.Password))

	case RPC.AuthService.VfsAuthToken:
		resp.Set(s.VfsAuthToken(ctx))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (UserService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"Count": {
				Description: `Count Users according to conditions in search params`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "search",
						Optional:    true,
						Description: `UserSearch`,
						Type:        smd.Object,
						TypeName:    "UserSearch",
						Properties: smd.PropertyList{
							{
								Name:     "id",
								Optional: true,
								Type:     smd.Integer,
							},
							{
								Name:     "login",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name:     "statusId",
								Optional: true,
								Type:     smd.Integer,
							},
							{
								Name:     "lastActivityAtFrom",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name:     "lastActivityAtTo",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name: "ids",
								Type: smd.Array,
								Items: map[string]string{
									"type": smd.Integer,
								},
							},
							{
								Name:     "notId",
								Optional: true,
								Type:     smd.Integer,
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Description: `int`,
					Type:        smd.Integer,
				},
				Errors: map[int]string{
					500: "Internal Error",
				},
			},
			"Get": {
				Description: `Get а list of Users according to conditions in search params`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "search",
						Optional:    true,
						Description: `UserSearch`,
						Type:        smd.Object,
						TypeName:    "UserSearch",
						Properties: smd.PropertyList{
							{
								Name:     "id",
								Optional: true,
								Type:     smd.Integer,
							},
							{
								Name:     "login",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name:     "statusId",
								Optional: true,
								Type:     smd.Integer,
							},
							{
								Name:     "lastActivityAtFrom",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name:     "lastActivityAtTo",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name: "ids",
								Type: smd.Array,
								Items: map[string]string{
									"type": smd.Integer,
								},
							},
							{
								Name:     "notId",
								Optional: true,
								Type:     smd.Integer,
							},
						},
					},
					{
						Name:        "viewOps",
						Optional:    true,
						Description: `ViewOps`,
						Type:        smd.Object,
						TypeName:    "ViewOps",
						Properties: smd.PropertyList{
							{
								Name:        "page",
								Description: `page number, default - 1`,
								Type:        smd.Integer,
							},
							{
								Name:        "pageSize",
								Description: `items count per page, max - 500`,
								Type:        smd.Integer,
							},
							{
								Name:        "sortColumn",
								Description: `sort by column name`,
								Type:        smd.String,
							},
							{
								Name:        "sortDesc",
								Description: `descending sort`,
								Type:        smd.Boolean,
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Description: `[]UserSummary`,
					Type:        smd.Array,
					TypeName:    "[]UserSummary",
					Items: map[string]string{
						"$ref": "#/definitions/UserSummary",
					},
					Definitions: map[string]smd.Definition{
						"UserSummary": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "createdAt",
									Type: smd.String,
								},
								{
									Name: "login",
									Type: smd.String,
								},
								{
									Name:     "lastActivityAt",
									Optional: true,
									Type:     smd.String,
								},
								{
									Name:     "status",
									Optional: true,
									Ref:      "#/definitions/Status",
									Type:     smd.Object,
								},
							},
						},
						"Status": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "alias",
									Type: smd.String,
								},
								{
									Name: "title",
									Type: smd.String,
								},
							},
						},
					},
				},
				Errors: map[int]string{
					500: "Internal Error",
				},
			},
			"GetByID": {
				Description: `GetByID returns a User by its ID.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "id",
						Description: `int`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: `User`,
					Optional:    true,
					Type:        smd.Object,
					TypeName:    "User",
					Properties: smd.PropertyList{
						{
							Name: "id",
							Type: smd.Integer,
						},
						{
							Name: "createdAt",
							Type: smd.String,
						},
						{
							Name: "login",
							Type: smd.String,
						},
						{
							Name: "password",
							Type: smd.String,
						},
						{
							Name:     "lastActivityAt",
							Optional: true,
							Type:     smd.String,
						},
						{
							Name: "statusId",
							Type: smd.Integer,
						},
						{
							Name:     "status",
							Optional: true,
							Ref:      "#/definitions/Status",
							Type:     smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"Status": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "alias",
									Type: smd.String,
								},
								{
									Name: "title",
									Type: smd.String,
								},
							},
						},
					},
				},
				Errors: map[int]string{
					500: "Internal Error",
					404: "Not Found",
				},
			},
			"Add": {
				Description: `Add a User from the query`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "user",
						Description: `User`,
						Type:        smd.Object,
						TypeName:    "User",
						Properties: smd.PropertyList{
							{
								Name: "id",
								Type: smd.Integer,
							},
							{
								Name: "createdAt",
								Type: smd.String,
							},
							{
								Name: "login",
								Type: smd.String,
							},
							{
								Name: "password",
								Type: smd.String,
							},
							{
								Name:     "lastActivityAt",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name: "statusId",
								Type: smd.Integer,
							},
							{
								Name:     "status",
								Optional: true,
								Ref:      "#/definitions/Status",
								Type:     smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"Status": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "id",
										Type: smd.Integer,
									},
									{
										Name: "alias",
										Type: smd.String,
									},
									{
										Name: "title",
										Type: smd.String,
									},
								},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Description: `User`,
					Optional:    true,
					Type:        smd.Object,
					TypeName:    "User",
					Properties: smd.PropertyList{
						{
							Name: "id",
							Type: smd.Integer,
						},
						{
							Name: "createdAt",
							Type: smd.String,
						},
						{
							Name: "login",
							Type: smd.String,
						},
						{
							Name: "password",
							Type: smd.String,
						},
						{
							Name:     "lastActivityAt",
							Optional: true,
							Type:     smd.String,
						},
						{
							Name: "statusId",
							Type: smd.Integer,
						},
						{
							Name:     "status",
							Optional: true,
							Ref:      "#/definitions/Status",
							Type:     smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"Status": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "alias",
									Type: smd.String,
								},
								{
									Name: "title",
									Type: smd.String,
								},
							},
						},
					},
				},
				Errors: map[int]string{
					500: "Internal Error",
					400: "Validation Error",
				},
			},
			"Update": {
				Description: `Update updates the User data identified by id from the query`,
				Parameters: []smd.JSONSchema{
					{
						Name:     "user",
						Type:     smd.Object,
						TypeName: "User",
						Properties: smd.PropertyList{
							{
								Name: "id",
								Type: smd.Integer,
							},
							{
								Name: "createdAt",
								Type: smd.String,
							},
							{
								Name: "login",
								Type: smd.String,
							},
							{
								Name: "password",
								Type: smd.String,
							},
							{
								Name:     "lastActivityAt",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name: "statusId",
								Type: smd.Integer,
							},
							{
								Name:     "status",
								Optional: true,
								Ref:      "#/definitions/Status",
								Type:     smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"Status": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "id",
										Type: smd.Integer,
									},
									{
										Name: "alias",
										Type: smd.String,
									},
									{
										Name: "title",
										Type: smd.String,
									},
								},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Description: `User`,
					Type:        smd.Boolean,
					TypeName:    "User",
				},
				Errors: map[int]string{
					500: "Internal Error",
					400: "Validation Error",
					404: "Not Found",
				},
			},
			"Delete": {
				Description: `Delete deletes the User by its ID.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "id",
						Description: `int`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: `isDeleted`,
					Type:        smd.Boolean,
				},
				Errors: map[int]string{
					500: "Internal Error",
					400: "Validation Error",
					404: "Not Found",
				},
			},
			"Validate": {
				Description: `Validate Verifies that User data is valid.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "user",
						Description: `User`,
						Type:        smd.Object,
						TypeName:    "User",
						Properties: smd.PropertyList{
							{
								Name: "id",
								Type: smd.Integer,
							},
							{
								Name: "createdAt",
								Type: smd.String,
							},
							{
								Name: "login",
								Type: smd.String,
							},
							{
								Name: "password",
								Type: smd.String,
							},
							{
								Name:     "lastActivityAt",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name: "statusId",
								Type: smd.Integer,
							},
							{
								Name:     "status",
								Optional: true,
								Ref:      "#/definitions/Status",
								Type:     smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"Status": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "id",
										Type: smd.Integer,
									},
									{
										Name: "alias",
										Type: smd.String,
									},
									{
										Name: "title",
										Type: smd.String,
									},
								},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Description: `[]FieldError`,
					Type:        smd.Array,
					TypeName:    "[]FieldError",
					Items: map[string]string{
						"$ref": "#/definitions/FieldError",
					},
					Definitions: map[string]smd.Definition{
						"FieldError": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "field",
									Type: smd.String,
								},
								{
									Name: "error",
									Type: smd.String,
								},
								{
									Name:        "constraint",
									Optional:    true,
									Description: `Help with generating an error message.`,
									Ref:         "#/definitions/FieldErrorConstraint",
									Type:        smd.Object,
								},
							},
						},
						"FieldErrorConstraint": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name:        "max",
									Description: `Max value for field.`,
									Type:        smd.Integer,
								},
								{
									Name:        "min",
									Description: `Min value for field.`,
									Type:        smd.Integer,
								},
							},
						},
					},
				},
				Errors: map[int]string{
					500: "Internal Error",
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s UserService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.UserService.Count:
		var args = struct {
			Search *UserSearch `json:"search"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"search"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Count(ctx, args.Search))

	case RPC.UserService.Get:
		var args = struct {
			Search  *UserSearch `json:"search"`
			ViewOps *ViewOps    `json:"viewOps"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"search", "viewOps"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Get(ctx, args.Search, args.ViewOps))

	case RPC.UserService.GetByID:
		var args = struct {
			Id int `json:"id"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"id"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.GetByID(ctx, args.Id))

	case RPC.UserService.Add:
		var args = struct {
			User User `json:"user"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"user"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Add(ctx, args.User))

	case RPC.UserService.Update:
		var args = struct {
			User User `json:"user"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"user"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Update(ctx, args.User))

	case RPC.UserService.Delete:
		var args = struct {
			Id int `json:"id"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"id"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Delete(ctx, args.Id))

	case RPC.UserService.Validate:
		var args = struct {
			User User `json:"user"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"user"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Validate(ctx, args.User))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
