// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolUserPoolAddOns
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html#cfn-cognito-userpool-userpooladdons-advancedsecuritymode
        /// </summary>
        public readonly string? AdvancedSecurityMode;

        [OutputConstructor]
        private UserPoolUserPoolAddOns(string? AdvancedSecurityMode)
        {
            this.AdvancedSecurityMode = AdvancedSecurityMode;
        }
    }
}