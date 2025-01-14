# ColleagueRequestCompoundV2

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

### NewColleagueRequestCompoundV2

`func NewColleagueRequestCompoundV2(fkiUserID int32, fkiUserIDColleague int32, bColleagueEzsignemail bool, bColleagueFinancial bool, bColleagueUsecloneemail bool, bColleagueAttachment bool, bColleagueCanafe bool, bColleaguePermission bool, bColleagueRealestatecompleted bool, eColleagueEzsign FieldEColleagueEzsign, eColleagueRealestateinprogress FieldEColleagueRealestateinprogess, ) *ColleagueRequestCompoundV2`

NewColleagueRequestCompoundV2 instantiates a new ColleagueRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColleagueRequestCompoundV2WithDefaults

`func NewColleagueRequestCompoundV2WithDefaults() *ColleagueRequestCompoundV2`

NewColleagueRequestCompoundV2WithDefaults instantiates a new ColleagueRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiColleagueID

`func (o *ColleagueRequestCompoundV2) GetPkiColleagueID() int32`

GetPkiColleagueID returns the PkiColleagueID field if non-nil, zero value otherwise.

### GetPkiColleagueIDOk

`func (o *ColleagueRequestCompoundV2) GetPkiColleagueIDOk() (*int32, bool)`

GetPkiColleagueIDOk returns a tuple with the PkiColleagueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiColleagueID

`func (o *ColleagueRequestCompoundV2) SetPkiColleagueID(v int32)`

SetPkiColleagueID sets PkiColleagueID field to given value.

### HasPkiColleagueID

`func (o *ColleagueRequestCompoundV2) HasPkiColleagueID() bool`

HasPkiColleagueID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *ColleagueRequestCompoundV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ColleagueRequestCompoundV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ColleagueRequestCompoundV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetFkiUserIDColleague

`func (o *ColleagueRequestCompoundV2) GetFkiUserIDColleague() int32`

GetFkiUserIDColleague returns the FkiUserIDColleague field if non-nil, zero value otherwise.

### GetFkiUserIDColleagueOk

`func (o *ColleagueRequestCompoundV2) GetFkiUserIDColleagueOk() (*int32, bool)`

GetFkiUserIDColleagueOk returns a tuple with the FkiUserIDColleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDColleague

`func (o *ColleagueRequestCompoundV2) SetFkiUserIDColleague(v int32)`

SetFkiUserIDColleague sets FkiUserIDColleague field to given value.


### GetBColleagueEzsignemail

`func (o *ColleagueRequestCompoundV2) GetBColleagueEzsignemail() bool`

GetBColleagueEzsignemail returns the BColleagueEzsignemail field if non-nil, zero value otherwise.

### GetBColleagueEzsignemailOk

`func (o *ColleagueRequestCompoundV2) GetBColleagueEzsignemailOk() (*bool, bool)`

GetBColleagueEzsignemailOk returns a tuple with the BColleagueEzsignemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueEzsignemail

`func (o *ColleagueRequestCompoundV2) SetBColleagueEzsignemail(v bool)`

SetBColleagueEzsignemail sets BColleagueEzsignemail field to given value.


### GetBColleagueFinancial

`func (o *ColleagueRequestCompoundV2) GetBColleagueFinancial() bool`

GetBColleagueFinancial returns the BColleagueFinancial field if non-nil, zero value otherwise.

### GetBColleagueFinancialOk

`func (o *ColleagueRequestCompoundV2) GetBColleagueFinancialOk() (*bool, bool)`

GetBColleagueFinancialOk returns a tuple with the BColleagueFinancial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueFinancial

`func (o *ColleagueRequestCompoundV2) SetBColleagueFinancial(v bool)`

SetBColleagueFinancial sets BColleagueFinancial field to given value.


### GetBColleagueUsecloneemail

`func (o *ColleagueRequestCompoundV2) GetBColleagueUsecloneemail() bool`

GetBColleagueUsecloneemail returns the BColleagueUsecloneemail field if non-nil, zero value otherwise.

### GetBColleagueUsecloneemailOk

`func (o *ColleagueRequestCompoundV2) GetBColleagueUsecloneemailOk() (*bool, bool)`

GetBColleagueUsecloneemailOk returns a tuple with the BColleagueUsecloneemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueUsecloneemail

`func (o *ColleagueRequestCompoundV2) SetBColleagueUsecloneemail(v bool)`

SetBColleagueUsecloneemail sets BColleagueUsecloneemail field to given value.


### GetBColleagueAttachment

`func (o *ColleagueRequestCompoundV2) GetBColleagueAttachment() bool`

GetBColleagueAttachment returns the BColleagueAttachment field if non-nil, zero value otherwise.

### GetBColleagueAttachmentOk

`func (o *ColleagueRequestCompoundV2) GetBColleagueAttachmentOk() (*bool, bool)`

GetBColleagueAttachmentOk returns a tuple with the BColleagueAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueAttachment

`func (o *ColleagueRequestCompoundV2) SetBColleagueAttachment(v bool)`

SetBColleagueAttachment sets BColleagueAttachment field to given value.


