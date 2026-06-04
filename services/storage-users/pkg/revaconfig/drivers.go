package revaconfig

import (
	"github.com/opencloud-eu/opencloud/services/storage-users/pkg/config"
)

// EOS is the config mapping for the EOS storage driver
func EOS(cfg *config.Config) map[string]any {
	return map[string]any{
		"namespace":              cfg.Drivers.EOS.Root,
		"shadow_namespace":       cfg.Drivers.EOS.ShadowNamespace,
		"uploads_namespace":      cfg.Drivers.EOS.UploadsNamespace,
		"share_folder":           cfg.Drivers.EOS.ShareFolder,
		"eos_binary":             cfg.Drivers.EOS.EosBinary,
		"xrdcopy_binary":         cfg.Drivers.EOS.XrdcopyBinary,
		"master_url":             cfg.Drivers.EOS.MasterURL,
		"slave_url":              cfg.Drivers.EOS.SlaveURL,
		"cache_directory":        cfg.Drivers.EOS.CacheDirectory,
		"sec_protocol":           cfg.Drivers.EOS.SecProtocol,
		"keytab":                 cfg.Drivers.EOS.Keytab,
		"single_username":        cfg.Drivers.EOS.SingleUsername,
		"enable_logging":         cfg.Drivers.EOS.EnableLogging,
		"show_hidden_sys_files":  cfg.Drivers.EOS.ShowHiddenSysFiles,
		"force_single_user_mode": cfg.Drivers.EOS.ForceSingleUserMode,
		"use_keytab":             cfg.Drivers.EOS.UseKeytab,
		"gatewaysvc":             cfg.Drivers.EOS.GatewaySVC,
	}
}

// EOSHome is the config mapping for the EOSHome storage driver
func EOSHome(cfg *config.Config) map[string]any {
	return map[string]any{
		"namespace":              cfg.Drivers.EOS.Root,
		"shadow_namespace":       cfg.Drivers.EOS.ShadowNamespace,
		"uploads_namespace":      cfg.Drivers.EOS.UploadsNamespace,
		"share_folder":           cfg.Drivers.EOS.ShareFolder,
		"eos_binary":             cfg.Drivers.EOS.EosBinary,
		"xrdcopy_binary":         cfg.Drivers.EOS.XrdcopyBinary,
		"master_url":             cfg.Drivers.EOS.MasterURL,
		"slave_url":              cfg.Drivers.EOS.SlaveURL,
		"cache_directory":        cfg.Drivers.EOS.CacheDirectory,
		"sec_protocol":           cfg.Drivers.EOS.SecProtocol,
		"keytab":                 cfg.Drivers.EOS.Keytab,
		"single_username":        cfg.Drivers.EOS.SingleUsername,
		"user_layout":            cfg.Drivers.EOS.UserLayout,
		"enable_logging":         cfg.Drivers.EOS.EnableLogging,
		"show_hidden_sys_files":  cfg.Drivers.EOS.ShowHiddenSysFiles,
		"force_single_user_mode": cfg.Drivers.EOS.ForceSingleUserMode,
		"use_keytab":             cfg.Drivers.EOS.UseKeytab,
		"gatewaysvc":             cfg.Drivers.EOS.GatewaySVC,
	}
}

// EOSGRPC is the config mapping for the EOSGRPC storage driver
func EOSGRPC(cfg *config.Config) map[string]any {
	return map[string]any{
		"namespace":              cfg.Drivers.EOS.Root,
		"shadow_namespace":       cfg.Drivers.EOS.ShadowNamespace,
		"share_folder":           cfg.Drivers.EOS.ShareFolder,
		"eos_binary":             cfg.Drivers.EOS.EosBinary,
		"xrdcopy_binary":         cfg.Drivers.EOS.XrdcopyBinary,
		"master_url":             cfg.Drivers.EOS.MasterURL,
		"master_grpc_uri":        cfg.Drivers.EOS.GRPCURI,
		"slave_url":              cfg.Drivers.EOS.SlaveURL,
		"cache_directory":        cfg.Drivers.EOS.CacheDirectory,
		"sec_protocol":           cfg.Drivers.EOS.SecProtocol,
		"keytab":                 cfg.Drivers.EOS.Keytab,
		"single_username":        cfg.Drivers.EOS.SingleUsername,
		"user_layout":            cfg.Drivers.EOS.UserLayout,
		"enable_logging":         cfg.Drivers.EOS.EnableLogging,
		"show_hidden_sys_files":  cfg.Drivers.EOS.ShowHiddenSysFiles,
		"force_single_user_mode": cfg.Drivers.EOS.ForceSingleUserMode,
		"use_keytab":             cfg.Drivers.EOS.UseKeytab,
		"enable_home":            false,
		"gatewaysvc":             cfg.Drivers.EOS.GatewaySVC,
	}
}

