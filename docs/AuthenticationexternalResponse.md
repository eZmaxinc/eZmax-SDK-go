# AuthenticationexternalResponse

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

### NewAuthenticationexternalResponse

`func NewAuthenticationexternalResponse(pkiAuthenticationexternalID int32, sAuthenticationexternalDescription string, eAuthenticationexternalType FieldEAuthenticationexternalType, objAudit CommonAudit, ) *AuthenticationexternalResponse`

NewAuthenticationexternalResponse instantiates a new AuthenticationexternalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationexternalResponseWithDefaults

`func NewAuthenticationexternalResponseWithDefaults() *AuthenticationexternalResponse`

NewAuthenticationexternalResponseWithDefaults instantiates a new AuthenticationexternalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAuthenticationexternalID

`func (o *AuthenticationexternalResponse) GetPkiAuthenticationexternalID() int32`

GetPkiAuthenticationexternalID returns the PkiAuthenticationexternalID field if non-nil, zero value otherwise.

### GetPkiAuthenticationexternalIDOk

`func (o *AuthenticationexternalResponse) GetPkiAuthenticationexternalIDOk() (*int32, bool)`

GetPkiAuthenticationexternalIDOk returns a tuple with the PkiAuthenticationexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAuthenticationexternalID

`func (o *AuthenticationexternalResponse) SetPkiAuthenticationexternalID(v int32)`

SetPkiAuthenticationexternalID sets PkiAuthenticationexternalID field to given value.


### GetSAuthenticationexternalDescription

`func (o *AuthenticationexternalResponse) GetSAuthenticationexternalDescription() string`

GetSAuthenticationexternalDescription returns the SAuthenticationexternalDescription field if non-nil, zero value otherwise.

### GetSAuthenticationexternalDescriptionOk

`func (o *AuthenticationexternalResponse) GetSAuthenticationexternalDescriptionOk() (*string, bool)`

GetSAuthenticationexternalDescriptionOk returns a tuple with the SAuthenticationexternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAuthenticationexternalDescription

`func (o *AuthenticationexternalResponse) SetSAuthenticationexternalDescription(v string)`

SetSAuthenticationexternalDescription sets SAuthenticationexternalDescription field to given value.


### GetEAuthenticationexternalType

`func (o *AuthenticationexternalResponse) GetEAuthenticationexternalType() FieldEAuthenticationexternalType`

GetEAuthenticationexternalType returns the EAuthenticationexternalType field if non-nil, zero value otherwise.

### GetEAuthenticationexternalTypeOk

`func (o *AuthenticationexternalResponse) GetEAuthenticationexternalTypeOk() (*FieldEAuthenticationexternalType, bool)`

GetEAuthenticationexternalTypeOk returns a tuple with the EAuthenticationexternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAuthenticationexternalType

`func (o *AuthenticationexternalResponse) SetEAuthenticationexternalType(v FieldEAuthenticationexternalType)`

SetEAuthenticationexternalType sets EAuthenticationexternalType field to given value.


### GetBAuthenticationexternalConnected

`func (o *AuthenticationexternalResponse) GetBAuthenticationexternalConnected() bool`

GetBAuthenticationexternalConnected returns the BAuthenticationexternalConnected field if non-nil, zero value otherwise.

### GetBAuthenticationexternalConnectedOk

`func (o *AuthenticationexternalResponse) GetBAuthenticationexternalConnectedOk() (*bool, bool)`

GetBAuthenticationexternalConnectedOk returns a tuple with the BAuthenticationexternalConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAuthenticationexternalConnected

`func (o *AuthenticationexternalResponse) SetBAuthenticationexternalConnected(v bool)`

SetBAuthenticationexternalConnected sets BAuthenticationexternalConnected field to given value.

### HasBAuthenticationexternalConnected

`func (o *AuthenticationexternalResponse) HasBAuthenticationexternalConnected() bool`

HasBAuthenticationexternalConnected returns a boolean if a field has been set.

### GetSAuthenticationexternalAuthorizationurl

`func (o *AuthenticationexternalResponse) GetSAuthenticationexternalAuthorizationurl() string`

GetSAuthenticationexternalAuthorizationurl returns the SAuthenticationexternalAuthorizationurl field if non-nil, zero value otherwise.

### GetSAuthenticationexternalAuthorizationurlOk

`func (o *AuthenticationexternalResponse) GetSAuthenticationexternalAuthorizationurlOk() (*string, bool)`

GetSAuthenticationexternalAuthorizationurlOk returns a tuple with the SAuthenticationexternalAuthorizationurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAuthenticationexternalAuthorizationurl

`func (o *AuthenticationexternalResponse) SetSAuthenticationexternalAuthorizationurl(v string)`

SetSAuthenticationexternalAuthorizationurl sets SAuthenticationexternalAuthorizationurl field to given value.

### HasSAuthenticationexternalAuthorizationurl

`func (o *AuthenticationexternalResponse) HasSAuthenticationexternalAuthorizationurl() bool`

HasSAuthenticationexternalAuthorizationurl returns a boolean if a field has been set.

### GetObjAudit

`func (o *AuthenticationexternalResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *AuthenticationexternalResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *AuthenticationexternalResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


