// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupActionDefinition
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-actiondefinition.html#cfn-networkfirewall-rulegroup-actiondefinition-publishmetricaction
        /// </summary>
        public readonly Outputs.RuleGroupPublishMetricAction? PublishMetricAction;

        [OutputConstructor]
        private RuleGroupActionDefinition(Outputs.RuleGroupPublishMetricAction? PublishMetricAction)
        {
            this.PublishMetricAction = PublishMetricAction;
        }
    }
}
