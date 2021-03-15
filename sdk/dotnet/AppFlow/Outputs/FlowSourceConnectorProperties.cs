// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowSourceConnectorProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-amplitude
        /// </summary>
        public readonly Outputs.FlowAmplitudeSourceProperties? Amplitude;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-datadog
        /// </summary>
        public readonly Outputs.FlowDatadogSourceProperties? Datadog;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-dynatrace
        /// </summary>
        public readonly Outputs.FlowDynatraceSourceProperties? Dynatrace;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-googleanalytics
        /// </summary>
        public readonly Outputs.FlowGoogleAnalyticsSourceProperties? GoogleAnalytics;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-infornexus
        /// </summary>
        public readonly Outputs.FlowInforNexusSourceProperties? InforNexus;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-marketo
        /// </summary>
        public readonly Outputs.FlowMarketoSourceProperties? Marketo;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-s3
        /// </summary>
        public readonly Outputs.FlowS3SourceProperties? S3;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-salesforce
        /// </summary>
        public readonly Outputs.FlowSalesforceSourceProperties? Salesforce;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-servicenow
        /// </summary>
        public readonly Outputs.FlowServiceNowSourceProperties? ServiceNow;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-singular
        /// </summary>
        public readonly Outputs.FlowSingularSourceProperties? Singular;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-slack
        /// </summary>
        public readonly Outputs.FlowSlackSourceProperties? Slack;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-trendmicro
        /// </summary>
        public readonly Outputs.FlowTrendmicroSourceProperties? Trendmicro;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-veeva
        /// </summary>
        public readonly Outputs.FlowVeevaSourceProperties? Veeva;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-zendesk
        /// </summary>
        public readonly Outputs.FlowZendeskSourceProperties? Zendesk;

        [OutputConstructor]
        private FlowSourceConnectorProperties(
            Outputs.FlowAmplitudeSourceProperties? amplitude,

            Outputs.FlowDatadogSourceProperties? datadog,

            Outputs.FlowDynatraceSourceProperties? dynatrace,

            Outputs.FlowGoogleAnalyticsSourceProperties? googleAnalytics,

            Outputs.FlowInforNexusSourceProperties? inforNexus,

            Outputs.FlowMarketoSourceProperties? marketo,

            Outputs.FlowS3SourceProperties? s3,

            Outputs.FlowSalesforceSourceProperties? salesforce,

            Outputs.FlowServiceNowSourceProperties? serviceNow,

            Outputs.FlowSingularSourceProperties? singular,

            Outputs.FlowSlackSourceProperties? slack,

            Outputs.FlowTrendmicroSourceProperties? trendmicro,

            Outputs.FlowVeevaSourceProperties? veeva,

            Outputs.FlowZendeskSourceProperties? zendesk)
        {
            Amplitude = amplitude;
            Datadog = datadog;
            Dynatrace = dynatrace;
            GoogleAnalytics = googleAnalytics;
            InforNexus = inforNexus;
            Marketo = marketo;
            S3 = s3;
            Salesforce = salesforce;
            ServiceNow = serviceNow;
            Singular = singular;
            Slack = slack;
            Trendmicro = trendmicro;
            Veeva = veeva;
            Zendesk = zendesk;
        }
    }
}
