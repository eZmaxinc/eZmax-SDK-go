# EzsigntemplatepublicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepublicID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepublic | [optional] 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiUserlogintypeID** | **int32** | The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| |6|**Embedded**|The Ezsignsigner will only be able to sign in the embedded solution. No email will be sent for invitation to sign. **Additional fee applies**|   |7|**Embedded with phone or SMS**|The Ezsignsigner will only be able to sign in the embedded solution and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|   |8|**No validation**|The Ezsignsigner will not receive an email and won&#39;t have to validate his connection using 2 factor. **Additional fee applies**|      |9|**Sms only**|The Ezsignsigner will not receive an email but will will need to authenticate using SMS. **Additional fee applies**|      | 
**FkiEzsigntemplateID** | Pointer to **int32** | The unique ID of the Ezsigntemplate | [optional] 
**FkiEzsigntemplatepackageID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepackage | [optional] 
**SEzsigntemplatepublicDescription** | **string** | The description of the Ezsigntemplatepublic | 
**BEzsigntemplatepublicIsactive** | **bool** | Whether the ezsigntemplatepublic is active or not | 
**TEzsigntemplatepublicNote** | **string** | The note of the Ezsigntemplatepublic | 
**EEzsigntemplatepublicLimittype** | [**FieldEEzsigntemplatepublicLimittype**](FieldEEzsigntemplatepublicLimittype.md) |  | 
**IEzsigntemplatepublicLimit** | **int32** | The limit of the Ezsigntemplatepublic | 

## Methods

### NewEzsigntemplatepublicRequest

`func NewEzsigntemplatepublicRequest(fkiEzsignfoldertypeID int32, fkiUserlogintypeID int32, sEzsigntemplatepublicDescription string, bEzsigntemplatepublicIsactive bool, tEzsigntemplatepublicNote string, eEzsigntemplatepublicLimittype FieldEEzsigntemplatepublicLimittype, iEzsigntemplatepublicLimit int32, ) *EzsigntemplatepublicRequest`

NewEzsigntemplatepublicRequest instantiates a new EzsigntemplatepublicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepublicRequestWithDefaults

`func NewEzsigntemplatepublicRequestWithDefaults() *EzsigntemplatepublicRequest`

NewEzsigntemplatepublicRequestWithDefaults instantiates a new EzsigntemplatepublicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicRequest) GetPkiEzsigntemplatepublicID() int32`

GetPkiEzsigntemplatepublicID returns the PkiEzsigntemplatepublicID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepublicIDOk

`func (o *EzsigntemplatepublicRequest) GetPkiEzsigntemplatepublicIDOk() (*int32, bool)`

GetPkiEzsigntemplatepublicIDOk returns a tuple with the PkiEzsigntemplatepublicID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicRequest) SetPkiEzsigntemplatepublicID(v int32)`

SetPkiEzsigntemplatepublicID sets PkiEzsigntemplatepublicID field to given value.

### HasPkiEzsigntemplatepublicID

`func (o *EzsigntemplatepublicRequest) HasPkiEzsigntemplatepublicID() bool`

HasPkiEzsigntemplatepublicID returns a boolean if a field has been set.

### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepublicRequest) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepublicRequest) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepublicRequest) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiUserlogintypeID

`func (o *EzsigntemplatepublicRequest) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsigntemplatepublicRequest) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsigntemplatepublicRequest) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatepublicRequest) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatepublicRequest) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatepublicRequest) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.

### HasFkiEzsigntemplateID

`func (o *EzsigntemplatepublicRequest) HasFkiEzsigntemplateID() bool`

HasFkiEzsigntemplateID returns a boolean if a field has been set.

### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicRequest) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepublicRequest) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicRequest) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.

### HasFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepublicRequest) HasFkiEzsigntemplatepackageID() bool`

HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.

### GetSEzsigntemplatepublicDescription

`func (o *EzsigntemplatepublicRequest) GetSEzsigntemplatepublicDescription() string`

GetSEzsigntemplatepublicDescription returns the SEzsigntemplatepublicDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepublicDescriptionOk

`func (o *EzsigntemplatepublicRequest) GetSEzsigntemplatepublicDescriptionOk() (*string, bool)`

GetSEzsigntemplatepublicDescriptionOk returns a tuple with the SEzsigntemplatepublicDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepublicDescription

`func (o *EzsigntemplatepublicRequest) SetSEzsigntemplatepublicDescription(v string)`

SetSEzsigntemplatepublicDescription sets SEzsigntemplatepublicDescription field to given value.


### GetBEzsigntemplatepublicIsactive

`func (o *EzsigntemplatepublicRequest) GetBEzsigntemplatepublicIsactive() bool`

GetBEzsigntemplatepublicIsactive returns the BEzsigntemplatepublicIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepublicIsactiveOk

`func (o *EzsigntemplatepublicRequest) GetBEzsigntemplatepublicIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepublicIsactiveOk returns a tuple with the BEzsigntemplatepublicIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepublicIsactive

`func (o *EzsigntemplatepublicRequest) SetBEzsigntemplatepublicIsactive(v bool)`

SetBEzsigntemplatepublicIsactive sets BEzsigntemplatepublicIsactive field to given value.


### GetTEzsigntemplatepublicNote

`func (o *EzsigntemplatepublicRequest) GetTEzsigntemplatepublicNote() string`

GetTEzsigntemplatepublicNote returns the TEzsigntemplatepublicNote field if non-nil, zero value otherwise.

### GetTEzsigntemplatepublicNoteOk

`func (o *EzsigntemplatepublicRequest) GetTEzsigntemplatepublicNoteOk() (*string, bool)`

GetTEzsigntemplatepublicNoteOk returns a tuple with the TEzsigntemplatepublicNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatepublicNote

`func (o *EzsigntemplatepublicRequest) SetTEzsigntemplatepublicNote(v string)`

SetTEzsigntemplatepublicNote sets TEzsigntemplatepublicNote field to given value.


### GetEEzsigntemplatepublicLimittype

`func (o *EzsigntemplatepublicRequest) GetEEzsigntemplatepublicLimittype() FieldEEzsigntemplatepublicLimittype`

GetEEzsigntemplatepublicLimittype returns the EEzsigntemplatepublicLimittype field if non-nil, zero value otherwise.

### GetEEzsigntemplatepublicLimittypeOk

`func (o *EzsigntemplatepublicRequest) GetEEzsigntemplatepublicLimittypeOk() (*FieldEEzsigntemplatepublicLimittype, bool)`

GetEEzsigntemplatepublicLimittypeOk returns a tuple with the EEzsigntemplatepublicLimittype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepublicLimittype

`func (o *EzsigntemplatepublicRequest) SetEEzsigntemplatepublicLimittype(v FieldEEzsigntemplatepublicLimittype)`

SetEEzsigntemplatepublicLimittype sets EEzsigntemplatepublicLimittype field to given value.


### GetIEzsigntemplatepublicLimit

`func (o *EzsigntemplatepublicRequest) GetIEzsigntemplatepublicLimit() int32`

GetIEzsigntemplatepublicLimit returns the IEzsigntemplatepublicLimit field if non-nil, zero value otherwise.

### GetIEzsigntemplatepublicLimitOk

`func (o *EzsigntemplatepublicRequest) GetIEzsigntemplatepublicLimitOk() (*int32, bool)`

GetIEzsigntemplatepublicLimitOk returns a tuple with the IEzsigntemplatepublicLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepublicLimit

`func (o *EzsigntemplatepublicRequest) SetIEzsigntemplatepublicLimit(v int32)`

SetIEzsigntemplatepublicLimit sets IEzsigntemplatepublicLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


