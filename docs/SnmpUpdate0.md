# SnmpUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to **string** |  | [optional] 
**Traps** | Pointer to **bool** |  | [optional] 
**V3** | Pointer to **bool** |  | [optional] 
**Community** | Pointer to **string** |  | [optional] 
**V3Username** | Pointer to **string** |  | [optional] 
**V3Authtype** | Pointer to **string** |  | [optional] 
**V3Password** | Pointer to **string** |  | [optional] 
**V3Privproto** | Pointer to **NullableString** |  | [optional] 
**V3Privpassphrase** | Pointer to **string** |  | [optional] 
**Loglevel** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 
**Zilstat** | Pointer to **bool** |  | [optional] 
**Iftop** | Pointer to **bool** |  | [optional] 

## Methods

### NewSnmpUpdate0

`func NewSnmpUpdate0() *SnmpUpdate0`

NewSnmpUpdate0 instantiates a new SnmpUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpUpdate0WithDefaults

`func NewSnmpUpdate0WithDefaults() *SnmpUpdate0`

NewSnmpUpdate0WithDefaults instantiates a new SnmpUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SnmpUpdate0) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SnmpUpdate0) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SnmpUpdate0) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SnmpUpdate0) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetContact

`func (o *SnmpUpdate0) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SnmpUpdate0) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SnmpUpdate0) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SnmpUpdate0) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetTraps

`func (o *SnmpUpdate0) GetTraps() bool`

GetTraps returns the Traps field if non-nil, zero value otherwise.

### GetTrapsOk

`func (o *SnmpUpdate0) GetTrapsOk() (*bool, bool)`

GetTrapsOk returns a tuple with the Traps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraps

`func (o *SnmpUpdate0) SetTraps(v bool)`

SetTraps sets Traps field to given value.

### HasTraps

`func (o *SnmpUpdate0) HasTraps() bool`

HasTraps returns a boolean if a field has been set.

### GetV3

`func (o *SnmpUpdate0) GetV3() bool`

GetV3 returns the V3 field if non-nil, zero value otherwise.

### GetV3Ok

`func (o *SnmpUpdate0) GetV3Ok() (*bool, bool)`

GetV3Ok returns a tuple with the V3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3

`func (o *SnmpUpdate0) SetV3(v bool)`

SetV3 sets V3 field to given value.

### HasV3

`func (o *SnmpUpdate0) HasV3() bool`

HasV3 returns a boolean if a field has been set.

### GetCommunity

`func (o *SnmpUpdate0) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *SnmpUpdate0) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *SnmpUpdate0) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *SnmpUpdate0) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetV3Username

`func (o *SnmpUpdate0) GetV3Username() string`

GetV3Username returns the V3Username field if non-nil, zero value otherwise.

### GetV3UsernameOk

`func (o *SnmpUpdate0) GetV3UsernameOk() (*string, bool)`

GetV3UsernameOk returns a tuple with the V3Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Username

`func (o *SnmpUpdate0) SetV3Username(v string)`

SetV3Username sets V3Username field to given value.

### HasV3Username

`func (o *SnmpUpdate0) HasV3Username() bool`

HasV3Username returns a boolean if a field has been set.

### GetV3Authtype

`func (o *SnmpUpdate0) GetV3Authtype() string`

GetV3Authtype returns the V3Authtype field if non-nil, zero value otherwise.

### GetV3AuthtypeOk

`func (o *SnmpUpdate0) GetV3AuthtypeOk() (*string, bool)`

GetV3AuthtypeOk returns a tuple with the V3Authtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Authtype

`func (o *SnmpUpdate0) SetV3Authtype(v string)`

SetV3Authtype sets V3Authtype field to given value.

### HasV3Authtype

`func (o *SnmpUpdate0) HasV3Authtype() bool`

HasV3Authtype returns a boolean if a field has been set.

### GetV3Password

`func (o *SnmpUpdate0) GetV3Password() string`

GetV3Password returns the V3Password field if non-nil, zero value otherwise.

### GetV3PasswordOk

