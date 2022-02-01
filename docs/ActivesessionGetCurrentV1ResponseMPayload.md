# ActivesessionGetCurrentV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SCustomerCode** | **string** | The customer code specific to the client in which the API request is being made | 
**EActivesessionSessiontype** | **string** | The type of session used for the API request call | 
**EActivesessionWeekdaystart** | [**FieldEActivesessionWeekdaystart**](FieldEActivesessionWeekdaystart.md) |  | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SCompanyNameX** | **string** | The name of the active Company in the current language | 
**SDepartmentNameX** | **string** | The name of the active Department in the current language | 
**ARegisteredModules** | **[]string** | An Array of Registered modules.  These are the modules that are Licensed to be used by the User or the API Key. | 
**APermissions** | **[]int32** | An array of permissions granted to the user or api key | 
**FkiUserID** | **int32** | The unique ID of the User | 
**FkiApikeyID** | **int32** | The unique ID of the Apikey | 

## Methods

### NewActivesessionGetCurrentV1ResponseMPayload

`func NewActivesessionGetCurrentV1ResponseMPayload(sCustomerCode string, eActivesessionSessiontype string, eActivesessionWeekdaystart FieldEActivesessionWeekdaystart, fkiLanguageID int32, sCompanyNameX string, sDepartmentNameX string, aRegisteredModules []string, aPermissions []int32, fkiUserID int32, fkiApikeyID int32, ) *ActivesessionGetCurrentV1ResponseMPayload`

NewActivesessionGetCurrentV1ResponseMPayload instantiates a new ActivesessionGetCurrentV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionGetCurrentV1ResponseMPayloadWithDefaults

`func NewActivesessionGetCurrentV1ResponseMPayloadWithDefaults() *ActivesessionGetCurrentV1ResponseMPayload`

NewActivesessionGetCurrentV1ResponseMPayloadWithDefaults instantiates a new ActivesessionGetCurrentV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSCustomerCode

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCustomerCode() string`

GetSCustomerCode returns the SCustomerCode field if non-nil, zero value otherwise.

### GetSCustomerCodeOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCustomerCodeOk() (*string, bool)`

GetSCustomerCodeOk returns a tuple with the SCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCustomerCode

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetSCustomerCode(v string)`

SetSCustomerCode sets SCustomerCode field to given value.


### GetEActivesessionSessiontype

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionSessiontype() string`

GetEActivesessionSessiontype returns the EActivesessionSessiontype field if non-nil, zero value otherwise.

### GetEActivesessionSessiontypeOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionSessiontypeOk() (*string, bool)`

GetEActivesessionSessiontypeOk returns a tuple with the EActivesessionSessiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionSessiontype

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionSessiontype(v string)`

SetEActivesessionSessiontype sets EActivesessionSessiontype field to given value.


### GetEActivesessionWeekdaystart

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionWeekdaystart() FieldEActivesessionWeekdaystart`

GetEActivesessionWeekdaystart returns the EActivesessionWeekdaystart field if non-nil, zero value otherwise.

### GetEActivesessionWeekdaystartOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionWeekdaystartOk() (*FieldEActivesessionWeekdaystart, bool)`

GetEActivesessionWeekdaystartOk returns a tuple with the EActivesessionWeekdaystart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionWeekdaystart

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionWeekdaystart(v FieldEActivesessionWeekdaystart)`

SetEActivesessionWeekdaystart sets EActivesessionWeekdaystart field to given value.


### GetFkiLanguageID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSCompanyNameX

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.


### GetSDepartmentNameX

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.


### GetARegisteredModules

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetARegisteredModules() []string`

GetARegisteredModules returns the ARegisteredModules field if non-nil, zero value otherwise.

### GetARegisteredModulesOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetARegisteredModulesOk() (*[]string, bool)`

GetARegisteredModulesOk returns a tuple with the ARegisteredModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetARegisteredModules

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetARegisteredModules(v []string)`

SetARegisteredModules sets ARegisteredModules field to given value.


### GetAPermissions

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetAPermissions() []int32`

GetAPermissions returns the APermissions field if non-nil, zero value otherwise.

### GetAPermissionsOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetAPermissionsOk() (*[]int32, bool)`

GetAPermissionsOk returns a tuple with the APermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPermissions

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetAPermissions(v []int32)`

SetAPermissions sets APermissions field to given value.


### GetFkiUserID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetFkiApikeyID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiApikeyID() int32`

GetFkiApikeyID returns the FkiApikeyID field if non-nil, zero value otherwise.

### GetFkiApikeyIDOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiApikeyIDOk() (*int32, bool)`

GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetFkiApikeyID(v int32)`

SetFkiApikeyID sets FkiApikeyID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


