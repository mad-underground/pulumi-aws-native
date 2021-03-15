// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowSalesforceSourceProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-salesforcesourceproperties.html#cfn-appflow-flow-salesforcesourceproperties-enabledynamicfieldupdate
        /// </summary>
        public readonly bool? EnableDynamicFieldUpdate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-salesforcesourceproperties.html#cfn-appflow-flow-salesforcesourceproperties-includedeletedrecords
        /// </summary>
        public readonly bool? IncludeDeletedRecords;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-salesforcesourceproperties.html#cfn-appflow-flow-salesforcesourceproperties-object
        /// </summary>
        public readonly string Object;

        [OutputConstructor]
        private FlowSalesforceSourceProperties(
            bool? enableDynamicFieldUpdate,

            bool? includeDeletedRecords,

            string @object)
        {
            EnableDynamicFieldUpdate = enableDynamicFieldUpdate;
            IncludeDeletedRecords = includeDeletedRecords;
            Object = @object;
        }
    }
}
