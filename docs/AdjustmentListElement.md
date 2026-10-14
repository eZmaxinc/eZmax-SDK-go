# AdjustmentListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiAdjustmentID** | **int32** | The unique ID of the Adjustment | 
**FkiPeriodID** | **int32** | The unique ID of the Period | 
**SPeriodYYYYMM** | Pointer to **string** | The YYYYMM of the Period | [optional] 
**FkiUserIDCreated** | Pointer to **int32** | The unique ID of the User | [optional] 
**SUserLoginname** | Pointer to **string** | The login name of the User. | [optional] 
**DtAdjustmentDate** | **string** | The date of the Adjustment | 
**SAdjustmentNumber** | **string** | The number of the Adjustment | 
**SAdjustmentComment** | **string** | The comment of the Adjustment | 
**DAdjustmentTotaldebit** | Pointer to **string** | The total debit of Adjustment | [optional] 
**DAdjustmentTotalcredit** | Pointer to **string** | The total credit of Adjustment | [optional] 

## Methods

### NewAdjustmentListElement

`func NewAdjustmentListElement(pkiAdjustmentID int32, fkiPeriodID int32, dtAdjustmentDate string, sAdjustmentNumber string, sAdjustmentComment string, ) *AdjustmentListElement`

NewAdjustmentListElement instantiates a new AdjustmentListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentListElementWithDefaults

`func NewAdjustmentListElementWithDefaults() *AdjustmentListElement`

NewAdjustmentListElementWithDefaults instantiates a new AdjustmentListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAdjustmentID

`func (o *AdjustmentListElement) GetPkiAdjustmentID() int32`

GetPkiAdjustmentID returns the PkiAdjustmentID field if non-nil, zero value otherwise.

### GetPkiAdjustmentIDOk

`func (o *AdjustmentListElement) GetPkiAdjustmentIDOk() (*int32, bool)`

GetPkiAdjustmentIDOk returns a tuple with the PkiAdjustmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAdjustmentID

`func (o *AdjustmentListElement) SetPkiAdjustmentID(v int32)`

SetPkiAdjustmentID sets PkiAdjustmentID field to given value.


### GetFkiPeriodID

`func (o *AdjustmentListElement) GetFkiPeriodID() int32`

GetFkiPeriodID returns the FkiPeriodID field if non-nil, zero value otherwise.

### GetFkiPeriodIDOk

`func (o *AdjustmentListElement) GetFkiPeriodIDOk() (*int32, bool)`

GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPeriodID

`func (o *AdjustmentListElement) SetFkiPeriodID(v int32)`

SetFkiPeriodID sets FkiPeriodID field to given value.


### GetSPeriodYYYYMM

`func (o *AdjustmentListElement) GetSPeriodYYYYMM() string`

GetSPeriodYYYYMM returns the SPeriodYYYYMM field if non-nil, zero value otherwise.

### GetSPeriodYYYYMMOk

`func (o *AdjustmentListElement) GetSPeriodYYYYMMOk() (*string, bool)`

GetSPeriodYYYYMMOk returns a tuple with the SPeriodYYYYMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPeriodYYYYMM

`func (o *AdjustmentListElement) SetSPeriodYYYYMM(v string)`

SetSPeriodYYYYMM sets SPeriodYYYYMM field to given value.

### HasSPeriodYYYYMM

`func (o *AdjustmentListElement) HasSPeriodYYYYMM() bool`

HasSPeriodYYYYMM returns a boolean if a field has been set.

### GetFkiUserIDCreated

`func (o *AdjustmentListElement) GetFkiUserIDCreated() int32`

GetFkiUserIDCreated returns the FkiUserIDCreated field if non-nil, zero value otherwise.

### GetFkiUserIDCreatedOk

`func (o *AdjustmentListElement) GetFkiUserIDCreatedOk() (*int32, bool)`

GetFkiUserIDCreatedOk returns a tuple with the FkiUserIDCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDCreated

`func (o *AdjustmentListElement) SetFkiUserIDCreated(v int32)`

SetFkiUserIDCreated sets FkiUserIDCreated field to given value.

