# EzsignfoldersignerassociationResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**BEzsignfoldersignerassociationDelayedsend** | **bool** | If this flag is true the signatory is part of a delayed send. | 
**EEzsignfoldersignerassociationRole** | [**FieldEEzsignfoldersignerassociationRole**](FieldEEzsignfoldersignerassociationRole.md) |  | 
**TEzsignfoldersignerassociationMessage** | **string** | A custom text message that will be added to the email sent. | 
**BEzsignfoldersignerassociationAllowsigninginperson** | **bool** | If the Ezsignfoldersignerassociation is allowed to sign in person or not | 
**ObjEzsignsignergroup** | Pointer to [**EzsignsignergroupResponseCompound**](EzsignsignergroupResponseCompound.md) |  | [optional] 
**ObjUser** | Pointer to [**EzsignfoldersignerassociationResponseCompoundUser**](EzsignfoldersignerassociationResponseCompoundUser.md) |  | [optional] 
**ObjEzsignsigner** | Pointer to [**EzsignsignerResponseCompound**](EzsignsignerResponseCompound.md) |  | [optional] 

## Methods

### NewEzsignfoldersignerassociationResponseCompoundV3

`func NewEzsignfoldersignerassociationResponseCompoundV3(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, bEzsignfoldersignerassociationDelayedsend bool, eEzsignfoldersignerassociationRole FieldEEzsignfoldersignerassociationRole, tEzsignfoldersignerassociationMessage string, bEzsignfoldersignerassociationAllowsigninginperson bool, ) *EzsignfoldersignerassociationResponseCompoundV3`

NewEzsignfoldersignerassociationResponseCompoundV3 instantiates a new EzsignfoldersignerassociationResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationResponseCompoundV3WithDefaults

`func NewEzsignfoldersignerassociationResponseCompoundV3WithDefaults() *EzsignfoldersignerassociationResponseCompoundV3`

NewEzsignfoldersignerassociationResponseCompoundV3WithDefaults instantiates a new EzsignfoldersignerassociationResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetBEzsignfoldersignerassociationDelayedsend

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetBEzsignfoldersignerassociationDelayedsend() bool`

GetBEzsignfoldersignerassociationDelayedsend returns the BEzsignfoldersignerassociationDelayedsend field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationDelayedsendOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetBEzsignfoldersignerassociationDelayedsendOk() (*bool, bool)`

GetBEzsignfoldersignerassociationDelayedsendOk returns a tuple with the BEzsignfoldersignerassociationDelayedsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationDelayedsend

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetBEzsignfoldersignerassociationDelayedsend(v bool)`

SetBEzsignfoldersignerassociationDelayedsend sets BEzsignfoldersignerassociationDelayedsend field to given value.


### GetEEzsignfoldersignerassociationRole

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetEEzsignfoldersignerassociationRole() FieldEEzsignfoldersignerassociationRole`

GetEEzsignfoldersignerassociationRole returns the EEzsignfoldersignerassociationRole field if non-nil, zero value otherwise.

### GetEEzsignfoldersignerassociationRoleOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetEEzsignfoldersignerassociationRoleOk() (*FieldEEzsignfoldersignerassociationRole, bool)`

GetEEzsignfoldersignerassociationRoleOk returns a tuple with the EEzsignfoldersignerassociationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldersignerassociationRole

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetEEzsignfoldersignerassociationRole(v FieldEEzsignfoldersignerassociationRole)`

SetEEzsignfoldersignerassociationRole sets EEzsignfoldersignerassociationRole field to given value.


### GetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.


### GetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetBEzsignfoldersignerassociationAllowsigninginperson() bool`

GetBEzsignfoldersignerassociationAllowsigninginperson returns the BEzsignfoldersignerassociationAllowsigninginperson field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationAllowsigninginpersonOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetBEzsignfoldersignerassociationAllowsigninginpersonOk() (*bool, bool)`

GetBEzsignfoldersignerassociationAllowsigninginpersonOk returns a tuple with the BEzsignfoldersignerassociationAllowsigninginperson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetBEzsignfoldersignerassociationAllowsigninginperson(v bool)`

SetBEzsignfoldersignerassociationAllowsigninginperson sets BEzsignfoldersignerassociationAllowsigninginperson field to given value.


### GetObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetObjEzsignsignergroup() EzsignsignergroupResponseCompound`

GetObjEzsignsignergroup returns the ObjEzsignsignergroup field if non-nil, zero value otherwise.

### GetObjEzsignsignergroupOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetObjEzsignsignergroupOk() (*EzsignsignergroupResponseCompound, bool)`

GetObjEzsignsignergroupOk returns a tuple with the ObjEzsignsignergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetObjEzsignsignergroup(v EzsignsignergroupResponseCompound)`

SetObjEzsignsignergroup sets ObjEzsignsignergroup field to given value.

### HasObjEzsignsignergroup

`func (o *EzsignfoldersignerassociationResponseCompoundV3) HasObjEzsignsignergroup() bool`

HasObjEzsignsignergroup returns a boolean if a field has been set.

### GetObjUser

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetObjUser() EzsignfoldersignerassociationResponseCompoundUser`

GetObjUser returns the ObjUser field if non-nil, zero value otherwise.

### GetObjUserOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetObjUserOk() (*EzsignfoldersignerassociationResponseCompoundUser, bool)`

GetObjUserOk returns a tuple with the ObjUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUser

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetObjUser(v EzsignfoldersignerassociationResponseCompoundUser)`

SetObjUser sets ObjUser field to given value.

### HasObjUser

`func (o *EzsignfoldersignerassociationResponseCompoundV3) HasObjUser() bool`

HasObjUser returns a boolean if a field has been set.

### GetObjEzsignsigner

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetObjEzsignsigner() EzsignsignerResponseCompound`

GetObjEzsignsigner returns the ObjEzsignsigner field if non-nil, zero value otherwise.

### GetObjEzsignsignerOk

`func (o *EzsignfoldersignerassociationResponseCompoundV3) GetObjEzsignsignerOk() (*EzsignsignerResponseCompound, bool)`

GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsigner

`func (o *EzsignfoldersignerassociationResponseCompoundV3) SetObjEzsignsigner(v EzsignsignerResponseCompound)`

SetObjEzsignsigner sets ObjEzsignsigner field to given value.

### HasObjEzsignsigner

`func (o *EzsignfoldersignerassociationResponseCompoundV3) HasObjEzsignsigner() bool`

HasObjEzsignsigner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


