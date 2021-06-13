# NetworkConfigurationUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] 
**HostnameB** | Pointer to **string** |  | [optional] 
**HostnameVirtual** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**ServiceAnnouncement** | Pointer to [**NetworkConfigurationUpdate0ServiceAnnouncement**](NetworkConfigurationUpdate0ServiceAnnouncement.md) |  | [optional] 
**Ipv4gateway** | Pointer to **string** |  | [optional] 
**Ipv6gateway** | Pointer to **string** |  | [optional] 
**Nameserver1** | Pointer to **string** |  | [optional] 
**Nameserver2** | Pointer to **string** |  | [optional] 
**Nameserver3** | Pointer to **string** |  | [optional] 
**Httpproxy** | Pointer to **string** |  | [optional] 
**NetwaitEnabled** | Pointer to **bool** |  | [optional] 
**NetwaitIp** | Pointer to **[]string** |  | [optional] 
**Hosts** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkConfigurationUpdate0

`func NewNetworkConfigurationUpdate0() *NetworkConfigurationUpdate0`

NewNetworkConfigurationUpdate0 instantiates a new NetworkConfigurationUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigurationUpdate0WithDefaults

`func NewNetworkConfigurationUpdate0WithDefaults() *NetworkConfigurationUpdate0`

NewNetworkConfigurationUpdate0WithDefaults instantiates a new NetworkConfigurationUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *NetworkConfigurationUpdate0) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NetworkConfigurationUpdate0) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NetworkConfigurationUpdate0) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NetworkConfigurationUpdate0) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHostnameB

`func (o *NetworkConfigurationUpdate0) GetHostnameB() string`

GetHostnameB returns the HostnameB field if non-nil, zero value otherwise.

### GetHostnameBOk

`func (o *NetworkConfigurationUpdate0) GetHostnameBOk() (*string, bool)`

GetHostnameBOk returns a tuple with the HostnameB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameB

`func (o *NetworkConfigurationUpdate0) SetHostnameB(v string)`

SetHostnameB sets HostnameB field to given value.

### HasHostnameB

`func (o *NetworkConfigurationUpdate0) HasHostnameB() bool`

HasHostnameB returns a boolean if a field has been set.

### GetHostnameVirtual

`func (o *NetworkConfigurationUpdate0) GetHostnameVirtual() string`

GetHostnameVirtual returns the HostnameVirtual field if non-nil, zero value otherwise.

### GetHostnameVirtualOk

`func (o *NetworkConfigurationUpdate0) GetHostnameVirtualOk() (*string, bool)`

GetHostnameVirtualOk returns a tuple with the HostnameVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVirtual

`func (o *NetworkConfigurationUpdate0) SetHostnameVirtual(v string)`

SetHostnameVirtual sets HostnameVirtual field to given value.

### HasHostnameVirtual

`func (o *NetworkConfigurationUpdate0) HasHostnameVirtual() bool`

HasHostnameVirtual returns a boolean if a field has been set.

### GetDomain

`func (o *NetworkConfigurationUpdate0) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NetworkConfigurationUpdate0) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NetworkConfigurationUpdate0) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NetworkConfigurationUpdate0) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomains

`func (o *NetworkConfigurationUpdate0) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *NetworkConfigurationUpdate0) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *NetworkConfigurationUpdate0) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *NetworkConfigurationUpdate0) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetServiceAnnouncement

`func (o *NetworkConfigurationUpdate0) GetServiceAnnouncement() NetworkConfigurationUpdate0ServiceAnnouncement`

GetServiceAnnouncement returns the ServiceAnnouncement field if non-nil, zero value otherwise.

### GetServiceAnnouncementOk

`func (o *NetworkConfigurationUpdate0) GetServiceAnnouncementOk() (*NetworkConfigurationUpdate0ServiceAnnouncement, bool)`

GetServiceAnnouncementOk returns a tuple with the ServiceAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAnnouncement

`func (o *NetworkConfigurationUpdate0) SetServiceAnnouncement(v NetworkConfigurationUpdate0ServiceAnnouncement)`

SetServiceAnnouncement sets ServiceAnnouncement field to given value.

### HasServiceAnnouncement

`func (o *NetworkConfigurationUpdate0) HasServiceAnnouncement() bool`

HasServiceAnnouncement returns a boolean if a field has been set.

### GetIpv4gateway

`func (o *NetworkConfigurationUpdate0) GetIpv4gateway() string`

GetIpv4gateway returns the Ipv4gateway field if non-nil, zero value otherwise.

### GetIpv4gatewayOk

`func (o *NetworkConfigurationUpdate0) GetIpv4gatewayOk() (*string, bool)`

GetIpv4gatewayOk returns a tuple with the Ipv4gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4gateway

`func (o *NetworkConfigurationUpdate0) SetIpv4gateway(v string)`

SetIpv4gateway sets Ipv4gateway field to given value.

### HasIpv4gateway

`func (o *NetworkConfigurationUpdate0) HasIpv4gateway() bool`

HasIpv4gateway returns a boolean if a field has been set.

### GetIpv6gateway

`func (o *NetworkConfigurationUpdate0) GetIpv6gateway() string`

