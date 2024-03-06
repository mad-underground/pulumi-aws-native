// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RDS::DBSecurityGroupIngress
 *
 * @deprecated DbSecurityGroupIngress is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DbSecurityGroupIngress extends pulumi.CustomResource {
    /**
     * Get an existing DbSecurityGroupIngress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DbSecurityGroupIngress {
        pulumi.log.warn("DbSecurityGroupIngress is deprecated: DbSecurityGroupIngress is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DbSecurityGroupIngress(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:DbSecurityGroupIngress';

    /**
     * Returns true if the given object is an instance of DbSecurityGroupIngress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbSecurityGroupIngress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbSecurityGroupIngress.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly cidrip!: pulumi.Output<string | undefined>;
    public readonly dbSecurityGroupName!: pulumi.Output<string>;
    public readonly ec2SecurityGroupId!: pulumi.Output<string | undefined>;
    public readonly ec2SecurityGroupName!: pulumi.Output<string | undefined>;
    public readonly ec2SecurityGroupOwnerId!: pulumi.Output<string | undefined>;

    /**
     * Create a DbSecurityGroupIngress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DbSecurityGroupIngress is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DbSecurityGroupIngressArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DbSecurityGroupIngress is deprecated: DbSecurityGroupIngress is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dbSecurityGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbSecurityGroupName'");
            }
            resourceInputs["cidrip"] = args ? args.cidrip : undefined;
            resourceInputs["dbSecurityGroupName"] = args ? args.dbSecurityGroupName : undefined;
            resourceInputs["ec2SecurityGroupId"] = args ? args.ec2SecurityGroupId : undefined;
            resourceInputs["ec2SecurityGroupName"] = args ? args.ec2SecurityGroupName : undefined;
            resourceInputs["ec2SecurityGroupOwnerId"] = args ? args.ec2SecurityGroupOwnerId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["cidrip"] = undefined /*out*/;
            resourceInputs["dbSecurityGroupName"] = undefined /*out*/;
            resourceInputs["ec2SecurityGroupId"] = undefined /*out*/;
            resourceInputs["ec2SecurityGroupName"] = undefined /*out*/;
            resourceInputs["ec2SecurityGroupOwnerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DbSecurityGroupIngress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DbSecurityGroupIngress resource.
 */
export interface DbSecurityGroupIngressArgs {
    cidrip?: pulumi.Input<string>;
    dbSecurityGroupName: pulumi.Input<string>;
    ec2SecurityGroupId?: pulumi.Input<string>;
    ec2SecurityGroupName?: pulumi.Input<string>;
    ec2SecurityGroupOwnerId?: pulumi.Input<string>;
}
