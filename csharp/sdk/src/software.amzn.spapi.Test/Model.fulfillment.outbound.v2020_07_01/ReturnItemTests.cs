/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * The version of the OpenAPI document: 2020-07-01
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using software.amzn.spapi.Model.fulfillment.outbound.v2020_07_01;
using software.amzn.spapi.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace software.amzn.spapi.Test.Model
{
    /// <summary>
    ///  Class for testing ReturnItem
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class ReturnItemTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for ReturnItem
        //private ReturnItem instance;

        public ReturnItemTests()
        {
            // TODO uncomment below to create an instance of ReturnItem
            //instance = new ReturnItem();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of ReturnItem
        /// </summary>
        [Fact]
        public void ReturnItemInstanceTest()
        {
            // TODO uncomment below to test "IsType" ReturnItem
            //Assert.IsType<ReturnItem>(instance);
        }

        /// <summary>
        /// Test the property 'SellerReturnItemId'
        /// </summary>
        [Fact]
        public void SellerReturnItemIdTest()
        {
            // TODO unit test for the property 'SellerReturnItemId'
        }

        /// <summary>
        /// Test the property 'SellerFulfillmentOrderItemId'
        /// </summary>
        [Fact]
        public void SellerFulfillmentOrderItemIdTest()
        {
            // TODO unit test for the property 'SellerFulfillmentOrderItemId'
        }

        /// <summary>
        /// Test the property 'AmazonShipmentId'
        /// </summary>
        [Fact]
        public void AmazonShipmentIdTest()
        {
            // TODO unit test for the property 'AmazonShipmentId'
        }

        /// <summary>
        /// Test the property 'SellerReturnReasonCode'
        /// </summary>
        [Fact]
        public void SellerReturnReasonCodeTest()
        {
            // TODO unit test for the property 'SellerReturnReasonCode'
        }

        /// <summary>
        /// Test the property 'ReturnComment'
        /// </summary>
        [Fact]
        public void ReturnCommentTest()
        {
            // TODO unit test for the property 'ReturnComment'
        }

        /// <summary>
        /// Test the property 'AmazonReturnReasonCode'
        /// </summary>
        [Fact]
        public void AmazonReturnReasonCodeTest()
        {
            // TODO unit test for the property 'AmazonReturnReasonCode'
        }

        /// <summary>
        /// Test the property 'Status'
        /// </summary>
        [Fact]
        public void StatusTest()
        {
            // TODO unit test for the property 'Status'
        }

        /// <summary>
        /// Test the property 'StatusChangedDate'
        /// </summary>
        [Fact]
        public void StatusChangedDateTest()
        {
            // TODO unit test for the property 'StatusChangedDate'
        }

        /// <summary>
        /// Test the property 'ReturnAuthorizationId'
        /// </summary>
        [Fact]
        public void ReturnAuthorizationIdTest()
        {
            // TODO unit test for the property 'ReturnAuthorizationId'
        }

        /// <summary>
        /// Test the property 'ReturnReceivedCondition'
        /// </summary>
        [Fact]
        public void ReturnReceivedConditionTest()
        {
            // TODO unit test for the property 'ReturnReceivedCondition'
        }

        /// <summary>
        /// Test the property 'FulfillmentCenterId'
        /// </summary>
        [Fact]
        public void FulfillmentCenterIdTest()
        {
            // TODO unit test for the property 'FulfillmentCenterId'
        }
    }
}
