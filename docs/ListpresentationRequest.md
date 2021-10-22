# ListpresentationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SListpresentationDescription** | **string** | A descriptive for the list presentation | 
**SListpresentationFilter** | **string** | The filter to apply to the request to limit results. | 
**SListpresentationOrderby** | **string** | The order by the user chose | 
**ASColumnName** | **[]string** | An array of column names that the user chose to bee visible | 
**IListpresentationRowMax** | **int32** | The maximum numbers of results to be returned | 
**IListpresentationRowOffset** | **int32** | The starting element from where to start retrieving the results. For example if you started at iRowOffset&#x3D;0 and asked for iRowMax&#x3D;100, to get the next 100 results, you could specify iRowOffset&#x3D;100&amp;iRowMax&#x3D;100, | 

## Methods

### NewListpresentationRequest

`func NewListpresentationRequest(sListpresentationDescription string, sListpresentationFilter string, sListpresentationOrderby string, aSColumnName []string, iListpresentationRowMax int32, iListpresentationRowOffset int32, ) *ListpresentationRequest`

NewListpresentationRequest instantiates a new ListpresentationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListpresentationRequestWithDefaults

`func NewListpresentationRequestWithDefaults() *ListpresentationRequest`

NewListpresentationRequestWithDefaults instantiates a new ListpresentationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSListpresentationDescription

`func (o *ListpresentationRequest) GetSListpresentationDescription() string`

GetSListpresentationDescription returns the SListpresentationDescription field if non-nil, zero value otherwise.

### GetSListpresentationDescriptionOk

`func (o *ListpresentationRequest) GetSListpresentationDescriptionOk() (*string, bool)`

GetSListpresentationDescriptionOk returns a tuple with the SListpresentationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSListpresentationDescription

`func (o *ListpresentationRequest) SetSListpresentationDescription(v string)`

SetSListpresentationDescription sets SListpresentationDescription field to given value.


### GetSListpresentationFilter

`func (o *ListpresentationRequest) GetSListpresentationFilter() string`

GetSListpresentationFilter returns the SListpresentationFilter field if non-nil, zero value otherwise.

### GetSListpresentationFilterOk

`func (o *ListpresentationRequest) GetSListpresentationFilterOk() (*string, bool)`

GetSListpresentationFilterOk returns a tuple with the SListpresentationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSListpresentationFilter

`func (o *ListpresentationRequest) SetSListpresentationFilter(v string)`

SetSListpresentationFilter sets SListpresentationFilter field to given value.


### GetSListpresentationOrderby

`func (o *ListpresentationRequest) GetSListpresentationOrderby() string`

GetSListpresentationOrderby returns the SListpresentationOrderby field if non-nil, zero value otherwise.

### GetSListpresentationOrderbyOk

`func (o *ListpresentationRequest) GetSListpresentationOrderbyOk() (*string, bool)`

GetSListpresentationOrderbyOk returns a tuple with the SListpresentationOrderby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSListpresentationOrderby

`func (o *ListpresentationRequest) SetSListpresentationOrderby(v string)`

SetSListpresentationOrderby sets SListpresentationOrderby field to given value.


### GetASColumnName

`func (o *ListpresentationRequest) GetASColumnName() []string`

GetASColumnName returns the ASColumnName field if non-nil, zero value otherwise.

### GetASColumnNameOk

`func (o *ListpresentationRequest) GetASColumnNameOk() (*[]string, bool)`

GetASColumnNameOk returns a tuple with the ASColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASColumnName

`func (o *ListpresentationRequest) SetASColumnName(v []string)`

SetASColumnName sets ASColumnName field to given value.


### GetIListpresentationRowMax

`func (o *ListpresentationRequest) GetIListpresentationRowMax() int32`

GetIListpresentationRowMax returns the IListpresentationRowMax field if non-nil, zero value otherwise.

### GetIListpresentationRowMaxOk

`func (o *ListpresentationRequest) GetIListpresentationRowMaxOk() (*int32, bool)`

GetIListpresentationRowMaxOk returns a tuple with the IListpresentationRowMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIListpresentationRowMax

`func (o *ListpresentationRequest) SetIListpresentationRowMax(v int32)`

SetIListpresentationRowMax sets IListpresentationRowMax field to given value.


### GetIListpresentationRowOffset

`func (o *ListpresentationRequest) GetIListpresentationRowOffset() int32`

GetIListpresentationRowOffset returns the IListpresentationRowOffset field if non-nil, zero value otherwise.

### GetIListpresentationRowOffsetOk

`func (o *ListpresentationRequest) GetIListpresentationRowOffsetOk() (*int32, bool)`

GetIListpresentationRowOffsetOk returns a tuple with the IListpresentationRowOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIListpresentationRowOffset

`func (o *ListpresentationRequest) SetIListpresentationRowOffset(v int32)`

SetIListpresentationRowOffset sets IListpresentationRowOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


