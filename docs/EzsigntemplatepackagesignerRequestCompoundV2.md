# EzsigntemplatepackagesignerRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackagesignerID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepackagesigner | [optional] 
**FkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**EEzsigntemplatepackagesignerRole** | Pointer to [**FieldEEzsigntemplatepackagesignerRole**](FieldEEzsigntemplatepackagesignerRole.md) |  | [optional] 
**EEzsigntemplatepackagesignerMapping** | Pointer to [**FieldEEzsigntemplatepackagesignerMapping**](FieldEEzsigntemplatepackagesignerMapping.md) |  | [optional] [default to MANUAL]
**SEzsigntemplatepackagesignerDescription** | **string** | The description of the Ezsigntemplatepackagesigner | 

## Methods

### NewEzsigntemplatepackagesignerRequestCompoundV2

`func NewEzsigntemplatepackagesignerRequestCompoundV2(fkiEzsigntemplatepackageID int32, sEzsigntemplatepackagesignerDescription string, ) *EzsigntemplatepackagesignerRequestCompoundV2`

NewEzsigntemplatepackagesignerRequestCompoundV2 instantiates a new EzsigntemplatepackagesignerRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackagesignerRequestCompoundV2WithDefaults

`func NewEzsigntemplatepackagesignerRequestCompoundV2WithDefaults() *EzsigntemplatepackagesignerRequestCompoundV2`

NewEzsigntemplatepackagesignerRequestCompoundV2WithDefaults instantiates a new EzsigntemplatepackagesignerRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetPkiEzsigntemplatepackagesignerID() int32`

GetPkiEzsigntemplatepackagesignerID returns the PkiEzsigntemplatepackagesignerID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackagesignerIDOk

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetPkiEzsigntemplatepackagesignerIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackagesignerIDOk returns a tuple with the PkiEzsigntemplatepackagesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) SetPkiEzsigntemplatepackagesignerID(v int32)`

SetPkiEzsigntemplatepackagesignerID sets PkiEzsigntemplatepackagesignerID field to given value.

### HasPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) HasPkiEzsigntemplatepackagesignerID() bool`

HasPkiEzsigntemplatepackagesignerID returns a boolean if a field has been set.

### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetEEzsigntemplatepackagesignerRole() FieldEEzsigntemplatepackagesignerRole`

GetEEzsigntemplatepackagesignerRole returns the EEzsigntemplatepackagesignerRole field if non-nil, zero value otherwise.

### GetEEzsigntemplatepackagesignerRoleOk

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetEEzsigntemplatepackagesignerRoleOk() (*FieldEEzsigntemplatepackagesignerRole, bool)`

GetEEzsigntemplatepackagesignerRoleOk returns a tuple with the EEzsigntemplatepackagesignerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) SetEEzsigntemplatepackagesignerRole(v FieldEEzsigntemplatepackagesignerRole)`

SetEEzsigntemplatepackagesignerRole sets EEzsigntemplatepackagesignerRole field to given value.

### HasEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) HasEEzsigntemplatepackagesignerRole() bool`

HasEEzsigntemplatepackagesignerRole returns a boolean if a field has been set.

### GetEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetEEzsigntemplatepackagesignerMapping() FieldEEzsigntemplatepackagesignerMapping`

GetEEzsigntemplatepackagesignerMapping returns the EEzsigntemplatepackagesignerMapping field if non-nil, zero value otherwise.

### GetEEzsigntemplatepackagesignerMappingOk

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetEEzsigntemplatepackagesignerMappingOk() (*FieldEEzsigntemplatepackagesignerMapping, bool)`

GetEEzsigntemplatepackagesignerMappingOk returns a tuple with the EEzsigntemplatepackagesignerMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) SetEEzsigntemplatepackagesignerMapping(v FieldEEzsigntemplatepackagesignerMapping)`

SetEEzsigntemplatepackagesignerMapping sets EEzsigntemplatepackagesignerMapping field to given value.

### HasEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) HasEEzsigntemplatepackagesignerMapping() bool`

HasEEzsigntemplatepackagesignerMapping returns a boolean if a field has been set.

### GetSEzsigntemplatepackagesignerDescription

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetSEzsigntemplatepackagesignerDescription() string`

GetSEzsigntemplatepackagesignerDescription returns the SEzsigntemplatepackagesignerDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackagesignerDescriptionOk

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) GetSEzsigntemplatepackagesignerDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackagesignerDescriptionOk returns a tuple with the SEzsigntemplatepackagesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackagesignerDescription

`func (o *EzsigntemplatepackagesignerRequestCompoundV2) SetSEzsigntemplatepackagesignerDescription(v string)`

SetSEzsigntemplatepackagesignerDescription sets SEzsigntemplatepackagesignerDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


