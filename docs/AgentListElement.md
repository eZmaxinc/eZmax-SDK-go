# AgentListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiAgentID** | **int32** | The unique ID of the Agent. | 
**FkiAgenttypeID** | **int32** | The unique ID of the Agenttype | 
**SAgenttypeNameX** | **string** | The name of the Agenttype in the language of the requester | 
**FkiAgentincorporationID** | Pointer to **int32** | The unique ID of the Agentincorporation. | [optional] 
**SAgentincorporationName** | Pointer to **string** | The name of the Agentincorporation | [optional] 
**FkiDepartmentID** | **int32** | The unique ID of the Department | 
**SDepartmentNameX** | **string** | The Name of the Department in the language of the requester | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**SRealestateboardnumberNumber** | Pointer to **string** | The number of the Realestateboardnumber | [optional] 
**SAgentCode** | **string** | The code of the Agent | 
**IAgentPhotocopiercode** | **int32** | The photocopiercode of the Agent | 
**IAgentLongdistancecode** | **int32** | The longdistancecode of the Agent | 
**IAgentBannernumber** | **int32** | The bannernumber of the Agent | 
**SAgentRealestateassociationlicense** | **string** | The realestateassociationlicense of the Agent | 
**DtAgentHiredate** | Pointer to **string** | The hiredate of the Agent | [optional] 
**DtAgentLeavedate** | Pointer to **string** | The leavedate of the Agent | [optional] 
**BAgentTranquillit** | **bool** | Whether if it&#39;s an tranquillit | 
**BAgentResidentiallicense** | **bool** | Whether if it&#39;s an residentiallicense | 
**BAgentCommerciallicense** | **bool** | Whether if it&#39;s an commerciallicense | 
**BAgentMortgagelicense** | **bool** | Whether if it&#39;s an mortgagelicense | 
**BAgentPaidbyofficetranquillit** | **bool** | Whether if it&#39;s an paidbyofficetranquillit | 
**DtAgentFintraccertification** | Pointer to **string** | The fintraccertification of the Agent | [optional] 
**BAgentIsactive** | **bool** | Whether the Agent is active or not | 
**SContactFirstname** | **string** | The First name of the contact | 
**SContactLastname** | **string** | The Last name of the contact | 
**DtContactBirthdate** | Pointer to **string** | The Birth Date of the contact | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SAddressCivic** | Pointer to **string** | The Civic number. | [optional] 
**SAddressStreet** | Pointer to **string** | The Street Name | [optional] 
**SAddressSuite** | Pointer to **string** | The Suite or appartment number | [optional] 
**SAddressCity** | Pointer to **string** | The City name | [optional] 
**SAddressZip** | Pointer to **string** | The Postal/Zip Code  The value must be entered without spaces | [optional] 
**SProvinceNameX** | Pointer to **string** | The name of the Province in the language of the requester | [optional] 
**SCountryNameX** | Pointer to **string** | The name of the Country in the language of the requester | [optional] 

## Methods

### NewAgentListElement

`func NewAgentListElement(pkiAgentID int32, fkiAgenttypeID int32, sAgenttypeNameX string, fkiDepartmentID int32, sDepartmentNameX string, fkiLanguageID int32, sLanguageNameX string, sAgentCode string, iAgentPhotocopiercode int32, iAgentLongdistancecode int32, iAgentBannernumber int32, sAgentRealestateassociationlicense string, bAgentTranquillit bool, bAgentResidentiallicense bool, bAgentCommerciallicense bool, bAgentMortgagelicense bool, bAgentPaidbyofficetranquillit bool, bAgentIsactive bool, sContactFirstname string, sContactLastname string, ) *AgentListElement`

NewAgentListElement instantiates a new AgentListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentListElementWithDefaults

`func NewAgentListElementWithDefaults() *AgentListElement`

NewAgentListElementWithDefaults instantiates a new AgentListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAgentID

