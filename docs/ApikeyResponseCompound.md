# ApikeyResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiApikeyID** | **int32** | The unique ID of the Apikey | 
**FkiUserID** | **int32** | The unique ID of the User | 
**ObjApikeyDescription** | [**MultilingualApikeyDescription**](MultilingualApikeyDescription.md) |  | 
**ObjContactName** | [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | 
**SApikeyApikey** | Pointer to **string** | The Apikey for the API key.  This will be hidden if we are not creating or regenerating the Apikey. | [optional] 
**SApikeySecret** | Pointer to **string** | The Secret for the API key.  This will be hidden if we are not creating or regenerating the Apikey. | [optional] 
**BApikeyIsactive** | **bool** | Whether the apikey is active or not | 
**BApikeyIssigned** | Pointer to **bool** | Whether the apikey is signed or not | [optional] 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewApikeyResponseCompound

`func NewApikeyResponseCompound(pkiApikeyID int32, fkiUserID int32, objApikeyDescription MultilingualApikeyDescription, objContactName CustomContactNameResponse, bApikeyIsactive bool, objAudit CommonAudit, ) *ApikeyResponseCompound`

NewApikeyResponseCompound instantiates a new ApikeyResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApikeyResponseCompoundWithDefaults

`func NewApikeyResponseCompoundWithDefaults() *ApikeyResponseCompound`

NewApikeyResponseCompoundWithDefaults instantiates a new ApikeyResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiApikeyID

`func (o *ApikeyResponseCompound) GetPkiApikeyID() int32`

GetPkiApikeyID returns the PkiApikeyID field if non-nil, zero value otherwise.

### GetPkiApikeyIDOk

`func (o *ApikeyResponseCompound) GetPkiApikeyIDOk() (*int32, bool)`

GetPkiApikeyIDOk returns a tuple with the PkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiApikeyID

`func (o *ApikeyResponseCompound) SetPkiApikeyID(v int32)`

SetPkiApikeyID sets PkiApikeyID field to given value.


### GetFkiUserID

`func (o *ApikeyResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ApikeyResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ApikeyResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetObjApikeyDescription

`func (o *ApikeyResponseCompound) GetObjApikeyDescription() MultilingualApikeyDescription`

GetObjApikeyDescription returns the ObjApikeyDescription field if non-nil, zero value otherwise.

### GetObjApikeyDescriptionOk

`func (o *ApikeyResponseCompound) GetObjApikeyDescriptionOk() (*MultilingualApikeyDescription, bool)`

GetObjApikeyDescriptionOk returns a tuple with the ObjApikeyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjApikeyDescription

`func (o *ApikeyResponseCompound) SetObjApikeyDescription(v MultilingualApikeyDescription)`

SetObjApikeyDescription sets ObjApikeyDescription field to given value.


### GetObjContactName

`func (o *ApikeyResponseCompound) GetObjContactName() CustomContactNameResponse`

GetObjContactName returns the ObjContactName field if non-nil, zero value otherwise.

### GetObjContactNameOk

`func (o *ApikeyResponseCompound) GetObjContactNameOk() (*CustomContactNameResponse, bool)`

GetObjContactNameOk returns a tuple with the ObjContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactName

`func (o *ApikeyResponseCompound) SetObjContactName(v CustomContactNameResponse)`

SetObjContactName sets ObjContactName field to given value.


### GetSApikeyApikey

`func (o *ApikeyResponseCompound) GetSApikeyApikey() string`

GetSApikeyApikey returns the SApikeyApikey field if non-nil, zero value otherwise.

### GetSApikeyApikeyOk

`func (o *ApikeyResponseCompound) GetSApikeyApikeyOk() (*string, bool)`

GetSApikeyApikeyOk returns a tuple with the SApikeyApikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSApikeyApikey

`func (o *ApikeyResponseCompound) SetSApikeyApikey(v string)`

SetSApikeyApikey sets SApikeyApikey field to given value.

### HasSApikeyApikey

`func (o *ApikeyResponseCompound) HasSApikeyApikey() bool`

HasSApikeyApikey returns a boolean if a field has been set.

### GetSApikeySecret

`func (o *ApikeyResponseCompound) GetSApikeySecret() string`

GetSApikeySecret returns the SApikeySecret field if non-nil, zero value otherwise.

### GetSApikeySecretOk

`func (o *ApikeyResponseCompound) GetSApikeySecretOk() (*string, bool)`

GetSApikeySecretOk returns a tuple with the SApikeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSApikeySecret

`func (o *ApikeyResponseCompound) SetSApikeySecret(v string)`

SetSApikeySecret sets SApikeySecret field to given value.

### HasSApikeySecret

`func (o *ApikeyResponseCompound) HasSApikeySecret() bool`

HasSApikeySecret returns a boolean if a field has been set.

### GetBApikeyIsactive

`func (o *ApikeyResponseCompound) GetBApikeyIsactive() bool`

GetBApikeyIsactive returns the BApikeyIsactive field if non-nil, zero value otherwise.

### GetBApikeyIsactiveOk

`func (o *ApikeyResponseCompound) GetBApikeyIsactiveOk() (*bool, bool)`

GetBApikeyIsactiveOk returns a tuple with the BApikeyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBApikeyIsactive

`func (o *ApikeyResponseCompound) SetBApikeyIsactive(v bool)`

SetBApikeyIsactive sets BApikeyIsactive field to given value.


### GetBApikeyIssigned

`func (o *ApikeyResponseCompound) GetBApikeyIssigned() bool`

GetBApikeyIssigned returns the BApikeyIssigned field if non-nil, zero value otherwise.

### GetBApikeyIssignedOk

`func (o *ApikeyResponseCompound) GetBApikeyIssignedOk() (*bool, bool)`

GetBApikeyIssignedOk returns a tuple with the BApikeyIssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBApikeyIssigned

`func (o *ApikeyResponseCompound) SetBApikeyIssigned(v bool)`

SetBApikeyIssigned sets BApikeyIssigned field to given value.

### HasBApikeyIssigned

`func (o *ApikeyResponseCompound) HasBApikeyIssigned() bool`

HasBApikeyIssigned returns a boolean if a field has been set.

### GetObjAudit

`func (o *ApikeyResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *ApikeyResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *ApikeyResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


