// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets
{
    public static class GetBudget
    {
        /// <summary>
        /// Resource Type definition for AWS::Budgets::Budget
        /// </summary>
        public static Task<GetBudgetResult> InvokeAsync(GetBudgetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBudgetResult>("aws-native:budgets:getBudget", args ?? new GetBudgetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Budgets::Budget
        /// </summary>
        public static Output<GetBudgetResult> Invoke(GetBudgetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBudgetResult>("aws-native:budgets:getBudget", args ?? new GetBudgetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBudgetArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetBudgetArgs()
        {
        }
    }

    public sealed class GetBudgetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetBudgetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBudgetResult
    {
        public readonly Outputs.BudgetData? BudgetValue;
        public readonly string? Id;

        [OutputConstructor]
        private GetBudgetResult(
            Outputs.BudgetData? budget,

            string? id)
        {
            BudgetValue = budget;
            Id = id;
        }
    }
}