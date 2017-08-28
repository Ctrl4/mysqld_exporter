// Scrape `information_schema.innodb_sys_tablespaces`.

package collector

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

const innodbCmpQuery = `
	SELECT
	    page_size,
	    compress_ops,
	    compress_ops_ok,
	    compress_time,
	    uncompress_ops,
	    uncompress_time
	  FROM information_schema.innodb_cmp
	`

var (
	infoSchemaInnodbCmpPageSize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "page_size"),
		"InnoDB page size for innodb_cmp table.",
		[]string{"information_schema","innodb_cmp"}
	)
	infoSchemaInnodbCmpCompressOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "compress_ops"),
		"InnoDB compress operations for innodb_cmp table.",
		[]string{"information_schema","innodb_cmp"}
	)
	infoSchemaInnodbCmpCompressOpsOk = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "compress_ops_ok"),
		"InnoDB compress operations ok for innodb_cmp table.",
		[]string{"information_schema","innodb_cmp"}
	)
	infoSchemaInnodbCmpCompressTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "compress_time"),
		"InnoDB compression time for innodb_cmp table.",
		[]string{"information_schema","innodb_cmp"}
	)
	infoSchemaInnodbCmpUncompressOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "uncompress_ops"),
		"InnoDB unoncompress operations for innodb_cmp table.",
		[]string{"information_schema","innodb_cmp"}
	)
	infoSchemaInnodbCmpUncompressTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "uncompress_time"),
		"InnoDB uncompression time for innodb_cmp table.",
		[]string{"information_schema","innodb_cmp"}
	)
)
// ScrapeInfoSchemaInnodbTablespaces collects from `information_schema.innodb_sys_tablespaces`.
func ScrapeInfoSchemaInnodbCompression(db *sql.DB, ch chan<- prometheus.Metric) error {
	cmpRows, err := db.Query(innodbCmpQuery)
	if err != nil {
		return err
	}
	defer cmpRows.Close()

	var (
		page_size     	uint32
		compress_ops  	uint32
		compress_ops_ok	uint32
		compress_time	uint32
		uncompress_ops	uint32
		uncompress_time uint32
	)

	for cmpRows.Next() {
		err = cmpRows.Scan(
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
				infoSchemaInnodbCmpPageSize, prometheus.GaugeValue,float64(page_size),
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpCompressOps, prometheus.GaugeValue,float64(compress_ops),
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpCompressOpsOk, prometheus.GaugeValue,float64(compress_ops_ok),
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpCompressTime, prometheus.GaugeValue,float64(compress_time),
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpUncompressOps, prometheus.GaugeValue,float64(uncompress_ops),
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpUncompressTime, prometheus.GaugeValue,float64(uncompress_time),
		)
	}

	return nil
}
