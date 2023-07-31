// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// Configures how to use the Account creation fraud prevention managed rule group in the web ACL
    /// </summary>
    public sealed class WebACLAWSManagedRulesACFPRuleSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("creationPath", required: true)]
        public Input<string> CreationPath { get; set; } = null!;

        [Input("enableRegexInPath")]
        public Input<bool>? EnableRegexInPath { get; set; }

        [Input("registrationPagePath", required: true)]
        public Input<string> RegistrationPagePath { get; set; } = null!;

        [Input("requestInspection", required: true)]
        public Input<Inputs.WebACLRequestInspectionACFPArgs> RequestInspection { get; set; } = null!;

        [Input("responseInspection")]
        public Input<Inputs.WebACLResponseInspectionArgs>? ResponseInspection { get; set; }

        public WebACLAWSManagedRulesACFPRuleSetArgs()
        {
        }
        public static new WebACLAWSManagedRulesACFPRuleSetArgs Empty => new WebACLAWSManagedRulesACFPRuleSetArgs();
    }
}
