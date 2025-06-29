# software.amzn.spapi.Model.orders.v0.ExceptionDates
Dates when the business is closed or open with a different time window.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceptionDate** | **string** | Date when the business is closed, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date format. | [optional] 
**IsOpen** | **bool** | Boolean indicating if the business is closed or open on that date. | [optional] 
**OpenIntervals** | [**List&lt;OpenInterval&gt;**](OpenInterval.md) | Time window during the day when the business is open. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

