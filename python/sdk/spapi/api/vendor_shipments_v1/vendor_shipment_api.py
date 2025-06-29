# coding: utf-8

"""
    Selling Partner API for Retail Procurement Shipments

    The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.

    The version of the OpenAPI document: v1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import absolute_import

import re  # noqa: F401

# python 2 and python 3 compatibility library
import six

from spapi.api_client import ApiClient


class VendorShipmentApi(object):
    """NOTE: This class is auto generated by the openapi generator.

    Do not edit the class manually.
    """

    api_models_module = "spapi.models.vendor_shipments_v1"

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.classFileName = 'vendor_shipment_api'

    def get_shipment_details(self, **kwargs):  # noqa: E501
        """GetShipmentDetails  # noqa: E501

        Returns the Details about Shipment, Carrier Details,  status of the shipment, container details and other details related to shipment based on the filter parameters value that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_shipment_details(async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param int limit: The limit to the number of records returned. Default value is 50 records.
        :param str sort_order: Sort in ascending or descending order by purchase order creation date.
        :param str next_token: Used for pagination when there are more shipments than the specified result size limit.
        :param datetime created_after: Get Shipment Details that became available after this timestamp will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime created_before: Get Shipment Details that became available before this timestamp will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipment_confirmed_before: Get Shipment Details by passing Shipment confirmed create Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipment_confirmed_after: Get Shipment Details by passing Shipment confirmed create Date After. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime package_label_created_before: Get Shipment Details by passing Package label create Date by buyer. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime package_label_created_after: Get Shipment Details by passing Package label create Date After by buyer. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipped_before: Get Shipment Details by passing Shipped Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipped_after: Get Shipment Details by passing Shipped Date After. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime estimated_delivery_before: Get Shipment Details by passing Estimated Delivery Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime estimated_delivery_after: Get Shipment Details by passing Estimated Delivery Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipment_delivery_before: Get Shipment Details by passing Shipment Delivery Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipment_delivery_after: Get Shipment Details by passing Shipment Delivery Date After. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime requested_pick_up_before: Get Shipment Details by passing Before Requested pickup date. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime requested_pick_up_after: Get Shipment Details by passing After Requested pickup date. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime scheduled_pick_up_before: Get Shipment Details by passing Before scheduled pickup date. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime scheduled_pick_up_after: Get Shipment Details by passing After Scheduled pickup date. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param str current_shipment_status: Get Shipment Details by passing Current shipment status.
        :param str vendor_shipment_identifier: Get Shipment Details by passing Vendor Shipment ID
        :param str buyer_reference_number: Get Shipment Details by passing buyer Reference ID
        :param str buyer_warehouse_code: Get Shipping Details based on buyer warehouse code. This value should be same as 'shipToParty.partyId' in the Shipment.
        :param str seller_warehouse_code: Get Shipping Details based on vendor warehouse code. This value should be same as 'sellingParty.partyId' in the Shipment.
        :return: GetShipmentDetailsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.get_shipment_details_with_http_info(**kwargs)  # noqa: E501
        else:
            (data) = self.get_shipment_details_with_http_info(**kwargs)  # noqa: E501
            return data

    def get_shipment_details_with_http_info(self, **kwargs):  # noqa: E501
        """GetShipmentDetails  # noqa: E501

        Returns the Details about Shipment, Carrier Details,  status of the shipment, container details and other details related to shipment based on the filter parameters value that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_shipment_details_with_http_info(async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param int limit: The limit to the number of records returned. Default value is 50 records.
        :param str sort_order: Sort in ascending or descending order by purchase order creation date.
        :param str next_token: Used for pagination when there are more shipments than the specified result size limit.
        :param datetime created_after: Get Shipment Details that became available after this timestamp will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime created_before: Get Shipment Details that became available before this timestamp will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipment_confirmed_before: Get Shipment Details by passing Shipment confirmed create Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipment_confirmed_after: Get Shipment Details by passing Shipment confirmed create Date After. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime package_label_created_before: Get Shipment Details by passing Package label create Date by buyer. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime package_label_created_after: Get Shipment Details by passing Package label create Date After by buyer. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipped_before: Get Shipment Details by passing Shipped Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipped_after: Get Shipment Details by passing Shipped Date After. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime estimated_delivery_before: Get Shipment Details by passing Estimated Delivery Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime estimated_delivery_after: Get Shipment Details by passing Estimated Delivery Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipment_delivery_before: Get Shipment Details by passing Shipment Delivery Date Before. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime shipment_delivery_after: Get Shipment Details by passing Shipment Delivery Date After. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime requested_pick_up_before: Get Shipment Details by passing Before Requested pickup date. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime requested_pick_up_after: Get Shipment Details by passing After Requested pickup date. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime scheduled_pick_up_before: Get Shipment Details by passing Before scheduled pickup date. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param datetime scheduled_pick_up_after: Get Shipment Details by passing After Scheduled pickup date. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
        :param str current_shipment_status: Get Shipment Details by passing Current shipment status.
        :param str vendor_shipment_identifier: Get Shipment Details by passing Vendor Shipment ID
        :param str buyer_reference_number: Get Shipment Details by passing buyer Reference ID
        :param str buyer_warehouse_code: Get Shipping Details based on buyer warehouse code. This value should be same as 'shipToParty.partyId' in the Shipment.
        :param str seller_warehouse_code: Get Shipping Details based on vendor warehouse code. This value should be same as 'sellingParty.partyId' in the Shipment.
        :return: GetShipmentDetailsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['limit', 'sort_order', 'next_token', 'created_after', 'created_before', 'shipment_confirmed_before', 'shipment_confirmed_after', 'package_label_created_before', 'package_label_created_after', 'shipped_before', 'shipped_after', 'estimated_delivery_before', 'estimated_delivery_after', 'shipment_delivery_before', 'shipment_delivery_after', 'requested_pick_up_before', 'requested_pick_up_after', 'scheduled_pick_up_before', 'scheduled_pick_up_after', 'current_shipment_status', 'vendor_shipment_identifier', 'buyer_reference_number', 'buyer_warehouse_code', 'seller_warehouse_code']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method get_shipment_details" % key
                )
            params[key] = val
        del params['kwargs']

        if self.api_client.client_side_validation and ('limit' in params and params['limit'] > 50):  # noqa: E501
            raise ValueError("Invalid value for parameter `limit` when calling `get_shipment_details`, must be a value less than or equal to `50`")  # noqa: E501
        if self.api_client.client_side_validation and ('limit' in params and params['limit'] < 1):  # noqa: E501
            raise ValueError("Invalid value for parameter `limit` when calling `get_shipment_details`, must be a value greater than or equal to `1`")  # noqa: E501
        collection_formats = {}

        path_params = {}

        query_params = []
        if 'limit' in params:
            query_params.append(('limit', params['limit']))  # noqa: E501
        if 'sort_order' in params:
            query_params.append(('sortOrder', params['sort_order']))  # noqa: E501
        if 'next_token' in params:
            query_params.append(('nextToken', params['next_token']))  # noqa: E501
        if 'created_after' in params:
            query_params.append(('createdAfter', params['created_after']))  # noqa: E501
        if 'created_before' in params:
            query_params.append(('createdBefore', params['created_before']))  # noqa: E501
        if 'shipment_confirmed_before' in params:
            query_params.append(('shipmentConfirmedBefore', params['shipment_confirmed_before']))  # noqa: E501
        if 'shipment_confirmed_after' in params:
            query_params.append(('shipmentConfirmedAfter', params['shipment_confirmed_after']))  # noqa: E501
        if 'package_label_created_before' in params:
            query_params.append(('packageLabelCreatedBefore', params['package_label_created_before']))  # noqa: E501
        if 'package_label_created_after' in params:
            query_params.append(('packageLabelCreatedAfter', params['package_label_created_after']))  # noqa: E501
        if 'shipped_before' in params:
            query_params.append(('shippedBefore', params['shipped_before']))  # noqa: E501
        if 'shipped_after' in params:
            query_params.append(('shippedAfter', params['shipped_after']))  # noqa: E501
        if 'estimated_delivery_before' in params:
            query_params.append(('estimatedDeliveryBefore', params['estimated_delivery_before']))  # noqa: E501
        if 'estimated_delivery_after' in params:
            query_params.append(('estimatedDeliveryAfter', params['estimated_delivery_after']))  # noqa: E501
        if 'shipment_delivery_before' in params:
            query_params.append(('shipmentDeliveryBefore', params['shipment_delivery_before']))  # noqa: E501
        if 'shipment_delivery_after' in params:
            query_params.append(('shipmentDeliveryAfter', params['shipment_delivery_after']))  # noqa: E501
        if 'requested_pick_up_before' in params:
            query_params.append(('requestedPickUpBefore', params['requested_pick_up_before']))  # noqa: E501
        if 'requested_pick_up_after' in params:
            query_params.append(('requestedPickUpAfter', params['requested_pick_up_after']))  # noqa: E501
        if 'scheduled_pick_up_before' in params:
            query_params.append(('scheduledPickUpBefore', params['scheduled_pick_up_before']))  # noqa: E501
        if 'scheduled_pick_up_after' in params:
            query_params.append(('scheduledPickUpAfter', params['scheduled_pick_up_after']))  # noqa: E501
        if 'current_shipment_status' in params:
            query_params.append(('currentShipmentStatus', params['current_shipment_status']))  # noqa: E501
        if 'vendor_shipment_identifier' in params:
            query_params.append(('vendorShipmentIdentifier', params['vendor_shipment_identifier']))  # noqa: E501
        if 'buyer_reference_number' in params:
            query_params.append(('buyerReferenceNumber', params['buyer_reference_number']))  # noqa: E501
        if 'buyer_warehouse_code' in params:
            query_params.append(('buyerWarehouseCode', params['buyer_warehouse_code']))  # noqa: E501
        if 'seller_warehouse_code' in params:
            query_params.append(('sellerWarehouseCode', params['seller_warehouse_code']))  # noqa: E501

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/vendor/shipping/v1/shipments', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='GetShipmentDetailsResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats, api_models_module=self.api_models_module)

    def get_shipment_labels(self, **kwargs):  # noqa: E501
        """get_shipment_labels  # noqa: E501

        Returns small parcel shipment labels based on the filters that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_shipment_labels(async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param int limit: The limit to the number of records returned. Default value is 50 records.
        :param str sort_order: Sort the list by shipment label creation date in ascending or descending order.
        :param str next_token: A token that you use to retrieve the next page of results. The response includes `nextToken` when the number of results exceeds the specified `pageSize` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until `nextToken` is null. Note that this operation can return empty pages.
        :param datetime label_created_after: Shipment labels created after this time will be included in the result. This field must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format.
        :param datetime label_created_before: Shipment labels created before this time will be included in the result. This field must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format.
        :param str buyer_reference_number: Get Shipment labels by passing buyer reference number.
        :param str vendor_shipment_identifier: Get Shipment labels by passing vendor shipment identifier.
        :param str seller_warehouse_code: Get Shipping labels based on vendor warehouse code. This value must be same as the `sellingParty.partyId` in the shipment.
        :return: GetShipmentLabels
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.get_shipment_labels_with_http_info(**kwargs)  # noqa: E501
        else:
            (data) = self.get_shipment_labels_with_http_info(**kwargs)  # noqa: E501
            return data

    def get_shipment_labels_with_http_info(self, **kwargs):  # noqa: E501
        """get_shipment_labels  # noqa: E501

        Returns small parcel shipment labels based on the filters that you specify.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_shipment_labels_with_http_info(async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param int limit: The limit to the number of records returned. Default value is 50 records.
        :param str sort_order: Sort the list by shipment label creation date in ascending or descending order.
        :param str next_token: A token that you use to retrieve the next page of results. The response includes `nextToken` when the number of results exceeds the specified `pageSize` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until `nextToken` is null. Note that this operation can return empty pages.
        :param datetime label_created_after: Shipment labels created after this time will be included in the result. This field must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format.
        :param datetime label_created_before: Shipment labels created before this time will be included in the result. This field must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format.
        :param str buyer_reference_number: Get Shipment labels by passing buyer reference number.
        :param str vendor_shipment_identifier: Get Shipment labels by passing vendor shipment identifier.
        :param str seller_warehouse_code: Get Shipping labels based on vendor warehouse code. This value must be same as the `sellingParty.partyId` in the shipment.
        :return: GetShipmentLabels
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['limit', 'sort_order', 'next_token', 'label_created_after', 'label_created_before', 'buyer_reference_number', 'vendor_shipment_identifier', 'seller_warehouse_code']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method get_shipment_labels" % key
                )
            params[key] = val
        del params['kwargs']

        if self.api_client.client_side_validation and ('limit' in params and params['limit'] > 50):  # noqa: E501
            raise ValueError("Invalid value for parameter `limit` when calling `get_shipment_labels`, must be a value less than or equal to `50`")  # noqa: E501
        if self.api_client.client_side_validation and ('limit' in params and params['limit'] < 1):  # noqa: E501
            raise ValueError("Invalid value for parameter `limit` when calling `get_shipment_labels`, must be a value greater than or equal to `1`")  # noqa: E501
        collection_formats = {}

        path_params = {}

        query_params = []
        if 'limit' in params:
            query_params.append(('limit', params['limit']))  # noqa: E501
        if 'sort_order' in params:
            query_params.append(('sortOrder', params['sort_order']))  # noqa: E501
        if 'next_token' in params:
            query_params.append(('nextToken', params['next_token']))  # noqa: E501
        if 'label_created_after' in params:
            query_params.append(('labelCreatedAfter', params['label_created_after']))  # noqa: E501
        if 'label_created_before' in params:
            query_params.append(('labelCreatedBefore', params['label_created_before']))  # noqa: E501
        if 'buyer_reference_number' in params:
            query_params.append(('buyerReferenceNumber', params['buyer_reference_number']))  # noqa: E501
        if 'vendor_shipment_identifier' in params:
            query_params.append(('vendorShipmentIdentifier', params['vendor_shipment_identifier']))  # noqa: E501
        if 'seller_warehouse_code' in params:
            query_params.append(('sellerWarehouseCode', params['seller_warehouse_code']))  # noqa: E501

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/vendor/shipping/v1/transportLabels', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='GetShipmentLabels',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats, api_models_module=self.api_models_module)

    def submit_shipment_confirmations(self, body, **kwargs):  # noqa: E501
        """SubmitShipmentConfirmations  # noqa: E501

        Submits one or more shipment confirmations for vendor orders.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.submit_shipment_confirmations(body, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param SubmitShipmentConfirmationsRequest body: A request to submit shipment confirmation. (required)
        :return: SubmitShipmentConfirmationsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.submit_shipment_confirmations_with_http_info(body, **kwargs)  # noqa: E501
        else:
            (data) = self.submit_shipment_confirmations_with_http_info(body, **kwargs)  # noqa: E501
            return data

    def submit_shipment_confirmations_with_http_info(self, body, **kwargs):  # noqa: E501
        """SubmitShipmentConfirmations  # noqa: E501

        Submits one or more shipment confirmations for vendor orders.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.submit_shipment_confirmations_with_http_info(body, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param SubmitShipmentConfirmationsRequest body: A request to submit shipment confirmation. (required)
        :return: SubmitShipmentConfirmationsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['body']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method submit_shipment_confirmations" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'body' is set
        if self.api_client.client_side_validation and ('body' not in params or
                                                       params['body'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `body` when calling `submit_shipment_confirmations`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'body' in params:
            body_params = params['body']
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/vendor/shipping/v1/shipmentConfirmations', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='SubmitShipmentConfirmationsResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats, api_models_module=self.api_models_module)

    def submit_shipments(self, body, **kwargs):  # noqa: E501
        """SubmitShipments  # noqa: E501

        Submits one or more shipment request for vendor Orders.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.submit_shipments(body, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param SubmitShipments body: A request to submit shipment request. (required)
        :return: SubmitShipmentConfirmationsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.submit_shipments_with_http_info(body, **kwargs)  # noqa: E501
        else:
            (data) = self.submit_shipments_with_http_info(body, **kwargs)  # noqa: E501
            return data

    def submit_shipments_with_http_info(self, body, **kwargs):  # noqa: E501
        """SubmitShipments  # noqa: E501

        Submits one or more shipment request for vendor Orders.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | 10 | 10 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.submit_shipments_with_http_info(body, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param SubmitShipments body: A request to submit shipment request. (required)
        :return: SubmitShipmentConfirmationsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['body']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method submit_shipments" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'body' is set
        if self.api_client.client_side_validation and ('body' not in params or
                                                       params['body'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `body` when calling `submit_shipments`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        if 'body' in params:
            body_params = params['body']
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # HTTP header `Content-Type`
        header_params['Content-Type'] = self.api_client.select_header_content_type(  # noqa: E501
            ['application/json'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/vendor/shipping/v1/shipments', 'POST',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='SubmitShipmentConfirmationsResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats, api_models_module=self.api_models_module)