`func (o *AgentListElement) GetPkiAgentID() int32`

GetPkiAgentID returns the PkiAgentID field if non-nil, zero value otherwise.

### GetPkiAgentIDOk

`func (o *AgentListElement) GetPkiAgentIDOk() (*int32, bool)`

GetPkiAgentIDOk returns a tuple with the PkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAgentID

`func (o *AgentListElement) SetPkiAgentID(v int32)`

SetPkiAgentID sets PkiAgentID field to given value.


### GetFkiAgenttypeID

`func (o *AgentListElement) GetFkiAgenttypeID() int32`

GetFkiAgenttypeID returns the FkiAgenttypeID field if non-nil, zero value otherwise.

### GetFkiAgenttypeIDOk

`func (o *AgentListElement) GetFkiAgenttypeIDOk() (*int32, bool)`

GetFkiAgenttypeIDOk returns a tuple with the FkiAgenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgenttypeID

`func (o *AgentListElement) SetFkiAgenttypeID(v int32)`

SetFkiAgenttypeID sets FkiAgenttypeID field to given value.


### GetSAgenttypeNameX

`func (o *AgentListElement) GetSAgenttypeNameX() string`

GetSAgenttypeNameX returns the SAgenttypeNameX field if non-nil, zero value otherwise.

### GetSAgenttypeNameXOk

`func (o *AgentListElement) GetSAgenttypeNameXOk() (*string, bool)`

GetSAgenttypeNameXOk returns a tuple with the SAgenttypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAgenttypeNameX

`func (o *AgentListElement) SetSAgenttypeNameX(v string)`

SetSAgenttypeNameX sets SAgenttypeNameX field to given value.


### GetFkiAgentincorporationID

`func (o *AgentListElement) GetFkiAgentincorporationID() int32`

GetFkiAgentincorporationID returns the FkiAgentincorporationID field if non-nil, zero value otherwise.

### GetFkiAgentincorporationIDOk

`func (o *AgentListElement) GetFkiAgentincorporationIDOk() (*int32, bool)`

GetFkiAgentincorporationIDOk returns a tuple with the FkiAgentincorporationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentincorporationID

`func (o *AgentListElement) SetFkiAgentincorporationID(v int32)`

SetFkiAgentincorporationID sets FkiAgentincorporationID field to given value.

### HasFkiAgentincorporationID

`func (o *AgentListElement) HasFkiAgentincorporationID() bool`

HasFkiAgentincorporationID returns a boolean if a field has been set.

### GetSAgentincorporationName

`func (o *AgentListElement) GetSAgentincorporationName() string`

GetSAgentincorporationName returns the SAgentincorporationName field if non-nil, zero value otherwise.

### GetSAgentincorporationNameOk

`func (o *AgentListElement) GetSAgentincorporationNameOk() (*string, bool)`

GetSAgentincorporationNameOk returns a tuple with the SAgentincorporationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAgentincorporationName

`func (o *AgentListElement) SetSAgentincorporationName(v string)`

SetSAgentincorporationName sets SAgentincorporationName field to given value.

### HasSAgentincorporationName

`func (o *AgentListElement) HasSAgentincorporationName() bool`

HasSAgentincorporationName returns a boolean if a field has been set.

### GetFkiDepartmentID

`func (o *AgentListElement) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *AgentListElement) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *AgentListElement) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### GetSDepartmentNameX

`func (o *AgentListElement) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *AgentListElement) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *AgentListElement) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.


### GetFkiLanguageID

`func (o *AgentListElement) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *AgentListElement) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *AgentListElement) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSLanguageNameX

`func (o *AgentListElement) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *AgentListElement) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *AgentListElement) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetSRealestateboardnumberNumber

`func (o *AgentListElement) GetSRealestateboardnumberNumber() string`

GetSRealestateboardnumberNumber returns the SRealestateboardnumberNumber field if non-nil, zero value otherwise.

