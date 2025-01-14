// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmIncidents.Outputs
{

    /// <summary>
    /// Value of the dynamic parameter to set when starting the SSM automation document.
    /// </summary>
    [OutputType]
    public sealed class ResponsePlanDynamicSsmParameterValue
    {
        public readonly Pulumi.AwsNative.SsmIncidents.ResponsePlanVariableType? Variable;

        [OutputConstructor]
        private ResponsePlanDynamicSsmParameterValue(Pulumi.AwsNative.SsmIncidents.ResponsePlanVariableType? variable)
        {
            Variable = variable;
        }
    }
}
