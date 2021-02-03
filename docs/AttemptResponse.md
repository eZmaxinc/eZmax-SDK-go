# AttemptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DtAttemptStart** | **string** | Represent a Date Time. The timezone is the one configured in the User&#39;s profile. | 
**SAttemptResult** | **string** | The Success or Failure message of the attempt when we tried to call the URL to deliver the webhook event. | 
**IAttemptDuration** | **int32** | The number of second it took to process the webhook or get an error | 

## Methods

### NewAttemptResponse

`func NewAttemptResponse(dtAttemptStart string, sAttemptResult string, iAttemptDuration int32, ) *AttemptResponse`

NewAttemptResponse instantiates a new AttemptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptResponseWithDefaults

`func NewAttemptResponseWithDefaults() *AttemptResponse`

NewAttemptResponseWithDefaults instantiates a new AttemptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDtAttemptStart

`func (o *AttemptResponse) GetDtAttemptStart() string`

GetDtAttemptStart returns the DtAttemptStart field if non-nil, zero value otherwise.

### GetDtAttemptStartOk

`func (o *AttemptResponse) GetDtAttemptStartOk() (*string, bool)`

GetDtAttemptStartOk returns a tuple with the DtAttemptStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtAttemptStart

`func (o *AttemptResponse) SetDtAttemptStart(v string)`

SetDtAttemptStart sets DtAttemptStart field to given value.


### GetSAttemptResult

`func (o *AttemptResponse) GetSAttemptResult() string`

GetSAttemptResult returns the SAttemptResult field if non-nil, zero value otherwise.

### GetSAttemptResultOk

`func (o *AttemptResponse) GetSAttemptResultOk() (*string, bool)`

GetSAttemptResultOk returns a tuple with the SAttemptResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttemptResult

`func (o *AttemptResponse) SetSAttemptResult(v string)`

SetSAttemptResult sets SAttemptResult field to given value.


### GetIAttemptDuration

`func (o *AttemptResponse) GetIAttemptDuration() int32`

GetIAttemptDuration returns the IAttemptDuration field if non-nil, zero value otherwise.

### GetIAttemptDurationOk

`func (o *AttemptResponse) GetIAttemptDurationOk() (*int32, bool)`

GetIAttemptDurationOk returns a tuple with the IAttemptDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAttemptDuration

`func (o *AttemptResponse) SetIAttemptDuration(v int32)`

SetIAttemptDuration sets IAttemptDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


