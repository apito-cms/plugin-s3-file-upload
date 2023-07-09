package main

import (
	"fmt"
	"github.com/apito-cms/buffers/extensions"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
)

// plugin internal state and implementation
type FileUpload struct {
}

// Init system Function Implementation
func (g FileUpload) Init(cred *protobuff.ThirdPartyCredential) error {
	return nil
}

// Migration system Function Implementation
func (g FileUpload) Migration() error {
	fmt.Println(fmt.Sprintf("Running Migration with"))
	return nil
}

// SchemaRegister system Function Implementation
func (g FileUpload) SchemaRegister() (*extensions.ThirdPartyGraphQLSchemas, error) {

	return &extensions.ThirdPartyGraphQLSchemas{
		Queries:   nil,
		Mutations: nil,
	}, nil
}

// RESTApiRegister system Function Implementation
func (g FileUpload) RESTApiRegister() ([]*extensions.ThirdPartyRESTApi, error) {

	fmt.Println(fmt.Sprintf("Running REST Api Register with"))

	return []*extensions.ThirdPartyRESTApi{
		{
			Path:       "/media/upload",
			Method:     "GET",
			Controller: g.ProviderRegister,
		},
	}, nil
}

// system Function Implementation
func (g FileUpload) ProviderRegister(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"message": "provider register",
		"code":    200,
	})
}

// Register system Function Implementation
func (g FileUpload) LoadConfiguration() (map[string]*protobuff.PluginDetails, error) {
	fmt.Println("Running Load Configuration")
	return map[string]*protobuff.PluginDetails{}, nil
}

// Register system Function Implementation
func (g FileUpload) Login(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println(fmt.Sprintf("Running Login"))
	return map[string]interface{}{
		"id_token":      "id token",
		"refresh_token": "refresh token",
	}, nil
}

// Register system Function Implementation
func (g FileUpload) Register(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("Running Email Register")
	return nil, nil
}

// ForgetPassword system Function Implementation
func (g FileUpload) ForgetPassword() {
	fmt.Println("Running Register")
}

// SendEmail system Function Implementation
func (g FileUpload) SendEmail() {
	fmt.Println("Running Register")
}

// SendOTP system Function Implementation
func (g FileUpload) SendOTP() {
	fmt.Println("Running Register")
}

// GetRegisterUser system Function Implementation
func (g FileUpload) GetRegisterUser() {
	fmt.Println("Running Register")
}

// GetLoggedInUser system Function Implementation
func (g FileUpload) GetLoggedInUser(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println(fmt.Sprintf("Running Register"))
	return nil, nil
}

// EmailAuth because plugin Name is email-auth exported
// This exported Name is case-sensitive for the extension to load
var S3FileUpload FileUpload
