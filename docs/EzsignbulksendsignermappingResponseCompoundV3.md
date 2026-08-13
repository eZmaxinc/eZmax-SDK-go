# EzsignbulksendsignermappingResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksendsignermappingID** | **int32** | The unique ID of the Ezsignbulksendsignermapping | 
**FkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**EEzsignbulksendsignermappingRole** | Pointer to [**FieldEEzsignbulksendsignermappingRole**](FieldEEzsignbulksendsignermappingRole.md) |  | [optional] 
**SEzsignbulksendsignermappingDescription** | **string** | The description of the Ezsignbulksendsignermapping | 

## Methods

### NewEzsignbulksendsignermappingResponseCompoundV3

`func NewEzsignbulksendsignermappingResponseCompoundV3(pkiEzsignbulksendsignermappingID int32, fkiEzsignbulksendID int32, sEzsignbulksendsignermappingDescription string, ) *EzsignbulksendsignermappingResponseCompoundV3`

NewEzsignbulksendsignermappingResponseCompoundV3 instantiates a new EzsignbulksendsignermappingResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendsignermappingResponseCompoundV3WithDefaults

`func NewEzsignbulksendsignermappingResponseCompoundV3WithDefaults() *EzsignbulksendsignermappingResponseCompoundV3`

NewEzsignbulksendsignermappingResponseCompoundV3WithDefaults instantiates a new EzsignbulksendsignermappingResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksendsignermappingID

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetPkiEzsignbulksendsignermappingID() int32`

GetPkiEzsignbulksendsignermappingID returns the PkiEzsignbulksendsignermappingID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksendsignermappingIDOk

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetPkiEzsignbulksendsignermappingIDOk() (*int32, bool)`

GetPkiEzsignbulksendsignermappingIDOk returns a tuple with the PkiEzsignbulksendsignermappingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksendsignermappingID

`func (o *EzsignbulksendsignermappingResponseCompoundV3) SetPkiEzsignbulksendsignermappingID(v int32)`

SetPkiEzsignbulksendsignermappingID sets PkiEzsignbulksendsignermappingID field to given value.


### GetFkiEzsignbulksendID

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetFkiEzsignbulksendID() int32`

GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetFkiEzsignbulksendIDOk

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetFkiEzsignbulksendIDOk() (*int32, bool)`

GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignbulksendID

`func (o *EzsignbulksendsignermappingResponseCompoundV3) SetFkiEzsignbulksendID(v int32)`

SetFkiEzsignbulksendID sets FkiEzsignbulksendID field to given value.


### GetFkiUserID

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsignbulksendsignermappingResponseCompoundV3) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsignbulksendsignermappingResponseCompoundV3) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetEEzsignbulksendsignermappingRole

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetEEzsignbulksendsignermappingRole() FieldEEzsignbulksendsignermappingRole`

GetEEzsignbulksendsignermappingRole returns the EEzsignbulksendsignermappingRole field if non-nil, zero value otherwise.

### GetEEzsignbulksendsignermappingRoleOk

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetEEzsignbulksendsignermappingRoleOk() (*FieldEEzsignbulksendsignermappingRole, bool)`

GetEEzsignbulksendsignermappingRoleOk returns a tuple with the EEzsignbulksendsignermappingRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignbulksendsignermappingRole

`func (o *EzsignbulksendsignermappingResponseCompoundV3) SetEEzsignbulksendsignermappingRole(v FieldEEzsignbulksendsignermappingRole)`

SetEEzsignbulksendsignermappingRole sets EEzsignbulksendsignermappingRole field to given value.

### HasEEzsignbulksendsignermappingRole

`func (o *EzsignbulksendsignermappingResponseCompoundV3) HasEEzsignbulksendsignermappingRole() bool`

HasEEzsignbulksendsignermappingRole returns a boolean if a field has been set.

### GetSEzsignbulksendsignermappingDescription

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetSEzsignbulksendsignermappingDescription() string`

GetSEzsignbulksendsignermappingDescription returns the SEzsignbulksendsignermappingDescription field if non-nil, zero value otherwise.

### GetSEzsignbulksendsignermappingDescriptionOk

`func (o *EzsignbulksendsignermappingResponseCompoundV3) GetSEzsignbulksendsignermappingDescriptionOk() (*string, bool)`

GetSEzsignbulksendsignermappingDescriptionOk returns a tuple with the SEzsignbulksendsignermappingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignbulksendsignermappingDescription

`func (o *EzsignbulksendsignermappingResponseCompoundV3) SetSEzsignbulksendsignermappingDescription(v string)`

SetSEzsignbulksendsignermappingDescription sets SEzsignbulksendsignermappingDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


