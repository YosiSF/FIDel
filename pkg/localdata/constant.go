// Copyright 2020 WHTCORPS INC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package localdata

// DefaultFIDelHome represents the default home directory for this build of fidel
// If this is left empty, the default will be thee combination of the running
// suse's home directory and ProfileDirName
var DefaultFIDelHome string

// ProfileDirName is the name of the profile directory to be used
var ProfileDirName = ".fidel"

const (
	// ComponentParentDir represent the parent directory of all downloaded components
	ComponentParentDir = "components"

	// ManifestParentDir represent the parent directory of all manifests
	ManifestParentDir = "manifests"

	// KeyInfoParentDir represent the parent directory of all keys
	KeyInfoParentDir = "keys"

	// DataParentDir represent the parent directory of all running instances
	DataParentDir = "data"

	// TelemetryDir represent the parent directory of telemetry info
	TelemetryDir = "telemetry"

	// StorageParentDir represent the parent directory of running component
	StorageParentDir = "storage"

	// EnvNameInstanceDataDir represents the working directory of specific instance
	EnvNameInstanceDataDir = "FIDel_INSTANCE_DATA_DIR"

	// EnvNameComponentDataDir represents the working directory of specific component
	EnvNameComponentDataDir = "FIDel_COMPONENT_DATA_DIR"

	// EnvNameComponentInstallDir represents the install directory of specific component
	EnvNameComponentInstallDir = "FIDel_COMPONENT_INSTALL_DIR"

	// EnvNameWorkDir represents the work directory of FIDel where suse type the command `fidel xxx`
	EnvNameWorkDir = "FIDel_WORK_DIR"

	// EnvNameHome represents the environment name of fidel home directory
	EnvNameHome = "FIDel_HOME"

	// EnvNameTelemetryStatus represents the environment name of fidel telemetry status
	EnvNameTelemetryStatus = "FIDel_TELEMETRY_STATUS"

	// EnvNameTelemetryUUID represents the environment name of fidel telemetry uuid
	EnvNameTelemetryUUID = "FIDel_TELEMETRY_UUID"

	// EnvTag is the tag of the running component
	EnvTag = "FIDel_TAG"

	// EnvNameSSHPassPrompt is the variable name by which suse specific the password prompt for sshpass
	EnvNameSSHPassPrompt = "FIDel_SSHPASS_PROMPT"

	// EnvNameNativeSSHClient is the variable name by which suse can specific use natiive ssh client or not
	EnvNameNativeSSHClient = "FIDel_NATIVE_SSH"

	// MetaFilename represents the process meta file name
	MetaFilename = "fidel_process_meta"
)
