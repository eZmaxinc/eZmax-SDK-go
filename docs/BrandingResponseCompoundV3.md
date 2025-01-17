# BrandingResponseCompoundV3

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
**EBrandingAlignlogo** | [**FieldEBrandingAlignlogo**](FieldEBrandingAlignlogo.md) |  | 
**IBrandingColor** | **int32** | The primary color. This is a RGB color converted into integer | 
**BBrandingIsactive** | **bool** | Whether the Branding is active or not | 
**SBrandingLogourl** | Pointer to **string** | The url of the picture used as logo in the Branding | [optional] 
**SBrandingLogoemailurl** | Pointer to **string** | The url of the picture used in email as logo in the Branding | [optional] 
**SBrandingLogointerfaceurl** | Pointer to **string** | The url of the picture used as logo in the Branding | [optional] 

## Methods

### NewBrandingResponseCompoundV3

`func NewBrandingResponseCompoundV3(pkiBrandingID int32, objBrandingDescription MultilingualBrandingDescription, sBrandingDescriptionX string, eBrandingLogo FieldEBrandingLogo, eBrandingAlignlogo FieldEBrandingAlignlogo, iBrandingColor int32, bBrandingIsactive bool, ) *BrandingResponseCompoundV3`

NewBrandingResponseCompoundV3 instantiates a new BrandingResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingResponseCompoundV3WithDefaults

`func NewBrandingResponseCompoundV3WithDefaults() *BrandingResponseCompoundV3`

NewBrandingResponseCompoundV3WithDefaults instantiates a new BrandingResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBrandingID

`func (o *BrandingResponseCompoundV3) GetPkiBrandingID() int32`

GetPkiBrandingID returns the PkiBrandingID field if non-nil, zero value otherwise.

### GetPkiBrandingIDOk

`func (o *BrandingResponseCompoundV3) GetPkiBrandingIDOk() (*int32, bool)`

GetPkiBrandingIDOk returns a tuple with the PkiBrandingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBrandingID

`func (o *BrandingResponseCompoundV3) SetPkiBrandingID(v int32)`

SetPkiBrandingID sets PkiBrandingID field to given value.


### GetFkiEmailID

`func (o *BrandingResponseCompoundV3) GetFkiEmailID() int32`

GetFkiEmailID returns the FkiEmailID field if non-nil, zero value otherwise.

### GetFkiEmailIDOk

`func (o *BrandingResponseCompoundV3) GetFkiEmailIDOk() (*int32, bool)`

GetFkiEmailIDOk returns a tuple with the FkiEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmailID

`func (o *BrandingResponseCompoundV3) SetFkiEmailID(v int32)`

SetFkiEmailID sets FkiEmailID field to given value.

### HasFkiEmailID

`func (o *BrandingResponseCompoundV3) HasFkiEmailID() bool`

HasFkiEmailID returns a boolean if a field has been set.

### GetObjBrandingDescription

`func (o *BrandingResponseCompoundV3) GetObjBrandingDescription() MultilingualBrandingDescription`

GetObjBrandingDescription returns the ObjBrandingDescription field if non-nil, zero value otherwise.

### GetObjBrandingDescriptionOk

`func (o *BrandingResponseCompoundV3) GetObjBrandingDescriptionOk() (*MultilingualBrandingDescription, bool)`

GetObjBrandingDescriptionOk returns a tuple with the ObjBrandingDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjBrandingDescription

`func (o *BrandingResponseCompoundV3) SetObjBrandingDescription(v MultilingualBrandingDescription)`

SetObjBrandingDescription sets ObjBrandingDescription field to given value.


### GetSBrandingDescriptionX

`func (o *BrandingResponseCompoundV3) GetSBrandingDescriptionX() string`

GetSBrandingDescriptionX returns the SBrandingDescriptionX field if non-nil, zero value otherwise.

### GetSBrandingDescriptionXOk

`func (o *BrandingResponseCompoundV3) GetSBrandingDescriptionXOk() (*string, bool)`

GetSBrandingDescriptionXOk returns a tuple with the SBrandingDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingDescriptionX

`func (o *BrandingResponseCompoundV3) SetSBrandingDescriptionX(v string)`

SetSBrandingDescriptionX sets SBrandingDescriptionX field to given value.


### GetSBrandingName

`func (o *BrandingResponseCompoundV3) GetSBrandingName() string`

GetSBrandingName returns the SBrandingName field if non-nil, zero value otherwise.

### GetSBrandingNameOk

`func (o *BrandingResponseCompoundV3) GetSBrandingNameOk() (*string, bool)`

GetSBrandingNameOk returns a tuple with the SBrandingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingName

`func (o *BrandingResponseCompoundV3) SetSBrandingName(v string)`

SetSBrandingName sets SBrandingName field to given value.

### HasSBrandingName

`func (o *BrandingResponseCompoundV3) HasSBrandingName() bool`

HasSBrandingName returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *BrandingResponseCompoundV3) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *BrandingResponseCompoundV3) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *BrandingResponseCompoundV3) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *BrandingResponseCompoundV3) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetEBrandingLogo

`func (o *BrandingResponseCompoundV3) GetEBrandingLogo() FieldEBrandingLogo`

