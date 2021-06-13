# InterfaceUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DisableOffloadCapabilities** | Pointer to **bool** |  | [optional] 
**Ipv4Dhcp** | Pointer to **bool** |  | [optional] 
**Ipv6Auto** | Pointer to **bool** |  | [optional] 
**Aliases** | Pointer to **[]map[string]interface{}** |  | [optional] 
**FailoverCritical** | Pointer to **bool** |  | [optional] 
**FailoverGroup** | Pointer to **NullableInt32** |  | [optional] 
**FailoverVhid** | Pointer to **NullableInt32** |  | [optional] 
**FailoverAliases** | Pointer to **[]map[string]interface{}** |  | [optional] 
**FailoverVirtualAliases** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BridgeMembers** | Pointer to **[]interface{}** |  | [optional] 
**LagProtocol** | Pointer to **string** |  | [optional] 
**LagPorts** | Pointer to **[]string** |  | [optional] 
**VlanParentInterface** | Pointer to **string** |  | [optional] 
**VlanTag** | Pointer to **int32** |  | [optional] 
**VlanPcp** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 

## Methods

### NewInterfaceUpdate1

`func NewInterfaceUpdate1() *InterfaceUpdate1`

NewInterfaceUpdate1 instantiates a new InterfaceUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceUpdate1WithDefaults

`func NewInterfaceUpdate1WithDefaults() *InterfaceUpdate1`

NewInterfaceUpdate1WithDefaults instantiates a new InterfaceUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InterfaceUpdate1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceUpdate1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceUpdate1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InterfaceUpdate1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InterfaceUpdate1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceUpdate1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceUpdate1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceUpdate1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InterfaceUpdate1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InterfaceUpdate1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisableOffloadCapabilities

`func (o *InterfaceUpdate1) GetDisableOffloadCapabilities() bool`

GetDisableOffloadCapabilities returns the DisableOffloadCapabilities field if non-nil, zero value otherwise.

### GetDisableOffloadCapabilitiesOk

`func (o *InterfaceUpdate1) GetDisableOffloadCapabilitiesOk() (*bool, bool)`

GetDisableOffloadCapabilitiesOk returns a tuple with the DisableOffloadCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOffloadCapabilities

`func (o *InterfaceUpdate1) SetDisableOffloadCapabilities(v bool)`

SetDisableOffloadCapabilities sets DisableOffloadCapabilities field to given value.

### HasDisableOffloadCapabilities

`func (o *InterfaceUpdate1) HasDisableOffloadCapabilities() bool`

HasDisableOffloadCapabilities returns a boolean if a field has been set.

### GetIpv4Dhcp

`func (o *InterfaceUpdate1) GetIpv4Dhcp() bool`

GetIpv4Dhcp returns the Ipv4Dhcp field if non-nil, zero value otherwise.

### GetIpv4DhcpOk

`func (o *InterfaceUpdate1) GetIpv4DhcpOk() (*bool, bool)`

GetIpv4DhcpOk returns a tuple with the Ipv4Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Dhcp

`func (o *InterfaceUpdate1) SetIpv4Dhcp(v bool)`

SetIpv4Dhcp sets Ipv4Dhcp field to given value.

### HasIpv4Dhcp

`func (o *InterfaceUpdate1) HasIpv4Dhcp() bool`

HasIpv4Dhcp returns a boolean if a field has been set.

### GetIpv6Auto

`func (o *InterfaceUpdate1) GetIpv6Auto() bool`

GetIpv6Auto returns the Ipv6Auto field if non-nil, zero value otherwise.

### GetIpv6AutoOk

`func (o *InterfaceUpdate1) GetIpv6AutoOk() (*bool, bool)`

GetIpv6AutoOk returns a tuple with the Ipv6Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Auto

`func (o *InterfaceUpdate1) SetIpv6Auto(v bool)`

SetIpv6Auto sets Ipv6Auto field to given value.

### HasIpv6Auto

`func (o *InterfaceUpdate1) HasIpv6Auto() bool`

HasIpv6Auto returns a boolean if a field has been set.

### GetAliases

