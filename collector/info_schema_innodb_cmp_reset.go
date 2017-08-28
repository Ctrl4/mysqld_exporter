// Scrape `information_schema.innodb_sys_tablespaces`.

package collector

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"

	"strconv"
)

const innodbCmpResetQuery = `
	SELECT
	    page_size,
	    compress_ops,
	    compress_ops_ok,
	    compress_time,
	    uncompress_ops,
	    uncompress_time
	  FROM information_schema.innodb_cmp_reset
	`

var (
	infoSchemaInnodbCmpResetPageSize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "cmp_page_size"),
		"InnoDB page size for innodb_cmp_reset table.",
		[]string{"page_size"}, nil,
	)
	infoSchemaInnodbCmpResetCompressOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "cmp_reset_compress_ops"),
		"InnoDB compress operations for innodb_cmp_reset table.",
		[]string{"page_size"}, nil,
	)
	infoSchemaInnodbCmpResetCompressOpsOk = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "cmp_reset_compress_ops_ok"),
		"InnoDB compress operations ok for innodb_cmp_reset table.",
		[]string{"page_size"}, nil,
	)
	infoSchemaInnodbCmpResetCompressTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "cmp_reset_compress_time"),
		"InnoDB compression time for innodb_cmp_reset table.",
		[]string{"page_size"}, nil,
	)
	infoSchemaInnodbCmpResetUncompressOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "cmp_reset_uncompress_ops"),
		"InnoDB unoncompress operations for innodb_cmp_reset table.",
		[]string{"page_size"}, nil,
	)
	infoSchemaInnodbCmpResetUncompressTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "cmp_reset_uncompress_time"),
		"InnoDB uncompression time for innodb_cmp_reset table.",
		[]string{"page_size"}, nil,
	)
)
// ScrapeInfoSchemaInnodbTablespaces collects from `information_schema.innodb_sys_tablespaces`.
func ScrapeInfoSchemaInnodbCompressionReset(db *sql.DB, ch chan<- prometheus.Metric) error {
	CmpResetRows, err := db.Query(innodbCmpResetQuery)
	if err != nil {
		return err
	}
	defer CmpResetRows.Close()

	var (
		page_size     	float64
		compress_ops  	float64
		compress_ops_ok	float64
		compress_time	float64
		uncompress_ops	float64
		uncompress_time float64
	)

	for CmpResetRows.Next() {
		err = CmpResetRows.Scan(
			&page_size,
			&compress_ops,
			&compress_ops_ok,
			&compress_time,
			&uncompress_ops,
			&uncompress_time,
		)
		if err != nil {
			return err
		}
	ch <- prometheus.MustNewConstMetric(
		infoSchemaInnodbCmpResetPageSize, prometheus.GaugeValue, page_size,
		strconv.FormatFloat(page_size,'f',-1,64),
	)
	ch <- prometheus.MustNewConstMetric(
		infoSchemaInnodbCmpResetCompressOps, prometheus.GaugeValue, compress_ops,
		strconv.FormatFloat(page_size,'f', -1, 64),
		)
	ch <- prometheus.MustNewConstMetric(
		infoSchemaInnodbCmpResetCompressOpsOk, prometheus.GaugeValue,compress_ops_ok,
		strconv.FormatFloat(page_size,'f', -1, 64),
	)
	ch <- prometheus.MustNewConstMetric(
		infoSchemaInnodbCmpResetCompressTime, prometheus.GaugeValue,compress_time,
		strconv.FormatFloat(page_size,'f', -1, 64),
	)
	ch <- prometheus.MustNewConstMetric(
		infoSchemaInnodbCmpResetUncompressOps, prometheus.GaugeValue,uncompress_ops,
		strconv.FormatFloat(page_size,'f', -1, 64),
		)
	ch <- prometheus.MustNewConstMetric(
		infoSchemaInnodbCmpResetUncompressTime, prometheus.GaugeValue,uncompress_time,
		strconv.FormatFloat(page_size,'f', -1, 64),
		)
	}

	return nil
}
