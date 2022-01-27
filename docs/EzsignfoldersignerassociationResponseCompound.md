# EzsignfoldersignerassociationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjUser** | Pointer to [**EzsignfoldersignerassociationResponseCompoundUser**](EzsignfoldersignerassociationResponseCompoundUser.md) |  | [optional] 
**ObjEzsignsigner** | Pointer to [**EzsignsignerResponseCompound**](EzsignsignerResponseCompound.md) |  | [optional] 
**PkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**BEzsignfoldersignerassociationReceivecopy** | **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | 

## Methods

### NewEzsignfoldersignerassociationResponseCompound

`func NewEzsignfoldersignerassociationResponseCompound(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, bEzsignfoldersignerassociationReceivecopy bool, ) *EzsignfoldersignerassociationResponseCompound`

NewEzsignfoldersignerassociationResponseCompound instantiates a new EzsignfoldersignerassociationResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationResponseCompoundWithDefaults

`func NewEzsignfoldersignerassociationResponseCompoundWithDefaults() *EzsignfoldersignerassociationResponseCompound`

NewEzsignfoldersignerassociationResponseCompoundWithDefaults instantiates a new EzsignfoldersignerassociationResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjUser

`func (o *EzsignfoldersignerassociationResponseCompound) GetObjUser() EzsignfoldersignerassociationResponseCompoundUser`

GetObjUser returns the ObjUser field if non-nil, zero value otherwise.

### GetObjUserOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetObjUserOk() (*EzsignfoldersignerassociationResponseCompoundUser, bool)`

GetObjUserOk returns a tuple with the ObjUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUser

`func (o *EzsignfoldersignerassociationResponseCompound) SetObjUser(v EzsignfoldersignerassociationResponseCompoundUser)`

SetObjUser sets ObjUser field to given value.

### HasObjUser

`func (o *EzsignfoldersignerassociationResponseCompound) HasObjUser() bool`

HasObjUser returns a boolean if a field has been set.

### GetObjEzsignsigner

`func (o *EzsignfoldersignerassociationResponseCompound) GetObjEzsignsigner() EzsignsignerResponseCompound`

GetObjEzsignsigner returns the ObjEzsignsigner field if non-nil, zero value otherwise.

### GetObjEzsignsignerOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetObjEzsignsignerOk() (*EzsignsignerResponseCompound, bool)`

GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsigner

`func (o *EzsignfoldersignerassociationResponseCompound) SetObjEzsignsigner(v EzsignsignerResponseCompound)`

SetObjEzsignsigner sets ObjEzsignsigner field to given value.

### HasObjEzsignsigner

`func (o *EzsignfoldersignerassociationResponseCompound) HasObjEzsignsigner() bool`

HasObjEzsignsigner returns a boolean if a field has been set.

### GetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationResponseCompound) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationResponseCompound) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationResponseCompound) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationResponseCompound) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationResponseCompound) GetBEzsignfoldersignerassociationReceivecopy() bool`

GetBEzsignfoldersignerassociationReceivecopy returns the BEzsignfoldersignerassociationReceivecopy field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationReceivecopyOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetBEzsignfoldersignerassociationReceivecopyOk() (*bool, bool)`

GetBEzsignfoldersignerassociationReceivecopyOk returns a tuple with the BEzsignfoldersignerassociationReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationResponseCompound) SetBEzsignfoldersignerassociationReceivecopy(v bool)`

SetBEzsignfoldersignerassociationReceivecopy sets BEzsignfoldersignerassociationReceivecopy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


