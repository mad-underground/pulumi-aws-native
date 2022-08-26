// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * A resource schema representing a Lake Formation Permission.
 */
export class PrincipalPermissions extends pulumi.CustomResource {
    /**
     * Get an existing PrincipalPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrincipalPermissions {
        return new PrincipalPermissions(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lakeformation:PrincipalPermissions';

    /**
     * Returns true if the given object is an instance of PrincipalPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrincipalPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrincipalPermissions.__pulumiType;
    }

    public readonly catalog!: pulumi.Output<string | undefined>;
    public readonly permissions!: pulumi.Output<enums.lakeformation.PrincipalPermissionsPermission[]>;
    public readonly permissionsWithGrantOption!: pulumi.Output<enums.lakeformation.PrincipalPermissionsPermission[]>;
    public readonly principal!: pulumi.Output<outputs.lakeformation.PrincipalPermissionsDataLakePrincipal>;
    public /*out*/ readonly principalIdentifier!: pulumi.Output<string>;
    public readonly resource!: pulumi.Output<outputs.lakeformation.PrincipalPermissionsResource>;
    public /*out*/ readonly resourceIdentifier!: pulumi.Output<string>;

    /**
     * Create a PrincipalPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrincipalPermissionsArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.permissionsWithGrantOption === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissionsWithGrantOption'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            resourceInputs["catalog"] = args ? args.catalog : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["permissionsWithGrantOption"] = args ? args.permissionsWithGrantOption : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["resource"] = args ? args.resource : undefined;
            resourceInputs["principalIdentifier"] = undefined /*out*/;
            resourceInputs["resourceIdentifier"] = undefined /*out*/;
        } else {
            resourceInputs["catalog"] = undefined /*out*/;
            resourceInputs["permissions"] = undefined /*out*/;
            resourceInputs["permissionsWithGrantOption"] = undefined /*out*/;
            resourceInputs["principal"] = undefined /*out*/;
            resourceInputs["principalIdentifier"] = undefined /*out*/;
            resourceInputs["resource"] = undefined /*out*/;
            resourceInputs["resourceIdentifier"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrincipalPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrincipalPermissions resource.
 */
export interface PrincipalPermissionsArgs {
    catalog?: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<enums.lakeformation.PrincipalPermissionsPermission>[]>;
    permissionsWithGrantOption: pulumi.Input<pulumi.Input<enums.lakeformation.PrincipalPermissionsPermission>[]>;
    principal: pulumi.Input<inputs.lakeformation.PrincipalPermissionsDataLakePrincipalArgs>;
    resource: pulumi.Input<inputs.lakeformation.PrincipalPermissionsResourceArgs>;
}