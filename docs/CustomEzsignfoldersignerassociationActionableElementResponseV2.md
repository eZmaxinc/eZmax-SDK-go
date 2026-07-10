# CustomEzsignfoldersignerassociationActionableElementResponseV2

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
**BEzsignfoldersignerassociationHasactionableelementsCurrent** | **bool** | Indicates if the Ezsignfoldersignerassociation has actionable elements in the current step | 
**BEzsignfoldersignerassociationHasactionableelementsFuture** | **bool** | Indicates if the Ezsignfoldersignerassociation has actionable elements in a future step | 

## Methods

### NewCustomEzsignfoldersignerassociationActionableElementResponseV2

`func NewCustomEzsignfoldersignerassociationActionableElementResponseV2(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, bEzsignfoldersignerassociationDelayedsend bool, eEzsignfoldersignerassociationRole FieldEEzsignfoldersignerassociationRole, tEzsignfoldersignerassociationMessage string, bEzsignfoldersignerassociationAllowsigninginperson bool, bEzsignfoldersignerassociationHasactionableelementsCurrent bool, bEzsignfoldersignerassociationHasactionableelementsFuture bool, ) *CustomEzsignfoldersignerassociationActionableElementResponseV2`

NewCustomEzsignfoldersignerassociationActionableElementResponseV2 instantiates a new CustomEzsignfoldersignerassociationActionableElementResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignfoldersignerassociationActionableElementResponseV2WithDefaults

`func NewCustomEzsignfoldersignerassociationActionableElementResponseV2WithDefaults() *CustomEzsignfoldersignerassociationActionableElementResponseV2`

NewCustomEzsignfoldersignerassociationActionableElementResponseV2WithDefaults instantiates a new CustomEzsignfoldersignerassociationActionableElementResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldersignerassociationID

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.


### GetFkiEzsignfolderID

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetBEzsignfoldersignerassociationDelayedsend

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetBEzsignfoldersignerassociationDelayedsend() bool`

GetBEzsignfoldersignerassociationDelayedsend returns the BEzsignfoldersignerassociationDelayedsend field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationDelayedsendOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetBEzsignfoldersignerassociationDelayedsendOk() (*bool, bool)`

GetBEzsignfoldersignerassociationDelayedsendOk returns a tuple with the BEzsignfoldersignerassociationDelayedsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationDelayedsend

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetBEzsignfoldersignerassociationDelayedsend(v bool)`

SetBEzsignfoldersignerassociationDelayedsend sets BEzsignfoldersignerassociationDelayedsend field to given value.


### GetEEzsignfoldersignerassociationRole

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetEEzsignfoldersignerassociationRole() FieldEEzsignfoldersignerassociationRole`

GetEEzsignfoldersignerassociationRole returns the EEzsignfoldersignerassociationRole field if non-nil, zero value otherwise.

### GetEEzsignfoldersignerassociationRoleOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetEEzsignfoldersignerassociationRoleOk() (*FieldEEzsignfoldersignerassociationRole, bool)`

GetEEzsignfoldersignerassociationRoleOk returns a tuple with the EEzsignfoldersignerassociationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignfoldersignerassociationRole

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetEEzsignfoldersignerassociationRole(v FieldEEzsignfoldersignerassociationRole)`

SetEEzsignfoldersignerassociationRole sets EEzsignfoldersignerassociationRole field to given value.


### GetTEzsignfoldersignerassociationMessage

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.


### GetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetBEzsignfoldersignerassociationAllowsigninginperson() bool`

GetBEzsignfoldersignerassociationAllowsigninginperson returns the BEzsignfoldersignerassociationAllowsigninginperson field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationAllowsigninginpersonOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetBEzsignfoldersignerassociationAllowsigninginpersonOk() (*bool, bool)`

GetBEzsignfoldersignerassociationAllowsigninginpersonOk returns a tuple with the BEzsignfoldersignerassociationAllowsigninginperson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetBEzsignfoldersignerassociationAllowsigninginperson(v bool)`

