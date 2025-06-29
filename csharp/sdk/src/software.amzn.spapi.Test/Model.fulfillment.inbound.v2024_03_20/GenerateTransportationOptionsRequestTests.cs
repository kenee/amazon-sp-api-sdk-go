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
    ///  Class for testing GenerateTransportationOptionsRequest
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class GenerateTransportationOptionsRequestTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for GenerateTransportationOptionsRequest
        //private GenerateTransportationOptionsRequest instance;

        public GenerateTransportationOptionsRequestTests()
        {
            // TODO uncomment below to create an instance of GenerateTransportationOptionsRequest
            //instance = new GenerateTransportationOptionsRequest();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of GenerateTransportationOptionsRequest
        /// </summary>
        [Fact]
        public void GenerateTransportationOptionsRequestInstanceTest()
        {
            // TODO uncomment below to test "IsType" GenerateTransportationOptionsRequest
            //Assert.IsType<GenerateTransportationOptionsRequest>(instance);
        }

        /// <summary>
        /// Test the property 'PlacementOptionId'
        /// </summary>
        [Fact]
        public void PlacementOptionIdTest()
        {
            // TODO unit test for the property 'PlacementOptionId'
        }

        /// <summary>
        /// Test the property 'ShipmentTransportationConfigurations'
        /// </summary>
        [Fact]
        public void ShipmentTransportationConfigurationsTest()
        {
            // TODO unit test for the property 'ShipmentTransportationConfigurations'
        }
    }
}
