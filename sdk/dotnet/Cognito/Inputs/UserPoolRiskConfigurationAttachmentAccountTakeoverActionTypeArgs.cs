// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("eventAction", required: true)]
        public Input<string> EventAction { get; set; } = null!;

        [Input("notify", required: true)]
        public Input<bool> Notify { get; set; } = null!;

        public UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeArgs()
        {
        }
        public static new UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeArgs Empty => new UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeArgs();
    }
}
