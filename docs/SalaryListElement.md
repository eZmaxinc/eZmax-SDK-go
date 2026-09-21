# SalaryListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSalaryID** | **int32** | The unique ID of the Salary | 
**FkiPeriodID** | **int32** | The unique ID of the Period | 
**FkiSalaryperiodID** | **int32** | The unique ID of the Salaryperiod | 
**FkiUSerIDCreated** | Pointer to **int32** | The unique ID of the User | [optional] 
**DtSalaryPaymentdate** | **string** | The paymentdate of the Salary | 
**SSalaryperiod** | Pointer to **string** | The period in Salaryperiodgroup | [optional] 
**SPeriodYYYYMM** | Pointer to **string** | The YYYYMM of the Period | [optional] 
**ISalarypreparationCount** | Pointer to **int32** | The count of Salarypreparation. | [optional] 
**DSalarypreparationTotal** | Pointer to **string** | The total of the Salarypreparation | [optional] 
**DtSalaryperiodStart** | Pointer to **string** | The start of the Salaryperiod | [optional] 
**DtSalaryperiodEnd** | Pointer to **string** | The end of the Salaryperiod | [optional] 
**SUserLoginname** | Pointer to **string** | The login name of the User. | [optional] 
**DtCreatedDate** | Pointer to **string** | The date and time at which the object was created | [optional] 
**BSalarypreparationIspaid** | Pointer to **bool** | Whether if it&#39;s an ispaid | [optional] 
**ISalaryperiodgroupYear** | Pointer to **int32** | The year of the Salaryperiodgroup | [optional] 

## Methods

### NewSalaryListElement

`func NewSalaryListElement(pkiSalaryID int32, fkiPeriodID int32, fkiSalaryperiodID int32, dtSalaryPaymentdate string, ) *SalaryListElement`

NewSalaryListElement instantiates a new SalaryListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalaryListElementWithDefaults

`func NewSalaryListElementWithDefaults() *SalaryListElement`

NewSalaryListElementWithDefaults instantiates a new SalaryListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSalaryID

`func (o *SalaryListElement) GetPkiSalaryID() int32`

GetPkiSalaryID returns the PkiSalaryID field if non-nil, zero value otherwise.

### GetPkiSalaryIDOk

`func (o *SalaryListElement) GetPkiSalaryIDOk() (*int32, bool)`

GetPkiSalaryIDOk returns a tuple with the PkiSalaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSalaryID

`func (o *SalaryListElement) SetPkiSalaryID(v int32)`

SetPkiSalaryID sets PkiSalaryID field to given value.


### GetFkiPeriodID

`func (o *SalaryListElement) GetFkiPeriodID() int32`

GetFkiPeriodID returns the FkiPeriodID field if non-nil, zero value otherwise.

### GetFkiPeriodIDOk

`func (o *SalaryListElement) GetFkiPeriodIDOk() (*int32, bool)`

GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPeriodID

`func (o *SalaryListElement) SetFkiPeriodID(v int32)`

SetFkiPeriodID sets FkiPeriodID field to given value.


### GetFkiSalaryperiodID

`func (o *SalaryListElement) GetFkiSalaryperiodID() int32`

GetFkiSalaryperiodID returns the FkiSalaryperiodID field if non-nil, zero value otherwise.

### GetFkiSalaryperiodIDOk

`func (o *SalaryListElement) GetFkiSalaryperiodIDOk() (*int32, bool)`

GetFkiSalaryperiodIDOk returns a tuple with the FkiSalaryperiodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSalaryperiodID

`func (o *SalaryListElement) SetFkiSalaryperiodID(v int32)`

SetFkiSalaryperiodID sets FkiSalaryperiodID field to given value.


### GetFkiUSerIDCreated

`func (o *SalaryListElement) GetFkiUSerIDCreated() int32`

GetFkiUSerIDCreated returns the FkiUSerIDCreated field if non-nil, zero value otherwise.

### GetFkiUSerIDCreatedOk

`func (o *SalaryListElement) GetFkiUSerIDCreatedOk() (*int32, bool)`

GetFkiUSerIDCreatedOk returns a tuple with the FkiUSerIDCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUSerIDCreated

