# TextstylestaticResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiTextstylestaticID** | Pointer to **int32** | The unique ID of the Textstylestatic | [optional] 
**FkiFontID** | **int32** | The unique ID of the Font | 
**SFontName** | **string** | The name of the Font | 
**BTextstylestaticBold** | **bool** | Whether the Textstylestatic is Bold or not | 
**BTextstylestaticUnderline** | **bool** | Whether the Textstylestatic is Underline or not | 
**BTextstylestaticItalic** | **bool** | Whether the Textstylestatic is Italic or not | 
**BTextstylestaticStrikethrough** | **bool** | Whether the Textstylestatic is Strikethrough or not | 
**ITextstylestaticFontcolor** | **int32** | The int32 representation of the Fontcolor. For example, RGB color #39435B would be 3752795 | 
**ITextstylestaticSize** | **int32** | The Size for the Font of the Textstylestatic | 

## Methods

### NewTextstylestaticResponse

`func NewTextstylestaticResponse(fkiFontID int32, sFontName string, bTextstylestaticBold bool, bTextstylestaticUnderline bool, bTextstylestaticItalic bool, bTextstylestaticStrikethrough bool, iTextstylestaticFontcolor int32, iTextstylestaticSize int32, ) *TextstylestaticResponse`

NewTextstylestaticResponse instantiates a new TextstylestaticResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextstylestaticResponseWithDefaults

`func NewTextstylestaticResponseWithDefaults() *TextstylestaticResponse`

NewTextstylestaticResponseWithDefaults instantiates a new TextstylestaticResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiTextstylestaticID

`func (o *TextstylestaticResponse) GetPkiTextstylestaticID() int32`

GetPkiTextstylestaticID returns the PkiTextstylestaticID field if non-nil, zero value otherwise.

### GetPkiTextstylestaticIDOk

`func (o *TextstylestaticResponse) GetPkiTextstylestaticIDOk() (*int32, bool)`

GetPkiTextstylestaticIDOk returns a tuple with the PkiTextstylestaticID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiTextstylestaticID

`func (o *TextstylestaticResponse) SetPkiTextstylestaticID(v int32)`

SetPkiTextstylestaticID sets PkiTextstylestaticID field to given value.

### HasPkiTextstylestaticID

`func (o *TextstylestaticResponse) HasPkiTextstylestaticID() bool`

HasPkiTextstylestaticID returns a boolean if a field has been set.

### GetFkiFontID

`func (o *TextstylestaticResponse) GetFkiFontID() int32`

GetFkiFontID returns the FkiFontID field if non-nil, zero value otherwise.

### GetFkiFontIDOk

`func (o *TextstylestaticResponse) GetFkiFontIDOk() (*int32, bool)`

GetFkiFontIDOk returns a tuple with the FkiFontID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontID

`func (o *TextstylestaticResponse) SetFkiFontID(v int32)`

SetFkiFontID sets FkiFontID field to given value.


### GetSFontName

`func (o *TextstylestaticResponse) GetSFontName() string`

GetSFontName returns the SFontName field if non-nil, zero value otherwise.

### GetSFontNameOk

`func (o *TextstylestaticResponse) GetSFontNameOk() (*string, bool)`

GetSFontNameOk returns a tuple with the SFontName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFontName

`func (o *TextstylestaticResponse) SetSFontName(v string)`

SetSFontName sets SFontName field to given value.


### GetBTextstylestaticBold

`func (o *TextstylestaticResponse) GetBTextstylestaticBold() bool`

GetBTextstylestaticBold returns the BTextstylestaticBold field if non-nil, zero value otherwise.

### GetBTextstylestaticBoldOk

`func (o *TextstylestaticResponse) GetBTextstylestaticBoldOk() (*bool, bool)`

GetBTextstylestaticBoldOk returns a tuple with the BTextstylestaticBold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTextstylestaticBold

`func (o *TextstylestaticResponse) SetBTextstylestaticBold(v bool)`

SetBTextstylestaticBold sets BTextstylestaticBold field to given value.


### GetBTextstylestaticUnderline

`func (o *TextstylestaticResponse) GetBTextstylestaticUnderline() bool`

GetBTextstylestaticUnderline returns the BTextstylestaticUnderline field if non-nil, zero value otherwise.

### GetBTextstylestaticUnderlineOk

`func (o *TextstylestaticResponse) GetBTextstylestaticUnderlineOk() (*bool, bool)`

GetBTextstylestaticUnderlineOk returns a tuple with the BTextstylestaticUnderline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTextstylestaticUnderline

`func (o *TextstylestaticResponse) SetBTextstylestaticUnderline(v bool)`

SetBTextstylestaticUnderline sets BTextstylestaticUnderline field to given value.


### GetBTextstylestaticItalic

`func (o *TextstylestaticResponse) GetBTextstylestaticItalic() bool`

GetBTextstylestaticItalic returns the BTextstylestaticItalic field if non-nil, zero value otherwise.

### GetBTextstylestaticItalicOk

`func (o *TextstylestaticResponse) GetBTextstylestaticItalicOk() (*bool, bool)`

GetBTextstylestaticItalicOk returns a tuple with the BTextstylestaticItalic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTextstylestaticItalic

`func (o *TextstylestaticResponse) SetBTextstylestaticItalic(v bool)`

SetBTextstylestaticItalic sets BTextstylestaticItalic field to given value.


### GetBTextstylestaticStrikethrough

`func (o *TextstylestaticResponse) GetBTextstylestaticStrikethrough() bool`

GetBTextstylestaticStrikethrough returns the BTextstylestaticStrikethrough field if non-nil, zero value otherwise.

### GetBTextstylestaticStrikethroughOk

`func (o *TextstylestaticResponse) GetBTextstylestaticStrikethroughOk() (*bool, bool)`

GetBTextstylestaticStrikethroughOk returns a tuple with the BTextstylestaticStrikethrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTextstylestaticStrikethrough

`func (o *TextstylestaticResponse) SetBTextstylestaticStrikethrough(v bool)`

SetBTextstylestaticStrikethrough sets BTextstylestaticStrikethrough field to given value.


### GetITextstylestaticFontcolor

`func (o *TextstylestaticResponse) GetITextstylestaticFontcolor() int32`

GetITextstylestaticFontcolor returns the ITextstylestaticFontcolor field if non-nil, zero value otherwise.

### GetITextstylestaticFontcolorOk

`func (o *TextstylestaticResponse) GetITextstylestaticFontcolorOk() (*int32, bool)`

GetITextstylestaticFontcolorOk returns a tuple with the ITextstylestaticFontcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITextstylestaticFontcolor

`func (o *TextstylestaticResponse) SetITextstylestaticFontcolor(v int32)`

SetITextstylestaticFontcolor sets ITextstylestaticFontcolor field to given value.


### GetITextstylestaticSize

`func (o *TextstylestaticResponse) GetITextstylestaticSize() int32`

GetITextstylestaticSize returns the ITextstylestaticSize field if non-nil, zero value otherwise.

### GetITextstylestaticSizeOk

`func (o *TextstylestaticResponse) GetITextstylestaticSizeOk() (*int32, bool)`

GetITextstylestaticSizeOk returns a tuple with the ITextstylestaticSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITextstylestaticSize

`func (o *TextstylestaticResponse) SetITextstylestaticSize(v int32)`

SetITextstylestaticSize sets ITextstylestaticSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