// Local is the config mapping for the Local storage driver
func Local(cfg *config.Config) map[string]any {
	return map[string]any{
		"root":         cfg.Drivers.Local.Root,
		"share_folder": cfg.Drivers.Local.ShareFolder,
	}
}

// Posix is the config mapping for the Posix storage driver
func Posix(cfg *config.Config, enableFSScan, enableFSWatch bool) map[string]any {
	return map[string]any{
		"root":                        cfg.Drivers.Posix.Root,
		"personalspacepath_template":  cfg.Drivers.Posix.PersonalSpacePathTemplate,
		"personalspacealias_template": cfg.Drivers.Posix.PersonalSpaceAliasTemplate,
		"generalspacepath_template":   cfg.Drivers.Posix.GeneralSpacePathTemplate,
		"generalspacealias_template":  cfg.Drivers.Posix.GeneralSpaceAliasTemplate,
		"permissionssvc":              cfg.Drivers.Posix.PermissionsEndpoint,
		"permissionssvc_tls_mode":     cfg.Commons.GRPCClientTLS.Mode,
		"treetime_accounting":         true,
		"treesize_accounting":         true,
		"asyncfileuploads":            cfg.Drivers.Posix.AsyncUploads,
		"scan_debounce_delay":         cfg.Drivers.Posix.ScanDebounceDelay,
		"max_quota":                   cfg.Drivers.Posix.MaxQuota,
		"disable_versioning":          cfg.Drivers.Posix.DisableVersioning,
		"multi_tenant_enabled":        cfg.Commons.MultiTenantEnabled,
		"propagator":                  cfg.Drivers.Posix.Propagator,
		"async_propagator_options": map[string]any{
			"propagation_delay": cfg.Drivers.Posix.AsyncPropagatorOptions.PropagationDelay,
		},
		"max_acquire_lock_cycles":    cfg.Drivers.Posix.MaxAcquireLockCycles,
		"lock_cycle_duration_factor": cfg.Drivers.Posix.LockCycleDurationFactor,
		"max_concurrency":            cfg.Drivers.Posix.MaxConcurrency,
		"idcache": map[string]any{
			"cache_store":               cfg.IDCache.Store,
			"cache_nodes":               cfg.IDCache.Nodes,
			"cache_database":            cfg.IDCache.Database,
			"cache_disable_persistence": cfg.IDCache.DisablePersistence,
			"cache_auth_username":       cfg.IDCache.AuthUsername,
			"cache_auth_password":       cfg.IDCache.AuthPassword,
		},
		"filemetadatacache": map[string]any{
			"cache_store":               cfg.FilemetadataCache.Store,
			"cache_nodes":               cfg.FilemetadataCache.Nodes,
			"cache_database":            cfg.FilemetadataCache.Database,
			"cache_ttl":                 cfg.FilemetadataCache.TTL,
			"cache_disable_persistence": cfg.FilemetadataCache.DisablePersistence,
			"cache_auth_username":       cfg.FilemetadataCache.AuthUsername,
			"cache_auth_password":       cfg.FilemetadataCache.AuthPassword,
		},
		"events": map[string]any{
			"numconsumers": cfg.Events.NumConsumers,
		},
		"tokens": map[string]any{
			"transfer_shared_secret": cfg.Commons.TransferSecret,
			"transfer_expires":       cfg.TransferExpires,
			"download_endpoint":      cfg.DataServerURL,
			"datagateway_endpoint":   cfg.DataGatewayURL,
		},
		"use_space_groups":           cfg.Drivers.Posix.UseSpaceGroups,
		"enable_fs_revisions":        cfg.Drivers.Posix.EnableFSRevisions,
		"scan_fs":                    enableFSScan,
		"watch_fs":                   enableFSWatch,
		"watch_type":                 cfg.Drivers.Posix.WatchType,
		"watch_path":                 cfg.Drivers.Posix.WatchPath,
		"watch_notification_brokers": cfg.Drivers.Posix.WatchNotificationBrokers,
		"watch_root":                 cfg.Drivers.Posix.WatchRoot,
		"inotify_stats_frequency":    cfg.Drivers.Posix.InotifyStatsFrequency,
	}
}

