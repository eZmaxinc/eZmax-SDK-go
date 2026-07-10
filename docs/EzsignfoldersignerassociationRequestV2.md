# EzsignfoldersignerassociationRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldersignerassociationID** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiEzsignsignergroupID** | Pointer to **int32** | The unique ID of the Ezsignsignergroup | [optional] 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**EEzsignfoldersignerassociationRole** | [**FieldEEzsignfoldersignerassociationRole**](FieldEEzsignfoldersignerassociationRole.md) |  | 
**TEzsignfoldersignerassociationMessage** | Pointer to **string** | A custom text message that will be added to the email sent. | [optional] 

## Methods

### NewEzsignfoldersignerassociationRequestV2

`func NewEzsignfoldersignerassociationRequestV2(fkiEzsignfolderID int32, eEzsignfoldersignerassociationRole FieldEEzsignfoldersignerassociationRole, ) *EzsignfoldersignerassociationRequestV2`

NewEzsignfoldersignerassociationRequestV2 instantiates a new EzsignfoldersignerassociationRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationRequestV2WithDefaults

`func NewEzsignfoldersignerassociationRequestV2WithDefaults() *EzsignfoldersignerassociationRequestV2`

NewEzsignfoldersignerassociationRequestV2WithDefaults instantiates a new EzsignfoldersignerassociationRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestV2) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfoldersignerassociationRequestV2) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestV2) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.

### HasPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestV2) HasPkiEzsignfoldersignerassociationID() bool`

HasPkiEzsignfoldersignerassociationID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *EzsignfoldersignerassociationRequestV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsignfoldersignerassociationRequestV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsignfoldersignerassociationRequestV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsignfoldersignerassociationRequestV2) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestV2) GetFkiEzsignsignergroupID() int32`

GetFkiEzsignsignergroupID returns the FkiEzsignsignergroupID field if non-nil, zero value otherwise.

### GetFkiEzsignsignergroupIDOk

`func (o *EzsignfoldersignerassociationRequestV2) GetFkiEzsignsignergroupIDOk() (*int32, bool)`

GetFkiEzsignsignergroupIDOk returns a tuple with the FkiEzsignsignergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestV2) SetFkiEzsignsignergroupID(v int32)`

SetFkiEzsignsignergroupID sets FkiEzsignsignergroupID field to given value.

### HasFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestV2) HasFkiEzsignsignergroupID() bool`

HasFkiEzsignsignergroupID returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationRequestV2) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignfoldersignerassociationRequestV2) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationRequestV2) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetEEzsignfoldersignerassociationRole

`func (o *EzsignfoldersignerassociationRequestV2) GetEEzsignfoldersignerassociationRole() FieldEEzsignfoldersignerassociationRole`

GetEEzsignfoldersignerassociationRole returns the EEzsignfoldersignerassociationRole field if non-nil, zero value otherwise.

### GetEEzsignfoldersignerassociationRoleOk

`func (o *EzsignfoldersignerassociationRequestV2) GetEEzsignfoldersignerassociationRoleOk() (*FieldEEzsignfoldersignerassociationRole, bool)`

GetEEzsignfoldersignerassociationRoleOk returns a tuple with the EEzsignfoldersignerassociationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldersignerassociationRole

`func (o *EzsignfoldersignerassociationRequestV2) SetEEzsignfoldersignerassociationRole(v FieldEEzsignfoldersignerassociationRole)`

SetEEzsignfoldersignerassociationRole sets EEzsignfoldersignerassociationRole field to given value.


### GetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestV2) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *EzsignfoldersignerassociationRequestV2) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestV2) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.

### HasTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestV2) HasTEzsignfoldersignerassociationMessage() bool`

HasTEzsignfoldersignerassociationMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


