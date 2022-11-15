/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// UserGroup struct for UserGroup
type UserGroup struct {
	Id                   *int32   `json:"id,omitempty"`
	BsdgrpGid            *int32   `json:"bsdgrp_gid,omitempty"`
	BsdgrpGroup          *string  `json:"bsdgrp_group,omitempty"`
	BsdgrpBuiltin        *bool    `json:"bsdgrp_builtin,omitempty"`
	BsdgrpSudo           *bool    `json:"bsdgrp_sudo,omitempty"`
	BsdgrpSudoNopasswd   *bool    `json:"bsdgrp_sudo_nopasswd,omitempty"`
	BsdgrpSudoCommands   []string `json:"bsdgrp_sudo_commands,omitempty"`
	BsdgrpSmb            *bool    `json:"bsdgrp_smb,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserGroup UserGroup

// NewUserGroup instantiates a new UserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroup() *UserGroup {
	this := UserGroup{}
	return &this
}

// NewUserGroupWithDefaults instantiates a new UserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupWithDefaults() *UserGroup {
	this := UserGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserGroup) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserGroup) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserGroup) SetId(v int32) {
	o.Id = &v
}

// GetBsdgrpGid returns the BsdgrpGid field value if set, zero value otherwise.
func (o *UserGroup) GetBsdgrpGid() int32 {
	if o == nil || isNil(o.BsdgrpGid) {
		var ret int32
		return ret
	}
	return *o.BsdgrpGid
}

// GetBsdgrpGidOk returns a tuple with the BsdgrpGid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetBsdgrpGidOk() (*int32, bool) {
	if o == nil || isNil(o.BsdgrpGid) {
		return nil, false
	}
	return o.BsdgrpGid, true
}

// HasBsdgrpGid returns a boolean if a field has been set.
func (o *UserGroup) HasBsdgrpGid() bool {
	if o != nil && !isNil(o.BsdgrpGid) {
		return true
	}

	return false
}

// SetBsdgrpGid gets a reference to the given int32 and assigns it to the BsdgrpGid field.
func (o *UserGroup) SetBsdgrpGid(v int32) {
	o.BsdgrpGid = &v
}

// GetBsdgrpGroup returns the BsdgrpGroup field value if set, zero value otherwise.
func (o *UserGroup) GetBsdgrpGroup() string {
	if o == nil || isNil(o.BsdgrpGroup) {
		var ret string
		return ret
	}
	return *o.BsdgrpGroup
}

// GetBsdgrpGroupOk returns a tuple with the BsdgrpGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetBsdgrpGroupOk() (*string, bool) {
	if o == nil || isNil(o.BsdgrpGroup) {
		return nil, false
	}
	return o.BsdgrpGroup, true
}

// HasBsdgrpGroup returns a boolean if a field has been set.
func (o *UserGroup) HasBsdgrpGroup() bool {
	if o != nil && !isNil(o.BsdgrpGroup) {
		return true
	}

	return false
}

// SetBsdgrpGroup gets a reference to the given string and assigns it to the BsdgrpGroup field.
func (o *UserGroup) SetBsdgrpGroup(v string) {
	o.BsdgrpGroup = &v
}

// GetBsdgrpBuiltin returns the BsdgrpBuiltin field value if set, zero value otherwise.
func (o *UserGroup) GetBsdgrpBuiltin() bool {
	if o == nil || isNil(o.BsdgrpBuiltin) {
		var ret bool
		return ret
	}
	return *o.BsdgrpBuiltin
}

// GetBsdgrpBuiltinOk returns a tuple with the BsdgrpBuiltin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetBsdgrpBuiltinOk() (*bool, bool) {
	if o == nil || isNil(o.BsdgrpBuiltin) {
		return nil, false
	}
	return o.BsdgrpBuiltin, true
}

// HasBsdgrpBuiltin returns a boolean if a field has been set.
func (o *UserGroup) HasBsdgrpBuiltin() bool {
	if o != nil && !isNil(o.BsdgrpBuiltin) {
		return true
	}

	return false
}

// SetBsdgrpBuiltin gets a reference to the given bool and assigns it to the BsdgrpBuiltin field.
func (o *UserGroup) SetBsdgrpBuiltin(v bool) {
	o.BsdgrpBuiltin = &v
}

// GetBsdgrpSudo returns the BsdgrpSudo field value if set, zero value otherwise.
func (o *UserGroup) GetBsdgrpSudo() bool {
	if o == nil || isNil(o.BsdgrpSudo) {
		var ret bool
		return ret
	}
	return *o.BsdgrpSudo
}

// GetBsdgrpSudoOk returns a tuple with the BsdgrpSudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetBsdgrpSudoOk() (*bool, bool) {
	if o == nil || isNil(o.BsdgrpSudo) {
		return nil, false
	}
	return o.BsdgrpSudo, true
}

// HasBsdgrpSudo returns a boolean if a field has been set.
func (o *UserGroup) HasBsdgrpSudo() bool {
	if o != nil && !isNil(o.BsdgrpSudo) {
		return true
	}

	return false
}

// SetBsdgrpSudo gets a reference to the given bool and assigns it to the BsdgrpSudo field.
func (o *UserGroup) SetBsdgrpSudo(v bool) {
	o.BsdgrpSudo = &v
}

// GetBsdgrpSudoNopasswd returns the BsdgrpSudoNopasswd field value if set, zero value otherwise.
func (o *UserGroup) GetBsdgrpSudoNopasswd() bool {
	if o == nil || isNil(o.BsdgrpSudoNopasswd) {
		var ret bool
		return ret
	}
	return *o.BsdgrpSudoNopasswd
}

// GetBsdgrpSudoNopasswdOk returns a tuple with the BsdgrpSudoNopasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetBsdgrpSudoNopasswdOk() (*bool, bool) {
	if o == nil || isNil(o.BsdgrpSudoNopasswd) {
		return nil, false
	}
	return o.BsdgrpSudoNopasswd, true
}

// HasBsdgrpSudoNopasswd returns a boolean if a field has been set.
func (o *UserGroup) HasBsdgrpSudoNopasswd() bool {
	if o != nil && !isNil(o.BsdgrpSudoNopasswd) {
		return true
	}

	return false
}

// SetBsdgrpSudoNopasswd gets a reference to the given bool and assigns it to the BsdgrpSudoNopasswd field.
func (o *UserGroup) SetBsdgrpSudoNopasswd(v bool) {
	o.BsdgrpSudoNopasswd = &v
}

// GetBsdgrpSudoCommands returns the BsdgrpSudoCommands field value if set, zero value otherwise.
func (o *UserGroup) GetBsdgrpSudoCommands() []string {
	if o == nil || isNil(o.BsdgrpSudoCommands) {
		var ret []string
		return ret
	}
	return o.BsdgrpSudoCommands
}

// GetBsdgrpSudoCommandsOk returns a tuple with the BsdgrpSudoCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetBsdgrpSudoCommandsOk() ([]string, bool) {
	if o == nil || isNil(o.BsdgrpSudoCommands) {
		return nil, false
	}
	return o.BsdgrpSudoCommands, true
}

// HasBsdgrpSudoCommands returns a boolean if a field has been set.
func (o *UserGroup) HasBsdgrpSudoCommands() bool {
	if o != nil && !isNil(o.BsdgrpSudoCommands) {
		return true
	}

	return false
}

// SetBsdgrpSudoCommands gets a reference to the given []string and assigns it to the BsdgrpSudoCommands field.
func (o *UserGroup) SetBsdgrpSudoCommands(v []string) {
	o.BsdgrpSudoCommands = v
}

// GetBsdgrpSmb returns the BsdgrpSmb field value if set, zero value otherwise.
func (o *UserGroup) GetBsdgrpSmb() bool {
	if o == nil || isNil(o.BsdgrpSmb) {
		var ret bool
		return ret
	}
	return *o.BsdgrpSmb
}

// GetBsdgrpSmbOk returns a tuple with the BsdgrpSmb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetBsdgrpSmbOk() (*bool, bool) {
	if o == nil || isNil(o.BsdgrpSmb) {
		return nil, false
	}
	return o.BsdgrpSmb, true
}

// HasBsdgrpSmb returns a boolean if a field has been set.
func (o *UserGroup) HasBsdgrpSmb() bool {
	if o != nil && !isNil(o.BsdgrpSmb) {
		return true
	}

	return false
}

// SetBsdgrpSmb gets a reference to the given bool and assigns it to the BsdgrpSmb field.
func (o *UserGroup) SetBsdgrpSmb(v bool) {
	o.BsdgrpSmb = &v
}

func (o UserGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.BsdgrpGid) {
		toSerialize["bsdgrp_gid"] = o.BsdgrpGid
	}
	if !isNil(o.BsdgrpGroup) {
		toSerialize["bsdgrp_group"] = o.BsdgrpGroup
	}
	if !isNil(o.BsdgrpBuiltin) {
		toSerialize["bsdgrp_builtin"] = o.BsdgrpBuiltin
	}
	if !isNil(o.BsdgrpSudo) {
		toSerialize["bsdgrp_sudo"] = o.BsdgrpSudo
	}
	if !isNil(o.BsdgrpSudoNopasswd) {
		toSerialize["bsdgrp_sudo_nopasswd"] = o.BsdgrpSudoNopasswd
	}
	if !isNil(o.BsdgrpSudoCommands) {
		toSerialize["bsdgrp_sudo_commands"] = o.BsdgrpSudoCommands
	}
	if !isNil(o.BsdgrpSmb) {
		toSerialize["bsdgrp_smb"] = o.BsdgrpSmb
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserGroup) UnmarshalJSON(bytes []byte) (err error) {
	varUserGroup := _UserGroup{}

	if err = json.Unmarshal(bytes, &varUserGroup); err == nil {
		*o = UserGroup(varUserGroup)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "bsdgrp_gid")
		delete(additionalProperties, "bsdgrp_group")
		delete(additionalProperties, "bsdgrp_builtin")
		delete(additionalProperties, "bsdgrp_sudo")
		delete(additionalProperties, "bsdgrp_sudo_nopasswd")
		delete(additionalProperties, "bsdgrp_sudo_commands")
		delete(additionalProperties, "bsdgrp_smb")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserGroup struct {
	value *UserGroup
	isSet bool
}

func (v NullableUserGroup) Get() *UserGroup {
	return v.value
}

func (v *NullableUserGroup) Set(val *UserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroup(val *UserGroup) *NullableUserGroup {
	return &NullableUserGroup{value: val, isSet: true}
}

func (v NullableUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
