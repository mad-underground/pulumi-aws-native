// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FileSystemArgs } from "./fileSystem";
export type FileSystem = import("./fileSystem").FileSystem;
export const FileSystem: typeof import("./fileSystem").FileSystem = null as any;

export { GetFileSystemArgs, GetFileSystemResult, GetFileSystemOutputArgs } from "./getFileSystem";
export const getFileSystem: typeof import("./getFileSystem").getFileSystem = null as any;
export const getFileSystemOutput: typeof import("./getFileSystem").getFileSystemOutput = null as any;

export { GetSnapshotArgs, GetSnapshotResult, GetSnapshotOutputArgs } from "./getSnapshot";
export const getSnapshot: typeof import("./getSnapshot").getSnapshot = null as any;
export const getSnapshotOutput: typeof import("./getSnapshot").getSnapshotOutput = null as any;

export { GetStorageVirtualMachineArgs, GetStorageVirtualMachineResult, GetStorageVirtualMachineOutputArgs } from "./getStorageVirtualMachine";
export const getStorageVirtualMachine: typeof import("./getStorageVirtualMachine").getStorageVirtualMachine = null as any;
export const getStorageVirtualMachineOutput: typeof import("./getStorageVirtualMachine").getStorageVirtualMachineOutput = null as any;

export { GetVolumeArgs, GetVolumeResult, GetVolumeOutputArgs } from "./getVolume";
export const getVolume: typeof import("./getVolume").getVolume = null as any;
export const getVolumeOutput: typeof import("./getVolume").getVolumeOutput = null as any;

export { SnapshotArgs } from "./snapshot";
export type Snapshot = import("./snapshot").Snapshot;
export const Snapshot: typeof import("./snapshot").Snapshot = null as any;

export { StorageVirtualMachineArgs } from "./storageVirtualMachine";
export type StorageVirtualMachine = import("./storageVirtualMachine").StorageVirtualMachine;
export const StorageVirtualMachine: typeof import("./storageVirtualMachine").StorageVirtualMachine = null as any;

export { VolumeArgs } from "./volume";
export type Volume = import("./volume").Volume;
export const Volume: typeof import("./volume").Volume = null as any;

utilities.lazyLoad(exports, ["FileSystem"], () => require("./fileSystem"));
utilities.lazyLoad(exports, ["getFileSystem","getFileSystemOutput"], () => require("./getFileSystem"));
utilities.lazyLoad(exports, ["getSnapshot","getSnapshotOutput"], () => require("./getSnapshot"));
utilities.lazyLoad(exports, ["getStorageVirtualMachine","getStorageVirtualMachineOutput"], () => require("./getStorageVirtualMachine"));
utilities.lazyLoad(exports, ["getVolume","getVolumeOutput"], () => require("./getVolume"));
utilities.lazyLoad(exports, ["Snapshot"], () => require("./snapshot"));
utilities.lazyLoad(exports, ["StorageVirtualMachine"], () => require("./storageVirtualMachine"));
utilities.lazyLoad(exports, ["Volume"], () => require("./volume"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:fsx:FileSystem":
                return new FileSystem(name, <any>undefined, { urn })
            case "aws-native:fsx:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "aws-native:fsx:StorageVirtualMachine":
                return new StorageVirtualMachine(name, <any>undefined, { urn })
            case "aws-native:fsx:Volume":
                return new Volume(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "fsx", _module)
