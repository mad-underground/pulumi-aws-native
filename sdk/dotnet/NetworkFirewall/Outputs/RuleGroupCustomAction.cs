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
    public sealed class RuleGroupCustomAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-customaction.html#cfn-networkfirewall-rulegroup-customaction-actiondefinition
        /// </summary>
        public readonly Outputs.RuleGroupActionDefinition ActionDefinition;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-customaction.html#cfn-networkfirewall-rulegroup-customaction-actionname
        /// </summary>
        public readonly string ActionName;

        [OutputConstructor]
        private RuleGroupCustomAction(
            Outputs.RuleGroupActionDefinition ActionDefinition,

            string ActionName)
        {
            this.ActionDefinition = ActionDefinition;
            this.ActionName = ActionName;
        }
    }
}
