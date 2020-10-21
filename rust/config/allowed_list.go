package config

var (
	// When adding a new path below, add a rustfmt.toml file at the root of
	// the repository and enable the rustfmt repo hook. See aosp/1458238
	// for an example.
	// TODO(b/160223496): enable rustfmt globally.
	RustAllowedPaths = []string{
		"external/minijail",
		"external/rust",
		"external/crosvm",
		"external/adhd",
		"frameworks/native/libs/binder/rust",
		"packages/modules/Virtualization",
		"prebuilts/rust",
		"system/extras/profcollectd",
		"system/hardware/interfaces/keystore2",
		"system/security",
		"system/tools/aidl",
	}

	RustModuleTypes = []string{
		"rust_binary",
		"rust_binary_host",
		"rust_library",
		"rust_library_dylib",
		"rust_library_rlib",
		"rust_ffi",
		"rust_ffi_shared",
		"rust_ffi_static",
		"rust_library_host",
		"rust_library_host_dylib",
		"rust_library_host_rlib",
		"rust_ffi_host",
		"rust_ffi_host_shared",
		"rust_ffi_host_static",
		"rust_proc_macro",
		"rust_test",
		"rust_test_host",
	}
)