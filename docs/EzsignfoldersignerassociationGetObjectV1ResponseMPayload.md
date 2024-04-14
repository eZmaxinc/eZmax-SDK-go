# EzsignfoldersignerassociationGetObjectV1ResponseMPayload

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

### NewEzsignfoldersignerassociationGetObjectV1ResponseMPayload

`func NewEzsignfoldersignerassociationGetObjectV1ResponseMPayload(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, bEzsignfoldersignerassociationDelayedsend bool, bEzsignfoldersignerassociationReceivecopy bool, tEzsignfoldersignerassociationMessage string, bEzsignfoldersignerassociationAllowsigninginperson bool, ) *EzsignfoldersignerassociationGetObjectV1ResponseMPayload`

NewEzsignfoldersignerassociationGetObjectV1ResponseMPayload instantiates a new EzsignfoldersignerassociationGetObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationGetObjectV1ResponseMPayloadWithDefaults

`func NewEzsignfoldersignerassociationGetObjectV1ResponseMPayloadWithDefaults() *EzsignfoldersignerassociationGetObjectV1ResponseMPayload`

NewEzsignfoldersignerassociationGetObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignfoldersignerassociationGetObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetBEzsignfoldersignerassociationDelayedsend

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetBEzsignfoldersignerassociationDelayedsend() bool`

GetBEzsignfoldersignerassociationDelayedsend returns the BEzsignfoldersignerassociationDelayedsend field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationDelayedsendOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetBEzsignfoldersignerassociationDelayedsendOk() (*bool, bool)`

GetBEzsignfoldersignerassociationDelayedsendOk returns a tuple with the BEzsignfoldersignerassociationDelayedsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationDelayedsend

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetBEzsignfoldersignerassociationDelayedsend(v bool)`

SetBEzsignfoldersignerassociationDelayedsend sets BEzsignfoldersignerassociationDelayedsend field to given value.


### GetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetBEzsignfoldersignerassociationReceivecopy() bool`

GetBEzsignfoldersignerassociationReceivecopy returns the BEzsignfoldersignerassociationReceivecopy field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationReceivecopyOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetBEzsignfoldersignerassociationReceivecopyOk() (*bool, bool)`

GetBEzsignfoldersignerassociationReceivecopyOk returns a tuple with the BEzsignfoldersignerassociationReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetBEzsignfoldersignerassociationReceivecopy(v bool)`

SetBEzsignfoldersignerassociationReceivecopy sets BEzsignfoldersignerassociationReceivecopy field to given value.


### GetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.


### GetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetBEzsignfoldersignerassociationAllowsigninginperson() bool`

GetBEzsignfoldersignerassociationAllowsigninginperson returns the BEzsignfoldersignerassociationAllowsigninginperson field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationAllowsigninginpersonOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetBEzsignfoldersignerassociationAllowsigninginpersonOk() (*bool, bool)`

GetBEzsignfoldersignerassociationAllowsigninginpersonOk returns a tuple with the BEzsignfoldersignerassociationAllowsigninginperson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetBEzsignfoldersignerassociationAllowsigninginperson(v bool)`

SetBEzsignfoldersignerassociationAllowsigninginperson sets BEzsignfoldersignerassociationAllowsigninginperson field to given value.


### GetObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetObjEzsignsignergroup() EzsignsignergroupResponseCompound`

GetObjEzsignsignergroup returns the ObjEzsignsignergroup field if non-nil, zero value otherwise.

### GetObjEzsignsignergroupOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetObjEzsignsignergroupOk() (*EzsignsignergroupResponseCompound, bool)`

GetObjEzsignsignergroupOk returns a tuple with the ObjEzsignsignergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetObjEzsignsignergroup(v EzsignsignergroupResponseCompound)`

SetObjEzsignsignergroup sets ObjEzsignsignergroup field to given value.

### HasObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) HasObjEzsignsignergroup() bool`

HasObjEzsignsignergroup returns a boolean if a field has been set.

### GetObjUser

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetObjUser() EzsignfoldersignerassociationResponseCompoundUser`

GetObjUser returns the ObjUser field if non-nil, zero value otherwise.

### GetObjUserOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetObjUserOk() (*EzsignfoldersignerassociationResponseCompoundUser, bool)`

GetObjUserOk returns a tuple with the ObjUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUser

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetObjUser(v EzsignfoldersignerassociationResponseCompoundUser)`

SetObjUser sets ObjUser field to given value.

### HasObjUser

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) HasObjUser() bool`

HasObjUser returns a boolean if a field has been set.

### GetObjEzsignsigner

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetObjEzsignsigner() EzsignsignerResponseCompound`

GetObjEzsignsigner returns the ObjEzsignsigner field if non-nil, zero value otherwise.

### GetObjEzsignsignerOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetObjEzsignsignerOk() (*EzsignsignerResponseCompound, bool)`

GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsigner

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetObjEzsignsigner(v EzsignsignerResponseCompound)`

SetObjEzsignsigner sets ObjEzsignsigner field to given value.

### HasObjEzsignsigner

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) HasObjEzsignsigner() bool`

HasObjEzsignsigner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


