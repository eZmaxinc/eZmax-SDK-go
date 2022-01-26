# Go API client for eZmaxApi

This API expose all the functionnalities for the eZmax and eZsign applications.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.4
- Package version: 1.1.4
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.ezmax.ca/en/contact](https://www.ezmax.ca/en/contact)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import eZmaxApi "github.com/ezmaxinc/ezmax-sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), eZmaxApi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), eZmaxApi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), eZmaxApi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), eZmaxApi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*GlobalCustomerApi* | [**GlobalCustomerGetEndpointV1**](docs/GlobalCustomerApi.md#globalcustomergetendpointv1) | **Get** /1/customer/{pksCustomerCode}/endpoint | Get customer endpoint
*ModuleAuthenticateApi* | [**AuthenticateAuthenticateV2**](docs/ModuleAuthenticateApi.md#authenticateauthenticatev2) | **Post** /2/module/authenticate/authenticate/{eSessionType} | Authenticate a user
*ModuleListApi* | [**ListGetListpresentationV1**](docs/ModuleListApi.md#listgetlistpresentationv1) | **Get** /1/module/list/listpresentation/{sListName} | Get all Listpresentation for a specific list
*ModuleListApi* | [**ListSaveListpresentationV1**](docs/ModuleListApi.md#listsavelistpresentationv1) | **Post** /1/module/list/listpresentation/{sListName} | Save all Listpresentation for a specific list
*ModuleSsprApi* | [**SsprResetPasswordRequestV1**](docs/ModuleSsprApi.md#ssprresetpasswordrequestv1) | **Post** /1/module/sspr/resetPasswordRequest | Reset Password Request
*ModuleSsprApi* | [**SsprResetPasswordV1**](docs/ModuleSsprApi.md#ssprresetpasswordv1) | **Post** /1/module/sspr/resetPassword | Reset Password
*ModuleSsprApi* | [**SsprSendUsernamesV1**](docs/ModuleSsprApi.md#ssprsendusernamesv1) | **Post** /1/module/sspr/sendUsernames | Send username(s)
*ModuleSsprApi* | [**SsprUnlockAccountRequestV1**](docs/ModuleSsprApi.md#ssprunlockaccountrequestv1) | **Post** /1/module/sspr/unlockAccountRequest | Unlock Account Request
*ModuleSsprApi* | [**SsprUnlockAccountV1**](docs/ModuleSsprApi.md#ssprunlockaccountv1) | **Post** /1/module/sspr/unlockAccount | Unlock Account
*ModuleSsprApi* | [**SsprValidateTokenV1**](docs/ModuleSsprApi.md#ssprvalidatetokenv1) | **Post** /1/module/sspr/validateToken | Validate Token
*ModuleUserApi* | [**UserCreateEzsignuserV1**](docs/ModuleUserApi.md#usercreateezsignuserv1) | **Post** /1/module/user/createezsignuser | Create a new User of type Ezsignuser
*ObjectActivesessionApi* | [**ActivesessionGetCurrentV1**](docs/ObjectActivesessionApi.md#activesessiongetcurrentv1) | **Get** /1/object/activesession/getCurrent | Get Current Activesession
*ObjectApikeyApi* | [**ApikeyCreateObjectV1**](docs/ObjectApikeyApi.md#apikeycreateobjectv1) | **Post** /1/object/apikey | Create a new Apikey
*ObjectEzsignbulksendApi* | [**EzsignbulksendGetListV1**](docs/ObjectEzsignbulksendApi.md#ezsignbulksendgetlistv1) | **Get** /1/object/ezsignbulksend/getList | Retrieve Ezsignbulksend list
*ObjectEzsigndocumentApi* | [**EzsigndocumentApplyEzsigntemplateV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentapplyezsigntemplatev1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate | Apply an Ezsign Template to the Ezsigndocument.
*ObjectEzsigndocumentApi* | [**EzsigndocumentApplyEzsigntemplateV2**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentapplyezsigntemplatev2) | **Post** /2/object/ezsigndocument/{pkiEzsigndocumentID}/applyEzsigntemplate | Apply an Ezsign Template to the Ezsigndocument.
*ObjectEzsigndocumentApi* | [**EzsigndocumentCreateObjectV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentcreateobjectv1) | **Post** /1/object/ezsigndocument | Create a new Ezsigndocument
*ObjectEzsigndocumentApi* | [**EzsigndocumentDeleteObjectV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentdeleteobjectv1) | **Delete** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Delete an existing Ezsigndocument
*ObjectEzsigndocumentApi* | [**EzsigndocumentGetDownloadUrlV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentgetdownloadurlv1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getDownloadUrl/{eDocumentType} | Retrieve a URL to download documents.
*ObjectEzsigndocumentApi* | [**EzsigndocumentGetEzsignpagesV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentgetezsignpagesv1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getEzsignpages | Retrieve an existing Ezsigndocument&#39;s Ezsignpages
*ObjectEzsigndocumentApi* | [**EzsigndocumentGetFormDataV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentgetformdatav1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getFormData | Retrieve an existing Ezsigndocument&#39;s Form Data
*ObjectEzsigndocumentApi* | [**EzsigndocumentGetObjectV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentgetobjectv1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Retrieve an existing Ezsigndocument
*ObjectEzsigndocumentApi* | [**EzsigndocumentGetWordsPositionsV1**](docs/ObjectEzsigndocumentApi.md#ezsigndocumentgetwordspositionsv1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getWordsPositions | Retrieve positions X,Y of given words from a Ezsigndocument
*ObjectEzsignfolderApi* | [**EzsignfolderCreateObjectV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldercreateobjectv1) | **Post** /1/object/ezsignfolder | Create a new Ezsignfolder
*ObjectEzsignfolderApi* | [**EzsignfolderDeleteObjectV1**](docs/ObjectEzsignfolderApi.md#ezsignfolderdeleteobjectv1) | **Delete** /1/object/ezsignfolder/{pkiEzsignfolderID} | Delete an existing Ezsignfolder
*ObjectEzsignfolderApi* | [**EzsignfolderGetEzsigndocumentsV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldergetezsigndocumentsv1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getEzsigndocuments | Retrieve an existing Ezsignfolder&#39;s Ezsigndocuments
*ObjectEzsignfolderApi* | [**EzsignfolderGetEzsignfoldersignerassociationsV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldergetezsignfoldersignerassociationsv1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getEzsignfoldersignerassociations | Retrieve an existing Ezsignfolder&#39;s Ezsignfoldersignerassociations
*ObjectEzsignfolderApi* | [**EzsignfolderGetFormsDataV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldergetformsdatav1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getFormsData | Retrieve an existing Ezsignfolder&#39;s forms data
*ObjectEzsignfolderApi* | [**EzsignfolderGetListV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldergetlistv1) | **Get** /1/object/ezsignfolder/getList | Retrieve Ezsignfolder list
*ObjectEzsignfolderApi* | [**EzsignfolderGetObjectV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldergetobjectv1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID} | Retrieve an existing Ezsignfolder
*ObjectEzsignfolderApi* | [**EzsignfolderSendV1**](docs/ObjectEzsignfolderApi.md#ezsignfoldersendv1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/send | Send the Ezsignfolder to the signatories for signature
*ObjectEzsignfolderApi* | [**EzsignfolderUnsendV1**](docs/ObjectEzsignfolderApi.md#ezsignfolderunsendv1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/unsend | Unsend the Ezsignfolder
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationCreateObjectV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationcreateobjectv1) | **Post** /1/object/ezsignfoldersignerassociation | Create a new Ezsignfoldersignerassociation
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationDeleteObjectV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationdeleteobjectv1) | **Delete** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Delete an existing Ezsignfoldersignerassociation
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationGetInPersonLoginUrlV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationgetinpersonloginurlv1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/getInPersonLoginUrl | Retrieve a Login Url to allow In-Person signing
*ObjectEzsignfoldersignerassociationApi* | [**EzsignfoldersignerassociationGetObjectV1**](docs/ObjectEzsignfoldersignerassociationApi.md#ezsignfoldersignerassociationgetobjectv1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Retrieve an existing Ezsignfoldersignerassociation
*ObjectEzsignfoldertypeApi* | [**EzsignfoldertypeGetAutocompleteV1**](docs/ObjectEzsignfoldertypeApi.md#ezsignfoldertypegetautocompletev1) | **Get** /1/object/ezsignfoldertype/getAutocomplete/{sSelector} | Retrieve Ezsignfoldertypes and IDs
*ObjectEzsignfoldertypeApi* | [**EzsignfoldertypeGetListV1**](docs/ObjectEzsignfoldertypeApi.md#ezsignfoldertypegetlistv1) | **Get** /1/object/ezsignfoldertype/getList | Retrieve Ezsignfoldertype list
*ObjectEzsignsignatureApi* | [**EzsignsignatureCreateObjectV1**](docs/ObjectEzsignsignatureApi.md#ezsignsignaturecreateobjectv1) | **Post** /1/object/ezsignsignature | Create a new Ezsignsignature
*ObjectEzsignsignatureApi* | [**EzsignsignatureDeleteObjectV1**](docs/ObjectEzsignsignatureApi.md#ezsignsignaturedeleteobjectv1) | **Delete** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Delete an existing Ezsignsignature
*ObjectEzsignsignatureApi* | [**EzsignsignatureGetObjectV1**](docs/ObjectEzsignsignatureApi.md#ezsignsignaturegetobjectv1) | **Get** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Retrieve an existing Ezsignsignature
*ObjectEzsigntemplatepackageApi* | [**EzsigntemplatepackageGetListV1**](docs/ObjectEzsigntemplatepackageApi.md#ezsigntemplatepackagegetlistv1) | **Get** /1/object/ezsigntemplatepackage/getList | Retrieve Ezsigntemplatepackage list
*ObjectFranchisebrokerApi* | [**FranchisebrokerGetAutocompleteV1**](docs/ObjectFranchisebrokerApi.md#franchisebrokergetautocompletev1) | **Get** /1/object/franchisebroker/getAutocomplete/{sSelector} | Retrieve Franchisebrokers and IDs
*ObjectFranchiseofficeApi* | [**FranchiseofficeGetAutocompleteV1**](docs/ObjectFranchiseofficeApi.md#franchiseofficegetautocompletev1) | **Get** /1/object/franchiseoffice/getAutocomplete/{sSelector} | Retrieve Franchiseoffices and IDs
*ObjectFranchisereferalincomeApi* | [**FranchisereferalincomeCreateObjectV1**](docs/ObjectFranchisereferalincomeApi.md#franchisereferalincomecreateobjectv1) | **Post** /1/object/franchisereferalincome | Create a new Franchisereferalincome
*ObjectPeriodApi* | [**PeriodGetAutocompleteV1**](docs/ObjectPeriodApi.md#periodgetautocompletev1) | **Get** /1/object/period/getAutocomplete/{sSelector} | Retrieve Periods and IDs


## Documentation For Models

 - [ActivesessionGetCurrentV1Response](docs/ActivesessionGetCurrentV1Response.md)
 - [ActivesessionGetCurrentV1ResponseAllOf](docs/ActivesessionGetCurrentV1ResponseAllOf.md)
 - [ActivesessionGetCurrentV1ResponseMPayload](docs/ActivesessionGetCurrentV1ResponseMPayload.md)
 - [AddressRequest](docs/AddressRequest.md)
 - [ApikeyCreateObjectV1Request](docs/ApikeyCreateObjectV1Request.md)
 - [ApikeyCreateObjectV1Response](docs/ApikeyCreateObjectV1Response.md)
 - [ApikeyCreateObjectV1ResponseAllOf](docs/ApikeyCreateObjectV1ResponseAllOf.md)
 - [ApikeyCreateObjectV1ResponseMPayload](docs/ApikeyCreateObjectV1ResponseMPayload.md)
 - [ApikeyRequest](docs/ApikeyRequest.md)
 - [ApikeyRequestCompound](docs/ApikeyRequestCompound.md)
 - [ApikeyResponse](docs/ApikeyResponse.md)
 - [AttemptResponse](docs/AttemptResponse.md)
 - [AuthenticateAuthenticateV2Request](docs/AuthenticateAuthenticateV2Request.md)
 - [AuthenticateAuthenticateV2Response](docs/AuthenticateAuthenticateV2Response.md)
 - [AuthenticateAuthenticateV2ResponseAllOf](docs/AuthenticateAuthenticateV2ResponseAllOf.md)
 - [AuthenticateAuthenticateV2ResponseMPayload](docs/AuthenticateAuthenticateV2ResponseMPayload.md)
 - [CommonAudit](docs/CommonAudit.md)
 - [CommonGetAutocompleteV1Response](docs/CommonGetAutocompleteV1Response.md)
 - [CommonGetAutocompleteV1ResponseAllOf](docs/CommonGetAutocompleteV1ResponseAllOf.md)
 - [CommonGetListV1ResponseMPayload](docs/CommonGetListV1ResponseMPayload.md)
 - [CommonResponse](docs/CommonResponse.md)
 - [CommonResponseError](docs/CommonResponseError.md)
 - [CommonResponseErrorSTemporaryFileUrl](docs/CommonResponseErrorSTemporaryFileUrl.md)
 - [CommonResponseErrorSTemporaryFileUrlAllOf](docs/CommonResponseErrorSTemporaryFileUrlAllOf.md)
 - [CommonResponseFilter](docs/CommonResponseFilter.md)
 - [CommonResponseGetList](docs/CommonResponseGetList.md)
 - [CommonResponseObjDebug](docs/CommonResponseObjDebug.md)
 - [CommonResponseObjDebugPayload](docs/CommonResponseObjDebugPayload.md)
 - [CommonResponseObjDebugPayloadGetList](docs/CommonResponseObjDebugPayloadGetList.md)
 - [CommonResponseObjDebugPayloadGetListAllOf](docs/CommonResponseObjDebugPayloadGetListAllOf.md)
 - [CommonResponseObjSQLQuery](docs/CommonResponseObjSQLQuery.md)
 - [CommonWebhook](docs/CommonWebhook.md)
 - [ContactRequest](docs/ContactRequest.md)
 - [ContactRequestCompound](docs/ContactRequestCompound.md)
 - [ContactRequestCompoundAllOf](docs/ContactRequestCompoundAllOf.md)
 - [ContactinformationsRequest](docs/ContactinformationsRequest.md)
 - [ContactinformationsRequestCompound](docs/ContactinformationsRequestCompound.md)
 - [ContactinformationsRequestCompoundAllOf](docs/ContactinformationsRequestCompoundAllOf.md)
 - [CustomAutocompleteElementResponse](docs/CustomAutocompleteElementResponse.md)
 - [CustomEzsignfoldersignerassociationstatusResponse](docs/CustomEzsignfoldersignerassociationstatusResponse.md)
 - [CustomEzsignsignaturestatusResponse](docs/CustomEzsignsignaturestatusResponse.md)
 - [CustomFormDataDocumentResponse](docs/CustomFormDataDocumentResponse.md)
 - [CustomFormDataSignerResponse](docs/CustomFormDataSignerResponse.md)
 - [CustomFormsDataFolderResponse](docs/CustomFormsDataFolderResponse.md)
 - [CustomWordPositionOccurenceResponse](docs/CustomWordPositionOccurenceResponse.md)
 - [CustomWordPositionWordResponse](docs/CustomWordPositionWordResponse.md)
 - [EmailRequest](docs/EmailRequest.md)
 - [EzsignbulksendGetListV1Response](docs/EzsignbulksendGetListV1Response.md)
 - [EzsignbulksendGetListV1ResponseAllOf](docs/EzsignbulksendGetListV1ResponseAllOf.md)
 - [EzsignbulksendGetListV1ResponseMPayload](docs/EzsignbulksendGetListV1ResponseMPayload.md)
 - [EzsignbulksendGetListV1ResponseMPayloadAllOf](docs/EzsignbulksendGetListV1ResponseMPayloadAllOf.md)
 - [EzsignbulksendListElement](docs/EzsignbulksendListElement.md)
 - [EzsigndocumentApplyEzsigntemplateV1Request](docs/EzsigndocumentApplyEzsigntemplateV1Request.md)
 - [EzsigndocumentApplyEzsigntemplateV1Response](docs/EzsigndocumentApplyEzsigntemplateV1Response.md)
 - [EzsigndocumentApplyEzsigntemplateV2Request](docs/EzsigndocumentApplyEzsigntemplateV2Request.md)
 - [EzsigndocumentApplyEzsigntemplateV2Response](docs/EzsigndocumentApplyEzsigntemplateV2Response.md)
 - [EzsigndocumentCreateObjectV1Request](docs/EzsigndocumentCreateObjectV1Request.md)
 - [EzsigndocumentCreateObjectV1Response](docs/EzsigndocumentCreateObjectV1Response.md)
 - [EzsigndocumentCreateObjectV1ResponseAllOf](docs/EzsigndocumentCreateObjectV1ResponseAllOf.md)
 - [EzsigndocumentCreateObjectV1ResponseMPayload](docs/EzsigndocumentCreateObjectV1ResponseMPayload.md)
 - [EzsigndocumentDeleteObjectV1Response](docs/EzsigndocumentDeleteObjectV1Response.md)
 - [EzsigndocumentGetDownloadUrlV1Response](docs/EzsigndocumentGetDownloadUrlV1Response.md)
 - [EzsigndocumentGetDownloadUrlV1ResponseAllOf](docs/EzsigndocumentGetDownloadUrlV1ResponseAllOf.md)
 - [EzsigndocumentGetDownloadUrlV1ResponseMPayload](docs/EzsigndocumentGetDownloadUrlV1ResponseMPayload.md)
 - [EzsigndocumentGetEzsignpagesV1Response](docs/EzsigndocumentGetEzsignpagesV1Response.md)
 - [EzsigndocumentGetEzsignpagesV1ResponseAllOf](docs/EzsigndocumentGetEzsignpagesV1ResponseAllOf.md)
 - [EzsigndocumentGetEzsignpagesV1ResponseMPayload](docs/EzsigndocumentGetEzsignpagesV1ResponseMPayload.md)
 - [EzsigndocumentGetFormDataV1Response](docs/EzsigndocumentGetFormDataV1Response.md)
 - [EzsigndocumentGetFormDataV1ResponseAllOf](docs/EzsigndocumentGetFormDataV1ResponseAllOf.md)
 - [EzsigndocumentGetFormDataV1ResponseMPayload](docs/EzsigndocumentGetFormDataV1ResponseMPayload.md)
 - [EzsigndocumentGetObjectV1Response](docs/EzsigndocumentGetObjectV1Response.md)
 - [EzsigndocumentGetObjectV1ResponseAllOf](docs/EzsigndocumentGetObjectV1ResponseAllOf.md)
 - [EzsigndocumentGetObjectV1ResponseMPayload](docs/EzsigndocumentGetObjectV1ResponseMPayload.md)
 - [EzsigndocumentGetWordsPositionsV1Request](docs/EzsigndocumentGetWordsPositionsV1Request.md)
 - [EzsigndocumentGetWordsPositionsV1Response](docs/EzsigndocumentGetWordsPositionsV1Response.md)
 - [EzsigndocumentGetWordsPositionsV1ResponseAllOf](docs/EzsigndocumentGetWordsPositionsV1ResponseAllOf.md)
 - [EzsigndocumentRequest](docs/EzsigndocumentRequest.md)
 - [EzsigndocumentRequestCompound](docs/EzsigndocumentRequestCompound.md)
 - [EzsigndocumentResponse](docs/EzsigndocumentResponse.md)
 - [EzsigndocumentResponseCompound](docs/EzsigndocumentResponseCompound.md)
 - [EzsignfolderCreateObjectV1Request](docs/EzsignfolderCreateObjectV1Request.md)
 - [EzsignfolderCreateObjectV1Response](docs/EzsignfolderCreateObjectV1Response.md)
 - [EzsignfolderCreateObjectV1ResponseAllOf](docs/EzsignfolderCreateObjectV1ResponseAllOf.md)
 - [EzsignfolderCreateObjectV1ResponseMPayload](docs/EzsignfolderCreateObjectV1ResponseMPayload.md)
 - [EzsignfolderDeleteObjectV1Response](docs/EzsignfolderDeleteObjectV1Response.md)
 - [EzsignfolderGetEzsigndocumentsV1Response](docs/EzsignfolderGetEzsigndocumentsV1Response.md)
 - [EzsignfolderGetEzsigndocumentsV1ResponseAllOf](docs/EzsignfolderGetEzsigndocumentsV1ResponseAllOf.md)
 - [EzsignfolderGetEzsigndocumentsV1ResponseMPayload](docs/EzsignfolderGetEzsigndocumentsV1ResponseMPayload.md)
 - [EzsignfolderGetEzsignfoldersignerassociationsV1Response](docs/EzsignfolderGetEzsignfoldersignerassociationsV1Response.md)
 - [EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf](docs/EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf.md)
 - [EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload](docs/EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload.md)
 - [EzsignfolderGetFormsDataV1Response](docs/EzsignfolderGetFormsDataV1Response.md)
 - [EzsignfolderGetFormsDataV1ResponseAllOf](docs/EzsignfolderGetFormsDataV1ResponseAllOf.md)
 - [EzsignfolderGetFormsDataV1ResponseMPayload](docs/EzsignfolderGetFormsDataV1ResponseMPayload.md)
 - [EzsignfolderGetListV1Response](docs/EzsignfolderGetListV1Response.md)
 - [EzsignfolderGetListV1ResponseAllOf](docs/EzsignfolderGetListV1ResponseAllOf.md)
 - [EzsignfolderGetListV1ResponseMPayload](docs/EzsignfolderGetListV1ResponseMPayload.md)
 - [EzsignfolderGetListV1ResponseMPayloadAllOf](docs/EzsignfolderGetListV1ResponseMPayloadAllOf.md)
 - [EzsignfolderGetObjectV1Response](docs/EzsignfolderGetObjectV1Response.md)
 - [EzsignfolderGetObjectV1ResponseAllOf](docs/EzsignfolderGetObjectV1ResponseAllOf.md)
 - [EzsignfolderGetObjectV1ResponseMPayload](docs/EzsignfolderGetObjectV1ResponseMPayload.md)
 - [EzsignfolderListElement](docs/EzsignfolderListElement.md)
 - [EzsignfolderRequest](docs/EzsignfolderRequest.md)
 - [EzsignfolderRequestCompound](docs/EzsignfolderRequestCompound.md)
 - [EzsignfolderResponse](docs/EzsignfolderResponse.md)
 - [EzsignfolderResponseCompound](docs/EzsignfolderResponseCompound.md)
 - [EzsignfolderSendV1Request](docs/EzsignfolderSendV1Request.md)
 - [EzsignfolderSendV1Response](docs/EzsignfolderSendV1Response.md)
 - [EzsignfolderUnsendV1Response](docs/EzsignfolderUnsendV1Response.md)
 - [EzsignfoldersignerassociationCreateObjectV1Request](docs/EzsignfoldersignerassociationCreateObjectV1Request.md)
 - [EzsignfoldersignerassociationCreateObjectV1Response](docs/EzsignfoldersignerassociationCreateObjectV1Response.md)
 - [EzsignfoldersignerassociationCreateObjectV1ResponseAllOf](docs/EzsignfoldersignerassociationCreateObjectV1ResponseAllOf.md)
 - [EzsignfoldersignerassociationCreateObjectV1ResponseMPayload](docs/EzsignfoldersignerassociationCreateObjectV1ResponseMPayload.md)
 - [EzsignfoldersignerassociationDeleteObjectV1Response](docs/EzsignfoldersignerassociationDeleteObjectV1Response.md)
 - [EzsignfoldersignerassociationGetInPersonLoginUrlV1Response](docs/EzsignfoldersignerassociationGetInPersonLoginUrlV1Response.md)
 - [EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf](docs/EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf.md)
 - [EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload](docs/EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload.md)
 - [EzsignfoldersignerassociationGetObjectV1Response](docs/EzsignfoldersignerassociationGetObjectV1Response.md)
 - [EzsignfoldersignerassociationGetObjectV1ResponseAllOf](docs/EzsignfoldersignerassociationGetObjectV1ResponseAllOf.md)
 - [EzsignfoldersignerassociationGetObjectV1ResponseMPayload](docs/EzsignfoldersignerassociationGetObjectV1ResponseMPayload.md)
 - [EzsignfoldersignerassociationRequest](docs/EzsignfoldersignerassociationRequest.md)
 - [EzsignfoldersignerassociationRequestCompound](docs/EzsignfoldersignerassociationRequestCompound.md)
 - [EzsignfoldersignerassociationRequestCompoundAllOf](docs/EzsignfoldersignerassociationRequestCompoundAllOf.md)
 - [EzsignfoldersignerassociationResponse](docs/EzsignfoldersignerassociationResponse.md)
 - [EzsignfoldersignerassociationResponseCompound](docs/EzsignfoldersignerassociationResponseCompound.md)
 - [EzsignfoldertypeGetListV1Response](docs/EzsignfoldertypeGetListV1Response.md)
 - [EzsignfoldertypeGetListV1ResponseAllOf](docs/EzsignfoldertypeGetListV1ResponseAllOf.md)
 - [EzsignfoldertypeGetListV1ResponseMPayload](docs/EzsignfoldertypeGetListV1ResponseMPayload.md)
 - [EzsignfoldertypeGetListV1ResponseMPayloadAllOf](docs/EzsignfoldertypeGetListV1ResponseMPayloadAllOf.md)
 - [EzsignfoldertypeListElement](docs/EzsignfoldertypeListElement.md)
 - [EzsignformfieldResponse](docs/EzsignformfieldResponse.md)
 - [EzsignformfieldResponseCompound](docs/EzsignformfieldResponseCompound.md)
 - [EzsignformfieldgroupResponse](docs/EzsignformfieldgroupResponse.md)
 - [EzsignformfieldgroupResponseCompound](docs/EzsignformfieldgroupResponseCompound.md)
 - [EzsignformfieldgroupResponseCompoundAllOf](docs/EzsignformfieldgroupResponseCompoundAllOf.md)
 - [EzsignpageResponse](docs/EzsignpageResponse.md)
 - [EzsignsignatureCreateObjectV1Request](docs/EzsignsignatureCreateObjectV1Request.md)
 - [EzsignsignatureCreateObjectV1Response](docs/EzsignsignatureCreateObjectV1Response.md)
 - [EzsignsignatureCreateObjectV1ResponseAllOf](docs/EzsignsignatureCreateObjectV1ResponseAllOf.md)
 - [EzsignsignatureCreateObjectV1ResponseMPayload](docs/EzsignsignatureCreateObjectV1ResponseMPayload.md)
 - [EzsignsignatureDeleteObjectV1Response](docs/EzsignsignatureDeleteObjectV1Response.md)
 - [EzsignsignatureGetObjectV1Response](docs/EzsignsignatureGetObjectV1Response.md)
 - [EzsignsignatureGetObjectV1ResponseAllOf](docs/EzsignsignatureGetObjectV1ResponseAllOf.md)
 - [EzsignsignatureRequest](docs/EzsignsignatureRequest.md)
 - [EzsignsignatureRequestCompound](docs/EzsignsignatureRequestCompound.md)
 - [EzsignsignaturecustomdateRequest](docs/EzsignsignaturecustomdateRequest.md)
 - [EzsignsignerRequest](docs/EzsignsignerRequest.md)
 - [EzsignsignerRequestCompound](docs/EzsignsignerRequestCompound.md)
 - [EzsignsignerRequestCompoundAllOf](docs/EzsignsignerRequestCompoundAllOf.md)
 - [EzsignsignerRequestCompoundContact](docs/EzsignsignerRequestCompoundContact.md)
 - [EzsignsignerResponse](docs/EzsignsignerResponse.md)
 - [EzsignsignerResponseCompound](docs/EzsignsignerResponseCompound.md)
 - [EzsignsignerResponseCompoundAllOf](docs/EzsignsignerResponseCompoundAllOf.md)
 - [EzsignsignerResponseCompoundContact](docs/EzsignsignerResponseCompoundContact.md)
 - [EzsigntemplatepackageGetListV1Response](docs/EzsigntemplatepackageGetListV1Response.md)
 - [EzsigntemplatepackageGetListV1ResponseAllOf](docs/EzsigntemplatepackageGetListV1ResponseAllOf.md)
 - [EzsigntemplatepackageGetListV1ResponseMPayload](docs/EzsigntemplatepackageGetListV1ResponseMPayload.md)
 - [EzsigntemplatepackageGetListV1ResponseMPayloadAllOf](docs/EzsigntemplatepackageGetListV1ResponseMPayloadAllOf.md)
 - [EzsigntemplatepackageListElement](docs/EzsigntemplatepackageListElement.md)
 - [FieldEEzsigndocumentStep](docs/FieldEEzsigndocumentStep.md)
 - [FieldEEzsignfolderSendreminderfrequency](docs/FieldEEzsignfolderSendreminderfrequency.md)
 - [FieldEEzsignfolderStep](docs/FieldEEzsignfolderStep.md)
 - [FieldEEzsignfoldertypePrivacylevel](docs/FieldEEzsignfoldertypePrivacylevel.md)
 - [FieldEEzsignsignatureType](docs/FieldEEzsignsignatureType.md)
 - [FieldEEzsigntemplatepackageType](docs/FieldEEzsigntemplatepackageType.md)
 - [FieldEPhoneType](docs/FieldEPhoneType.md)
 - [FieldEUserType](docs/FieldEUserType.md)
 - [FieldEUserTypeSSPR](docs/FieldEUserTypeSSPR.md)
 - [FranchisereferalincomeCreateObjectV1Request](docs/FranchisereferalincomeCreateObjectV1Request.md)
 - [FranchisereferalincomeCreateObjectV1Response](docs/FranchisereferalincomeCreateObjectV1Response.md)
 - [FranchisereferalincomeCreateObjectV1ResponseAllOf](docs/FranchisereferalincomeCreateObjectV1ResponseAllOf.md)
 - [FranchisereferalincomeCreateObjectV1ResponseMPayload](docs/FranchisereferalincomeCreateObjectV1ResponseMPayload.md)
 - [FranchisereferalincomeRequest](docs/FranchisereferalincomeRequest.md)
 - [FranchisereferalincomeRequestCompound](docs/FranchisereferalincomeRequestCompound.md)
 - [FranchisereferalincomeRequestCompoundAllOf](docs/FranchisereferalincomeRequestCompoundAllOf.md)
 - [GlobalCustomerGetEndpointV1Response](docs/GlobalCustomerGetEndpointV1Response.md)
 - [HeaderAcceptLanguage](docs/HeaderAcceptLanguage.md)
 - [ListGetListpresentationV1Response](docs/ListGetListpresentationV1Response.md)
 - [ListGetListpresentationV1ResponseAllOf](docs/ListGetListpresentationV1ResponseAllOf.md)
 - [ListGetListpresentationV1ResponseMPayload](docs/ListGetListpresentationV1ResponseMPayload.md)
 - [ListSaveListpresentationV1Request](docs/ListSaveListpresentationV1Request.md)
 - [ListSaveListpresentationV1Response](docs/ListSaveListpresentationV1Response.md)
 - [ListpresentationRequest](docs/ListpresentationRequest.md)
 - [MultilingualApikeyDescription](docs/MultilingualApikeyDescription.md)
 - [PhoneRequest](docs/PhoneRequest.md)
 - [SsprResetPasswordRequestV1Request](docs/SsprResetPasswordRequestV1Request.md)
 - [SsprResetPasswordV1Request](docs/SsprResetPasswordV1Request.md)
 - [SsprSendUsernamesV1Request](docs/SsprSendUsernamesV1Request.md)
 - [SsprUnlockAccountRequestV1Request](docs/SsprUnlockAccountRequestV1Request.md)
 - [SsprUnlockAccountV1Request](docs/SsprUnlockAccountV1Request.md)
 - [SsprValidateTokenV1Request](docs/SsprValidateTokenV1Request.md)
 - [UNUSEDEzsigndocumentEditObjectV1Request](docs/UNUSEDEzsigndocumentEditObjectV1Request.md)
 - [UNUSEDEzsigndocumentEditObjectV1Response](docs/UNUSEDEzsigndocumentEditObjectV1Response.md)
 - [UNUSEDEzsignfolderEditObjectV1Request](docs/UNUSEDEzsignfolderEditObjectV1Request.md)
 - [UNUSEDEzsignfolderEditObjectV1Response](docs/UNUSEDEzsignfolderEditObjectV1Response.md)
 - [UNUSEDEzsignfoldersignerassociationEditObjectV1Request](docs/UNUSEDEzsignfoldersignerassociationEditObjectV1Request.md)
 - [UNUSEDEzsignfoldersignerassociationEditObjectV1Response](docs/UNUSEDEzsignfoldersignerassociationEditObjectV1Response.md)
 - [UNUSEDEzsignsignatureEditObjectV1Request](docs/UNUSEDEzsignsignatureEditObjectV1Request.md)
 - [UNUSEDEzsignsignatureEditObjectV1Response](docs/UNUSEDEzsignsignatureEditObjectV1Response.md)
 - [UserCreateEzsignuserV1Request](docs/UserCreateEzsignuserV1Request.md)
 - [UserCreateEzsignuserV1Response](docs/UserCreateEzsignuserV1Response.md)
 - [UserCreateEzsignuserV1ResponseAllOf](docs/UserCreateEzsignuserV1ResponseAllOf.md)
 - [UserCreateEzsignuserV1ResponseMPayload](docs/UserCreateEzsignuserV1ResponseMPayload.md)
 - [UserResponse](docs/UserResponse.md)
 - [WebhookEzsignDocumentCompleted](docs/WebhookEzsignDocumentCompleted.md)
 - [WebhookEzsignDocumentCompletedAllOf](docs/WebhookEzsignDocumentCompletedAllOf.md)
 - [WebhookEzsignFolderCompleted](docs/WebhookEzsignFolderCompleted.md)
 - [WebhookEzsignFolderCompletedAllOf](docs/WebhookEzsignFolderCompletedAllOf.md)
 - [WebhookResponse](docs/WebhookResponse.md)
 - [WebhookUserUserCreated](docs/WebhookUserUserCreated.md)
 - [WebhookUserUserCreatedAllOf](docs/WebhookUserUserCreatedAllOf.md)
 - [WebsiteRequest](docs/WebsiteRequest.md)


## Documentation For Authorization



### Authorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support-api@ezmax.ca

