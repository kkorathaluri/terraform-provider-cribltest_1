// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type HTTPOutAuthConfig struct {
	Disabled *bool   `json:"disabled,omitempty"`
	Password *string `json:"password,omitempty"`
	Token    *string `json:"token,omitempty"`
	Username *string `json:"username,omitempty"`
}

func (o *HTTPOutAuthConfig) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *HTTPOutAuthConfig) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *HTTPOutAuthConfig) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *HTTPOutAuthConfig) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
