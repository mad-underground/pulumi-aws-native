// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ApplicationInsights.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentconfiguration.html
    /// </summary>
    public sealed class ApplicationComponentConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentconfiguration.html#cfn-applicationinsights-application-componentconfiguration-configurationdetails
        /// </summary>
        [Input("ConfigurationDetails")]
        public Input<Inputs.ApplicationConfigurationDetailsArgs>? ConfigurationDetails { get; set; }

        [Input("SubComponentTypeConfigurations")]
        private InputList<Inputs.ApplicationSubComponentTypeConfigurationArgs>? _SubComponentTypeConfigurations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentconfiguration.html#cfn-applicationinsights-application-componentconfiguration-subcomponenttypeconfigurations
        /// </summary>
        public InputList<Inputs.ApplicationSubComponentTypeConfigurationArgs> SubComponentTypeConfigurations
        {
            get => _SubComponentTypeConfigurations ?? (_SubComponentTypeConfigurations = new InputList<Inputs.ApplicationSubComponentTypeConfigurationArgs>());
            set => _SubComponentTypeConfigurations = value;
        }

        public ApplicationComponentConfigurationArgs()
        {
        }
    }
}
