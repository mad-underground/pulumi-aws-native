// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rum
{
    /// <summary>
    /// Resource Type definition for AWS::RUM::AppMonitor
    /// </summary>
    [AwsNativeResourceType("aws-native:rum:AppMonitor")]
    public partial class AppMonitor : global::Pulumi.CustomResource
    {
        [Output("appMonitorConfiguration")]
        public Output<Outputs.AppMonitorConfiguration?> AppMonitorConfiguration { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the new app monitor.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("customEvents")]
        public Output<Outputs.AppMonitorCustomEvents?> CustomEvents { get; private set; } = null!;

        /// <summary>
        /// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
        /// </summary>
        [Output("cwLogEnabled")]
        public Output<bool?> CwLogEnabled { get; private set; } = null!;

        /// <summary>
        /// The top-level internet domain name for which your application has administrative authority.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// A name for the app monitor
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AppMonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppMonitor(string name, AppMonitorArgs args, CustomResourceOptions? options = null)
            : base("aws-native:rum:AppMonitor", name, args ?? new AppMonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppMonitor(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:rum:AppMonitor", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppMonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppMonitor Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppMonitor(name, id, options);
        }
    }

    public sealed class AppMonitorArgs : global::Pulumi.ResourceArgs
    {
        [Input("appMonitorConfiguration")]
        public Input<Inputs.AppMonitorConfigurationArgs>? AppMonitorConfiguration { get; set; }

        [Input("customEvents")]
        public Input<Inputs.AppMonitorCustomEventsArgs>? CustomEvents { get; set; }

        /// <summary>
        /// Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
        /// </summary>
        [Input("cwLogEnabled")]
        public Input<bool>? CwLogEnabled { get; set; }

        /// <summary>
        /// The top-level internet domain name for which your application has administrative authority.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// A name for the app monitor
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public AppMonitorArgs()
        {
        }
        public static new AppMonitorArgs Empty => new AppMonitorArgs();
    }
}
