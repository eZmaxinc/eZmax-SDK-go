# BrandingRequest

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

### NewBrandingRequest

`func NewBrandingRequest(objBrandingDescription MultilingualBrandingDescription, eBrandingLogo FieldEBrandingLogo, iBrandingColortext int32, iBrandingColortextlinkbox int32, iBrandingColortextbutton int32, iBrandingColorbackground int32, iBrandingColorbackgroundbutton int32, iBrandingColorbackgroundsmallbox int32, bBrandingIsactive bool, ) *BrandingRequest`

NewBrandingRequest instantiates a new BrandingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingRequestWithDefaults

`func NewBrandingRequestWithDefaults() *BrandingRequest`

NewBrandingRequestWithDefaults instantiates a new BrandingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrandingID

`func (o *BrandingRequest) GetPkiBrandingID() int32`

GetPkiBrandingID returns the PkiBrandingID field if non-nil, zero value otherwise.

### GetPkiBrandingIDOk

`func (o *BrandingRequest) GetPkiBrandingIDOk() (*int32, bool)`

GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrandingID

`func (o *BrandingRequest) SetPkiBrandingID(v int32)`

SetPkiBrandingID sets PkiBrandingID field to given value.

### HasPkiBrandingID

`func (o *BrandingRequest) HasPkiBrandingID() bool`

HasPkiBrandingID returns a boolean if a field has been set.

### GetObjBrandingDescription

`func (o *BrandingRequest) GetObjBrandingDescription() MultilingualBrandingDescription`

GetObjBrandingDescription returns the ObjBrandingDescription field if non-nil, zero value otherwise.

### GetObjBrandingDescriptionOk

`func (o *BrandingRequest) GetObjBrandingDescriptionOk() (*MultilingualBrandingDescription, bool)`

GetObjBrandingDescriptionOk returns a tuple with the ObjBrandingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjBrandingDescription

`func (o *BrandingRequest) SetObjBrandingDescription(v MultilingualBrandingDescription)`

SetObjBrandingDescription sets ObjBrandingDescription field to given value.


### GetEBrandingLogo

`func (o *BrandingRequest) GetEBrandingLogo() FieldEBrandingLogo`

GetEBrandingLogo returns the EBrandingLogo field if non-nil, zero value otherwise.

### GetEBrandingLogoOk

`func (o *BrandingRequest) GetEBrandingLogoOk() (*FieldEBrandingLogo, bool)`

GetEBrandingLogoOk returns a tuple with the EBrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingLogo

`func (o *BrandingRequest) SetEBrandingLogo(v FieldEBrandingLogo)`

SetEBrandingLogo sets EBrandingLogo field to given value.


### GetSBrandingBase64

`func (o *BrandingRequest) GetSBrandingBase64() string`

GetSBrandingBase64 returns the SBrandingBase64 field if non-nil, zero value otherwise.

### GetSBrandingBase64Ok

`func (o *BrandingRequest) GetSBrandingBase64Ok() (*string, bool)`

GetSBrandingBase64Ok returns a tuple with the SBrandingBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingBase64

`func (o *BrandingRequest) SetSBrandingBase64(v string)`

SetSBrandingBase64 sets SBrandingBase64 field to given value.

### HasSBrandingBase64

`func (o *BrandingRequest) HasSBrandingBase64() bool`

HasSBrandingBase64 returns a boolean if a field has been set.

### GetIBrandingColortext

`func (o *BrandingRequest) GetIBrandingColortext() int32`

GetIBrandingColortext returns the IBrandingColortext field if non-nil, zero value otherwise.

### GetIBrandingColortextOk

`func (o *BrandingRequest) GetIBrandingColortextOk() (*int32, bool)`

GetIBrandingColortextOk returns a tuple with the IBrandingColortext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortext

`func (o *BrandingRequest) SetIBrandingColortext(v int32)`

SetIBrandingColortext sets IBrandingColortext field to given value.


### GetIBrandingColortextlinkbox

`func (o *BrandingRequest) GetIBrandingColortextlinkbox() int32`

GetIBrandingColortextlinkbox returns the IBrandingColortextlinkbox field if non-nil, zero value otherwise.

### GetIBrandingColortextlinkboxOk

