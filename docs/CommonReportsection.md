# CommonReportsection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjReportsubsection** | [**[]CommonReportsubsection**](CommonReportsubsection.md) |  | 
**AObjReportcolumn** | [**[]CommonReportcolumn**](CommonReportcolumn.md) |  | 
**EReportsectionHorizontalalignment** | [**EnumHorizontalalignment**](EnumHorizontalalignment.md) |  | 
**IReportsectionColumncount** | **int32** | The number of Reportcolumns in the Reportsection | 
**IReportsectionWidth** | **int32** | The combined width of all the Reportcolumns in the Reportsection | 
**SReportsectionTitle** | Pointer to **string** | The title of this Reportsection | [optional] 

## Methods

### NewCommonReportsection

`func NewCommonReportsection(aObjReportsubsection []CommonReportsubsection, aObjReportcolumn []CommonReportcolumn, eReportsectionHorizontalalignment EnumHorizontalalignment, iReportsectionColumncount int32, iReportsectionWidth int32, ) *CommonReportsection`

NewCommonReportsection instantiates a new CommonReportsection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonReportsectionWithDefaults

`func NewCommonReportsectionWithDefaults() *CommonReportsection`

NewCommonReportsectionWithDefaults instantiates a new CommonReportsection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjReportsubsection

`func (o *CommonReportsection) GetAObjReportsubsection() []CommonReportsubsection`

GetAObjReportsubsection returns the AObjReportsubsection field if non-nil, zero value otherwise.

### GetAObjReportsubsectionOk

`func (o *CommonReportsection) GetAObjReportsubsectionOk() (*[]CommonReportsubsection, bool)`

GetAObjReportsubsectionOk returns a tuple with the AObjReportsubsection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjReportsubsection

`func (o *CommonReportsection) SetAObjReportsubsection(v []CommonReportsubsection)`

SetAObjReportsubsection sets AObjReportsubsection field to given value.


### GetAObjReportcolumn

`func (o *CommonReportsection) GetAObjReportcolumn() []CommonReportcolumn`

GetAObjReportcolumn returns the AObjReportcolumn field if non-nil, zero value otherwise.

### GetAObjReportcolumnOk

`func (o *CommonReportsection) GetAObjReportcolumnOk() (*[]CommonReportcolumn, bool)`

GetAObjReportcolumnOk returns a tuple with the AObjReportcolumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjReportcolumn

`func (o *CommonReportsection) SetAObjReportcolumn(v []CommonReportcolumn)`

SetAObjReportcolumn sets AObjReportcolumn field to given value.


### GetEReportsectionHorizontalalignment

`func (o *CommonReportsection) GetEReportsectionHorizontalalignment() EnumHorizontalalignment`

GetEReportsectionHorizontalalignment returns the EReportsectionHorizontalalignment field if non-nil, zero value otherwise.

### GetEReportsectionHorizontalalignmentOk

`func (o *CommonReportsection) GetEReportsectionHorizontalalignmentOk() (*EnumHorizontalalignment, bool)`

GetEReportsectionHorizontalalignmentOk returns a tuple with the EReportsectionHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEReportsectionHorizontalalignment

`func (o *CommonReportsection) SetEReportsectionHorizontalalignment(v EnumHorizontalalignment)`

SetEReportsectionHorizontalalignment sets EReportsectionHorizontalalignment field to given value.


### GetIReportsectionColumncount

`func (o *CommonReportsection) GetIReportsectionColumncount() int32`

GetIReportsectionColumncount returns the IReportsectionColumncount field if non-nil, zero value otherwise.

### GetIReportsectionColumncountOk

`func (o *CommonReportsection) GetIReportsectionColumncountOk() (*int32, bool)`

GetIReportsectionColumncountOk returns a tuple with the IReportsectionColumncount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIReportsectionColumncount

`func (o *CommonReportsection) SetIReportsectionColumncount(v int32)`

SetIReportsectionColumncount sets IReportsectionColumncount field to given value.


### GetIReportsectionWidth

`func (o *CommonReportsection) GetIReportsectionWidth() int32`

GetIReportsectionWidth returns the IReportsectionWidth field if non-nil, zero value otherwise.

### GetIReportsectionWidthOk

`func (o *CommonReportsection) GetIReportsectionWidthOk() (*int32, bool)`

GetIReportsectionWidthOk returns a tuple with the IReportsectionWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIReportsectionWidth

`func (o *CommonReportsection) SetIReportsectionWidth(v int32)`

SetIReportsectionWidth sets IReportsectionWidth field to given value.


### GetSReportsectionTitle

`func (o *CommonReportsection) GetSReportsectionTitle() string`

GetSReportsectionTitle returns the SReportsectionTitle field if non-nil, zero value otherwise.

### GetSReportsectionTitleOk

`func (o *CommonReportsection) GetSReportsectionTitleOk() (*string, bool)`

GetSReportsectionTitleOk returns a tuple with the SReportsectionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSReportsectionTitle

`func (o *CommonReportsection) SetSReportsectionTitle(v string)`

SetSReportsectionTitle sets SReportsectionTitle field to given value.

### HasSReportsectionTitle

`func (o *CommonReportsection) HasSReportsectionTitle() bool`

HasSReportsectionTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


