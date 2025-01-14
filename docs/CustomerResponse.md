# CustomerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCustomerID** | **int32** | The unique ID of the Customer. | 
**FkiCompanyID** | **int32** | The unique ID of the Company | 
**FkiCustomergroupID** | **int32** | The unique ID of the Customergroup | 
**SCustomerName** | **string** | The name of the Customer | 
**FkiContactinformationsID** | **int32** | The unique ID of the Contactinformations | 
**FkiContactcontainerID** | **int32** | The unique ID of the Contactcontainer | 
**FkiImageID** | **int32** | The unique ID of the Image | 
**FkiGlaccountcontainerID** | **int32** | The unique ID of the Glaccountcontainer | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**FkiDepartmentID** | **int32** | The unique ID of the Department | 
**FkiPaymentmethodID** | **int32** | The unique ID of the Paymentmethod | 
**FkiElectronicfundstransferbankaccountID** | **int32** | The unique ID of the Electronicfundstransferbankaccount | 
**FkiElectronicfundstransferbankaccountIDDirectdebit** | **int32** | The unique ID of the Electronicfundstransferbankaccount | 
**FkiSendingmethodID** | **int32** | The unique ID of the Sendingmethod | 
**FkiTaxassignmentID** | **int32** | The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable| | 
**FkiAttendancestatusID** | **int32** | The unique ID of the Attendancestatus | 
**FkiAgentIDVariableexpensechargeto** | **int32** | The unique ID of the Agent. | 
**FkiBrokerIDVariableexpensechargeto** | **int32** | The unique ID of the Broker. | 
**FkiCustomerIDVariableexpensechargeto** | **int32** | The unique ID of the Customer. | 
**FkiGlaccountcontainerIDVariableexpensechargeto** | **int32** | The unique ID of the Glaccountcontainer | 
**FkiAgentIDSupplychargechargeto** | **int32** | The unique ID of the Agent. | 
**FkiBrokerIDSupplychargechargeto** | **int32** | The unique ID of the Broker. | 
**FkiCustomerIDSupplychargechargeto** | **int32** | The unique ID of the Customer. | 
**FkiGlaccountcontainerIDSupplychargechargeto** | **int32** | The unique ID of the Glaccountcontainer | 
**FkiInvoicealternatelogoID** | **int32** | The unique ID of the Invoicealternatelogo | 
**FkiSynchronizationlinkserverID** | **int32** | The unique ID of the Synchronizationlinkserver | 
**EfkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**EfksCustomerCode** | Pointer to **string** | The code of the Customer | [optional] 
**SCustomerCode** | **string** | The code of the Customer | 
**DCustomerFulltimeequivalent** | **string** | The fulltimeequivalent of the Customer | 
**ICustomerPhotocopiercode** | **int32** | The photocopiercode of the Customer | 
**ICustomerLongdistancecode** | **int32** | The longdistancecode of the Customer | 
**ICustomerTimewindowstart** | **int32** | The timewindowstart of the Customer | 
**ICustomerTimewindowend** | **int32** | The timewindowend of the Customer | 
**DCustomerMinimumchargeableinterests** | **string** | The minimumchargeableinterests of the Customer | 
**DtCustomerBirthdate** | **string** | The birthdate of the Customer | 
**DtCustomerTransfer** | **string** | The transfer of the Customer | 
**DtCustomerTransferappointment** | **string** | The transferappointment of the Customer | 
**DtCustomerTransfersurvey** | **string** | The transfersurvey of the Customer | 
**BCustomerIsactive** | **bool** | Whether the customer is active or not | 
**BCustomerVariableexpensefinanced** | **bool** | Whether if it&#39;s an variableexpensefinanced | 
**BCustomerVariableexpensefinancedtaxes** | **bool** | Whether if it&#39;s an variableexpensefinancedtaxes | 
**BCustomerSupplychargefinanced** | **bool** | Whether if it&#39;s an supplychargefinanced | 
**BCustomerSupplychargefinancedtaxes** | **bool** | Whether if it&#39;s an supplychargefinancedtaxes | 
**BCustomerAttendance** | **bool** | Whether if it&#39;s an attendance | 
**ECustomerType** | [**FieldECustomerType**](FieldECustomerType.md) |  | 
**ECustomerMarketingcorrespondence** | [**FieldECustomerMarketingcorrespondence**](FieldECustomerMarketingcorrespondence.md) |  | 
**BCustomerBlackcopycarbon** | **bool** | Whether if it&#39;s an blackcopycarbon | 
**BCustomerUnsubscribeinfo** | **bool** | Whether if it&#39;s an unsubscribeinfo | 
**TCustomerComment** | **string** | The comment of the Customer | 
**IMPORTID** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerResponse

