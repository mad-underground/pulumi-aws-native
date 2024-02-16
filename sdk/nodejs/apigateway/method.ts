// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::ApiGateway::Method`` resource creates API Gateway methods that define the parameters and body that clients must send in their requests.
 */
export class Method extends pulumi.CustomResource {
    /**
     * Get an existing Method resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Method {
        return new Method(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigateway:Method';

    /**
     * Returns true if the given object is an instance of Method.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Method {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Method.__pulumiType;
    }

    /**
     * A boolean flag specifying whether a valid ApiKey is required to invoke this method.
     */
    public readonly apiKeyRequired!: pulumi.Output<boolean | undefined>;
    /**
     * A list of authorization scopes configured on the method. The scopes are used with a ``COGNITO_USER_POOLS`` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.
     */
    public readonly authorizationScopes!: pulumi.Output<string[] | undefined>;
    /**
     * The method's authorization type. This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html) in the *API Gateway API Reference*.
     *   If you specify the ``AuthorizerId`` property, specify ``CUSTOM`` or ``COGNITO_USER_POOLS`` for this property.
     */
    public readonly authorizationType!: pulumi.Output<enums.apigateway.MethodAuthorizationType | undefined>;
    /**
     * The identifier of an authorizer to use on this method. The method's authorization type must be ``CUSTOM`` or ``COGNITO_USER_POOLS``.
     */
    public readonly authorizerId!: pulumi.Output<string | undefined>;
    /**
     * The method's HTTP verb.
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * Represents an ``HTTP``, ``HTTP_PROXY``, ``AWS``, ``AWS_PROXY``, or Mock integration.
     */
    public readonly integration!: pulumi.Output<outputs.apigateway.MethodIntegration | undefined>;
    /**
     * Gets a method response associated with a given HTTP status code.
     */
    public readonly methodResponses!: pulumi.Output<outputs.apigateway.MethodResponse[] | undefined>;
    /**
     * A human-friendly operation identifier for the method. For example, you can assign the ``operationName`` of ``ListPets`` for the ``GET /pets`` method in the ``PetStore`` example.
     */
    public readonly operationName!: pulumi.Output<string | undefined>;
    /**
     * A key-value map specifying data schemas, represented by Model resources, (as the mapped value) of the request payloads of given content types (as the mapping key).
     */
    public readonly requestModels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A key-value map defining required or optional method request parameters that can be accepted by API Gateway. A key is a method request parameter name matching the pattern of ``method.request.{location}.{name}``, where ``location`` is ``querystring``, ``path``, or ``header`` and ``name`` is a valid and unique parameter name. The value associated with the key is a Boolean flag indicating whether the parameter is required (``true``) or optional (``false``). The method request parameter names defined here are available in Integration to be mapped to integration request parameters or templates.
     */
    public readonly requestParameters!: pulumi.Output<{[key: string]: boolean} | undefined>;
    /**
     * The identifier of a RequestValidator for request validation.
     */
    public readonly requestValidatorId!: pulumi.Output<string | undefined>;
    /**
     * The Resource identifier for the MethodResponse resource.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The string identifier of the associated RestApi.
     */
    public readonly restApiId!: pulumi.Output<string>;

    /**
     * Create a Method resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MethodArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.httpMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'httpMethod'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.restApiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApiId'");
            }
            resourceInputs["apiKeyRequired"] = args ? args.apiKeyRequired : undefined;
            resourceInputs["authorizationScopes"] = args ? args.authorizationScopes : undefined;
            resourceInputs["authorizationType"] = args ? args.authorizationType : undefined;
            resourceInputs["authorizerId"] = args ? args.authorizerId : undefined;
            resourceInputs["httpMethod"] = args ? args.httpMethod : undefined;
            resourceInputs["integration"] = args ? args.integration : undefined;
            resourceInputs["methodResponses"] = args ? args.methodResponses : undefined;
            resourceInputs["operationName"] = args ? args.operationName : undefined;
            resourceInputs["requestModels"] = args ? args.requestModels : undefined;
            resourceInputs["requestParameters"] = args ? args.requestParameters : undefined;
            resourceInputs["requestValidatorId"] = args ? args.requestValidatorId : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["restApiId"] = args ? args.restApiId : undefined;
        } else {
            resourceInputs["apiKeyRequired"] = undefined /*out*/;
            resourceInputs["authorizationScopes"] = undefined /*out*/;
            resourceInputs["authorizationType"] = undefined /*out*/;
            resourceInputs["authorizerId"] = undefined /*out*/;
            resourceInputs["httpMethod"] = undefined /*out*/;
            resourceInputs["integration"] = undefined /*out*/;
            resourceInputs["methodResponses"] = undefined /*out*/;
            resourceInputs["operationName"] = undefined /*out*/;
            resourceInputs["requestModels"] = undefined /*out*/;
            resourceInputs["requestParameters"] = undefined /*out*/;
            resourceInputs["requestValidatorId"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["restApiId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["httpMethod", "resourceId", "restApiId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Method.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Method resource.
 */
export interface MethodArgs {
    /**
     * A boolean flag specifying whether a valid ApiKey is required to invoke this method.
     */
    apiKeyRequired?: pulumi.Input<boolean>;
    /**
     * A list of authorization scopes configured on the method. The scopes are used with a ``COGNITO_USER_POOLS`` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.
     */
    authorizationScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The method's authorization type. This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html) in the *API Gateway API Reference*.
     *   If you specify the ``AuthorizerId`` property, specify ``CUSTOM`` or ``COGNITO_USER_POOLS`` for this property.
     */
    authorizationType?: pulumi.Input<enums.apigateway.MethodAuthorizationType>;
    /**
     * The identifier of an authorizer to use on this method. The method's authorization type must be ``CUSTOM`` or ``COGNITO_USER_POOLS``.
     */
    authorizerId?: pulumi.Input<string>;
    /**
     * The method's HTTP verb.
     */
    httpMethod: pulumi.Input<string>;
    /**
     * Represents an ``HTTP``, ``HTTP_PROXY``, ``AWS``, ``AWS_PROXY``, or Mock integration.
     */
    integration?: pulumi.Input<inputs.apigateway.MethodIntegrationArgs>;
    /**
     * Gets a method response associated with a given HTTP status code.
     */
    methodResponses?: pulumi.Input<pulumi.Input<inputs.apigateway.MethodResponseArgs>[]>;
    /**
     * A human-friendly operation identifier for the method. For example, you can assign the ``operationName`` of ``ListPets`` for the ``GET /pets`` method in the ``PetStore`` example.
     */
    operationName?: pulumi.Input<string>;
    /**
     * A key-value map specifying data schemas, represented by Model resources, (as the mapped value) of the request payloads of given content types (as the mapping key).
     */
    requestModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A key-value map defining required or optional method request parameters that can be accepted by API Gateway. A key is a method request parameter name matching the pattern of ``method.request.{location}.{name}``, where ``location`` is ``querystring``, ``path``, or ``header`` and ``name`` is a valid and unique parameter name. The value associated with the key is a Boolean flag indicating whether the parameter is required (``true``) or optional (``false``). The method request parameter names defined here are available in Integration to be mapped to integration request parameters or templates.
     */
    requestParameters?: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
    /**
     * The identifier of a RequestValidator for request validation.
     */
    requestValidatorId?: pulumi.Input<string>;
    /**
     * The Resource identifier for the MethodResponse resource.
     */
    resourceId: pulumi.Input<string>;
    /**
     * The string identifier of the associated RestApi.
     */
    restApiId: pulumi.Input<string>;
}
