# AdjustmentreminderListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiAdjustmentreminderID** | **int32** | The unique ID of the Adjustmentreminder | 
**FkiAdjustmentID** | **int32** | The unique ID of the Adjustment | 
**FkiPeriodID** | **int32** | The unique ID of the Period | 
**SPeriodYYYYMM** | Pointer to **string** | The YYYYMM of the Period | [optional] 
**FkiUserIDCreated** | Pointer to **int32** | The unique ID of the User | [optional] 
**SUserLoginname** | Pointer to **string** | The login name of the User. | [optional] 
**EAdjustmentreminderType** | [**FieldEAdjustmentreminderType**](FieldEAdjustmentreminderType.md) |  | 
**SAdjustmentNumber** | Pointer to **string** | The number of the Adjustment | [optional] 
**SAdjustmentComment** | Pointer to **string** | The comment of the Adjustment | [optional] 
**DAdjustmentTotaldebit** | Pointer to **string** | The total debit of Adjustment | [optional] 
**DAdjustmentTotalcredit** | Pointer to **string** | The total credit of Adjustment | [optional] 

## Methods

### NewAdjustmentreminderListElement

`func NewAdjustmentreminderListElement(pkiAdjustmentreminderID int32, fkiAdjustmentID int32, fkiPeriodID int32, eAdjustmentreminderType FieldEAdjustmentreminderType, ) *AdjustmentreminderListElement`

NewAdjustmentreminderListElement instantiates a new AdjustmentreminderListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentreminderListElementWithDefaults

`func NewAdjustmentreminderListElementWithDefaults() *AdjustmentreminderListElement`

NewAdjustmentreminderListElementWithDefaults instantiates a new AdjustmentreminderListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAdjustmentreminderID

`func (o *AdjustmentreminderListElement) GetPkiAdjustmentreminderID() int32`

GetPkiAdjustmentreminderID returns the PkiAdjustmentreminderID field if non-nil, zero value otherwise.

### GetPkiAdjustmentreminderIDOk

`func (o *AdjustmentreminderListElement) GetPkiAdjustmentreminderIDOk() (*int32, bool)`

GetPkiAdjustmentreminderIDOk returns a tuple with the PkiAdjustmentreminderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAdjustmentreminderID

`func (o *AdjustmentreminderListElement) SetPkiAdjustmentreminderID(v int32)`

SetPkiAdjustmentreminderID sets PkiAdjustmentreminderID field to given value.


### GetFkiAdjustmentID

`func (o *AdjustmentreminderListElement) GetFkiAdjustmentID() int32`

GetFkiAdjustmentID returns the FkiAdjustmentID field if non-nil, zero value otherwise.

### GetFkiAdjustmentIDOk

`func (o *AdjustmentreminderListElement) GetFkiAdjustmentIDOk() (*int32, bool)`

GetFkiAdjustmentIDOk returns a tuple with the FkiAdjustmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAdjustmentID

`func (o *AdjustmentreminderListElement) SetFkiAdjustmentID(v int32)`

SetFkiAdjustmentID sets FkiAdjustmentID field to given value.


### GetFkiPeriodID

`func (o *AdjustmentreminderListElement) GetFkiPeriodID() int32`

GetFkiPeriodID returns the FkiPeriodID field if non-nil, zero value otherwise.

### GetFkiPeriodIDOk

`func (o *AdjustmentreminderListElement) GetFkiPeriodIDOk() (*int32, bool)`

GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPeriodID

`func (o *AdjustmentreminderListElement) SetFkiPeriodID(v int32)`

SetFkiPeriodID sets FkiPeriodID field to given value.


### GetSPeriodYYYYMM

`func (o *AdjustmentreminderListElement) GetSPeriodYYYYMM() string`

GetSPeriodYYYYMM returns the SPeriodYYYYMM field if non-nil, zero value otherwise.

### GetSPeriodYYYYMMOk

`func (o *AdjustmentreminderListElement) GetSPeriodYYYYMMOk() (*string, bool)`

GetSPeriodYYYYMMOk returns a tuple with the SPeriodYYYYMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPeriodYYYYMM

`func (o *AdjustmentreminderListElement) SetSPeriodYYYYMM(v string)`

SetSPeriodYYYYMM sets SPeriodYYYYMM field to given value.

### HasSPeriodYYYYMM

`func (o *AdjustmentreminderListElement) HasSPeriodYYYYMM() bool`

HasSPeriodYYYYMM returns a boolean if a field has been set.

### GetFkiUserIDCreated

`func (o *AdjustmentreminderListElement) GetFkiUserIDCreated() int32`

GetFkiUserIDCreated returns the FkiUserIDCreated field if non-nil, zero value otherwise.

### GetFkiUserIDCreatedOk

`func (o *AdjustmentreminderListElement) GetFkiUserIDCreatedOk() (*int32, bool)`

GetFkiUserIDCreatedOk returns a tuple with the FkiUserIDCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDCreated

