// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    /// <summary>
    /// The provider type for the yandex package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [YandexResourceType("pulumi:providers:yandex")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// ID of Yandex.Cloud tenant.
        /// </summary>
        [Output("cloudId")]
        public Output<string?> CloudId { get; private set; } = null!;

        /// <summary>
        /// The API endpoint for Yandex.Cloud SDK client.
        /// </summary>
        [Output("endpoint")]
        public Output<string?> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The default folder ID where resources will be placed.
        /// </summary>
        [Output("folderId")]
        public Output<string?> FolderId { get; private set; } = null!;

        [Output("organizationId")]
        public Output<string?> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Profile to use in the shared credentials file. Default value is `default`.
        /// </summary>
        [Output("profile")]
        public Output<string?> Profile { get; private set; } = null!;

        /// <summary>
        /// The region where operations will take place. Examples are ru-central1
        /// </summary>
        [Output("regionId")]
        public Output<string?> RegionId { get; private set; } = null!;

        /// <summary>
        /// Either the path to or the contents of a Service Account key file in JSON format.
        /// </summary>
        [Output("serviceAccountKeyFile")]
        public Output<string?> ServiceAccountKeyFile { get; private set; } = null!;

        /// <summary>
        /// Path to shared credentials file.
        /// </summary>
        [Output("sharedCredentialsFile")]
        public Output<string?> SharedCredentialsFile { get; private set; } = null!;

        /// <summary>
        /// Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly
        /// specified.
        /// </summary>
        [Output("storageAccessKey")]
        public Output<string?> StorageAccessKey { get; private set; } = null!;

        /// <summary>
        /// Yandex.Cloud storage service endpoint. Default is storage.yandexcloud.net
        /// </summary>
        [Output("storageEndpoint")]
        public Output<string?> StorageEndpoint { get; private set; } = null!;

        /// <summary>
        /// Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly
        /// specified.
        /// </summary>
        [Output("storageSecretKey")]
        public Output<string?> StorageSecretKey { get; private set; } = null!;

        /// <summary>
        /// The access token for API operations.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;

        /// <summary>
        /// Yandex.Cloud Message Queue service access key. Used when a message queue resource doesn't have an access key explicitly
        /// specified.
        /// </summary>
        [Output("ymqAccessKey")]
        public Output<string?> YmqAccessKey { get; private set; } = null!;

        /// <summary>
        /// Yandex.Cloud Message Queue service endpoint. Default is message-queue.api.cloud.yandex.net
        /// </summary>
        [Output("ymqEndpoint")]
        public Output<string?> YmqEndpoint { get; private set; } = null!;

        /// <summary>
        /// Yandex.Cloud Message Queue service secret key. Used when a message queue resource doesn't have a secret key explicitly
        /// specified.
        /// </summary>
        [Output("ymqSecretKey")]
        public Output<string?> YmqSecretKey { get; private set; } = null!;

        /// <summary>
        /// The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
        /// </summary>
        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("yandex", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/regrau/pulumi-yandex/releases",
                AdditionalSecretOutputs =
                {
                    "storageSecretKey",
                    "token",
                    "ymqSecretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of Yandex.Cloud tenant.
        /// </summary>
        [Input("cloudId")]
        public Input<string>? CloudId { get; set; }

        /// <summary>
        /// The API endpoint for Yandex.Cloud SDK client.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The default folder ID where resources will be placed.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
        /// </summary>
        [Input("insecure", json: true)]
        public Input<bool>? Insecure { get; set; }

        /// <summary>
        /// The maximum number of times an API request is being executed. If the API request still fails, an error is thrown.
        /// </summary>
        [Input("maxRetries", json: true)]
        public Input<int>? MaxRetries { get; set; }

        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Disable use of TLS. Default value is `false`.
        /// </summary>
        [Input("plaintext", json: true)]
        public Input<bool>? Plaintext { get; set; }

        /// <summary>
        /// Profile to use in the shared credentials file. Default value is `default`.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// The region where operations will take place. Examples are ru-central1
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// Either the path to or the contents of a Service Account key file in JSON format.
        /// </summary>
        [Input("serviceAccountKeyFile")]
        public Input<string>? ServiceAccountKeyFile { get; set; }

        /// <summary>
        /// Path to shared credentials file.
        /// </summary>
        [Input("sharedCredentialsFile")]
        public Input<string>? SharedCredentialsFile { get; set; }

        /// <summary>
        /// Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly
        /// specified.
        /// </summary>
        [Input("storageAccessKey")]
        public Input<string>? StorageAccessKey { get; set; }

        /// <summary>
        /// Yandex.Cloud storage service endpoint. Default is storage.yandexcloud.net
        /// </summary>
        [Input("storageEndpoint")]
        public Input<string>? StorageEndpoint { get; set; }

        [Input("storageSecretKey")]
        private Input<string>? _storageSecretKey;

        /// <summary>
        /// Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly
        /// specified.
        /// </summary>
        public Input<string>? StorageSecretKey
        {
            get => _storageSecretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _storageSecretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The access token for API operations.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Yandex.Cloud Message Queue service access key. Used when a message queue resource doesn't have an access key explicitly
        /// specified.
        /// </summary>
        [Input("ymqAccessKey")]
        public Input<string>? YmqAccessKey { get; set; }

        /// <summary>
        /// Yandex.Cloud Message Queue service endpoint. Default is message-queue.api.cloud.yandex.net
        /// </summary>
        [Input("ymqEndpoint")]
        public Input<string>? YmqEndpoint { get; set; }

        [Input("ymqSecretKey")]
        private Input<string>? _ymqSecretKey;

        /// <summary>
        /// Yandex.Cloud Message Queue service secret key. Used when a message queue resource doesn't have a secret key explicitly
        /// specified.
        /// </summary>
        public Input<string>? YmqSecretKey
        {
            get => _ymqSecretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ymqSecretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
