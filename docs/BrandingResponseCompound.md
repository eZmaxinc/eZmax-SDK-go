# BrandingResponseCompound

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
**IBrandingColortext** | **int32** | The color of the text. This is a RGB color converted into integer | 
**IBrandingColortextlinkbox** | **int32** | The color of the text in the link box. This is a RGB color converted into integer | 
**IBrandingColortextbutton** | **int32** | The color of the text in the button. This is a RGB color converted into integer | 
**IBrandingColorbackground** | **int32** | The color of the background. This is a RGB color converted into integer | 
**IBrandingColorbackgroundbutton** | **int32** | The color of the background of the button. This is a RGB color converted into integer | 
**IBrandingColorbackgroundsmallbox** | **int32** | The color of the background of the small box. This is a RGB color converted into integer | 
**BBrandingIsactive** | **bool** | Whether the Branding is active or not | 
**SBrandingLogourl** | Pointer to **string** | The url of the picture used as logo in the Branding | [optional] 

## Methods

### NewBrandingResponseCompound

`func NewBrandingResponseCompound(pkiBrandingID int32, objBrandingDescription MultilingualBrandingDescription, sBrandingDescriptionX string, eBrandingLogo FieldEBrandingLogo, iBrandingColortext int32, iBrandingColortextlinkbox int32, iBrandingColortextbutton int32, iBrandingColorbackground int32, iBrandingColorbackgroundbutton int32, iBrandingColorbackgroundsmallbox int32, bBrandingIsactive bool, ) *BrandingResponseCompound`

NewBrandingResponseCompound instantiates a new BrandingResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingResponseCompoundWithDefaults

`func NewBrandingResponseCompoundWithDefaults() *BrandingResponseCompound`

NewBrandingResponseCompoundWithDefaults instantiates a new BrandingResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrandingID

`func (o *BrandingResponseCompound) GetPkiBrandingID() int32`

GetPkiBrandingID returns the PkiBrandingID field if non-nil, zero value otherwise.

### GetPkiBrandingIDOk

`func (o *BrandingResponseCompound) GetPkiBrandingIDOk() (*int32, bool)`

GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrandingID

`func (o *BrandingResponseCompound) SetPkiBrandingID(v int32)`

SetPkiBrandingID sets PkiBrandingID field to given value.


### GetFkiEmailID

`func (o *BrandingResponseCompound) GetFkiEmailID() int32`

GetFkiEmailID returns the FkiEmailID field if non-nil, zero value otherwise.

### GetFkiEmailIDOk

`func (o *BrandingResponseCompound) GetFkiEmailIDOk() (*int32, bool)`

GetFkiEmailIDOk returns a tuple with the FkiEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmailID

`func (o *BrandingResponseCompound) SetFkiEmailID(v int32)`

SetFkiEmailID sets FkiEmailID field to given value.

### HasFkiEmailID

`func (o *BrandingResponseCompound) HasFkiEmailID() bool`

HasFkiEmailID returns a boolean if a field has been set.

### GetObjBrandingDescription

`func (o *BrandingResponseCompound) GetObjBrandingDescription() MultilingualBrandingDescription`

GetObjBrandingDescription returns the ObjBrandingDescription field if non-nil, zero value otherwise.

### GetObjBrandingDescriptionOk

`func (o *BrandingResponseCompound) GetObjBrandingDescriptionOk() (*MultilingualBrandingDescription, bool)`

GetObjBrandingDescriptionOk returns a tuple with the ObjBrandingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjBrandingDescription

`func (o *BrandingResponseCompound) SetObjBrandingDescription(v MultilingualBrandingDescription)`

SetObjBrandingDescription sets ObjBrandingDescription field to given value.


### GetSBrandingDescriptionX

`func (o *BrandingResponseCompound) GetSBrandingDescriptionX() string`

GetSBrandingDescriptionX returns the SBrandingDescriptionX field if non-nil, zero value otherwise.

### GetSBrandingDescriptionXOk

`func (o *BrandingResponseCompound) GetSBrandingDescriptionXOk() (*string, bool)`

GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingDescriptionX

`func (o *BrandingResponseCompound) SetSBrandingDescriptionX(v string)`

SetSBrandingDescriptionX sets SBrandingDescriptionX field to given value.


### GetSBrandingName

`func (o *BrandingResponseCompound) GetSBrandingName() string`

GetSBrandingName returns the SBrandingName field if non-nil, zero value otherwise.

### GetSBrandingNameOk

`func (o *BrandingResponseCompound) GetSBrandingNameOk() (*string, bool)`

GetSBrandingNameOk returns a tuple with the SBrandingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingName

`func (o *BrandingResponseCompound) SetSBrandingName(v string)`

SetSBrandingName sets SBrandingName field to given value.

### HasSBrandingName

`func (o *BrandingResponseCompound) HasSBrandingName() bool`

HasSBrandingName returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *BrandingResponseCompound) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *BrandingResponseCompound) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *BrandingResponseCompound) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *BrandingResponseCompound) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetEBrandingLogo

`func (o *BrandingResponseCompound) GetEBrandingLogo() FieldEBrandingLogo`

GetEBrandingLogo returns the EBrandingLogo field if non-nil, zero value otherwise.

### GetEBrandingLogoOk

`func (o *BrandingResponseCompound) GetEBrandingLogoOk() (*FieldEBrandingLogo, bool)`

GetEBrandingLogoOk returns a tuple with the EBrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingLogo

