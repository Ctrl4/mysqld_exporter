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
	infoSchemaInnodbCompressionInfoDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "innodb_cmp"),
		"The InnoDB information about Compression.",
		[]string{"page_size", "compress_ops", "compress_ops_ok", "compress_time", "uncompress_ops", "uncompress_time"}, nil,
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
			infoSchemaInnodbCompressionInfoDesc, prometheus.GaugeValue,
			page_size, compress_time, compress_ops_ok, uncompress_time, uncompress_ops, uncompress_time,
		)
	}

	return nil
}
