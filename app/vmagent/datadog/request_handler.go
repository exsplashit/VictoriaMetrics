package datadog

import (
	"net/http"

	"github.com/exsplashit/VictoriaMetrics/app/vmagent/common"
	"github.com/exsplashit/VictoriaMetrics/app/vmagent/remotewrite"
	"github.com/exsplashit/VictoriaMetrics/lib/auth"
	"github.com/exsplashit/VictoriaMetrics/lib/prompbmarshal"
	parserCommon "github.com/exsplashit/VictoriaMetrics/lib/protoparser/common"
	parser "github.com/exsplashit/VictoriaMetrics/lib/protoparser/datadog"
	"github.com/exsplashit/VictoriaMetrics/lib/protoparser/datadog/stream"
	"github.com/exsplashit/VictoriaMetrics/lib/tenantmetrics"
	"github.com/VictoriaMetrics/metrics"
)

var (
	rowsInserted       = metrics.NewCounter(`vmagent_rows_inserted_total{type="datadog"}`)
	rowsTenantInserted = tenantmetrics.NewCounterMap(`vmagent_tenant_inserted_rows_total{type="datadog"}`)
	rowsPerInsert      = metrics.NewHistogram(`vmagent_rows_per_insert{type="datadog"}`)
)

// InsertHandlerForHTTP processes remote write for DataDog POST /api/v1/series request.
//
// See https://docs.datadoghq.com/api/latest/metrics/#submit-metrics
func InsertHandlerForHTTP(at *auth.Token, req *http.Request) error {
	extraLabels, err := parserCommon.GetExtraLabels(req)
	if err != nil {
		return err
	}
	ce := req.Header.Get("Content-Encoding")
	return stream.Parse(req.Body, ce, func(series []parser.Series) error {
		return insertRows(at, series, extraLabels)
	})
}

func insertRows(at *auth.Token, series []parser.Series, extraLabels []prompbmarshal.Label) error {
	ctx := common.GetPushCtx()
	defer common.PutPushCtx(ctx)

	rowsTotal := 0
	tssDst := ctx.WriteRequest.Timeseries[:0]
	labels := ctx.Labels[:0]
	samples := ctx.Samples[:0]
	for i := range series {
		ss := &series[i]
		rowsTotal += len(ss.Points)
		labelsLen := len(labels)
		labels = append(labels, prompbmarshal.Label{
			Name:  "__name__",
			Value: ss.Metric,
		})
		if ss.Host != "" {
			labels = append(labels, prompbmarshal.Label{
				Name:  "host",
				Value: ss.Host,
			})
		}
		if ss.Device != "" {
			labels = append(labels, prompbmarshal.Label{
				Name:  "device",
				Value: ss.Device,
			})
		}
		for _, tag := range ss.Tags {
			name, value := parser.SplitTag(tag)
			if name == "host" {
				name = "exported_host"
			}
			labels = append(labels, prompbmarshal.Label{
				Name:  name,
				Value: value,
			})
		}
		labels = append(labels, extraLabels...)
		samplesLen := len(samples)
		for _, pt := range ss.Points {
			samples = append(samples, prompbmarshal.Sample{
				Timestamp: pt.Timestamp(),
				Value:     pt.Value(),
			})
		}
		tssDst = append(tssDst, prompbmarshal.TimeSeries{
			Labels:  labels[labelsLen:],
			Samples: samples[samplesLen:],
		})
	}
	ctx.WriteRequest.Timeseries = tssDst
	ctx.Labels = labels
	ctx.Samples = samples
	remotewrite.Push(at, &ctx.WriteRequest)
	rowsInserted.Add(rowsTotal)
	if at != nil {
		rowsTenantInserted.Get(at).Add(rowsTotal)
	}
	rowsPerInsert.Update(float64(rowsTotal))
	return nil
}
