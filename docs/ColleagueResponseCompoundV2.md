# ColleagueResponseCompoundV2

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

### NewColleagueResponseCompoundV2

`func NewColleagueResponseCompoundV2(pkiColleagueID int32, fkiUserID int32, fkiUserIDColleague int32, bColleagueEzsignemail bool, bColleagueFinancial bool, bColleagueUsecloneemail bool, bColleagueAttachment bool, bColleagueCanafe bool, bColleaguePermission bool, bColleagueRealestatecompleted bool, eColleagueEzsign FieldEColleagueEzsign, eColleagueRealestateinprogress FieldEColleagueRealestateinprogess, objUserName CustomUserNameResponse, objAudit CommonAudit, ) *ColleagueResponseCompoundV2`

NewColleagueResponseCompoundV2 instantiates a new ColleagueResponseCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColleagueResponseCompoundV2WithDefaults

`func NewColleagueResponseCompoundV2WithDefaults() *ColleagueResponseCompoundV2`

NewColleagueResponseCompoundV2WithDefaults instantiates a new ColleagueResponseCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiColleagueID

`func (o *ColleagueResponseCompoundV2) GetPkiColleagueID() int32`

GetPkiColleagueID returns the PkiColleagueID field if non-nil, zero value otherwise.

### GetPkiColleagueIDOk

`func (o *ColleagueResponseCompoundV2) GetPkiColleagueIDOk() (*int32, bool)`

GetPkiColleagueIDOk returns a tuple with the PkiColleagueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiColleagueID

`func (o *ColleagueResponseCompoundV2) SetPkiColleagueID(v int32)`

SetPkiColleagueID sets PkiColleagueID field to given value.


### GetFkiUserID

