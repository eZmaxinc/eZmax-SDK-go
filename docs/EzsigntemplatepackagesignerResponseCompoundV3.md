# EzsigntemplatepackagesignerResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackagesignerID** | **int32** | The unique ID of the Ezsigntemplatepackagesigner | 
**FkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**SEzdoctemplatedocumentNameX** | Pointer to **string** | The name of the Ezdoctemplatedocument in the language of the requester | [optional] 
**EEzsigntemplatepackagesignerRole** | Pointer to [**FieldEEzsigntemplatepackagesignerRole**](FieldEEzsigntemplatepackagesignerRole.md) |  | [optional] 
**EEzsigntemplatepackagesignerMapping** | Pointer to [**FieldEEzsigntemplatepackagesignerMapping**](FieldEEzsigntemplatepackagesignerMapping.md) |  | [optional] [default to MANUAL]
**SEzsigntemplatepackagesignerDescription** | **string** | The description of the Ezsigntemplatepackagesigner | 
**SUserName** | Pointer to **string** | The description of the User in the language of the requester | [optional] 
**SUsergroupNameX** | Pointer to **string** | The Name of the Usergroup in the language of the requester | [optional] 

## Methods

### NewEzsigntemplatepackagesignerResponseCompoundV3

`func NewEzsigntemplatepackagesignerResponseCompoundV3(pkiEzsigntemplatepackagesignerID int32, fkiEzsigntemplatepackageID int32, sEzsigntemplatepackagesignerDescription string, ) *EzsigntemplatepackagesignerResponseCompoundV3`

NewEzsigntemplatepackagesignerResponseCompoundV3 instantiates a new EzsigntemplatepackagesignerResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackagesignerResponseCompoundV3WithDefaults

`func NewEzsigntemplatepackagesignerResponseCompoundV3WithDefaults() *EzsigntemplatepackagesignerResponseCompoundV3`

NewEzsigntemplatepackagesignerResponseCompoundV3WithDefaults instantiates a new EzsigntemplatepackagesignerResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetPkiEzsigntemplatepackagesignerID() int32`

GetPkiEzsigntemplatepackagesignerID returns the PkiEzsigntemplatepackagesignerID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackagesignerIDOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetPkiEzsigntemplatepackagesignerIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackagesignerIDOk returns a tuple with the PkiEzsigntemplatepackagesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetPkiEzsigntemplatepackagesignerID(v int32)`

SetPkiEzsigntemplatepackagesignerID sets PkiEzsigntemplatepackagesignerID field to given value.


### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetSEzdoctemplatedocumentNameX() string`

GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentNameXOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetSEzdoctemplatedocumentNameXOk() (*string, bool)`

GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetSEzdoctemplatedocumentNameX(v string)`

SetSEzdoctemplatedocumentNameX sets SEzdoctemplatedocumentNameX field to given value.

### HasSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) HasSEzdoctemplatedocumentNameX() bool`

HasSEzdoctemplatedocumentNameX returns a boolean if a field has been set.

### GetEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetEEzsigntemplatepackagesignerRole() FieldEEzsigntemplatepackagesignerRole`

GetEEzsigntemplatepackagesignerRole returns the EEzsigntemplatepackagesignerRole field if non-nil, zero value otherwise.

### GetEEzsigntemplatepackagesignerRoleOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetEEzsigntemplatepackagesignerRoleOk() (*FieldEEzsigntemplatepackagesignerRole, bool)`

GetEEzsigntemplatepackagesignerRoleOk returns a tuple with the EEzsigntemplatepackagesignerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetEEzsigntemplatepackagesignerRole(v FieldEEzsigntemplatepackagesignerRole)`

SetEEzsigntemplatepackagesignerRole sets EEzsigntemplatepackagesignerRole field to given value.

### HasEEzsigntemplatepackagesignerRole

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) HasEEzsigntemplatepackagesignerRole() bool`

HasEEzsigntemplatepackagesignerRole returns a boolean if a field has been set.

### GetEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetEEzsigntemplatepackagesignerMapping() FieldEEzsigntemplatepackagesignerMapping`

GetEEzsigntemplatepackagesignerMapping returns the EEzsigntemplatepackagesignerMapping field if non-nil, zero value otherwise.

### GetEEzsigntemplatepackagesignerMappingOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetEEzsigntemplatepackagesignerMappingOk() (*FieldEEzsigntemplatepackagesignerMapping, bool)`

GetEEzsigntemplatepackagesignerMappingOk returns a tuple with the EEzsigntemplatepackagesignerMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetEEzsigntemplatepackagesignerMapping(v FieldEEzsigntemplatepackagesignerMapping)`

SetEEzsigntemplatepackagesignerMapping sets EEzsigntemplatepackagesignerMapping field to given value.

### HasEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) HasEEzsigntemplatepackagesignerMapping() bool`

HasEEzsigntemplatepackagesignerMapping returns a boolean if a field has been set.

### GetSEzsigntemplatepackagesignerDescription

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetSEzsigntemplatepackagesignerDescription() string`

GetSEzsigntemplatepackagesignerDescription returns the SEzsigntemplatepackagesignerDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackagesignerDescriptionOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetSEzsigntemplatepackagesignerDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackagesignerDescriptionOk returns a tuple with the SEzsigntemplatepackagesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackagesignerDescription

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetSEzsigntemplatepackagesignerDescription(v string)`

SetSEzsigntemplatepackagesignerDescription sets SEzsigntemplatepackagesignerDescription field to given value.


### GetSUserName

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetSUserName() string`

GetSUserName returns the SUserName field if non-nil, zero value otherwise.

### GetSUserNameOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetSUserNameOk() (*string, bool)`

GetSUserNameOk returns a tuple with the SUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserName

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetSUserName(v string)`

SetSUserName sets SUserName field to given value.

### HasSUserName

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) HasSUserName() bool`

HasSUserName returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.

### HasSUsergroupNameX

`func (o *EzsigntemplatepackagesignerResponseCompoundV3) HasSUsergroupNameX() bool`

HasSUsergroupNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


