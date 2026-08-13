# EzsignbulksendsignermappingRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksendsignermappingID** | Pointer to **int32** | The unique ID of the Ezsignbulksendsignermapping | [optional] 
**FkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**EEzsignbulksendsignermappingRole** | Pointer to [**FieldEEzsignbulksendsignermappingRole**](FieldEEzsignbulksendsignermappingRole.md) |  | [optional] 
**SEzsignbulksendsignermappingDescription** | **string** | The description of the Ezsignbulksendsignermapping | 

## Methods

### NewEzsignbulksendsignermappingRequestV2

`func NewEzsignbulksendsignermappingRequestV2(fkiEzsignbulksendID int32, sEzsignbulksendsignermappingDescription string, ) *EzsignbulksendsignermappingRequestV2`

NewEzsignbulksendsignermappingRequestV2 instantiates a new EzsignbulksendsignermappingRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendsignermappingRequestV2WithDefaults

`func NewEzsignbulksendsignermappingRequestV2WithDefaults() *EzsignbulksendsignermappingRequestV2`

NewEzsignbulksendsignermappingRequestV2WithDefaults instantiates a new EzsignbulksendsignermappingRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksendsignermappingID

`func (o *EzsignbulksendsignermappingRequestV2) GetPkiEzsignbulksendsignermappingID() int32`

GetPkiEzsignbulksendsignermappingID returns the PkiEzsignbulksendsignermappingID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksendsignermappingIDOk

`func (o *EzsignbulksendsignermappingRequestV2) GetPkiEzsignbulksendsignermappingIDOk() (*int32, bool)`

GetPkiEzsignbulksendsignermappingIDOk returns a tuple with the PkiEzsignbulksendsignermappingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksendsignermappingID

`func (o *EzsignbulksendsignermappingRequestV2) SetPkiEzsignbulksendsignermappingID(v int32)`

SetPkiEzsignbulksendsignermappingID sets PkiEzsignbulksendsignermappingID field to given value.

### HasPkiEzsignbulksendsignermappingID

`func (o *EzsignbulksendsignermappingRequestV2) HasPkiEzsignbulksendsignermappingID() bool`

HasPkiEzsignbulksendsignermappingID returns a boolean if a field has been set.

### GetFkiEzsignbulksendID

`func (o *EzsignbulksendsignermappingRequestV2) GetFkiEzsignbulksendID() int32`

GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetFkiEzsignbulksendIDOk

`func (o *EzsignbulksendsignermappingRequestV2) GetFkiEzsignbulksendIDOk() (*int32, bool)`

GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignbulksendID

`func (o *EzsignbulksendsignermappingRequestV2) SetFkiEzsignbulksendID(v int32)`

SetFkiEzsignbulksendID sets FkiEzsignbulksendID field to given value.


### GetFkiUserID

`func (o *EzsignbulksendsignermappingRequestV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsignbulksendsignermappingRequestV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsignbulksendsignermappingRequestV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsignbulksendsignermappingRequestV2) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetEEzsignbulksendsignermappingRole

`func (o *EzsignbulksendsignermappingRequestV2) GetEEzsignbulksendsignermappingRole() FieldEEzsignbulksendsignermappingRole`

GetEEzsignbulksendsignermappingRole returns the EEzsignbulksendsignermappingRole field if non-nil, zero value otherwise.

### GetEEzsignbulksendsignermappingRoleOk

`func (o *EzsignbulksendsignermappingRequestV2) GetEEzsignbulksendsignermappingRoleOk() (*FieldEEzsignbulksendsignermappingRole, bool)`

GetEEzsignbulksendsignermappingRoleOk returns a tuple with the EEzsignbulksendsignermappingRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignbulksendsignermappingRole

`func (o *EzsignbulksendsignermappingRequestV2) SetEEzsignbulksendsignermappingRole(v FieldEEzsignbulksendsignermappingRole)`

SetEEzsignbulksendsignermappingRole sets EEzsignbulksendsignermappingRole field to given value.

### HasEEzsignbulksendsignermappingRole

`func (o *EzsignbulksendsignermappingRequestV2) HasEEzsignbulksendsignermappingRole() bool`

HasEEzsignbulksendsignermappingRole returns a boolean if a field has been set.

### GetSEzsignbulksendsignermappingDescription

`func (o *EzsignbulksendsignermappingRequestV2) GetSEzsignbulksendsignermappingDescription() string`

GetSEzsignbulksendsignermappingDescription returns the SEzsignbulksendsignermappingDescription field if non-nil, zero value otherwise.

### GetSEzsignbulksendsignermappingDescriptionOk

`func (o *EzsignbulksendsignermappingRequestV2) GetSEzsignbulksendsignermappingDescriptionOk() (*string, bool)`

GetSEzsignbulksendsignermappingDescriptionOk returns a tuple with the SEzsignbulksendsignermappingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignbulksendsignermappingDescription

`func (o *EzsignbulksendsignermappingRequestV2) SetSEzsignbulksendsignermappingDescription(v string)`

SetSEzsignbulksendsignermappingDescription sets SEzsignbulksendsignermappingDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