// LocalHome is the config mapping for the LocalHome storage driver
func LocalHome(cfg *config.Config) map[string]any {
	return map[string]any{
		"root":         cfg.Drivers.Local.Root,
		"share_folder": cfg.Drivers.Local.ShareFolder,
		"user_layout":  cfg.Drivers.Local.UserLayout,
	}
}

// OwnCloudSQL is the config mapping for the OwnCloudSQL storage driver
func OwnCloudSQL(cfg *config.Config) map[string]any {
	return map[string]any{
		"datadirectory":   cfg.Drivers.OwnCloudSQL.Root,
		"upload_info_dir": cfg.Drivers.OwnCloudSQL.UploadInfoDir,
		"share_folder":    cfg.Drivers.OwnCloudSQL.ShareFolder,
		"user_layout":     cfg.Drivers.OwnCloudSQL.UserLayout,
		"enable_home":     false,
		"dbusername":      cfg.Drivers.OwnCloudSQL.DBUsername,
		"dbpassword":      cfg.Drivers.OwnCloudSQL.DBPassword,
		"dbhost":          cfg.Drivers.OwnCloudSQL.DBHost,
		"dbport":          cfg.Drivers.OwnCloudSQL.DBPort,
		"dbname":          cfg.Drivers.OwnCloudSQL.DBName,
		"userprovidersvc": cfg.Drivers.OwnCloudSQL.UsersProviderEndpoint,
		"tokens": map[string]any{
			"download_endpoint":      cfg.DataServerURL,
			"datagateway_endpoint":   cfg.DataGatewayURL,
			"transfer_shared_secret": cfg.Commons.TransferSecret,
			"transfer_expires":       cfg.TransferExpires,
		},
	}
}