`func (o *BrandingRequest) GetIBrandingColortextlinkboxOk() (*int32, bool)`

GetIBrandingColortextlinkboxOk returns a tuple with the IBrandingColortextlinkbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextlinkbox

`func (o *BrandingRequest) SetIBrandingColortextlinkbox(v int32)`

SetIBrandingColortextlinkbox sets IBrandingColortextlinkbox field to given value.


### GetIBrandingColortextbutton

`func (o *BrandingRequest) GetIBrandingColortextbutton() int32`

GetIBrandingColortextbutton returns the IBrandingColortextbutton field if non-nil, zero value otherwise.

### GetIBrandingColortextbuttonOk

`func (o *BrandingRequest) GetIBrandingColortextbuttonOk() (*int32, bool)`

GetIBrandingColortextbuttonOk returns a tuple with the IBrandingColortextbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextbutton

`func (o *BrandingRequest) SetIBrandingColortextbutton(v int32)`

SetIBrandingColortextbutton sets IBrandingColortextbutton field to given value.


### GetIBrandingColorbackground

`func (o *BrandingRequest) GetIBrandingColorbackground() int32`

GetIBrandingColorbackground returns the IBrandingColorbackground field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundOk

`func (o *BrandingRequest) GetIBrandingColorbackgroundOk() (*int32, bool)`

GetIBrandingColorbackgroundOk returns a tuple with the IBrandingColorbackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackground

`func (o *BrandingRequest) SetIBrandingColorbackground(v int32)`

SetIBrandingColorbackground sets IBrandingColorbackground field to given value.


### GetIBrandingColorbackgroundbutton

`func (o *BrandingRequest) GetIBrandingColorbackgroundbutton() int32`

GetIBrandingColorbackgroundbutton returns the IBrandingColorbackgroundbutton field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundbuttonOk

`func (o *BrandingRequest) GetIBrandingColorbackgroundbuttonOk() (*int32, bool)`

GetIBrandingColorbackgroundbuttonOk returns a tuple with the IBrandingColorbackgroundbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundbutton

`func (o *BrandingRequest) SetIBrandingColorbackgroundbutton(v int32)`

SetIBrandingColorbackgroundbutton sets IBrandingColorbackgroundbutton field to given value.


### GetIBrandingColorbackgroundsmallbox

`func (o *BrandingRequest) GetIBrandingColorbackgroundsmallbox() int32`

GetIBrandingColorbackgroundsmallbox returns the IBrandingColorbackgroundsmallbox field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundsmallboxOk

`func (o *BrandingRequest) GetIBrandingColorbackgroundsmallboxOk() (*int32, bool)`

GetIBrandingColorbackgroundsmallboxOk returns a tuple with the IBrandingColorbackgroundsmallbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundsmallbox

`func (o *BrandingRequest) SetIBrandingColorbackgroundsmallbox(v int32)`

SetIBrandingColorbackgroundsmallbox sets IBrandingColorbackgroundsmallbox field to given value.


### GetSBrandingName

`func (o *BrandingRequest) GetSBrandingName() string`

GetSBrandingName returns the SBrandingName field if non-nil, zero value otherwise.

### GetSBrandingNameOk

`func (o *BrandingRequest) GetSBrandingNameOk() (*string, bool)`

GetSBrandingNameOk returns a tuple with the SBrandingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingName

`func (o *BrandingRequest) SetSBrandingName(v string)`

SetSBrandingName sets SBrandingName field to given value.

### HasSBrandingName

`func (o *BrandingRequest) HasSBrandingName() bool`

HasSBrandingName returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *BrandingRequest) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *BrandingRequest) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *BrandingRequest) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *BrandingRequest) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetBBrandingIsactive

`func (o *BrandingRequest) GetBBrandingIsactive() bool`

GetBBrandingIsactive returns the BBrandingIsactive field if non-nil, zero value otherwise.

### GetBBrandingIsactiveOk

`func (o *BrandingRequest) GetBBrandingIsactiveOk() (*bool, bool)`

GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrandingIsactive

`func (o *BrandingRequest) SetBBrandingIsactive(v bool)`

SetBBrandingIsactive sets BBrandingIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


