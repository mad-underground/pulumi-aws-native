// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer.Inputs
{

    public sealed class ServerWorkflowDetailArgs : global::Pulumi.ResourceArgs
    {
        [Input("executionRole", required: true)]
        public Input<string> ExecutionRole { get; set; } = null!;

        [Input("workflowId", required: true)]
        public Input<string> WorkflowId { get; set; } = null!;

        public ServerWorkflowDetailArgs()
        {
        }
        public static new ServerWorkflowDetailArgs Empty => new ServerWorkflowDetailArgs();
    }
}