GetIpv6gateway returns the Ipv6gateway field if non-nil, zero value otherwise.

### GetIpv6gatewayOk

`func (o *NetworkConfigurationUpdate0) GetIpv6gatewayOk() (*string, bool)`

GetIpv6gatewayOk returns a tuple with the Ipv6gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6gateway

`func (o *NetworkConfigurationUpdate0) SetIpv6gateway(v string)`

SetIpv6gateway sets Ipv6gateway field to given value.

### HasIpv6gateway

`func (o *NetworkConfigurationUpdate0) HasIpv6gateway() bool`

HasIpv6gateway returns a boolean if a field has been set.

### GetNameserver1

`func (o *NetworkConfigurationUpdate0) GetNameserver1() string`

GetNameserver1 returns the Nameserver1 field if non-nil, zero value otherwise.

### GetNameserver1Ok

`func (o *NetworkConfigurationUpdate0) GetNameserver1Ok() (*string, bool)`

GetNameserver1Ok returns a tuple with the Nameserver1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver1

`func (o *NetworkConfigurationUpdate0) SetNameserver1(v string)`

SetNameserver1 sets Nameserver1 field to given value.

### HasNameserver1

`func (o *NetworkConfigurationUpdate0) HasNameserver1() bool`

HasNameserver1 returns a boolean if a field has been set.

### GetNameserver2

`func (o *NetworkConfigurationUpdate0) GetNameserver2() string`

GetNameserver2 returns the Nameserver2 field if non-nil, zero value otherwise.

### GetNameserver2Ok

`func (o *NetworkConfigurationUpdate0) GetNameserver2Ok() (*string, bool)`

GetNameserver2Ok returns a tuple with the Nameserver2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver2

`func (o *NetworkConfigurationUpdate0) SetNameserver2(v string)`

SetNameserver2 sets Nameserver2 field to given value.

### HasNameserver2

`func (o *NetworkConfigurationUpdate0) HasNameserver2() bool`

HasNameserver2 returns a boolean if a field has been set.

### GetNameserver3

`func (o *NetworkConfigurationUpdate0) GetNameserver3() string`

GetNameserver3 returns the Nameserver3 field if non-nil, zero value otherwise.

### GetNameserver3Ok

`func (o *NetworkConfigurationUpdate0) GetNameserver3Ok() (*string, bool)`

GetNameserver3Ok returns a tuple with the Nameserver3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver3

`func (o *NetworkConfigurationUpdate0) SetNameserver3(v string)`

SetNameserver3 sets Nameserver3 field to given value.

### HasNameserver3

`func (o *NetworkConfigurationUpdate0) HasNameserver3() bool`

HasNameserver3 returns a boolean if a field has been set.

### GetHttpproxy

`func (o *NetworkConfigurationUpdate0) GetHttpproxy() string`

GetHttpproxy returns the Httpproxy field if non-nil, zero value otherwise.

### GetHttpproxyOk

`func (o *NetworkConfigurationUpdate0) GetHttpproxyOk() (*string, bool)`

GetHttpproxyOk returns a tuple with the Httpproxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpproxy

`func (o *NetworkConfigurationUpdate0) SetHttpproxy(v string)`

SetHttpproxy sets Httpproxy field to given value.

### HasHttpproxy

`func (o *NetworkConfigurationUpdate0) HasHttpproxy() bool`

HasHttpproxy returns a boolean if a field has been set.

### GetNetwaitEnabled

`func (o *NetworkConfigurationUpdate0) GetNetwaitEnabled() bool`

GetNetwaitEnabled returns the NetwaitEnabled field if non-nil, zero value otherwise.

### GetNetwaitEnabledOk

`func (o *NetworkConfigurationUpdate0) GetNetwaitEnabledOk() (*bool, bool)`

GetNetwaitEnabledOk returns a tuple with the NetwaitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwaitEnabled

`func (o *NetworkConfigurationUpdate0) SetNetwaitEnabled(v bool)`

SetNetwaitEnabled sets NetwaitEnabled field to given value.

### HasNetwaitEnabled

`func (o *NetworkConfigurationUpdate0) HasNetwaitEnabled() bool`

HasNetwaitEnabled returns a boolean if a field has been set.

### GetNetwaitIp

`func (o *NetworkConfigurationUpdate0) GetNetwaitIp() []string`

GetNetwaitIp returns the NetwaitIp field if non-nil, zero value otherwise.

### GetNetwaitIpOk

`func (o *NetworkConfigurationUpdate0) GetNetwaitIpOk() (*[]string, bool)`

GetNetwaitIpOk returns a tuple with the NetwaitIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwaitIp

`func (o *NetworkConfigurationUpdate0) SetNetwaitIp(v []string)`

SetNetwaitIp sets NetwaitIp field to given value.

### HasNetwaitIp

`func (o *NetworkConfigurationUpdate0) HasNetwaitIp() bool`

HasNetwaitIp returns a boolean if a field has been set.

### GetHosts

`func (o *NetworkConfigurationUpdate0) GetHosts() string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *NetworkConfigurationUpdate0) GetHostsOk() (*string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *NetworkConfigurationUpdate0) SetHosts(v string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *NetworkConfigurationUpdate0) HasHosts() bool`

HasHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