`func (o *InterfaceUpdate1) GetAliases() []map[string]interface{}`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InterfaceUpdate1) GetAliasesOk() (*[]map[string]interface{}, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InterfaceUpdate1) SetAliases(v []map[string]interface{})`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InterfaceUpdate1) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetFailoverCritical

`func (o *InterfaceUpdate1) GetFailoverCritical() bool`

GetFailoverCritical returns the FailoverCritical field if non-nil, zero value otherwise.

### GetFailoverCriticalOk

`func (o *InterfaceUpdate1) GetFailoverCriticalOk() (*bool, bool)`

GetFailoverCriticalOk returns a tuple with the FailoverCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverCritical

`func (o *InterfaceUpdate1) SetFailoverCritical(v bool)`

SetFailoverCritical sets FailoverCritical field to given value.

### HasFailoverCritical

`func (o *InterfaceUpdate1) HasFailoverCritical() bool`

HasFailoverCritical returns a boolean if a field has been set.

### GetFailoverGroup

`func (o *InterfaceUpdate1) GetFailoverGroup() int32`

GetFailoverGroup returns the FailoverGroup field if non-nil, zero value otherwise.

### GetFailoverGroupOk

`func (o *InterfaceUpdate1) GetFailoverGroupOk() (*int32, bool)`

GetFailoverGroupOk returns a tuple with the FailoverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverGroup

`func (o *InterfaceUpdate1) SetFailoverGroup(v int32)`

SetFailoverGroup sets FailoverGroup field to given value.

### HasFailoverGroup

`func (o *InterfaceUpdate1) HasFailoverGroup() bool`

HasFailoverGroup returns a boolean if a field has been set.

### SetFailoverGroupNil

`func (o *InterfaceUpdate1) SetFailoverGroupNil(b bool)`

 SetFailoverGroupNil sets the value for FailoverGroup to be an explicit nil

### UnsetFailoverGroup
`func (o *InterfaceUpdate1) UnsetFailoverGroup()`

UnsetFailoverGroup ensures that no value is present for FailoverGroup, not even an explicit nil
### GetFailoverVhid

`func (o *InterfaceUpdate1) GetFailoverVhid() int32`

GetFailoverVhid returns the FailoverVhid field if non-nil, zero value otherwise.

### GetFailoverVhidOk

`func (o *InterfaceUpdate1) GetFailoverVhidOk() (*int32, bool)`

GetFailoverVhidOk returns a tuple with the FailoverVhid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverVhid

`func (o *InterfaceUpdate1) SetFailoverVhid(v int32)`

SetFailoverVhid sets FailoverVhid field to given value.

### HasFailoverVhid

`func (o *InterfaceUpdate1) HasFailoverVhid() bool`

HasFailoverVhid returns a boolean if a field has been set.

### SetFailoverVhidNil

`func (o *InterfaceUpdate1) SetFailoverVhidNil(b bool)`

 SetFailoverVhidNil sets the value for FailoverVhid to be an explicit nil

### UnsetFailoverVhid
`func (o *InterfaceUpdate1) UnsetFailoverVhid()`

UnsetFailoverVhid ensures that no value is present for FailoverVhid, not even an explicit nil
### GetFailoverAliases

`func (o *InterfaceUpdate1) GetFailoverAliases() []map[string]interface{}`

GetFailoverAliases returns the FailoverAliases field if non-nil, zero value otherwise.

### GetFailoverAliasesOk

`func (o *InterfaceUpdate1) GetFailoverAliasesOk() (*[]map[string]interface{}, bool)`

GetFailoverAliasesOk returns a tuple with the FailoverAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAliases

`func (o *InterfaceUpdate1) SetFailoverAliases(v []map[string]interface{})`

SetFailoverAliases sets FailoverAliases field to given value.

### HasFailoverAliases

`func (o *InterfaceUpdate1) HasFailoverAliases() bool`

HasFailoverAliases returns a boolean if a field has been set.

### GetFailoverVirtualAliases

`func (o *InterfaceUpdate1) GetFailoverVirtualAliases() []map[string]interface{}`

GetFailoverVirtualAliases returns the FailoverVirtualAliases field if non-nil, zero value otherwise.

### GetFailoverVirtualAliasesOk

`func (o *InterfaceUpdate1) GetFailoverVirtualAliasesOk() (*[]map[string]interface{}, bool)`

GetFailoverVirtualAliasesOk returns a tuple with the FailoverVirtualAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverVirtualAliases

`func (o *InterfaceUpdate1) SetFailoverVirtualAliases(v []map[string]interface{})`

SetFailoverVirtualAliases sets FailoverVirtualAliases field to given value.

### HasFailoverVirtualAliases

`func (o *InterfaceUpdate1) HasFailoverVirtualAliases() bool`

HasFailoverVirtualAliases returns a boolean if a field has been set.

### GetBridgeMembers

`func (o *InterfaceUpdate1) GetBridgeMembers() []interface{}`

GetBridgeMembers returns the BridgeMembers field if non-nil, zero value otherwise.

### GetBridgeMembersOk

`func (o *InterfaceUpdate1) GetBridgeMembersOk() (*[]interface{}, bool)`

GetBridgeMembersOk returns a tuple with the BridgeMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeMembers

`func (o *InterfaceUpdate1) SetBridgeMembers(v []interface{})`

SetBridgeMembers sets BridgeMembers field to given value.

### HasBridgeMembers

`func (o *InterfaceUpdate1) HasBridgeMembers() bool`

HasBridgeMembers returns a boolean if a field has been set.

### GetLagProtocol

`func (o *InterfaceUpdate1) GetLagProtocol() string`

GetLagProtocol returns the LagProtocol field if non-nil, zero value otherwise.

### GetLagProtocolOk

`func (o *InterfaceUpdate1) GetLagProtocolOk() (*string, bool)`

GetLagProtocolOk returns a tuple with the LagProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagProtocol

`func (o *InterfaceUpdate1) SetLagProtocol(v string)`

SetLagProtocol sets LagProtocol field to given value.

### HasLagProtocol

`func (o *InterfaceUpdate1) HasLagProtocol() bool`

HasLagProtocol returns a boolean if a field has been set.

### GetLagPorts

`func (o *InterfaceUpdate1) GetLagPorts() []string`

GetLagPorts returns the LagPorts field if non-nil, zero value otherwise.

### GetLagPortsOk

`func (o *InterfaceUpdate1) GetLagPortsOk() (*[]string, bool)`

GetLagPortsOk returns a tuple with the LagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPorts

`func (o *InterfaceUpdate1) SetLagPorts(v []string)`

SetLagPorts sets LagPorts field to given value.

### HasLagPorts

`func (o *InterfaceUpdate1) HasLagPorts() bool`

HasLagPorts returns a boolean if a field has been set.

### GetVlanParentInterface

`func (o *InterfaceUpdate1) GetVlanParentInterface() string`

GetVlanParentInterface returns the VlanParentInterface field if non-nil, zero value otherwise.

### GetVlanParentInterfaceOk

`func (o *InterfaceUpdate1) GetVlanParentInterfaceOk() (*string, bool)`

GetVlanParentInterfaceOk returns a tuple with the VlanParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParentInterface

`func (o *InterfaceUpdate1) SetVlanParentInterface(v string)`

SetVlanParentInterface sets VlanParentInterface field to given value.

### HasVlanParentInterface

`func (o *InterfaceUpdate1) HasVlanParentInterface() bool`

HasVlanParentInterface returns a boolean if a field has been set.

### GetVlanTag

`func (o *InterfaceUpdate1) GetVlanTag() int32`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *InterfaceUpdate1) GetVlanTagOk() (*int32, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *InterfaceUpdate1) SetVlanTag(v int32)`

