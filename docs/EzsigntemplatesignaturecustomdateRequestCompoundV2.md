# EzsigntemplatesignaturecustomdateRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignaturecustomdateID** | Pointer to **int32** | The unique ID of the Ezsigntemplatesignaturecustomdate | [optional] 
**IEzsigntemplatesignaturecustomdateOffsetx** | **int32** | The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the left of the signature, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsigntemplatesignaturecustomdateOffsety** | **int32** | The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the top of the signature, you would use \&quot;200\&quot; for the Y coordinate. | 
**SEzsigntemplatesignaturecustomdateFormat** | **string** | The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \&quot;Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\&quot; would become \&quot;Signature date: 01/06/2022 08:07\&quot;  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST |  | 

## Methods

### NewEzsigntemplatesignaturecustomdateRequestCompoundV2

`func NewEzsigntemplatesignaturecustomdateRequestCompoundV2(iEzsigntemplatesignaturecustomdateOffsetx int32, iEzsigntemplatesignaturecustomdateOffsety int32, sEzsigntemplatesignaturecustomdateFormat string, ) *EzsigntemplatesignaturecustomdateRequestCompoundV2`

NewEzsigntemplatesignaturecustomdateRequestCompoundV2 instantiates a new EzsigntemplatesignaturecustomdateRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignaturecustomdateRequestCompoundV2WithDefaults

`func NewEzsigntemplatesignaturecustomdateRequestCompoundV2WithDefaults() *EzsigntemplatesignaturecustomdateRequestCompoundV2`

NewEzsigntemplatesignaturecustomdateRequestCompoundV2WithDefaults instantiates a new EzsigntemplatesignaturecustomdateRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignaturecustomdateID

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) GetPkiEzsigntemplatesignaturecustomdateID() int32`

GetPkiEzsigntemplatesignaturecustomdateID returns the PkiEzsigntemplatesignaturecustomdateID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignaturecustomdateIDOk

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) GetPkiEzsigntemplatesignaturecustomdateIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignaturecustomdateIDOk returns a tuple with the PkiEzsigntemplatesignaturecustomdateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignaturecustomdateID

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) SetPkiEzsigntemplatesignaturecustomdateID(v int32)`

SetPkiEzsigntemplatesignaturecustomdateID sets PkiEzsigntemplatesignaturecustomdateID field to given value.

### HasPkiEzsigntemplatesignaturecustomdateID

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) HasPkiEzsigntemplatesignaturecustomdateID() bool`

HasPkiEzsigntemplatesignaturecustomdateID returns a boolean if a field has been set.

### GetIEzsigntemplatesignaturecustomdateOffsetx

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) GetIEzsigntemplatesignaturecustomdateOffsetx() int32`

GetIEzsigntemplatesignaturecustomdateOffsetx returns the IEzsigntemplatesignaturecustomdateOffsetx field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignaturecustomdateOffsetxOk

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) GetIEzsigntemplatesignaturecustomdateOffsetxOk() (*int32, bool)`

GetIEzsigntemplatesignaturecustomdateOffsetxOk returns a tuple with the IEzsigntemplatesignaturecustomdateOffsetx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignaturecustomdateOffsetx

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) SetIEzsigntemplatesignaturecustomdateOffsetx(v int32)`

SetIEzsigntemplatesignaturecustomdateOffsetx sets IEzsigntemplatesignaturecustomdateOffsetx field to given value.


### GetIEzsigntemplatesignaturecustomdateOffsety

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) GetIEzsigntemplatesignaturecustomdateOffsety() int32`

GetIEzsigntemplatesignaturecustomdateOffsety returns the IEzsigntemplatesignaturecustomdateOffsety field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignaturecustomdateOffsetyOk

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) GetIEzsigntemplatesignaturecustomdateOffsetyOk() (*int32, bool)`

GetIEzsigntemplatesignaturecustomdateOffsetyOk returns a tuple with the IEzsigntemplatesignaturecustomdateOffsety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignaturecustomdateOffsety

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) SetIEzsigntemplatesignaturecustomdateOffsety(v int32)`

SetIEzsigntemplatesignaturecustomdateOffsety sets IEzsigntemplatesignaturecustomdateOffsety field to given value.


### GetSEzsigntemplatesignaturecustomdateFormat

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) GetSEzsigntemplatesignaturecustomdateFormat() string`

GetSEzsigntemplatesignaturecustomdateFormat returns the SEzsigntemplatesignaturecustomdateFormat field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignaturecustomdateFormatOk

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) GetSEzsigntemplatesignaturecustomdateFormatOk() (*string, bool)`

GetSEzsigntemplatesignaturecustomdateFormatOk returns a tuple with the SEzsigntemplatesignaturecustomdateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignaturecustomdateFormat

`func (o *EzsigntemplatesignaturecustomdateRequestCompoundV2) SetSEzsigntemplatesignaturecustomdateFormat(v string)`

SetSEzsigntemplatesignaturecustomdateFormat sets SEzsigntemplatesignaturecustomdateFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