`func (o *AdjustmentreminderListElement) SetFkiUserIDCreated(v int32)`

SetFkiUserIDCreated sets FkiUserIDCreated field to given value.

### HasFkiUserIDCreated

`func (o *AdjustmentreminderListElement) HasFkiUserIDCreated() bool`

HasFkiUserIDCreated returns a boolean if a field has been set.

### GetSUserLoginname

`func (o *AdjustmentreminderListElement) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *AdjustmentreminderListElement) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *AdjustmentreminderListElement) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.

### HasSUserLoginname

`func (o *AdjustmentreminderListElement) HasSUserLoginname() bool`

HasSUserLoginname returns a boolean if a field has been set.

### GetEAdjustmentreminderType

`func (o *AdjustmentreminderListElement) GetEAdjustmentreminderType() FieldEAdjustmentreminderType`

GetEAdjustmentreminderType returns the EAdjustmentreminderType field if non-nil, zero value otherwise.

### GetEAdjustmentreminderTypeOk

`func (o *AdjustmentreminderListElement) GetEAdjustmentreminderTypeOk() (*FieldEAdjustmentreminderType, bool)`

GetEAdjustmentreminderTypeOk returns a tuple with the EAdjustmentreminderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAdjustmentreminderType

`func (o *AdjustmentreminderListElement) SetEAdjustmentreminderType(v FieldEAdjustmentreminderType)`

SetEAdjustmentreminderType sets EAdjustmentreminderType field to given value.


### GetSAdjustmentNumber

`func (o *AdjustmentreminderListElement) GetSAdjustmentNumber() string`

GetSAdjustmentNumber returns the SAdjustmentNumber field if non-nil, zero value otherwise.

### GetSAdjustmentNumberOk

`func (o *AdjustmentreminderListElement) GetSAdjustmentNumberOk() (*string, bool)`

GetSAdjustmentNumberOk returns a tuple with the SAdjustmentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAdjustmentNumber

`func (o *AdjustmentreminderListElement) SetSAdjustmentNumber(v string)`

SetSAdjustmentNumber sets SAdjustmentNumber field to given value.

### HasSAdjustmentNumber

`func (o *AdjustmentreminderListElement) HasSAdjustmentNumber() bool`

HasSAdjustmentNumber returns a boolean if a field has been set.

### GetSAdjustmentComment

`func (o *AdjustmentreminderListElement) GetSAdjustmentComment() string`

GetSAdjustmentComment returns the SAdjustmentComment field if non-nil, zero value otherwise.

### GetSAdjustmentCommentOk

`func (o *AdjustmentreminderListElement) GetSAdjustmentCommentOk() (*string, bool)`

GetSAdjustmentCommentOk returns a tuple with the SAdjustmentComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAdjustmentComment

`func (o *AdjustmentreminderListElement) SetSAdjustmentComment(v string)`

SetSAdjustmentComment sets SAdjustmentComment field to given value.

### HasSAdjustmentComment

`func (o *AdjustmentreminderListElement) HasSAdjustmentComment() bool`

HasSAdjustmentComment returns a boolean if a field has been set.

### GetDAdjustmentTotaldebit

`func (o *AdjustmentreminderListElement) GetDAdjustmentTotaldebit() string`

GetDAdjustmentTotaldebit returns the DAdjustmentTotaldebit field if non-nil, zero value otherwise.

### GetDAdjustmentTotaldebitOk

`func (o *AdjustmentreminderListElement) GetDAdjustmentTotaldebitOk() (*string, bool)`

GetDAdjustmentTotaldebitOk returns a tuple with the DAdjustmentTotaldebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDAdjustmentTotaldebit

`func (o *AdjustmentreminderListElement) SetDAdjustmentTotaldebit(v string)`

SetDAdjustmentTotaldebit sets DAdjustmentTotaldebit field to given value.

### HasDAdjustmentTotaldebit

`func (o *AdjustmentreminderListElement) HasDAdjustmentTotaldebit() bool`

HasDAdjustmentTotaldebit returns a boolean if a field has been set.

### GetDAdjustmentTotalcredit

`func (o *AdjustmentreminderListElement) GetDAdjustmentTotalcredit() string`

GetDAdjustmentTotalcredit returns the DAdjustmentTotalcredit field if non-nil, zero value otherwise.

### GetDAdjustmentTotalcreditOk

`func (o *AdjustmentreminderListElement) GetDAdjustmentTotalcreditOk() (*string, bool)`

GetDAdjustmentTotalcreditOk returns a tuple with the DAdjustmentTotalcredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDAdjustmentTotalcredit

`func (o *AdjustmentreminderListElement) SetDAdjustmentTotalcredit(v string)`

SetDAdjustmentTotalcredit sets DAdjustmentTotalcredit field to given value.

### HasDAdjustmentTotalcredit

`func (o *AdjustmentreminderListElement) HasDAdjustmentTotalcredit() bool`

HasDAdjustmentTotalcredit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


