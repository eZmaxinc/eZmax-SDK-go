# AuthenticationexternalListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiAuthenticationexternalID** | **int32** | The unique ID of the Authenticationexternal | 
**SAuthenticationexternalDescription** | **string** | The description of the Authenticationexternal | 
**EAuthenticationexternalType** | [**FieldEAuthenticationexternalType**](FieldEAuthenticationexternalType.md) |  | 
**BAuthenticationexternalConnected** | **bool** | Whether the Authenticationexternal has been connected or not | 

## Methods

### NewAuthenticationexternalListElement

`func NewAuthenticationexternalListElement(pkiAuthenticationexternalID int32, sAuthenticationexternalDescription string, eAuthenticationexternalType FieldEAuthenticationexternalType, bAuthenticationexternalConnected bool, ) *AuthenticationexternalListElement`

NewAuthenticationexternalListElement instantiates a new AuthenticationexternalListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationexternalListElementWithDefaults

`func NewAuthenticationexternalListElementWithDefaults() *AuthenticationexternalListElement`

NewAuthenticationexternalListElementWithDefaults instantiates a new AuthenticationexternalListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAuthenticationexternalID

`func (o *AuthenticationexternalListElement) GetPkiAuthenticationexternalID() int32`

GetPkiAuthenticationexternalID returns the PkiAuthenticationexternalID field if non-nil, zero value otherwise.

### GetPkiAuthenticationexternalIDOk

`func (o *AuthenticationexternalListElement) GetPkiAuthenticationexternalIDOk() (*int32, bool)`

GetPkiAuthenticationexternalIDOk returns a tuple with the PkiAuthenticationexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAuthenticationexternalID

`func (o *AuthenticationexternalListElement) SetPkiAuthenticationexternalID(v int32)`

SetPkiAuthenticationexternalID sets PkiAuthenticationexternalID field to given value.


### GetSAuthenticationexternalDescription

`func (o *AuthenticationexternalListElement) GetSAuthenticationexternalDescription() string`

GetSAuthenticationexternalDescription returns the SAuthenticationexternalDescription field if non-nil, zero value otherwise.

### GetSAuthenticationexternalDescriptionOk

`func (o *AuthenticationexternalListElement) GetSAuthenticationexternalDescriptionOk() (*string, bool)`

GetSAuthenticationexternalDescriptionOk returns a tuple with the SAuthenticationexternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAuthenticationexternalDescription

`func (o *AuthenticationexternalListElement) SetSAuthenticationexternalDescription(v string)`

SetSAuthenticationexternalDescription sets SAuthenticationexternalDescription field to given value.


### GetEAuthenticationexternalType

`func (o *AuthenticationexternalListElement) GetEAuthenticationexternalType() FieldEAuthenticationexternalType`

GetEAuthenticationexternalType returns the EAuthenticationexternalType field if non-nil, zero value otherwise.

### GetEAuthenticationexternalTypeOk

`func (o *AuthenticationexternalListElement) GetEAuthenticationexternalTypeOk() (*FieldEAuthenticationexternalType, bool)`

GetEAuthenticationexternalTypeOk returns a tuple with the EAuthenticationexternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAuthenticationexternalType

`func (o *AuthenticationexternalListElement) SetEAuthenticationexternalType(v FieldEAuthenticationexternalType)`

SetEAuthenticationexternalType sets EAuthenticationexternalType field to given value.


### GetBAuthenticationexternalConnected

`func (o *AuthenticationexternalListElement) GetBAuthenticationexternalConnected() bool`

GetBAuthenticationexternalConnected returns the BAuthenticationexternalConnected field if non-nil, zero value otherwise.

### GetBAuthenticationexternalConnectedOk

`func (o *AuthenticationexternalListElement) GetBAuthenticationexternalConnectedOk() (*bool, bool)`

GetBAuthenticationexternalConnectedOk returns a tuple with the BAuthenticationexternalConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAuthenticationexternalConnected

`func (o *AuthenticationexternalListElement) SetBAuthenticationexternalConnected(v bool)`

SetBAuthenticationexternalConnected sets BAuthenticationexternalConnected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


