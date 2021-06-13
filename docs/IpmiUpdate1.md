# IpmiUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipaddress** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Dhcp** | Pointer to **bool** |  | [optional] 
**Vlan** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewIpmiUpdate1

`func NewIpmiUpdate1() *IpmiUpdate1`

NewIpmiUpdate1 instantiates a new IpmiUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiUpdate1WithDefaults

`func NewIpmiUpdate1WithDefaults() *IpmiUpdate1`

NewIpmiUpdate1WithDefaults instantiates a new IpmiUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpaddress

`func (o *IpmiUpdate1) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *IpmiUpdate1) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *IpmiUpdate1) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *IpmiUpdate1) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetNetmask

`func (o *IpmiUpdate1) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IpmiUpdate1) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IpmiUpdate1) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IpmiUpdate1) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetGateway

`func (o *IpmiUpdate1) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IpmiUpdate1) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IpmiUpdate1) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IpmiUpdate1) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPassword

`func (o *IpmiUpdate1) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IpmiUpdate1) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IpmiUpdate1) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IpmiUpdate1) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDhcp

`func (o *IpmiUpdate1) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *IpmiUpdate1) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *IpmiUpdate1) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *IpmiUpdate1) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetVlan

`func (o *IpmiUpdate1) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *IpmiUpdate1) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *IpmiUpdate1) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *IpmiUpdate1) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *IpmiUpdate1) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *IpmiUpdate1) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