// Decomposed is the config mapping for the Decomposed storage driver
func Decomposed(cfg *config.Config) map[string]any {
	return map[string]any{
		"metadata_backend": "messagepack",
		"propagator":       cfg.Drivers.Decomposed.Propagator,
		"async_propagator_options": map[string]any{
			"propagation_delay": cfg.Drivers.Decomposed.AsyncPropagatorOptions.PropagationDelay,
		},
		"root":                        cfg.Drivers.Decomposed.Root,
		"user_layout":                 cfg.Drivers.Decomposed.UserLayout,
		"share_folder":                cfg.Drivers.Decomposed.ShareFolder,
		"personalspacealias_template": cfg.Drivers.Decomposed.PersonalSpaceAliasTemplate,
		"personalspacepath_template":  cfg.Drivers.Decomposed.PersonalSpacePathTemplate,
		"generalspacealias_template":  cfg.Drivers.Decomposed.GeneralSpaceAliasTemplate,
		"generalspacepath_template":   cfg.Drivers.Decomposed.GeneralSpacePathTemplate,
		"treetime_accounting":         true,
		"treesize_accounting":         true,
		"permissionssvc":              cfg.Drivers.Decomposed.PermissionsEndpoint,
		"permissionssvc_tls_mode":     cfg.Commons.GRPCClientTLS.Mode,
		"max_acquire_lock_cycles":     cfg.Drivers.Decomposed.MaxAcquireLockCycles,
		"lock_cycle_duration_factor":  cfg.Drivers.Decomposed.LockCycleDurationFactor,
		"max_concurrency":             cfg.Drivers.Decomposed.MaxConcurrency,
		"asyncfileuploads":            cfg.Drivers.Decomposed.AsyncUploads,
		"max_quota":                   cfg.Drivers.Decomposed.MaxQuota,
		"disable_versioning":          cfg.Drivers.Decomposed.DisableVersioning,
		"multi_tenant_enabled":        cfg.Commons.MultiTenantEnabled,
		"filemetadatacache": map[string]any{
			"cache_store":               cfg.FilemetadataCache.Store,
			"cache_nodes":               cfg.FilemetadataCache.Nodes,
			"cache_database":            cfg.FilemetadataCache.Database,
			"cache_ttl":                 cfg.FilemetadataCache.TTL,
			"cache_disable_persistence": cfg.FilemetadataCache.DisablePersistence,
			"cache_auth_username":       cfg.FilemetadataCache.AuthUsername,
			"cache_auth_password":       cfg.FilemetadataCache.AuthPassword,
		},
		"idcache": map[string]any{
			"cache_store":               cfg.IDCache.Store,
			"cache_nodes":               cfg.IDCache.Nodes,
			"cache_database":            cfg.IDCache.Database,
			"cache_disable_persistence": cfg.IDCache.DisablePersistence,
			"cache_auth_username":       cfg.IDCache.AuthUsername,
			"cache_auth_password":       cfg.IDCache.AuthPassword,
		},
		"events": map[string]any{
			"numconsumers": cfg.Events.NumConsumers,
		},
		"tokens": map[string]any{
			"transfer_shared_secret": cfg.Commons.TransferSecret,
			"transfer_expires":       cfg.TransferExpires,
			"download_endpoint":      cfg.DataServerURL,
			"datagateway_endpoint":   cfg.DataGatewayURL,
		},
	}
}

// DecomposedsNoEvents is the config mapping for the Decomposed storage driver emitting no events
func DecomposedNoEvents(cfg *config.Config) map[string]any {
	return map[string]any{
		"metadata_backend": "messagepack",
		"propagator":       cfg.Drivers.Decomposed.Propagator,
		"async_propagator_options": map[string]any{
			"propagation_delay": cfg.Drivers.Decomposed.AsyncPropagatorOptions.PropagationDelay,
		},
		"root":                        cfg.Drivers.Decomposed.Root,
		"user_layout":                 cfg.Drivers.Decomposed.UserLayout,
		"share_folder":                cfg.Drivers.Decomposed.ShareFolder,
		"personalspacealias_template": cfg.Drivers.Decomposed.PersonalSpaceAliasTemplate,
		"personalspacepath_template":  cfg.Drivers.Decomposed.PersonalSpacePathTemplate,
		"generalspacealias_template":  cfg.Drivers.Decomposed.GeneralSpaceAliasTemplate,
		"generalspacepath_template":   cfg.Drivers.Decomposed.GeneralSpacePathTemplate,
		"treetime_accounting":         true,
		"treesize_accounting":         true,
		"permissionssvc":              cfg.Drivers.Decomposed.PermissionsEndpoint,
		"permissionssvc_tls_mode":     cfg.Commons.GRPCClientTLS.Mode,
		"max_acquire_lock_cycles":     cfg.Drivers.Decomposed.MaxAcquireLockCycles,
		"lock_cycle_duration_factor":  cfg.Drivers.Decomposed.LockCycleDurationFactor,
		"max_concurrency":             cfg.Drivers.Decomposed.MaxConcurrency,
		"max_quota":                   cfg.Drivers.Decomposed.MaxQuota,
		"disable_versioning":          cfg.Drivers.Decomposed.DisableVersioning,
		"multi_tenant_enabled":        cfg.Commons.MultiTenantEnabled,
		"filemetadatacache": map[string]any{
			"cache_store":               cfg.FilemetadataCache.Store,
			"cache_nodes":               cfg.FilemetadataCache.Nodes,
			"cache_database":            cfg.FilemetadataCache.Database,
			"cache_ttl":                 cfg.FilemetadataCache.TTL,
			"cache_disable_persistence": cfg.FilemetadataCache.DisablePersistence,
			"cache_auth_username":       cfg.FilemetadataCache.AuthUsername,
			"cache_auth_password":       cfg.FilemetadataCache.AuthPassword,
		},
		"idcache": map[string]any{
			"cache_store":               cfg.IDCache.Store,
			"cache_nodes":               cfg.IDCache.Nodes,
			"cache_database":            cfg.IDCache.Database,
			"cache_disable_persistence": cfg.IDCache.DisablePersistence,
			"cache_auth_username":       cfg.IDCache.AuthUsername,
			"cache_auth_password":       cfg.IDCache.AuthPassword,
		},
	}
}

