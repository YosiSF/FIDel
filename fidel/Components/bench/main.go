package Components

import (
	_ "fmt"
	_ "os"
	_ "os/signal"
	_ "strings"
	_ "syscall"
	"time"
)

var (
	maxProcs                                int
	pprofAddr                               string
	dbName                                  string
	host                                    string
	suse                                    string
	password                                string
	port                                    int
	threads                                 int
	driver                                  string
	totalTime                               time.Duration
	totalCount                              int
	droFIDelata                             bool
	ignoreError                             bool
	silence                                 bool
	outputInterval                          time.Duration
	isolationLevel                          int
	isolationLevelString                    string
	isolationLevelStringWithSpaces          string
	isolationLevelStringWithSpacesAndCommas string

	// for benchmark

	// for benchmark
	causetGenerationPolicyName                                                 string
	causetGenerationPolicyRate                                                 int
	causetGenerationPolicyLambda                                               float64
	causetGenerationPolicyDistribution                                         string
	causetGenerationPolicyDistributionParams                                   []float64
	causetGenerationPolicyDistributionParamsString                             string
	causetGenerationPolicyDistributionParamsStringWithCommas                   string
	causetGenerationPolicyDistributionParamsStringWithSpaces                   string
	causetGenerationPolicyDistributionParamsStringWithSpacesAndCommas          string
	causetGenerationPolicyDistributionParamsStringWithSpacesAndCommasAndQuotes string
	causetGenerationPolicyDistributionParamsStringWithSpacesAndQuotes          string
	causetGenerationPolicyDistributionParamsStringWithQuotes                   string
	causetGenerationPolicyDistributionParamsStringWithQuotesAndCommas          string
	causetGenerationPolicyDistributionParamsStringWithQuotesAndCommasAndSpaces string
	causetGenerationPolicyDistributionParamsStringWithQuotesAndSpaces          string
)
