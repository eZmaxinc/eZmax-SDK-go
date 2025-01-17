# EzsignfoldersignerassociationRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldersignerassociationID** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiEzsignsignergroupID** | Pointer to **int32** | The unique ID of the Ezsignsignergroup | [optional] 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**BEzsignfoldersignerassociationReceivecopy** | Pointer to **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | [optional] 
**TEzsignfoldersignerassociationMessage** | Pointer to **string** | A custom text message that will be added to the email sent. | [optional] 
**ObjEzsignsigner** | Pointer to [**EzsignsignerRequestCompound**](EzsignsignerRequestCompound.md) |  | [optional] 

## Methods

### NewEzsignfoldersignerassociationRequestCompound

`func NewEzsignfoldersignerassociationRequestCompound(fkiEzsignfolderID int32, ) *EzsignfoldersignerassociationRequestCompound`

NewEzsignfoldersignerassociationRequestCompound instantiates a new EzsignfoldersignerassociationRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationRequestCompoundWithDefaults

`func NewEzsignfoldersignerassociationRequestCompoundWithDefaults() *EzsignfoldersignerassociationRequestCompound`

NewEzsignfoldersignerassociationRequestCompoundWithDefaults instantiates a new EzsignfoldersignerassociationRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestCompound) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfoldersignerassociationRequestCompound) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestCompound) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.

### HasPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestCompound) HasPkiEzsignfoldersignerassociationID() bool`

HasPkiEzsignfoldersignerassociationID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *EzsignfoldersignerassociationRequestCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsignfoldersignerassociationRequestCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsignfoldersignerassociationRequestCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsignfoldersignerassociationRequestCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignsignergroupID() int32`

GetFkiEzsignsignergroupID returns the FkiEzsignsignergroupID field if non-nil, zero value otherwise.

### GetFkiEzsignsignergroupIDOk

`func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignsignergroupIDOk() (*int32, bool)`

GetFkiEzsignsignergroupIDOk returns a tuple with the FkiEzsignsignergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestCompound) SetFkiEzsignsignergroupID(v int32)`

SetFkiEzsignsignergroupID sets FkiEzsignsignergroupID field to given value.

### HasFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestCompound) HasFkiEzsignsignergroupID() bool`

HasFkiEzsignsignergroupID returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationRequestCompound) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationRequestCompound) GetBEzsignfoldersignerassociationReceivecopy() bool`

GetBEzsignfoldersignerassociationReceivecopy returns the BEzsignfoldersignerassociationReceivecopy field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationReceivecopyOk

`func (o *EzsignfoldersignerassociationRequestCompound) GetBEzsignfoldersignerassociationReceivecopyOk() (*bool, bool)`

GetBEzsignfoldersignerassociationReceivecopyOk returns a tuple with the BEzsignfoldersignerassociationReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationRequestCompound) SetBEzsignfoldersignerassociationReceivecopy(v bool)`

SetBEzsignfoldersignerassociationReceivecopy sets BEzsignfoldersignerassociationReceivecopy field to given value.

### HasBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationRequestCompound) HasBEzsignfoldersignerassociationReceivecopy() bool`

HasBEzsignfoldersignerassociationReceivecopy returns a boolean if a field has been set.

### GetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestCompound) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *EzsignfoldersignerassociationRequestCompound) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestCompound) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.

### HasTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestCompound) HasTEzsignfoldersignerassociationMessage() bool`

HasTEzsignfoldersignerassociationMessage returns a boolean if a field has been set.

### GetObjEzsignsigner

`func (o *EzsignfoldersignerassociationRequestCompound) GetObjEzsignsigner() EzsignsignerRequestCompound`

GetObjEzsignsigner returns the ObjEzsignsigner field if non-nil, zero value otherwise.

### GetObjEzsignsignerOk

`func (o *EzsignfoldersignerassociationRequestCompound) GetObjEzsignsignerOk() (*EzsignsignerRequestCompound, bool)`

GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsigner

`func (o *EzsignfoldersignerassociationRequestCompound) SetObjEzsignsigner(v EzsignsignerRequestCompound)`

SetObjEzsignsigner sets ObjEzsignsigner field to given value.

### HasObjEzsignsigner

`func (o *EzsignfoldersignerassociationRequestCompound) HasObjEzsignsigner() bool`

HasObjEzsignsigner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