SetVlanTag sets VlanTag field to given value.

### HasVlanTag

`func (o *InterfaceUpdate1) HasVlanTag() bool`

HasVlanTag returns a boolean if a field has been set.

### GetVlanPcp

`func (o *InterfaceUpdate1) GetVlanPcp() int32`

GetVlanPcp returns the VlanPcp field if non-nil, zero value otherwise.

### GetVlanPcpOk

`func (o *InterfaceUpdate1) GetVlanPcpOk() (*int32, bool)`

GetVlanPcpOk returns a tuple with the VlanPcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPcp

`func (o *InterfaceUpdate1) SetVlanPcp(v int32)`

SetVlanPcp sets VlanPcp field to given value.

### HasVlanPcp

`func (o *InterfaceUpdate1) HasVlanPcp() bool`

HasVlanPcp returns a boolean if a field has been set.

### SetVlanPcpNil

`func (o *InterfaceUpdate1) SetVlanPcpNil(b bool)`

 SetVlanPcpNil sets the value for VlanPcp to be an explicit nil

### UnsetVlanPcp
`func (o *InterfaceUpdate1) UnsetVlanPcp()`

UnsetVlanPcp ensures that no value is present for VlanPcp, not even an explicit nil
### GetMtu

`func (o *InterfaceUpdate1) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InterfaceUpdate1) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InterfaceUpdate1) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InterfaceUpdate1) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InterfaceUpdate1) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InterfaceUpdate1) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetOptions

`func (o *InterfaceUpdate1) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InterfaceUpdate1) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InterfaceUpdate1) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InterfaceUpdate1) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