### GetSRealestateboardnumberNumberOk

`func (o *AgentListElement) GetSRealestateboardnumberNumberOk() (*string, bool)`

GetSRealestateboardnumberNumberOk returns a tuple with the SRealestateboardnumberNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSRealestateboardnumberNumber

`func (o *AgentListElement) SetSRealestateboardnumberNumber(v string)`

SetSRealestateboardnumberNumber sets SRealestateboardnumberNumber field to given value.

### HasSRealestateboardnumberNumber

`func (o *AgentListElement) HasSRealestateboardnumberNumber() bool`

HasSRealestateboardnumberNumber returns a boolean if a field has been set.

### GetSAgentCode

`func (o *AgentListElement) GetSAgentCode() string`

GetSAgentCode returns the SAgentCode field if non-nil, zero value otherwise.

### GetSAgentCodeOk

`func (o *AgentListElement) GetSAgentCodeOk() (*string, bool)`

GetSAgentCodeOk returns a tuple with the SAgentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAgentCode

`func (o *AgentListElement) SetSAgentCode(v string)`

SetSAgentCode sets SAgentCode field to given value.


### GetIAgentPhotocopiercode

`func (o *AgentListElement) GetIAgentPhotocopiercode() int32`

GetIAgentPhotocopiercode returns the IAgentPhotocopiercode field if non-nil, zero value otherwise.

### GetIAgentPhotocopiercodeOk

`func (o *AgentListElement) GetIAgentPhotocopiercodeOk() (*int32, bool)`

GetIAgentPhotocopiercodeOk returns a tuple with the IAgentPhotocopiercode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAgentPhotocopiercode

`func (o *AgentListElement) SetIAgentPhotocopiercode(v int32)`

SetIAgentPhotocopiercode sets IAgentPhotocopiercode field to given value.


### GetIAgentLongdistancecode

`func (o *AgentListElement) GetIAgentLongdistancecode() int32`

GetIAgentLongdistancecode returns the IAgentLongdistancecode field if non-nil, zero value otherwise.

### GetIAgentLongdistancecodeOk

`func (o *AgentListElement) GetIAgentLongdistancecodeOk() (*int32, bool)`

GetIAgentLongdistancecodeOk returns a tuple with the IAgentLongdistancecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAgentLongdistancecode

`func (o *AgentListElement) SetIAgentLongdistancecode(v int32)`

SetIAgentLongdistancecode sets IAgentLongdistancecode field to given value.


### GetIAgentBannernumber

`func (o *AgentListElement) GetIAgentBannernumber() int32`

GetIAgentBannernumber returns the IAgentBannernumber field if non-nil, zero value otherwise.

### GetIAgentBannernumberOk

`func (o *AgentListElement) GetIAgentBannernumberOk() (*int32, bool)`

GetIAgentBannernumberOk returns a tuple with the IAgentBannernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAgentBannernumber

`func (o *AgentListElement) SetIAgentBannernumber(v int32)`

SetIAgentBannernumber sets IAgentBannernumber field to given value.


### GetSAgentRealestateassociationlicense

`func (o *AgentListElement) GetSAgentRealestateassociationlicense() string`

GetSAgentRealestateassociationlicense returns the SAgentRealestateassociationlicense field if non-nil, zero value otherwise.

### GetSAgentRealestateassociationlicenseOk

`func (o *AgentListElement) GetSAgentRealestateassociationlicenseOk() (*string, bool)`

GetSAgentRealestateassociationlicenseOk returns a tuple with the SAgentRealestateassociationlicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAgentRealestateassociationlicense

`func (o *AgentListElement) SetSAgentRealestateassociationlicense(v string)`

SetSAgentRealestateassociationlicense sets SAgentRealestateassociationlicense field to given value.


### GetDtAgentHiredate

`func (o *AgentListElement) GetDtAgentHiredate() string`

