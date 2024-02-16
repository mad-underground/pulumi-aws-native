// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::ApiGatewayV2::DomainName`` resource specifies a custom domain name for your API in Amazon API Gateway (API Gateway).
 *  You can use a custom domain name to provide a URL that's more intuitive and easier to recall. For more information about using custom domain names, see [Set up Custom Domain Name for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html) in the *API Gateway Developer Guide*.
 */
export class DomainName extends pulumi.CustomResource {
    /**
     * Get an existing DomainName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DomainName {
        return new DomainName(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigatewayv2:DomainName';

    /**
     * Returns true if the given object is an instance of DomainName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainName.__pulumiType;
    }

    /**
     * The custom domain name for your API in Amazon API Gateway. Uppercase letters are not supported.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The domain name configurations.
     */
    public readonly domainNameConfigurations!: pulumi.Output<outputs.apigatewayv2.DomainNameConfiguration[] | undefined>;
    /**
     * The mutual TLS authentication configuration for a custom domain name.
     */
    public readonly mutualTlsAuthentication!: pulumi.Output<outputs.apigatewayv2.DomainNameMutualTlsAuthentication | undefined>;
    public /*out*/ readonly regionalDomainName!: pulumi.Output<string>;
    public /*out*/ readonly regionalHostedZoneId!: pulumi.Output<string>;
    /**
     * The collection of tags associated with a domain name.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a DomainName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainNameArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["domainNameConfigurations"] = args ? args.domainNameConfigurations : undefined;
            resourceInputs["mutualTlsAuthentication"] = args ? args.mutualTlsAuthentication : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["regionalDomainName"] = undefined /*out*/;
            resourceInputs["regionalHostedZoneId"] = undefined /*out*/;
        } else {
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["domainNameConfigurations"] = undefined /*out*/;
            resourceInputs["mutualTlsAuthentication"] = undefined /*out*/;
            resourceInputs["regionalDomainName"] = undefined /*out*/;
            resourceInputs["regionalHostedZoneId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["domainName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DomainName.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DomainName resource.
 */
export interface DomainNameArgs {
    /**
     * The custom domain name for your API in Amazon API Gateway. Uppercase letters are not supported.
     */
    domainName: pulumi.Input<string>;
    /**
     * The domain name configurations.
     */
    domainNameConfigurations?: pulumi.Input<pulumi.Input<inputs.apigatewayv2.DomainNameConfigurationArgs>[]>;
    /**
     * The mutual TLS authentication configuration for a custom domain name.
     */
    mutualTlsAuthentication?: pulumi.Input<inputs.apigatewayv2.DomainNameMutualTlsAuthenticationArgs>;
    /**
     * The collection of tags associated with a domain name.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
