# CustomEzsignfoldersignerassociationActionableElementResponse

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
**BEzsignfoldersignerassociationHasactionableelementsCurrent** | **bool** | Indicates if the Ezsignfoldersignerassociation has actionable elements in the current step | 
**BEzsignfoldersignerassociationHasactionableelementsFuture** | **bool** | Indicates if the Ezsignfoldersignerassociation has actionable elements in a future step | 

## Methods

### NewCustomEzsignfoldersignerassociationActionableElementResponse

`func NewCustomEzsignfoldersignerassociationActionableElementResponse(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, bEzsignfoldersignerassociationDelayedsend bool, bEzsignfoldersignerassociationReceivecopy bool, tEzsignfoldersignerassociationMessage string, bEzsignfoldersignerassociationAllowsigninginperson bool, bEzsignfoldersignerassociationHasactionableelementsCurrent bool, bEzsignfoldersignerassociationHasactionableelementsFuture bool, ) *CustomEzsignfoldersignerassociationActionableElementResponse`

NewCustomEzsignfoldersignerassociationActionableElementResponse instantiates a new CustomEzsignfoldersignerassociationActionableElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignfoldersignerassociationActionableElementResponseWithDefaults

`func NewCustomEzsignfoldersignerassociationActionableElementResponseWithDefaults() *CustomEzsignfoldersignerassociationActionableElementResponse`

NewCustomEzsignfoldersignerassociationActionableElementResponseWithDefaults instantiates a new CustomEzsignfoldersignerassociationActionableElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldersignerassociationID

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.


### GetFkiEzsignfolderID

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetBEzsignfoldersignerassociationDelayedsend

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationDelayedsend() bool`

GetBEzsignfoldersignerassociationDelayedsend returns the BEzsignfoldersignerassociationDelayedsend field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationDelayedsendOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationDelayedsendOk() (*bool, bool)`

GetBEzsignfoldersignerassociationDelayedsendOk returns a tuple with the BEzsignfoldersignerassociationDelayedsend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationDelayedsend

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetBEzsignfoldersignerassociationDelayedsend(v bool)`

SetBEzsignfoldersignerassociationDelayedsend sets BEzsignfoldersignerassociationDelayedsend field to given value.


### GetBEzsignfoldersignerassociationReceivecopy

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationReceivecopy() bool`

GetBEzsignfoldersignerassociationReceivecopy returns the BEzsignfoldersignerassociationReceivecopy field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationReceivecopyOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationReceivecopyOk() (*bool, bool)`

GetBEzsignfoldersignerassociationReceivecopyOk returns a tuple with the BEzsignfoldersignerassociationReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationReceivecopy

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetBEzsignfoldersignerassociationReceivecopy(v bool)`

SetBEzsignfoldersignerassociationReceivecopy sets BEzsignfoldersignerassociationReceivecopy field to given value.


### GetTEzsignfoldersignerassociationMessage

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.


### GetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationAllowsigninginperson() bool`

GetBEzsignfoldersignerassociationAllowsigninginperson returns the BEzsignfoldersignerassociationAllowsigninginperson field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationAllowsigninginpersonOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationAllowsigninginpersonOk() (*bool, bool)`

GetBEzsignfoldersignerassociationAllowsigninginpersonOk returns a tuple with the BEzsignfoldersignerassociationAllowsigninginperson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationAllowsigninginperson

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetBEzsignfoldersignerassociationAllowsigninginperson(v bool)`

SetBEzsignfoldersignerassociationAllowsigninginperson sets BEzsignfoldersignerassociationAllowsigninginperson field to given value.


### GetObjEzsignsignergroup

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetObjEzsignsignergroup() EzsignsignergroupResponseCompound`

GetObjEzsignsignergroup returns the ObjEzsignsignergroup field if non-nil, zero value otherwise.

### GetObjEzsignsignergroupOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetObjEzsignsignergroupOk() (*EzsignsignergroupResponseCompound, bool)`

GetObjEzsignsignergroupOk returns a tuple with the ObjEzsignsignergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsignergroup

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetObjEzsignsignergroup(v EzsignsignergroupResponseCompound)`

SetObjEzsignsignergroup sets ObjEzsignsignergroup field to given value.

### HasObjEzsignsignergroup

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) HasObjEzsignsignergroup() bool`

HasObjEzsignsignergroup returns a boolean if a field has been set.

### GetObjUser

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetObjUser() EzsignfoldersignerassociationResponseCompoundUser`

GetObjUser returns the ObjUser field if non-nil, zero value otherwise.

### GetObjUserOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetObjUserOk() (*EzsignfoldersignerassociationResponseCompoundUser, bool)`

GetObjUserOk returns a tuple with the ObjUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUser

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetObjUser(v EzsignfoldersignerassociationResponseCompoundUser)`

SetObjUser sets ObjUser field to given value.

### HasObjUser

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) HasObjUser() bool`

HasObjUser returns a boolean if a field has been set.

### GetObjEzsignsigner

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetObjEzsignsigner() EzsignsignerResponseCompound`

GetObjEzsignsigner returns the ObjEzsignsigner field if non-nil, zero value otherwise.

### GetObjEzsignsignerOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetObjEzsignsignerOk() (*EzsignsignerResponseCompound, bool)`

GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsigner

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetObjEzsignsigner(v EzsignsignerResponseCompound)`

SetObjEzsignsigner sets ObjEzsignsigner field to given value.

### HasObjEzsignsigner

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) HasObjEzsignsigner() bool`

HasObjEzsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldersignerassociationHasactionableelementsCurrent

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationHasactionableelementsCurrent() bool`

GetBEzsignfoldersignerassociationHasactionableelementsCurrent returns the BEzsignfoldersignerassociationHasactionableelementsCurrent field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationHasactionableelementsCurrentOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationHasactionableelementsCurrentOk() (*bool, bool)`

GetBEzsignfoldersignerassociationHasactionableelementsCurrentOk returns a tuple with the BEzsignfoldersignerassociationHasactionableelementsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationHasactionableelementsCurrent

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetBEzsignfoldersignerassociationHasactionableelementsCurrent(v bool)`

SetBEzsignfoldersignerassociationHasactionableelementsCurrent sets BEzsignfoldersignerassociationHasactionableelementsCurrent field to given value.


### GetBEzsignfoldersignerassociationHasactionableelementsFuture

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationHasactionableelementsFuture() bool`

GetBEzsignfoldersignerassociationHasactionableelementsFuture returns the BEzsignfoldersignerassociationHasactionableelementsFuture field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationHasactionableelementsFutureOk

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) GetBEzsignfoldersignerassociationHasactionableelementsFutureOk() (*bool, bool)`

GetBEzsignfoldersignerassociationHasactionableelementsFutureOk returns a tuple with the BEzsignfoldersignerassociationHasactionableelementsFuture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationHasactionableelementsFuture

`func (o *CustomEzsignfoldersignerassociationActionableElementResponse) SetBEzsignfoldersignerassociationHasactionableelementsFuture(v bool)`

SetBEzsignfoldersignerassociationHasactionableelementsFuture sets BEzsignfoldersignerassociationHasactionableelementsFuture field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


