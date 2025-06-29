/*
 * The Selling Partner API for FBA inbound operations.
 *
 * The Selling Partner API for Fulfillment By Amazon (FBA) Inbound. The FBA Inbound API enables building inbound workflows to create, manage, and send shipments into Amazon's fulfillment network. The API has interoperability with the Send-to-Amazon user interface.
 *
 * The version of the OpenAPI document: 2024-03-20
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using software.amzn.spapi.Model.fulfillment.inbound.v2024_03_20;
using software.amzn.spapi.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace software.amzn.spapi.Test.Model
{
    /// <summary>
    ///  Class for testing SpdTrackingItem
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class SpdTrackingItemTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for SpdTrackingItem
        //private SpdTrackingItem instance;

        public SpdTrackingItemTests()
        {
            // TODO uncomment below to create an instance of SpdTrackingItem
            //instance = new SpdTrackingItem();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of SpdTrackingItem
        /// </summary>
        [Fact]
        public void SpdTrackingItemInstanceTest()
        {
            // TODO uncomment below to test "IsType" SpdTrackingItem
            //Assert.IsType<SpdTrackingItem>(instance);
        }

        /// <summary>
        /// Test the property 'BoxId'
        /// </summary>
        [Fact]
        public void BoxIdTest()
        {
            // TODO unit test for the property 'BoxId'
        }

        /// <summary>
        /// Test the property 'TrackingId'
        /// </summary>
        [Fact]
        public void TrackingIdTest()
        {
            // TODO unit test for the property 'TrackingId'
        }

        /// <summary>
        /// Test the property 'TrackingNumberValidationStatus'
        /// </summary>
        [Fact]
        public void TrackingNumberValidationStatusTest()
        {
            // TODO unit test for the property 'TrackingNumberValidationStatus'
        }
    }
}