### GetBColleagueCanafe

`func (o *ColleagueRequestCompoundV2) GetBColleagueCanafe() bool`

GetBColleagueCanafe returns the BColleagueCanafe field if non-nil, zero value otherwise.

### GetBColleagueCanafeOk

`func (o *ColleagueRequestCompoundV2) GetBColleagueCanafeOk() (*bool, bool)`

GetBColleagueCanafeOk returns a tuple with the BColleagueCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueCanafe

`func (o *ColleagueRequestCompoundV2) SetBColleagueCanafe(v bool)`

SetBColleagueCanafe sets BColleagueCanafe field to given value.


### GetBColleaguePermission

`func (o *ColleagueRequestCompoundV2) GetBColleaguePermission() bool`

GetBColleaguePermission returns the BColleaguePermission field if non-nil, zero value otherwise.

### GetBColleaguePermissionOk

`func (o *ColleagueRequestCompoundV2) GetBColleaguePermissionOk() (*bool, bool)`

GetBColleaguePermissionOk returns a tuple with the BColleaguePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleaguePermission

`func (o *ColleagueRequestCompoundV2) SetBColleaguePermission(v bool)`

SetBColleaguePermission sets BColleaguePermission field to given value.


### GetBColleagueRealestatecompleted

`func (o *ColleagueRequestCompoundV2) GetBColleagueRealestatecompleted() bool`

GetBColleagueRealestatecompleted returns the BColleagueRealestatecompleted field if non-nil, zero value otherwise.

### GetBColleagueRealestatecompletedOk

`func (o *ColleagueRequestCompoundV2) GetBColleagueRealestatecompletedOk() (*bool, bool)`

GetBColleagueRealestatecompletedOk returns a tuple with the BColleagueRealestatecompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueRealestatecompleted

`func (o *ColleagueRequestCompoundV2) SetBColleagueRealestatecompleted(v bool)`

SetBColleagueRealestatecompleted sets BColleagueRealestatecompleted field to given value.


### GetDtColleagueFrom

`func (o *ColleagueRequestCompoundV2) GetDtColleagueFrom() string`

GetDtColleagueFrom returns the DtColleagueFrom field if non-nil, zero value otherwise.

### GetDtColleagueFromOk

`func (o *ColleagueRequestCompoundV2) GetDtColleagueFromOk() (*string, bool)`

GetDtColleagueFromOk returns a tuple with the DtColleagueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtColleagueFrom

`func (o *ColleagueRequestCompoundV2) SetDtColleagueFrom(v string)`

SetDtColleagueFrom sets DtColleagueFrom field to given value.

### HasDtColleagueFrom

`func (o *ColleagueRequestCompoundV2) HasDtColleagueFrom() bool`

HasDtColleagueFrom returns a boolean if a field has been set.

### GetDtColleagueTo

`func (o *ColleagueRequestCompoundV2) GetDtColleagueTo() string`

GetDtColleagueTo returns the DtColleagueTo field if non-nil, zero value otherwise.

### GetDtColleagueToOk

`func (o *ColleagueRequestCompoundV2) GetDtColleagueToOk() (*string, bool)`

GetDtColleagueToOk returns a tuple with the DtColleagueTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtColleagueTo

`func (o *ColleagueRequestCompoundV2) SetDtColleagueTo(v string)`

SetDtColleagueTo sets DtColleagueTo field to given value.

### HasDtColleagueTo

`func (o *ColleagueRequestCompoundV2) HasDtColleagueTo() bool`

HasDtColleagueTo returns a boolean if a field has been set.

### GetEColleagueEzsign

`func (o *ColleagueRequestCompoundV2) GetEColleagueEzsign() FieldEColleagueEzsign`

GetEColleagueEzsign returns the EColleagueEzsign field if non-nil, zero value otherwise.

### GetEColleagueEzsignOk

`func (o *ColleagueRequestCompoundV2) GetEColleagueEzsignOk() (*FieldEColleagueEzsign, bool)`

GetEColleagueEzsignOk returns a tuple with the EColleagueEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEColleagueEzsign

`func (o *ColleagueRequestCompoundV2) SetEColleagueEzsign(v FieldEColleagueEzsign)`

SetEColleagueEzsign sets EColleagueEzsign field to given value.


### GetEColleagueRealestateinprogress

`func (o *ColleagueRequestCompoundV2) GetEColleagueRealestateinprogress() FieldEColleagueRealestateinprogess`

GetEColleagueRealestateinprogress returns the EColleagueRealestateinprogress field if non-nil, zero value otherwise.

### GetEColleagueRealestateinprogressOk

`func (o *ColleagueRequestCompoundV2) GetEColleagueRealestateinprogressOk() (*FieldEColleagueRealestateinprogess, bool)`

GetEColleagueRealestateinprogressOk returns a tuple with the EColleagueRealestateinprogress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEColleagueRealestateinprogress

`func (o *ColleagueRequestCompoundV2) SetEColleagueRealestateinprogress(v FieldEColleagueRealestateinprogess)`

SetEColleagueRealestateinprogress sets EColleagueRealestateinprogress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


