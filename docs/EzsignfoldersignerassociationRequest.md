# EzsignfoldersignerassociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldersignerassociationID** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiEzsignsignergroupID** | Pointer to **int32** | The unique ID of the Ezsignsignergroup | [optional] 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**BEzsignfoldersignerassociationReceivecopy** | Pointer to **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | [optional] 
**TEzsignfoldersignerassociationMessage** | Pointer to **string** | A custom text message that will be added to the email sent. | [optional] 

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

### GetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequest) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfoldersignerassociationRequest) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequest) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.

### HasPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationRequest) HasPkiEzsignfoldersignerassociationID() bool`

HasPkiEzsignfoldersignerassociationID returns a boolean if a field has been set.

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

### GetFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequest) GetFkiEzsignsignergroupID() int32`

GetFkiEzsignsignergroupID returns the FkiEzsignsignergroupID field if non-nil, zero value otherwise.

### GetFkiEzsignsignergroupIDOk

`func (o *EzsignfoldersignerassociationRequest) GetFkiEzsignsignergroupIDOk() (*int32, bool)`

GetFkiEzsignsignergroupIDOk returns a tuple with the FkiEzsignsignergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequest) SetFkiEzsignsignergroupID(v int32)`

SetFkiEzsignsignergroupID sets FkiEzsignsignergroupID field to given value.

### HasFkiEzsignsignergroupID

`func (o *EzsignfoldersignerassociationRequest) HasFkiEzsignsignergroupID() bool`

HasFkiEzsignsignergroupID returns a boolean if a field has been set.

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


### GetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationRequest) GetBEzsignfoldersignerassociationReceivecopy() bool`

GetBEzsignfoldersignerassociationReceivecopy returns the BEzsignfoldersignerassociationReceivecopy field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationReceivecopyOk

`func (o *EzsignfoldersignerassociationRequest) GetBEzsignfoldersignerassociationReceivecopyOk() (*bool, bool)`

GetBEzsignfoldersignerassociationReceivecopyOk returns a tuple with the BEzsignfoldersignerassociationReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationRequest) SetBEzsignfoldersignerassociationReceivecopy(v bool)`

SetBEzsignfoldersignerassociationReceivecopy sets BEzsignfoldersignerassociationReceivecopy field to given value.

### HasBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationRequest) HasBEzsignfoldersignerassociationReceivecopy() bool`

HasBEzsignfoldersignerassociationReceivecopy returns a boolean if a field has been set.

### GetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequest) GetTEzsignfoldersignerassociationMessage() string`

GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field if non-nil, zero value otherwise.

### GetTEzsignfoldersignerassociationMessageOk

`func (o *EzsignfoldersignerassociationRequest) GetTEzsignfoldersignerassociationMessageOk() (*string, bool)`

GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequest) SetTEzsignfoldersignerassociationMessage(v string)`

SetTEzsignfoldersignerassociationMessage sets TEzsignfoldersignerassociationMessage field to given value.

### HasTEzsignfoldersignerassociationMessage

`func (o *EzsignfoldersignerassociationRequest) HasTEzsignfoldersignerassociationMessage() bool`

HasTEzsignfoldersignerassociationMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


