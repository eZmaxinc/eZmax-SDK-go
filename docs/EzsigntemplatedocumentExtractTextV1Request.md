# EzsigntemplatedocumentExtractTextV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IPage** | **int32** | The page where the area is located | 
**ESection** | Pointer to **string** | The section of the page | [optional] 
**IX** | Pointer to **int32** | The X coordinate (Horizontal). Require when eSection &#x3D; &#39;Region&#39; or eSection is not set. | [optional] 
**IY** | Pointer to **int32** | The Y coordinate (Vertical). Require when eSection &#x3D; &#39;Region&#39; or eSection is not set. | [optional] 
**IWidth** | Pointer to **int32** | Area&#39;s width. Require when eSection &#x3D; &#39;Region&#39; or eSection is not set. | [optional] 
**IHeight** | Pointer to **int32** | Area&#39;s height. Require when eSection &#x3D; &#39;Region&#39; or eSection is not set. | [optional] 

## Methods

### NewEzsigntemplatedocumentExtractTextV1Request

`func NewEzsigntemplatedocumentExtractTextV1Request(iPage int32, ) *EzsigntemplatedocumentExtractTextV1Request`

NewEzsigntemplatedocumentExtractTextV1Request instantiates a new EzsigntemplatedocumentExtractTextV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatedocumentExtractTextV1RequestWithDefaults

`func NewEzsigntemplatedocumentExtractTextV1RequestWithDefaults() *EzsigntemplatedocumentExtractTextV1Request`

NewEzsigntemplatedocumentExtractTextV1RequestWithDefaults instantiates a new EzsigntemplatedocumentExtractTextV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIPage

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIPage() int32`

GetIPage returns the IPage field if non-nil, zero value otherwise.

### GetIPageOk

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIPageOk() (*int32, bool)`

GetIPageOk returns a tuple with the IPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPage

`func (o *EzsigntemplatedocumentExtractTextV1Request) SetIPage(v int32)`

SetIPage sets IPage field to given value.


### GetESection

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetESection() string`

GetESection returns the ESection field if non-nil, zero value otherwise.

### GetESectionOk

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetESectionOk() (*string, bool)`

GetESectionOk returns a tuple with the ESection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESection

`func (o *EzsigntemplatedocumentExtractTextV1Request) SetESection(v string)`

SetESection sets ESection field to given value.

### HasESection

`func (o *EzsigntemplatedocumentExtractTextV1Request) HasESection() bool`

HasESection returns a boolean if a field has been set.

### GetIX

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIX() int32`

GetIX returns the IX field if non-nil, zero value otherwise.

### GetIXOk

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIXOk() (*int32, bool)`

GetIXOk returns a tuple with the IX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIX

`func (o *EzsigntemplatedocumentExtractTextV1Request) SetIX(v int32)`

SetIX sets IX field to given value.

### HasIX

`func (o *EzsigntemplatedocumentExtractTextV1Request) HasIX() bool`

HasIX returns a boolean if a field has been set.

### GetIY

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIY() int32`

GetIY returns the IY field if non-nil, zero value otherwise.

### GetIYOk

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIYOk() (*int32, bool)`

GetIYOk returns a tuple with the IY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIY

`func (o *EzsigntemplatedocumentExtractTextV1Request) SetIY(v int32)`

SetIY sets IY field to given value.

### HasIY

`func (o *EzsigntemplatedocumentExtractTextV1Request) HasIY() bool`

HasIY returns a boolean if a field has been set.

### GetIWidth

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIWidth() int32`

GetIWidth returns the IWidth field if non-nil, zero value otherwise.

### GetIWidthOk

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIWidthOk() (*int32, bool)`

GetIWidthOk returns a tuple with the IWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIWidth

`func (o *EzsigntemplatedocumentExtractTextV1Request) SetIWidth(v int32)`

SetIWidth sets IWidth field to given value.

### HasIWidth

`func (o *EzsigntemplatedocumentExtractTextV1Request) HasIWidth() bool`

HasIWidth returns a boolean if a field has been set.

### GetIHeight

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIHeight() int32`

GetIHeight returns the IHeight field if non-nil, zero value otherwise.

### GetIHeightOk

`func (o *EzsigntemplatedocumentExtractTextV1Request) GetIHeightOk() (*int32, bool)`

GetIHeightOk returns a tuple with the IHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIHeight

`func (o *EzsigntemplatedocumentExtractTextV1Request) SetIHeight(v int32)`

SetIHeight sets IHeight field to given value.

### HasIHeight

`func (o *EzsigntemplatedocumentExtractTextV1Request) HasIHeight() bool`

HasIHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


