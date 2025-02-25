// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    /// <summary>
    /// Operation to be performed on provided source fields
    /// </summary>
    [OutputType]
    public sealed class FlowConnectorOperator
    {
        public readonly Pulumi.AwsNative.AppFlow.FlowAmplitudeConnectorOperator? Amplitude;
        public readonly Pulumi.AwsNative.AppFlow.FlowOperator? CustomConnector;
        public readonly Pulumi.AwsNative.AppFlow.FlowDatadogConnectorOperator? Datadog;
        public readonly Pulumi.AwsNative.AppFlow.FlowDynatraceConnectorOperator? Dynatrace;
        public readonly Pulumi.AwsNative.AppFlow.FlowGoogleAnalyticsConnectorOperator? GoogleAnalytics;
        public readonly Pulumi.AwsNative.AppFlow.FlowInforNexusConnectorOperator? InforNexus;
        public readonly Pulumi.AwsNative.AppFlow.FlowMarketoConnectorOperator? Marketo;
        public readonly Pulumi.AwsNative.AppFlow.FlowPardotConnectorOperator? Pardot;
        public readonly Pulumi.AwsNative.AppFlow.FlowS3ConnectorOperator? S3;
        public readonly Pulumi.AwsNative.AppFlow.FlowSalesforceConnectorOperator? Salesforce;
        public readonly Pulumi.AwsNative.AppFlow.FlowSapoDataConnectorOperator? SapoData;
        public readonly Pulumi.AwsNative.AppFlow.FlowServiceNowConnectorOperator? ServiceNow;
        public readonly Pulumi.AwsNative.AppFlow.FlowSingularConnectorOperator? Singular;
        public readonly Pulumi.AwsNative.AppFlow.FlowSlackConnectorOperator? Slack;
        public readonly Pulumi.AwsNative.AppFlow.FlowTrendmicroConnectorOperator? Trendmicro;
        public readonly Pulumi.AwsNative.AppFlow.FlowVeevaConnectorOperator? Veeva;
        public readonly Pulumi.AwsNative.AppFlow.FlowZendeskConnectorOperator? Zendesk;

        [OutputConstructor]
        private FlowConnectorOperator(
            Pulumi.AwsNative.AppFlow.FlowAmplitudeConnectorOperator? amplitude,

            Pulumi.AwsNative.AppFlow.FlowOperator? customConnector,

            Pulumi.AwsNative.AppFlow.FlowDatadogConnectorOperator? datadog,

            Pulumi.AwsNative.AppFlow.FlowDynatraceConnectorOperator? dynatrace,

            Pulumi.AwsNative.AppFlow.FlowGoogleAnalyticsConnectorOperator? googleAnalytics,

            Pulumi.AwsNative.AppFlow.FlowInforNexusConnectorOperator? inforNexus,

            Pulumi.AwsNative.AppFlow.FlowMarketoConnectorOperator? marketo,

            Pulumi.AwsNative.AppFlow.FlowPardotConnectorOperator? pardot,

            Pulumi.AwsNative.AppFlow.FlowS3ConnectorOperator? s3,

            Pulumi.AwsNative.AppFlow.FlowSalesforceConnectorOperator? salesforce,

            Pulumi.AwsNative.AppFlow.FlowSapoDataConnectorOperator? sapoData,

            Pulumi.AwsNative.AppFlow.FlowServiceNowConnectorOperator? serviceNow,

            Pulumi.AwsNative.AppFlow.FlowSingularConnectorOperator? singular,

            Pulumi.AwsNative.AppFlow.FlowSlackConnectorOperator? slack,

            Pulumi.AwsNative.AppFlow.FlowTrendmicroConnectorOperator? trendmicro,

            Pulumi.AwsNative.AppFlow.FlowVeevaConnectorOperator? veeva,

            Pulumi.AwsNative.AppFlow.FlowZendeskConnectorOperator? zendesk)
        {
            Amplitude = amplitude;
            CustomConnector = customConnector;
            Datadog = datadog;
            Dynatrace = dynatrace;
            GoogleAnalytics = googleAnalytics;
            InforNexus = inforNexus;
            Marketo = marketo;
            Pardot = pardot;
            S3 = s3;
            Salesforce = salesforce;
            SapoData = sapoData;
            ServiceNow = serviceNow;
            Singular = singular;
            Slack = slack;
            Trendmicro = trendmicro;
            Veeva = veeva;
            Zendesk = zendesk;
        }
    }
}