`func NewCustomerResponse(pkiCustomerID int32, fkiCompanyID int32, fkiCustomergroupID int32, sCustomerName string, fkiContactinformationsID int32, fkiContactcontainerID int32, fkiImageID int32, fkiGlaccountcontainerID int32, fkiLanguageID int32, fkiDepartmentID int32, fkiPaymentmethodID int32, fkiElectronicfundstransferbankaccountID int32, fkiElectronicfundstransferbankaccountIDDirectdebit int32, fkiSendingmethodID int32, fkiTaxassignmentID int32, fkiAttendancestatusID int32, fkiAgentIDVariableexpensechargeto int32, fkiBrokerIDVariableexpensechargeto int32, fkiCustomerIDVariableexpensechargeto int32, fkiGlaccountcontainerIDVariableexpensechargeto int32, fkiAgentIDSupplychargechargeto int32, fkiBrokerIDSupplychargechargeto int32, fkiCustomerIDSupplychargechargeto int32, fkiGlaccountcontainerIDSupplychargechargeto int32, fkiInvoicealternatelogoID int32, fkiSynchronizationlinkserverID int32, sCustomerCode string, dCustomerFulltimeequivalent string, iCustomerPhotocopiercode int32, iCustomerLongdistancecode int32, iCustomerTimewindowstart int32, iCustomerTimewindowend int32, dCustomerMinimumchargeableinterests string, dtCustomerBirthdate string, dtCustomerTransfer string, dtCustomerTransferappointment string, dtCustomerTransfersurvey string, bCustomerIsactive bool, bCustomerVariableexpensefinanced bool, bCustomerVariableexpensefinancedtaxes bool, bCustomerSupplychargefinanced bool, bCustomerSupplychargefinancedtaxes bool, bCustomerAttendance bool, eCustomerType FieldECustomerType, eCustomerMarketingcorrespondence FieldECustomerMarketingcorrespondence, bCustomerBlackcopycarbon bool, bCustomerUnsubscribeinfo bool, tCustomerComment string, ) *CustomerResponse`

NewCustomerResponse instantiates a new CustomerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResponseWithDefaults

`func NewCustomerResponseWithDefaults() *CustomerResponse`

NewCustomerResponseWithDefaults instantiates a new CustomerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCustomerID

`func (o *CustomerResponse) GetPkiCustomerID() int32`

GetPkiCustomerID returns the PkiCustomerID field if non-nil, zero value otherwise.

### GetPkiCustomerIDOk

`func (o *CustomerResponse) GetPkiCustomerIDOk() (*int32, bool)`

GetPkiCustomerIDOk returns a tuple with the PkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCustomerID

`func (o *CustomerResponse) SetPkiCustomerID(v int32)`

SetPkiCustomerID sets PkiCustomerID field to given value.


### GetFkiCompanyID

`func (o *CustomerResponse) GetFkiCompanyID() int32`

GetFkiCompanyID returns the FkiCompanyID field if non-nil, zero value otherwise.

### GetFkiCompanyIDOk

`func (o *CustomerResponse) GetFkiCompanyIDOk() (*int32, bool)`

GetFkiCompanyIDOk returns a tuple with the FkiCompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCompanyID

`func (o *CustomerResponse) SetFkiCompanyID(v int32)`

SetFkiCompanyID sets FkiCompanyID field to given value.


### GetFkiCustomergroupID

`func (o *CustomerResponse) GetFkiCustomergroupID() int32`

GetFkiCustomergroupID returns the FkiCustomergroupID field if non-nil, zero value otherwise.

### GetFkiCustomergroupIDOk

`func (o *CustomerResponse) GetFkiCustomergroupIDOk() (*int32, bool)`

GetFkiCustomergroupIDOk returns a tuple with the FkiCustomergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomergroupID

`func (o *CustomerResponse) SetFkiCustomergroupID(v int32)`

SetFkiCustomergroupID sets FkiCustomergroupID field to given value.


### GetSCustomerName

`func (o *CustomerResponse) GetSCustomerName() string`

GetSCustomerName returns the SCustomerName field if non-nil, zero value otherwise.

### GetSCustomerNameOk

`func (o *CustomerResponse) GetSCustomerNameOk() (*string, bool)`

GetSCustomerNameOk returns a tuple with the SCustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCustomerName

`func (o *CustomerResponse) SetSCustomerName(v string)`

SetSCustomerName sets SCustomerName field to given value.


### GetFkiContactinformationsID

`func (o *CustomerResponse) GetFkiContactinformationsID() int32`

GetFkiContactinformationsID returns the FkiContactinformationsID field if non-nil, zero value otherwise.

### GetFkiContactinformationsIDOk

`func (o *CustomerResponse) GetFkiContactinformationsIDOk() (*int32, bool)`

GetFkiContactinformationsIDOk returns a tuple with the FkiContactinformationsID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactinformationsID

`func (o *CustomerResponse) SetFkiContactinformationsID(v int32)`

SetFkiContactinformationsID sets FkiContactinformationsID field to given value.


### GetFkiContactcontainerID

`func (o *CustomerResponse) GetFkiContactcontainerID() int32`

GetFkiContactcontainerID returns the FkiContactcontainerID field if non-nil, zero value otherwise.

### GetFkiContactcontainerIDOk

`func (o *CustomerResponse) GetFkiContactcontainerIDOk() (*int32, bool)`

GetFkiContactcontainerIDOk returns a tuple with the FkiContactcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactcontainerID

`func (o *CustomerResponse) SetFkiContactcontainerID(v int32)`

