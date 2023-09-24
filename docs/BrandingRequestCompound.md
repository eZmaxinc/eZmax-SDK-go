# BrandingRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBrandingID** | Pointer to **int32** | The unique ID of the Branding | [optional] 
**ObjBrandingDescription** | [**MultilingualBrandingDescription**](MultilingualBrandingDescription.md) |  | 
**EBrandingLogo** | [**FieldEBrandingLogo**](FieldEBrandingLogo.md) |  | 
**SBrandingBase64** | Pointer to **string** | The Base64 encoded binary content of the branding logo. This need to match image type selected in eBrandingLogo if you supply an image. If you select &#39;Default&#39;, the logo will be deleted and the default one will be used. | [optional] 
**IBrandingColortext** | **int32** | The color of the text. This is a RGB color converted into integer | 
**IBrandingColortextlinkbox** | **int32** | The color of the text in the link box. This is a RGB color converted into integer | 
**IBrandingColortextbutton** | **int32** | The color of the text in the button. This is a RGB color converted into integer | 
**IBrandingColorbackground** | **int32** | The color of the background. This is a RGB color converted into integer | 
**IBrandingColorbackgroundbutton** | **int32** | The color of the background of the button. This is a RGB color converted into integer | 
**IBrandingColorbackgroundsmallbox** | **int32** | The color of the background of the small box. This is a RGB color converted into integer | 
**SBrandingName** | Pointer to **string** | The name of the Branding  This value will only be set if you wish to overwrite the default name. If you want to keep the default name, leave this property empty | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**BBrandingIsactive** | **bool** | Whether the Branding is active or not | 

## Methods

### NewBrandingRequestCompound

`func NewBrandingRequestCompound(objBrandingDescription MultilingualBrandingDescription, eBrandingLogo FieldEBrandingLogo, iBrandingColortext int32, iBrandingColortextlinkbox int32, iBrandingColortextbutton int32, iBrandingColorbackground int32, iBrandingColorbackgroundbutton int32, iBrandingColorbackgroundsmallbox int32, bBrandingIsactive bool, ) *BrandingRequestCompound`

NewBrandingRequestCompound instantiates a new BrandingRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingRequestCompoundWithDefaults

`func NewBrandingRequestCompoundWithDefaults() *BrandingRequestCompound`

NewBrandingRequestCompoundWithDefaults instantiates a new BrandingRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrandingID

`func (o *BrandingRequestCompound) GetPkiBrandingID() int32`

GetPkiBrandingID returns the PkiBrandingID field if non-nil, zero value otherwise.

### GetPkiBrandingIDOk

`func (o *BrandingRequestCompound) GetPkiBrandingIDOk() (*int32, bool)`

GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrandingID

`func (o *BrandingRequestCompound) SetPkiBrandingID(v int32)`

SetPkiBrandingID sets PkiBrandingID field to given value.

### HasPkiBrandingID

`func (o *BrandingRequestCompound) HasPkiBrandingID() bool`

HasPkiBrandingID returns a boolean if a field has been set.

### GetObjBrandingDescription

`func (o *BrandingRequestCompound) GetObjBrandingDescription() MultilingualBrandingDescription`

GetObjBrandingDescription returns the ObjBrandingDescription field if non-nil, zero value otherwise.

### GetObjBrandingDescriptionOk

`func (o *BrandingRequestCompound) GetObjBrandingDescriptionOk() (*MultilingualBrandingDescription, bool)`

GetObjBrandingDescriptionOk returns a tuple with the ObjBrandingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjBrandingDescription

`func (o *BrandingRequestCompound) SetObjBrandingDescription(v MultilingualBrandingDescription)`

SetObjBrandingDescription sets ObjBrandingDescription field to given value.


### GetEBrandingLogo

`func (o *BrandingRequestCompound) GetEBrandingLogo() FieldEBrandingLogo`

GetEBrandingLogo returns the EBrandingLogo field if non-nil, zero value otherwise.

### GetEBrandingLogoOk

`func (o *BrandingRequestCompound) GetEBrandingLogoOk() (*FieldEBrandingLogo, bool)`

GetEBrandingLogoOk returns a tuple with the EBrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingLogo

`func (o *BrandingRequestCompound) SetEBrandingLogo(v FieldEBrandingLogo)`

SetEBrandingLogo sets EBrandingLogo field to given value.


### GetSBrandingBase64

`func (o *BrandingRequestCompound) GetSBrandingBase64() string`

GetSBrandingBase64 returns the SBrandingBase64 field if non-nil, zero value otherwise.

### GetSBrandingBase64Ok

`func (o *BrandingRequestCompound) GetSBrandingBase64Ok() (*string, bool)`

GetSBrandingBase64Ok returns a tuple with the SBrandingBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingBase64

`func (o *BrandingRequestCompound) SetSBrandingBase64(v string)`

SetSBrandingBase64 sets SBrandingBase64 field to given value.

### HasSBrandingBase64

`func (o *BrandingRequestCompound) HasSBrandingBase64() bool`

HasSBrandingBase64 returns a boolean if a field has been set.

### GetIBrandingColortext

