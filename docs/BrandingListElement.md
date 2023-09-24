# BrandingListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBrandingID** | **int32** | The unique ID of the Branding | 
**SBrandingDescriptionX** | **string** | The Description of the Branding in the language of the requester | 
**IBrandingColortext** | **int32** | The color of the text. This is a RGB color converted into integer | 
**IBrandingColortextlinkbox** | **int32** | The color of the text in the link box. This is a RGB color converted into integer | 
**IBrandingColortextbutton** | **int32** | The color of the text in the button. This is a RGB color converted into integer | 
**IBrandingColorbackground** | **int32** | The color of the background. This is a RGB color converted into integer | 
**IBrandingColorbackgroundbutton** | **int32** | The color of the background of the button. This is a RGB color converted into integer | 
**IBrandingColorbackgroundsmallbox** | **int32** | The color of the background of the small box. This is a RGB color converted into integer | 
**BBrandingIsactive** | **bool** | Whether the Branding is active or not | 

## Methods

### NewBrandingListElement

`func NewBrandingListElement(pkiBrandingID int32, sBrandingDescriptionX string, iBrandingColortext int32, iBrandingColortextlinkbox int32, iBrandingColortextbutton int32, iBrandingColorbackground int32, iBrandingColorbackgroundbutton int32, iBrandingColorbackgroundsmallbox int32, bBrandingIsactive bool, ) *BrandingListElement`

NewBrandingListElement instantiates a new BrandingListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingListElementWithDefaults

`func NewBrandingListElementWithDefaults() *BrandingListElement`

NewBrandingListElementWithDefaults instantiates a new BrandingListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrandingID

`func (o *BrandingListElement) GetPkiBrandingID() int32`

GetPkiBrandingID returns the PkiBrandingID field if non-nil, zero value otherwise.

### GetPkiBrandingIDOk

`func (o *BrandingListElement) GetPkiBrandingIDOk() (*int32, bool)`

GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrandingID

`func (o *BrandingListElement) SetPkiBrandingID(v int32)`

SetPkiBrandingID sets PkiBrandingID field to given value.


### GetSBrandingDescriptionX

`func (o *BrandingListElement) GetSBrandingDescriptionX() string`

GetSBrandingDescriptionX returns the SBrandingDescriptionX field if non-nil, zero value otherwise.

### GetSBrandingDescriptionXOk

`func (o *BrandingListElement) GetSBrandingDescriptionXOk() (*string, bool)`

GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingDescriptionX

`func (o *BrandingListElement) SetSBrandingDescriptionX(v string)`

SetSBrandingDescriptionX sets SBrandingDescriptionX field to given value.


### GetIBrandingColortext

`func (o *BrandingListElement) GetIBrandingColortext() int32`

GetIBrandingColortext returns the IBrandingColortext field if non-nil, zero value otherwise.

### GetIBrandingColortextOk

`func (o *BrandingListElement) GetIBrandingColortextOk() (*int32, bool)`

GetIBrandingColortextOk returns a tuple with the IBrandingColortext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortext

`func (o *BrandingListElement) SetIBrandingColortext(v int32)`

SetIBrandingColortext sets IBrandingColortext field to given value.


### GetIBrandingColortextlinkbox

`func (o *BrandingListElement) GetIBrandingColortextlinkbox() int32`

GetIBrandingColortextlinkbox returns the IBrandingColortextlinkbox field if non-nil, zero value otherwise.

### GetIBrandingColortextlinkboxOk

`func (o *BrandingListElement) GetIBrandingColortextlinkboxOk() (*int32, bool)`

GetIBrandingColortextlinkboxOk returns a tuple with the IBrandingColortextlinkbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextlinkbox

`func (o *BrandingListElement) SetIBrandingColortextlinkbox(v int32)`

SetIBrandingColortextlinkbox sets IBrandingColortextlinkbox field to given value.


### GetIBrandingColortextbutton

`func (o *BrandingListElement) GetIBrandingColortextbutton() int32`

GetIBrandingColortextbutton returns the IBrandingColortextbutton field if non-nil, zero value otherwise.

### GetIBrandingColortextbuttonOk

`func (o *BrandingListElement) GetIBrandingColortextbuttonOk() (*int32, bool)`

GetIBrandingColortextbuttonOk returns a tuple with the IBrandingColortextbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextbutton

`func (o *BrandingListElement) SetIBrandingColortextbutton(v int32)`

SetIBrandingColortextbutton sets IBrandingColortextbutton field to given value.


### GetIBrandingColorbackground

`func (o *BrandingListElement) GetIBrandingColorbackground() int32`

GetIBrandingColorbackground returns the IBrandingColorbackground field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundOk

`func (o *BrandingListElement) GetIBrandingColorbackgroundOk() (*int32, bool)`

GetIBrandingColorbackgroundOk returns a tuple with the IBrandingColorbackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackground

`func (o *BrandingListElement) SetIBrandingColorbackground(v int32)`

SetIBrandingColorbackground sets IBrandingColorbackground field to given value.


### GetIBrandingColorbackgroundbutton

`func (o *BrandingListElement) GetIBrandingColorbackgroundbutton() int32`

GetIBrandingColorbackgroundbutton returns the IBrandingColorbackgroundbutton field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundbuttonOk

`func (o *BrandingListElement) GetIBrandingColorbackgroundbuttonOk() (*int32, bool)`

GetIBrandingColorbackgroundbuttonOk returns a tuple with the IBrandingColorbackgroundbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundbutton

`func (o *BrandingListElement) SetIBrandingColorbackgroundbutton(v int32)`

SetIBrandingColorbackgroundbutton sets IBrandingColorbackgroundbutton field to given value.


### GetIBrandingColorbackgroundsmallbox

`func (o *BrandingListElement) GetIBrandingColorbackgroundsmallbox() int32`

GetIBrandingColorbackgroundsmallbox returns the IBrandingColorbackgroundsmallbox field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundsmallboxOk

`func (o *BrandingListElement) GetIBrandingColorbackgroundsmallboxOk() (*int32, bool)`

GetIBrandingColorbackgroundsmallboxOk returns a tuple with the IBrandingColorbackgroundsmallbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundsmallbox

`func (o *BrandingListElement) SetIBrandingColorbackgroundsmallbox(v int32)`

SetIBrandingColorbackgroundsmallbox sets IBrandingColorbackgroundsmallbox field to given value.


### GetBBrandingIsactive

`func (o *BrandingListElement) GetBBrandingIsactive() bool`

GetBBrandingIsactive returns the BBrandingIsactive field if non-nil, zero value otherwise.

### GetBBrandingIsactiveOk

`func (o *BrandingListElement) GetBBrandingIsactiveOk() (*bool, bool)`

GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrandingIsactive

`func (o *BrandingListElement) SetBBrandingIsactive(v bool)`

SetBBrandingIsactive sets BBrandingIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


