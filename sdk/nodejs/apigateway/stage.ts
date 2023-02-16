// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::Stage
 */
export class Stage extends pulumi.CustomResource {
    /**
     * Get an existing Stage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Stage {
        return new Stage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigateway:Stage';

    /**
     * Returns true if the given object is an instance of Stage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stage.__pulumiType;
    }

    /**
     * Specifies settings for logging access in this stage.
     */
    public readonly accessLogSetting!: pulumi.Output<outputs.apigateway.StageAccessLogSetting | undefined>;
    /**
     * Indicates whether cache clustering is enabled for the stage.
     */
    public readonly cacheClusterEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The stage's cache cluster size.
     */
    public readonly cacheClusterSize!: pulumi.Output<string | undefined>;
    /**
     * Specifies settings for the canary deployment in this stage.
     */
    public readonly canarySetting!: pulumi.Output<outputs.apigateway.StageCanarySetting | undefined>;
    /**
     * The ID of the client certificate that API Gateway uses to call your integration endpoints in the stage. 
     */
    public readonly clientCertificateId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the deployment that the stage is associated with. This parameter is required to create a stage. 
     */
    public readonly deploymentId!: pulumi.Output<string | undefined>;
    /**
     * A description of the stage.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The version ID of the API documentation snapshot.
     */
    public readonly documentationVersion!: pulumi.Output<string | undefined>;
    /**
     * Settings for all methods in the stage.
     */
    public readonly methodSettings!: pulumi.Output<outputs.apigateway.StageMethodSetting[] | undefined>;
    /**
     * The ID of the RestApi resource that you're deploying with this stage.
     */
    public readonly restApiId!: pulumi.Output<string>;
    /**
     * The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).
     */
    public readonly stageName!: pulumi.Output<string | undefined>;
    /**
     * An array of arbitrary tags (key-value pairs) to associate with the stage.
     */
    public readonly tags!: pulumi.Output<outputs.apigateway.StageTag[] | undefined>;
    /**
     * Specifies whether active X-Ray tracing is enabled for this stage.
     */
    public readonly tracingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.
     */
    public readonly variables!: pulumi.Output<any | undefined>;

    /**
     * Create a Stage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.restApiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApiId'");
            }
            resourceInputs["accessLogSetting"] = args ? args.accessLogSetting : undefined;
            resourceInputs["cacheClusterEnabled"] = args ? args.cacheClusterEnabled : undefined;
            resourceInputs["cacheClusterSize"] = args ? args.cacheClusterSize : undefined;
            resourceInputs["canarySetting"] = args ? args.canarySetting : undefined;
            resourceInputs["clientCertificateId"] = args ? args.clientCertificateId : undefined;
            resourceInputs["deploymentId"] = args ? args.deploymentId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["documentationVersion"] = args ? args.documentationVersion : undefined;
            resourceInputs["methodSettings"] = args ? args.methodSettings : undefined;
            resourceInputs["restApiId"] = args ? args.restApiId : undefined;
            resourceInputs["stageName"] = args ? args.stageName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tracingEnabled"] = args ? args.tracingEnabled : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
        } else {
            resourceInputs["accessLogSetting"] = undefined /*out*/;
            resourceInputs["cacheClusterEnabled"] = undefined /*out*/;
            resourceInputs["cacheClusterSize"] = undefined /*out*/;
            resourceInputs["canarySetting"] = undefined /*out*/;
            resourceInputs["clientCertificateId"] = undefined /*out*/;
            resourceInputs["deploymentId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["documentationVersion"] = undefined /*out*/;
            resourceInputs["methodSettings"] = undefined /*out*/;
            resourceInputs["restApiId"] = undefined /*out*/;
            resourceInputs["stageName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["tracingEnabled"] = undefined /*out*/;
            resourceInputs["variables"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Stage resource.
 */
export interface StageArgs {
    /**
     * Specifies settings for logging access in this stage.
     */
    accessLogSetting?: pulumi.Input<inputs.apigateway.StageAccessLogSettingArgs>;
    /**
     * Indicates whether cache clustering is enabled for the stage.
     */
    cacheClusterEnabled?: pulumi.Input<boolean>;
    /**
     * The stage's cache cluster size.
     */
    cacheClusterSize?: pulumi.Input<string>;
    /**
     * Specifies settings for the canary deployment in this stage.
     */
    canarySetting?: pulumi.Input<inputs.apigateway.StageCanarySettingArgs>;
    /**
     * The ID of the client certificate that API Gateway uses to call your integration endpoints in the stage. 
     */
    clientCertificateId?: pulumi.Input<string>;
    /**
     * The ID of the deployment that the stage is associated with. This parameter is required to create a stage. 
     */
    deploymentId?: pulumi.Input<string>;
    /**
     * A description of the stage.
     */
    description?: pulumi.Input<string>;
    /**
     * The version ID of the API documentation snapshot.
     */
    documentationVersion?: pulumi.Input<string>;
    /**
     * Settings for all methods in the stage.
     */
    methodSettings?: pulumi.Input<pulumi.Input<inputs.apigateway.StageMethodSettingArgs>[]>;
    /**
     * The ID of the RestApi resource that you're deploying with this stage.
     */
    restApiId: pulumi.Input<string>;
    /**
     * The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).
     */
    stageName?: pulumi.Input<string>;
    /**
     * An array of arbitrary tags (key-value pairs) to associate with the stage.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.apigateway.StageTagArgs>[]>;
    /**
     * Specifies whether active X-Ray tracing is enabled for this stage.
     */
    tracingEnabled?: pulumi.Input<boolean>;
    /**
     * A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.
     */
    variables?: any;
}
