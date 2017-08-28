// Scrape `information_schema.innodb_sys_tablespaces`.

package collector

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"

	"strconv"
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
		[]string{"page size"}, nil,
	)
	infoSchemaInnodbCmpCompressOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "compress_ops"),
		"InnoDB compress operations for innodb_cmp table.",
		[]string{"schema","table","column"}, nil,
	)
	infoSchemaInnodbCmpCompressOpsOk = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "compress_ops_ok"),
		"InnoDB compress operations ok for innodb_cmp table.",
		[]string{"schema","table","column"}, nil,
	)
	infoSchemaInnodbCmpCompressTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "compress_time"),
		"InnoDB compression time for innodb_cmp table.",
		[]string{"schema","table","column"}, nil,
	)
	infoSchemaInnodbCmpUncompressOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "uncompress_ops"),
		"InnoDB unoncompress operations for innodb_cmp table.",
		[]string{"schema","table","column"}, nil,
	)
	infoSchemaInnodbCmpUncompressTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, informationSchema, "uncompress_time"),
		"InnoDB uncompression time for innodb_cmp table.",
		[]string{"schema","table","column"}, nil,
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
		page_size     	float64
		compress_ops  	float64
		compress_ops_ok	float64
		compress_time	float64
		uncompress_ops	float64
		uncompress_time float64
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
		infoSchemaInnodbCmpPageSize, prometheus.GaugeValue, page_size,
		strconv.FormatFloat(page_size,'f',-1,64),
	)
/*
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpPageSize, prometheus.GaugeValue,float64(page_size),
				"information_schema", "innodb_cmp",
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpCompressOps, prometheus.GaugeValue,float64(compress_ops),
				"information_schema", "innodb_cmp",
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpCompressOpsOk, prometheus.GaugeValue,float64(compress_ops_ok),
				"information_schema", "innodb_cmp",
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpCompressTime, prometheus.GaugeValue,float64(compress_time),
				"information_schema", "innodb_cmp",
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpUncompressOps, prometheus.GaugeValue,float64(uncompress_ops),
				"information_schema", "innodb_cmp",
		)
	ch <- prometheus.MustNewConstMetric(
				infoSchemaInnodbCmpUncompressTime, prometheus.GaugeValue,float64(uncompress_time),
				"information_schema", "innodb_cmp",
		)*/
	}

	return nil
}
