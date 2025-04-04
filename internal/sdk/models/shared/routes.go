// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type Groups struct {
	Name string `json:"name"`
	// Short description of this group
	Description *string `json:"description,omitempty"`
	// Whether this group is disabled
	Disabled *bool `json:"disabled,omitempty"`
}

func (o *Groups) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Groups) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Groups) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

type Comments struct {
	// Optional, short description of this Route's purpose
	Comment              *string `json:"comment,omitempty"`
	AdditionalProperties any     `additionalProperties:"true" json:"-"`
}

func (c Comments) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Comments) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Comments) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *Comments) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

type Routes struct {
	// Routes ID
	ID *string `json:"id,omitempty"`
	// Pipeline routing rules
	Routes []RoutesRoute     `json:"routes"`
	Groups map[string]Groups `json:"groups,omitempty"`
	// Comments
	Comments []Comments `json:"comments,omitempty"`
}

func (o *Routes) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Routes) GetRoutes() []RoutesRoute {
	if o == nil {
		return []RoutesRoute{}
	}
	return o.Routes
}

func (o *Routes) GetGroups() map[string]Groups {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *Routes) GetComments() []Comments {
	if o == nil {
		return nil
	}
	return o.Comments
}

type RoutesInput struct {
	// Routes ID
	ID *string `json:"id,omitempty"`
	// Pipeline routing rules
	Routes []RoutesRouteInput `json:"routes"`
	Groups map[string]Groups  `json:"groups,omitempty"`
	// Comments
	Comments []Comments `json:"comments,omitempty"`
}

func (o *RoutesInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RoutesInput) GetRoutes() []RoutesRouteInput {
	if o == nil {
		return []RoutesRouteInput{}
	}
	return o.Routes
}

func (o *RoutesInput) GetGroups() map[string]Groups {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *RoutesInput) GetComments() []Comments {
	if o == nil {
		return nil
	}
	return o.Comments
}
