# EzsignsignerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignerID** | **int32** | The unique ID of the Ezsignsigner | 
**FkiTaxassignmentID** | **int32** | The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable| | 
**FkiSecretquestionID** | Pointer to **int32** | The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father&#39;s middle name| |15|Your mother&#39;s maiden name| |16|Name of your eldest child| |17|Your spouse&#39;s middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat&#39;s name| |22|Date of Birth (YYYY-MM-DD)| |22|Secret Code| |22|Your reference code| | [optional] 
**FkiUserlogintypeID** | **int32** | The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| |6|**Embedded**|The Ezsignsigner will only be able to sign in the embedded solution. No email will be sent for invitation to sign. **Additional fee applies**|   |7|**Embedded with phone or SMS**|The Ezsignsigner will only be able to sign in the embedded solution and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|   |8|**No validation**|The Ezsignsigner will not receive an email and won&#39;t have to validate his connection using 2 factor. **Additional fee applies**|      |9|**Sms only**|The Ezsignsigner will not receive an email but will will need to authenticate using SMS. **Additional fee applies**|      | 
**SUserlogintypeDescriptionX** | **string** | The description of the Userlogintype in the language of the requester | 

## Methods

### NewEzsignsignerResponse

`func NewEzsignsignerResponse(pkiEzsignsignerID int32, fkiTaxassignmentID int32, fkiUserlogintypeID int32, sUserlogintypeDescriptionX string, ) *EzsignsignerResponse`

NewEzsignsignerResponse instantiates a new EzsignsignerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignerResponseWithDefaults

`func NewEzsignsignerResponseWithDefaults() *EzsignsignerResponse`

NewEzsignsignerResponseWithDefaults instantiates a new EzsignsignerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignerID

`func (o *EzsignsignerResponse) GetPkiEzsignsignerID() int32`

GetPkiEzsignsignerID returns the PkiEzsignsignerID field if non-nil, zero value otherwise.

### GetPkiEzsignsignerIDOk

`func (o *EzsignsignerResponse) GetPkiEzsignsignerIDOk() (*int32, bool)`

GetPkiEzsignsignerIDOk returns a tuple with the PkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignerID

`func (o *EzsignsignerResponse) SetPkiEzsignsignerID(v int32)`

SetPkiEzsignsignerID sets PkiEzsignsignerID field to given value.


### GetFkiTaxassignmentID

`func (o *EzsignsignerResponse) GetFkiTaxassignmentID() int32`

GetFkiTaxassignmentID returns the FkiTaxassignmentID field if non-nil, zero value otherwise.

### GetFkiTaxassignmentIDOk

`func (o *EzsignsignerResponse) GetFkiTaxassignmentIDOk() (*int32, bool)`

GetFkiTaxassignmentIDOk returns a tuple with the FkiTaxassignmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTaxassignmentID

`func (o *EzsignsignerResponse) SetFkiTaxassignmentID(v int32)`

SetFkiTaxassignmentID sets FkiTaxassignmentID field to given value.


### GetFkiSecretquestionID

`func (o *EzsignsignerResponse) GetFkiSecretquestionID() int32`

GetFkiSecretquestionID returns the FkiSecretquestionID field if non-nil, zero value otherwise.

### GetFkiSecretquestionIDOk

`func (o *EzsignsignerResponse) GetFkiSecretquestionIDOk() (*int32, bool)`

GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSecretquestionID

`func (o *EzsignsignerResponse) SetFkiSecretquestionID(v int32)`

SetFkiSecretquestionID sets FkiSecretquestionID field to given value.

### HasFkiSecretquestionID

`func (o *EzsignsignerResponse) HasFkiSecretquestionID() bool`

HasFkiSecretquestionID returns a boolean if a field has been set.

### GetFkiUserlogintypeID

`func (o *EzsignsignerResponse) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsignsignerResponse) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsignsignerResponse) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.


### GetSUserlogintypeDescriptionX

`func (o *EzsignsignerResponse) GetSUserlogintypeDescriptionX() string`

GetSUserlogintypeDescriptionX returns the SUserlogintypeDescriptionX field if non-nil, zero value otherwise.

### GetSUserlogintypeDescriptionXOk

`func (o *EzsignsignerResponse) GetSUserlogintypeDescriptionXOk() (*string, bool)`

GetSUserlogintypeDescriptionXOk returns a tuple with the SUserlogintypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserlogintypeDescriptionX

`func (o *EzsignsignerResponse) SetSUserlogintypeDescriptionX(v string)`

SetSUserlogintypeDescriptionX sets SUserlogintypeDescriptionX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


