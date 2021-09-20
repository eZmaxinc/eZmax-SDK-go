# WordPositionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IPage** | Pointer to **int32** | The page where the word occurence was found | [optional] 
**IX** | Pointer to **int32** | The X coordinate (Horizontal) where the Word occurence was found.  Coordinate is calculated at 100dpi (dot per inch). | [optional] 
**IY** | Pointer to **int32** | The Y coordinate (Vertical) where the Word occurence was found.  Coordinate is calculated at 100dpi (dot per inch). | [optional] 

## Methods

### NewWordPositionResponse

`func NewWordPositionResponse() *WordPositionResponse`

NewWordPositionResponse instantiates a new WordPositionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordPositionResponseWithDefaults

`func NewWordPositionResponseWithDefaults() *WordPositionResponse`

NewWordPositionResponseWithDefaults instantiates a new WordPositionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIPage

`func (o *WordPositionResponse) GetIPage() int32`

GetIPage returns the IPage field if non-nil, zero value otherwise.

### GetIPageOk

`func (o *WordPositionResponse) GetIPageOk() (*int32, bool)`

GetIPageOk returns a tuple with the IPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPage

`func (o *WordPositionResponse) SetIPage(v int32)`

SetIPage sets IPage field to given value.

### HasIPage

`func (o *WordPositionResponse) HasIPage() bool`

HasIPage returns a boolean if a field has been set.

### GetIX

`func (o *WordPositionResponse) GetIX() int32`

GetIX returns the IX field if non-nil, zero value otherwise.

### GetIXOk

`func (o *WordPositionResponse) GetIXOk() (*int32, bool)`

GetIXOk returns a tuple with the IX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIX

`func (o *WordPositionResponse) SetIX(v int32)`

SetIX sets IX field to given value.

### HasIX

`func (o *WordPositionResponse) HasIX() bool`

HasIX returns a boolean if a field has been set.

### GetIY

`func (o *WordPositionResponse) GetIY() int32`

GetIY returns the IY field if non-nil, zero value otherwise.

### GetIYOk

`func (o *WordPositionResponse) GetIYOk() (*int32, bool)`

GetIYOk returns a tuple with the IY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIY

`func (o *WordPositionResponse) SetIY(v int32)`

SetIY sets IY field to given value.

### HasIY

`func (o *WordPositionResponse) HasIY() bool`

HasIY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


