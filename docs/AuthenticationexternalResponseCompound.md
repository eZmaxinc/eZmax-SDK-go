# AuthenticationexternalResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiAuthenticationexternalID** | **int32** | The unique ID of the Authenticationexternal | 
**SAuthenticationexternalDescription** | **string** | The description of the Authenticationexternal | 
**EAuthenticationexternalType** | [**FieldEAuthenticationexternalType**](FieldEAuthenticationexternalType.md) |  | 
**BAuthenticationexternalConnected** | Pointer to **bool** | Whether the Authenticationexternal has been connected or not | [optional] 
**SAuthenticationexternalAuthorizationurl** | Pointer to **string** | The url to authorize the Authenticationexternal | [optional] 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewAuthenticationexternalResponseCompound

`func NewAuthenticationexternalResponseCompound(pkiAuthenticationexternalID int32, sAuthenticationexternalDescription string, eAuthenticationexternalType FieldEAuthenticationexternalType, objAudit CommonAudit, ) *AuthenticationexternalResponseCompound`

NewAuthenticationexternalResponseCompound instantiates a new AuthenticationexternalResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationexternalResponseCompoundWithDefaults

`func NewAuthenticationexternalResponseCompoundWithDefaults() *AuthenticationexternalResponseCompound`

NewAuthenticationexternalResponseCompoundWithDefaults instantiates a new AuthenticationexternalResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAuthenticationexternalID

`func (o *AuthenticationexternalResponseCompound) GetPkiAuthenticationexternalID() int32`

GetPkiAuthenticationexternalID returns the PkiAuthenticationexternalID field if non-nil, zero value otherwise.

### GetPkiAuthenticationexternalIDOk

`func (o *AuthenticationexternalResponseCompound) GetPkiAuthenticationexternalIDOk() (*int32, bool)`

GetPkiAuthenticationexternalIDOk returns a tuple with the PkiAuthenticationexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAuthenticationexternalID

`func (o *AuthenticationexternalResponseCompound) SetPkiAuthenticationexternalID(v int32)`

SetPkiAuthenticationexternalID sets PkiAuthenticationexternalID field to given value.


### GetSAuthenticationexternalDescription

`func (o *AuthenticationexternalResponseCompound) GetSAuthenticationexternalDescription() string`

GetSAuthenticationexternalDescription returns the SAuthenticationexternalDescription field if non-nil, zero value otherwise.

### GetSAuthenticationexternalDescriptionOk

`func (o *AuthenticationexternalResponseCompound) GetSAuthenticationexternalDescriptionOk() (*string, bool)`

GetSAuthenticationexternalDescriptionOk returns a tuple with the SAuthenticationexternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAuthenticationexternalDescription

`func (o *AuthenticationexternalResponseCompound) SetSAuthenticationexternalDescription(v string)`

SetSAuthenticationexternalDescription sets SAuthenticationexternalDescription field to given value.


### GetEAuthenticationexternalType

`func (o *AuthenticationexternalResponseCompound) GetEAuthenticationexternalType() FieldEAuthenticationexternalType`

GetEAuthenticationexternalType returns the EAuthenticationexternalType field if non-nil, zero value otherwise.

### GetEAuthenticationexternalTypeOk

`func (o *AuthenticationexternalResponseCompound) GetEAuthenticationexternalTypeOk() (*FieldEAuthenticationexternalType, bool)`

GetEAuthenticationexternalTypeOk returns a tuple with the EAuthenticationexternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAuthenticationexternalType

`func (o *AuthenticationexternalResponseCompound) SetEAuthenticationexternalType(v FieldEAuthenticationexternalType)`

SetEAuthenticationexternalType sets EAuthenticationexternalType field to given value.


### GetBAuthenticationexternalConnected

`func (o *AuthenticationexternalResponseCompound) GetBAuthenticationexternalConnected() bool`

GetBAuthenticationexternalConnected returns the BAuthenticationexternalConnected field if non-nil, zero value otherwise.

### GetBAuthenticationexternalConnectedOk

`func (o *AuthenticationexternalResponseCompound) GetBAuthenticationexternalConnectedOk() (*bool, bool)`

GetBAuthenticationexternalConnectedOk returns a tuple with the BAuthenticationexternalConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAuthenticationexternalConnected

`func (o *AuthenticationexternalResponseCompound) SetBAuthenticationexternalConnected(v bool)`

SetBAuthenticationexternalConnected sets BAuthenticationexternalConnected field to given value.

### HasBAuthenticationexternalConnected

`func (o *AuthenticationexternalResponseCompound) HasBAuthenticationexternalConnected() bool`

HasBAuthenticationexternalConnected returns a boolean if a field has been set.

### GetSAuthenticationexternalAuthorizationurl

`func (o *AuthenticationexternalResponseCompound) GetSAuthenticationexternalAuthorizationurl() string`

GetSAuthenticationexternalAuthorizationurl returns the SAuthenticationexternalAuthorizationurl field if non-nil, zero value otherwise.

### GetSAuthenticationexternalAuthorizationurlOk

`func (o *AuthenticationexternalResponseCompound) GetSAuthenticationexternalAuthorizationurlOk() (*string, bool)`

GetSAuthenticationexternalAuthorizationurlOk returns a tuple with the SAuthenticationexternalAuthorizationurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAuthenticationexternalAuthorizationurl

`func (o *AuthenticationexternalResponseCompound) SetSAuthenticationexternalAuthorizationurl(v string)`

SetSAuthenticationexternalAuthorizationurl sets SAuthenticationexternalAuthorizationurl field to given value.

### HasSAuthenticationexternalAuthorizationurl

`func (o *AuthenticationexternalResponseCompound) HasSAuthenticationexternalAuthorizationurl() bool`

HasSAuthenticationexternalAuthorizationurl returns a boolean if a field has been set.

### GetObjAudit

`func (o *AuthenticationexternalResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *AuthenticationexternalResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *AuthenticationexternalResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


