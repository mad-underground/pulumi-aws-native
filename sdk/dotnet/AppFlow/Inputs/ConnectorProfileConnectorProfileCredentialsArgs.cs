// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html
    /// </summary>
    public sealed class ConnectorProfileConnectorProfileCredentialsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-amplitude
        /// </summary>
        [Input("amplitude")]
        public Input<Inputs.ConnectorProfileAmplitudeConnectorProfileCredentialsArgs>? Amplitude { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-datadog
        /// </summary>
        [Input("datadog")]
        public Input<Inputs.ConnectorProfileDatadogConnectorProfileCredentialsArgs>? Datadog { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-dynatrace
        /// </summary>
        [Input("dynatrace")]
        public Input<Inputs.ConnectorProfileDynatraceConnectorProfileCredentialsArgs>? Dynatrace { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-googleanalytics
        /// </summary>
        [Input("googleAnalytics")]
        public Input<Inputs.ConnectorProfileGoogleAnalyticsConnectorProfileCredentialsArgs>? GoogleAnalytics { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-infornexus
        /// </summary>
        [Input("inforNexus")]
        public Input<Inputs.ConnectorProfileInforNexusConnectorProfileCredentialsArgs>? InforNexus { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-marketo
        /// </summary>
        [Input("marketo")]
        public Input<Inputs.ConnectorProfileMarketoConnectorProfileCredentialsArgs>? Marketo { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-redshift
        /// </summary>
        [Input("redshift")]
        public Input<Inputs.ConnectorProfileRedshiftConnectorProfileCredentialsArgs>? Redshift { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-salesforce
        /// </summary>
        [Input("salesforce")]
        public Input<Inputs.ConnectorProfileSalesforceConnectorProfileCredentialsArgs>? Salesforce { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-servicenow
        /// </summary>
        [Input("serviceNow")]
        public Input<Inputs.ConnectorProfileServiceNowConnectorProfileCredentialsArgs>? ServiceNow { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-singular
        /// </summary>
        [Input("singular")]
        public Input<Inputs.ConnectorProfileSingularConnectorProfileCredentialsArgs>? Singular { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-slack
        /// </summary>
        [Input("slack")]
        public Input<Inputs.ConnectorProfileSlackConnectorProfileCredentialsArgs>? Slack { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-snowflake
        /// </summary>
        [Input("snowflake")]
        public Input<Inputs.ConnectorProfileSnowflakeConnectorProfileCredentialsArgs>? Snowflake { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-trendmicro
        /// </summary>
        [Input("trendmicro")]
        public Input<Inputs.ConnectorProfileTrendmicroConnectorProfileCredentialsArgs>? Trendmicro { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-veeva
        /// </summary>
        [Input("veeva")]
        public Input<Inputs.ConnectorProfileVeevaConnectorProfileCredentialsArgs>? Veeva { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html#cfn-appflow-connectorprofile-connectorprofilecredentials-zendesk
        /// </summary>
        [Input("zendesk")]
        public Input<Inputs.ConnectorProfileZendeskConnectorProfileCredentialsArgs>? Zendesk { get; set; }

        public ConnectorProfileConnectorProfileCredentialsArgs()
        {
        }
    }
}
