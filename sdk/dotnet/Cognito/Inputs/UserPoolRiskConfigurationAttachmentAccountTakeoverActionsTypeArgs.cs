// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolRiskConfigurationAttachmentAccountTakeoverActionsTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("highAction")]
        public Input<Inputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeArgs>? HighAction { get; set; }

        [Input("lowAction")]
        public Input<Inputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeArgs>? LowAction { get; set; }

        [Input("mediumAction")]
        public Input<Inputs.UserPoolRiskConfigurationAttachmentAccountTakeoverActionTypeArgs>? MediumAction { get; set; }

        public UserPoolRiskConfigurationAttachmentAccountTakeoverActionsTypeArgs()
        {
        }
        public static new UserPoolRiskConfigurationAttachmentAccountTakeoverActionsTypeArgs Empty => new UserPoolRiskConfigurationAttachmentAccountTakeoverActionsTypeArgs();
    }
}
