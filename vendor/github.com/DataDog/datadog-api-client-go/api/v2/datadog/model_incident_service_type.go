/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// IncidentServiceType Incident Service resource type.
type IncidentServiceType string

// List of IncidentServiceType
const (
	INCIDENTSERVICETYPE_SERVICES IncidentServiceType = "services"
)

func (v *IncidentServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IncidentServiceType(value)
	for _, existing := range []IncidentServiceType{"services"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IncidentServiceType", value)
}

// Ptr returns reference to IncidentServiceType value
func (v IncidentServiceType) Ptr() *IncidentServiceType {
	return &v
}

type NullableIncidentServiceType struct {
	value *IncidentServiceType
	isSet bool
}

func (v NullableIncidentServiceType) Get() *IncidentServiceType {
	return v.value
}

func (v *NullableIncidentServiceType) Set(val *IncidentServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentServiceType(val *IncidentServiceType) *NullableIncidentServiceType {
	return &NullableIncidentServiceType{value: val, isSet: true}
}

func (v NullableIncidentServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