// DecomposedS3 is the config mapping for the decomposeds3 storage driver
func DecomposedS3(cfg *config.Config) map[string]any {
	return map[string]any{
		"metadata_backend": "messagepack",
		"propagator":       cfg.Drivers.DecomposedS3.Propagator,
		"async_propagator_options": map[string]any{
			"propagation_delay": cfg.Drivers.DecomposedS3.AsyncPropagatorOptions.PropagationDelay,
		},
		"root":                        cfg.Drivers.DecomposedS3.Root,
		"user_layout":                 cfg.Drivers.DecomposedS3.UserLayout,
		"share_folder":                cfg.Drivers.DecomposedS3.ShareFolder,
		"personalspacealias_template": cfg.Drivers.DecomposedS3.PersonalSpaceAliasTemplate,
		"personalspacepath_template":  cfg.Drivers.DecomposedS3.PersonalSpacePathTemplate,
		"generalspacealias_template":  cfg.Drivers.DecomposedS3.GeneralSpaceAliasTemplate,
		"generalspacepath_template":   cfg.Drivers.DecomposedS3.GeneralSpacePathTemplate,
		"treetime_accounting":         true,
		"treesize_accounting":         true,
		"permissionssvc":              cfg.Drivers.DecomposedS3.PermissionsEndpoint,
		"permissionssvc_tls_mode":     cfg.Commons.GRPCClientTLS.Mode,
		"s3.region":                   cfg.Drivers.DecomposedS3.Region,
		"s3.access_key":               cfg.Drivers.DecomposedS3.AccessKey,
		"s3.secret_key":               cfg.Drivers.DecomposedS3.SecretKey,
		"s3.endpoint":                 cfg.Drivers.DecomposedS3.Endpoint,
		"s3.bucket":                   cfg.Drivers.DecomposedS3.Bucket,
		"s3.disable_content_sha254":   cfg.Drivers.DecomposedS3.DisableContentSha256,
		"s3.disable_multipart":        cfg.Drivers.DecomposedS3.DisableMultipart,
		"s3.send_content_md5":         cfg.Drivers.DecomposedS3.SendContentMd5,
		"s3.concurrent_stream_parts":  cfg.Drivers.DecomposedS3.ConcurrentStreamParts,
		"s3.num_threads":              cfg.Drivers.DecomposedS3.NumThreads,
		"s3.part_size":                cfg.Drivers.DecomposedS3.PartSize,
		"max_acquire_lock_cycles":     cfg.Drivers.DecomposedS3.MaxAcquireLockCycles,
		"lock_cycle_duration_factor":  cfg.Drivers.DecomposedS3.LockCycleDurationFactor,
		"max_concurrency":             cfg.Drivers.DecomposedS3.MaxConcurrency,
		"disable_versioning":          cfg.Drivers.DecomposedS3.DisableVersioning,
		"multi_tenant_enabled":        cfg.Commons.MultiTenantEnabled,
		"asyncfileuploads":            cfg.Drivers.DecomposedS3.AsyncUploads,
		"filemetadatacache": map[string]any{
			"cache_store":               cfg.FilemetadataCache.Store,
			"cache_nodes":               cfg.FilemetadataCache.Nodes,
			"cache_database":            cfg.FilemetadataCache.Database,
			"cache_ttl":                 cfg.FilemetadataCache.TTL,
			"cache_disable_persistence": cfg.FilemetadataCache.DisablePersistence,
			"cache_auth_username":       cfg.FilemetadataCache.AuthUsername,
			"cache_auth_password":       cfg.FilemetadataCache.AuthPassword,
		},
		"idcache": map[string]any{
			"cache_store":               cfg.IDCache.Store,
			"cache_nodes":               cfg.IDCache.Nodes,
			"cache_database":            cfg.IDCache.Database,
			"cache_disable_persistence": cfg.IDCache.DisablePersistence,
			"cache_auth_username":       cfg.IDCache.AuthUsername,
			"cache_auth_password":       cfg.IDCache.AuthPassword,
		},
		"events": map[string]any{
			"numconsumers": cfg.Events.NumConsumers,
		},
		"tokens": map[string]any{
			"transfer_shared_secret": cfg.Commons.TransferSecret,
			"transfer_expires":       cfg.TransferExpires,
			"download_endpoint":      cfg.DataServerURL,
			"datagateway_endpoint":   cfg.DataGatewayURL,
		},
	}
}

