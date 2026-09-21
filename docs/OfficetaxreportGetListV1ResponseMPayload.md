# OfficetaxreportGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjOfficetaxreport** | [**[]OfficetaxreportListElement**](OfficetaxreportListElement.md) |  | 

## Methods

### NewOfficetaxreportGetListV1ResponseMPayload

`func NewOfficetaxreportGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjOfficetaxreport []OfficetaxreportListElement, ) *OfficetaxreportGetListV1ResponseMPayload`

NewOfficetaxreportGetListV1ResponseMPayload instantiates a new OfficetaxreportGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfficetaxreportGetListV1ResponseMPayloadWithDefaults

`func NewOfficetaxreportGetListV1ResponseMPayloadWithDefaults() *OfficetaxreportGetListV1ResponseMPayload`

NewOfficetaxreportGetListV1ResponseMPayloadWithDefaults instantiates a new OfficetaxreportGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *OfficetaxreportGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *OfficetaxreportGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *OfficetaxreportGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *OfficetaxreportGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *OfficetaxreportGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *OfficetaxreportGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjOfficetaxreport

`func (o *OfficetaxreportGetListV1ResponseMPayload) GetAObjOfficetaxreport() []OfficetaxreportListElement`

GetAObjOfficetaxreport returns the AObjOfficetaxreport field if non-nil, zero value otherwise.

### GetAObjOfficetaxreportOk

`func (o *OfficetaxreportGetListV1ResponseMPayload) GetAObjOfficetaxreportOk() (*[]OfficetaxreportListElement, bool)`

GetAObjOfficetaxreportOk returns a tuple with the AObjOfficetaxreport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjOfficetaxreport

`func (o *OfficetaxreportGetListV1ResponseMPayload) SetAObjOfficetaxreport(v []OfficetaxreportListElement)`

SetAObjOfficetaxreport sets AObjOfficetaxreport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


