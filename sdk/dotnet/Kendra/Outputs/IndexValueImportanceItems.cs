// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class IndexValueImportanceItems
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-valueimportanceitems.html#cfn-kendra-index-valueimportanceitems-valueimportanceitems
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexValueImportanceItem> ValueImportanceItems;

        [OutputConstructor]
        private IndexValueImportanceItems(ImmutableArray<Outputs.IndexValueImportanceItem> valueImportanceItems)
        {
            ValueImportanceItems = valueImportanceItems;
        }
    }
}
