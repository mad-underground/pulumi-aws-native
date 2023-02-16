// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceCatalog::PortfolioPrincipalAssociation
 *
 * @deprecated PortfolioPrincipalAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class PortfolioPrincipalAssociation extends pulumi.CustomResource {
    /**
     * Get an existing PortfolioPrincipalAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PortfolioPrincipalAssociation {
        pulumi.log.warn("PortfolioPrincipalAssociation is deprecated: PortfolioPrincipalAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new PortfolioPrincipalAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:servicecatalog:PortfolioPrincipalAssociation';

    /**
     * Returns true if the given object is an instance of PortfolioPrincipalAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PortfolioPrincipalAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PortfolioPrincipalAssociation.__pulumiType;
    }

    public readonly acceptLanguage!: pulumi.Output<string | undefined>;
    public readonly portfolioId!: pulumi.Output<string>;
    public readonly principalARN!: pulumi.Output<string>;
    public readonly principalType!: pulumi.Output<string>;

    /**
     * Create a PortfolioPrincipalAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated PortfolioPrincipalAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: PortfolioPrincipalAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("PortfolioPrincipalAssociation is deprecated: PortfolioPrincipalAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.portfolioId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portfolioId'");
            }
            if ((!args || args.principalARN === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalARN'");
            }
            if ((!args || args.principalType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalType'");
            }
            resourceInputs["acceptLanguage"] = args ? args.acceptLanguage : undefined;
            resourceInputs["portfolioId"] = args ? args.portfolioId : undefined;
            resourceInputs["principalARN"] = args ? args.principalARN : undefined;
            resourceInputs["principalType"] = args ? args.principalType : undefined;
        } else {
            resourceInputs["acceptLanguage"] = undefined /*out*/;
            resourceInputs["portfolioId"] = undefined /*out*/;
            resourceInputs["principalARN"] = undefined /*out*/;
            resourceInputs["principalType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PortfolioPrincipalAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PortfolioPrincipalAssociation resource.
 */
export interface PortfolioPrincipalAssociationArgs {
    acceptLanguage?: pulumi.Input<string>;
    portfolioId: pulumi.Input<string>;
    principalARN: pulumi.Input<string>;
    principalType: pulumi.Input<string>;
}
