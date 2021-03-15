// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResourceGroups.Outputs
{

    [OutputType]
    public sealed class GroupResourceQuery
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-resourcequery.html#cfn-resourcegroups-group-resourcequery-query
        /// </summary>
        public readonly Outputs.GroupQuery? Query;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-resourcequery.html#cfn-resourcegroups-group-resourcequery-type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GroupResourceQuery(
            Outputs.GroupQuery? query,

            string? type)
        {
            Query = query;
            Type = type;
        }
    }
}
