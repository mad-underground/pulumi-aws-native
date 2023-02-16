// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DAX::Cluster
 *
 * @deprecated Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        pulumi.log.warn("Cluster is deprecated: Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dax:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly availabilityZones!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly clusterDiscoveryEndpoint!: pulumi.Output<string>;
    public /*out*/ readonly clusterDiscoveryEndpointURL!: pulumi.Output<string>;
    public readonly clusterEndpointEncryptionType!: pulumi.Output<string | undefined>;
    public readonly clusterName!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly iAMRoleARN!: pulumi.Output<string>;
    public readonly nodeType!: pulumi.Output<string>;
    public readonly notificationTopicARN!: pulumi.Output<string | undefined>;
    public readonly parameterGroupName!: pulumi.Output<string | undefined>;
    public readonly preferredMaintenanceWindow!: pulumi.Output<string | undefined>;
    public readonly replicationFactor!: pulumi.Output<number>;
    public readonly sSESpecification!: pulumi.Output<outputs.dax.ClusterSSESpecification | undefined>;
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    public readonly subnetGroupName!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<any | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Cluster is deprecated: Cluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.iAMRoleARN === undefined) && !opts.urn) {
                throw new Error("Missing required property 'iAMRoleARN'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            if ((!args || args.replicationFactor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicationFactor'");
            }
            resourceInputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            resourceInputs["clusterEndpointEncryptionType"] = args ? args.clusterEndpointEncryptionType : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["iAMRoleARN"] = args ? args.iAMRoleARN : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["notificationTopicARN"] = args ? args.notificationTopicARN : undefined;
            resourceInputs["parameterGroupName"] = args ? args.parameterGroupName : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["replicationFactor"] = args ? args.replicationFactor : undefined;
            resourceInputs["sSESpecification"] = args ? args.sSESpecification : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["subnetGroupName"] = args ? args.subnetGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["clusterDiscoveryEndpoint"] = undefined /*out*/;
            resourceInputs["clusterDiscoveryEndpointURL"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["availabilityZones"] = undefined /*out*/;
            resourceInputs["clusterDiscoveryEndpoint"] = undefined /*out*/;
            resourceInputs["clusterDiscoveryEndpointURL"] = undefined /*out*/;
            resourceInputs["clusterEndpointEncryptionType"] = undefined /*out*/;
            resourceInputs["clusterName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["iAMRoleARN"] = undefined /*out*/;
            resourceInputs["nodeType"] = undefined /*out*/;
            resourceInputs["notificationTopicARN"] = undefined /*out*/;
            resourceInputs["parameterGroupName"] = undefined /*out*/;
            resourceInputs["preferredMaintenanceWindow"] = undefined /*out*/;
            resourceInputs["replicationFactor"] = undefined /*out*/;
            resourceInputs["sSESpecification"] = undefined /*out*/;
            resourceInputs["securityGroupIds"] = undefined /*out*/;
            resourceInputs["subnetGroupName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    clusterEndpointEncryptionType?: pulumi.Input<string>;
    clusterName?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    iAMRoleARN: pulumi.Input<string>;
    nodeType: pulumi.Input<string>;
    notificationTopicARN?: pulumi.Input<string>;
    parameterGroupName?: pulumi.Input<string>;
    preferredMaintenanceWindow?: pulumi.Input<string>;
    replicationFactor: pulumi.Input<number>;
    sSESpecification?: pulumi.Input<inputs.dax.ClusterSSESpecificationArgs>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    subnetGroupName?: pulumi.Input<string>;
    tags?: any;
}