GetDtAgentHiredate returns the DtAgentHiredate field if non-nil, zero value otherwise.

### GetDtAgentHiredateOk

`func (o *AgentListElement) GetDtAgentHiredateOk() (*string, bool)`

GetDtAgentHiredateOk returns a tuple with the DtAgentHiredate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtAgentHiredate

`func (o *AgentListElement) SetDtAgentHiredate(v string)`

SetDtAgentHiredate sets DtAgentHiredate field to given value.

### HasDtAgentHiredate

`func (o *AgentListElement) HasDtAgentHiredate() bool`

HasDtAgentHiredate returns a boolean if a field has been set.

### GetDtAgentLeavedate

`func (o *AgentListElement) GetDtAgentLeavedate() string`

GetDtAgentLeavedate returns the DtAgentLeavedate field if non-nil, zero value otherwise.

### GetDtAgentLeavedateOk

`func (o *AgentListElement) GetDtAgentLeavedateOk() (*string, bool)`

GetDtAgentLeavedateOk returns a tuple with the DtAgentLeavedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtAgentLeavedate

`func (o *AgentListElement) SetDtAgentLeavedate(v string)`

SetDtAgentLeavedate sets DtAgentLeavedate field to given value.

### HasDtAgentLeavedate

`func (o *AgentListElement) HasDtAgentLeavedate() bool`

HasDtAgentLeavedate returns a boolean if a field has been set.

### GetBAgentTranquillit

`func (o *AgentListElement) GetBAgentTranquillit() bool`

GetBAgentTranquillit returns the BAgentTranquillit field if non-nil, zero value otherwise.

### GetBAgentTranquillitOk

`func (o *AgentListElement) GetBAgentTranquillitOk() (*bool, bool)`

GetBAgentTranquillitOk returns a tuple with the BAgentTranquillit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAgentTranquillit

`func (o *AgentListElement) SetBAgentTranquillit(v bool)`

SetBAgentTranquillit sets BAgentTranquillit field to given value.


### GetBAgentResidentiallicense

`func (o *AgentListElement) GetBAgentResidentiallicense() bool`

GetBAgentResidentiallicense returns the BAgentResidentiallicense field if non-nil, zero value otherwise.

### GetBAgentResidentiallicenseOk

`func (o *AgentListElement) GetBAgentResidentiallicenseOk() (*bool, bool)`

GetBAgentResidentiallicenseOk returns a tuple with the BAgentResidentiallicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAgentResidentiallicense

`func (o *AgentListElement) SetBAgentResidentiallicense(v bool)`

SetBAgentResidentiallicense sets BAgentResidentiallicense field to given value.


### GetBAgentCommerciallicense

`func (o *AgentListElement) GetBAgentCommerciallicense() bool`

GetBAgentCommerciallicense returns the BAgentCommerciallicense field if non-nil, zero value otherwise.

### GetBAgentCommerciallicenseOk

`func (o *AgentListElement) GetBAgentCommerciallicenseOk() (*bool, bool)`

GetBAgentCommerciallicenseOk returns a tuple with the BAgentCommerciallicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAgentCommerciallicense

`func (o *AgentListElement) SetBAgentCommerciallicense(v bool)`

SetBAgentCommerciallicense sets BAgentCommerciallicense field to given value.


### GetBAgentMortgagelicense

`func (o *AgentListElement) GetBAgentMortgagelicense() bool`

GetBAgentMortgagelicense returns the BAgentMortgagelicense field if non-nil, zero value otherwise.

### GetBAgentMortgagelicenseOk

`func (o *AgentListElement) GetBAgentMortgagelicenseOk() (*bool, bool)`

GetBAgentMortgagelicenseOk returns a tuple with the BAgentMortgagelicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAgentMortgagelicense

`func (o *AgentListElement) SetBAgentMortgagelicense(v bool)`

SetBAgentMortgagelicense sets BAgentMortgagelicense field to given value.


