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
    public sealed class FaqTagList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-faq-taglist.html#cfn-kendra-faq-taglist-taglist
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> TagList;

        [OutputConstructor]
        private FaqTagList(ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tagList)
        {
            TagList = tagList;
        }
    }
}
