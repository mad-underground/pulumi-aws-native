// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Budgets.Outputs
{

    [OutputType]
    public sealed class BudgetProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-budget
        /// </summary>
        public readonly Outputs.BudgetBudgetData Budget;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-notificationswithsubscribers
        /// </summary>
        public readonly ImmutableArray<Outputs.BudgetNotificationWithSubscribers> NotificationsWithSubscribers;

        [OutputConstructor]
        private BudgetProperties(
            Outputs.BudgetBudgetData Budget,

            ImmutableArray<Outputs.BudgetNotificationWithSubscribers> NotificationsWithSubscribers)
        {
            this.Budget = Budget;
            this.NotificationsWithSubscribers = NotificationsWithSubscribers;
        }
    }
}