### HasFkiUserIDCreated

`func (o *AdjustmentListElement) HasFkiUserIDCreated() bool`

HasFkiUserIDCreated returns a boolean if a field has been set.

### GetSUserLoginname

`func (o *AdjustmentListElement) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *AdjustmentListElement) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *AdjustmentListElement) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.

### HasSUserLoginname

`func (o *AdjustmentListElement) HasSUserLoginname() bool`

HasSUserLoginname returns a boolean if a field has been set.

### GetDtAdjustmentDate

`func (o *AdjustmentListElement) GetDtAdjustmentDate() string`

GetDtAdjustmentDate returns the DtAdjustmentDate field if non-nil, zero value otherwise.

### GetDtAdjustmentDateOk

`func (o *AdjustmentListElement) GetDtAdjustmentDateOk() (*string, bool)`

GetDtAdjustmentDateOk returns a tuple with the DtAdjustmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtAdjustmentDate

`func (o *AdjustmentListElement) SetDtAdjustmentDate(v string)`

SetDtAdjustmentDate sets DtAdjustmentDate field to given value.


### GetSAdjustmentNumber

`func (o *AdjustmentListElement) GetSAdjustmentNumber() string`

GetSAdjustmentNumber returns the SAdjustmentNumber field if non-nil, zero value otherwise.

### GetSAdjustmentNumberOk

`func (o *AdjustmentListElement) GetSAdjustmentNumberOk() (*string, bool)`

GetSAdjustmentNumberOk returns a tuple with the SAdjustmentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAdjustmentNumber

`func (o *AdjustmentListElement) SetSAdjustmentNumber(v string)`

SetSAdjustmentNumber sets SAdjustmentNumber field to given value.


### GetSAdjustmentComment

`func (o *AdjustmentListElement) GetSAdjustmentComment() string`

GetSAdjustmentComment returns the SAdjustmentComment field if non-nil, zero value otherwise.

### GetSAdjustmentCommentOk

`func (o *AdjustmentListElement) GetSAdjustmentCommentOk() (*string, bool)`

GetSAdjustmentCommentOk returns a tuple with the SAdjustmentComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAdjustmentComment

`func (o *AdjustmentListElement) SetSAdjustmentComment(v string)`

SetSAdjustmentComment sets SAdjustmentComment field to given value.


### GetDAdjustmentTotaldebit

`func (o *AdjustmentListElement) GetDAdjustmentTotaldebit() string`

GetDAdjustmentTotaldebit returns the DAdjustmentTotaldebit field if non-nil, zero value otherwise.

### GetDAdjustmentTotaldebitOk

`func (o *AdjustmentListElement) GetDAdjustmentTotaldebitOk() (*string, bool)`

GetDAdjustmentTotaldebitOk returns a tuple with the DAdjustmentTotaldebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDAdjustmentTotaldebit

`func (o *AdjustmentListElement) SetDAdjustmentTotaldebit(v string)`

SetDAdjustmentTotaldebit sets DAdjustmentTotaldebit field to given value.

### HasDAdjustmentTotaldebit

`func (o *AdjustmentListElement) HasDAdjustmentTotaldebit() bool`

HasDAdjustmentTotaldebit returns a boolean if a field has been set.

### GetDAdjustmentTotalcredit

`func (o *AdjustmentListElement) GetDAdjustmentTotalcredit() string`

GetDAdjustmentTotalcredit returns the DAdjustmentTotalcredit field if non-nil, zero value otherwise.

### GetDAdjustmentTotalcreditOk

`func (o *AdjustmentListElement) GetDAdjustmentTotalcreditOk() (*string, bool)`

GetDAdjustmentTotalcreditOk returns a tuple with the DAdjustmentTotalcredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDAdjustmentTotalcredit

`func (o *AdjustmentListElement) SetDAdjustmentTotalcredit(v string)`

SetDAdjustmentTotalcredit sets DAdjustmentTotalcredit field to given value.

### HasDAdjustmentTotalcredit

`func (o *AdjustmentListElement) HasDAdjustmentTotalcredit() bool`

HasDAdjustmentTotalcredit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