`func (o *BrandingResponseCompound) SetEBrandingLogo(v FieldEBrandingLogo)`

SetEBrandingLogo sets EBrandingLogo field to given value.


### GetIBrandingColortext

`func (o *BrandingResponseCompound) GetIBrandingColortext() int32`

GetIBrandingColortext returns the IBrandingColortext field if non-nil, zero value otherwise.

### GetIBrandingColortextOk

`func (o *BrandingResponseCompound) GetIBrandingColortextOk() (*int32, bool)`

GetIBrandingColortextOk returns a tuple with the IBrandingColortext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortext

`func (o *BrandingResponseCompound) SetIBrandingColortext(v int32)`

SetIBrandingColortext sets IBrandingColortext field to given value.


### GetIBrandingColortextlinkbox

`func (o *BrandingResponseCompound) GetIBrandingColortextlinkbox() int32`

GetIBrandingColortextlinkbox returns the IBrandingColortextlinkbox field if non-nil, zero value otherwise.

### GetIBrandingColortextlinkboxOk

`func (o *BrandingResponseCompound) GetIBrandingColortextlinkboxOk() (*int32, bool)`

GetIBrandingColortextlinkboxOk returns a tuple with the IBrandingColortextlinkbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextlinkbox

`func (o *BrandingResponseCompound) SetIBrandingColortextlinkbox(v int32)`

SetIBrandingColortextlinkbox sets IBrandingColortextlinkbox field to given value.


### GetIBrandingColortextbutton

`func (o *BrandingResponseCompound) GetIBrandingColortextbutton() int32`

GetIBrandingColortextbutton returns the IBrandingColortextbutton field if non-nil, zero value otherwise.

### GetIBrandingColortextbuttonOk

`func (o *BrandingResponseCompound) GetIBrandingColortextbuttonOk() (*int32, bool)`

GetIBrandingColortextbuttonOk returns a tuple with the IBrandingColortextbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColortextbutton

`func (o *BrandingResponseCompound) SetIBrandingColortextbutton(v int32)`

SetIBrandingColortextbutton sets IBrandingColortextbutton field to given value.


### GetIBrandingColorbackground

`func (o *BrandingResponseCompound) GetIBrandingColorbackground() int32`

GetIBrandingColorbackground returns the IBrandingColorbackground field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundOk

`func (o *BrandingResponseCompound) GetIBrandingColorbackgroundOk() (*int32, bool)`

GetIBrandingColorbackgroundOk returns a tuple with the IBrandingColorbackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackground

`func (o *BrandingResponseCompound) SetIBrandingColorbackground(v int32)`

SetIBrandingColorbackground sets IBrandingColorbackground field to given value.


### GetIBrandingColorbackgroundbutton

`func (o *BrandingResponseCompound) GetIBrandingColorbackgroundbutton() int32`

GetIBrandingColorbackgroundbutton returns the IBrandingColorbackgroundbutton field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundbuttonOk

`func (o *BrandingResponseCompound) GetIBrandingColorbackgroundbuttonOk() (*int32, bool)`

GetIBrandingColorbackgroundbuttonOk returns a tuple with the IBrandingColorbackgroundbutton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundbutton

`func (o *BrandingResponseCompound) SetIBrandingColorbackgroundbutton(v int32)`

SetIBrandingColorbackgroundbutton sets IBrandingColorbackgroundbutton field to given value.


### GetIBrandingColorbackgroundsmallbox

`func (o *BrandingResponseCompound) GetIBrandingColorbackgroundsmallbox() int32`

GetIBrandingColorbackgroundsmallbox returns the IBrandingColorbackgroundsmallbox field if non-nil, zero value otherwise.

### GetIBrandingColorbackgroundsmallboxOk

`func (o *BrandingResponseCompound) GetIBrandingColorbackgroundsmallboxOk() (*int32, bool)`

GetIBrandingColorbackgroundsmallboxOk returns a tuple with the IBrandingColorbackgroundsmallbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColorbackgroundsmallbox

`func (o *BrandingResponseCompound) SetIBrandingColorbackgroundsmallbox(v int32)`

SetIBrandingColorbackgroundsmallbox sets IBrandingColorbackgroundsmallbox field to given value.


### GetBBrandingIsactive

`func (o *BrandingResponseCompound) GetBBrandingIsactive() bool`

GetBBrandingIsactive returns the BBrandingIsactive field if non-nil, zero value otherwise.

### GetBBrandingIsactiveOk

`func (o *BrandingResponseCompound) GetBBrandingIsactiveOk() (*bool, bool)`

GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrandingIsactive

`func (o *BrandingResponseCompound) SetBBrandingIsactive(v bool)`

SetBBrandingIsactive sets BBrandingIsactive field to given value.


### GetSBrandingLogourl

`func (o *BrandingResponseCompound) GetSBrandingLogourl() string`

GetSBrandingLogourl returns the SBrandingLogourl field if non-nil, zero value otherwise.

### GetSBrandingLogourlOk

`func (o *BrandingResponseCompound) GetSBrandingLogourlOk() (*string, bool)`

GetSBrandingLogourlOk returns a tuple with the SBrandingLogourl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingLogourl

`func (o *BrandingResponseCompound) SetSBrandingLogourl(v string)`

SetSBrandingLogourl sets SBrandingLogourl field to given value.

### HasSBrandingLogourl

`func (o *BrandingResponseCompound) HasSBrandingLogourl() bool`

HasSBrandingLogourl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


