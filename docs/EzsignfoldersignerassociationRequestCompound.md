# EzsignfoldersignerassociationRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjEzsignsigner** | Pointer to [**EzsignsignerRequestCompound**](EzsignsignerRequestCompound.md) |  | [optional] 
**FkiUserID** | Pointer to **int32** | A reference to a valid User.  This is only used if the signatory will be a user from the system. | [optional] 
**FkiEzsignfolderID** | **int32** | A reference to a valid Ezsignfolder.  That value is returned after a successful Ezsignfolder Creation. | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


