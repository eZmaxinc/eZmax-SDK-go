# BrandingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBrandingID** | **int32** | The unique ID of the Branding | 
**FkiEmailID** | Pointer to **int32** | The unique ID of the Email | [optional] 
**ObjBrandingDescription** | [**MultilingualBrandingDescription**](MultilingualBrandingDescription.md) |  | 
**SBrandingDescriptionX** | **string** | The Description of the Branding in the language of the requester | 
**SBrandingName** | Pointer to **string** | The name of the Branding  This value will only be set if you wish to overwrite the default name. If you want to keep the default name, leave this property empty | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**EBrandingLogo** | [**FieldEBrandingLogo**](FieldEBrandingLogo.md) |  | 
**EBrandingLogointerface** | Pointer to [**FieldEBrandingLogointerface**](FieldEBrandingLogointerface.md) |  | [optional] [default to DEFAULT]
**IBrandingColortext** | **int32** | The color of the text. This is a RGB color converted into integer | 
**IBrandingColortextlinkbox** | **int32** | The color of the text in the link box. This is a RGB color converted into integer | 
**IBrandingColortextbutton** | **int32** | The color of the text in the button. This is a RGB color converted into integer | 
**IBrandingColorbackground** | **int32** | The color of the background. This is a RGB color converted into integer | 
**IBrandingColorbackgroundbutton** | **int32** | The color of the background of the button. This is a RGB color converted into integer | 
**IBrandingColorbackgroundsmallbox** | **int32** | The color of the background of the small box. This is a RGB color converted into integer | 
**IBrandingInterfacecolor** | Pointer to **int32** | The color of the interface. This is a RGB color converted into integer | [optional] 
**BBrandingIsactive** | **bool** | Whether the Branding is active or not | 

## Methods

### NewBrandingResponse

`func NewBrandingResponse(pkiBrandingID int32, objBrandingDescription MultilingualBrandingDescription, sBrandingDescriptionX string, eBrandingLogo FieldEBrandingLogo, iBrandingColortext int32, iBrandingColortextlinkbox int32, iBrandingColortextbutton int32, iBrandingColorbackground int32, iBrandingColorbackgroundbutton int32, iBrandingColorbackgroundsmallbox int32, bBrandingIsactive bool, ) *BrandingResponse`

NewBrandingResponse instantiates a new BrandingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingResponseWithDefaults

`func NewBrandingResponseWithDefaults() *BrandingResponse`

NewBrandingResponseWithDefaults instantiates a new BrandingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrandingID

`func (o *BrandingResponse) GetPkiBrandingID() int32`

GetPkiBrandingID returns the PkiBrandingID field if non-nil, zero value otherwise.

### GetPkiBrandingIDOk

`func (o *BrandingResponse) GetPkiBrandingIDOk() (*int32, bool)`

GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrandingID

`func (o *BrandingResponse) SetPkiBrandingID(v int32)`

SetPkiBrandingID sets PkiBrandingID field to given value.


### GetFkiEmailID

`func (o *BrandingResponse) GetFkiEmailID() int32`

GetFkiEmailID returns the FkiEmailID field if non-nil, zero value otherwise.

### GetFkiEmailIDOk

`func (o *BrandingResponse) GetFkiEmailIDOk() (*int32, bool)`

GetFkiEmailIDOk returns a tuple with the FkiEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmailID

`func (o *BrandingResponse) SetFkiEmailID(v int32)`

SetFkiEmailID sets FkiEmailID field to given value.

### HasFkiEmailID

`func (o *BrandingResponse) HasFkiEmailID() bool`

HasFkiEmailID returns a boolean if a field has been set.

### GetObjBrandingDescription

`func (o *BrandingResponse) GetObjBrandingDescription() MultilingualBrandingDescription`

GetObjBrandingDescription returns the ObjBrandingDescription field if non-nil, zero value otherwise.

### GetObjBrandingDescriptionOk

`func (o *BrandingResponse) GetObjBrandingDescriptionOk() (*MultilingualBrandingDescription, bool)`

