// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationObjectStorage.
 */
export class LocationObjectStorage extends pulumi.CustomResource {
    /**
     * Get an existing LocationObjectStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LocationObjectStorage {
        return new LocationObjectStorage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:datasync:LocationObjectStorage';

    /**
     * Returns true if the given object is an instance of LocationObjectStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocationObjectStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocationObjectStorage.__pulumiType;
    }

    /**
     * Optional. The access key is used if credentials are required to access the self-managed object storage server.
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
     */
    public readonly agentArns!: pulumi.Output<string[]>;
    /**
     * The name of the bucket on the self-managed object storage server.
     */
    public readonly bucketName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the location that is created.
     */
    public /*out*/ readonly locationArn!: pulumi.Output<string>;
    /**
     * The URL of the object storage location that was described.
     */
    public /*out*/ readonly locationUri!: pulumi.Output<string>;
    /**
     * Optional. The secret key is used if credentials are required to access the self-managed object storage server.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * X.509 PEM content containing a certificate authority or chain to trust.
     */
    public readonly serverCertificate!: pulumi.Output<string | undefined>;
    /**
     * The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.
     */
    public readonly serverHostname!: pulumi.Output<string | undefined>;
    /**
     * The port that your self-managed server accepts inbound network traffic on.
     */
    public readonly serverPort!: pulumi.Output<number | undefined>;
    /**
     * The protocol that the object storage server uses to communicate.
     */
    public readonly serverProtocol!: pulumi.Output<enums.datasync.LocationObjectStorageServerProtocol | undefined>;
    /**
     * The subdirectory in the self-managed object storage server that is used to read data from.
     */
    public readonly subdirectory!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.datasync.LocationObjectStorageTag[] | undefined>;

    /**
     * Create a LocationObjectStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocationObjectStorageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentArns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentArns'");
            }
            resourceInputs["accessKey"] = args ? args.accessKey : undefined;
            resourceInputs["agentArns"] = args ? args.agentArns : undefined;
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["secretKey"] = args ? args.secretKey : undefined;
            resourceInputs["serverCertificate"] = args ? args.serverCertificate : undefined;
            resourceInputs["serverHostname"] = args ? args.serverHostname : undefined;
            resourceInputs["serverPort"] = args ? args.serverPort : undefined;
            resourceInputs["serverProtocol"] = args ? args.serverProtocol : undefined;
            resourceInputs["subdirectory"] = args ? args.subdirectory : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["locationArn"] = undefined /*out*/;
            resourceInputs["locationUri"] = undefined /*out*/;
        } else {
            resourceInputs["accessKey"] = undefined /*out*/;
            resourceInputs["agentArns"] = undefined /*out*/;
            resourceInputs["bucketName"] = undefined /*out*/;
            resourceInputs["locationArn"] = undefined /*out*/;
            resourceInputs["locationUri"] = undefined /*out*/;
            resourceInputs["secretKey"] = undefined /*out*/;
            resourceInputs["serverCertificate"] = undefined /*out*/;
            resourceInputs["serverHostname"] = undefined /*out*/;
            resourceInputs["serverPort"] = undefined /*out*/;
            resourceInputs["serverProtocol"] = undefined /*out*/;
            resourceInputs["subdirectory"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LocationObjectStorage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LocationObjectStorage resource.
 */
export interface LocationObjectStorageArgs {
    /**
     * Optional. The access key is used if credentials are required to access the self-managed object storage server.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
     */
    agentArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the bucket on the self-managed object storage server.
     */
    bucketName?: pulumi.Input<string>;
    /**
     * Optional. The secret key is used if credentials are required to access the self-managed object storage server.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * X.509 PEM content containing a certificate authority or chain to trust.
     */
    serverCertificate?: pulumi.Input<string>;
    /**
     * The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.
     */
    serverHostname?: pulumi.Input<string>;
    /**
     * The port that your self-managed server accepts inbound network traffic on.
     */
    serverPort?: pulumi.Input<number>;
    /**
     * The protocol that the object storage server uses to communicate.
     */
    serverProtocol?: pulumi.Input<enums.datasync.LocationObjectStorageServerProtocol>;
    /**
     * The subdirectory in the self-managed object storage server that is used to read data from.
     */
    subdirectory?: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.datasync.LocationObjectStorageTagArgs>[]>;
}
