// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolInviteMessageTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("emailMessage")]
        public Input<string>? EmailMessage { get; set; }

        [Input("emailSubject")]
        public Input<string>? EmailSubject { get; set; }

        [Input("smsMessage")]
        public Input<string>? SmsMessage { get; set; }

        public UserPoolInviteMessageTemplateArgs()
        {
        }
        public static new UserPoolInviteMessageTemplateArgs Empty => new UserPoolInviteMessageTemplateArgs();
    }
}
