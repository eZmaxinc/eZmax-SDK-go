# EzsignfoldersignerassociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiUserID** | Pointer to **int32** | A reference to a valid User.  This is only used if the signatory will be a user from the system. | [optional] 
**FkiEzsignfolderID** | **int32** | A reference to a valid Ezsignfolder.  That value is returned after a successful Ezsignfolder Creation. | 

## Methods

### NewEzsignfoldersignerassociationRequest

`func NewEzsignfoldersignerassociationRequest(fkiEzsignfolderID int32, ) *EzsignfoldersignerassociationRequest`

NewEzsignfoldersignerassociationRequest instantiates a new EzsignfoldersignerassociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationRequestWithDefaults

`func NewEzsignfoldersignerassociationRequestWithDefaults() *EzsignfoldersignerassociationRequest`

NewEzsignfoldersignerassociationRequestWithDefaults instantiates a new EzsignfoldersignerassociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiUserID

`func (o *EzsignfoldersignerassociationRequest) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsignfoldersignerassociationRequest) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsignfoldersignerassociationRequest) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsignfoldersignerassociationRequest) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationRequest) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignfoldersignerassociationRequest) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationRequest) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


