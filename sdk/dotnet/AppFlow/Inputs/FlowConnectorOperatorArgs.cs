// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    /// <summary>
    /// Operation to be performed on provided source fields
    /// </summary>
    public sealed class FlowConnectorOperatorArgs : global::Pulumi.ResourceArgs
    {
        [Input("amplitude")]
        public Input<Pulumi.AwsNative.AppFlow.FlowAmplitudeConnectorOperator>? Amplitude { get; set; }

        [Input("customConnector")]
        public Input<Pulumi.AwsNative.AppFlow.FlowOperator>? CustomConnector { get; set; }

        [Input("datadog")]
        public Input<Pulumi.AwsNative.AppFlow.FlowDatadogConnectorOperator>? Datadog { get; set; }

        [Input("dynatrace")]
        public Input<Pulumi.AwsNative.AppFlow.FlowDynatraceConnectorOperator>? Dynatrace { get; set; }

        [Input("googleAnalytics")]
        public Input<Pulumi.AwsNative.AppFlow.FlowGoogleAnalyticsConnectorOperator>? GoogleAnalytics { get; set; }

        [Input("inforNexus")]
        public Input<Pulumi.AwsNative.AppFlow.FlowInforNexusConnectorOperator>? InforNexus { get; set; }

        [Input("marketo")]
        public Input<Pulumi.AwsNative.AppFlow.FlowMarketoConnectorOperator>? Marketo { get; set; }

        [Input("pardot")]
        public Input<Pulumi.AwsNative.AppFlow.FlowPardotConnectorOperator>? Pardot { get; set; }

        [Input("s3")]
        public Input<Pulumi.AwsNative.AppFlow.FlowS3ConnectorOperator>? S3 { get; set; }

        [Input("sAPOData")]
        public Input<Pulumi.AwsNative.AppFlow.FlowSAPODataConnectorOperator>? SAPOData { get; set; }

        [Input("salesforce")]
        public Input<Pulumi.AwsNative.AppFlow.FlowSalesforceConnectorOperator>? Salesforce { get; set; }

        [Input("serviceNow")]
        public Input<Pulumi.AwsNative.AppFlow.FlowServiceNowConnectorOperator>? ServiceNow { get; set; }

        [Input("singular")]
        public Input<Pulumi.AwsNative.AppFlow.FlowSingularConnectorOperator>? Singular { get; set; }

        [Input("slack")]
        public Input<Pulumi.AwsNative.AppFlow.FlowSlackConnectorOperator>? Slack { get; set; }

        [Input("trendmicro")]
        public Input<Pulumi.AwsNative.AppFlow.FlowTrendmicroConnectorOperator>? Trendmicro { get; set; }

        [Input("veeva")]
        public Input<Pulumi.AwsNative.AppFlow.FlowVeevaConnectorOperator>? Veeva { get; set; }

        [Input("zendesk")]
        public Input<Pulumi.AwsNative.AppFlow.FlowZendeskConnectorOperator>? Zendesk { get; set; }

        public FlowConnectorOperatorArgs()
        {
        }
        public static new FlowConnectorOperatorArgs Empty => new FlowConnectorOperatorArgs();
    }
}
