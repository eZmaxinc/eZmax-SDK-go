# EzsignfoldersignerassociationGetObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**FkiEzsignsignerID** | **NullableInt32** | The unique ID of the Ezsignsigner | 
**FkiUserID** | **NullableInt32** | The unique ID of the User | 
**BEzsignfoldersignerassociationReceivecopy** | **bool** | If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain&#39;t required to sign the document. | 

## Methods

### NewEzsignfoldersignerassociationGetObjectV1ResponseMPayload

`func NewEzsignfoldersignerassociationGetObjectV1ResponseMPayload(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, fkiEzsignsignerID NullableInt32, fkiUserID NullableInt32, bEzsignfoldersignerassociationReceivecopy bool, ) *EzsignfoldersignerassociationGetObjectV1ResponseMPayload`

NewEzsignfoldersignerassociationGetObjectV1ResponseMPayload instantiates a new EzsignfoldersignerassociationGetObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationGetObjectV1ResponseMPayloadWithDefaults

`func NewEzsignfoldersignerassociationGetObjectV1ResponseMPayloadWithDefaults() *EzsignfoldersignerassociationGetObjectV1ResponseMPayload`

NewEzsignfoldersignerassociationGetObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignfoldersignerassociationGetObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetPkiEzsignfoldersignerassociationID() int32`

GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldersignerassociationID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetPkiEzsignfoldersignerassociationID(v int32)`

SetPkiEzsignfoldersignerassociationID sets PkiEzsignfoldersignerassociationID field to given value.


### GetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetFkiEzsignsignerID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetFkiEzsignsignerID() int32`

GetFkiEzsignsignerID returns the FkiEzsignsignerID field if non-nil, zero value otherwise.

### GetFkiEzsignsignerIDOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetFkiEzsignsignerIDOk() (*int32, bool)`

GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignerID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetFkiEzsignsignerID(v int32)`

SetFkiEzsignsignerID sets FkiEzsignsignerID field to given value.


### SetFkiEzsignsignerIDNil

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetFkiEzsignsignerIDNil(b bool)`

 SetFkiEzsignsignerIDNil sets the value for FkiEzsignsignerID to be an explicit nil

### UnsetFkiEzsignsignerID
`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) UnsetFkiEzsignsignerID()`

UnsetFkiEzsignsignerID ensures that no value is present for FkiEzsignsignerID, not even an explicit nil
### GetFkiUserID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### SetFkiUserIDNil

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetFkiUserIDNil(b bool)`

 SetFkiUserIDNil sets the value for FkiUserID to be an explicit nil

### UnsetFkiUserID
`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) UnsetFkiUserID()`

UnsetFkiUserID ensures that no value is present for FkiUserID, not even an explicit nil
### GetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetBEzsignfoldersignerassociationReceivecopy() bool`

GetBEzsignfoldersignerassociationReceivecopy returns the BEzsignfoldersignerassociationReceivecopy field if non-nil, zero value otherwise.

### GetBEzsignfoldersignerassociationReceivecopyOk

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) GetBEzsignfoldersignerassociationReceivecopyOk() (*bool, bool)`

GetBEzsignfoldersignerassociationReceivecopyOk returns a tuple with the BEzsignfoldersignerassociationReceivecopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldersignerassociationReceivecopy

`func (o *EzsignfoldersignerassociationGetObjectV1ResponseMPayload) SetBEzsignfoldersignerassociationReceivecopy(v bool)`

SetBEzsignfoldersignerassociationReceivecopy sets BEzsignfoldersignerassociationReceivecopy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


