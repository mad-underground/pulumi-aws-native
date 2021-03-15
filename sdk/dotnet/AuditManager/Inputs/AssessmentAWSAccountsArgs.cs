// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AuditManager.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-awsaccounts.html
    /// </summary>
    public sealed class AssessmentAWSAccountsArgs : Pulumi.ResourceArgs
    {
        [Input("aWSAccounts")]
        private InputList<Inputs.AssessmentAWSAccountArgs>? _aWSAccounts;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-awsaccounts.html#cfn-auditmanager-assessment-awsaccounts-awsaccounts
        /// </summary>
        public InputList<Inputs.AssessmentAWSAccountArgs> AWSAccounts
        {
            get => _aWSAccounts ?? (_aWSAccounts = new InputList<Inputs.AssessmentAWSAccountArgs>());
            set => _aWSAccounts = value;
        }

        public AssessmentAWSAccountsArgs()
        {
        }
    }
}
