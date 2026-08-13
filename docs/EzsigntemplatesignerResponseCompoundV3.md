# EzsigntemplatesignerResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignerID** | **int32** | The unique ID of the Ezsigntemplatesigner | 
**FkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**EEzsigntemplatesignerRole** | Pointer to [**FieldEEzsigntemplatesignerRole**](FieldEEzsigntemplatesignerRole.md) |  | [optional] 
**EEzsigntemplatesignerMapping** | Pointer to [**FieldEEzsigntemplatesignerMapping**](FieldEEzsigntemplatesignerMapping.md) |  | [optional] 
**SEzsigntemplatesignerDescription** | **string** | The description of the Ezsigntemplatesigner | 
**SUserName** | Pointer to **string** | The description of the User in the language of the requester | [optional] 
**SUsergroupNameX** | Pointer to **string** | The Name of the Usergroup in the language of the requester | [optional] 

## Methods

### NewEzsigntemplatesignerResponseCompoundV3

`func NewEzsigntemplatesignerResponseCompoundV3(pkiEzsigntemplatesignerID int32, fkiEzsigntemplateID int32, sEzsigntemplatesignerDescription string, ) *EzsigntemplatesignerResponseCompoundV3`

NewEzsigntemplatesignerResponseCompoundV3 instantiates a new EzsigntemplatesignerResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignerResponseCompoundV3WithDefaults

`func NewEzsigntemplatesignerResponseCompoundV3WithDefaults() *EzsigntemplatesignerResponseCompoundV3`

NewEzsigntemplatesignerResponseCompoundV3WithDefaults instantiates a new EzsigntemplatesignerResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerResponseCompoundV3) GetPkiEzsigntemplatesignerID() int32`

GetPkiEzsigntemplatesignerID returns the PkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetPkiEzsigntemplatesignerIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignerIDOk returns a tuple with the PkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerResponseCompoundV3) SetPkiEzsigntemplatesignerID(v int32)`

SetPkiEzsigntemplatesignerID sets PkiEzsigntemplatesignerID field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatesignerResponseCompoundV3) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatesignerResponseCompoundV3) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.


### GetFkiUserID

`func (o *EzsigntemplatesignerResponseCompoundV3) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigntemplatesignerResponseCompoundV3) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigntemplatesignerResponseCompoundV3) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsigntemplatesignerResponseCompoundV3) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsigntemplatesignerResponseCompoundV3) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsigntemplatesignerResponseCompoundV3) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerResponseCompoundV3) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerResponseCompoundV3) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerResponseCompoundV3) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetEEzsigntemplatesignerRole

`func (o *EzsigntemplatesignerResponseCompoundV3) GetEEzsigntemplatesignerRole() FieldEEzsigntemplatesignerRole`

GetEEzsigntemplatesignerRole returns the EEzsigntemplatesignerRole field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignerRoleOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetEEzsigntemplatesignerRoleOk() (*FieldEEzsigntemplatesignerRole, bool)`

GetEEzsigntemplatesignerRoleOk returns a tuple with the EEzsigntemplatesignerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignerRole

`func (o *EzsigntemplatesignerResponseCompoundV3) SetEEzsigntemplatesignerRole(v FieldEEzsigntemplatesignerRole)`

SetEEzsigntemplatesignerRole sets EEzsigntemplatesignerRole field to given value.

### HasEEzsigntemplatesignerRole

`func (o *EzsigntemplatesignerResponseCompoundV3) HasEEzsigntemplatesignerRole() bool`

HasEEzsigntemplatesignerRole returns a boolean if a field has been set.

### GetEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerResponseCompoundV3) GetEEzsigntemplatesignerMapping() FieldEEzsigntemplatesignerMapping`

GetEEzsigntemplatesignerMapping returns the EEzsigntemplatesignerMapping field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignerMappingOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetEEzsigntemplatesignerMappingOk() (*FieldEEzsigntemplatesignerMapping, bool)`

GetEEzsigntemplatesignerMappingOk returns a tuple with the EEzsigntemplatesignerMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerResponseCompoundV3) SetEEzsigntemplatesignerMapping(v FieldEEzsigntemplatesignerMapping)`

SetEEzsigntemplatesignerMapping sets EEzsigntemplatesignerMapping field to given value.

### HasEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerResponseCompoundV3) HasEEzsigntemplatesignerMapping() bool`

HasEEzsigntemplatesignerMapping returns a boolean if a field has been set.

### GetSEzsigntemplatesignerDescription

`func (o *EzsigntemplatesignerResponseCompoundV3) GetSEzsigntemplatesignerDescription() string`

GetSEzsigntemplatesignerDescription returns the SEzsigntemplatesignerDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignerDescriptionOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetSEzsigntemplatesignerDescriptionOk() (*string, bool)`

GetSEzsigntemplatesignerDescriptionOk returns a tuple with the SEzsigntemplatesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignerDescription

`func (o *EzsigntemplatesignerResponseCompoundV3) SetSEzsigntemplatesignerDescription(v string)`

SetSEzsigntemplatesignerDescription sets SEzsigntemplatesignerDescription field to given value.


### GetSUserName

`func (o *EzsigntemplatesignerResponseCompoundV3) GetSUserName() string`

GetSUserName returns the SUserName field if non-nil, zero value otherwise.

### GetSUserNameOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetSUserNameOk() (*string, bool)`

GetSUserNameOk returns a tuple with the SUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserName

`func (o *EzsigntemplatesignerResponseCompoundV3) SetSUserName(v string)`

SetSUserName sets SUserName field to given value.

### HasSUserName

`func (o *EzsigntemplatesignerResponseCompoundV3) HasSUserName() bool`

HasSUserName returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *EzsigntemplatesignerResponseCompoundV3) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *EzsigntemplatesignerResponseCompoundV3) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *EzsigntemplatesignerResponseCompoundV3) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.

### HasSUsergroupNameX

`func (o *EzsigntemplatesignerResponseCompoundV3) HasSUsergroupNameX() bool`

HasSUsergroupNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


