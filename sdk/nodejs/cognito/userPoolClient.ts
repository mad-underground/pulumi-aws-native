// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPoolClient
 */
export class UserPoolClient extends pulumi.CustomResource {
    /**
     * Get an existing UserPoolClient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserPoolClient {
        return new UserPoolClient(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cognito:UserPoolClient';

    /**
     * Returns true if the given object is an instance of UserPoolClient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPoolClient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPoolClient.__pulumiType;
    }

    public readonly accessTokenValidity!: pulumi.Output<number | undefined>;
    public readonly allowedOAuthFlows!: pulumi.Output<string[] | undefined>;
    public readonly allowedOAuthFlowsUserPoolClient!: pulumi.Output<boolean | undefined>;
    public readonly allowedOAuthScopes!: pulumi.Output<string[] | undefined>;
    public readonly analyticsConfiguration!: pulumi.Output<outputs.cognito.UserPoolClientAnalyticsConfiguration | undefined>;
    public readonly authSessionValidity!: pulumi.Output<number | undefined>;
    public readonly callbackUrls!: pulumi.Output<string[] | undefined>;
    public readonly clientName!: pulumi.Output<string | undefined>;
    public /*out*/ readonly clientSecret!: pulumi.Output<string>;
    public readonly defaultRedirectUri!: pulumi.Output<string | undefined>;
    public readonly enablePropagateAdditionalUserContextData!: pulumi.Output<boolean | undefined>;
    public readonly enableTokenRevocation!: pulumi.Output<boolean | undefined>;
    public readonly explicitAuthFlows!: pulumi.Output<string[] | undefined>;
    public readonly generateSecret!: pulumi.Output<boolean | undefined>;
    public readonly idTokenValidity!: pulumi.Output<number | undefined>;
    public readonly logoutUrls!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly preventUserExistenceErrors!: pulumi.Output<string | undefined>;
    public readonly readAttributes!: pulumi.Output<string[] | undefined>;
    public readonly refreshTokenValidity!: pulumi.Output<number | undefined>;
    public readonly supportedIdentityProviders!: pulumi.Output<string[] | undefined>;
    public readonly tokenValidityUnits!: pulumi.Output<outputs.cognito.UserPoolClientTokenValidityUnits | undefined>;
    public readonly userPoolId!: pulumi.Output<string>;
    public readonly writeAttributes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a UserPoolClient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPoolClientArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.userPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPoolId'");
            }
            resourceInputs["accessTokenValidity"] = args ? args.accessTokenValidity : undefined;
            resourceInputs["allowedOAuthFlows"] = args ? args.allowedOAuthFlows : undefined;
            resourceInputs["allowedOAuthFlowsUserPoolClient"] = args ? args.allowedOAuthFlowsUserPoolClient : undefined;
            resourceInputs["allowedOAuthScopes"] = args ? args.allowedOAuthScopes : undefined;
            resourceInputs["analyticsConfiguration"] = args ? args.analyticsConfiguration : undefined;
            resourceInputs["authSessionValidity"] = args ? args.authSessionValidity : undefined;
            resourceInputs["callbackUrls"] = args ? args.callbackUrls : undefined;
            resourceInputs["clientName"] = args ? args.clientName : undefined;
            resourceInputs["defaultRedirectUri"] = args ? args.defaultRedirectUri : undefined;
            resourceInputs["enablePropagateAdditionalUserContextData"] = args ? args.enablePropagateAdditionalUserContextData : undefined;
            resourceInputs["enableTokenRevocation"] = args ? args.enableTokenRevocation : undefined;
            resourceInputs["explicitAuthFlows"] = args ? args.explicitAuthFlows : undefined;
            resourceInputs["generateSecret"] = args ? args.generateSecret : undefined;
            resourceInputs["idTokenValidity"] = args ? args.idTokenValidity : undefined;
            resourceInputs["logoutUrls"] = args ? args.logoutUrls : undefined;
            resourceInputs["preventUserExistenceErrors"] = args ? args.preventUserExistenceErrors : undefined;
            resourceInputs["readAttributes"] = args ? args.readAttributes : undefined;
            resourceInputs["refreshTokenValidity"] = args ? args.refreshTokenValidity : undefined;
            resourceInputs["supportedIdentityProviders"] = args ? args.supportedIdentityProviders : undefined;
            resourceInputs["tokenValidityUnits"] = args ? args.tokenValidityUnits : undefined;
            resourceInputs["userPoolId"] = args ? args.userPoolId : undefined;
            resourceInputs["writeAttributes"] = args ? args.writeAttributes : undefined;
            resourceInputs["clientSecret"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["accessTokenValidity"] = undefined /*out*/;
            resourceInputs["allowedOAuthFlows"] = undefined /*out*/;
            resourceInputs["allowedOAuthFlowsUserPoolClient"] = undefined /*out*/;
            resourceInputs["allowedOAuthScopes"] = undefined /*out*/;
            resourceInputs["analyticsConfiguration"] = undefined /*out*/;
            resourceInputs["authSessionValidity"] = undefined /*out*/;
            resourceInputs["callbackUrls"] = undefined /*out*/;
            resourceInputs["clientName"] = undefined /*out*/;
            resourceInputs["clientSecret"] = undefined /*out*/;
            resourceInputs["defaultRedirectUri"] = undefined /*out*/;
            resourceInputs["enablePropagateAdditionalUserContextData"] = undefined /*out*/;
            resourceInputs["enableTokenRevocation"] = undefined /*out*/;
            resourceInputs["explicitAuthFlows"] = undefined /*out*/;
            resourceInputs["generateSecret"] = undefined /*out*/;
            resourceInputs["idTokenValidity"] = undefined /*out*/;
            resourceInputs["logoutUrls"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["preventUserExistenceErrors"] = undefined /*out*/;
            resourceInputs["readAttributes"] = undefined /*out*/;
            resourceInputs["refreshTokenValidity"] = undefined /*out*/;
            resourceInputs["supportedIdentityProviders"] = undefined /*out*/;
            resourceInputs["tokenValidityUnits"] = undefined /*out*/;
            resourceInputs["userPoolId"] = undefined /*out*/;
            resourceInputs["writeAttributes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["generateSecret", "userPoolId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(UserPoolClient.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserPoolClient resource.
 */
export interface UserPoolClientArgs {
    accessTokenValidity?: pulumi.Input<number>;
    allowedOAuthFlows?: pulumi.Input<pulumi.Input<string>[]>;
    allowedOAuthFlowsUserPoolClient?: pulumi.Input<boolean>;
    allowedOAuthScopes?: pulumi.Input<pulumi.Input<string>[]>;
    analyticsConfiguration?: pulumi.Input<inputs.cognito.UserPoolClientAnalyticsConfigurationArgs>;
    authSessionValidity?: pulumi.Input<number>;
    callbackUrls?: pulumi.Input<pulumi.Input<string>[]>;
    clientName?: pulumi.Input<string>;
    defaultRedirectUri?: pulumi.Input<string>;
    enablePropagateAdditionalUserContextData?: pulumi.Input<boolean>;
    enableTokenRevocation?: pulumi.Input<boolean>;
    explicitAuthFlows?: pulumi.Input<pulumi.Input<string>[]>;
    generateSecret?: pulumi.Input<boolean>;
    idTokenValidity?: pulumi.Input<number>;
    logoutUrls?: pulumi.Input<pulumi.Input<string>[]>;
    preventUserExistenceErrors?: pulumi.Input<string>;
    readAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    refreshTokenValidity?: pulumi.Input<number>;
    supportedIdentityProviders?: pulumi.Input<pulumi.Input<string>[]>;
    tokenValidityUnits?: pulumi.Input<inputs.cognito.UserPoolClientTokenValidityUnitsArgs>;
    userPoolId: pulumi.Input<string>;
    writeAttributes?: pulumi.Input<pulumi.Input<string>[]>;
}
