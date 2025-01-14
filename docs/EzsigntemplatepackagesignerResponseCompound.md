# EzsigntemplatepackagesignerResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackagesignerID** | **int32** | The unique ID of the Ezsigntemplatepackagesigner | 
**FkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**FkiEzdoctemplatedocumentID** | Pointer to **int32** | The unique ID of the Ezdoctemplatedocument | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**SEzdoctemplatedocumentNameX** | Pointer to **string** | The name of the Ezdoctemplatedocument in the language of the requester | [optional] 
**BEzsigntemplatepackagesignerReceivecopy** | Pointer to **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | [optional] 
**EEzsigntemplatepackagesignerMapping** | Pointer to [**FieldEEzsigntemplatepackagesignerMapping**](FieldEEzsigntemplatepackagesignerMapping.md) |  | [optional] [default to MANUAL]
**SEzsigntemplatepackagesignerDescription** | **string** | The description of the Ezsigntemplatepackagesigner | 
**SUserName** | Pointer to **string** | The description of the User in the language of the requester | [optional] 
**SUsergroupNameX** | Pointer to **string** | The Name of the Usergroup in the language of the requester | [optional] 

## Methods

### NewEzsigntemplatepackagesignerResponseCompound

`func NewEzsigntemplatepackagesignerResponseCompound(pkiEzsigntemplatepackagesignerID int32, fkiEzsigntemplatepackageID int32, sEzsigntemplatepackagesignerDescription string, ) *EzsigntemplatepackagesignerResponseCompound`

NewEzsigntemplatepackagesignerResponseCompound instantiates a new EzsigntemplatepackagesignerResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackagesignerResponseCompoundWithDefaults

`func NewEzsigntemplatepackagesignerResponseCompoundWithDefaults() *EzsigntemplatepackagesignerResponseCompound`

NewEzsigntemplatepackagesignerResponseCompoundWithDefaults instantiates a new EzsigntemplatepackagesignerResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerResponseCompound) GetPkiEzsigntemplatepackagesignerID() int32`

GetPkiEzsigntemplatepackagesignerID returns the PkiEzsigntemplatepackagesignerID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackagesignerIDOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetPkiEzsigntemplatepackagesignerIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackagesignerIDOk returns a tuple with the PkiEzsigntemplatepackagesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackagesignerID

`func (o *EzsigntemplatepackagesignerResponseCompound) SetPkiEzsigntemplatepackagesignerID(v int32)`

SetPkiEzsigntemplatepackagesignerID sets PkiEzsigntemplatepackagesignerID field to given value.


### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagesignerResponseCompound) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagesignerResponseCompound) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.


### GetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerResponseCompound) GetFkiEzdoctemplatedocumentID() int32`

GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzdoctemplatedocumentIDOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool)`

GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerResponseCompound) SetFkiEzdoctemplatedocumentID(v int32)`

SetFkiEzdoctemplatedocumentID sets FkiEzdoctemplatedocumentID field to given value.

### HasFkiEzdoctemplatedocumentID

`func (o *EzsigntemplatepackagesignerResponseCompound) HasFkiEzdoctemplatedocumentID() bool`

HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *EzsigntemplatepackagesignerResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigntemplatepackagesignerResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigntemplatepackagesignerResponseCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *EzsigntemplatepackagesignerResponseCompound) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *EzsigntemplatepackagesignerResponseCompound) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *EzsigntemplatepackagesignerResponseCompound) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackagesignerResponseCompound) GetSEzdoctemplatedocumentNameX() string`

GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentNameXOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetSEzdoctemplatedocumentNameXOk() (*string, bool)`

GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackagesignerResponseCompound) SetSEzdoctemplatedocumentNameX(v string)`

SetSEzdoctemplatedocumentNameX sets SEzdoctemplatedocumentNameX field to given value.

### HasSEzdoctemplatedocumentNameX

`func (o *EzsigntemplatepackagesignerResponseCompound) HasSEzdoctemplatedocumentNameX() bool`

HasSEzdoctemplatedocumentNameX returns a boolean if a field has been set.

### GetBEzsigntemplatepackagesignerReceivecopy

`func (o *EzsigntemplatepackagesignerResponseCompound) GetBEzsigntemplatepackagesignerReceivecopy() bool`

GetBEzsigntemplatepackagesignerReceivecopy returns the BEzsigntemplatepackagesignerReceivecopy field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackagesignerReceivecopyOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetBEzsigntemplatepackagesignerReceivecopyOk() (*bool, bool)`

GetBEzsigntemplatepackagesignerReceivecopyOk returns a tuple with the BEzsigntemplatepackagesignerReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackagesignerReceivecopy

`func (o *EzsigntemplatepackagesignerResponseCompound) SetBEzsigntemplatepackagesignerReceivecopy(v bool)`

SetBEzsigntemplatepackagesignerReceivecopy sets BEzsigntemplatepackagesignerReceivecopy field to given value.

### HasBEzsigntemplatepackagesignerReceivecopy

`func (o *EzsigntemplatepackagesignerResponseCompound) HasBEzsigntemplatepackagesignerReceivecopy() bool`

HasBEzsigntemplatepackagesignerReceivecopy returns a boolean if a field has been set.

### GetEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerResponseCompound) GetEEzsigntemplatepackagesignerMapping() FieldEEzsigntemplatepackagesignerMapping`

GetEEzsigntemplatepackagesignerMapping returns the EEzsigntemplatepackagesignerMapping field if non-nil, zero value otherwise.

### GetEEzsigntemplatepackagesignerMappingOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetEEzsigntemplatepackagesignerMappingOk() (*FieldEEzsigntemplatepackagesignerMapping, bool)`

GetEEzsigntemplatepackagesignerMappingOk returns a tuple with the EEzsigntemplatepackagesignerMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerResponseCompound) SetEEzsigntemplatepackagesignerMapping(v FieldEEzsigntemplatepackagesignerMapping)`

SetEEzsigntemplatepackagesignerMapping sets EEzsigntemplatepackagesignerMapping field to given value.

### HasEEzsigntemplatepackagesignerMapping

`func (o *EzsigntemplatepackagesignerResponseCompound) HasEEzsigntemplatepackagesignerMapping() bool`

HasEEzsigntemplatepackagesignerMapping returns a boolean if a field has been set.

### GetSEzsigntemplatepackagesignerDescription

`func (o *EzsigntemplatepackagesignerResponseCompound) GetSEzsigntemplatepackagesignerDescription() string`

GetSEzsigntemplatepackagesignerDescription returns the SEzsigntemplatepackagesignerDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackagesignerDescriptionOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetSEzsigntemplatepackagesignerDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackagesignerDescriptionOk returns a tuple with the SEzsigntemplatepackagesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackagesignerDescription

`func (o *EzsigntemplatepackagesignerResponseCompound) SetSEzsigntemplatepackagesignerDescription(v string)`

SetSEzsigntemplatepackagesignerDescription sets SEzsigntemplatepackagesignerDescription field to given value.


### GetSUserName

`func (o *EzsigntemplatepackagesignerResponseCompound) GetSUserName() string`

GetSUserName returns the SUserName field if non-nil, zero value otherwise.

### GetSUserNameOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetSUserNameOk() (*string, bool)`

GetSUserNameOk returns a tuple with the SUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserName

`func (o *EzsigntemplatepackagesignerResponseCompound) SetSUserName(v string)`

SetSUserName sets SUserName field to given value.

### HasSUserName

`func (o *EzsigntemplatepackagesignerResponseCompound) HasSUserName() bool`

HasSUserName returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *EzsigntemplatepackagesignerResponseCompound) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *EzsigntemplatepackagesignerResponseCompound) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *EzsigntemplatepackagesignerResponseCompound) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.

### HasSUsergroupNameX

`func (o *EzsigntemplatepackagesignerResponseCompound) HasSUsergroupNameX() bool`

HasSUsergroupNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


