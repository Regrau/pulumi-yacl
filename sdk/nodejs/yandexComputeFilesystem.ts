// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * File storage is a virtual file system that can be attached to multiple Compute Cloud VMs in the same availability zone.
 *
 * Users can share files in storage and use them from different VMs.
 *
 * Storage is attached to a VM through the [Filesystem in Userspace](https://en.wikipedia.org/wiki/Filesystem_in_Userspace) (FUSE) interface as a [virtiofs](https://www.kernel.org/doc/html/latest/filesystems/virtiofs.html) device that is not linked to the host file system directly.
 *
 * For more information about filesystems in Yandex.Cloud, see:
 *
 * * [Documentation](https://cloud.yandex.com/docs/compute/concepts/filesystem)
 * * How-to Guides
 *     * [Attach filesystem to a VM](https://cloud.yandex.com/en-ru/docs/compute/operations/filesystem/attach-to-vm)
 *     * [Detach filesystem from VM](https://cloud.yandex.com/en-ru/docs/compute/operations/filesystem/detach-from-vm)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const defaultyandexComputeFilesystem = new yandex.yandexComputeFilesystem("default", {
 *     labels: {
 *         environment: "test",
 *     },
 *     size: 10,
 *     type: "network-ssd",
 *     zone: "ru-central1-a",
 * });
 * ```
 *
 * ## Import
 *
 * A filesystem can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import yandex:index/yandexComputeFilesystem:yandexComputeFilesystem default filesystem_id
 * ```
 */
export class YandexComputeFilesystem extends pulumi.CustomResource {
    /**
     * Get an existing YandexComputeFilesystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: YandexComputeFilesystemState, opts?: pulumi.CustomResourceOptions): YandexComputeFilesystem {
        return new YandexComputeFilesystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/yandexComputeFilesystem:yandexComputeFilesystem';

    /**
     * Returns true if the given object is an instance of YandexComputeFilesystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is YandexComputeFilesystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === YandexComputeFilesystem.__pulumiType;
    }

    /**
     * Block size of the filesystem, specified in bytes.
     */
    public readonly blockSize!: pulumi.Output<number | undefined>;
    /**
     * Creation timestamp of the filesystem.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Description of the filesystem. Provide this property when you create a resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the folder that the filesystem belongs to. If it is not provided, the default 
     * provider folder is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Labels to assign to this filesystem. A list of key/value pairs. For details about the concept, 
     * see [documentation](https://cloud.yandex.com/docs/overview/concepts/services#labels).
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the filesystem. Provide this property when you create a resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Size of the filesystem, specified in GB.
     */
    public readonly size!: pulumi.Output<number | undefined>;
    /**
     * The status of the filesystem.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Type of filesystem to create. Type `network-hdd` is set by default.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Availability zone where the filesystem will reside.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a YandexComputeFilesystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: YandexComputeFilesystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: YandexComputeFilesystemArgs | YandexComputeFilesystemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as YandexComputeFilesystemState | undefined;
            resourceInputs["blockSize"] = state ? state.blockSize : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as YandexComputeFilesystemArgs | undefined;
            resourceInputs["blockSize"] = args ? args.blockSize : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(YandexComputeFilesystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering yandexComputeFilesystem resources.
 */
export interface YandexComputeFilesystemState {
    /**
     * Block size of the filesystem, specified in bytes.
     */
    blockSize?: pulumi.Input<number>;
    /**
     * Creation timestamp of the filesystem.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Description of the filesystem. Provide this property when you create a resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the folder that the filesystem belongs to. If it is not provided, the default 
     * provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this filesystem. A list of key/value pairs. For details about the concept, 
     * see [documentation](https://cloud.yandex.com/docs/overview/concepts/services#labels).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the filesystem. Provide this property when you create a resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Size of the filesystem, specified in GB.
     */
    size?: pulumi.Input<number>;
    /**
     * The status of the filesystem.
     */
    status?: pulumi.Input<string>;
    /**
     * Type of filesystem to create. Type `network-hdd` is set by default.
     */
    type?: pulumi.Input<string>;
    /**
     * Availability zone where the filesystem will reside.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a YandexComputeFilesystem resource.
 */
export interface YandexComputeFilesystemArgs {
    /**
     * Block size of the filesystem, specified in bytes.
     */
    blockSize?: pulumi.Input<number>;
    /**
     * Description of the filesystem. Provide this property when you create a resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the folder that the filesystem belongs to. If it is not provided, the default 
     * provider folder is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Labels to assign to this filesystem. A list of key/value pairs. For details about the concept, 
     * see [documentation](https://cloud.yandex.com/docs/overview/concepts/services#labels).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the filesystem. Provide this property when you create a resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Size of the filesystem, specified in GB.
     */
    size?: pulumi.Input<number>;
    /**
     * Type of filesystem to create. Type `network-hdd` is set by default.
     */
    type?: pulumi.Input<string>;
    /**
     * Availability zone where the filesystem will reside.
     */
    zone?: pulumi.Input<string>;
}