`func (o *ColleagueResponseCompoundV2) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ColleagueResponseCompoundV2) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ColleagueResponseCompoundV2) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetFkiUserIDColleague

`func (o *ColleagueResponseCompoundV2) GetFkiUserIDColleague() int32`

GetFkiUserIDColleague returns the FkiUserIDColleague field if non-nil, zero value otherwise.

### GetFkiUserIDColleagueOk

`func (o *ColleagueResponseCompoundV2) GetFkiUserIDColleagueOk() (*int32, bool)`

GetFkiUserIDColleagueOk returns a tuple with the FkiUserIDColleague field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDColleague

`func (o *ColleagueResponseCompoundV2) SetFkiUserIDColleague(v int32)`

SetFkiUserIDColleague sets FkiUserIDColleague field to given value.


### GetBColleagueEzsignemail

`func (o *ColleagueResponseCompoundV2) GetBColleagueEzsignemail() bool`

GetBColleagueEzsignemail returns the BColleagueEzsignemail field if non-nil, zero value otherwise.

### GetBColleagueEzsignemailOk

`func (o *ColleagueResponseCompoundV2) GetBColleagueEzsignemailOk() (*bool, bool)`

GetBColleagueEzsignemailOk returns a tuple with the BColleagueEzsignemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueEzsignemail

`func (o *ColleagueResponseCompoundV2) SetBColleagueEzsignemail(v bool)`

SetBColleagueEzsignemail sets BColleagueEzsignemail field to given value.


### GetBColleagueFinancial

`func (o *ColleagueResponseCompoundV2) GetBColleagueFinancial() bool`

GetBColleagueFinancial returns the BColleagueFinancial field if non-nil, zero value otherwise.

### GetBColleagueFinancialOk

`func (o *ColleagueResponseCompoundV2) GetBColleagueFinancialOk() (*bool, bool)`

GetBColleagueFinancialOk returns a tuple with the BColleagueFinancial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueFinancial

`func (o *ColleagueResponseCompoundV2) SetBColleagueFinancial(v bool)`

SetBColleagueFinancial sets BColleagueFinancial field to given value.


### GetBColleagueUsecloneemail

`func (o *ColleagueResponseCompoundV2) GetBColleagueUsecloneemail() bool`

GetBColleagueUsecloneemail returns the BColleagueUsecloneemail field if non-nil, zero value otherwise.

### GetBColleagueUsecloneemailOk

`func (o *ColleagueResponseCompoundV2) GetBColleagueUsecloneemailOk() (*bool, bool)`

GetBColleagueUsecloneemailOk returns a tuple with the BColleagueUsecloneemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueUsecloneemail

`func (o *ColleagueResponseCompoundV2) SetBColleagueUsecloneemail(v bool)`

SetBColleagueUsecloneemail sets BColleagueUsecloneemail field to given value.


### GetBColleagueAttachment

`func (o *ColleagueResponseCompoundV2) GetBColleagueAttachment() bool`

GetBColleagueAttachment returns the BColleagueAttachment field if non-nil, zero value otherwise.

### GetBColleagueAttachmentOk

`func (o *ColleagueResponseCompoundV2) GetBColleagueAttachmentOk() (*bool, bool)`

GetBColleagueAttachmentOk returns a tuple with the BColleagueAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueAttachment

`func (o *ColleagueResponseCompoundV2) SetBColleagueAttachment(v bool)`

SetBColleagueAttachment sets BColleagueAttachment field to given value.


### GetBColleagueCanafe

`func (o *ColleagueResponseCompoundV2) GetBColleagueCanafe() bool`

GetBColleagueCanafe returns the BColleagueCanafe field if non-nil, zero value otherwise.

### GetBColleagueCanafeOk

`func (o *ColleagueResponseCompoundV2) GetBColleagueCanafeOk() (*bool, bool)`

GetBColleagueCanafeOk returns a tuple with the BColleagueCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueCanafe

`func (o *ColleagueResponseCompoundV2) SetBColleagueCanafe(v bool)`

SetBColleagueCanafe sets BColleagueCanafe field to given value.


### GetBColleaguePermission

`func (o *ColleagueResponseCompoundV2) GetBColleaguePermission() bool`

GetBColleaguePermission returns the BColleaguePermission field if non-nil, zero value otherwise.

### GetBColleaguePermissionOk

`func (o *ColleagueResponseCompoundV2) GetBColleaguePermissionOk() (*bool, bool)`

GetBColleaguePermissionOk returns a tuple with the BColleaguePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleaguePermission

`func (o *ColleagueResponseCompoundV2) SetBColleaguePermission(v bool)`

SetBColleaguePermission sets BColleaguePermission field to given value.


### GetBColleagueRealestatecompleted

`func (o *ColleagueResponseCompoundV2) GetBColleagueRealestatecompleted() bool`

GetBColleagueRealestatecompleted returns the BColleagueRealestatecompleted field if non-nil, zero value otherwise.

### GetBColleagueRealestatecompletedOk

`func (o *ColleagueResponseCompoundV2) GetBColleagueRealestatecompletedOk() (*bool, bool)`

GetBColleagueRealestatecompletedOk returns a tuple with the BColleagueRealestatecompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBColleagueRealestatecompleted

`func (o *ColleagueResponseCompoundV2) SetBColleagueRealestatecompleted(v bool)`

SetBColleagueRealestatecompleted sets BColleagueRealestatecompleted field to given value.


### GetDtColleagueFrom

`func (o *ColleagueResponseCompoundV2) GetDtColleagueFrom() string`

GetDtColleagueFrom returns the DtColleagueFrom field if non-nil, zero value otherwise.

### GetDtColleagueFromOk

`func (o *ColleagueResponseCompoundV2) GetDtColleagueFromOk() (*string, bool)`

GetDtColleagueFromOk returns a tuple with the DtColleagueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtColleagueFrom

`func (o *ColleagueResponseCompoundV2) SetDtColleagueFrom(v string)`

SetDtColleagueFrom sets DtColleagueFrom field to given value.

### HasDtColleagueFrom

`func (o *ColleagueResponseCompoundV2) HasDtColleagueFrom() bool`

HasDtColleagueFrom returns a boolean if a field has been set.

### GetDtColleagueTo

`func (o *ColleagueResponseCompoundV2) GetDtColleagueTo() string`

GetDtColleagueTo returns the DtColleagueTo field if non-nil, zero value otherwise.

### GetDtColleagueToOk

`func (o *ColleagueResponseCompoundV2) GetDtColleagueToOk() (*string, bool)`

GetDtColleagueToOk returns a tuple with the DtColleagueTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtColleagueTo

`func (o *ColleagueResponseCompoundV2) SetDtColleagueTo(v string)`

SetDtColleagueTo sets DtColleagueTo field to given value.

### HasDtColleagueTo

`func (o *ColleagueResponseCompoundV2) HasDtColleagueTo() bool`

HasDtColleagueTo returns a boolean if a field has been set.

### GetEColleagueEzsign

`func (o *ColleagueResponseCompoundV2) GetEColleagueEzsign() FieldEColleagueEzsign`

GetEColleagueEzsign returns the EColleagueEzsign field if non-nil, zero value otherwise.

### GetEColleagueEzsignOk

`func (o *ColleagueResponseCompoundV2) GetEColleagueEzsignOk() (*FieldEColleagueEzsign, bool)`

GetEColleagueEzsignOk returns a tuple with the EColleagueEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEColleagueEzsign

`func (o *ColleagueResponseCompoundV2) SetEColleagueEzsign(v FieldEColleagueEzsign)`

SetEColleagueEzsign sets EColleagueEzsign field to given value.


### GetEColleagueRealestateinprogress

`func (o *ColleagueResponseCompoundV2) GetEColleagueRealestateinprogress() FieldEColleagueRealestateinprogess`

GetEColleagueRealestateinprogress returns the EColleagueRealestateinprogress field if non-nil, zero value otherwise.

### GetEColleagueRealestateinprogressOk

`func (o *ColleagueResponseCompoundV2) GetEColleagueRealestateinprogressOk() (*FieldEColleagueRealestateinprogess, bool)`

GetEColleagueRealestateinprogressOk returns a tuple with the EColleagueRealestateinprogress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEColleagueRealestateinprogress

`func (o *ColleagueResponseCompoundV2) SetEColleagueRealestateinprogress(v FieldEColleagueRealestateinprogess)`

SetEColleagueRealestateinprogress sets EColleagueRealestateinprogress field to given value.


### GetObjUserName

`func (o *ColleagueResponseCompoundV2) GetObjUserName() CustomUserNameResponse`

GetObjUserName returns the ObjUserName field if non-nil, zero value otherwise.

### GetObjUserNameOk

`func (o *ColleagueResponseCompoundV2) GetObjUserNameOk() (*CustomUserNameResponse, bool)`

GetObjUserNameOk returns a tuple with the ObjUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUserName

`func (o *ColleagueResponseCompoundV2) SetObjUserName(v CustomUserNameResponse)`

SetObjUserName sets ObjUserName field to given value.


### GetObjAudit

`func (o *ColleagueResponseCompoundV2) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *ColleagueResponseCompoundV2) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *ColleagueResponseCompoundV2) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


