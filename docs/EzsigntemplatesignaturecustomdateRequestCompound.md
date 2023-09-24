# EzsigntemplatesignaturecustomdateRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignaturecustomdateID** | Pointer to **int32** | The unique ID of the Ezsigntemplatesignaturecustomdate | [optional] 
**IEzsigntemplatesignaturecustomdateX** | **int32** | The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsigntemplatesignaturecustomdateY** | **int32** | The Y coordinate (Vertical) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**SEzsigntemplatesignaturecustomdateFormat** | **string** | The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \&quot;Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\&quot; would become \&quot;Signature date: 01/06/2022 08:07\&quot;  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST |  | 

## Methods

### NewEzsigntemplatesignaturecustomdateRequestCompound

`func NewEzsigntemplatesignaturecustomdateRequestCompound(iEzsigntemplatesignaturecustomdateX int32, iEzsigntemplatesignaturecustomdateY int32, sEzsigntemplatesignaturecustomdateFormat string, ) *EzsigntemplatesignaturecustomdateRequestCompound`

NewEzsigntemplatesignaturecustomdateRequestCompound instantiates a new EzsigntemplatesignaturecustomdateRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignaturecustomdateRequestCompoundWithDefaults

`func NewEzsigntemplatesignaturecustomdateRequestCompoundWithDefaults() *EzsigntemplatesignaturecustomdateRequestCompound`

NewEzsigntemplatesignaturecustomdateRequestCompoundWithDefaults instantiates a new EzsigntemplatesignaturecustomdateRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignaturecustomdateID

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetPkiEzsigntemplatesignaturecustomdateID() int32`

GetPkiEzsigntemplatesignaturecustomdateID returns the PkiEzsigntemplatesignaturecustomdateID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignaturecustomdateIDOk

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetPkiEzsigntemplatesignaturecustomdateIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignaturecustomdateIDOk returns a tuple with the PkiEzsigntemplatesignaturecustomdateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignaturecustomdateID

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) SetPkiEzsigntemplatesignaturecustomdateID(v int32)`

SetPkiEzsigntemplatesignaturecustomdateID sets PkiEzsigntemplatesignaturecustomdateID field to given value.

### HasPkiEzsigntemplatesignaturecustomdateID

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) HasPkiEzsigntemplatesignaturecustomdateID() bool`

HasPkiEzsigntemplatesignaturecustomdateID returns a boolean if a field has been set.

### GetIEzsigntemplatesignaturecustomdateX

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetIEzsigntemplatesignaturecustomdateX() int32`

GetIEzsigntemplatesignaturecustomdateX returns the IEzsigntemplatesignaturecustomdateX field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignaturecustomdateXOk

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetIEzsigntemplatesignaturecustomdateXOk() (*int32, bool)`

GetIEzsigntemplatesignaturecustomdateXOk returns a tuple with the IEzsigntemplatesignaturecustomdateX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignaturecustomdateX

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) SetIEzsigntemplatesignaturecustomdateX(v int32)`

SetIEzsigntemplatesignaturecustomdateX sets IEzsigntemplatesignaturecustomdateX field to given value.


### GetIEzsigntemplatesignaturecustomdateY

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetIEzsigntemplatesignaturecustomdateY() int32`

GetIEzsigntemplatesignaturecustomdateY returns the IEzsigntemplatesignaturecustomdateY field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignaturecustomdateYOk

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetIEzsigntemplatesignaturecustomdateYOk() (*int32, bool)`

GetIEzsigntemplatesignaturecustomdateYOk returns a tuple with the IEzsigntemplatesignaturecustomdateY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignaturecustomdateY

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) SetIEzsigntemplatesignaturecustomdateY(v int32)`

SetIEzsigntemplatesignaturecustomdateY sets IEzsigntemplatesignaturecustomdateY field to given value.


### GetSEzsigntemplatesignaturecustomdateFormat

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetSEzsigntemplatesignaturecustomdateFormat() string`

GetSEzsigntemplatesignaturecustomdateFormat returns the SEzsigntemplatesignaturecustomdateFormat field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignaturecustomdateFormatOk

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetSEzsigntemplatesignaturecustomdateFormatOk() (*string, bool)`

GetSEzsigntemplatesignaturecustomdateFormatOk returns a tuple with the SEzsigntemplatesignaturecustomdateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignaturecustomdateFormat

`func (o *EzsigntemplatesignaturecustomdateRequestCompound) SetSEzsigntemplatesignaturecustomdateFormat(v string)`

SetSEzsigntemplatesignaturecustomdateFormat sets SEzsigntemplatesignaturecustomdateFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