GetObjBrandingDescriptionOk returns a tuple with the ObjBrandingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjBrandingDescription

`func (o *BrandingResponse) SetObjBrandingDescription(v MultilingualBrandingDescription)`

SetObjBrandingDescription sets ObjBrandingDescription field to given value.


### GetSBrandingDescriptionX

`func (o *BrandingResponse) GetSBrandingDescriptionX() string`

GetSBrandingDescriptionX returns the SBrandingDescriptionX field if non-nil, zero value otherwise.

### GetSBrandingDescriptionXOk

`func (o *BrandingResponse) GetSBrandingDescriptionXOk() (*string, bool)`

GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingDescriptionX

`func (o *BrandingResponse) SetSBrandingDescriptionX(v string)`

SetSBrandingDescriptionX sets SBrandingDescriptionX field to given value.


### GetSBrandingName

`func (o *BrandingResponse) GetSBrandingName() string`

GetSBrandingName returns the SBrandingName field if non-nil, zero value otherwise.

### GetSBrandingNameOk

`func (o *BrandingResponse) GetSBrandingNameOk() (*string, bool)`

GetSBrandingNameOk returns a tuple with the SBrandingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingName

`func (o *BrandingResponse) SetSBrandingName(v string)`

SetSBrandingName sets SBrandingName field to given value.

### HasSBrandingName

`func (o *BrandingResponse) HasSBrandingName() bool`

HasSBrandingName returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *BrandingResponse) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *BrandingResponse) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *BrandingResponse) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *BrandingResponse) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetEBrandingLogo

`func (o *BrandingResponse) GetEBrandingLogo() FieldEBrandingLogo`

GetEBrandingLogo returns the EBrandingLogo field if non-nil, zero value otherwise.

### GetEBrandingLogoOk

`func (o *BrandingResponse) GetEBrandingLogoOk() (*FieldEBrandingLogo, bool)`

GetEBrandingLogoOk returns a tuple with the EBrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingLogo

`func (o *BrandingResponse) SetEBrandingLogo(v FieldEBrandingLogo)`

SetEBrandingLogo sets EBrandingLogo field to given value.


### GetEBrandingLogointerface

`func (o *BrandingResponse) GetEBrandingLogointerface() FieldEBrandingLogointerface`

GetEBrandingLogointerface returns the EBrandingLogointerface field if non-nil, zero value otherwise.

### GetEBrandingLogointerfaceOk

`func (o *BrandingResponse) GetEBrandingLogointerfaceOk() (*FieldEBrandingLogointerface, bool)`

GetEBrandingLogointerfaceOk returns a tuple with the EBrandingLogointerface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingLogointerface

`func (o *BrandingResponse) SetEBrandingLogointerface(v FieldEBrandingLogointerface)`

SetEBrandingLogointerface sets EBrandingLogointerface field to given value.

### HasEBrandingLogointerface

`func (o *BrandingResponse) HasEBrandingLogointerface() bool`

HasEBrandingLogointerface returns a boolean if a field has been set.

### GetIBrandingColortext

`func (o *BrandingResponse) GetIBrandingColortext() int32`

GetIBrandingColortext returns the IBrandingColortext field if non-nil, zero value otherwise.

### GetIBrandingColortextOk

`func (o *BrandingResponse) GetIBrandingColortextOk() (*int32, bool)`

GetIBrandingColortextOk returns a tuple with the IBrandingColortext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortext

`func (o *BrandingResponse) SetIBrandingColortext(v int32)`

SetIBrandingColortext sets IBrandingColortext field to given value.


### GetIBrandingColortextlinkbox

`func (o *BrandingResponse) GetIBrandingColortextlinkbox() int32`

GetIBrandingColortextlinkbox returns the IBrandingColortextlinkbox field if non-nil, zero value otherwise.

### GetIBrandingColortextlinkboxOk

`func (o *BrandingResponse) GetIBrandingColortextlinkboxOk() (*int32, bool)`

GetIBrandingColortextlinkboxOk returns a tuple with the IBrandingColortextlinkbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextlinkbox

`func (o *BrandingResponse) SetIBrandingColortextlinkbox(v int32)`

SetIBrandingColortextlinkbox sets IBrandingColortextlinkbox field to given value.


### GetIBrandingColortextbutton

`func (o *BrandingResponse) GetIBrandingColortextbutton() int32`

GetIBrandingColortextbutton returns the IBrandingColortextbutton field if non-nil, zero value otherwise.

### GetIBrandingColortextbuttonOk

`func (o *BrandingResponse) GetIBrandingColortextbuttonOk() (*int32, bool)`

GetIBrandingColortextbuttonOk returns a tuple with the IBrandingColortextbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextbutton

`func (o *BrandingResponse) SetIBrandingColortextbutton(v int32)`

SetIBrandingColortextbutton sets IBrandingColortextbutton field to given value.


### GetIBrandingColorbackground

`func (o *BrandingResponse) GetIBrandingColorbackground() int32`

GetIBrandingColorbackground returns the IBrandingColorbackground field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundOk

`func (o *BrandingResponse) GetIBrandingColorbackgroundOk() (*int32, bool)`

GetIBrandingColorbackgroundOk returns a tuple with the IBrandingColorbackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackground

`func (o *BrandingResponse) SetIBrandingColorbackground(v int32)`

SetIBrandingColorbackground sets IBrandingColorbackground field to given value.


### GetIBrandingColorbackgroundbutton

`func (o *BrandingResponse) GetIBrandingColorbackgroundbutton() int32`

GetIBrandingColorbackgroundbutton returns the IBrandingColorbackgroundbutton field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundbuttonOk

`func (o *BrandingResponse) GetIBrandingColorbackgroundbuttonOk() (*int32, bool)`

GetIBrandingColorbackgroundbuttonOk returns a tuple with the IBrandingColorbackgroundbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundbutton

`func (o *BrandingResponse) SetIBrandingColorbackgroundbutton(v int32)`

SetIBrandingColorbackgroundbutton sets IBrandingColorbackgroundbutton field to given value.


### GetIBrandingColorbackgroundsmallbox

`func (o *BrandingResponse) GetIBrandingColorbackgroundsmallbox() int32`

GetIBrandingColorbackgroundsmallbox returns the IBrandingColorbackgroundsmallbox field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundsmallboxOk

`func (o *BrandingResponse) GetIBrandingColorbackgroundsmallboxOk() (*int32, bool)`

GetIBrandingColorbackgroundsmallboxOk returns a tuple with the IBrandingColorbackgroundsmallbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundsmallbox

`func (o *BrandingResponse) SetIBrandingColorbackgroundsmallbox(v int32)`

SetIBrandingColorbackgroundsmallbox sets IBrandingColorbackgroundsmallbox field to given value.


### GetIBrandingInterfacecolor

`func (o *BrandingResponse) GetIBrandingInterfacecolor() int32`

GetIBrandingInterfacecolor returns the IBrandingInterfacecolor field if non-nil, zero value otherwise.

### GetIBrandingInterfacecolorOk

`func (o *BrandingResponse) GetIBrandingInterfacecolorOk() (*int32, bool)`

GetIBrandingInterfacecolorOk returns a tuple with the IBrandingInterfacecolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingInterfacecolor

`func (o *BrandingResponse) SetIBrandingInterfacecolor(v int32)`

SetIBrandingInterfacecolor sets IBrandingInterfacecolor field to given value.

### HasIBrandingInterfacecolor

`func (o *BrandingResponse) HasIBrandingInterfacecolor() bool`

HasIBrandingInterfacecolor returns a boolean if a field has been set.

### GetBBrandingIsactive

`func (o *BrandingResponse) GetBBrandingIsactive() bool`

GetBBrandingIsactive returns the BBrandingIsactive field if non-nil, zero value otherwise.

### GetBBrandingIsactiveOk

`func (o *BrandingResponse) GetBBrandingIsactiveOk() (*bool, bool)`

GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrandingIsactive

`func (o *BrandingResponse) SetBBrandingIsactive(v bool)`

SetBBrandingIsactive sets BBrandingIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


