// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleIotSiteWiseActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("putAssetPropertyValueEntries", required: true)]
        private InputList<Inputs.TopicRulePutAssetPropertyValueEntryArgs>? _putAssetPropertyValueEntries;
        public InputList<Inputs.TopicRulePutAssetPropertyValueEntryArgs> PutAssetPropertyValueEntries
        {
            get => _putAssetPropertyValueEntries ?? (_putAssetPropertyValueEntries = new InputList<Inputs.TopicRulePutAssetPropertyValueEntryArgs>());
            set => _putAssetPropertyValueEntries = value;
        }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public TopicRuleIotSiteWiseActionArgs()
        {
        }
        public static new TopicRuleIotSiteWiseActionArgs Empty => new TopicRuleIotSiteWiseActionArgs();
    }
}
