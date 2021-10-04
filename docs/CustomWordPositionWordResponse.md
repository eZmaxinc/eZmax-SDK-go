# CustomWordPositionWordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SWord** | **string** | The searched word | 
**AObjWordPositionOccurence** | [**[]CustomWordPositionOccurenceResponse**](CustomWordPositionOccurenceResponse.md) | The found occurences for the seached word | 

## Methods

### NewCustomWordPositionWordResponse

`func NewCustomWordPositionWordResponse(sWord string, aObjWordPositionOccurence []CustomWordPositionOccurenceResponse, ) *CustomWordPositionWordResponse`

NewCustomWordPositionWordResponse instantiates a new CustomWordPositionWordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWordPositionWordResponseWithDefaults

`func NewCustomWordPositionWordResponseWithDefaults() *CustomWordPositionWordResponse`

NewCustomWordPositionWordResponseWithDefaults instantiates a new CustomWordPositionWordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSWord

`func (o *CustomWordPositionWordResponse) GetSWord() string`

GetSWord returns the SWord field if non-nil, zero value otherwise.

### GetSWordOk

`func (o *CustomWordPositionWordResponse) GetSWordOk() (*string, bool)`

GetSWordOk returns a tuple with the SWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWord

`func (o *CustomWordPositionWordResponse) SetSWord(v string)`

SetSWord sets SWord field to given value.


### GetAObjWordPositionOccurence

`func (o *CustomWordPositionWordResponse) GetAObjWordPositionOccurence() []CustomWordPositionOccurenceResponse`

GetAObjWordPositionOccurence returns the AObjWordPositionOccurence field if non-nil, zero value otherwise.

### GetAObjWordPositionOccurenceOk

`func (o *CustomWordPositionWordResponse) GetAObjWordPositionOccurenceOk() (*[]CustomWordPositionOccurenceResponse, bool)`

GetAObjWordPositionOccurenceOk returns a tuple with the AObjWordPositionOccurence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjWordPositionOccurence

`func (o *CustomWordPositionWordResponse) SetAObjWordPositionOccurence(v []CustomWordPositionOccurenceResponse)`

SetAObjWordPositionOccurence sets AObjWordPositionOccurence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