`func (o *SalaryListElement) SetFkiUSerIDCreated(v int32)`

SetFkiUSerIDCreated sets FkiUSerIDCreated field to given value.

### HasFkiUSerIDCreated

`func (o *SalaryListElement) HasFkiUSerIDCreated() bool`

HasFkiUSerIDCreated returns a boolean if a field has been set.

### GetDtSalaryPaymentdate

`func (o *SalaryListElement) GetDtSalaryPaymentdate() string`

GetDtSalaryPaymentdate returns the DtSalaryPaymentdate field if non-nil, zero value otherwise.

### GetDtSalaryPaymentdateOk

`func (o *SalaryListElement) GetDtSalaryPaymentdateOk() (*string, bool)`

GetDtSalaryPaymentdateOk returns a tuple with the DtSalaryPaymentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSalaryPaymentdate

`func (o *SalaryListElement) SetDtSalaryPaymentdate(v string)`

SetDtSalaryPaymentdate sets DtSalaryPaymentdate field to given value.


### GetSSalaryperiod

`func (o *SalaryListElement) GetSSalaryperiod() string`

GetSSalaryperiod returns the SSalaryperiod field if non-nil, zero value otherwise.

### GetSSalaryperiodOk

`func (o *SalaryListElement) GetSSalaryperiodOk() (*string, bool)`

GetSSalaryperiodOk returns a tuple with the SSalaryperiod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSalaryperiod

`func (o *SalaryListElement) SetSSalaryperiod(v string)`

SetSSalaryperiod sets SSalaryperiod field to given value.

### HasSSalaryperiod

`func (o *SalaryListElement) HasSSalaryperiod() bool`

HasSSalaryperiod returns a boolean if a field has been set.

### GetSPeriodYYYYMM

`func (o *SalaryListElement) GetSPeriodYYYYMM() string`

GetSPeriodYYYYMM returns the SPeriodYYYYMM field if non-nil, zero value otherwise.

### GetSPeriodYYYYMMOk

`func (o *SalaryListElement) GetSPeriodYYYYMMOk() (*string, bool)`

GetSPeriodYYYYMMOk returns a tuple with the SPeriodYYYYMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPeriodYYYYMM

`func (o *SalaryListElement) SetSPeriodYYYYMM(v string)`

SetSPeriodYYYYMM sets SPeriodYYYYMM field to given value.

### HasSPeriodYYYYMM

`func (o *SalaryListElement) HasSPeriodYYYYMM() bool`

HasSPeriodYYYYMM returns a boolean if a field has been set.

### GetISalarypreparationCount

`func (o *SalaryListElement) GetISalarypreparationCount() int32`

GetISalarypreparationCount returns the ISalarypreparationCount field if non-nil, zero value otherwise.

### GetISalarypreparationCountOk

`func (o *SalaryListElement) GetISalarypreparationCountOk() (*int32, bool)`

GetISalarypreparationCountOk returns a tuple with the ISalarypreparationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISalarypreparationCount

`func (o *SalaryListElement) SetISalarypreparationCount(v int32)`

SetISalarypreparationCount sets ISalarypreparationCount field to given value.

### HasISalarypreparationCount

`func (o *SalaryListElement) HasISalarypreparationCount() bool`

HasISalarypreparationCount returns a boolean if a field has been set.

### GetDSalarypreparationTotal

`func (o *SalaryListElement) GetDSalarypreparationTotal() string`

GetDSalarypreparationTotal returns the DSalarypreparationTotal field if non-nil, zero value otherwise.

### GetDSalarypreparationTotalOk

`func (o *SalaryListElement) GetDSalarypreparationTotalOk() (*string, bool)`

GetDSalarypreparationTotalOk returns a tuple with the DSalarypreparationTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSalarypreparationTotal

`func (o *SalaryListElement) SetDSalarypreparationTotal(v string)`

SetDSalarypreparationTotal sets DSalarypreparationTotal field to given value.

### HasDSalarypreparationTotal

`func (o *SalaryListElement) HasDSalarypreparationTotal() bool`

HasDSalarypreparationTotal returns a boolean if a field has been set.

### GetDtSalaryperiodStart

`func (o *SalaryListElement) GetDtSalaryperiodStart() string`

