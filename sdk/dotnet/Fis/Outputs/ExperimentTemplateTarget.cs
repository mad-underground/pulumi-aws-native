// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fis.Outputs
{

    /// <summary>
    /// Specifies a target for an experiment.
    /// </summary>
    [OutputType]
    public sealed class ExperimentTemplateTarget
    {
        public readonly ImmutableArray<Outputs.ExperimentTemplateTargetFilter> Filters;
        public readonly ImmutableDictionary<string, string>? Parameters;
        public readonly ImmutableArray<string> ResourceArns;
        public readonly ImmutableDictionary<string, string>? ResourceTags;
        public readonly string ResourceType;
        public readonly string SelectionMode;

        [OutputConstructor]
        private ExperimentTemplateTarget(
            ImmutableArray<Outputs.ExperimentTemplateTargetFilter> filters,

            ImmutableDictionary<string, string>? parameters,

            ImmutableArray<string> resourceArns,

            ImmutableDictionary<string, string>? resourceTags,

            string resourceType,

            string selectionMode)
        {
            Filters = filters;
            Parameters = parameters;
            ResourceArns = resourceArns;
            ResourceTags = resourceTags;
            ResourceType = resourceType;
            SelectionMode = selectionMode;
        }
    }
}
