// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    public sealed class RuleNetworkConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("awsVpcConfiguration")]
        public Input<Inputs.RuleAwsVpcConfigurationArgs>? AwsVpcConfiguration { get; set; }

        public RuleNetworkConfigurationArgs()
        {
        }
    }
}