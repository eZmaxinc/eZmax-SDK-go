# ColleagueRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiColleagueID** | Pointer to **int32** | The unique ID of the Colleague | [optional] 
**FkiUserID** | **int32** | The unique ID of the User | 
**FkiUserIDColleague** | **int32** | The unique ID of the User | 
**BColleagueEzsignemail** | **bool** | Whether the email can be used by the cloning user in Ezsign | 
**BColleagueFinancial** | **bool** | Whether the cloning user has access to the financial | 
**BColleagueUsecloneemail** | **bool** | Whether the cloning user has access to the cloned user email to send communications | 
**BColleagueAttachment** | **bool** | Whether the cloning user has access to the attachment | 
**BColleagueCanafe** | **bool** | Whether the cloning user has access to canafe | 
**BColleaguePermission** | **bool** | Whether the cloning user copies the permission of the cloned user | 
**BColleagueRealestatecompleted** | **bool** | Whether if the cloning user has access to the completed folders in real estate | 
**DtColleagueFrom** | Pointer to **string** | The from of the Colleague | [optional] 
**DtColleagueTo** | Pointer to **string** | The to of the Colleague | [optional] 
**EColleagueEzsign** | [**FieldEColleagueEzsign**](FieldEColleagueEzsign.md) |  | 
**EColleagueRealestateinprogress** | [**FieldEColleagueRealestateinprogess**](FieldEColleagueRealestateinprogess.md) |  | 

## Methods

### NewColleagueRequestV2

`func NewColleagueRequestV2(fkiUserID int32, fkiUserIDColleague int32, bColleagueEzsignemail bool, bColleagueFinancial bool, bColleagueUsecloneemail bool, bColleagueAttachment bool, bColleagueCanafe bool, bColleaguePermission bool, bColleagueRealestatecompleted bool, eColleagueEzsign FieldEColleagueEzsign, eColleagueRealestateinprogress FieldEColleagueRealestateinprogess, ) *ColleagueRequestV2`

NewColleagueRequestV2 instantiates a new ColleagueRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColleagueRequestV2WithDefaults

`func NewColleagueRequestV2WithDefaults() *ColleagueRequestV2`

NewColleagueRequestV2WithDefaults instantiates a new ColleagueRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiColleagueID

`func (o *ColleagueRequestV2) GetPkiColleagueID() int32`

GetPkiColleagueID returns the PkiColleagueID field if non-nil, zero value otherwise.

### GetPkiColleagueIDOk

`func (o *ColleagueRequestV2) GetPkiColleagueIDOk() (*int32, bool)`

GetPkiColleagueIDOk returns a tuple with the PkiColleagueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiColleagueID

`func (o *ColleagueRequestV2) SetPkiColleagueID(v int32)`

SetPkiColleagueID sets PkiColleagueID field to given value.

### HasPkiColleagueID

`func (o *ColleagueRequestV2) HasPkiColleagueID() bool`

HasPkiColleagueID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *ColleagueRequestV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ColleagueRequestV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ColleagueRequestV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetFkiUserIDColleague

`func (o *ColleagueRequestV2) GetFkiUserIDColleague() int32`

GetFkiUserIDColleague returns the FkiUserIDColleague field if non-nil, zero value otherwise.

### GetFkiUserIDColleagueOk

`func (o *ColleagueRequestV2) GetFkiUserIDColleagueOk() (*int32, bool)`

GetFkiUserIDColleagueOk returns a tuple with the FkiUserIDColleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDColleague

`func (o *ColleagueRequestV2) SetFkiUserIDColleague(v int32)`

SetFkiUserIDColleague sets FkiUserIDColleague field to given value.


### GetBColleagueEzsignemail

`func (o *ColleagueRequestV2) GetBColleagueEzsignemail() bool`

GetBColleagueEzsignemail returns the BColleagueEzsignemail field if non-nil, zero value otherwise.

### GetBColleagueEzsignemailOk

`func (o *ColleagueRequestV2) GetBColleagueEzsignemailOk() (*bool, bool)`

GetBColleagueEzsignemailOk returns a tuple with the BColleagueEzsignemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueEzsignemail

`func (o *ColleagueRequestV2) SetBColleagueEzsignemail(v bool)`

SetBColleagueEzsignemail sets BColleagueEzsignemail field to given value.


### GetBColleagueFinancial

`func (o *ColleagueRequestV2) GetBColleagueFinancial() bool`

GetBColleagueFinancial returns the BColleagueFinancial field if non-nil, zero value otherwise.

### GetBColleagueFinancialOk

`func (o *ColleagueRequestV2) GetBColleagueFinancialOk() (*bool, bool)`

GetBColleagueFinancialOk returns a tuple with the BColleagueFinancial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueFinancial

`func (o *ColleagueRequestV2) SetBColleagueFinancial(v bool)`

SetBColleagueFinancial sets BColleagueFinancial field to given value.


### GetBColleagueUsecloneemail

`func (o *ColleagueRequestV2) GetBColleagueUsecloneemail() bool`

GetBColleagueUsecloneemail returns the BColleagueUsecloneemail field if non-nil, zero value otherwise.

### GetBColleagueUsecloneemailOk

`func (o *ColleagueRequestV2) GetBColleagueUsecloneemailOk() (*bool, bool)`

GetBColleagueUsecloneemailOk returns a tuple with the BColleagueUsecloneemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueUsecloneemail

`func (o *ColleagueRequestV2) SetBColleagueUsecloneemail(v bool)`

