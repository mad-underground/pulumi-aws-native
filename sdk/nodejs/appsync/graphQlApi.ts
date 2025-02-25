// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::GraphQLApi
 *
 * @deprecated GraphQlApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class GraphQlApi extends pulumi.CustomResource {
    /**
     * Get an existing GraphQlApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GraphQlApi {
        pulumi.log.warn("GraphQlApi is deprecated: GraphQlApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new GraphQlApi(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appsync:GraphQlApi';

    /**
     * Returns true if the given object is an instance of GraphQlApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GraphQlApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GraphQlApi.__pulumiType;
    }

    public readonly additionalAuthenticationProviders!: pulumi.Output<outputs.appsync.GraphQlApiAdditionalAuthenticationProvider[] | undefined>;
    public /*out*/ readonly apiId!: pulumi.Output<string>;
    public readonly apiType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly authenticationType!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly enhancedMetricsConfig!: pulumi.Output<outputs.appsync.GraphQlApiEnhancedMetricsConfig | undefined>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::AppSync::GraphQLApi` for more information about the expected schema for this property.
     */
    public readonly environmentVariables!: pulumi.Output<any | undefined>;
    public /*out*/ readonly graphQlDns!: pulumi.Output<string>;
    public /*out*/ readonly graphQlEndpointArn!: pulumi.Output<string>;
    public /*out*/ readonly graphQlUrl!: pulumi.Output<string>;
    public readonly introspectionConfig!: pulumi.Output<string | undefined>;
    public readonly lambdaAuthorizerConfig!: pulumi.Output<outputs.appsync.GraphQlApiLambdaAuthorizerConfig | undefined>;
    public readonly logConfig!: pulumi.Output<outputs.appsync.GraphQlApiLogConfig | undefined>;
    public readonly mergedApiExecutionRoleArn!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly openIdConnectConfig!: pulumi.Output<outputs.appsync.GraphQlApiOpenIdConnectConfig | undefined>;
    public readonly ownerContact!: pulumi.Output<string | undefined>;
    public readonly queryDepthLimit!: pulumi.Output<number | undefined>;
    public /*out*/ readonly realtimeDns!: pulumi.Output<string>;
    public /*out*/ readonly realtimeUrl!: pulumi.Output<string>;
    public readonly resolverCountLimit!: pulumi.Output<number | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public readonly userPoolConfig!: pulumi.Output<outputs.appsync.GraphQlApiUserPoolConfig | undefined>;
    public readonly visibility!: pulumi.Output<string | undefined>;
    public readonly xrayEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GraphQlApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated GraphQlApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: GraphQlApiArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GraphQlApi is deprecated: GraphQlApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.authenticationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authenticationType'");
            }
            resourceInputs["additionalAuthenticationProviders"] = args ? args.additionalAuthenticationProviders : undefined;
            resourceInputs["apiType"] = args ? args.apiType : undefined;
            resourceInputs["authenticationType"] = args ? args.authenticationType : undefined;
            resourceInputs["enhancedMetricsConfig"] = args ? args.enhancedMetricsConfig : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["introspectionConfig"] = args ? args.introspectionConfig : undefined;
            resourceInputs["lambdaAuthorizerConfig"] = args ? args.lambdaAuthorizerConfig : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["mergedApiExecutionRoleArn"] = args ? args.mergedApiExecutionRoleArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["openIdConnectConfig"] = args ? args.openIdConnectConfig : undefined;
            resourceInputs["ownerContact"] = args ? args.ownerContact : undefined;
            resourceInputs["queryDepthLimit"] = args ? args.queryDepthLimit : undefined;
            resourceInputs["resolverCountLimit"] = args ? args.resolverCountLimit : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userPoolConfig"] = args ? args.userPoolConfig : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["xrayEnabled"] = args ? args.xrayEnabled : undefined;
            resourceInputs["apiId"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["graphQlDns"] = undefined /*out*/;
            resourceInputs["graphQlEndpointArn"] = undefined /*out*/;
            resourceInputs["graphQlUrl"] = undefined /*out*/;
            resourceInputs["realtimeDns"] = undefined /*out*/;
            resourceInputs["realtimeUrl"] = undefined /*out*/;
        } else {
            resourceInputs["additionalAuthenticationProviders"] = undefined /*out*/;
            resourceInputs["apiId"] = undefined /*out*/;
            resourceInputs["apiType"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["authenticationType"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["enhancedMetricsConfig"] = undefined /*out*/;
            resourceInputs["environmentVariables"] = undefined /*out*/;
            resourceInputs["graphQlDns"] = undefined /*out*/;
            resourceInputs["graphQlEndpointArn"] = undefined /*out*/;
            resourceInputs["graphQlUrl"] = undefined /*out*/;
            resourceInputs["introspectionConfig"] = undefined /*out*/;
            resourceInputs["lambdaAuthorizerConfig"] = undefined /*out*/;
            resourceInputs["logConfig"] = undefined /*out*/;
            resourceInputs["mergedApiExecutionRoleArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["openIdConnectConfig"] = undefined /*out*/;
            resourceInputs["ownerContact"] = undefined /*out*/;
            resourceInputs["queryDepthLimit"] = undefined /*out*/;
            resourceInputs["realtimeDns"] = undefined /*out*/;
            resourceInputs["realtimeUrl"] = undefined /*out*/;
            resourceInputs["resolverCountLimit"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["userPoolConfig"] = undefined /*out*/;
            resourceInputs["visibility"] = undefined /*out*/;
            resourceInputs["xrayEnabled"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GraphQlApi.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GraphQlApi resource.
 */
export interface GraphQlApiArgs {
    additionalAuthenticationProviders?: pulumi.Input<pulumi.Input<inputs.appsync.GraphQlApiAdditionalAuthenticationProviderArgs>[]>;
    apiType?: pulumi.Input<string>;
    authenticationType: pulumi.Input<string>;
    enhancedMetricsConfig?: pulumi.Input<inputs.appsync.GraphQlApiEnhancedMetricsConfigArgs>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::AppSync::GraphQLApi` for more information about the expected schema for this property.
     */
    environmentVariables?: any;
    introspectionConfig?: pulumi.Input<string>;
    lambdaAuthorizerConfig?: pulumi.Input<inputs.appsync.GraphQlApiLambdaAuthorizerConfigArgs>;
    logConfig?: pulumi.Input<inputs.appsync.GraphQlApiLogConfigArgs>;
    mergedApiExecutionRoleArn?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    openIdConnectConfig?: pulumi.Input<inputs.appsync.GraphQlApiOpenIdConnectConfigArgs>;
    ownerContact?: pulumi.Input<string>;
    queryDepthLimit?: pulumi.Input<number>;
    resolverCountLimit?: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    userPoolConfig?: pulumi.Input<inputs.appsync.GraphQlApiUserPoolConfigArgs>;
    visibility?: pulumi.Input<string>;
    xrayEnabled?: pulumi.Input<boolean>;
}