### GetBAgentPaidbyofficetranquillit

`func (o *AgentListElement) GetBAgentPaidbyofficetranquillit() bool`

GetBAgentPaidbyofficetranquillit returns the BAgentPaidbyofficetranquillit field if non-nil, zero value otherwise.

### GetBAgentPaidbyofficetranquillitOk

`func (o *AgentListElement) GetBAgentPaidbyofficetranquillitOk() (*bool, bool)`

GetBAgentPaidbyofficetranquillitOk returns a tuple with the BAgentPaidbyofficetranquillit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAgentPaidbyofficetranquillit

`func (o *AgentListElement) SetBAgentPaidbyofficetranquillit(v bool)`

SetBAgentPaidbyofficetranquillit sets BAgentPaidbyofficetranquillit field to given value.


### GetDtAgentFintraccertification

`func (o *AgentListElement) GetDtAgentFintraccertification() string`

GetDtAgentFintraccertification returns the DtAgentFintraccertification field if non-nil, zero value otherwise.

### GetDtAgentFintraccertificationOk

`func (o *AgentListElement) GetDtAgentFintraccertificationOk() (*string, bool)`

GetDtAgentFintraccertificationOk returns a tuple with the DtAgentFintraccertification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtAgentFintraccertification

`func (o *AgentListElement) SetDtAgentFintraccertification(v string)`

SetDtAgentFintraccertification sets DtAgentFintraccertification field to given value.

### HasDtAgentFintraccertification

`func (o *AgentListElement) HasDtAgentFintraccertification() bool`

HasDtAgentFintraccertification returns a boolean if a field has been set.

### GetBAgentIsactive

`func (o *AgentListElement) GetBAgentIsactive() bool`

GetBAgentIsactive returns the BAgentIsactive field if non-nil, zero value otherwise.

### GetBAgentIsactiveOk

`func (o *AgentListElement) GetBAgentIsactiveOk() (*bool, bool)`

GetBAgentIsactiveOk returns a tuple with the BAgentIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAgentIsactive

`func (o *AgentListElement) SetBAgentIsactive(v bool)`

SetBAgentIsactive sets BAgentIsactive field to given value.


### GetSContactFirstname

`func (o *AgentListElement) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *AgentListElement) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *AgentListElement) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.


### GetSContactLastname

`func (o *AgentListElement) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *AgentListElement) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *AgentListElement) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.


### GetDtContactBirthdate

`func (o *AgentListElement) GetDtContactBirthdate() string`

GetDtContactBirthdate returns the DtContactBirthdate field if non-nil, zero value otherwise.

### GetDtContactBirthdateOk

`func (o *AgentListElement) GetDtContactBirthdateOk() (*string, bool)`

GetDtContactBirthdateOk returns a tuple with the DtContactBirthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtContactBirthdate

`func (o *AgentListElement) SetDtContactBirthdate(v string)`

SetDtContactBirthdate sets DtContactBirthdate field to given value.

### HasDtContactBirthdate

`func (o *AgentListElement) HasDtContactBirthdate() bool`

HasDtContactBirthdate returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *AgentListElement) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *AgentListElement) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *AgentListElement) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *AgentListElement) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *AgentListElement) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *AgentListElement) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *AgentListElement) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *AgentListElement) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSAddressCivic

`func (o *AgentListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *AgentListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *AgentListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *AgentListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *AgentListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *AgentListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *AgentListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *AgentListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *AgentListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *AgentListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *AgentListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *AgentListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *AgentListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *AgentListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *AgentListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *AgentListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSAddressZip

`func (o *AgentListElement) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *AgentListElement) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *AgentListElement) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.

### HasSAddressZip

`func (o *AgentListElement) HasSAddressZip() bool`

HasSAddressZip returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *AgentListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *AgentListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *AgentListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *AgentListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *AgentListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *AgentListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *AgentListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *AgentListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


