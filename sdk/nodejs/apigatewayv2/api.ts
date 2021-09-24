// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGatewayV2::Api
 *
 * @deprecated Api is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Api {
        pulumi.log.warn("Api is deprecated: Api is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Api(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigatewayv2:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    public /*out*/ readonly apiEndpoint!: pulumi.Output<string>;
    public readonly apiKeySelectionExpression!: pulumi.Output<string | undefined>;
    public readonly basePath!: pulumi.Output<string | undefined>;
    public readonly body!: pulumi.Output<any | undefined>;
    public readonly bodyS3Location!: pulumi.Output<outputs.apigatewayv2.ApiBodyS3Location | undefined>;
    public readonly corsConfiguration!: pulumi.Output<outputs.apigatewayv2.ApiCors | undefined>;
    public readonly credentialsArn!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly disableExecuteApiEndpoint!: pulumi.Output<boolean | undefined>;
    public readonly disableSchemaValidation!: pulumi.Output<boolean | undefined>;
    public readonly failOnWarnings!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly protocolType!: pulumi.Output<string | undefined>;
    public readonly routeKey!: pulumi.Output<string | undefined>;
    public readonly routeSelectionExpression!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<any | undefined>;
    public readonly target!: pulumi.Output<string | undefined>;
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Api is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: ApiArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Api is deprecated: Api is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["apiKeySelectionExpression"] = args ? args.apiKeySelectionExpression : undefined;
            inputs["basePath"] = args ? args.basePath : undefined;
            inputs["body"] = args ? args.body : undefined;
            inputs["bodyS3Location"] = args ? args.bodyS3Location : undefined;
            inputs["corsConfiguration"] = args ? args.corsConfiguration : undefined;
            inputs["credentialsArn"] = args ? args.credentialsArn : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disableExecuteApiEndpoint"] = args ? args.disableExecuteApiEndpoint : undefined;
            inputs["disableSchemaValidation"] = args ? args.disableSchemaValidation : undefined;
            inputs["failOnWarnings"] = args ? args.failOnWarnings : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protocolType"] = args ? args.protocolType : undefined;
            inputs["routeKey"] = args ? args.routeKey : undefined;
            inputs["routeSelectionExpression"] = args ? args.routeSelectionExpression : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["apiEndpoint"] = undefined /*out*/;
        } else {
            inputs["apiEndpoint"] = undefined /*out*/;
            inputs["apiKeySelectionExpression"] = undefined /*out*/;
            inputs["basePath"] = undefined /*out*/;
            inputs["body"] = undefined /*out*/;
            inputs["bodyS3Location"] = undefined /*out*/;
            inputs["corsConfiguration"] = undefined /*out*/;
            inputs["credentialsArn"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["disableExecuteApiEndpoint"] = undefined /*out*/;
            inputs["disableSchemaValidation"] = undefined /*out*/;
            inputs["failOnWarnings"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["protocolType"] = undefined /*out*/;
            inputs["routeKey"] = undefined /*out*/;
            inputs["routeSelectionExpression"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["target"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Api.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    apiKeySelectionExpression?: pulumi.Input<string>;
    basePath?: pulumi.Input<string>;
    body?: any;
    bodyS3Location?: pulumi.Input<inputs.apigatewayv2.ApiBodyS3LocationArgs>;
    corsConfiguration?: pulumi.Input<inputs.apigatewayv2.ApiCorsArgs>;
    credentialsArn?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    disableExecuteApiEndpoint?: pulumi.Input<boolean>;
    disableSchemaValidation?: pulumi.Input<boolean>;
    failOnWarnings?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    protocolType?: pulumi.Input<string>;
    routeKey?: pulumi.Input<string>;
    routeSelectionExpression?: pulumi.Input<string>;
    tags?: any;
    target?: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}