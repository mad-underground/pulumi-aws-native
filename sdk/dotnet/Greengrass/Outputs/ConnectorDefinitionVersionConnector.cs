// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Greengrass.Outputs
{

    [OutputType]
    public sealed class ConnectorDefinitionVersionConnector
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinitionversion-connector.html#cfn-greengrass-connectordefinitionversion-connector-connectorarn
        /// </summary>
        public readonly string ConnectorArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinitionversion-connector.html#cfn-greengrass-connectordefinitionversion-connector-id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-connectordefinitionversion-connector.html#cfn-greengrass-connectordefinitionversion-connector-parameters
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? Parameters;

        [OutputConstructor]
        private ConnectorDefinitionVersionConnector(
            string ConnectorArn,

            string Id,

            Union<System.Text.Json.JsonElement, string>? Parameters)
        {
            this.ConnectorArn = ConnectorArn;
            this.Id = Id;
            this.Parameters = Parameters;
        }
    }
}