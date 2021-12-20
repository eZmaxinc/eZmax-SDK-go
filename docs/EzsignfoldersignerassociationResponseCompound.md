# EzsignfoldersignerassociationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsignsignerID** | **NullableInt32** | The unique ID of the Ezsignsigner | 
**FkiUserID** | **NullableInt32** | The unique ID of the User | 
**BEzsignfoldersignerassociationReceivecopy** | **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | 

## Methods

### NewEzsignfoldersignerassociationResponseCompound

`func NewEzsignfoldersignerassociationResponseCompound(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, fkiEzsignsignerID NullableInt32, fkiUserID NullableInt32, bEzsignfoldersignerassociationReceivecopy bool, ) *EzsignfoldersignerassociationResponseCompound`

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


### GetFkiEzsignsignerID

`func (o *EzsignfoldersignerassociationResponseCompound) GetFkiEzsignsignerID() int32`

GetFkiEzsignsignerID returns the FkiEzsignsignerID field if non-nil, zero value otherwise.

### GetFkiEzsignsignerIDOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetFkiEzsignsignerIDOk() (*int32, bool)`

GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignerID

`func (o *EzsignfoldersignerassociationResponseCompound) SetFkiEzsignsignerID(v int32)`

SetFkiEzsignsignerID sets FkiEzsignsignerID field to given value.


### SetFkiEzsignsignerIDNil

`func (o *EzsignfoldersignerassociationResponseCompound) SetFkiEzsignsignerIDNil(b bool)`

 SetFkiEzsignsignerIDNil sets the value for FkiEzsignsignerID to be an explicit nil

### UnsetFkiEzsignsignerID
`func (o *EzsignfoldersignerassociationResponseCompound) UnsetFkiEzsignsignerID()`

UnsetFkiEzsignsignerID ensures that no value is present for FkiEzsignsignerID, not even an explicit nil
### GetFkiUserID

`func (o *EzsignfoldersignerassociationResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsignfoldersignerassociationResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsignfoldersignerassociationResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### SetFkiUserIDNil

`func (o *EzsignfoldersignerassociationResponseCompound) SetFkiUserIDNil(b bool)`

 SetFkiUserIDNil sets the value for FkiUserID to be an explicit nil

### UnsetFkiUserID
`func (o *EzsignfoldersignerassociationResponseCompound) UnsetFkiUserID()`

UnsetFkiUserID ensures that no value is present for FkiUserID, not even an explicit nil
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


