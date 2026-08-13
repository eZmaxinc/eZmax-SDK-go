# EzsigntemplatepackagesignerRequestV2

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

### NewEzsigntemplatepackagesignerRequestV2

`func NewEzsigntemplatepackagesignerRequestV2(fkiEzsigntemplatepackageID int32, sEzsigntemplatepackagesignerDescription string, ) *EzsigntemplatepackagesignerRequestV2`

NewEzsigntemplatepackagesignerRequestV2 instantiates a new EzsigntemplatepackagesignerRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackagesignerRequestV2WithDefaults

`func NewEzsigntemplatepackagesignerRequestV2WithDefaults() *EzsigntemplatepackagesignerRequestV2`

NewEzsigntemplatepackagesignerRequestV2WithDefaults instantiates a new EzsigntemplatepackagesignerRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerRequestV2) GetPkiEzsigntemplatepackagesignerID() int32`

GetPkiEzsigntemplatepackagesignerID returns the PkiEzsigntemplatepackagesignerID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackagesignerIDOk

`func (o *EzsigntemplatepackagesignerRequestV2) GetPkiEzsigntemplatepackagesignerIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackagesignerIDOk returns a tuple with the PkiEzsigntemplatepackagesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerRequestV2) SetPkiEzsigntemplatepackagesignerID(v int32)`

SetPkiEzsigntemplatepackagesignerID sets PkiEzsigntemplatepackagesignerID field to given value.

### HasPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerRequestV2) HasPkiEzsigntemplatepackagesignerID() bool`

HasPkiEzsigntemplatepackagesignerID returns a boolean if a field has been set.

### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagesignerRequestV2) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackagesignerRequestV2) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagesignerRequestV2) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerRequestV2) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatepackagesignerRequestV2) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerRequestV2) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerRequestV2) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *EzsigntemplatepackagesignerRequestV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigntemplatepackagesignerRequestV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigntemplatepackagesignerRequestV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigntemplatepackagesignerRequestV2) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsigntemplatepackagesignerRequestV2) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsigntemplatepackagesignerRequestV2) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsigntemplatepackagesignerRequestV2) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsigntemplatepackagesignerRequestV2) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerRequestV2) GetEEzsigntemplatepackagesignerRole() FieldEEzsigntemplatepackagesignerRole`

GetEEzsigntemplatepackagesignerRole returns the EEzsigntemplatepackagesignerRole field if non-nil, zero value otherwise.

### GetEEzsigntemplatepackagesignerRoleOk

`func (o *EzsigntemplatepackagesignerRequestV2) GetEEzsigntemplatepackagesignerRoleOk() (*FieldEEzsigntemplatepackagesignerRole, bool)`

GetEEzsigntemplatepackagesignerRoleOk returns a tuple with the EEzsigntemplatepackagesignerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerRequestV2) SetEEzsigntemplatepackagesignerRole(v FieldEEzsigntemplatepackagesignerRole)`

SetEEzsigntemplatepackagesignerRole sets EEzsigntemplatepackagesignerRole field to given value.

### HasEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerRequestV2) HasEEzsigntemplatepackagesignerRole() bool`

HasEEzsigntemplatepackagesignerRole returns a boolean if a field has been set.

### GetEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerRequestV2) GetEEzsigntemplatepackagesignerMapping() FieldEEzsigntemplatepackagesignerMapping`

GetEEzsigntemplatepackagesignerMapping returns the EEzsigntemplatepackagesignerMapping field if non-nil, zero value otherwise.

### GetEEzsigntemplatepackagesignerMappingOk

`func (o *EzsigntemplatepackagesignerRequestV2) GetEEzsigntemplatepackagesignerMappingOk() (*FieldEEzsigntemplatepackagesignerMapping, bool)`

GetEEzsigntemplatepackagesignerMappingOk returns a tuple with the EEzsigntemplatepackagesignerMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerRequestV2) SetEEzsigntemplatepackagesignerMapping(v FieldEEzsigntemplatepackagesignerMapping)`

SetEEzsigntemplatepackagesignerMapping sets EEzsigntemplatepackagesignerMapping field to given value.

### HasEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerRequestV2) HasEEzsigntemplatepackagesignerMapping() bool`

HasEEzsigntemplatepackagesignerMapping returns a boolean if a field has been set.

### GetSEzsigntemplatepackagesignerDescription

`func (o *EzsigntemplatepackagesignerRequestV2) GetSEzsigntemplatepackagesignerDescription() string`

GetSEzsigntemplatepackagesignerDescription returns the SEzsigntemplatepackagesignerDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackagesignerDescriptionOk

`func (o *EzsigntemplatepackagesignerRequestV2) GetSEzsigntemplatepackagesignerDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackagesignerDescriptionOk returns a tuple with the SEzsigntemplatepackagesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackagesignerDescription

`func (o *EzsigntemplatepackagesignerRequestV2) SetSEzsigntemplatepackagesignerDescription(v string)`

SetSEzsigntemplatepackagesignerDescription sets SEzsigntemplatepackagesignerDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


