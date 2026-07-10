# EzsignfoldersignerassociationRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldersignerassociationID** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiEzsignsignergroupID** | Pointer to **int32** | The unique ID of the Ezsignsignergroup | [optional] 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**EEzsignfoldersignerassociationRole** | [**FieldEEzsignfoldersignerassociationRole**](FieldEEzsignfoldersignerassociationRole.md) |  | 
**TEzsignfoldersignerassociationMessage** | Pointer to **string** | A custom text message that will be added to the email sent. | [optional] 
**ObjEzsignsigner** | Pointer to [**EzsignsignerRequestCompound**](EzsignsignerRequestCompound.md) |  | [optional] 

## Methods

### NewEzsignfoldersignerassociationRequestCompoundV2

`func NewEzsignfoldersignerassociationRequestCompoundV2(fkiEzsignfolderID int32, eEzsignfoldersignerassociationRole FieldEEzsignfoldersignerassociationRole, ) *EzsignfoldersignerassociationRequestCompoundV2`

NewEzsignfoldersignerassociationRequestCompoundV2 instantiates a new EzsignfoldersignerassociationRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationRequestCompoundV2WithDefaults

`func NewEzsignfoldersignerassociationRequestCompoundV2WithDefaults() *EzsignfoldersignerassociationRequestCompoundV2`

NewEzsignfoldersignerassociationRequestCompoundV2WithDefaults instantiates a new EzsignfoldersignerassociationRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.

### HasPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) HasPkiEzsignfoldersignerassociationID() bool`

HasPkiEzsignfoldersignerassociationID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetFkiEzsignsignergroupID() int32`

GetFkiEzsignsignergroupID returns the FkiEzsignsignergroupID field if non-nil, zero value otherwise.

### GetFkiEzsignsignergroupIDOk

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetFkiEzsignsignergroupIDOk() (*int32, bool)`

GetFkiEzsignsignergroupIDOk returns a tuple with the FkiEzsignsignergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) SetFkiEzsignsignergroupID(v int32)`

SetFkiEzsignsignergroupID sets FkiEzsignsignergroupID field to given value.

### HasFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) HasFkiEzsignsignergroupID() bool`

HasFkiEzsignsignergroupID returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationRequestCompoundV2) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetEEzsignfoldersignerassociationRole

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetEEzsignfoldersignerassociationRole() FieldEEzsignfoldersignerassociationRole`

GetEEzsignfoldersignerassociationRole returns the EEzsignfoldersignerassociationRole field if non-nil, zero value otherwise.

### GetEEzsignfoldersignerassociationRoleOk

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetEEzsignfoldersignerassociationRoleOk() (*FieldEEzsignfoldersignerassociationRole, bool)`

GetEEzsignfoldersignerassociationRoleOk returns a tuple with the EEzsignfoldersignerassociationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldersignerassociationRole

`func (o *EzsignfoldersignerassociationRequestCompoundV2) SetEEzsignfoldersignerassociationRole(v FieldEEzsignfoldersignerassociationRole)`

SetEEzsignfoldersignerassociationRole sets EEzsignfoldersignerassociationRole field to given value.


### GetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestCompoundV2) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.

### HasTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequestCompoundV2) HasTEzsignfoldersignerassociationMessage() bool`

HasTEzsignfoldersignerassociationMessage returns a boolean if a field has been set.

### GetObjEzsignsigner

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetObjEzsignsigner() EzsignsignerRequestCompound`

GetObjEzsignsigner returns the ObjEzsignsigner field if non-nil, zero value otherwise.

### GetObjEzsignsignerOk

`func (o *EzsignfoldersignerassociationRequestCompoundV2) GetObjEzsignsignerOk() (*EzsignsignerRequestCompound, bool)`

GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsigner

`func (o *EzsignfoldersignerassociationRequestCompoundV2) SetObjEzsignsigner(v EzsignsignerRequestCompound)`

SetObjEzsignsigner sets ObjEzsignsigner field to given value.

### HasObjEzsignsigner

`func (o *EzsignfoldersignerassociationRequestCompoundV2) HasObjEzsignsigner() bool`

HasObjEzsignsigner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


