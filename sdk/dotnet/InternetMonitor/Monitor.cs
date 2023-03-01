// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InternetMonitor
{
    /// <summary>
    /// Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application
    /// </summary>
    [AwsNativeResourceType("aws-native:internetmonitor:Monitor")]
    public partial class Monitor : global::Pulumi.CustomResource
    {
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("maxCityNetworksToMonitor")]
        public Output<int?> MaxCityNetworksToMonitor { get; private set; } = null!;

        [Output("modifiedAt")]
        public Output<string> ModifiedAt { get; private set; } = null!;

        [Output("monitorArn")]
        public Output<string> MonitorArn { get; private set; } = null!;

        [Output("monitorName")]
        public Output<string?> MonitorName { get; private set; } = null!;

        [Output("processingStatus")]
        public Output<Pulumi.AwsNative.InternetMonitor.MonitorProcessingStatusCode> ProcessingStatus { get; private set; } = null!;

        [Output("processingStatusInfo")]
        public Output<string> ProcessingStatusInfo { get; private set; } = null!;

        [Output("resources")]
        public Output<ImmutableArray<string>> Resources { get; private set; } = null!;

        [Output("resourcesToAdd")]
        public Output<ImmutableArray<string>> ResourcesToAdd { get; private set; } = null!;

        [Output("resourcesToRemove")]
        public Output<ImmutableArray<string>> ResourcesToRemove { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.AwsNative.InternetMonitor.MonitorConfigState?> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.MonitorTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Monitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Monitor(string name, MonitorArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:internetmonitor:Monitor", name, args ?? new MonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Monitor(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:internetmonitor:Monitor", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Monitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Monitor Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Monitor(name, id, options);
        }
    }

    public sealed class MonitorArgs : global::Pulumi.ResourceArgs
    {
        [Input("maxCityNetworksToMonitor")]
        public Input<int>? MaxCityNetworksToMonitor { get; set; }

        [Input("monitorName")]
        public Input<string>? MonitorName { get; set; }

        [Input("resources")]
        private InputList<string>? _resources;
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        [Input("resourcesToAdd")]
        private InputList<string>? _resourcesToAdd;
        public InputList<string> ResourcesToAdd
        {
            get => _resourcesToAdd ?? (_resourcesToAdd = new InputList<string>());
            set => _resourcesToAdd = value;
        }

        [Input("resourcesToRemove")]
        private InputList<string>? _resourcesToRemove;
        public InputList<string> ResourcesToRemove
        {
            get => _resourcesToRemove ?? (_resourcesToRemove = new InputList<string>());
            set => _resourcesToRemove = value;
        }

        [Input("status")]
        public Input<Pulumi.AwsNative.InternetMonitor.MonitorConfigState>? Status { get; set; }

        [Input("tags")]
        private InputList<Inputs.MonitorTagArgs>? _tags;
        public InputList<Inputs.MonitorTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.MonitorTagArgs>());
            set => _tags = value;
        }

        public MonitorArgs()
        {
        }
        public static new MonitorArgs Empty => new MonitorArgs();
    }
}
