// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Yandex
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("yandex");

        private static readonly __Value<string?> _cloudId = new __Value<string?>(() => __config.Get("cloudId"));
        /// <summary>
        /// ID of Yandex.Cloud tenant.
        /// </summary>
        public static string? CloudId
        {
            get => _cloudId.Get();
            set => _cloudId.Set(value);
        }

        private static readonly __Value<string?> _endpoint = new __Value<string?>(() => __config.Get("endpoint"));
        /// <summary>
        /// The API endpoint for Yandex.Cloud SDK client.
        /// </summary>
        public static string? Endpoint
        {
            get => _endpoint.Get();
            set => _endpoint.Set(value);
        }

        private static readonly __Value<string?> _folderId = new __Value<string?>(() => __config.Get("folderId"));
        /// <summary>
        /// The default folder ID where resources will be placed.
        /// </summary>
        public static string? FolderId
        {
            get => _folderId.Get();
            set => _folderId.Set(value);
        }

        private static readonly __Value<bool?> _insecure = new __Value<bool?>(() => __config.GetBoolean("insecure"));
        /// <summary>
        /// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
        /// </summary>
        public static bool? Insecure
        {
            get => _insecure.Get();
            set => _insecure.Set(value);
        }

        private static readonly __Value<int?> _maxRetries = new __Value<int?>(() => __config.GetInt32("maxRetries"));
        /// <summary>
        /// The maximum number of times an API request is being executed. If the API request still fails, an error is thrown.
        /// </summary>
        public static int? MaxRetries
        {
            get => _maxRetries.Get();
            set => _maxRetries.Set(value);
        }

        private static readonly __Value<string?> _organizationId = new __Value<string?>(() => __config.Get("organizationId"));
        public static string? OrganizationId
        {
            get => _organizationId.Get();
            set => _organizationId.Set(value);
        }

        private static readonly __Value<bool?> _plaintext = new __Value<bool?>(() => __config.GetBoolean("plaintext"));
        /// <summary>
        /// Disable use of TLS. Default value is `false`.
        /// </summary>
        public static bool? Plaintext
        {
            get => _plaintext.Get();
            set => _plaintext.Set(value);
        }

        private static readonly __Value<string?> _regionId = new __Value<string?>(() => __config.Get("regionId"));
        public static string? RegionId
        {
            get => _regionId.Get();
            set => _regionId.Set(value);
        }

        private static readonly __Value<string?> _serviceAccountKeyFile = new __Value<string?>(() => __config.Get("serviceAccountKeyFile"));
        /// <summary>
        /// Either the path to or the contents of a Service Account key file in JSON format.
        /// </summary>
        public static string? ServiceAccountKeyFile
        {
            get => _serviceAccountKeyFile.Get();
            set => _serviceAccountKeyFile.Set(value);
        }

        private static readonly __Value<string?> _storageAccessKey = new __Value<string?>(() => __config.Get("storageAccessKey"));
        /// <summary>
        /// Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly
        /// specified.
        /// </summary>
        public static string? StorageAccessKey
        {
            get => _storageAccessKey.Get();
            set => _storageAccessKey.Set(value);
        }

        private static readonly __Value<string?> _storageEndpoint = new __Value<string?>(() => __config.Get("storageEndpoint"));
        /// <summary>
        /// Yandex.Cloud storage service endpoint. Default is storage.yandexcloud.net
        /// </summary>
        public static string? StorageEndpoint
        {
            get => _storageEndpoint.Get();
            set => _storageEndpoint.Set(value);
        }

        private static readonly __Value<string?> _storageSecretKey = new __Value<string?>(() => __config.Get("storageSecretKey"));
        /// <summary>
        /// Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly
        /// specified.
        /// </summary>
        public static string? StorageSecretKey
        {
            get => _storageSecretKey.Get();
            set => _storageSecretKey.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        /// <summary>
        /// The access token for API operations.
        /// </summary>
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

        private static readonly __Value<string?> _ymqAccessKey = new __Value<string?>(() => __config.Get("ymqAccessKey"));
        /// <summary>
        /// Yandex.Cloud Message Queue service access key. Used when a message queue resource doesn't have an access key explicitly
        /// specified.
        /// </summary>
        public static string? YmqAccessKey
        {
            get => _ymqAccessKey.Get();
            set => _ymqAccessKey.Set(value);
        }

        private static readonly __Value<string?> _ymqEndpoint = new __Value<string?>(() => __config.Get("ymqEndpoint"));
        /// <summary>
        /// Yandex.Cloud Message Queue service endpoint. Default is message-queue.api.cloud.yandex.net
        /// </summary>
        public static string? YmqEndpoint
        {
            get => _ymqEndpoint.Get();
            set => _ymqEndpoint.Set(value);
        }

        private static readonly __Value<string?> _ymqSecretKey = new __Value<string?>(() => __config.Get("ymqSecretKey"));
        /// <summary>
        /// Yandex.Cloud Message Queue service secret key. Used when a message queue resource doesn't have a secret key explicitly
        /// specified.
        /// </summary>
        public static string? YmqSecretKey
        {
            get => _ymqSecretKey.Get();
            set => _ymqSecretKey.Set(value);
        }

        private static readonly __Value<string?> _zone = new __Value<string?>(() => __config.Get("zone"));
        /// <summary>
        /// The zone where operations will take place. Examples are ru-central1-a, ru-central2-c, etc.
        /// </summary>
        public static string? Zone
        {
            get => _zone.Get();
            set => _zone.Set(value);
        }

    }
}