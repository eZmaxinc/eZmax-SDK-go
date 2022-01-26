# EzsignsignaturecustomdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IEzsignsignaturecustomdateX** | **int32** | The X coordinate (Horizontal) where to put the custom date block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the custom date block 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsignsignaturecustomdateY** | **int32** | The Y coordinate (Vertical) where to put the custom date block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the custom date block 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**SEzsignsignaturecustomdateFormat** | **string** | The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \&quot;Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\&quot; would become \&quot;Signature date: 01/06/2022 08:07\&quot;  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST |  | 

## Methods

### NewEzsignsignaturecustomdateRequest

`func NewEzsignsignaturecustomdateRequest(iEzsignsignaturecustomdateX int32, iEzsignsignaturecustomdateY int32, sEzsignsignaturecustomdateFormat string, ) *EzsignsignaturecustomdateRequest`

NewEzsignsignaturecustomdateRequest instantiates a new EzsignsignaturecustomdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignaturecustomdateRequestWithDefaults

`func NewEzsignsignaturecustomdateRequestWithDefaults() *EzsignsignaturecustomdateRequest`

NewEzsignsignaturecustomdateRequestWithDefaults instantiates a new EzsignsignaturecustomdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIEzsignsignaturecustomdateX

`func (o *EzsignsignaturecustomdateRequest) GetIEzsignsignaturecustomdateX() int32`

GetIEzsignsignaturecustomdateX returns the IEzsignsignaturecustomdateX field if non-nil, zero value otherwise.

### GetIEzsignsignaturecustomdateXOk

`func (o *EzsignsignaturecustomdateRequest) GetIEzsignsignaturecustomdateXOk() (*int32, bool)`

GetIEzsignsignaturecustomdateXOk returns a tuple with the IEzsignsignaturecustomdateX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignaturecustomdateX

`func (o *EzsignsignaturecustomdateRequest) SetIEzsignsignaturecustomdateX(v int32)`

SetIEzsignsignaturecustomdateX sets IEzsignsignaturecustomdateX field to given value.


### GetIEzsignsignaturecustomdateY

`func (o *EzsignsignaturecustomdateRequest) GetIEzsignsignaturecustomdateY() int32`

GetIEzsignsignaturecustomdateY returns the IEzsignsignaturecustomdateY field if non-nil, zero value otherwise.

### GetIEzsignsignaturecustomdateYOk

`func (o *EzsignsignaturecustomdateRequest) GetIEzsignsignaturecustomdateYOk() (*int32, bool)`

GetIEzsignsignaturecustomdateYOk returns a tuple with the IEzsignsignaturecustomdateY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignaturecustomdateY

`func (o *EzsignsignaturecustomdateRequest) SetIEzsignsignaturecustomdateY(v int32)`

SetIEzsignsignaturecustomdateY sets IEzsignsignaturecustomdateY field to given value.


### GetSEzsignsignaturecustomdateFormat

`func (o *EzsignsignaturecustomdateRequest) GetSEzsignsignaturecustomdateFormat() string`

GetSEzsignsignaturecustomdateFormat returns the SEzsignsignaturecustomdateFormat field if non-nil, zero value otherwise.

### GetSEzsignsignaturecustomdateFormatOk

`func (o *EzsignsignaturecustomdateRequest) GetSEzsignsignaturecustomdateFormatOk() (*string, bool)`

GetSEzsignsignaturecustomdateFormatOk returns a tuple with the SEzsignsignaturecustomdateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignaturecustomdateFormat

`func (o *EzsignsignaturecustomdateRequest) SetSEzsignsignaturecustomdateFormat(v string)`

SetSEzsignsignaturecustomdateFormat sets SEzsignsignaturecustomdateFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