GetEBrandingLogo returns the EBrandingLogo field if non-nil, zero value otherwise.

### GetEBrandingLogoOk

`func (o *BrandingResponseCompoundV3) GetEBrandingLogoOk() (*FieldEBrandingLogo, bool)`

GetEBrandingLogoOk returns a tuple with the EBrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingLogo

`func (o *BrandingResponseCompoundV3) SetEBrandingLogo(v FieldEBrandingLogo)`

SetEBrandingLogo sets EBrandingLogo field to given value.


### GetEBrandingAlignlogo

`func (o *BrandingResponseCompoundV3) GetEBrandingAlignlogo() FieldEBrandingAlignlogo`

GetEBrandingAlignlogo returns the EBrandingAlignlogo field if non-nil, zero value otherwise.

### GetEBrandingAlignlogoOk

`func (o *BrandingResponseCompoundV3) GetEBrandingAlignlogoOk() (*FieldEBrandingAlignlogo, bool)`

GetEBrandingAlignlogoOk returns a tuple with the EBrandingAlignlogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBrandingAlignlogo

`func (o *BrandingResponseCompoundV3) SetEBrandingAlignlogo(v FieldEBrandingAlignlogo)`

SetEBrandingAlignlogo sets EBrandingAlignlogo field to given value.


### GetIBrandingColor

`func (o *BrandingResponseCompoundV3) GetIBrandingColor() int32`

GetIBrandingColor returns the IBrandingColor field if non-nil, zero value otherwise.

### GetIBrandingColorOk

`func (o *BrandingResponseCompoundV3) GetIBrandingColorOk() (*int32, bool)`

GetIBrandingColorOk returns a tuple with the IBrandingColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBrandingColor

`func (o *BrandingResponseCompoundV3) SetIBrandingColor(v int32)`

SetIBrandingColor sets IBrandingColor field to given value.


### GetBBrandingIsactive

`func (o *BrandingResponseCompoundV3) GetBBrandingIsactive() bool`

GetBBrandingIsactive returns the BBrandingIsactive field if non-nil, zero value otherwise.

### GetBBrandingIsactiveOk

`func (o *BrandingResponseCompoundV3) GetBBrandingIsactiveOk() (*bool, bool)`

GetBBrandingIsactiveOk returns a tuple with the BBrandingIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBrandingIsactive

`func (o *BrandingResponseCompoundV3) SetBBrandingIsactive(v bool)`

SetBBrandingIsactive sets BBrandingIsactive field to given value.


### GetSBrandingLogourl

`func (o *BrandingResponseCompoundV3) GetSBrandingLogourl() string`

GetSBrandingLogourl returns the SBrandingLogourl field if non-nil, zero value otherwise.

### GetSBrandingLogourlOk

`func (o *BrandingResponseCompoundV3) GetSBrandingLogourlOk() (*string, bool)`

GetSBrandingLogourlOk returns a tuple with the SBrandingLogourl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingLogourl

`func (o *BrandingResponseCompoundV3) SetSBrandingLogourl(v string)`

SetSBrandingLogourl sets SBrandingLogourl field to given value.

### HasSBrandingLogourl

`func (o *BrandingResponseCompoundV3) HasSBrandingLogourl() bool`

HasSBrandingLogourl returns a boolean if a field has been set.

### GetSBrandingLogoemailurl

`func (o *BrandingResponseCompoundV3) GetSBrandingLogoemailurl() string`

GetSBrandingLogoemailurl returns the SBrandingLogoemailurl field if non-nil, zero value otherwise.

### GetSBrandingLogoemailurlOk

`func (o *BrandingResponseCompoundV3) GetSBrandingLogoemailurlOk() (*string, bool)`

GetSBrandingLogoemailurlOk returns a tuple with the SBrandingLogoemailurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingLogoemailurl

`func (o *BrandingResponseCompoundV3) SetSBrandingLogoemailurl(v string)`

SetSBrandingLogoemailurl sets SBrandingLogoemailurl field to given value.

### HasSBrandingLogoemailurl

`func (o *BrandingResponseCompoundV3) HasSBrandingLogoemailurl() bool`

HasSBrandingLogoemailurl returns a boolean if a field has been set.

### GetSBrandingLogointerfaceurl

`func (o *BrandingResponseCompoundV3) GetSBrandingLogointerfaceurl() string`

GetSBrandingLogointerfaceurl returns the SBrandingLogointerfaceurl field if non-nil, zero value otherwise.

### GetSBrandingLogointerfaceurlOk

`func (o *BrandingResponseCompoundV3) GetSBrandingLogointerfaceurlOk() (*string, bool)`

GetSBrandingLogointerfaceurlOk returns a tuple with the SBrandingLogointerfaceurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBrandingLogointerfaceurl

`func (o *BrandingResponseCompoundV3) SetSBrandingLogointerfaceurl(v string)`

SetSBrandingLogointerfaceurl sets SBrandingLogointerfaceurl field to given value.

### HasSBrandingLogointerfaceurl

`func (o *BrandingResponseCompoundV3) HasSBrandingLogointerfaceurl() bool`

HasSBrandingLogointerfaceurl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


