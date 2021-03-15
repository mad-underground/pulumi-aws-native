// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Outputs
{

    [OutputType]
    public sealed class AssociationParameterValues
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html#cfn-ssm-association-parametervalues-parametervalues
        /// </summary>
        public readonly ImmutableArray<string> ParameterValues;

        [OutputConstructor]
        private AssociationParameterValues(ImmutableArray<string> parameterValues)
        {
            ParameterValues = parameterValues;
        }
    }
}