// DecomposedS3NoEvents is the config mapping for the decomposeds3 storage driver emitting no events
func DecomposedS3NoEvents(cfg *config.Config) map[string]any {
	return map[string]any{
		"metadata_backend": "messagepack",
		"propagator":       cfg.Drivers.DecomposedS3.Propagator,
		"async_propagator_options": map[string]any{
			"propagation_delay": cfg.Drivers.DecomposedS3.AsyncPropagatorOptions.PropagationDelay,
		},
		"root":                        cfg.Drivers.DecomposedS3.Root,
		"user_layout":                 cfg.Drivers.DecomposedS3.UserLayout,
		"share_folder":                cfg.Drivers.DecomposedS3.ShareFolder,
		"personalspacealias_template": cfg.Drivers.Decomposed.PersonalSpaceAliasTemplate,
		"personalspacepath_template":  cfg.Drivers.Decomposed.PersonalSpacePathTemplate,
		"generalspacealias_template":  cfg.Drivers.Decomposed.GeneralSpaceAliasTemplate,
		"generalspacepath_template":   cfg.Drivers.Decomposed.GeneralSpacePathTemplate,
		"treetime_accounting":         true,
		"treesize_accounting":         true,
		"permissionssvc":              cfg.Drivers.DecomposedS3.PermissionsEndpoint,
		"permissionssvc_tls_mode":     cfg.Commons.GRPCClientTLS.Mode,
		"s3.region":                   cfg.Drivers.DecomposedS3.Region,
		"s3.access_key":               cfg.Drivers.DecomposedS3.AccessKey,
		"s3.secret_key":               cfg.Drivers.DecomposedS3.SecretKey,
		"s3.endpoint":                 cfg.Drivers.DecomposedS3.Endpoint,
		"s3.bucket":                   cfg.Drivers.DecomposedS3.Bucket,
		"max_acquire_lock_cycles":     cfg.Drivers.DecomposedS3.MaxAcquireLockCycles,
		"max_concurrency":             cfg.Drivers.DecomposedS3.MaxConcurrency,
		"disable_versioning":          cfg.Drivers.DecomposedS3.DisableVersioning,
		"multi_tenant_enabled":        cfg.Commons.MultiTenantEnabled,
		"lock_cycle_duration_factor":  cfg.Drivers.DecomposedS3.LockCycleDurationFactor,
		"filemetadatacache": map[string]any{
			"cache_store":               cfg.FilemetadataCache.Store,
			"cache_nodes":               cfg.FilemetadataCache.Nodes,
			"cache_database":            cfg.FilemetadataCache.Database,
			"cache_ttl":                 cfg.FilemetadataCache.TTL,
			"cache_disable_persistence": cfg.FilemetadataCache.DisablePersistence,
			"cache_auth_username":       cfg.FilemetadataCache.AuthUsername,
			"cache_auth_password":       cfg.FilemetadataCache.AuthPassword,
		},
		"idcache": map[string]any{
			"cache_store":               cfg.IDCache.Store,
			"cache_nodes":               cfg.IDCache.Nodes,
			"cache_database":            cfg.IDCache.Database,
			"cache_disable_persistence": cfg.IDCache.DisablePersistence,
			"cache_auth_username":       cfg.IDCache.AuthUsername,
			"cache_auth_password":       cfg.IDCache.AuthPassword,
		},
	}
}

