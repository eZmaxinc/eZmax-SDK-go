# EzsignpageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignpageID** | **int32** | The unique ID of the Ezsignpage | 
**IEzsignpageWidthimage** | **int32** | The Width of the page&#39;s image in pixels calculated at 100 DPI | 
**IEzsignpageHeightimage** | **int32** | The Height of the page&#39;s image in pixels calculated at 100 DPI | 
**IEzsignpageWidthpdf** | **int32** | The Width of the page in points calculated at 72 DPI | 
**IEzsignpageHeightpdf** | **int32** | The Height of the page in points calculated at 72 DPI | 
**IEzsignpagePagenumber** | **int32** | The page number in the Ezsigndocument | 
**SComputedImageurl** | **string** | The Url to the Ezsignpage&#39;s rasterized image.  Url will expire after 5 minutes. | 

## Methods

### NewEzsignpageResponse

`func NewEzsignpageResponse(pkiEzsignpageID int32, iEzsignpageWidthimage int32, iEzsignpageHeightimage int32, iEzsignpageWidthpdf int32, iEzsignpageHeightpdf int32, iEzsignpagePagenumber int32, sComputedImageurl string, ) *EzsignpageResponse`

NewEzsignpageResponse instantiates a new EzsignpageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignpageResponseWithDefaults

`func NewEzsignpageResponseWithDefaults() *EzsignpageResponse`

NewEzsignpageResponseWithDefaults instantiates a new EzsignpageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignpageID

`func (o *EzsignpageResponse) GetPkiEzsignpageID() int32`

GetPkiEzsignpageID returns the PkiEzsignpageID field if non-nil, zero value otherwise.

### GetPkiEzsignpageIDOk

`func (o *EzsignpageResponse) GetPkiEzsignpageIDOk() (*int32, bool)`

GetPkiEzsignpageIDOk returns a tuple with the PkiEzsignpageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignpageID

`func (o *EzsignpageResponse) SetPkiEzsignpageID(v int32)`

SetPkiEzsignpageID sets PkiEzsignpageID field to given value.


### GetIEzsignpageWidthimage

`func (o *EzsignpageResponse) GetIEzsignpageWidthimage() int32`

GetIEzsignpageWidthimage returns the IEzsignpageWidthimage field if non-nil, zero value otherwise.

### GetIEzsignpageWidthimageOk

`func (o *EzsignpageResponse) GetIEzsignpageWidthimageOk() (*int32, bool)`

GetIEzsignpageWidthimageOk returns a tuple with the IEzsignpageWidthimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpageWidthimage

`func (o *EzsignpageResponse) SetIEzsignpageWidthimage(v int32)`

SetIEzsignpageWidthimage sets IEzsignpageWidthimage field to given value.


### GetIEzsignpageHeightimage

`func (o *EzsignpageResponse) GetIEzsignpageHeightimage() int32`

GetIEzsignpageHeightimage returns the IEzsignpageHeightimage field if non-nil, zero value otherwise.

### GetIEzsignpageHeightimageOk

`func (o *EzsignpageResponse) GetIEzsignpageHeightimageOk() (*int32, bool)`

GetIEzsignpageHeightimageOk returns a tuple with the IEzsignpageHeightimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpageHeightimage

`func (o *EzsignpageResponse) SetIEzsignpageHeightimage(v int32)`

SetIEzsignpageHeightimage sets IEzsignpageHeightimage field to given value.


### GetIEzsignpageWidthpdf

`func (o *EzsignpageResponse) GetIEzsignpageWidthpdf() int32`

GetIEzsignpageWidthpdf returns the IEzsignpageWidthpdf field if non-nil, zero value otherwise.

### GetIEzsignpageWidthpdfOk

`func (o *EzsignpageResponse) GetIEzsignpageWidthpdfOk() (*int32, bool)`

GetIEzsignpageWidthpdfOk returns a tuple with the IEzsignpageWidthpdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpageWidthpdf

`func (o *EzsignpageResponse) SetIEzsignpageWidthpdf(v int32)`

SetIEzsignpageWidthpdf sets IEzsignpageWidthpdf field to given value.


### GetIEzsignpageHeightpdf

`func (o *EzsignpageResponse) GetIEzsignpageHeightpdf() int32`

GetIEzsignpageHeightpdf returns the IEzsignpageHeightpdf field if non-nil, zero value otherwise.

### GetIEzsignpageHeightpdfOk

`func (o *EzsignpageResponse) GetIEzsignpageHeightpdfOk() (*int32, bool)`

GetIEzsignpageHeightpdfOk returns a tuple with the IEzsignpageHeightpdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpageHeightpdf

`func (o *EzsignpageResponse) SetIEzsignpageHeightpdf(v int32)`

SetIEzsignpageHeightpdf sets IEzsignpageHeightpdf field to given value.


### GetIEzsignpagePagenumber

`func (o *EzsignpageResponse) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *EzsignpageResponse) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *EzsignpageResponse) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.


### GetSComputedImageurl

`func (o *EzsignpageResponse) GetSComputedImageurl() string`

GetSComputedImageurl returns the SComputedImageurl field if non-nil, zero value otherwise.

### GetSComputedImageurlOk

`func (o *EzsignpageResponse) GetSComputedImageurlOk() (*string, bool)`

GetSComputedImageurlOk returns a tuple with the SComputedImageurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSComputedImageurl

`func (o *EzsignpageResponse) SetSComputedImageurl(v string)`

SetSComputedImageurl sets SComputedImageurl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


