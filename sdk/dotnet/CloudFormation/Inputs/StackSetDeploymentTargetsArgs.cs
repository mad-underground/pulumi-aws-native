// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFormation.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html
    /// </summary>
    public sealed class StackSetDeploymentTargetsArgs : Pulumi.ResourceArgs
    {
        [Input("Accounts")]
        private InputList<string>? _Accounts;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-accounts
        /// </summary>
        public InputList<string> Accounts
        {
            get => _Accounts ?? (_Accounts = new InputList<string>());
            set => _Accounts = value;
        }

        [Input("OrganizationalUnitIds")]
        private InputList<string>? _OrganizationalUnitIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-organizationalunitids
        /// </summary>
        public InputList<string> OrganizationalUnitIds
        {
            get => _OrganizationalUnitIds ?? (_OrganizationalUnitIds = new InputList<string>());
            set => _OrganizationalUnitIds = value;
        }

        public StackSetDeploymentTargetsArgs()
        {
        }
    }
}