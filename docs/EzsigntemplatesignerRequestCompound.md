# EzsigntemplatesignerRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignerID** | Pointer to **int32** | The unique ID of the Ezsigntemplatesigner | [optional] 
**FkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**BEzsigntemplatesignerReceivecopy** | Pointer to **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | [optional] 
**EEzsigntemplatesignerMapping** | Pointer to [**FieldEEzsigntemplatesignerMapping**](FieldEEzsigntemplatesignerMapping.md) |  | [optional] 
**SEzsigntemplatesignerDescription** | **string** | The description of the Ezsigntemplatesigner | 

## Methods

### NewEzsigntemplatesignerRequestCompound

`func NewEzsigntemplatesignerRequestCompound(fkiEzsigntemplateID int32, sEzsigntemplatesignerDescription string, ) *EzsigntemplatesignerRequestCompound`

NewEzsigntemplatesignerRequestCompound instantiates a new EzsigntemplatesignerRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignerRequestCompoundWithDefaults

`func NewEzsigntemplatesignerRequestCompoundWithDefaults() *EzsigntemplatesignerRequestCompound`

NewEzsigntemplatesignerRequestCompoundWithDefaults instantiates a new EzsigntemplatesignerRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerRequestCompound) GetPkiEzsigntemplatesignerID() int32`

GetPkiEzsigntemplatesignerID returns the PkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatesignerRequestCompound) GetPkiEzsigntemplatesignerIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignerIDOk returns a tuple with the PkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerRequestCompound) SetPkiEzsigntemplatesignerID(v int32)`

SetPkiEzsigntemplatesignerID sets PkiEzsigntemplatesignerID field to given value.

### HasPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerRequestCompound) HasPkiEzsigntemplatesignerID() bool`

HasPkiEzsigntemplatesignerID returns a boolean if a field has been set.

### GetFkiEzsigntemplateID

`func (o *EzsigntemplatesignerRequestCompound) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatesignerRequestCompound) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatesignerRequestCompound) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.


### GetFkiUserID

`func (o *EzsigntemplatesignerRequestCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigntemplatesignerRequestCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigntemplatesignerRequestCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigntemplatesignerRequestCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsigntemplatesignerRequestCompound) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsigntemplatesignerRequestCompound) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsigntemplatesignerRequestCompound) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsigntemplatesignerRequestCompound) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerRequestCompound) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatesignerRequestCompound) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerRequestCompound) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatesignerRequestCompound) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetBEzsigntemplatesignerReceivecopy

`func (o *EzsigntemplatesignerRequestCompound) GetBEzsigntemplatesignerReceivecopy() bool`

GetBEzsigntemplatesignerReceivecopy returns the BEzsigntemplatesignerReceivecopy field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignerReceivecopyOk

`func (o *EzsigntemplatesignerRequestCompound) GetBEzsigntemplatesignerReceivecopyOk() (*bool, bool)`

GetBEzsigntemplatesignerReceivecopyOk returns a tuple with the BEzsigntemplatesignerReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignerReceivecopy

`func (o *EzsigntemplatesignerRequestCompound) SetBEzsigntemplatesignerReceivecopy(v bool)`

SetBEzsigntemplatesignerReceivecopy sets BEzsigntemplatesignerReceivecopy field to given value.

### HasBEzsigntemplatesignerReceivecopy

`func (o *EzsigntemplatesignerRequestCompound) HasBEzsigntemplatesignerReceivecopy() bool`

HasBEzsigntemplatesignerReceivecopy returns a boolean if a field has been set.

### GetEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerRequestCompound) GetEEzsigntemplatesignerMapping() FieldEEzsigntemplatesignerMapping`

GetEEzsigntemplatesignerMapping returns the EEzsigntemplatesignerMapping field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignerMappingOk

`func (o *EzsigntemplatesignerRequestCompound) GetEEzsigntemplatesignerMappingOk() (*FieldEEzsigntemplatesignerMapping, bool)`

GetEEzsigntemplatesignerMappingOk returns a tuple with the EEzsigntemplatesignerMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerRequestCompound) SetEEzsigntemplatesignerMapping(v FieldEEzsigntemplatesignerMapping)`

SetEEzsigntemplatesignerMapping sets EEzsigntemplatesignerMapping field to given value.

### HasEEzsigntemplatesignerMapping

`func (o *EzsigntemplatesignerRequestCompound) HasEEzsigntemplatesignerMapping() bool`

HasEEzsigntemplatesignerMapping returns a boolean if a field has been set.

### GetSEzsigntemplatesignerDescription

`func (o *EzsigntemplatesignerRequestCompound) GetSEzsigntemplatesignerDescription() string`

GetSEzsigntemplatesignerDescription returns the SEzsigntemplatesignerDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignerDescriptionOk

`func (o *EzsigntemplatesignerRequestCompound) GetSEzsigntemplatesignerDescriptionOk() (*string, bool)`

GetSEzsigntemplatesignerDescriptionOk returns a tuple with the SEzsigntemplatesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignerDescription

`func (o *EzsigntemplatesignerRequestCompound) SetSEzsigntemplatesignerDescription(v string)`

SetSEzsigntemplatesignerDescription sets SEzsigntemplatesignerDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


