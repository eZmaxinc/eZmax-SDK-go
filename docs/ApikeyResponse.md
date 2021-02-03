# ApikeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjApikeyDescription** | [**MultilingualApikeyDescription**](Multilingual-ApikeyDescription.md) |  | 
**SComputedToken** | Pointer to **string** | The secret token for the API key.  This will be returned only on creation. | [optional] 
**PkiApikeyID** | **int32** | The unique ID of the Apikey | 
**ObjAudit** | [**CommonAudit**](Common-Audit.md) |  | 

## Methods

### NewApikeyResponse

`func NewApikeyResponse(objApikeyDescription MultilingualApikeyDescription, pkiApikeyID int32, objAudit CommonAudit, ) *ApikeyResponse`

NewApikeyResponse instantiates a new ApikeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApikeyResponseWithDefaults

`func NewApikeyResponseWithDefaults() *ApikeyResponse`

NewApikeyResponseWithDefaults instantiates a new ApikeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjApikeyDescription

`func (o *ApikeyResponse) GetObjApikeyDescription() MultilingualApikeyDescription`

GetObjApikeyDescription returns the ObjApikeyDescription field if non-nil, zero value otherwise.

### GetObjApikeyDescriptionOk

`func (o *ApikeyResponse) GetObjApikeyDescriptionOk() (*MultilingualApikeyDescription, bool)`

GetObjApikeyDescriptionOk returns a tuple with the ObjApikeyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjApikeyDescription

`func (o *ApikeyResponse) SetObjApikeyDescription(v MultilingualApikeyDescription)`

SetObjApikeyDescription sets ObjApikeyDescription field to given value.


### GetSComputedToken

`func (o *ApikeyResponse) GetSComputedToken() string`

GetSComputedToken returns the SComputedToken field if non-nil, zero value otherwise.

### GetSComputedTokenOk

`func (o *ApikeyResponse) GetSComputedTokenOk() (*string, bool)`

GetSComputedTokenOk returns a tuple with the SComputedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSComputedToken

`func (o *ApikeyResponse) SetSComputedToken(v string)`

SetSComputedToken sets SComputedToken field to given value.

### HasSComputedToken

`func (o *ApikeyResponse) HasSComputedToken() bool`

HasSComputedToken returns a boolean if a field has been set.

### GetPkiApikeyID

`func (o *ApikeyResponse) GetPkiApikeyID() int32`

GetPkiApikeyID returns the PkiApikeyID field if non-nil, zero value otherwise.

### GetPkiApikeyIDOk

`func (o *ApikeyResponse) GetPkiApikeyIDOk() (*int32, bool)`

GetPkiApikeyIDOk returns a tuple with the PkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiApikeyID

`func (o *ApikeyResponse) SetPkiApikeyID(v int32)`

SetPkiApikeyID sets PkiApikeyID field to given value.


### GetObjAudit

`func (o *ApikeyResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *ApikeyResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *ApikeyResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


