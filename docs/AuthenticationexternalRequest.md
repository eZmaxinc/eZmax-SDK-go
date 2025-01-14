# AuthenticationexternalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiAuthenticationexternalID** | Pointer to **int32** | The unique ID of the Authenticationexternal | [optional] 
**SAuthenticationexternalDescription** | **string** | The description of the Authenticationexternal | 
**EAuthenticationexternalType** | [**FieldEAuthenticationexternalType**](FieldEAuthenticationexternalType.md) |  | 

## Methods

### NewAuthenticationexternalRequest

`func NewAuthenticationexternalRequest(sAuthenticationexternalDescription string, eAuthenticationexternalType FieldEAuthenticationexternalType, ) *AuthenticationexternalRequest`

NewAuthenticationexternalRequest instantiates a new AuthenticationexternalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationexternalRequestWithDefaults

`func NewAuthenticationexternalRequestWithDefaults() *AuthenticationexternalRequest`

NewAuthenticationexternalRequestWithDefaults instantiates a new AuthenticationexternalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAuthenticationexternalID

`func (o *AuthenticationexternalRequest) GetPkiAuthenticationexternalID() int32`

GetPkiAuthenticationexternalID returns the PkiAuthenticationexternalID field if non-nil, zero value otherwise.

### GetPkiAuthenticationexternalIDOk

`func (o *AuthenticationexternalRequest) GetPkiAuthenticationexternalIDOk() (*int32, bool)`

GetPkiAuthenticationexternalIDOk returns a tuple with the PkiAuthenticationexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAuthenticationexternalID

`func (o *AuthenticationexternalRequest) SetPkiAuthenticationexternalID(v int32)`

SetPkiAuthenticationexternalID sets PkiAuthenticationexternalID field to given value.

### HasPkiAuthenticationexternalID

`func (o *AuthenticationexternalRequest) HasPkiAuthenticationexternalID() bool`

HasPkiAuthenticationexternalID returns a boolean if a field has been set.

### GetSAuthenticationexternalDescription

`func (o *AuthenticationexternalRequest) GetSAuthenticationexternalDescription() string`

GetSAuthenticationexternalDescription returns the SAuthenticationexternalDescription field if non-nil, zero value otherwise.

### GetSAuthenticationexternalDescriptionOk

`func (o *AuthenticationexternalRequest) GetSAuthenticationexternalDescriptionOk() (*string, bool)`

GetSAuthenticationexternalDescriptionOk returns a tuple with the SAuthenticationexternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAuthenticationexternalDescription

`func (o *AuthenticationexternalRequest) SetSAuthenticationexternalDescription(v string)`

SetSAuthenticationexternalDescription sets SAuthenticationexternalDescription field to given value.


### GetEAuthenticationexternalType

`func (o *AuthenticationexternalRequest) GetEAuthenticationexternalType() FieldEAuthenticationexternalType`

GetEAuthenticationexternalType returns the EAuthenticationexternalType field if non-nil, zero value otherwise.

### GetEAuthenticationexternalTypeOk

`func (o *AuthenticationexternalRequest) GetEAuthenticationexternalTypeOk() (*FieldEAuthenticationexternalType, bool)`

GetEAuthenticationexternalTypeOk returns a tuple with the EAuthenticationexternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAuthenticationexternalType

`func (o *AuthenticationexternalRequest) SetEAuthenticationexternalType(v FieldEAuthenticationexternalType)`

SetEAuthenticationexternalType sets EAuthenticationexternalType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


