package handler

import (
	"fmt"
	"strings"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateCard

type CreateCardRequest struct {
	Name        string   `json:"name"`
	Cost        *int64   `json:"cost"`
	Power       *int64   `json:"power"`
	Description string   `json:"description"`
	Source      string   `json:"source"`
	Image       string   `json:"image"`
	Tags        []string `json:"tags"`
}

func (r *CreateCardRequest) Validate() error {
	// Empty or malformed body
	if r.Name == "" && r.Cost == nil && r.Power == nil && r.Description == "" && r.Source == "" && r.Image == "" {
		return fmt.Errorf("empty or malformed body")
	}
	// Single fields
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Cost == nil {
		return errParamIsRequired("cost", "int64")
	}
	if r.Power == nil {
		return errParamIsRequired("power", "int64")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Source == "" {
		return errParamIsRequired("source", "string")
	}
	if r.Image == "" {
		return errParamIsRequired("image", "string")
	}
	if r.Tags == nil {
		return errParamIsRequired("tags", "string[]")
	} else {
		// Check if any tag is null
		for i := 0; i < len(r.Tags); i++ {
			if r.Tags[i] == "" {
				return errParamIsRequired("tag", "string")
			}
			if strings.Contains(r.Tags[i], ",") {
				return fmt.Errorf("don't use commas on tags")
			}
		}
	}
	return nil
}

// UpdateCard

type UpdateCardRequest struct {
	Name        string `json:"name"`
	Cost        *int64 `json:"cost"`
	Power       *int64 `json:"power"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Image       string `json:"image"`
}

func (r *UpdateCardRequest) Validate() error {
	// If any field is provided validation is truthy
	if r.Name != "" || r.Cost != nil || r.Power != nil || r.Description != "" || r.Source != "" || r.Image != "" {
		return nil
	}
	return fmt.Errorf("at least one parameter is needed or malformed body")
}
