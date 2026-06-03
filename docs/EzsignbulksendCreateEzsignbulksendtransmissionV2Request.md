# EzsignbulksendCreateEzsignbulksendtransmissionV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiUserlogintypeID** | **int32** | The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| |6|**Embedded**|The Ezsignsigner will only be able to sign in the embedded solution. No email will be sent for invitation to sign. **Additional fee applies**|   |7|**Embedded with phone or SMS**|The Ezsignsigner will only be able to sign in the embedded solution and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|   |8|**No validation**|The Ezsignsigner will not receive an email and won&#39;t have to validate his connection using 2 factor. **Additional fee applies**|      |9|**Sms only**|The Ezsignsigner will not receive an email but will will need to authenticate using SMS. **Additional fee applies**|      | 
**FkiSecretquestionID** | Pointer to **int32** | The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father&#39;s middle name| |15|Your mother&#39;s maiden name| |16|Name of your eldest child| |17|Your spouse&#39;s middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat&#39;s name| |22|Date of Birth (YYYY-MM-DD)| |23|Secret Code| |24|Your reference code| |25|What are the last 4 digits of your SIN| |26|What is your postal code| |27|What is your employee number| |28|What is your manager’s first name| |29|What is your file number| |30|What is your client/member number| |31|What is your license number| |32|What are the last 4 digits of your phone number| |33|What is your student number| | [optional] 
**FkiEzsigntsarequirementID** | Pointer to **int32** | The unique ID of the Ezsigntsarequirement.  Determine if a Time Stamping Authority should add a timestamp on each of the signature. Valid values:  |Value|Description| |-|-| |1|No. TSA Timestamping will requested. This will make all signatures a lot faster since no round-trip to the TSA server will be required. Timestamping will be made using eZsign server&#39;s time.| |2|Best effort. Timestamping from a Time Stamping Authority will be requested but is not mandatory. In the very improbable case it cannot be completed, the timestamping will be made using eZsign server&#39;s time. **Additional fee applies**| |3|Mandatory. Timestamping from a Time Stamping Authority will be requested and is mandatory. In the very improbable case it cannot be completed, the signature will fail and the user will be asked to retry. **Additional fee applies**| | [optional] 
**SEzsignbulksendtransmissionDescription** | **string** | The description of the Ezsignbulksendtransmission | 
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the Ezsigndocument can be signed. | 
**IEzsignfolderSendreminderfirstdays** | **int32** | The number of days before the first reminder sending | 
**IEzsignfolderSendreminderotherdays** | **int32** | The number of days after the first reminder sending | 
**TExtraMessage** | **string** | A custom text message that will be added to the email sent. | 
**SCsvBase64** | **string** | The Base64 encoded binary content of the CSV file. | 

## Methods

### NewEzsignbulksendCreateEzsignbulksendtransmissionV2Request

`func NewEzsignbulksendCreateEzsignbulksendtransmissionV2Request(fkiUserlogintypeID int32, sEzsignbulksendtransmissionDescription string, dtEzsigndocumentDuedate string, iEzsignfolderSendreminderfirstdays int32, iEzsignfolderSendreminderotherdays int32, tExtraMessage string, sCsvBase64 string, ) *EzsignbulksendCreateEzsignbulksendtransmissionV2Request`

NewEzsignbulksendCreateEzsignbulksendtransmissionV2Request instantiates a new EzsignbulksendCreateEzsignbulksendtransmissionV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendCreateEzsignbulksendtransmissionV2RequestWithDefaults

`func NewEzsignbulksendCreateEzsignbulksendtransmissionV2RequestWithDefaults() *EzsignbulksendCreateEzsignbulksendtransmissionV2Request`

NewEzsignbulksendCreateEzsignbulksendtransmissionV2RequestWithDefaults instantiates a new EzsignbulksendCreateEzsignbulksendtransmissionV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiUserlogintypeID

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.


### GetFkiSecretquestionID

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetFkiSecretquestionID() int32`

GetFkiSecretquestionID returns the FkiSecretquestionID field if non-nil, zero value otherwise.

### GetFkiSecretquestionIDOk

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetFkiSecretquestionIDOk() (*int32, bool)`

GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSecretquestionID

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetFkiSecretquestionID(v int32)`

SetFkiSecretquestionID sets FkiSecretquestionID field to given value.

### HasFkiSecretquestionID

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) HasFkiSecretquestionID() bool`

HasFkiSecretquestionID returns a boolean if a field has been set.

### GetFkiEzsigntsarequirementID

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetFkiEzsigntsarequirementID() int32`

GetFkiEzsigntsarequirementID returns the FkiEzsigntsarequirementID field if non-nil, zero value otherwise.

### GetFkiEzsigntsarequirementIDOk

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetFkiEzsigntsarequirementIDOk() (*int32, bool)`

GetFkiEzsigntsarequirementIDOk returns a tuple with the FkiEzsigntsarequirementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntsarequirementID

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetFkiEzsigntsarequirementID(v int32)`

SetFkiEzsigntsarequirementID sets FkiEzsigntsarequirementID field to given value.

### HasFkiEzsigntsarequirementID

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) HasFkiEzsigntsarequirementID() bool`

HasFkiEzsigntsarequirementID returns a boolean if a field has been set.

### GetSEzsignbulksendtransmissionDescription

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetSEzsignbulksendtransmissionDescription() string`

GetSEzsignbulksendtransmissionDescription returns the SEzsignbulksendtransmissionDescription field if non-nil, zero value otherwise.

### GetSEzsignbulksendtransmissionDescriptionOk

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetSEzsignbulksendtransmissionDescriptionOk() (*string, bool)`

GetSEzsignbulksendtransmissionDescriptionOk returns a tuple with the SEzsignbulksendtransmissionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignbulksendtransmissionDescription

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetSEzsignbulksendtransmissionDescription(v string)`

SetSEzsignbulksendtransmissionDescription sets SEzsignbulksendtransmissionDescription field to given value.


### GetDtEzsigndocumentDuedate

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetIEzsignfolderSendreminderfirstdays

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetIEzsignfolderSendreminderfirstdays() int32`

GetIEzsignfolderSendreminderfirstdays returns the IEzsignfolderSendreminderfirstdays field if non-nil, zero value otherwise.

### GetIEzsignfolderSendreminderfirstdaysOk

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetIEzsignfolderSendreminderfirstdaysOk() (*int32, bool)`

GetIEzsignfolderSendreminderfirstdaysOk returns a tuple with the IEzsignfolderSendreminderfirstdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSendreminderfirstdays

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetIEzsignfolderSendreminderfirstdays(v int32)`

SetIEzsignfolderSendreminderfirstdays sets IEzsignfolderSendreminderfirstdays field to given value.


### GetIEzsignfolderSendreminderotherdays

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetIEzsignfolderSendreminderotherdays() int32`

GetIEzsignfolderSendreminderotherdays returns the IEzsignfolderSendreminderotherdays field if non-nil, zero value otherwise.

### GetIEzsignfolderSendreminderotherdaysOk

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetIEzsignfolderSendreminderotherdaysOk() (*int32, bool)`

GetIEzsignfolderSendreminderotherdaysOk returns a tuple with the IEzsignfolderSendreminderotherdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolderSendreminderotherdays

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetIEzsignfolderSendreminderotherdays(v int32)`

SetIEzsignfolderSendreminderotherdays sets IEzsignfolderSendreminderotherdays field to given value.


### GetTExtraMessage

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetTExtraMessage() string`

GetTExtraMessage returns the TExtraMessage field if non-nil, zero value otherwise.

### GetTExtraMessageOk

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetTExtraMessageOk() (*string, bool)`

GetTExtraMessageOk returns a tuple with the TExtraMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTExtraMessage

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetTExtraMessage(v string)`

SetTExtraMessage sets TExtraMessage field to given value.


### GetSCsvBase64

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetSCsvBase64() string`

GetSCsvBase64 returns the SCsvBase64 field if non-nil, zero value otherwise.

### GetSCsvBase64Ok

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) GetSCsvBase64Ok() (*string, bool)`

GetSCsvBase64Ok returns a tuple with the SCsvBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCsvBase64

`func (o *EzsignbulksendCreateEzsignbulksendtransmissionV2Request) SetSCsvBase64(v string)`

SetSCsvBase64 sets SCsvBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