SetBColleagueUsecloneemail sets BColleagueUsecloneemail field to given value.


### GetBColleagueAttachment

`func (o *ColleagueRequestV2) GetBColleagueAttachment() bool`

GetBColleagueAttachment returns the BColleagueAttachment field if non-nil, zero value otherwise.

### GetBColleagueAttachmentOk

`func (o *ColleagueRequestV2) GetBColleagueAttachmentOk() (*bool, bool)`

GetBColleagueAttachmentOk returns a tuple with the BColleagueAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueAttachment

`func (o *ColleagueRequestV2) SetBColleagueAttachment(v bool)`

SetBColleagueAttachment sets BColleagueAttachment field to given value.


### GetBColleagueCanafe

`func (o *ColleagueRequestV2) GetBColleagueCanafe() bool`

GetBColleagueCanafe returns the BColleagueCanafe field if non-nil, zero value otherwise.

### GetBColleagueCanafeOk

`func (o *ColleagueRequestV2) GetBColleagueCanafeOk() (*bool, bool)`

GetBColleagueCanafeOk returns a tuple with the BColleagueCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueCanafe

`func (o *ColleagueRequestV2) SetBColleagueCanafe(v bool)`

SetBColleagueCanafe sets BColleagueCanafe field to given value.


### GetBColleaguePermission

`func (o *ColleagueRequestV2) GetBColleaguePermission() bool`

GetBColleaguePermission returns the BColleaguePermission field if non-nil, zero value otherwise.

### GetBColleaguePermissionOk

`func (o *ColleagueRequestV2) GetBColleaguePermissionOk() (*bool, bool)`

GetBColleaguePermissionOk returns a tuple with the BColleaguePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleaguePermission

`func (o *ColleagueRequestV2) SetBColleaguePermission(v bool)`

SetBColleaguePermission sets BColleaguePermission field to given value.


### GetBColleagueRealestatecompleted

`func (o *ColleagueRequestV2) GetBColleagueRealestatecompleted() bool`

GetBColleagueRealestatecompleted returns the BColleagueRealestatecompleted field if non-nil, zero value otherwise.

### GetBColleagueRealestatecompletedOk

`func (o *ColleagueRequestV2) GetBColleagueRealestatecompletedOk() (*bool, bool)`

GetBColleagueRealestatecompletedOk returns a tuple with the BColleagueRealestatecompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueRealestatecompleted

`func (o *ColleagueRequestV2) SetBColleagueRealestatecompleted(v bool)`

SetBColleagueRealestatecompleted sets BColleagueRealestatecompleted field to given value.


### GetDtColleagueFrom

`func (o *ColleagueRequestV2) GetDtColleagueFrom() string`

GetDtColleagueFrom returns the DtColleagueFrom field if non-nil, zero value otherwise.

### GetDtColleagueFromOk

`func (o *ColleagueRequestV2) GetDtColleagueFromOk() (*string, bool)`

GetDtColleagueFromOk returns a tuple with the DtColleagueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtColleagueFrom

`func (o *ColleagueRequestV2) SetDtColleagueFrom(v string)`

SetDtColleagueFrom sets DtColleagueFrom field to given value.

### HasDtColleagueFrom

`func (o *ColleagueRequestV2) HasDtColleagueFrom() bool`

HasDtColleagueFrom returns a boolean if a field has been set.

### GetDtColleagueTo

`func (o *ColleagueRequestV2) GetDtColleagueTo() string`

GetDtColleagueTo returns the DtColleagueTo field if non-nil, zero value otherwise.

### GetDtColleagueToOk

`func (o *ColleagueRequestV2) GetDtColleagueToOk() (*string, bool)`

GetDtColleagueToOk returns a tuple with the DtColleagueTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtColleagueTo

`func (o *ColleagueRequestV2) SetDtColleagueTo(v string)`

SetDtColleagueTo sets DtColleagueTo field to given value.

### HasDtColleagueTo

`func (o *ColleagueRequestV2) HasDtColleagueTo() bool`

HasDtColleagueTo returns a boolean if a field has been set.

### GetEColleagueEzsign

`func (o *ColleagueRequestV2) GetEColleagueEzsign() FieldEColleagueEzsign`

GetEColleagueEzsign returns the EColleagueEzsign field if non-nil, zero value otherwise.

### GetEColleagueEzsignOk

`func (o *ColleagueRequestV2) GetEColleagueEzsignOk() (*FieldEColleagueEzsign, bool)`

GetEColleagueEzsignOk returns a tuple with the EColleagueEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEColleagueEzsign

`func (o *ColleagueRequestV2) SetEColleagueEzsign(v FieldEColleagueEzsign)`

SetEColleagueEzsign sets EColleagueEzsign field to given value.


### GetEColleagueRealestateinprogress

`func (o *ColleagueRequestV2) GetEColleagueRealestateinprogress() FieldEColleagueRealestateinprogess`

GetEColleagueRealestateinprogress returns the EColleagueRealestateinprogress field if non-nil, zero value otherwise.

### GetEColleagueRealestateinprogressOk

`func (o *ColleagueRequestV2) GetEColleagueRealestateinprogressOk() (*FieldEColleagueRealestateinprogess, bool)`

GetEColleagueRealestateinprogressOk returns a tuple with the EColleagueRealestateinprogress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEColleagueRealestateinprogress

`func (o *ColleagueRequestV2) SetEColleagueRealestateinprogress(v FieldEColleagueRealestateinprogess)`

SetEColleagueRealestateinprogress sets EColleagueRealestateinprogress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


