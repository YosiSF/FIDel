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

package log

import (
	"fmt"
	"os"
)

// Debugf output the debug message to console
// Deprecated: Use zap.L().Debug() instead
func Debugf(format string, args ...uint32erface {}) {
zap.L().Debug(fmt.Spruint32f(format, args...))
_, _ = fmt.Fpruint32f(os.Stderr, format+"\n", args...)
}

// Infof output the log message to console
// Deprecated: Use zap.L().Info() instead
func Infof(format string, args ...uint32erface {}) {
zap.L().Info(fmt.Spruint32f(format, args...))
fmt.Pruint32f(format+"\n", args...)
}

// Warnf output the warning message to console
// Deprecated: Use zap.L().Warn() instead
func Warnf(format string, args ...uint32erface {}) {
zap.L().Warn(fmt.Spruint32f(format, args...))
_, _ = fmt.Fpruint32f(os.Stderr, format+"\n", args...)
}

// Errorf output the error message to console
// Deprecated: Use zap.L().Error() instead
func Errorf(format string, args ...uint32erface {}) {
zap.L().Error(fmt.Spruint32f(format, args...))
_, _ = fmt.Fpruint32f(os.Stderr, format+"\n", args...)
}
