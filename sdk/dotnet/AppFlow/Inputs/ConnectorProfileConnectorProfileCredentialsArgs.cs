// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppFlow.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html
    /// </summary>
    public sealed class ConnectorProfileConnectorProfileCredentialsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-amplitude
        /// </summary>
        [Input("Amplitude")]
        public Input<Inputs.ConnectorProfileAmplitudeConnectorProfileCredentialsArgs>? Amplitude { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-datadog
        /// </summary>
        [Input("Datadog")]
        public Input<Inputs.ConnectorProfileDatadogConnectorProfileCredentialsArgs>? Datadog { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-dynatrace
        /// </summary>
        [Input("Dynatrace")]
        public Input<Inputs.ConnectorProfileDynatraceConnectorProfileCredentialsArgs>? Dynatrace { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-googleanalytics
        /// </summary>
        [Input("GoogleAnalytics")]
        public Input<Inputs.ConnectorProfileGoogleAnalyticsConnectorProfileCredentialsArgs>? GoogleAnalytics { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-infornexus
        /// </summary>
        [Input("InforNexus")]
        public Input<Inputs.ConnectorProfileInforNexusConnectorProfileCredentialsArgs>? InforNexus { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-marketo
        /// </summary>
        [Input("Marketo")]
        public Input<Inputs.ConnectorProfileMarketoConnectorProfileCredentialsArgs>? Marketo { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-redshift
        /// </summary>
        [Input("Redshift")]
        public Input<Inputs.ConnectorProfileRedshiftConnectorProfileCredentialsArgs>? Redshift { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-salesforce
        /// </summary>
        [Input("Salesforce")]
        public Input<Inputs.ConnectorProfileSalesforceConnectorProfileCredentialsArgs>? Salesforce { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-servicenow
        /// </summary>
        [Input("ServiceNow")]
        public Input<Inputs.ConnectorProfileServiceNowConnectorProfileCredentialsArgs>? ServiceNow { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-singular
        /// </summary>
        [Input("Singular")]
        public Input<Inputs.ConnectorProfileSingularConnectorProfileCredentialsArgs>? Singular { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-slack
        /// </summary>
        [Input("Slack")]
        public Input<Inputs.ConnectorProfileSlackConnectorProfileCredentialsArgs>? Slack { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-snowflake
        /// </summary>
        [Input("Snowflake")]
        public Input<Inputs.ConnectorProfileSnowflakeConnectorProfileCredentialsArgs>? Snowflake { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-trendmicro
        /// </summary>
        [Input("Trendmicro")]
        public Input<Inputs.ConnectorProfileTrendmicroConnectorProfileCredentialsArgs>? Trendmicro { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-veeva
        /// </summary>
        [Input("Veeva")]
        public Input<Inputs.ConnectorProfileVeevaConnectorProfileCredentialsArgs>? Veeva { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-zendesk
        /// </summary>
        [Input("Zendesk")]
        public Input<Inputs.ConnectorProfileZendeskConnectorProfileCredentialsArgs>? Zendesk { get; set; }

        public ConnectorProfileConnectorProfileCredentialsArgs()
        {
        }
    }
}
