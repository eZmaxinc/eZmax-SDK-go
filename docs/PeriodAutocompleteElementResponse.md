# PeriodAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SPeriodYYYYMM** | **string** | The YYYYMM of the Period | 
**PkiPeriodID** | **int32** | The unique ID of the Period | 
**BPeriodIsactive** | **bool** | Whether the Period is active or not | 

## Methods

### NewPeriodAutocompleteElementResponse

`func NewPeriodAutocompleteElementResponse(sPeriodYYYYMM string, pkiPeriodID int32, bPeriodIsactive bool, ) *PeriodAutocompleteElementResponse`

NewPeriodAutocompleteElementResponse instantiates a new PeriodAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodAutocompleteElementResponseWithDefaults

`func NewPeriodAutocompleteElementResponseWithDefaults() *PeriodAutocompleteElementResponse`

NewPeriodAutocompleteElementResponseWithDefaults instantiates a new PeriodAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSPeriodYYYYMM

`func (o *PeriodAutocompleteElementResponse) GetSPeriodYYYYMM() string`

GetSPeriodYYYYMM returns the SPeriodYYYYMM field if non-nil, zero value otherwise.

### GetSPeriodYYYYMMOk

`func (o *PeriodAutocompleteElementResponse) GetSPeriodYYYYMMOk() (*string, bool)`

GetSPeriodYYYYMMOk returns a tuple with the SPeriodYYYYMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPeriodYYYYMM

`func (o *PeriodAutocompleteElementResponse) SetSPeriodYYYYMM(v string)`

SetSPeriodYYYYMM sets SPeriodYYYYMM field to given value.


### GetPkiPeriodID

`func (o *PeriodAutocompleteElementResponse) GetPkiPeriodID() int32`

GetPkiPeriodID returns the PkiPeriodID field if non-nil, zero value otherwise.

### GetPkiPeriodIDOk

`func (o *PeriodAutocompleteElementResponse) GetPkiPeriodIDOk() (*int32, bool)`

GetPkiPeriodIDOk returns a tuple with the PkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPeriodID

`func (o *PeriodAutocompleteElementResponse) SetPkiPeriodID(v int32)`

SetPkiPeriodID sets PkiPeriodID field to given value.


### GetBPeriodIsactive

`func (o *PeriodAutocompleteElementResponse) GetBPeriodIsactive() bool`

GetBPeriodIsactive returns the BPeriodIsactive field if non-nil, zero value otherwise.

### GetBPeriodIsactiveOk

`func (o *PeriodAutocompleteElementResponse) GetBPeriodIsactiveOk() (*bool, bool)`

GetBPeriodIsactiveOk returns a tuple with the BPeriodIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPeriodIsactive

`func (o *PeriodAutocompleteElementResponse) SetBPeriodIsactive(v bool)`

SetBPeriodIsactive sets BPeriodIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


