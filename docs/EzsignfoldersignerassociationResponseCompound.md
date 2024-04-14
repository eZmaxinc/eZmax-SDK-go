# EzsignfoldersignerassociationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**BEzsignfoldersignerassociationDelayedsend** | **bool** | If this flag is true the signatory is part of a delayed send. | 
**BEzsignfoldersignerassociationReceivecopy** | **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | 
**TEzsignfoldersignerassociationMessage** | **string** | A custom text message that will be added to the email sent. | 
**BEzsignfoldersignerassociationAllowsigninginperson** | **bool** | If the Ezsignfoldersignerassociation is allowed to sign in person or not | 
**ObjEzsignsignergroup** | Pointer to [**EzsignsignergroupResponseCompound**](EzsignsignergroupResponseCompound.md) |  | [optional] 
**ObjUser** | Pointer to [**EzsignfoldersignerassociationResponseCompoundUser**](EzsignfoldersignerassociationResponseCompoundUser.md) |  | [optional] 
**ObjEzsignsigner** | Pointer to [**EzsignsignerResponseCompound**](EzsignsignerResponseCompound.md) |  | [optional] 

## Methods

### NewEzsignfoldersignerassociationResponseCompound

`func NewEzsignfoldersignerassociationResponseCompound(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, bEzsignfoldersignerassociationDelayedsend bool, bEzsignfoldersignerassociationReceivecopy bool, tEzsignfoldersignerassociationMessage string, bEzsignfoldersignerassociationAllowsigninginperson bool, ) *EzsignfoldersignerassociationResponseCompound`

NewEzsignfoldersignerassociationResponseCompound instantiates a new EzsignfoldersignerassociationResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationResponseCompoundWithDefaults

`func NewEzsignfoldersignerassociationResponseCompoundWithDefaults() *EzsignfoldersignerassociationResponseCompound`

NewEzsignfoldersignerassociationResponseCompoundWithDefaults instantiates a new EzsignfoldersignerassociationResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetBEzsignfoldersignerassociationDelayedsend

`func (o *EzsignfoldersignerassociationResponseCompound) GetBEzsignfoldersignerassociationDelayedsend() bool`

GetBEzsignfoldersignerassociationDelayedsend returns the BEzsignfoldersignerassociationDelayedsend field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationDelayedsendOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetBEzsignfoldersignerassociationDelayedsendOk() (*bool, bool)`

GetBEzsignfoldersignerassociationDelayedsendOk returns a tuple with the BEzsignfoldersignerassociationDelayedsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationDelayedsend

`func (o *EzsignfoldersignerassociationResponseCompound) SetBEzsignfoldersignerassociationDelayedsend(v bool)`

SetBEzsignfoldersignerassociationDelayedsend sets BEzsignfoldersignerassociationDelayedsend field to given value.


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


### GetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationResponseCompound) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationResponseCompound) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.


### GetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *EzsignfoldersignerassociationResponseCompound) GetBEzsignfoldersignerassociationAllowsigninginperson() bool`

GetBEzsignfoldersignerassociationAllowsigninginperson returns the BEzsignfoldersignerassociationAllowsigninginperson field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationAllowsigninginpersonOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetBEzsignfoldersignerassociationAllowsigninginpersonOk() (*bool, bool)`

GetBEzsignfoldersignerassociationAllowsigninginpersonOk returns a tuple with the BEzsignfoldersignerassociationAllowsigninginperson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *EzsignfoldersignerassociationResponseCompound) SetBEzsignfoldersignerassociationAllowsigninginperson(v bool)`

SetBEzsignfoldersignerassociationAllowsigninginperson sets BEzsignfoldersignerassociationAllowsigninginperson field to given value.


### GetObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationResponseCompound) GetObjEzsignsignergroup() EzsignsignergroupResponseCompound`

GetObjEzsignsignergroup returns the ObjEzsignsignergroup field if non-nil, zero value otherwise.

### GetObjEzsignsignergroupOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetObjEzsignsignergroupOk() (*EzsignsignergroupResponseCompound, bool)`

GetObjEzsignsignergroupOk returns a tuple with the ObjEzsignsignergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationResponseCompound) SetObjEzsignsignergroup(v EzsignsignergroupResponseCompound)`

SetObjEzsignsignergroup sets ObjEzsignsignergroup field to given value.

### HasObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationResponseCompound) HasObjEzsignsignergroup() bool`

HasObjEzsignsignergroup returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


