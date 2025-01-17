# EzsigntemplatesignerResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignerID** | **int32** | The unique ID of the Ezsigntemplatesigner | 
**FkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**BEzsigntemplatesignerReceivecopy** | Pointer to **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | [optional] 
**EEzsigntemplatesignerMapping** | Pointer to [**FieldEEzsigntemplatesignerMapping**](FieldEEzsigntemplatesignerMapping.md) |  | [optional] 
**SEzsigntemplatesignerDescription** | **string** | The description of the Ezsigntemplatesigner | 
**SUserName** | Pointer to **string** | The description of the User in the language of the requester | [optional] 
**SUsergroupNameX** | Pointer to **string** | The Name of the Usergroup in the language of the requester | [optional] 

## Methods

### NewEzsigntemplatesignerResponseCompound

`func NewEzsigntemplatesignerResponseCompound(pkiEzsigntemplatesignerID int32, fkiEzsigntemplateID int32, sEzsigntemplatesignerDescription string, ) *EzsigntemplatesignerResponseCompound`

NewEzsigntemplatesignerResponseCompound instantiates a new EzsigntemplatesignerResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignerResponseCompoundWithDefaults

`func NewEzsigntemplatesignerResponseCompoundWithDefaults() *EzsigntemplatesignerResponseCompound`

NewEzsigntemplatesignerResponseCompoundWithDefaults instantiates a new EzsigntemplatesignerResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerResponseCompound) GetPkiEzsigntemplatesignerID() int32`

GetPkiEzsigntemplatesignerID returns the PkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatesignerResponseCompound) GetPkiEzsigntemplatesignerIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignerIDOk returns a tuple with the PkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerResponseCompound) SetPkiEzsigntemplatesignerID(v int32)`

SetPkiEzsigntemplatesignerID sets PkiEzsigntemplatesignerID field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatesignerResponseCompound) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatesignerResponseCompound) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatesignerResponseCompound) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.


### GetFkiUserID

`func (o *EzsigntemplatesignerResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigntemplatesignerResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigntemplatesignerResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigntemplatesignerResponseCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsigntemplatesignerResponseCompound) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsigntemplatesignerResponseCompound) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsigntemplatesignerResponseCompound) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsigntemplatesignerResponseCompound) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerResponseCompound) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatesignerResponseCompound) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerResponseCompound) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerResponseCompound) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetBEzsigntemplatesignerReceivecopy

`func (o *EzsigntemplatesignerResponseCompound) GetBEzsigntemplatesignerReceivecopy() bool`

GetBEzsigntemplatesignerReceivecopy returns the BEzsigntemplatesignerReceivecopy field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignerReceivecopyOk

`func (o *EzsigntemplatesignerResponseCompound) GetBEzsigntemplatesignerReceivecopyOk() (*bool, bool)`

GetBEzsigntemplatesignerReceivecopyOk returns a tuple with the BEzsigntemplatesignerReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignerReceivecopy

`func (o *EzsigntemplatesignerResponseCompound) SetBEzsigntemplatesignerReceivecopy(v bool)`

SetBEzsigntemplatesignerReceivecopy sets BEzsigntemplatesignerReceivecopy field to given value.

### HasBEzsigntemplatesignerReceivecopy

`func (o *EzsigntemplatesignerResponseCompound) HasBEzsigntemplatesignerReceivecopy() bool`

HasBEzsigntemplatesignerReceivecopy returns a boolean if a field has been set.

### GetEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerResponseCompound) GetEEzsigntemplatesignerMapping() FieldEEzsigntemplatesignerMapping`

GetEEzsigntemplatesignerMapping returns the EEzsigntemplatesignerMapping field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignerMappingOk

`func (o *EzsigntemplatesignerResponseCompound) GetEEzsigntemplatesignerMappingOk() (*FieldEEzsigntemplatesignerMapping, bool)`

GetEEzsigntemplatesignerMappingOk returns a tuple with the EEzsigntemplatesignerMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerResponseCompound) SetEEzsigntemplatesignerMapping(v FieldEEzsigntemplatesignerMapping)`

SetEEzsigntemplatesignerMapping sets EEzsigntemplatesignerMapping field to given value.

### HasEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerResponseCompound) HasEEzsigntemplatesignerMapping() bool`

HasEEzsigntemplatesignerMapping returns a boolean if a field has been set.

### GetSEzsigntemplatesignerDescription

`func (o *EzsigntemplatesignerResponseCompound) GetSEzsigntemplatesignerDescription() string`

GetSEzsigntemplatesignerDescription returns the SEzsigntemplatesignerDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignerDescriptionOk

`func (o *EzsigntemplatesignerResponseCompound) GetSEzsigntemplatesignerDescriptionOk() (*string, bool)`

GetSEzsigntemplatesignerDescriptionOk returns a tuple with the SEzsigntemplatesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignerDescription

`func (o *EzsigntemplatesignerResponseCompound) SetSEzsigntemplatesignerDescription(v string)`

SetSEzsigntemplatesignerDescription sets SEzsigntemplatesignerDescription field to given value.


### GetSUserName

`func (o *EzsigntemplatesignerResponseCompound) GetSUserName() string`

GetSUserName returns the SUserName field if non-nil, zero value otherwise.

### GetSUserNameOk

`func (o *EzsigntemplatesignerResponseCompound) GetSUserNameOk() (*string, bool)`

GetSUserNameOk returns a tuple with the SUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserName

`func (o *EzsigntemplatesignerResponseCompound) SetSUserName(v string)`

SetSUserName sets SUserName field to given value.

### HasSUserName

`func (o *EzsigntemplatesignerResponseCompound) HasSUserName() bool`

HasSUserName returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *EzsigntemplatesignerResponseCompound) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *EzsigntemplatesignerResponseCompound) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *EzsigntemplatesignerResponseCompound) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.

### HasSUsergroupNameX

`func (o *EzsigntemplatesignerResponseCompound) HasSUsergroupNameX() bool`

HasSUsergroupNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


