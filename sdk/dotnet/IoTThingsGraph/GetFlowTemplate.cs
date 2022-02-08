// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTThingsGraph
{
    public static class GetFlowTemplate
    {
        /// <summary>
        /// Resource Type definition for AWS::IoTThingsGraph::FlowTemplate
        /// </summary>
        public static Task<GetFlowTemplateResult> InvokeAsync(GetFlowTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFlowTemplateResult>("aws-native:iotthingsgraph:getFlowTemplate", args ?? new GetFlowTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IoTThingsGraph::FlowTemplate
        /// </summary>
        public static Output<GetFlowTemplateResult> Invoke(GetFlowTemplateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFlowTemplateResult>("aws-native:iotthingsgraph:getFlowTemplate", args ?? new GetFlowTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFlowTemplateArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetFlowTemplateArgs()
        {
        }
    }

    public sealed class GetFlowTemplateInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetFlowTemplateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFlowTemplateResult
    {
        public readonly double? CompatibleNamespaceVersion;
        public readonly Outputs.FlowTemplateDefinitionDocument? Definition;
        public readonly string? Id;

        [OutputConstructor]
        private GetFlowTemplateResult(
            double? compatibleNamespaceVersion,

            Outputs.FlowTemplateDefinitionDocument? definition,

            string? id)
        {
            CompatibleNamespaceVersion = compatibleNamespaceVersion;
            Definition = definition;
            Id = id;
        }
    }
}
