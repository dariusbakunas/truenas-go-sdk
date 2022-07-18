# NetworkSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | Pointer to [**map[string]NetworkSummaryIpsValue**](NetworkSummaryIpsValue.md) |  | [optional] 
**DefaultRoutes** | Pointer to **[]string** |  | [optional] 
**Nameservers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkSummary

`func NewNetworkSummary() *NetworkSummary`

NewNetworkSummary instantiates a new NetworkSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSummaryWithDefaults

`func NewNetworkSummaryWithDefaults() *NetworkSummary`

NewNetworkSummaryWithDefaults instantiates a new NetworkSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *NetworkSummary) GetIps() map[string]NetworkSummaryIpsValue`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *NetworkSummary) GetIpsOk() (*map[string]NetworkSummaryIpsValue, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *NetworkSummary) SetIps(v map[string]NetworkSummaryIpsValue)`

SetIps sets Ips field to given value.

### HasIps

`func (o *NetworkSummary) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetDefaultRoutes

`func (o *NetworkSummary) GetDefaultRoutes() []string`

GetDefaultRoutes returns the DefaultRoutes field if non-nil, zero value otherwise.

### GetDefaultRoutesOk

`func (o *NetworkSummary) GetDefaultRoutesOk() (*[]string, bool)`

GetDefaultRoutesOk returns a tuple with the DefaultRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoutes

`func (o *NetworkSummary) SetDefaultRoutes(v []string)`

SetDefaultRoutes sets DefaultRoutes field to given value.

### HasDefaultRoutes

`func (o *NetworkSummary) HasDefaultRoutes() bool`

HasDefaultRoutes returns a boolean if a field has been set.

### GetNameservers

`func (o *NetworkSummary) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *NetworkSummary) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *NetworkSummary) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *NetworkSummary) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


