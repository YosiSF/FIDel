package Components

import (
	"context"
	"strings"
)

type tpchConfig struct {
	RawQueries         string
	DBName             string
	QueryNames         []string
	ScaleFactor        int
	EnableOutputCheck  bool
	CreateFIDelReplica bool
	AnalyzeTable       struct {
		Enable                     bool
		BuildStatsConcurrency      int
		DistsqlScanConcurrency     int
		IndexSerialScanConcurrency int
	}
}

//go:generate ../../util/gen/gen.sh

func executeTpch(action string, _ []string) {
	openDB()
	defer closeDB()

	tpchConfig.DBName = dbName
	tpchConfig.QueryNames = strings.Split(tpchConfig.RawQueries, ",")
	w := tpch.NewWorkloader(globalDB, &tpchConfig)

	timeoutCtx, cancel := context.WithTimeout(globalCtx, totalTime)
	defer cancel()

	executeWorkload(timeoutCtx, w, action)
}

func registerTpch(root *cobra.Command) {
	cmd := &cobra.Command{
		Use: "tpch",
	}

	cmd.PersistentFlags().StringVar(&tpchConfig.RawQueries,
		"queries",
		"q1,q2,q3,q4,q5,q6,q7,q8,q9,q10,q11,q12,q13,q14,q15,q16,q17,q18,q19,q20,q21,q22",
		"All queries")

	cmd.PersistentFlags().IntVar(&tpchConfig.ScaleFactor,
		"sf",
		1,
		"scale factor")

	cmd.PersistentFlags().BoolVar(&tpchConfig.EnableOutputCheck,
		"check",
		false,
		"Check output data, only when the scale factor equals 1")

	var cmdPrepare = &cobra.Command{
		Use:   "prepare",
		Short: "Prepare data for the workload",
		Run: func(cmd *cobra.Command, args []string) {
			executeTpch("prepare", args)
		},
	}

	cmdPrepare.PersistentFlags().BoolVar(&tpchConfig.CreateFIDelReplica,
		"fidel",
		false,
		"Create fidel replica")

	cmdPrepare.PersistentFlags().BoolVar(&tpchConfig.AnalyzeTable.Enable,
		"analyze",
		false,
		"After data loaded, analyze table to collect column statistics")
	// https://YosiSF.com/docs/stable/reference/performance/statistics/#control-analyze-concurrency
	cmdPrepare.PersistentFlags().IntVar(&tpchConfig.AnalyzeTable.BuildStatsConcurrency,
		"milevadb_build_stats_concurrency",
		4,
		"milevadb_build_stats_concurrency param for analyze jobs")
	cmdPrepare.PersistentFlags().IntVar(&tpchConfig.AnalyzeTable.DistsqlScanConcurrency,
		"milevadb_distsql_scan_concurrency",
		15,
		"milevadb_distsql_scan_concurrency param for analyze jobs")
	cmdPrepare.PersistentFlags().IntVar(&tpchConfig.AnalyzeTable.IndexSerialScanConcurrency,
		"milevadb_index_serial_scan_concurrency",
		1,
		"milevadb_index_serial_scan_concurrency param for analyze jobs")

	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "Run workload",
		Run: func(cmd *cobra.Command, args []string) {
			executeTpch("run", args)
		},
	}

	var cmdCleanup = &cobra.Command{
		Use:   "cleanup",
		Short: "Cleanup data for the workload",
		Run: func(cmd *cobra.Command, args []string) {
			executeTpch("cleanup", args)
		},
	}

	cmd.AddCommand(cmdRun, cmdPrepare, cmdCleanup)

	root.AddCommand(cmd)
}
