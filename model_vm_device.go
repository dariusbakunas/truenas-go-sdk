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

// VMDevice struct for VMDevice
type VMDevice struct {
	Id                   int32                  `json:"id"`
	Dtype                string                 `json:"dtype"`
	Order                *int32                 `json:"order,omitempty"`
	Vm                   *int32                 `json:"vm,omitempty"`
	Attributes           map[string]interface{} `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VMDevice VMDevice

// NewVMDevice instantiates a new VMDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMDevice(id int32, dtype string) *VMDevice {
	this := VMDevice{}
	this.Id = id
	this.Dtype = dtype
	return &this
}

// NewVMDeviceWithDefaults instantiates a new VMDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMDeviceWithDefaults() *VMDevice {
	this := VMDevice{}
	return &this
}

// GetId returns the Id field value
func (o *VMDevice) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMDevice) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMDevice) SetId(v int32) {
	o.Id = v
}

// GetDtype returns the Dtype field value
func (o *VMDevice) GetDtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dtype
}

// GetDtypeOk returns a tuple with the Dtype field value
// and a boolean to check if the value has been set.
func (o *VMDevice) GetDtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dtype, true
}

// SetDtype sets field value
func (o *VMDevice) SetDtype(v string) {
	o.Dtype = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VMDevice) GetOrder() int32 {
	if o == nil || isNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMDevice) GetOrderOk() (*int32, bool) {
	if o == nil || isNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VMDevice) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *VMDevice) SetOrder(v int32) {
	o.Order = &v
}

// GetVm returns the Vm field value if set, zero value otherwise.
func (o *VMDevice) GetVm() int32 {
	if o == nil || isNil(o.Vm) {
		var ret int32
		return ret
	}
	return *o.Vm
}

// GetVmOk returns a tuple with the Vm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMDevice) GetVmOk() (*int32, bool) {
	if o == nil || isNil(o.Vm) {
		return nil, false
	}
	return o.Vm, true
}

// HasVm returns a boolean if a field has been set.
func (o *VMDevice) HasVm() bool {
	if o != nil && !isNil(o.Vm) {
		return true
	}

	return false
}

// SetVm gets a reference to the given int32 and assigns it to the Vm field.
func (o *VMDevice) SetVm(v int32) {
	o.Vm = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *VMDevice) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMDevice) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *VMDevice) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *VMDevice) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o VMDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["dtype"] = o.Dtype
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !isNil(o.Vm) {
		toSerialize["vm"] = o.Vm
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VMDevice) UnmarshalJSON(bytes []byte) (err error) {
	varVMDevice := _VMDevice{}

	if err = json.Unmarshal(bytes, &varVMDevice); err == nil {
		*o = VMDevice(varVMDevice)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "dtype")
		delete(additionalProperties, "order")
		delete(additionalProperties, "vm")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMDevice struct {
	value *VMDevice
	isSet bool
}

func (v NullableVMDevice) Get() *VMDevice {
	return v.value
}

func (v *NullableVMDevice) Set(val *VMDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableVMDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableVMDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMDevice(val *VMDevice) *NullableVMDevice {
	return &NullableVMDevice{value: val, isSet: true}
}

func (v NullableVMDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
