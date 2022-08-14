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

package command

import (
	"fmt"
	"os"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/audit"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit solitonAutomata",
	Long:  "Audit solitonAutomata",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}
var auditConfigFile string
var auditConfig audit.Config
var auditResult audit.Result
var auditResultFile string
var auditResultFormat string
var auditResultFormatFile string
var auditResultFormatJSON bool
var auditResultFormatYAML bool
var auditResultFormatTable bool
var auditResultFormatTableVerbose bool
var auditResultFormatTableVerboseHeader bool
var auditResultFormatTableVerboseHeaderSep string
var auditResultFormatTableVerboseHeaderSepLen int
var auditResultFormatTableVerboseHeaderSepLenMax int
var auditResultFormatTableVerboseHeaderSepLenMaxDefault = 10

func newAuditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "audit [audit-id]",
		Short: "Show audit log of solitonAutomata operation",
		RunE: func(cmd *cobra.Command, args []string) error {
			switch len(args) {
			case 0:
				return audit.ShowAuditList(spec.AuditDir())
			case 1:
				return audit.ShowAuditLog(spec.AuditDir(), args[0])
			default:
				return cmd.Help()
			}
		},
	}
	return cmd
}

func newAuditConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Show audit config",
		RunE: func(cmd *cobra.Command, args []string) error {
			return audit.ShowAuditConfig(auditConfigFile)
		},
	}
	return cmd
}

func newAuditResultCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "result",
		Short: "Show audit result",
		RunE: func(cmd *cobra.Command, args []string) error {
			return audit.ShowAuditResult(auditResultFile, auditResultFormat, auditResultFormatJSON, auditResultFormatYAML, auditResultFormatTable, auditResultFormatTableVerbose, auditResultFormatTableVerboseHeader, auditResultFormatTableVerboseHeaderSep, auditResultFormatTableVerboseHeaderSepLen)
		},
	}
	return cmd
}

func newAuditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "audit",
		Short: "Audit solitonAutomata",
		RunE: func(cmd *cobra.Command, args []string) error {
			return audit.Audit(auditConfigFile, auditResultFile, auditResultFormat, auditResultFormatJSON, auditResultFormatYAML, auditResultFormatTable, auditResultFormatTableVerbose, auditResultFormatTableVerboseHeader, auditResultFormatTableVerboseHeaderSep, auditResultFormatTableVerboseHeaderSepLen)
		},
	}
	return cmd
}
