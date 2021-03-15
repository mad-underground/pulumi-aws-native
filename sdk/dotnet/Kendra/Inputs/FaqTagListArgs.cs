// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-faq-taglist.html
    /// </summary>
    public sealed class FaqTagListArgs : Pulumi.ResourceArgs
    {
        [Input("tagList")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tagList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-faq-taglist.html#cfn-kendra-faq-taglist-taglist
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> TagList
        {
            get => _tagList ?? (_tagList = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tagList = value;
        }

        public FaqTagListArgs()
        {
        }
    }
}
