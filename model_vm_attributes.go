/*
 * TrueNAS RESTful API
 *
 * Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// VMAttributes struct for VMAttributes
type VMAttributes struct {
	Type                 *string `json:"type,omitempty"`
	Mac                  *string `json:"mac,omitempty"`
	NicAttach            *string `json:"nic_attach,omitempty"`
	Path                 *string `json:"path,omitempty"`
	Wait                 *bool   `json:"wait,omitempty"`
	VncPort              *int32  `json:"vnc_port,omitempty"`
	VncResolution        *string `json:"vnc_resolution,omitempty"`
	VncBind              *string `json:"vnc_bind,omitempty"`
	VncPassword          *string `json:"vnc_password,omitempty"`
	VncWeb               *bool   `json:"vnc_web,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VMAttributes VMAttributes

// NewVMAttributes instantiates a new VMAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMAttributes() *VMAttributes {
	this := VMAttributes{}
	return &this
}

// NewVMAttributesWithDefaults instantiates a new VMAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMAttributesWithDefaults() *VMAttributes {
	this := VMAttributes{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VMAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VMAttributes) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VMAttributes) SetType(v string) {
	o.Type = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *VMAttributes) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *VMAttributes) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *VMAttributes) SetMac(v string) {
	o.Mac = &v
}

// GetNicAttach returns the NicAttach field value if set, zero value otherwise.
func (o *VMAttributes) GetNicAttach() string {
	if o == nil || o.NicAttach == nil {
		var ret string
		return ret
	}
	return *o.NicAttach
}

// GetNicAttachOk returns a tuple with the NicAttach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetNicAttachOk() (*string, bool) {
	if o == nil || o.NicAttach == nil {
		return nil, false
	}
	return o.NicAttach, true
}

// HasNicAttach returns a boolean if a field has been set.
func (o *VMAttributes) HasNicAttach() bool {
	if o != nil && o.NicAttach != nil {
		return true
	}

	return false
}

// SetNicAttach gets a reference to the given string and assigns it to the NicAttach field.
func (o *VMAttributes) SetNicAttach(v string) {
	o.NicAttach = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VMAttributes) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VMAttributes) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VMAttributes) SetPath(v string) {
	o.Path = &v
}

// GetWait returns the Wait field value if set, zero value otherwise.
func (o *VMAttributes) GetWait() bool {
	if o == nil || o.Wait == nil {
		var ret bool
		return ret
	}
	return *o.Wait
}

// GetWaitOk returns a tuple with the Wait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetWaitOk() (*bool, bool) {
	if o == nil || o.Wait == nil {
		return nil, false
	}
	return o.Wait, true
}

// HasWait returns a boolean if a field has been set.
func (o *VMAttributes) HasWait() bool {
	if o != nil && o.Wait != nil {
		return true
	}

	return false
}

// SetWait gets a reference to the given bool and assigns it to the Wait field.
func (o *VMAttributes) SetWait(v bool) {
	o.Wait = &v
}

// GetVncPort returns the VncPort field value if set, zero value otherwise.
func (o *VMAttributes) GetVncPort() int32 {
	if o == nil || o.VncPort == nil {
		var ret int32
		return ret
	}
	return *o.VncPort
}

// GetVncPortOk returns a tuple with the VncPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetVncPortOk() (*int32, bool) {
	if o == nil || o.VncPort == nil {
		return nil, false
	}
	return o.VncPort, true
}

// HasVncPort returns a boolean if a field has been set.
func (o *VMAttributes) HasVncPort() bool {
	if o != nil && o.VncPort != nil {
		return true
	}

	return false
}

// SetVncPort gets a reference to the given int32 and assigns it to the VncPort field.
func (o *VMAttributes) SetVncPort(v int32) {
	o.VncPort = &v
}

// GetVncResolution returns the VncResolution field value if set, zero value otherwise.
func (o *VMAttributes) GetVncResolution() string {
	if o == nil || o.VncResolution == nil {
		var ret string
		return ret
	}
	return *o.VncResolution
}

// GetVncResolutionOk returns a tuple with the VncResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetVncResolutionOk() (*string, bool) {
	if o == nil || o.VncResolution == nil {
		return nil, false
	}
	return o.VncResolution, true
}

// HasVncResolution returns a boolean if a field has been set.
func (o *VMAttributes) HasVncResolution() bool {
	if o != nil && o.VncResolution != nil {
		return true
	}

	return false
}

// SetVncResolution gets a reference to the given string and assigns it to the VncResolution field.
func (o *VMAttributes) SetVncResolution(v string) {
	o.VncResolution = &v
}

// GetVncBind returns the VncBind field value if set, zero value otherwise.
func (o *VMAttributes) GetVncBind() string {
	if o == nil || o.VncBind == nil {
		var ret string
		return ret
	}
	return *o.VncBind
}

// GetVncBindOk returns a tuple with the VncBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetVncBindOk() (*string, bool) {
	if o == nil || o.VncBind == nil {
		return nil, false
	}
	return o.VncBind, true
}

// HasVncBind returns a boolean if a field has been set.
func (o *VMAttributes) HasVncBind() bool {
	if o != nil && o.VncBind != nil {
		return true
	}

	return false
}

// SetVncBind gets a reference to the given string and assigns it to the VncBind field.
func (o *VMAttributes) SetVncBind(v string) {
	o.VncBind = &v
}

// GetVncPassword returns the VncPassword field value if set, zero value otherwise.
func (o *VMAttributes) GetVncPassword() string {
	if o == nil || o.VncPassword == nil {
		var ret string
		return ret
	}
	return *o.VncPassword
}

// GetVncPasswordOk returns a tuple with the VncPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetVncPasswordOk() (*string, bool) {
	if o == nil || o.VncPassword == nil {
		return nil, false
	}
	return o.VncPassword, true
}

// HasVncPassword returns a boolean if a field has been set.
func (o *VMAttributes) HasVncPassword() bool {
	if o != nil && o.VncPassword != nil {
		return true
	}

	return false
}

// SetVncPassword gets a reference to the given string and assigns it to the VncPassword field.
func (o *VMAttributes) SetVncPassword(v string) {
	o.VncPassword = &v
}

// GetVncWeb returns the VncWeb field value if set, zero value otherwise.
func (o *VMAttributes) GetVncWeb() bool {
	if o == nil || o.VncWeb == nil {
		var ret bool
		return ret
	}
	return *o.VncWeb
}

// GetVncWebOk returns a tuple with the VncWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMAttributes) GetVncWebOk() (*bool, bool) {
	if o == nil || o.VncWeb == nil {
		return nil, false
	}
	return o.VncWeb, true
}

// HasVncWeb returns a boolean if a field has been set.
func (o *VMAttributes) HasVncWeb() bool {
	if o != nil && o.VncWeb != nil {
		return true
	}

	return false
}

// SetVncWeb gets a reference to the given bool and assigns it to the VncWeb field.
func (o *VMAttributes) SetVncWeb(v bool) {
	o.VncWeb = &v
}

func (o VMAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Mac != nil {
		toSerialize["mac"] = o.Mac
	}
	if o.NicAttach != nil {
		toSerialize["nic_attach"] = o.NicAttach
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Wait != nil {
		toSerialize["wait"] = o.Wait
	}
	if o.VncPort != nil {
		toSerialize["vnc_port"] = o.VncPort
	}
	if o.VncResolution != nil {
		toSerialize["vnc_resolution"] = o.VncResolution
	}
	if o.VncBind != nil {
		toSerialize["vnc_bind"] = o.VncBind
	}
	if o.VncPassword != nil {
		toSerialize["vnc_password"] = o.VncPassword
	}
	if o.VncWeb != nil {
		toSerialize["vnc_web"] = o.VncWeb
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VMAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varVMAttributes := _VMAttributes{}

	if err = json.Unmarshal(bytes, &varVMAttributes); err == nil {
		*o = VMAttributes(varVMAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "nic_attach")
		delete(additionalProperties, "path")
		delete(additionalProperties, "wait")
		delete(additionalProperties, "vnc_port")
		delete(additionalProperties, "vnc_resolution")
		delete(additionalProperties, "vnc_bind")
		delete(additionalProperties, "vnc_password")
		delete(additionalProperties, "vnc_web")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMAttributes struct {
	value *VMAttributes
	isSet bool
}

func (v NullableVMAttributes) Get() *VMAttributes {
	return v.value
}

func (v *NullableVMAttributes) Set(val *VMAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableVMAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableVMAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMAttributes(val *VMAttributes) *NullableVMAttributes {
	return &NullableVMAttributes{value: val, isSet: true}
}

func (v NullableVMAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
