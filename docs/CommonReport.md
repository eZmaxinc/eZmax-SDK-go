# CommonReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjReportsection** | [**[]CommonReportsection**](CommonReportsection.md) |  | 
**BReportPaginate** | Pointer to **bool** | Whether we display pagination in the report | [optional] 
**SReportTitle** | Pointer to **string** | The title of this Report | [optional] 

## Methods

### NewCommonReport

`func NewCommonReport(aObjReportsection []CommonReportsection, ) *CommonReport`

NewCommonReport instantiates a new CommonReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonReportWithDefaults

`func NewCommonReportWithDefaults() *CommonReport`

NewCommonReportWithDefaults instantiates a new CommonReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjReportsection

`func (o *CommonReport) GetAObjReportsection() []CommonReportsection`

GetAObjReportsection returns the AObjReportsection field if non-nil, zero value otherwise.

### GetAObjReportsectionOk

`func (o *CommonReport) GetAObjReportsectionOk() (*[]CommonReportsection, bool)`

GetAObjReportsectionOk returns a tuple with the AObjReportsection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjReportsection

`func (o *CommonReport) SetAObjReportsection(v []CommonReportsection)`

SetAObjReportsection sets AObjReportsection field to given value.


### GetBReportPaginate

`func (o *CommonReport) GetBReportPaginate() bool`

GetBReportPaginate returns the BReportPaginate field if non-nil, zero value otherwise.

### GetBReportPaginateOk

`func (o *CommonReport) GetBReportPaginateOk() (*bool, bool)`

GetBReportPaginateOk returns a tuple with the BReportPaginate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBReportPaginate

`func (o *CommonReport) SetBReportPaginate(v bool)`

SetBReportPaginate sets BReportPaginate field to given value.

### HasBReportPaginate

`func (o *CommonReport) HasBReportPaginate() bool`

HasBReportPaginate returns a boolean if a field has been set.

### GetSReportTitle

`func (o *CommonReport) GetSReportTitle() string`

GetSReportTitle returns the SReportTitle field if non-nil, zero value otherwise.

### GetSReportTitleOk

`func (o *CommonReport) GetSReportTitleOk() (*string, bool)`

GetSReportTitleOk returns a tuple with the SReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSReportTitle

`func (o *CommonReport) SetSReportTitle(v string)`

SetSReportTitle sets SReportTitle field to given value.

### HasSReportTitle

`func (o *CommonReport) HasSReportTitle() bool`

HasSReportTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