SetFkiContactcontainerID sets FkiContactcontainerID field to given value.


### GetFkiImageID

`func (o *CustomerResponse) GetFkiImageID() int32`

GetFkiImageID returns the FkiImageID field if non-nil, zero value otherwise.

### GetFkiImageIDOk

`func (o *CustomerResponse) GetFkiImageIDOk() (*int32, bool)`

GetFkiImageIDOk returns a tuple with the FkiImageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiImageID

`func (o *CustomerResponse) SetFkiImageID(v int32)`

SetFkiImageID sets FkiImageID field to given value.


### GetFkiGlaccountcontainerID

`func (o *CustomerResponse) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *CustomerResponse) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *CustomerResponse) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.


### GetFkiLanguageID

`func (o *CustomerResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *CustomerResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *CustomerResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiDepartmentID

`func (o *CustomerResponse) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *CustomerResponse) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *CustomerResponse) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### GetFkiPaymentmethodID

`func (o *CustomerResponse) GetFkiPaymentmethodID() int32`

GetFkiPaymentmethodID returns the FkiPaymentmethodID field if non-nil, zero value otherwise.

### GetFkiPaymentmethodIDOk

`func (o *CustomerResponse) GetFkiPaymentmethodIDOk() (*int32, bool)`

GetFkiPaymentmethodIDOk returns a tuple with the FkiPaymentmethodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentmethodID

`func (o *CustomerResponse) SetFkiPaymentmethodID(v int32)`

SetFkiPaymentmethodID sets FkiPaymentmethodID field to given value.


### GetFkiElectronicfundstransferbankaccountID

`func (o *CustomerResponse) GetFkiElectronicfundstransferbankaccountID() int32`

GetFkiElectronicfundstransferbankaccountID returns the FkiElectronicfundstransferbankaccountID field if non-nil, zero value otherwise.

### GetFkiElectronicfundstransferbankaccountIDOk

`func (o *CustomerResponse) GetFkiElectronicfundstransferbankaccountIDOk() (*int32, bool)`

GetFkiElectronicfundstransferbankaccountIDOk returns a tuple with the FkiElectronicfundstransferbankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiElectronicfundstransferbankaccountID

`func (o *CustomerResponse) SetFkiElectronicfundstransferbankaccountID(v int32)`

SetFkiElectronicfundstransferbankaccountID sets FkiElectronicfundstransferbankaccountID field to given value.


### GetFkiElectronicfundstransferbankaccountIDDirectdebit

`func (o *CustomerResponse) GetFkiElectronicfundstransferbankaccountIDDirectdebit() int32`

GetFkiElectronicfundstransferbankaccountIDDirectdebit returns the FkiElectronicfundstransferbankaccountIDDirectdebit field if non-nil, zero value otherwise.

### GetFkiElectronicfundstransferbankaccountIDDirectdebitOk

`func (o *CustomerResponse) GetFkiElectronicfundstransferbankaccountIDDirectdebitOk() (*int32, bool)`

GetFkiElectronicfundstransferbankaccountIDDirectdebitOk returns a tuple with the FkiElectronicfundstransferbankaccountIDDirectdebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiElectronicfundstransferbankaccountIDDirectdebit

`func (o *CustomerResponse) SetFkiElectronicfundstransferbankaccountIDDirectdebit(v int32)`

SetFkiElectronicfundstransferbankaccountIDDirectdebit sets FkiElectronicfundstransferbankaccountIDDirectdebit field to given value.


### GetFkiSendingmethodID

`func (o *CustomerResponse) GetFkiSendingmethodID() int32`

GetFkiSendingmethodID returns the FkiSendingmethodID field if non-nil, zero value otherwise.

### GetFkiSendingmethodIDOk

`func (o *CustomerResponse) GetFkiSendingmethodIDOk() (*int32, bool)`

GetFkiSendingmethodIDOk returns a tuple with the FkiSendingmethodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSendingmethodID

`func (o *CustomerResponse) SetFkiSendingmethodID(v int32)`

SetFkiSendingmethodID sets FkiSendingmethodID field to given value.


### GetFkiTaxassignmentID

`func (o *CustomerResponse) GetFkiTaxassignmentID() int32`

GetFkiTaxassignmentID returns the FkiTaxassignmentID field if non-nil, zero value otherwise.

### GetFkiTaxassignmentIDOk

`func (o *CustomerResponse) GetFkiTaxassignmentIDOk() (*int32, bool)`

GetFkiTaxassignmentIDOk returns a tuple with the FkiTaxassignmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTaxassignmentID

`func (o *CustomerResponse) SetFkiTaxassignmentID(v int32)`

SetFkiTaxassignmentID sets FkiTaxassignmentID field to given value.


### GetFkiAttendancestatusID

`func (o *CustomerResponse) GetFkiAttendancestatusID() int32`

GetFkiAttendancestatusID returns the FkiAttendancestatusID field if non-nil, zero value otherwise.

### GetFkiAttendancestatusIDOk

`func (o *CustomerResponse) GetFkiAttendancestatusIDOk() (*int32, bool)`

GetFkiAttendancestatusIDOk returns a tuple with the FkiAttendancestatusID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAttendancestatusID

`func (o *CustomerResponse) SetFkiAttendancestatusID(v int32)`

SetFkiAttendancestatusID sets FkiAttendancestatusID field to given value.


### GetFkiAgentIDVariableexpensechargeto

`func (o *CustomerResponse) GetFkiAgentIDVariableexpensechargeto() int32`

GetFkiAgentIDVariableexpensechargeto returns the FkiAgentIDVariableexpensechargeto field if non-nil, zero value otherwise.

### GetFkiAgentIDVariableexpensechargetoOk

`func (o *CustomerResponse) GetFkiAgentIDVariableexpensechargetoOk() (*int32, bool)`

GetFkiAgentIDVariableexpensechargetoOk returns a tuple with the FkiAgentIDVariableexpensechargeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentIDVariableexpensechargeto

`func (o *CustomerResponse) SetFkiAgentIDVariableexpensechargeto(v int32)`

SetFkiAgentIDVariableexpensechargeto sets FkiAgentIDVariableexpensechargeto field to given value.


### GetFkiBrokerIDVariableexpensechargeto

`func (o *CustomerResponse) GetFkiBrokerIDVariableexpensechargeto() int32`

GetFkiBrokerIDVariableexpensechargeto returns the FkiBrokerIDVariableexpensechargeto field if non-nil, zero value otherwise.

### GetFkiBrokerIDVariableexpensechargetoOk

`func (o *CustomerResponse) GetFkiBrokerIDVariableexpensechargetoOk() (*int32, bool)`

GetFkiBrokerIDVariableexpensechargetoOk returns a tuple with the FkiBrokerIDVariableexpensechargeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerIDVariableexpensechargeto

`func (o *CustomerResponse) SetFkiBrokerIDVariableexpensechargeto(v int32)`

SetFkiBrokerIDVariableexpensechargeto sets FkiBrokerIDVariableexpensechargeto field to given value.


### GetFkiCustomerIDVariableexpensechargeto

`func (o *CustomerResponse) GetFkiCustomerIDVariableexpensechargeto() int32`

GetFkiCustomerIDVariableexpensechargeto returns the FkiCustomerIDVariableexpensechargeto field if non-nil, zero value otherwise.

### GetFkiCustomerIDVariableexpensechargetoOk

`func (o *CustomerResponse) GetFkiCustomerIDVariableexpensechargetoOk() (*int32, bool)`

GetFkiCustomerIDVariableexpensechargetoOk returns a tuple with the FkiCustomerIDVariableexpensechargeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerIDVariableexpensechargeto

`func (o *CustomerResponse) SetFkiCustomerIDVariableexpensechargeto(v int32)`

SetFkiCustomerIDVariableexpensechargeto sets FkiCustomerIDVariableexpensechargeto field to given value.


### GetFkiGlaccountcontainerIDVariableexpensechargeto

`func (o *CustomerResponse) GetFkiGlaccountcontainerIDVariableexpensechargeto() int32`

GetFkiGlaccountcontainerIDVariableexpensechargeto returns the FkiGlaccountcontainerIDVariableexpensechargeto field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDVariableexpensechargetoOk

`func (o *CustomerResponse) GetFkiGlaccountcontainerIDVariableexpensechargetoOk() (*int32, bool)`

GetFkiGlaccountcontainerIDVariableexpensechargetoOk returns a tuple with the FkiGlaccountcontainerIDVariableexpensechargeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerIDVariableexpensechargeto

`func (o *CustomerResponse) SetFkiGlaccountcontainerIDVariableexpensechargeto(v int32)`

SetFkiGlaccountcontainerIDVariableexpensechargeto sets FkiGlaccountcontainerIDVariableexpensechargeto field to given value.


### GetFkiAgentIDSupplychargechargeto

`func (o *CustomerResponse) GetFkiAgentIDSupplychargechargeto() int32`

GetFkiAgentIDSupplychargechargeto returns the FkiAgentIDSupplychargechargeto field if non-nil, zero value otherwise.

### GetFkiAgentIDSupplychargechargetoOk

`func (o *CustomerResponse) GetFkiAgentIDSupplychargechargetoOk() (*int32, bool)`

GetFkiAgentIDSupplychargechargetoOk returns a tuple with the FkiAgentIDSupplychargechargeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentIDSupplychargechargeto

`func (o *CustomerResponse) SetFkiAgentIDSupplychargechargeto(v int32)`

SetFkiAgentIDSupplychargechargeto sets FkiAgentIDSupplychargechargeto field to given value.


### GetFkiBrokerIDSupplychargechargeto

`func (o *CustomerResponse) GetFkiBrokerIDSupplychargechargeto() int32`

GetFkiBrokerIDSupplychargechargeto returns the FkiBrokerIDSupplychargechargeto field if non-nil, zero value otherwise.

### GetFkiBrokerIDSupplychargechargetoOk

`func (o *CustomerResponse) GetFkiBrokerIDSupplychargechargetoOk() (*int32, bool)`

GetFkiBrokerIDSupplychargechargetoOk returns a tuple with the FkiBrokerIDSupplychargechargeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerIDSupplychargechargeto

`func (o *CustomerResponse) SetFkiBrokerIDSupplychargechargeto(v int32)`

SetFkiBrokerIDSupplychargechargeto sets FkiBrokerIDSupplychargechargeto field to given value.


### GetFkiCustomerIDSupplychargechargeto

`func (o *CustomerResponse) GetFkiCustomerIDSupplychargechargeto() int32`

GetFkiCustomerIDSupplychargechargeto returns the FkiCustomerIDSupplychargechargeto field if non-nil, zero value otherwise.

### GetFkiCustomerIDSupplychargechargetoOk

`func (o *CustomerResponse) GetFkiCustomerIDSupplychargechargetoOk() (*int32, bool)`

GetFkiCustomerIDSupplychargechargetoOk returns a tuple with the FkiCustomerIDSupplychargechargeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerIDSupplychargechargeto

`func (o *CustomerResponse) SetFkiCustomerIDSupplychargechargeto(v int32)`

SetFkiCustomerIDSupplychargechargeto sets FkiCustomerIDSupplychargechargeto field to given value.


### GetFkiGlaccountcontainerIDSupplychargechargeto

`func (o *CustomerResponse) GetFkiGlaccountcontainerIDSupplychargechargeto() int32`

GetFkiGlaccountcontainerIDSupplychargechargeto returns the FkiGlaccountcontainerIDSupplychargechargeto field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDSupplychargechargetoOk

`func (o *CustomerResponse) GetFkiGlaccountcontainerIDSupplychargechargetoOk() (*int32, bool)`

GetFkiGlaccountcontainerIDSupplychargechargetoOk returns a tuple with the FkiGlaccountcontainerIDSupplychargechargeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerIDSupplychargechargeto

`func (o *CustomerResponse) SetFkiGlaccountcontainerIDSupplychargechargeto(v int32)`

SetFkiGlaccountcontainerIDSupplychargechargeto sets FkiGlaccountcontainerIDSupplychargechargeto field to given value.


### GetFkiInvoicealternatelogoID

`func (o *CustomerResponse) GetFkiInvoicealternatelogoID() int32`

GetFkiInvoicealternatelogoID returns the FkiInvoicealternatelogoID field if non-nil, zero value otherwise.

### GetFkiInvoicealternatelogoIDOk

`func (o *CustomerResponse) GetFkiInvoicealternatelogoIDOk() (*int32, bool)`

GetFkiInvoicealternatelogoIDOk returns a tuple with the FkiInvoicealternatelogoID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInvoicealternatelogoID

`func (o *CustomerResponse) SetFkiInvoicealternatelogoID(v int32)`

SetFkiInvoicealternatelogoID sets FkiInvoicealternatelogoID field to given value.


### GetFkiSynchronizationlinkserverID

`func (o *CustomerResponse) GetFkiSynchronizationlinkserverID() int32`

GetFkiSynchronizationlinkserverID returns the FkiSynchronizationlinkserverID field if non-nil, zero value otherwise.

### GetFkiSynchronizationlinkserverIDOk

`func (o *CustomerResponse) GetFkiSynchronizationlinkserverIDOk() (*int32, bool)`

GetFkiSynchronizationlinkserverIDOk returns a tuple with the FkiSynchronizationlinkserverID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSynchronizationlinkserverID

`func (o *CustomerResponse) SetFkiSynchronizationlinkserverID(v int32)`

SetFkiSynchronizationlinkserverID sets FkiSynchronizationlinkserverID field to given value.


### GetEfkiUserID

`func (o *CustomerResponse) GetEfkiUserID() int32`

GetEfkiUserID returns the EfkiUserID field if non-nil, zero value otherwise.

### GetEfkiUserIDOk

`func (o *CustomerResponse) GetEfkiUserIDOk() (*int32, bool)`

GetEfkiUserIDOk returns a tuple with the EfkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfkiUserID

`func (o *CustomerResponse) SetEfkiUserID(v int32)`

SetEfkiUserID sets EfkiUserID field to given value.

### HasEfkiUserID

`func (o *CustomerResponse) HasEfkiUserID() bool`

HasEfkiUserID returns a boolean if a field has been set.

### GetEfksCustomerCode

`func (o *CustomerResponse) GetEfksCustomerCode() string`

GetEfksCustomerCode returns the EfksCustomerCode field if non-nil, zero value otherwise.

### GetEfksCustomerCodeOk

`func (o *CustomerResponse) GetEfksCustomerCodeOk() (*string, bool)`

GetEfksCustomerCodeOk returns a tuple with the EfksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfksCustomerCode

`func (o *CustomerResponse) SetEfksCustomerCode(v string)`

SetEfksCustomerCode sets EfksCustomerCode field to given value.

### HasEfksCustomerCode

`func (o *CustomerResponse) HasEfksCustomerCode() bool`

HasEfksCustomerCode returns a boolean if a field has been set.

### GetSCustomerCode

`func (o *CustomerResponse) GetSCustomerCode() string`

GetSCustomerCode returns the SCustomerCode field if non-nil, zero value otherwise.

### GetSCustomerCodeOk

`func (o *CustomerResponse) GetSCustomerCodeOk() (*string, bool)`

GetSCustomerCodeOk returns a tuple with the SCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCustomerCode

`func (o *CustomerResponse) SetSCustomerCode(v string)`

SetSCustomerCode sets SCustomerCode field to given value.


### GetDCustomerFulltimeequivalent

`func (o *CustomerResponse) GetDCustomerFulltimeequivalent() string`

GetDCustomerFulltimeequivalent returns the DCustomerFulltimeequivalent field if non-nil, zero value otherwise.

### GetDCustomerFulltimeequivalentOk

`func (o *CustomerResponse) GetDCustomerFulltimeequivalentOk() (*string, bool)`

GetDCustomerFulltimeequivalentOk returns a tuple with the DCustomerFulltimeequivalent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDCustomerFulltimeequivalent

`func (o *CustomerResponse) SetDCustomerFulltimeequivalent(v string)`

SetDCustomerFulltimeequivalent sets DCustomerFulltimeequivalent field to given value.


### GetICustomerPhotocopiercode

`func (o *CustomerResponse) GetICustomerPhotocopiercode() int32`

GetICustomerPhotocopiercode returns the ICustomerPhotocopiercode field if non-nil, zero value otherwise.

### GetICustomerPhotocopiercodeOk

`func (o *CustomerResponse) GetICustomerPhotocopiercodeOk() (*int32, bool)`

GetICustomerPhotocopiercodeOk returns a tuple with the ICustomerPhotocopiercode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICustomerPhotocopiercode

`func (o *CustomerResponse) SetICustomerPhotocopiercode(v int32)`

SetICustomerPhotocopiercode sets ICustomerPhotocopiercode field to given value.


### GetICustomerLongdistancecode

`func (o *CustomerResponse) GetICustomerLongdistancecode() int32`

GetICustomerLongdistancecode returns the ICustomerLongdistancecode field if non-nil, zero value otherwise.

### GetICustomerLongdistancecodeOk

`func (o *CustomerResponse) GetICustomerLongdistancecodeOk() (*int32, bool)`

GetICustomerLongdistancecodeOk returns a tuple with the ICustomerLongdistancecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICustomerLongdistancecode

`func (o *CustomerResponse) SetICustomerLongdistancecode(v int32)`

SetICustomerLongdistancecode sets ICustomerLongdistancecode field to given value.


### GetICustomerTimewindowstart

`func (o *CustomerResponse) GetICustomerTimewindowstart() int32`

GetICustomerTimewindowstart returns the ICustomerTimewindowstart field if non-nil, zero value otherwise.

### GetICustomerTimewindowstartOk

`func (o *CustomerResponse) GetICustomerTimewindowstartOk() (*int32, bool)`

GetICustomerTimewindowstartOk returns a tuple with the ICustomerTimewindowstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICustomerTimewindowstart

`func (o *CustomerResponse) SetICustomerTimewindowstart(v int32)`

SetICustomerTimewindowstart sets ICustomerTimewindowstart field to given value.


### GetICustomerTimewindowend

`func (o *CustomerResponse) GetICustomerTimewindowend() int32`

GetICustomerTimewindowend returns the ICustomerTimewindowend field if non-nil, zero value otherwise.

### GetICustomerTimewindowendOk

`func (o *CustomerResponse) GetICustomerTimewindowendOk() (*int32, bool)`

GetICustomerTimewindowendOk returns a tuple with the ICustomerTimewindowend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICustomerTimewindowend

`func (o *CustomerResponse) SetICustomerTimewindowend(v int32)`

SetICustomerTimewindowend sets ICustomerTimewindowend field to given value.


### GetDCustomerMinimumchargeableinterests

`func (o *CustomerResponse) GetDCustomerMinimumchargeableinterests() string`

GetDCustomerMinimumchargeableinterests returns the DCustomerMinimumchargeableinterests field if non-nil, zero value otherwise.

### GetDCustomerMinimumchargeableinterestsOk

`func (o *CustomerResponse) GetDCustomerMinimumchargeableinterestsOk() (*string, bool)`

GetDCustomerMinimumchargeableinterestsOk returns a tuple with the DCustomerMinimumchargeableinterests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDCustomerMinimumchargeableinterests

`func (o *CustomerResponse) SetDCustomerMinimumchargeableinterests(v string)`

SetDCustomerMinimumchargeableinterests sets DCustomerMinimumchargeableinterests field to given value.


### GetDtCustomerBirthdate

`func (o *CustomerResponse) GetDtCustomerBirthdate() string`

GetDtCustomerBirthdate returns the DtCustomerBirthdate field if non-nil, zero value otherwise.

### GetDtCustomerBirthdateOk

`func (o *CustomerResponse) GetDtCustomerBirthdateOk() (*string, bool)`

GetDtCustomerBirthdateOk returns a tuple with the DtCustomerBirthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCustomerBirthdate

`func (o *CustomerResponse) SetDtCustomerBirthdate(v string)`

SetDtCustomerBirthdate sets DtCustomerBirthdate field to given value.


### GetDtCustomerTransfer

`func (o *CustomerResponse) GetDtCustomerTransfer() string`

GetDtCustomerTransfer returns the DtCustomerTransfer field if non-nil, zero value otherwise.

### GetDtCustomerTransferOk

`func (o *CustomerResponse) GetDtCustomerTransferOk() (*string, bool)`

GetDtCustomerTransferOk returns a tuple with the DtCustomerTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCustomerTransfer

`func (o *CustomerResponse) SetDtCustomerTransfer(v string)`

SetDtCustomerTransfer sets DtCustomerTransfer field to given value.


### GetDtCustomerTransferappointment

`func (o *CustomerResponse) GetDtCustomerTransferappointment() string`

GetDtCustomerTransferappointment returns the DtCustomerTransferappointment field if non-nil, zero value otherwise.

### GetDtCustomerTransferappointmentOk

`func (o *CustomerResponse) GetDtCustomerTransferappointmentOk() (*string, bool)`

GetDtCustomerTransferappointmentOk returns a tuple with the DtCustomerTransferappointment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCustomerTransferappointment

`func (o *CustomerResponse) SetDtCustomerTransferappointment(v string)`

SetDtCustomerTransferappointment sets DtCustomerTransferappointment field to given value.


### GetDtCustomerTransfersurvey

`func (o *CustomerResponse) GetDtCustomerTransfersurvey() string`

GetDtCustomerTransfersurvey returns the DtCustomerTransfersurvey field if non-nil, zero value otherwise.

### GetDtCustomerTransfersurveyOk

`func (o *CustomerResponse) GetDtCustomerTransfersurveyOk() (*string, bool)`

GetDtCustomerTransfersurveyOk returns a tuple with the DtCustomerTransfersurvey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCustomerTransfersurvey

`func (o *CustomerResponse) SetDtCustomerTransfersurvey(v string)`

SetDtCustomerTransfersurvey sets DtCustomerTransfersurvey field to given value.


### GetBCustomerIsactive

`func (o *CustomerResponse) GetBCustomerIsactive() bool`

GetBCustomerIsactive returns the BCustomerIsactive field if non-nil, zero value otherwise.

### GetBCustomerIsactiveOk

`func (o *CustomerResponse) GetBCustomerIsactiveOk() (*bool, bool)`

GetBCustomerIsactiveOk returns a tuple with the BCustomerIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerIsactive

`func (o *CustomerResponse) SetBCustomerIsactive(v bool)`

SetBCustomerIsactive sets BCustomerIsactive field to given value.


### GetBCustomerVariableexpensefinanced

`func (o *CustomerResponse) GetBCustomerVariableexpensefinanced() bool`

GetBCustomerVariableexpensefinanced returns the BCustomerVariableexpensefinanced field if non-nil, zero value otherwise.

### GetBCustomerVariableexpensefinancedOk

`func (o *CustomerResponse) GetBCustomerVariableexpensefinancedOk() (*bool, bool)`

GetBCustomerVariableexpensefinancedOk returns a tuple with the BCustomerVariableexpensefinanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerVariableexpensefinanced

`func (o *CustomerResponse) SetBCustomerVariableexpensefinanced(v bool)`

SetBCustomerVariableexpensefinanced sets BCustomerVariableexpensefinanced field to given value.


### GetBCustomerVariableexpensefinancedtaxes

`func (o *CustomerResponse) GetBCustomerVariableexpensefinancedtaxes() bool`

GetBCustomerVariableexpensefinancedtaxes returns the BCustomerVariableexpensefinancedtaxes field if non-nil, zero value otherwise.

### GetBCustomerVariableexpensefinancedtaxesOk

`func (o *CustomerResponse) GetBCustomerVariableexpensefinancedtaxesOk() (*bool, bool)`

GetBCustomerVariableexpensefinancedtaxesOk returns a tuple with the BCustomerVariableexpensefinancedtaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerVariableexpensefinancedtaxes

`func (o *CustomerResponse) SetBCustomerVariableexpensefinancedtaxes(v bool)`

SetBCustomerVariableexpensefinancedtaxes sets BCustomerVariableexpensefinancedtaxes field to given value.


### GetBCustomerSupplychargefinanced

`func (o *CustomerResponse) GetBCustomerSupplychargefinanced() bool`

GetBCustomerSupplychargefinanced returns the BCustomerSupplychargefinanced field if non-nil, zero value otherwise.

### GetBCustomerSupplychargefinancedOk

`func (o *CustomerResponse) GetBCustomerSupplychargefinancedOk() (*bool, bool)`

GetBCustomerSupplychargefinancedOk returns a tuple with the BCustomerSupplychargefinanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerSupplychargefinanced

`func (o *CustomerResponse) SetBCustomerSupplychargefinanced(v bool)`

SetBCustomerSupplychargefinanced sets BCustomerSupplychargefinanced field to given value.


### GetBCustomerSupplychargefinancedtaxes

`func (o *CustomerResponse) GetBCustomerSupplychargefinancedtaxes() bool`

GetBCustomerSupplychargefinancedtaxes returns the BCustomerSupplychargefinancedtaxes field if non-nil, zero value otherwise.

### GetBCustomerSupplychargefinancedtaxesOk

`func (o *CustomerResponse) GetBCustomerSupplychargefinancedtaxesOk() (*bool, bool)`

GetBCustomerSupplychargefinancedtaxesOk returns a tuple with the BCustomerSupplychargefinancedtaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerSupplychargefinancedtaxes

`func (o *CustomerResponse) SetBCustomerSupplychargefinancedtaxes(v bool)`

SetBCustomerSupplychargefinancedtaxes sets BCustomerSupplychargefinancedtaxes field to given value.


### GetBCustomerAttendance

`func (o *CustomerResponse) GetBCustomerAttendance() bool`

GetBCustomerAttendance returns the BCustomerAttendance field if non-nil, zero value otherwise.

### GetBCustomerAttendanceOk

`func (o *CustomerResponse) GetBCustomerAttendanceOk() (*bool, bool)`

GetBCustomerAttendanceOk returns a tuple with the BCustomerAttendance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerAttendance

`func (o *CustomerResponse) SetBCustomerAttendance(v bool)`

SetBCustomerAttendance sets BCustomerAttendance field to given value.


### GetECustomerType

`func (o *CustomerResponse) GetECustomerType() FieldECustomerType`

GetECustomerType returns the ECustomerType field if non-nil, zero value otherwise.

### GetECustomerTypeOk

`func (o *CustomerResponse) GetECustomerTypeOk() (*FieldECustomerType, bool)`

GetECustomerTypeOk returns a tuple with the ECustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECustomerType

`func (o *CustomerResponse) SetECustomerType(v FieldECustomerType)`

SetECustomerType sets ECustomerType field to given value.


### GetECustomerMarketingcorrespondence

`func (o *CustomerResponse) GetECustomerMarketingcorrespondence() FieldECustomerMarketingcorrespondence`

GetECustomerMarketingcorrespondence returns the ECustomerMarketingcorrespondence field if non-nil, zero value otherwise.

### GetECustomerMarketingcorrespondenceOk

`func (o *CustomerResponse) GetECustomerMarketingcorrespondenceOk() (*FieldECustomerMarketingcorrespondence, bool)`

GetECustomerMarketingcorrespondenceOk returns a tuple with the ECustomerMarketingcorrespondence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECustomerMarketingcorrespondence

`func (o *CustomerResponse) SetECustomerMarketingcorrespondence(v FieldECustomerMarketingcorrespondence)`

SetECustomerMarketingcorrespondence sets ECustomerMarketingcorrespondence field to given value.


### GetBCustomerBlackcopycarbon

`func (o *CustomerResponse) GetBCustomerBlackcopycarbon() bool`

GetBCustomerBlackcopycarbon returns the BCustomerBlackcopycarbon field if non-nil, zero value otherwise.

### GetBCustomerBlackcopycarbonOk

`func (o *CustomerResponse) GetBCustomerBlackcopycarbonOk() (*bool, bool)`

GetBCustomerBlackcopycarbonOk returns a tuple with the BCustomerBlackcopycarbon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerBlackcopycarbon

`func (o *CustomerResponse) SetBCustomerBlackcopycarbon(v bool)`

SetBCustomerBlackcopycarbon sets BCustomerBlackcopycarbon field to given value.


### GetBCustomerUnsubscribeinfo

`func (o *CustomerResponse) GetBCustomerUnsubscribeinfo() bool`

GetBCustomerUnsubscribeinfo returns the BCustomerUnsubscribeinfo field if non-nil, zero value otherwise.

### GetBCustomerUnsubscribeinfoOk

`func (o *CustomerResponse) GetBCustomerUnsubscribeinfoOk() (*bool, bool)`

GetBCustomerUnsubscribeinfoOk returns a tuple with the BCustomerUnsubscribeinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerUnsubscribeinfo

`func (o *CustomerResponse) SetBCustomerUnsubscribeinfo(v bool)`

SetBCustomerUnsubscribeinfo sets BCustomerUnsubscribeinfo field to given value.


### GetTCustomerComment

`func (o *CustomerResponse) GetTCustomerComment() string`

GetTCustomerComment returns the TCustomerComment field if non-nil, zero value otherwise.

### GetTCustomerCommentOk

`func (o *CustomerResponse) GetTCustomerCommentOk() (*string, bool)`

GetTCustomerCommentOk returns a tuple with the TCustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTCustomerComment

`func (o *CustomerResponse) SetTCustomerComment(v string)`

SetTCustomerComment sets TCustomerComment field to given value.


### GetIMPORTID

`func (o *CustomerResponse) GetIMPORTID() string`

GetIMPORTID returns the IMPORTID field if non-nil, zero value otherwise.

### GetIMPORTIDOk

`func (o *CustomerResponse) GetIMPORTIDOk() (*string, bool)`

GetIMPORTIDOk returns a tuple with the IMPORTID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTID

`func (o *CustomerResponse) SetIMPORTID(v string)`

SetIMPORTID sets IMPORTID field to given value.

### HasIMPORTID

`func (o *CustomerResponse) HasIMPORTID() bool`

HasIMPORTID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