GetDtSalaryperiodStart returns the DtSalaryperiodStart field if non-nil, zero value otherwise.

### GetDtSalaryperiodStartOk

`func (o *SalaryListElement) GetDtSalaryperiodStartOk() (*string, bool)`

GetDtSalaryperiodStartOk returns a tuple with the DtSalaryperiodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSalaryperiodStart

`func (o *SalaryListElement) SetDtSalaryperiodStart(v string)`

SetDtSalaryperiodStart sets DtSalaryperiodStart field to given value.

### HasDtSalaryperiodStart

`func (o *SalaryListElement) HasDtSalaryperiodStart() bool`

HasDtSalaryperiodStart returns a boolean if a field has been set.

### GetDtSalaryperiodEnd

`func (o *SalaryListElement) GetDtSalaryperiodEnd() string`

GetDtSalaryperiodEnd returns the DtSalaryperiodEnd field if non-nil, zero value otherwise.

### GetDtSalaryperiodEndOk

`func (o *SalaryListElement) GetDtSalaryperiodEndOk() (*string, bool)`

GetDtSalaryperiodEndOk returns a tuple with the DtSalaryperiodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSalaryperiodEnd

`func (o *SalaryListElement) SetDtSalaryperiodEnd(v string)`

SetDtSalaryperiodEnd sets DtSalaryperiodEnd field to given value.

### HasDtSalaryperiodEnd

`func (o *SalaryListElement) HasDtSalaryperiodEnd() bool`

HasDtSalaryperiodEnd returns a boolean if a field has been set.

### GetSUserLoginname

`func (o *SalaryListElement) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *SalaryListElement) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *SalaryListElement) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.

### HasSUserLoginname

`func (o *SalaryListElement) HasSUserLoginname() bool`

HasSUserLoginname returns a boolean if a field has been set.

### GetDtCreatedDate

`func (o *SalaryListElement) GetDtCreatedDate() string`

GetDtCreatedDate returns the DtCreatedDate field if non-nil, zero value otherwise.

### GetDtCreatedDateOk

`func (o *SalaryListElement) GetDtCreatedDateOk() (*string, bool)`

GetDtCreatedDateOk returns a tuple with the DtCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCreatedDate

`func (o *SalaryListElement) SetDtCreatedDate(v string)`

SetDtCreatedDate sets DtCreatedDate field to given value.

### HasDtCreatedDate

`func (o *SalaryListElement) HasDtCreatedDate() bool`

HasDtCreatedDate returns a boolean if a field has been set.

### GetBSalarypreparationIspaid

`func (o *SalaryListElement) GetBSalarypreparationIspaid() bool`

GetBSalarypreparationIspaid returns the BSalarypreparationIspaid field if non-nil, zero value otherwise.

### GetBSalarypreparationIspaidOk

`func (o *SalaryListElement) GetBSalarypreparationIspaidOk() (*bool, bool)`

GetBSalarypreparationIspaidOk returns a tuple with the BSalarypreparationIspaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSalarypreparationIspaid

`func (o *SalaryListElement) SetBSalarypreparationIspaid(v bool)`

SetBSalarypreparationIspaid sets BSalarypreparationIspaid field to given value.

### HasBSalarypreparationIspaid

`func (o *SalaryListElement) HasBSalarypreparationIspaid() bool`

HasBSalarypreparationIspaid returns a boolean if a field has been set.

### GetISalaryperiodgroupYear

`func (o *SalaryListElement) GetISalaryperiodgroupYear() int32`

GetISalaryperiodgroupYear returns the ISalaryperiodgroupYear field if non-nil, zero value otherwise.

### GetISalaryperiodgroupYearOk

`func (o *SalaryListElement) GetISalaryperiodgroupYearOk() (*int32, bool)`

GetISalaryperiodgroupYearOk returns a tuple with the ISalaryperiodgroupYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISalaryperiodgroupYear

`func (o *SalaryListElement) SetISalaryperiodgroupYear(v int32)`

SetISalaryperiodgroupYear sets ISalaryperiodgroupYear field to given value.

### HasISalaryperiodgroupYear

`func (o *SalaryListElement) HasISalaryperiodgroupYear() bool`

HasISalaryperiodgroupYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


