# ColleagueResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiColleagueID** | **int32** | The unique ID of the Colleague | 
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
**ObjUserName** | [**CustomUserNameResponse**](CustomUserNameResponse.md) |  | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewColleagueResponseV2

`func NewColleagueResponseV2(pkiColleagueID int32, fkiUserID int32, fkiUserIDColleague int32, bColleagueEzsignemail bool, bColleagueFinancial bool, bColleagueUsecloneemail bool, bColleagueAttachment bool, bColleagueCanafe bool, bColleaguePermission bool, bColleagueRealestatecompleted bool, eColleagueEzsign FieldEColleagueEzsign, eColleagueRealestateinprogress FieldEColleagueRealestateinprogess, objUserName CustomUserNameResponse, objAudit CommonAudit, ) *ColleagueResponseV2`

NewColleagueResponseV2 instantiates a new ColleagueResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColleagueResponseV2WithDefaults

`func NewColleagueResponseV2WithDefaults() *ColleagueResponseV2`

NewColleagueResponseV2WithDefaults instantiates a new ColleagueResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiColleagueID

`func (o *ColleagueResponseV2) GetPkiColleagueID() int32`

GetPkiColleagueID returns the PkiColleagueID field if non-nil, zero value otherwise.

### GetPkiColleagueIDOk

`func (o *ColleagueResponseV2) GetPkiColleagueIDOk() (*int32, bool)`

GetPkiColleagueIDOk returns a tuple with the PkiColleagueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiColleagueID

`func (o *ColleagueResponseV2) SetPkiColleagueID(v int32)`

SetPkiColleagueID sets PkiColleagueID field to given value.


### GetFkiUserID

