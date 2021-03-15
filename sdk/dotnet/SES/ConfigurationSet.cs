// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html
    /// </summary>
    [AwsNativeResourceType("aws-native:SES:ConfigurationSet")]
    public partial class ConfigurationSet : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationSet(string name, ConfigurationSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:SES:ConfigurationSet", name, args ?? new ConfigurationSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:SES:ConfigurationSet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigurationSet(name, id, options);
        }
    }

    public sealed class ConfigurationSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationset.html#cfn-ses-configurationset-name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ConfigurationSetArgs()
        {
        }
    }
}
