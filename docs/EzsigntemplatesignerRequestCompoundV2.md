# EzsigntemplatesignerRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignerID** | Pointer to **int32** | The unique ID of the Ezsigntemplatesigner | [optional] 
**FkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**EEzsigntemplatesignerRole** | Pointer to [**FieldEEzsigntemplatesignerRole**](FieldEEzsigntemplatesignerRole.md) |  | [optional] 
**EEzsigntemplatesignerMapping** | Pointer to [**FieldEEzsigntemplatesignerMapping**](FieldEEzsigntemplatesignerMapping.md) |  | [optional] 
**SEzsigntemplatesignerDescription** | **string** | The description of the Ezsigntemplatesigner | 

## Methods

### NewEzsigntemplatesignerRequestCompoundV2

`func NewEzsigntemplatesignerRequestCompoundV2(fkiEzsigntemplateID int32, sEzsigntemplatesignerDescription string, ) *EzsigntemplatesignerRequestCompoundV2`

NewEzsigntemplatesignerRequestCompoundV2 instantiates a new EzsigntemplatesignerRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignerRequestCompoundV2WithDefaults

`func NewEzsigntemplatesignerRequestCompoundV2WithDefaults() *EzsigntemplatesignerRequestCompoundV2`

NewEzsigntemplatesignerRequestCompoundV2WithDefaults instantiates a new EzsigntemplatesignerRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerRequestCompoundV2) GetPkiEzsigntemplatesignerID() int32`

GetPkiEzsigntemplatesignerID returns the PkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatesignerRequestCompoundV2) GetPkiEzsigntemplatesignerIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignerIDOk returns a tuple with the PkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerRequestCompoundV2) SetPkiEzsigntemplatesignerID(v int32)`

SetPkiEzsigntemplatesignerID sets PkiEzsigntemplatesignerID field to given value.

### HasPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerRequestCompoundV2) HasPkiEzsigntemplatesignerID() bool`

HasPkiEzsigntemplatesignerID returns a boolean if a field has been set.

### GetFkiEzsigntemplateID

`func (o *EzsigntemplatesignerRequestCompoundV2) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatesignerRequestCompoundV2) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatesignerRequestCompoundV2) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.


### GetFkiUserID

`func (o *EzsigntemplatesignerRequestCompoundV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigntemplatesignerRequestCompoundV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigntemplatesignerRequestCompoundV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigntemplatesignerRequestCompoundV2) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsigntemplatesignerRequestCompoundV2) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsigntemplatesignerRequestCompoundV2) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsigntemplatesignerRequestCompoundV2) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsigntemplatesignerRequestCompoundV2) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerRequestCompoundV2) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatesignerRequestCompoundV2) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerRequestCompoundV2) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerRequestCompoundV2) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetEEzsigntemplatesignerRole

`func (o *EzsigntemplatesignerRequestCompoundV2) GetEEzsigntemplatesignerRole() FieldEEzsigntemplatesignerRole`

GetEEzsigntemplatesignerRole returns the EEzsigntemplatesignerRole field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignerRoleOk

`func (o *EzsigntemplatesignerRequestCompoundV2) GetEEzsigntemplatesignerRoleOk() (*FieldEEzsigntemplatesignerRole, bool)`

GetEEzsigntemplatesignerRoleOk returns a tuple with the EEzsigntemplatesignerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignerRole

`func (o *EzsigntemplatesignerRequestCompoundV2) SetEEzsigntemplatesignerRole(v FieldEEzsigntemplatesignerRole)`

SetEEzsigntemplatesignerRole sets EEzsigntemplatesignerRole field to given value.

### HasEEzsigntemplatesignerRole

`func (o *EzsigntemplatesignerRequestCompoundV2) HasEEzsigntemplatesignerRole() bool`

HasEEzsigntemplatesignerRole returns a boolean if a field has been set.

### GetEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerRequestCompoundV2) GetEEzsigntemplatesignerMapping() FieldEEzsigntemplatesignerMapping`

GetEEzsigntemplatesignerMapping returns the EEzsigntemplatesignerMapping field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignerMappingOk

`func (o *EzsigntemplatesignerRequestCompoundV2) GetEEzsigntemplatesignerMappingOk() (*FieldEEzsigntemplatesignerMapping, bool)`

GetEEzsigntemplatesignerMappingOk returns a tuple with the EEzsigntemplatesignerMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerRequestCompoundV2) SetEEzsigntemplatesignerMapping(v FieldEEzsigntemplatesignerMapping)`

SetEEzsigntemplatesignerMapping sets EEzsigntemplatesignerMapping field to given value.

### HasEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerRequestCompoundV2) HasEEzsigntemplatesignerMapping() bool`

HasEEzsigntemplatesignerMapping returns a boolean if a field has been set.

### GetSEzsigntemplatesignerDescription

`func (o *EzsigntemplatesignerRequestCompoundV2) GetSEzsigntemplatesignerDescription() string`

GetSEzsigntemplatesignerDescription returns the SEzsigntemplatesignerDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignerDescriptionOk

`func (o *EzsigntemplatesignerRequestCompoundV2) GetSEzsigntemplatesignerDescriptionOk() (*string, bool)`

GetSEzsigntemplatesignerDescriptionOk returns a tuple with the SEzsigntemplatesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignerDescription

`func (o *EzsigntemplatesignerRequestCompoundV2) SetSEzsigntemplatesignerDescription(v string)`

SetSEzsigntemplatesignerDescription sets SEzsigntemplatesignerDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


