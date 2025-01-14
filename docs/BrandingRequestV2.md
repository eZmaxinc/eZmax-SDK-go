# BrandingRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBrandingID** | Pointer to **int32** | The unique ID of the Branding | [optional] 
**ObjBrandingDescription** | [**MultilingualBrandingDescription**](MultilingualBrandingDescription.md) |  | 
**EBrandingLogo** | [**FieldEBrandingLogo**](FieldEBrandingLogo.md) |  | 
**EBrandingAlignlogo** | Pointer to [**FieldEBrandingAlignlogo**](FieldEBrandingAlignlogo.md) |  | [optional] 
**SBrandingBase64** | Pointer to **string** | The Base64 encoded binary content of the branding logo. This need to match image type selected in eBrandingLogo if you supply an image. If you select &#39;Default&#39;, the logo will be deleted and the default one will be used. | [optional] 
**IBrandingColor** | **int32** | The primary color. This is a RGB color converted into integer | 
**SBrandingName** | Pointer to **string** | The name of the Branding  This value will only be set if you wish to overwrite the default name. If you want to keep the default name, leave this property empty | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**BBrandingIsactive** | **bool** | Whether the Branding is active or not | 

## Methods

### NewBrandingRequestV2

`func NewBrandingRequestV2(objBrandingDescription MultilingualBrandingDescription, eBrandingLogo FieldEBrandingLogo, iBrandingColor int32, bBrandingIsactive bool, ) *BrandingRequestV2`

NewBrandingRequestV2 instantiates a new BrandingRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingRequestV2WithDefaults

`func NewBrandingRequestV2WithDefaults() *BrandingRequestV2`

NewBrandingRequestV2WithDefaults instantiates a new BrandingRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrandingID

`func (o *BrandingRequestV2) GetPkiBrandingID() int32`

GetPkiBrandingID returns the PkiBrandingID field if non-nil, zero value otherwise.

### GetPkiBrandingIDOk

`func (o *BrandingRequestV2) GetPkiBrandingIDOk() (*int32, bool)`

GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrandingID

`func (o *BrandingRequestV2) SetPkiBrandingID(v int32)`

SetPkiBrandingID sets PkiBrandingID field to given value.

### HasPkiBrandingID

`func (o *BrandingRequestV2) HasPkiBrandingID() bool`

HasPkiBrandingID returns a boolean if a field has been set.

### GetObjBrandingDescription

`func (o *BrandingRequestV2) GetObjBrandingDescription() MultilingualBrandingDescription`

GetObjBrandingDescription returns the ObjBrandingDescription field if non-nil, zero value otherwise.

### GetObjBrandingDescriptionOk

`func (o *BrandingRequestV2) GetObjBrandingDescriptionOk() (*MultilingualBrandingDescription, bool)`

GetObjBrandingDescriptionOk returns a tuple with the ObjBrandingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjBrandingDescription

`func (o *BrandingRequestV2) SetObjBrandingDescription(v MultilingualBrandingDescription)`

SetObjBrandingDescription sets ObjBrandingDescription field to given value.


### GetEBrandingLogo

`func (o *BrandingRequestV2) GetEBrandingLogo() FieldEBrandingLogo`

GetEBrandingLogo returns the EBrandingLogo field if non-nil, zero value otherwise.

### GetEBrandingLogoOk

`func (o *BrandingRequestV2) GetEBrandingLogoOk() (*FieldEBrandingLogo, bool)`

GetEBrandingLogoOk returns a tuple with the EBrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingLogo

`func (o *BrandingRequestV2) SetEBrandingLogo(v FieldEBrandingLogo)`

SetEBrandingLogo sets EBrandingLogo field to given value.


### GetEBrandingAlignlogo

`func (o *BrandingRequestV2) GetEBrandingAlignlogo() FieldEBrandingAlignlogo`

GetEBrandingAlignlogo returns the EBrandingAlignlogo field if non-nil, zero value otherwise.

### GetEBrandingAlignlogoOk

`func (o *BrandingRequestV2) GetEBrandingAlignlogoOk() (*FieldEBrandingAlignlogo, bool)`

GetEBrandingAlignlogoOk returns a tuple with the EBrandingAlignlogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingAlignlogo

`func (o *BrandingRequestV2) SetEBrandingAlignlogo(v FieldEBrandingAlignlogo)`

SetEBrandingAlignlogo sets EBrandingAlignlogo field to given value.

### HasEBrandingAlignlogo

`func (o *BrandingRequestV2) HasEBrandingAlignlogo() bool`

HasEBrandingAlignlogo returns a boolean if a field has been set.

### GetSBrandingBase64

`func (o *BrandingRequestV2) GetSBrandingBase64() string`

GetSBrandingBase64 returns the SBrandingBase64 field if non-nil, zero value otherwise.

### GetSBrandingBase64Ok

`func (o *BrandingRequestV2) GetSBrandingBase64Ok() (*string, bool)`

GetSBrandingBase64Ok returns a tuple with the SBrandingBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingBase64

`func (o *BrandingRequestV2) SetSBrandingBase64(v string)`

SetSBrandingBase64 sets SBrandingBase64 field to given value.

### HasSBrandingBase64

`func (o *BrandingRequestV2) HasSBrandingBase64() bool`

HasSBrandingBase64 returns a boolean if a field has been set.

### GetIBrandingColor

`func (o *BrandingRequestV2) GetIBrandingColor() int32`

GetIBrandingColor returns the IBrandingColor field if non-nil, zero value otherwise.

### GetIBrandingColorOk

`func (o *BrandingRequestV2) GetIBrandingColorOk() (*int32, bool)`

GetIBrandingColorOk returns a tuple with the IBrandingColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColor

`func (o *BrandingRequestV2) SetIBrandingColor(v int32)`

SetIBrandingColor sets IBrandingColor field to given value.


### GetSBrandingName

`func (o *BrandingRequestV2) GetSBrandingName() string`

GetSBrandingName returns the SBrandingName field if non-nil, zero value otherwise.

### GetSBrandingNameOk

`func (o *BrandingRequestV2) GetSBrandingNameOk() (*string, bool)`

GetSBrandingNameOk returns a tuple with the SBrandingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingName

`func (o *BrandingRequestV2) SetSBrandingName(v string)`

SetSBrandingName sets SBrandingName field to given value.

### HasSBrandingName

`func (o *BrandingRequestV2) HasSBrandingName() bool`

HasSBrandingName returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *BrandingRequestV2) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *BrandingRequestV2) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *BrandingRequestV2) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *BrandingRequestV2) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetBBrandingIsactive

`func (o *BrandingRequestV2) GetBBrandingIsactive() bool`

GetBBrandingIsactive returns the BBrandingIsactive field if non-nil, zero value otherwise.

### GetBBrandingIsactiveOk

`func (o *BrandingRequestV2) GetBBrandingIsactiveOk() (*bool, bool)`

GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrandingIsactive

`func (o *BrandingRequestV2) SetBBrandingIsactive(v bool)`

SetBBrandingIsactive sets BBrandingIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


