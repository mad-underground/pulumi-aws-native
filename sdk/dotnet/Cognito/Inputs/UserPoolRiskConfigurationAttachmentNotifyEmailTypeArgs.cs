// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolRiskConfigurationAttachmentNotifyEmailTypeArgs : Pulumi.ResourceArgs
    {
        [Input("htmlBody")]
        public Input<string>? HtmlBody { get; set; }

        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        [Input("textBody")]
        public Input<string>? TextBody { get; set; }

        public UserPoolRiskConfigurationAttachmentNotifyEmailTypeArgs()
        {
        }
    }
}