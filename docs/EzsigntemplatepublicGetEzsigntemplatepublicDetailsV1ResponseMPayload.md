# EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjBranding** | Pointer to **map[string]interface{}** | A Custom Branding Object | [optional] 
**FkiUserlogintypeID** | **int32** | The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| |6|**Embedded**|The Ezsignsigner will only be able to sign in the embedded solution. No email will be sent for invitation to sign. **Additional fee applies**|   |7|**Embedded with phone or SMS**|The Ezsignsigner will only be able to sign in the embedded solution and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|   |8|**No validation**|The Ezsignsigner will not receive an email and won&#39;t have to validate his connection using 2 factor. **Additional fee applies**|      |9|**Sms only**|The Ezsignsigner will not receive an email but will will need to authenticate using SMS. **Additional fee applies**|      | 
**ASEzsigntemplatesignerDescription** | **[]string** |  | 

## Methods

### NewEzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload

`func NewEzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload(fkiUserlogintypeID int32, aSEzsigntemplatesignerDescription []string, ) *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload`

NewEzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload instantiates a new EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayloadWithDefaults

`func NewEzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayloadWithDefaults() *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload`

NewEzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjBranding

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) GetObjBranding() map[string]interface{}`

GetObjBranding returns the ObjBranding field if non-nil, zero value otherwise.

### GetObjBrandingOk

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) GetObjBrandingOk() (*map[string]interface{}, bool)`

GetObjBrandingOk returns a tuple with the ObjBranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjBranding

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) SetObjBranding(v map[string]interface{})`

SetObjBranding sets ObjBranding field to given value.

### HasObjBranding

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) HasObjBranding() bool`

HasObjBranding returns a boolean if a field has been set.

### GetFkiUserlogintypeID

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.


### GetASEzsigntemplatesignerDescription

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) GetASEzsigntemplatesignerDescription() []string`

GetASEzsigntemplatesignerDescription returns the ASEzsigntemplatesignerDescription field if non-nil, zero value otherwise.

### GetASEzsigntemplatesignerDescriptionOk

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) GetASEzsigntemplatesignerDescriptionOk() (*[]string, bool)`

GetASEzsigntemplatesignerDescriptionOk returns a tuple with the ASEzsigntemplatesignerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASEzsigntemplatesignerDescription

`func (o *EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1ResponseMPayload) SetASEzsigntemplatesignerDescription(v []string)`

SetASEzsigntemplatesignerDescription sets ASEzsigntemplatesignerDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


