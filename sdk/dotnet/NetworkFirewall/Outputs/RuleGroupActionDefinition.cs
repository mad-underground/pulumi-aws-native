// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupActionDefinition
    {
        public readonly Outputs.RuleGroupPublishMetricAction? PublishMetricAction;

        [OutputConstructor]
        private RuleGroupActionDefinition(Outputs.RuleGroupPublishMetricAction? publishMetricAction)
        {
            PublishMetricAction = publishMetricAction;
        }
    }
}
