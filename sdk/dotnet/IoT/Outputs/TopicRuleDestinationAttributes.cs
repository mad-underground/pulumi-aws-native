// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoT.Outputs
{

    [OutputType]
    public sealed class TopicRuleDestinationAttributes
    {
        public readonly string Arn;
        public readonly string StatusReason;

        [OutputConstructor]
        private TopicRuleDestinationAttributes(
            string Arn,

            string StatusReason)
        {
            this.Arn = Arn;
            this.StatusReason = StatusReason;
        }
    }
}