// KVFS is the config mapping for the KVFS storage driver
func KVFS(cfg *config.Config) map[string]interface{} {
	// The GC resolves space-owner liveness through the gateway IN-PROCESS, so it
	// must address the gateway by its registry SERVICE NAME (e.g.
	// "eu.opencloud.api.gateway") — the registry-backed client pool resolves the
	// id as a service name, so a bind address like "127.0.0.1:9142" yields
	// "service not found". cfg.Reva.Address is exactly what the other in-process
	// services (graph, webdav) pass to the gateway selector; RevaGatewayGRPCAddr
	// is the bind address used only by no-registry CLI contexts (e.g. trash-bin
	// purge). Falling back to the bind address keeps split/no-registry setups working.
	gatewayAddr := cfg.RevaGatewayGRPCAddr
	if cfg.Reva != nil && cfg.Reva.Address != "" {
		gatewayAddr = cfg.Reva.Address
	}
	return map[string]interface{}{
		"nats_nodes":         cfg.Drivers.KVFS.NATSNodes,
		"nats_username":      cfg.Drivers.KVFS.NATSUsername,
		"nats_password":      cfg.Drivers.KVFS.NATSPassword,
		"nats_replicas":      cfg.Drivers.KVFS.NATSReplicas,
		"bucket_prefix":      cfg.Drivers.KVFS.BucketPrefix,
		"s3.endpoint":        cfg.Drivers.KVFS.S3Endpoint,
		"s3.region":          cfg.Drivers.KVFS.S3Region,
		"s3.bucket":          cfg.Drivers.KVFS.S3Bucket,
		"s3.access_key":      cfg.Drivers.KVFS.S3AccessKey,
		"s3.secret_key":      cfg.Drivers.KVFS.S3SecretKey,
		"disable_versioning": cfg.Drivers.KVFS.DisableVersioning,
		"max_versions":       cfg.Drivers.KVFS.MaxVersions,
		"gc_enabled":         cfg.Drivers.KVFS.GCEnabled,
		"gc_interval":        cfg.Drivers.KVFS.GCInterval,
		"gc_dry_run":         cfg.Drivers.KVFS.GCDryRun,
		"gc_min_age":         cfg.Drivers.KVFS.GCMinAge,
		"gc_run_on_start":    cfg.Drivers.KVFS.GCRunOnStart,
		// GC identity-orphan reaping: the driver resolves space-owner liveness
		// through the gateway as the service account. storage-system leaves the
		// service-account fields empty (no personal spaces) so its resolver stays
		// disabled. gateway_addr is the registry service name (computed above).
		"gateway_addr":                gatewayAddr,
		"service_account_id":          cfg.ServiceAccount.ServiceAccountID,
		"service_account_secret":      cfg.ServiceAccount.ServiceAccountSecret,
		"max_cas_retries":             cfg.Drivers.KVFS.MaxCASRetries,
		"children_max_value_size":     cfg.Drivers.KVFS.ChildrenMaxValueSize,
		"small_file_threshold":        cfg.Drivers.KVFS.SmallFileThreshold,
		"upload_backend":              cfg.Drivers.KVFS.UploadBackend,
		"upload_tmp_dir":              cfg.Drivers.KVFS.UploadTmpDir,
		"upload_nats_stream":          cfg.Drivers.KVFS.UploadNATSStream,
		"upload_nats_subject_prefix":  cfg.Drivers.KVFS.UploadNATSSubjectPrefix,
		"upload_nats_storage":         cfg.Drivers.KVFS.UploadNATSStorage,
		"upload_nats_replicas":        cfg.Drivers.KVFS.UploadNATSReplicas,
		"upload_nats_max_age":         cfg.Drivers.KVFS.UploadNATSMaxAge,
		"upload_nats_max_bytes":       cfg.Drivers.KVFS.UploadNATSMaxBytes,
		"upload_nats_max_chunk_bytes": cfg.Drivers.KVFS.UploadNATSMaxChunkBytes,
		"mount_id":                    cfg.MountID,
	}
}
