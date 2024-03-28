# EzsignsignaturecustomdateResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignaturecustomdateID** | **int32** | The unique ID of the Ezsignsignaturecustomdate | 
**IEzsignsignaturecustomdateX** | Pointer to **int32** | The X coordinate (Horizontal) where to put the Ezsignsignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignaturecustomdate block 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | [optional] 
**IEzsignsignaturecustomdateY** | Pointer to **int32** | The Y coordinate (Vertical) where to put the Ezsignsignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignaturecustomdate block 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | [optional] 
**IEzsignsignaturecustomdateOffsetx** | Pointer to **int32** | The X coordinate (Horizontal) where to put the Ezsignsignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignaturecustomdate block 2 inches from the left of the signature, you would use \&quot;200\&quot; for the X coordinate. | [optional] 
**IEzsignsignaturecustomdateOffsety** | Pointer to **int32** | The Y coordinate (Vertical) where to put the Ezsignsignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignaturecustomdate block 3 inches from the top of the signature, you would use \&quot;300\&quot; for the Y coordinate. | [optional] 
**SEzsignsignaturecustomdateFormat** | **string** | The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \&quot;Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\&quot; would become \&quot;Signature date: 01/06/2022 08:07\&quot;  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST |  | 

## Methods

### NewEzsignsignaturecustomdateResponseCompound

`func NewEzsignsignaturecustomdateResponseCompound(pkiEzsignsignaturecustomdateID int32, sEzsignsignaturecustomdateFormat string, ) *EzsignsignaturecustomdateResponseCompound`

NewEzsignsignaturecustomdateResponseCompound instantiates a new EzsignsignaturecustomdateResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignaturecustomdateResponseCompoundWithDefaults

`func NewEzsignsignaturecustomdateResponseCompoundWithDefaults() *EzsignsignaturecustomdateResponseCompound`

NewEzsignsignaturecustomdateResponseCompoundWithDefaults instantiates a new EzsignsignaturecustomdateResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignaturecustomdateID

`func (o *EzsignsignaturecustomdateResponseCompound) GetPkiEzsignsignaturecustomdateID() int32`

GetPkiEzsignsignaturecustomdateID returns the PkiEzsignsignaturecustomdateID field if non-nil, zero value otherwise.

### GetPkiEzsignsignaturecustomdateIDOk

`func (o *EzsignsignaturecustomdateResponseCompound) GetPkiEzsignsignaturecustomdateIDOk() (*int32, bool)`

GetPkiEzsignsignaturecustomdateIDOk returns a tuple with the PkiEzsignsignaturecustomdateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignaturecustomdateID

`func (o *EzsignsignaturecustomdateResponseCompound) SetPkiEzsignsignaturecustomdateID(v int32)`

SetPkiEzsignsignaturecustomdateID sets PkiEzsignsignaturecustomdateID field to given value.


### GetIEzsignsignaturecustomdateX

`func (o *EzsignsignaturecustomdateResponseCompound) GetIEzsignsignaturecustomdateX() int32`

GetIEzsignsignaturecustomdateX returns the IEzsignsignaturecustomdateX field if non-nil, zero value otherwise.

### GetIEzsignsignaturecustomdateXOk

`func (o *EzsignsignaturecustomdateResponseCompound) GetIEzsignsignaturecustomdateXOk() (*int32, bool)`

GetIEzsignsignaturecustomdateXOk returns a tuple with the IEzsignsignaturecustomdateX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignaturecustomdateX

`func (o *EzsignsignaturecustomdateResponseCompound) SetIEzsignsignaturecustomdateX(v int32)`

SetIEzsignsignaturecustomdateX sets IEzsignsignaturecustomdateX field to given value.

### HasIEzsignsignaturecustomdateX

`func (o *EzsignsignaturecustomdateResponseCompound) HasIEzsignsignaturecustomdateX() bool`

HasIEzsignsignaturecustomdateX returns a boolean if a field has been set.

### GetIEzsignsignaturecustomdateY

`func (o *EzsignsignaturecustomdateResponseCompound) GetIEzsignsignaturecustomdateY() int32`

GetIEzsignsignaturecustomdateY returns the IEzsignsignaturecustomdateY field if non-nil, zero value otherwise.