SetBEzsignfoldersignerassociationAllowsigninginperson sets BEzsignfoldersignerassociationAllowsigninginperson field to given value.


### GetObjEzsignsignergroup

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetObjEzsignsignergroup() EzsignsignergroupResponseCompound`

GetObjEzsignsignergroup returns the ObjEzsignsignergroup field if non-nil, zero value otherwise.

### GetObjEzsignsignergroupOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetObjEzsignsignergroupOk() (*EzsignsignergroupResponseCompound, bool)`

GetObjEzsignsignergroupOk returns a tuple with the ObjEzsignsignergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsignergroup

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetObjEzsignsignergroup(v EzsignsignergroupResponseCompound)`

SetObjEzsignsignergroup sets ObjEzsignsignergroup field to given value.

### HasObjEzsignsignergroup

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) HasObjEzsignsignergroup() bool`

HasObjEzsignsignergroup returns a boolean if a field has been set.

### GetObjUser

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetObjUser() EzsignfoldersignerassociationResponseCompoundUser`

GetObjUser returns the ObjUser field if non-nil, zero value otherwise.

### GetObjUserOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetObjUserOk() (*EzsignfoldersignerassociationResponseCompoundUser, bool)`

GetObjUserOk returns a tuple with the ObjUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUser

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetObjUser(v EzsignfoldersignerassociationResponseCompoundUser)`

SetObjUser sets ObjUser field to given value.

### HasObjUser

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) HasObjUser() bool`

HasObjUser returns a boolean if a field has been set.

### GetObjEzsignsigner

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetObjEzsignsigner() EzsignsignerResponseCompound`

GetObjEzsignsigner returns the ObjEzsignsigner field if non-nil, zero value otherwise.

### GetObjEzsignsignerOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetObjEzsignsignerOk() (*EzsignsignerResponseCompound, bool)`

GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsigner

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetObjEzsignsigner(v EzsignsignerResponseCompound)`

SetObjEzsignsigner sets ObjEzsignsigner field to given value.

### HasObjEzsignsigner

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) HasObjEzsignsigner() bool`

HasObjEzsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldersignerassociationHasactionableelementsCurrent

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetBEzsignfoldersignerassociationHasactionableelementsCurrent() bool`

GetBEzsignfoldersignerassociationHasactionableelementsCurrent returns the BEzsignfoldersignerassociationHasactionableelementsCurrent field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationHasactionableelementsCurrentOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetBEzsignfoldersignerassociationHasactionableelementsCurrentOk() (*bool, bool)`

GetBEzsignfoldersignerassociationHasactionableelementsCurrentOk returns a tuple with the BEzsignfoldersignerassociationHasactionableelementsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationHasactionableelementsCurrent

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetBEzsignfoldersignerassociationHasactionableelementsCurrent(v bool)`

SetBEzsignfoldersignerassociationHasactionableelementsCurrent sets BEzsignfoldersignerassociationHasactionableelementsCurrent field to given value.


### GetBEzsignfoldersignerassociationHasactionableelementsFuture

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetBEzsignfoldersignerassociationHasactionableelementsFuture() bool`

GetBEzsignfoldersignerassociationHasactionableelementsFuture returns the BEzsignfoldersignerassociationHasactionableelementsFuture field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationHasactionableelementsFutureOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) GetBEzsignfoldersignerassociationHasactionableelementsFutureOk() (*bool, bool)`

GetBEzsignfoldersignerassociationHasactionableelementsFutureOk returns a tuple with the BEzsignfoldersignerassociationHasactionableelementsFuture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationHasactionableelementsFuture

`func (o *CustomEzsignfoldersignerassociationActionableElementResponseV2) SetBEzsignfoldersignerassociationHasactionableelementsFuture(v bool)`

SetBEzsignfoldersignerassociationHasactionableelementsFuture sets BEzsignfoldersignerassociationHasactionableelementsFuture field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


