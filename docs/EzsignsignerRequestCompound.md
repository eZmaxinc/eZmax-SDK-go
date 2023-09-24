# EzsignsignerRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiUserlogintypeID** | Pointer to **int32** | The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| | [optional] 
**FkiTaxassignmentID** | **int32** | The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable| | 
**FkiSecretquestionID** | Pointer to **int32** | The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father&#39;s middle name| |15|Your mother&#39;s maiden name| |16|Name of your eldest child| |17|Your spouse&#39;s middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat&#39;s name| |22|Date of Birth (YYYY-MM-DD)| | [optional] 
**EEzsignsignerLogintype** | Pointer to **string** | The method the Ezsignsigner will authenticate to the signing platform.  1. **Password** means the Ezsignsigner will receive a secure link by email. 2. **PasswordPhone** means the Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**. 3. **PasswordQuestion** means the Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer. 4. **InPersonPhone** means the Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**. 5. **InPerson** means the Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type. | [optional] 
**SEzsignsignerSecretanswer** | Pointer to **string** | The predefined answer to the secret question the Ezsignsigner will need to provide to successfully authenticate. | [optional] 
**ObjContact** | [**EzsignsignerRequestCompoundContact**](EzsignsignerRequestCompoundContact.md) |  | 

## Methods

### NewEzsignsignerRequestCompound

`func NewEzsignsignerRequestCompound(fkiTaxassignmentID int32, objContact EzsignsignerRequestCompoundContact, ) *EzsignsignerRequestCompound`

NewEzsignsignerRequestCompound instantiates a new EzsignsignerRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignerRequestCompoundWithDefaults

`func NewEzsignsignerRequestCompoundWithDefaults() *EzsignsignerRequestCompound`

NewEzsignsignerRequestCompoundWithDefaults instantiates a new EzsignsignerRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiUserlogintypeID

`func (o *EzsignsignerRequestCompound) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsignsignerRequestCompound) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsignsignerRequestCompound) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.

### HasFkiUserlogintypeID

`func (o *EzsignsignerRequestCompound) HasFkiUserlogintypeID() bool`

HasFkiUserlogintypeID returns a boolean if a field has been set.

### GetFkiTaxassignmentID

`func (o *EzsignsignerRequestCompound) GetFkiTaxassignmentID() int32`

GetFkiTaxassignmentID returns the FkiTaxassignmentID field if non-nil, zero value otherwise.

### GetFkiTaxassignmentIDOk

`func (o *EzsignsignerRequestCompound) GetFkiTaxassignmentIDOk() (*int32, bool)`

GetFkiTaxassignmentIDOk returns a tuple with the FkiTaxassignmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTaxassignmentID

`func (o *EzsignsignerRequestCompound) SetFkiTaxassignmentID(v int32)`

SetFkiTaxassignmentID sets FkiTaxassignmentID field to given value.


### GetFkiSecretquestionID

`func (o *EzsignsignerRequestCompound) GetFkiSecretquestionID() int32`

GetFkiSecretquestionID returns the FkiSecretquestionID field if non-nil, zero value otherwise.

### GetFkiSecretquestionIDOk

`func (o *EzsignsignerRequestCompound) GetFkiSecretquestionIDOk() (*int32, bool)`

GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSecretquestionID

`func (o *EzsignsignerRequestCompound) SetFkiSecretquestionID(v int32)`

SetFkiSecretquestionID sets FkiSecretquestionID field to given value.

### HasFkiSecretquestionID

`func (o *EzsignsignerRequestCompound) HasFkiSecretquestionID() bool`

HasFkiSecretquestionID returns a boolean if a field has been set.

### GetEEzsignsignerLogintype

`func (o *EzsignsignerRequestCompound) GetEEzsignsignerLogintype() string`

GetEEzsignsignerLogintype returns the EEzsignsignerLogintype field if non-nil, zero value otherwise.

### GetEEzsignsignerLogintypeOk

`func (o *EzsignsignerRequestCompound) GetEEzsignsignerLogintypeOk() (*string, bool)`

GetEEzsignsignerLogintypeOk returns a tuple with the EEzsignsignerLogintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignerLogintype

`func (o *EzsignsignerRequestCompound) SetEEzsignsignerLogintype(v string)`

SetEEzsignsignerLogintype sets EEzsignsignerLogintype field to given value.

### HasEEzsignsignerLogintype

`func (o *EzsignsignerRequestCompound) HasEEzsignsignerLogintype() bool`

HasEEzsignsignerLogintype returns a boolean if a field has been set.

### GetSEzsignsignerSecretanswer

`func (o *EzsignsignerRequestCompound) GetSEzsignsignerSecretanswer() string`

GetSEzsignsignerSecretanswer returns the SEzsignsignerSecretanswer field if non-nil, zero value otherwise.

### GetSEzsignsignerSecretanswerOk

`func (o *EzsignsignerRequestCompound) GetSEzsignsignerSecretanswerOk() (*string, bool)`

GetSEzsignsignerSecretanswerOk returns a tuple with the SEzsignsignerSecretanswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignerSecretanswer

`func (o *EzsignsignerRequestCompound) SetSEzsignsignerSecretanswer(v string)`

SetSEzsignsignerSecretanswer sets SEzsignsignerSecretanswer field to given value.

### HasSEzsignsignerSecretanswer

`func (o *EzsignsignerRequestCompound) HasSEzsignsignerSecretanswer() bool`

HasSEzsignsignerSecretanswer returns a boolean if a field has been set.

### GetObjContact

`func (o *EzsignsignerRequestCompound) GetObjContact() EzsignsignerRequestCompoundContact`

GetObjContact returns the ObjContact field if non-nil, zero value otherwise.

### GetObjContactOk

`func (o *EzsignsignerRequestCompound) GetObjContactOk() (*EzsignsignerRequestCompoundContact, bool)`

GetObjContactOk returns a tuple with the ObjContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContact

`func (o *EzsignsignerRequestCompound) SetObjContact(v EzsignsignerRequestCompoundContact)`

SetObjContact sets ObjContact field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