`func (o *BrandingRequestCompound) GetIBrandingColortext() int32`

GetIBrandingColortext returns the IBrandingColortext field if non-nil, zero value otherwise.

### GetIBrandingColortextOk

`func (o *BrandingRequestCompound) GetIBrandingColortextOk() (*int32, bool)`

GetIBrandingColortextOk returns a tuple with the IBrandingColortext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortext

`func (o *BrandingRequestCompound) SetIBrandingColortext(v int32)`

SetIBrandingColortext sets IBrandingColortext field to given value.


### GetIBrandingColortextlinkbox

`func (o *BrandingRequestCompound) GetIBrandingColortextlinkbox() int32`

GetIBrandingColortextlinkbox returns the IBrandingColortextlinkbox field if non-nil, zero value otherwise.

### GetIBrandingColortextlinkboxOk

`func (o *BrandingRequestCompound) GetIBrandingColortextlinkboxOk() (*int32, bool)`

GetIBrandingColortextlinkboxOk returns a tuple with the IBrandingColortextlinkbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextlinkbox

`func (o *BrandingRequestCompound) SetIBrandingColortextlinkbox(v int32)`

SetIBrandingColortextlinkbox sets IBrandingColortextlinkbox field to given value.


### GetIBrandingColortextbutton

`func (o *BrandingRequestCompound) GetIBrandingColortextbutton() int32`

GetIBrandingColortextbutton returns the IBrandingColortextbutton field if non-nil, zero value otherwise.

### GetIBrandingColortextbuttonOk

`func (o *BrandingRequestCompound) GetIBrandingColortextbuttonOk() (*int32, bool)`

GetIBrandingColortextbuttonOk returns a tuple with the IBrandingColortextbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextbutton

`func (o *BrandingRequestCompound) SetIBrandingColortextbutton(v int32)`

SetIBrandingColortextbutton sets IBrandingColortextbutton field to given value.


### GetIBrandingColorbackground

`func (o *BrandingRequestCompound) GetIBrandingColorbackground() int32`

GetIBrandingColorbackground returns the IBrandingColorbackground field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundOk

`func (o *BrandingRequestCompound) GetIBrandingColorbackgroundOk() (*int32, bool)`

GetIBrandingColorbackgroundOk returns a tuple with the IBrandingColorbackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackground

`func (o *BrandingRequestCompound) SetIBrandingColorbackground(v int32)`

SetIBrandingColorbackground sets IBrandingColorbackground field to given value.


### GetIBrandingColorbackgroundbutton

`func (o *BrandingRequestCompound) GetIBrandingColorbackgroundbutton() int32`

GetIBrandingColorbackgroundbutton returns the IBrandingColorbackgroundbutton field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundbuttonOk

`func (o *BrandingRequestCompound) GetIBrandingColorbackgroundbuttonOk() (*int32, bool)`

GetIBrandingColorbackgroundbuttonOk returns a tuple with the IBrandingColorbackgroundbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundbutton

`func (o *BrandingRequestCompound) SetIBrandingColorbackgroundbutton(v int32)`

SetIBrandingColorbackgroundbutton sets IBrandingColorbackgroundbutton field to given value.


### GetIBrandingColorbackgroundsmallbox

`func (o *BrandingRequestCompound) GetIBrandingColorbackgroundsmallbox() int32`

GetIBrandingColorbackgroundsmallbox returns the IBrandingColorbackgroundsmallbox field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundsmallboxOk

`func (o *BrandingRequestCompound) GetIBrandingColorbackgroundsmallboxOk() (*int32, bool)`

GetIBrandingColorbackgroundsmallboxOk returns a tuple with the IBrandingColorbackgroundsmallbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundsmallbox

`func (o *BrandingRequestCompound) SetIBrandingColorbackgroundsmallbox(v int32)`

SetIBrandingColorbackgroundsmallbox sets IBrandingColorbackgroundsmallbox field to given value.


### GetSBrandingName

`func (o *BrandingRequestCompound) GetSBrandingName() string`

GetSBrandingName returns the SBrandingName field if non-nil, zero value otherwise.

### GetSBrandingNameOk

`func (o *BrandingRequestCompound) GetSBrandingNameOk() (*string, bool)`

GetSBrandingNameOk returns a tuple with the SBrandingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingName

`func (o *BrandingRequestCompound) SetSBrandingName(v string)`

SetSBrandingName sets SBrandingName field to given value.

### HasSBrandingName

`func (o *BrandingRequestCompound) HasSBrandingName() bool`

HasSBrandingName returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *BrandingRequestCompound) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *BrandingRequestCompound) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *BrandingRequestCompound) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *BrandingRequestCompound) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetBBrandingIsactive

`func (o *BrandingRequestCompound) GetBBrandingIsactive() bool`

GetBBrandingIsactive returns the BBrandingIsactive field if non-nil, zero value otherwise.

### GetBBrandingIsactiveOk

`func (o *BrandingRequestCompound) GetBBrandingIsactiveOk() (*bool, bool)`

GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrandingIsactive

`func (o *BrandingRequestCompound) SetBBrandingIsactive(v bool)`

SetBBrandingIsactive sets BBrandingIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