`func (o *SnmpUpdate0) GetV3PasswordOk() (*string, bool)`

GetV3PasswordOk returns a tuple with the V3Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Password

`func (o *SnmpUpdate0) SetV3Password(v string)`

SetV3Password sets V3Password field to given value.

### HasV3Password

`func (o *SnmpUpdate0) HasV3Password() bool`

HasV3Password returns a boolean if a field has been set.

### GetV3Privproto

`func (o *SnmpUpdate0) GetV3Privproto() string`

GetV3Privproto returns the V3Privproto field if non-nil, zero value otherwise.

### GetV3PrivprotoOk

`func (o *SnmpUpdate0) GetV3PrivprotoOk() (*string, bool)`

GetV3PrivprotoOk returns a tuple with the V3Privproto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Privproto

`func (o *SnmpUpdate0) SetV3Privproto(v string)`

SetV3Privproto sets V3Privproto field to given value.

### HasV3Privproto

`func (o *SnmpUpdate0) HasV3Privproto() bool`

HasV3Privproto returns a boolean if a field has been set.

### SetV3PrivprotoNil

`func (o *SnmpUpdate0) SetV3PrivprotoNil(b bool)`

 SetV3PrivprotoNil sets the value for V3Privproto to be an explicit nil

### UnsetV3Privproto
`func (o *SnmpUpdate0) UnsetV3Privproto()`

UnsetV3Privproto ensures that no value is present for V3Privproto, not even an explicit nil
### GetV3Privpassphrase

`func (o *SnmpUpdate0) GetV3Privpassphrase() string`

GetV3Privpassphrase returns the V3Privpassphrase field if non-nil, zero value otherwise.

### GetV3PrivpassphraseOk

`func (o *SnmpUpdate0) GetV3PrivpassphraseOk() (*string, bool)`

GetV3PrivpassphraseOk returns a tuple with the V3Privpassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Privpassphrase

`func (o *SnmpUpdate0) SetV3Privpassphrase(v string)`

SetV3Privpassphrase sets V3Privpassphrase field to given value.

### HasV3Privpassphrase

`func (o *SnmpUpdate0) HasV3Privpassphrase() bool`

HasV3Privpassphrase returns a boolean if a field has been set.

### GetLoglevel

`func (o *SnmpUpdate0) GetLoglevel() int32`

GetLoglevel returns the Loglevel field if non-nil, zero value otherwise.

### GetLoglevelOk

`func (o *SnmpUpdate0) GetLoglevelOk() (*int32, bool)`

GetLoglevelOk returns a tuple with the Loglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoglevel

`func (o *SnmpUpdate0) SetLoglevel(v int32)`

SetLoglevel sets Loglevel field to given value.

### HasLoglevel

`func (o *SnmpUpdate0) HasLoglevel() bool`

HasLoglevel returns a boolean if a field has been set.

### GetOptions

`func (o *SnmpUpdate0) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SnmpUpdate0) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SnmpUpdate0) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SnmpUpdate0) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetZilstat

`func (o *SnmpUpdate0) GetZilstat() bool`

GetZilstat returns the Zilstat field if non-nil, zero value otherwise.

### GetZilstatOk

`func (o *SnmpUpdate0) GetZilstatOk() (*bool, bool)`

GetZilstatOk returns a tuple with the Zilstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZilstat

`func (o *SnmpUpdate0) SetZilstat(v bool)`

SetZilstat sets Zilstat field to given value.

### HasZilstat

`func (o *SnmpUpdate0) HasZilstat() bool`

HasZilstat returns a boolean if a field has been set.

### GetIftop

`func (o *SnmpUpdate0) GetIftop() bool`

GetIftop returns the Iftop field if non-nil, zero value otherwise.

### GetIftopOk

`func (o *SnmpUpdate0) GetIftopOk() (*bool, bool)`

GetIftopOk returns a tuple with the Iftop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIftop

`func (o *SnmpUpdate0) SetIftop(v bool)`

SetIftop sets Iftop field to given value.

### HasIftop

`func (o *SnmpUpdate0) HasIftop() bool`

HasIftop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


