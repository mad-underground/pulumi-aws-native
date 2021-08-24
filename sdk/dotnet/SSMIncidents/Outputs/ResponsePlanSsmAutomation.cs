// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSMIncidents.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html
    /// </summary>
    [OutputType]
    public sealed class ResponsePlanSsmAutomation
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-documentname
        /// </summary>
        public readonly string DocumentName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-documentversion
        /// </summary>
        public readonly string? DocumentVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-parameters
        /// </summary>
        public readonly ImmutableArray<Outputs.ResponsePlanSsmParameter> Parameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-rolearn
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-targetaccount
        /// </summary>
        public readonly string? TargetAccount;

        [OutputConstructor]
        private ResponsePlanSsmAutomation(
            string documentName,

            string? documentVersion,

            ImmutableArray<Outputs.ResponsePlanSsmParameter> parameters,

            string roleArn,

            string? targetAccount)
        {
            DocumentName = documentName;
            DocumentVersion = documentVersion;
            Parameters = parameters;
            RoleArn = roleArn;
            TargetAccount = targetAccount;
        }
    }
}