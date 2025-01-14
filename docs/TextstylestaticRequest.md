# TextstylestaticRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiFontID** | **int32** | The unique ID of the Font | 
**BTextstylestaticBold** | **bool** | Whether the Textstylestatic is Bold or not | 
**BTextstylestaticUnderline** | **bool** | Whether the Textstylestatic is Underline or not | 
**BTextstylestaticItalic** | **bool** | Whether the Textstylestatic is Italic or not | 
**BTextstylestaticStrikethrough** | **bool** | Whether the Textstylestatic is Strikethrough or not | 
**ITextstylestaticFontcolor** | **int32** | The int32 representation of the Fontcolor. For example, RGB color #39435B would be 3752795 | 
**ITextstylestaticSize** | **int32** | The Size for the Font of the Textstylestatic | 

## Methods

### NewTextstylestaticRequest

`func NewTextstylestaticRequest(fkiFontID int32, bTextstylestaticBold bool, bTextstylestaticUnderline bool, bTextstylestaticItalic bool, bTextstylestaticStrikethrough bool, iTextstylestaticFontcolor int32, iTextstylestaticSize int32, ) *TextstylestaticRequest`

NewTextstylestaticRequest instantiates a new TextstylestaticRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextstylestaticRequestWithDefaults

`func NewTextstylestaticRequestWithDefaults() *TextstylestaticRequest`

NewTextstylestaticRequestWithDefaults instantiates a new TextstylestaticRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiFontID

`func (o *TextstylestaticRequest) GetFkiFontID() int32`

GetFkiFontID returns the FkiFontID field if non-nil, zero value otherwise.

### GetFkiFontIDOk

`func (o *TextstylestaticRequest) GetFkiFontIDOk() (*int32, bool)`

GetFkiFontIDOk returns a tuple with the FkiFontID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontID

`func (o *TextstylestaticRequest) SetFkiFontID(v int32)`

SetFkiFontID sets FkiFontID field to given value.


### GetBTextstylestaticBold

`func (o *TextstylestaticRequest) GetBTextstylestaticBold() bool`

GetBTextstylestaticBold returns the BTextstylestaticBold field if non-nil, zero value otherwise.

### GetBTextstylestaticBoldOk

`func (o *TextstylestaticRequest) GetBTextstylestaticBoldOk() (*bool, bool)`

GetBTextstylestaticBoldOk returns a tuple with the BTextstylestaticBold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTextstylestaticBold

`func (o *TextstylestaticRequest) SetBTextstylestaticBold(v bool)`

SetBTextstylestaticBold sets BTextstylestaticBold field to given value.


### GetBTextstylestaticUnderline

`func (o *TextstylestaticRequest) GetBTextstylestaticUnderline() bool`

GetBTextstylestaticUnderline returns the BTextstylestaticUnderline field if non-nil, zero value otherwise.

### GetBTextstylestaticUnderlineOk

`func (o *TextstylestaticRequest) GetBTextstylestaticUnderlineOk() (*bool, bool)`

GetBTextstylestaticUnderlineOk returns a tuple with the BTextstylestaticUnderline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTextstylestaticUnderline

`func (o *TextstylestaticRequest) SetBTextstylestaticUnderline(v bool)`

SetBTextstylestaticUnderline sets BTextstylestaticUnderline field to given value.


### GetBTextstylestaticItalic

`func (o *TextstylestaticRequest) GetBTextstylestaticItalic() bool`

GetBTextstylestaticItalic returns the BTextstylestaticItalic field if non-nil, zero value otherwise.

### GetBTextstylestaticItalicOk

`func (o *TextstylestaticRequest) GetBTextstylestaticItalicOk() (*bool, bool)`

GetBTextstylestaticItalicOk returns a tuple with the BTextstylestaticItalic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTextstylestaticItalic

`func (o *TextstylestaticRequest) SetBTextstylestaticItalic(v bool)`

SetBTextstylestaticItalic sets BTextstylestaticItalic field to given value.


### GetBTextstylestaticStrikethrough

`func (o *TextstylestaticRequest) GetBTextstylestaticStrikethrough() bool`

GetBTextstylestaticStrikethrough returns the BTextstylestaticStrikethrough field if non-nil, zero value otherwise.

### GetBTextstylestaticStrikethroughOk

`func (o *TextstylestaticRequest) GetBTextstylestaticStrikethroughOk() (*bool, bool)`

GetBTextstylestaticStrikethroughOk returns a tuple with the BTextstylestaticStrikethrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTextstylestaticStrikethrough

`func (o *TextstylestaticRequest) SetBTextstylestaticStrikethrough(v bool)`

SetBTextstylestaticStrikethrough sets BTextstylestaticStrikethrough field to given value.


### GetITextstylestaticFontcolor

`func (o *TextstylestaticRequest) GetITextstylestaticFontcolor() int32`

GetITextstylestaticFontcolor returns the ITextstylestaticFontcolor field if non-nil, zero value otherwise.

### GetITextstylestaticFontcolorOk

`func (o *TextstylestaticRequest) GetITextstylestaticFontcolorOk() (*int32, bool)`

GetITextstylestaticFontcolorOk returns a tuple with the ITextstylestaticFontcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITextstylestaticFontcolor

`func (o *TextstylestaticRequest) SetITextstylestaticFontcolor(v int32)`

SetITextstylestaticFontcolor sets ITextstylestaticFontcolor field to given value.


### GetITextstylestaticSize

`func (o *TextstylestaticRequest) GetITextstylestaticSize() int32`

GetITextstylestaticSize returns the ITextstylestaticSize field if non-nil, zero value otherwise.

### GetITextstylestaticSizeOk

`func (o *TextstylestaticRequest) GetITextstylestaticSizeOk() (*int32, bool)`

GetITextstylestaticSizeOk returns a tuple with the ITextstylestaticSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITextstylestaticSize

`func (o *TextstylestaticRequest) SetITextstylestaticSize(v int32)`

SetITextstylestaticSize sets ITextstylestaticSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