`func (o *ColleagueResponseV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ColleagueResponseV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ColleagueResponseV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetFkiUserIDColleague

`func (o *ColleagueResponseV2) GetFkiUserIDColleague() int32`

GetFkiUserIDColleague returns the FkiUserIDColleague field if non-nil, zero value otherwise.

### GetFkiUserIDColleagueOk

`func (o *ColleagueResponseV2) GetFkiUserIDColleagueOk() (*int32, bool)`

GetFkiUserIDColleagueOk returns a tuple with the FkiUserIDColleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDColleague

`func (o *ColleagueResponseV2) SetFkiUserIDColleague(v int32)`

SetFkiUserIDColleague sets FkiUserIDColleague field to given value.


### GetBColleagueEzsignemail

`func (o *ColleagueResponseV2) GetBColleagueEzsignemail() bool`

GetBColleagueEzsignemail returns the BColleagueEzsignemail field if non-nil, zero value otherwise.

### GetBColleagueEzsignemailOk

`func (o *ColleagueResponseV2) GetBColleagueEzsignemailOk() (*bool, bool)`

GetBColleagueEzsignemailOk returns a tuple with the BColleagueEzsignemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueEzsignemail

`func (o *ColleagueResponseV2) SetBColleagueEzsignemail(v bool)`

SetBColleagueEzsignemail sets BColleagueEzsignemail field to given value.


### GetBColleagueFinancial

`func (o *ColleagueResponseV2) GetBColleagueFinancial() bool`

GetBColleagueFinancial returns the BColleagueFinancial field if non-nil, zero value otherwise.

### GetBColleagueFinancialOk

`func (o *ColleagueResponseV2) GetBColleagueFinancialOk() (*bool, bool)`

GetBColleagueFinancialOk returns a tuple with the BColleagueFinancial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueFinancial

`func (o *ColleagueResponseV2) SetBColleagueFinancial(v bool)`

SetBColleagueFinancial sets BColleagueFinancial field to given value.


### GetBColleagueUsecloneemail

`func (o *ColleagueResponseV2) GetBColleagueUsecloneemail() bool`

GetBColleagueUsecloneemail returns the BColleagueUsecloneemail field if non-nil, zero value otherwise.

### GetBColleagueUsecloneemailOk

`func (o *ColleagueResponseV2) GetBColleagueUsecloneemailOk() (*bool, bool)`

GetBColleagueUsecloneemailOk returns a tuple with the BColleagueUsecloneemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueUsecloneemail

`func (o *ColleagueResponseV2) SetBColleagueUsecloneemail(v bool)`

SetBColleagueUsecloneemail sets BColleagueUsecloneemail field to given value.


### GetBColleagueAttachment

`func (o *ColleagueResponseV2) GetBColleagueAttachment() bool`

GetBColleagueAttachment returns the BColleagueAttachment field if non-nil, zero value otherwise.

### GetBColleagueAttachmentOk

`func (o *ColleagueResponseV2) GetBColleagueAttachmentOk() (*bool, bool)`

GetBColleagueAttachmentOk returns a tuple with the BColleagueAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueAttachment

`func (o *ColleagueResponseV2) SetBColleagueAttachment(v bool)`

SetBColleagueAttachment sets BColleagueAttachment field to given value.


### GetBColleagueCanafe

`func (o *ColleagueResponseV2) GetBColleagueCanafe() bool`

GetBColleagueCanafe returns the BColleagueCanafe field if non-nil, zero value otherwise.

### GetBColleagueCanafeOk

`func (o *ColleagueResponseV2) GetBColleagueCanafeOk() (*bool, bool)`

GetBColleagueCanafeOk returns a tuple with the BColleagueCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueCanafe

`func (o *ColleagueResponseV2) SetBColleagueCanafe(v bool)`

SetBColleagueCanafe sets BColleagueCanafe field to given value.


### GetBColleaguePermission

`func (o *ColleagueResponseV2) GetBColleaguePermission() bool`

GetBColleaguePermission returns the BColleaguePermission field if non-nil, zero value otherwise.

### GetBColleaguePermissionOk

`func (o *ColleagueResponseV2) GetBColleaguePermissionOk() (*bool, bool)`

GetBColleaguePermissionOk returns a tuple with the BColleaguePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleaguePermission

`func (o *ColleagueResponseV2) SetBColleaguePermission(v bool)`

SetBColleaguePermission sets BColleaguePermission field to given value.


### GetBColleagueRealestatecompleted

`func (o *ColleagueResponseV2) GetBColleagueRealestatecompleted() bool`

GetBColleagueRealestatecompleted returns the BColleagueRealestatecompleted field if non-nil, zero value otherwise.

### GetBColleagueRealestatecompletedOk

`func (o *ColleagueResponseV2) GetBColleagueRealestatecompletedOk() (*bool, bool)`

GetBColleagueRealestatecompletedOk returns a tuple with the BColleagueRealestatecompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueRealestatecompleted

`func (o *ColleagueResponseV2) SetBColleagueRealestatecompleted(v bool)`

SetBColleagueRealestatecompleted sets BColleagueRealestatecompleted field to given value.


### GetDtColleagueFrom

`func (o *ColleagueResponseV2) GetDtColleagueFrom() string`

GetDtColleagueFrom returns the DtColleagueFrom field if non-nil, zero value otherwise.

### GetDtColleagueFromOk

`func (o *ColleagueResponseV2) GetDtColleagueFromOk() (*string, bool)`

GetDtColleagueFromOk returns a tuple with the DtColleagueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtColleagueFrom

`func (o *ColleagueResponseV2) SetDtColleagueFrom(v string)`

SetDtColleagueFrom sets DtColleagueFrom field to given value.

### HasDtColleagueFrom

`func (o *ColleagueResponseV2) HasDtColleagueFrom() bool`

HasDtColleagueFrom returns a boolean if a field has been set.

### GetDtColleagueTo

`func (o *ColleagueResponseV2) GetDtColleagueTo() string`

GetDtColleagueTo returns the DtColleagueTo field if non-nil, zero value otherwise.

### GetDtColleagueToOk

`func (o *ColleagueResponseV2) GetDtColleagueToOk() (*string, bool)`

GetDtColleagueToOk returns a tuple with the DtColleagueTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtColleagueTo

`func (o *ColleagueResponseV2) SetDtColleagueTo(v string)`

SetDtColleagueTo sets DtColleagueTo field to given value.

### HasDtColleagueTo

`func (o *ColleagueResponseV2) HasDtColleagueTo() bool`

HasDtColleagueTo returns a boolean if a field has been set.

### GetEColleagueEzsign

`func (o *ColleagueResponseV2) GetEColleagueEzsign() FieldEColleagueEzsign`

GetEColleagueEzsign returns the EColleagueEzsign field if non-nil, zero value otherwise.

### GetEColleagueEzsignOk

`func (o *ColleagueResponseV2) GetEColleagueEzsignOk() (*FieldEColleagueEzsign, bool)`

GetEColleagueEzsignOk returns a tuple with the EColleagueEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEColleagueEzsign

`func (o *ColleagueResponseV2) SetEColleagueEzsign(v FieldEColleagueEzsign)`

SetEColleagueEzsign sets EColleagueEzsign field to given value.


### GetEColleagueRealestateinprogress

`func (o *ColleagueResponseV2) GetEColleagueRealestateinprogress() FieldEColleagueRealestateinprogess`

GetEColleagueRealestateinprogress returns the EColleagueRealestateinprogress field if non-nil, zero value otherwise.

### GetEColleagueRealestateinprogressOk

`func (o *ColleagueResponseV2) GetEColleagueRealestateinprogressOk() (*FieldEColleagueRealestateinprogess, bool)`

GetEColleagueRealestateinprogressOk returns a tuple with the EColleagueRealestateinprogress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEColleagueRealestateinprogress

`func (o *ColleagueResponseV2) SetEColleagueRealestateinprogress(v FieldEColleagueRealestateinprogess)`

SetEColleagueRealestateinprogress sets EColleagueRealestateinprogress field to given value.


### GetObjUserName

`func (o *ColleagueResponseV2) GetObjUserName() CustomUserNameResponse`

GetObjUserName returns the ObjUserName field if non-nil, zero value otherwise.

### GetObjUserNameOk

`func (o *ColleagueResponseV2) GetObjUserNameOk() (*CustomUserNameResponse, bool)`

GetObjUserNameOk returns a tuple with the ObjUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUserName

`func (o *ColleagueResponseV2) SetObjUserName(v CustomUserNameResponse)`

SetObjUserName sets ObjUserName field to given value.


### GetObjAudit

`func (o *ColleagueResponseV2) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *ColleagueResponseV2) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *ColleagueResponseV2) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


