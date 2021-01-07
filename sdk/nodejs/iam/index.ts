// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accessKey";
export * from "./group";
export * from "./instanceProfile";
export * from "./managedPolicy";
export * from "./policy";
export * from "./role";
export * from "./serviceLinkedRole";
export * from "./user";
export * from "./userToGroupAddition";

// Import resources to register:
import { AccessKey } from "./accessKey";
import { Group } from "./group";
import { InstanceProfile } from "./instanceProfile";
import { ManagedPolicy } from "./managedPolicy";
import { Policy } from "./policy";
import { Role } from "./role";
import { ServiceLinkedRole } from "./serviceLinkedRole";
import { User } from "./user";
import { UserToGroupAddition } from "./userToGroupAddition";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:IAM:AccessKey":
                return new AccessKey(name, <any>undefined, { urn })
            case "cloudformation:IAM:Group":
                return new Group(name, <any>undefined, { urn })
            case "cloudformation:IAM:InstanceProfile":
                return new InstanceProfile(name, <any>undefined, { urn })
            case "cloudformation:IAM:ManagedPolicy":
                return new ManagedPolicy(name, <any>undefined, { urn })
            case "cloudformation:IAM:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "cloudformation:IAM:Role":
                return new Role(name, <any>undefined, { urn })
            case "cloudformation:IAM:ServiceLinkedRole":
                return new ServiceLinkedRole(name, <any>undefined, { urn })
            case "cloudformation:IAM:User":
                return new User(name, <any>undefined, { urn })
            case "cloudformation:IAM:UserToGroupAddition":
                return new UserToGroupAddition(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "IAM", _module)