### GetIEzsignsignaturecustomdateYOk

`func (o *EzsignsignaturecustomdateResponseCompound) GetIEzsignsignaturecustomdateYOk() (*int32, bool)`

GetIEzsignsignaturecustomdateYOk returns a tuple with the IEzsignsignaturecustomdateY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignaturecustomdateY

`func (o *EzsignsignaturecustomdateResponseCompound) SetIEzsignsignaturecustomdateY(v int32)`

SetIEzsignsignaturecustomdateY sets IEzsignsignaturecustomdateY field to given value.

### HasIEzsignsignaturecustomdateY

`func (o *EzsignsignaturecustomdateResponseCompound) HasIEzsignsignaturecustomdateY() bool`

HasIEzsignsignaturecustomdateY returns a boolean if a field has been set.

### GetIEzsignsignaturecustomdateOffsetx

`func (o *EzsignsignaturecustomdateResponseCompound) GetIEzsignsignaturecustomdateOffsetx() int32`

GetIEzsignsignaturecustomdateOffsetx returns the IEzsignsignaturecustomdateOffsetx field if non-nil, zero value otherwise.

### GetIEzsignsignaturecustomdateOffsetxOk

`func (o *EzsignsignaturecustomdateResponseCompound) GetIEzsignsignaturecustomdateOffsetxOk() (*int32, bool)`

GetIEzsignsignaturecustomdateOffsetxOk returns a tuple with the IEzsignsignaturecustomdateOffsetx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignaturecustomdateOffsetx

`func (o *EzsignsignaturecustomdateResponseCompound) SetIEzsignsignaturecustomdateOffsetx(v int32)`

SetIEzsignsignaturecustomdateOffsetx sets IEzsignsignaturecustomdateOffsetx field to given value.

### HasIEzsignsignaturecustomdateOffsetx

`func (o *EzsignsignaturecustomdateResponseCompound) HasIEzsignsignaturecustomdateOffsetx() bool`

HasIEzsignsignaturecustomdateOffsetx returns a boolean if a field has been set.

### GetIEzsignsignaturecustomdateOffsety

`func (o *EzsignsignaturecustomdateResponseCompound) GetIEzsignsignaturecustomdateOffsety() int32`

GetIEzsignsignaturecustomdateOffsety returns the IEzsignsignaturecustomdateOffsety field if non-nil, zero value otherwise.

### GetIEzsignsignaturecustomdateOffsetyOk

`func (o *EzsignsignaturecustomdateResponseCompound) GetIEzsignsignaturecustomdateOffsetyOk() (*int32, bool)`

GetIEzsignsignaturecustomdateOffsetyOk returns a tuple with the IEzsignsignaturecustomdateOffsety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignaturecustomdateOffsety

`func (o *EzsignsignaturecustomdateResponseCompound) SetIEzsignsignaturecustomdateOffsety(v int32)`

SetIEzsignsignaturecustomdateOffsety sets IEzsignsignaturecustomdateOffsety field to given value.

### HasIEzsignsignaturecustomdateOffsety

`func (o *EzsignsignaturecustomdateResponseCompound) HasIEzsignsignaturecustomdateOffsety() bool`

HasIEzsignsignaturecustomdateOffsety returns a boolean if a field has been set.

### GetSEzsignsignaturecustomdateFormat

`func (o *EzsignsignaturecustomdateResponseCompound) GetSEzsignsignaturecustomdateFormat() string`

GetSEzsignsignaturecustomdateFormat returns the SEzsignsignaturecustomdateFormat field if non-nil, zero value otherwise.

### GetSEzsignsignaturecustomdateFormatOk

`func (o *EzsignsignaturecustomdateResponseCompound) GetSEzsignsignaturecustomdateFormatOk() (*string, bool)`

GetSEzsignsignaturecustomdateFormatOk returns a tuple with the SEzsignsignaturecustomdateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignaturecustomdateFormat

`func (o *EzsignsignaturecustomdateResponseCompound) SetSEzsignsignaturecustomdateFormat(v string)`

SetSEzsignsignaturecustomdateFormat sets SEzsignsignaturecustomdateFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


