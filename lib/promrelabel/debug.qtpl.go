// Code generated by qtc from "debug.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line lib/promrelabel/debug.qtpl:1
package promrelabel

//line lib/promrelabel/debug.qtpl:1
import (
	"fmt"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/htmlcomponents"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promutils"
)

//line lib/promrelabel/debug.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/promrelabel/debug.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/promrelabel/debug.qtpl:9
func StreamRelabelDebugSteps(qw422016 *qt422016.Writer, targetURL, targetID, format string, dss []DebugStep, metric, relabelConfigs string, err error) {
//line lib/promrelabel/debug.qtpl:10
	if format == "json" {
//line lib/promrelabel/debug.qtpl:11
		StreamRelabelDebugStepsJSON(qw422016, targetURL, targetID, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:12
	} else {
//line lib/promrelabel/debug.qtpl:13
		StreamRelabelDebugStepsHTML(qw422016, targetURL, targetID, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:14
	}
//line lib/promrelabel/debug.qtpl:15
}

//line lib/promrelabel/debug.qtpl:15
func WriteRelabelDebugSteps(qq422016 qtio422016.Writer, targetURL, targetID, format string, dss []DebugStep, metric, relabelConfigs string, err error) {
//line lib/promrelabel/debug.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:15
	StreamRelabelDebugSteps(qw422016, targetURL, targetID, format, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:15
}

//line lib/promrelabel/debug.qtpl:15
func RelabelDebugSteps(targetURL, targetID, format string, dss []DebugStep, metric, relabelConfigs string, err error) string {
//line lib/promrelabel/debug.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:15
	WriteRelabelDebugSteps(qb422016, targetURL, targetID, format, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:15
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:15
	return qs422016
//line lib/promrelabel/debug.qtpl:15
}

//line lib/promrelabel/debug.qtpl:17
func StreamRelabelDebugStepsHTML(qw422016 *qt422016.Writer, targetURL, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) {
//line lib/promrelabel/debug.qtpl:17
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head>`)
//line lib/promrelabel/debug.qtpl:21
	htmlcomponents.StreamCommonHeader(qw422016)
//line lib/promrelabel/debug.qtpl:21
	qw422016.N().S(`<title>Metric relabel debug</title><script>function submitRelabelDebugForm(e) {var form = e.target;var method = "GET";if (form.elements["relabel_configs"].value.length + form.elements["metric"].value.length > 1000) {method = "POST";}form.method = method;}</script></head><body>`)
//line lib/promrelabel/debug.qtpl:35
	htmlcomponents.StreamNavbar(qw422016)
//line lib/promrelabel/debug.qtpl:35
	qw422016.N().S(`<div class="container-fluid"><a href="https://docs.victoriametrics.com/relabeling.html" target="_blank">Relabeling docs</a>`)
//line lib/promrelabel/debug.qtpl:37
	qw422016.N().S(` `)
//line lib/promrelabel/debug.qtpl:39
	if targetURL != "" {
//line lib/promrelabel/debug.qtpl:39
		qw422016.N().S(`<a href="metric-relabel-debug`)
//line lib/promrelabel/debug.qtpl:40
		if targetID != "" {
//line lib/promrelabel/debug.qtpl:40
			qw422016.N().S(`?id=`)
//line lib/promrelabel/debug.qtpl:40
			qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:40
		}
//line lib/promrelabel/debug.qtpl:40
		qw422016.N().S(`">Metric relabel debug</a>`)
//line lib/promrelabel/debug.qtpl:41
	} else {
//line lib/promrelabel/debug.qtpl:41
		qw422016.N().S(`<a href="target-relabel-debug`)
//line lib/promrelabel/debug.qtpl:42
		if targetID != "" {
//line lib/promrelabel/debug.qtpl:42
			qw422016.N().S(`?id=`)
//line lib/promrelabel/debug.qtpl:42
			qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:42
		}
//line lib/promrelabel/debug.qtpl:42
		qw422016.N().S(`">Target relabel debug</a>`)
//line lib/promrelabel/debug.qtpl:43
	}
//line lib/promrelabel/debug.qtpl:43
	qw422016.N().S(`<br>`)
//line lib/promrelabel/debug.qtpl:46
	if err != nil {
//line lib/promrelabel/debug.qtpl:47
		htmlcomponents.StreamErrorNotification(qw422016, err)
//line lib/promrelabel/debug.qtpl:48
	}
//line lib/promrelabel/debug.qtpl:48
	qw422016.N().S(`<div class="m-3"><form method="POST" onsubmit="submitRelabelDebugForm(event)">`)
//line lib/promrelabel/debug.qtpl:52
	streamrelabelDebugFormInputs(qw422016, metric, relabelConfigs)
//line lib/promrelabel/debug.qtpl:53
	if targetID != "" {
//line lib/promrelabel/debug.qtpl:53
		qw422016.N().S(`<input type="hidden" name="id" value="`)
//line lib/promrelabel/debug.qtpl:54
		qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:54
		qw422016.N().S(`" />`)
//line lib/promrelabel/debug.qtpl:55
	}
//line lib/promrelabel/debug.qtpl:55
	qw422016.N().S(`<input type="submit" value="Submit" class="btn btn-primary m-1" />`)
//line lib/promrelabel/debug.qtpl:57
	if targetID != "" {
//line lib/promrelabel/debug.qtpl:57
		qw422016.N().S(`<button type="button" onclick="location.href='?id=`)
//line lib/promrelabel/debug.qtpl:58
		qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:58
		qw422016.N().S(`'" class="btn btn-secondary m-1">Reset</button>`)
//line lib/promrelabel/debug.qtpl:59
	}
//line lib/promrelabel/debug.qtpl:59
	qw422016.N().S(`</form></div><div class="row"><main class="col-12">`)
//line lib/promrelabel/debug.qtpl:65
	streamrelabelDebugSteps(qw422016, dss, targetURL, targetID)
//line lib/promrelabel/debug.qtpl:65
	qw422016.N().S(`</main></div></div></body></html>`)
//line lib/promrelabel/debug.qtpl:71
}

//line lib/promrelabel/debug.qtpl:71
func WriteRelabelDebugStepsHTML(qq422016 qtio422016.Writer, targetURL, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) {
//line lib/promrelabel/debug.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:71
	StreamRelabelDebugStepsHTML(qw422016, targetURL, targetID, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:71
}

//line lib/promrelabel/debug.qtpl:71
func RelabelDebugStepsHTML(targetURL, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) string {
//line lib/promrelabel/debug.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:71
	WriteRelabelDebugStepsHTML(qb422016, targetURL, targetID, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:71
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:71
	return qs422016
//line lib/promrelabel/debug.qtpl:71
}

//line lib/promrelabel/debug.qtpl:73
func streamrelabelDebugFormInputs(qw422016 *qt422016.Writer, metric, relabelConfigs string) {
//line lib/promrelabel/debug.qtpl:73
	qw422016.N().S(`<div>Relabel configs:<br/><textarea name="relabel_configs" style="width: 100%; height: 15em" class="m-1">`)
//line lib/promrelabel/debug.qtpl:76
	qw422016.E().S(relabelConfigs)
//line lib/promrelabel/debug.qtpl:76
	qw422016.N().S(`</textarea></div><div>Labels:<br/><textarea name="metric" style="width: 100%; height: 5em" class="m-1">`)
//line lib/promrelabel/debug.qtpl:81
	qw422016.E().S(metric)
//line lib/promrelabel/debug.qtpl:81
	qw422016.N().S(`</textarea></div>`)
//line lib/promrelabel/debug.qtpl:83
}

//line lib/promrelabel/debug.qtpl:83
func writerelabelDebugFormInputs(qq422016 qtio422016.Writer, metric, relabelConfigs string) {
//line lib/promrelabel/debug.qtpl:83
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:83
	streamrelabelDebugFormInputs(qw422016, metric, relabelConfigs)
//line lib/promrelabel/debug.qtpl:83
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:83
}

//line lib/promrelabel/debug.qtpl:83
func relabelDebugFormInputs(metric, relabelConfigs string) string {
//line lib/promrelabel/debug.qtpl:83
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:83
	writerelabelDebugFormInputs(qb422016, metric, relabelConfigs)
//line lib/promrelabel/debug.qtpl:83
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:83
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:83
	return qs422016
//line lib/promrelabel/debug.qtpl:83
}

//line lib/promrelabel/debug.qtpl:85
func streamrelabelDebugSteps(qw422016 *qt422016.Writer, dss []DebugStep, targetURL, targetID string) {
//line lib/promrelabel/debug.qtpl:86
	if len(dss) > 0 {
//line lib/promrelabel/debug.qtpl:86
		qw422016.N().S(`<div class="m-3"><b>Original labels:</b> <samp>`)
//line lib/promrelabel/debug.qtpl:88
		streammustFormatLabels(qw422016, dss[0].In)
//line lib/promrelabel/debug.qtpl:88
		qw422016.N().S(`</samp></div>`)
//line lib/promrelabel/debug.qtpl:90
	}
//line lib/promrelabel/debug.qtpl:90
	qw422016.N().S(`<table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col" style="width: 5%">Step</th><th scope="col" style="width: 25%">Relabeling Rule</th><th scope="col" style="width: 35%">Input Labels</th><th scope="col" stile="width: 35%">Output labels</a></tr></thead><tbody>`)
//line lib/promrelabel/debug.qtpl:101
	for i, ds := range dss {
//line lib/promrelabel/debug.qtpl:103
		inLabels := promutils.MustNewLabelsFromString(ds.In)
		outLabels := promutils.MustNewLabelsFromString(ds.Out)
		changedLabels := getChangedLabelNames(inLabels, outLabels)

//line lib/promrelabel/debug.qtpl:106
		qw422016.N().S(`<tr><td>`)
//line lib/promrelabel/debug.qtpl:108
		qw422016.N().D(i)
//line lib/promrelabel/debug.qtpl:108
		qw422016.N().S(`</td><td><b><pre class="m-2">`)
//line lib/promrelabel/debug.qtpl:109
		qw422016.E().S(ds.Rule)
//line lib/promrelabel/debug.qtpl:109
		qw422016.N().S(`</pre></b></td><td><div class="m-2" style="font-size: 0.9em" title="deleted and updated labels highlighted in red">`)
//line lib/promrelabel/debug.qtpl:112
		streamlabelsWithHighlight(qw422016, inLabels, changedLabels, "red")
//line lib/promrelabel/debug.qtpl:112
		qw422016.N().S(`</div></td><td><div class="m-2" style="font-size: 0.9em" title="added and updated labels highlighted in blue">`)
//line lib/promrelabel/debug.qtpl:117
		streamlabelsWithHighlight(qw422016, outLabels, changedLabels, "blue")
//line lib/promrelabel/debug.qtpl:117
		qw422016.N().S(`</div></td></tr>`)
//line lib/promrelabel/debug.qtpl:121
	}
//line lib/promrelabel/debug.qtpl:121
	qw422016.N().S(`</tbody></table>`)
//line lib/promrelabel/debug.qtpl:124
	if len(dss) > 0 {
//line lib/promrelabel/debug.qtpl:124
		qw422016.N().S(`<div class="m-3"><b>Resulting labels:</b> <samp>`)
//line lib/promrelabel/debug.qtpl:126
		streammustFormatLabels(qw422016, dss[len(dss)-1].Out)
//line lib/promrelabel/debug.qtpl:126
		qw422016.N().S(`</samp>`)
//line lib/promrelabel/debug.qtpl:127
		if targetURL != "" {
//line lib/promrelabel/debug.qtpl:127
			qw422016.N().S(`<div><b>Target URL:</b>`)
//line lib/promrelabel/debug.qtpl:129
			qw422016.N().S(` `)
//line lib/promrelabel/debug.qtpl:129
			qw422016.N().S(`<a href="`)
//line lib/promrelabel/debug.qtpl:129
			qw422016.E().S(targetURL)
//line lib/promrelabel/debug.qtpl:129
			qw422016.N().S(`" target="_blank">`)
//line lib/promrelabel/debug.qtpl:129
			qw422016.E().S(targetURL)
//line lib/promrelabel/debug.qtpl:129
			qw422016.N().S(`</a>`)
//line lib/promrelabel/debug.qtpl:130
			if targetID != "" {
//line lib/promrelabel/debug.qtpl:131
				qw422016.N().S(` `)
//line lib/promrelabel/debug.qtpl:131
				qw422016.N().S(`(<a href="target_response?id=`)
//line lib/promrelabel/debug.qtpl:132
				qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:132
				qw422016.N().S(`" target="_blank" title="click to fetch target response on behalf of the scraper">response</a>)`)
//line lib/promrelabel/debug.qtpl:133
			}
//line lib/promrelabel/debug.qtpl:133
			qw422016.N().S(`</div>`)
//line lib/promrelabel/debug.qtpl:135
		}
//line lib/promrelabel/debug.qtpl:135
		qw422016.N().S(`</div>`)
//line lib/promrelabel/debug.qtpl:137
	}
//line lib/promrelabel/debug.qtpl:138
}

//line lib/promrelabel/debug.qtpl:138
func writerelabelDebugSteps(qq422016 qtio422016.Writer, dss []DebugStep, targetURL, targetID string) {
//line lib/promrelabel/debug.qtpl:138
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:138
	streamrelabelDebugSteps(qw422016, dss, targetURL, targetID)
//line lib/promrelabel/debug.qtpl:138
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:138
}

//line lib/promrelabel/debug.qtpl:138
func relabelDebugSteps(dss []DebugStep, targetURL, targetID string) string {
//line lib/promrelabel/debug.qtpl:138
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:138
	writerelabelDebugSteps(qb422016, dss, targetURL, targetID)
//line lib/promrelabel/debug.qtpl:138
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:138
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:138
	return qs422016
//line lib/promrelabel/debug.qtpl:138
}

//line lib/promrelabel/debug.qtpl:140
func StreamRelabelDebugStepsJSON(qw422016 *qt422016.Writer, targetURL, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) {
//line lib/promrelabel/debug.qtpl:140
	qw422016.N().S(`{`)
//line lib/promrelabel/debug.qtpl:142
	if err != nil {
//line lib/promrelabel/debug.qtpl:142
		qw422016.N().S(`"status": "error","error":`)
//line lib/promrelabel/debug.qtpl:144
		qw422016.N().Q(fmt.Sprintf("Error: %s", err))
//line lib/promrelabel/debug.qtpl:145
	} else {
//line lib/promrelabel/debug.qtpl:145
		qw422016.N().S(`"status": "success",`)
//line lib/promrelabel/debug.qtpl:147
		if len(dss) > 0 {
//line lib/promrelabel/debug.qtpl:147
			qw422016.N().S(`"originalLabels":`)
//line lib/promrelabel/debug.qtpl:148
			qw422016.N().Q(mustFormatLabels(dss[0].In))
//line lib/promrelabel/debug.qtpl:148
			qw422016.N().S(`,"resultingLabels":`)
//line lib/promrelabel/debug.qtpl:149
			qw422016.N().Q(mustFormatLabels(dss[len(dss)-1].Out))
//line lib/promrelabel/debug.qtpl:149
			qw422016.N().S(`,`)
//line lib/promrelabel/debug.qtpl:150
		}
//line lib/promrelabel/debug.qtpl:150
		qw422016.N().S(`"steps": [`)
//line lib/promrelabel/debug.qtpl:152
		for i, ds := range dss {
//line lib/promrelabel/debug.qtpl:154
			inLabels := promutils.MustNewLabelsFromString(ds.In)
			outLabels := promutils.MustNewLabelsFromString(ds.Out)
			changedLabels := getChangedLabelNames(inLabels, outLabels)

//line lib/promrelabel/debug.qtpl:157
			qw422016.N().S(`{"inLabels":`)
//line lib/promrelabel/debug.qtpl:159
			qw422016.N().Q(labelsWithHighlight(inLabels, changedLabels, "red"))
//line lib/promrelabel/debug.qtpl:159
			qw422016.N().S(`,"outLabels":`)
//line lib/promrelabel/debug.qtpl:160
			qw422016.N().Q(labelsWithHighlight(outLabels, changedLabels, "blue"))
//line lib/promrelabel/debug.qtpl:160
			qw422016.N().S(`,"rule":`)
//line lib/promrelabel/debug.qtpl:161
			qw422016.N().Q(ds.Rule)
//line lib/promrelabel/debug.qtpl:161
			qw422016.N().S(`}`)
//line lib/promrelabel/debug.qtpl:163
			if i != len(dss)-1 {
//line lib/promrelabel/debug.qtpl:163
				qw422016.N().S(`,`)
//line lib/promrelabel/debug.qtpl:163
			}
//line lib/promrelabel/debug.qtpl:164
		}
//line lib/promrelabel/debug.qtpl:164
		qw422016.N().S(`]`)
//line lib/promrelabel/debug.qtpl:166
	}
//line lib/promrelabel/debug.qtpl:166
	qw422016.N().S(`}`)
//line lib/promrelabel/debug.qtpl:168
}

//line lib/promrelabel/debug.qtpl:168
func WriteRelabelDebugStepsJSON(qq422016 qtio422016.Writer, targetURL, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) {
//line lib/promrelabel/debug.qtpl:168
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:168
	StreamRelabelDebugStepsJSON(qw422016, targetURL, targetID, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:168
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:168
}

//line lib/promrelabel/debug.qtpl:168
func RelabelDebugStepsJSON(targetURL, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) string {
//line lib/promrelabel/debug.qtpl:168
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:168
	WriteRelabelDebugStepsJSON(qb422016, targetURL, targetID, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:168
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:168
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:168
	return qs422016
//line lib/promrelabel/debug.qtpl:168
}

//line lib/promrelabel/debug.qtpl:170
func streamlabelsWithHighlight(qw422016 *qt422016.Writer, labels *promutils.Labels, highlight map[string]struct{}, color string) {
//line lib/promrelabel/debug.qtpl:172
	labelsList := labels.GetLabels()
	metricName := ""
	for i, label := range labelsList {
		if label.Name == "__name__" {
			metricName = label.Value
			labelsList = append(labelsList[:i], labelsList[i+1:]...)
			break
		}
	}

//line lib/promrelabel/debug.qtpl:182
	if metricName != "" {
//line lib/promrelabel/debug.qtpl:183
		if _, ok := highlight["__name__"]; ok {
//line lib/promrelabel/debug.qtpl:183
			qw422016.N().S(`<span style="font-weight:bold;color:`)
//line lib/promrelabel/debug.qtpl:184
			qw422016.E().S(color)
//line lib/promrelabel/debug.qtpl:184
			qw422016.N().S(`">`)
//line lib/promrelabel/debug.qtpl:184
			qw422016.E().S(metricName)
//line lib/promrelabel/debug.qtpl:184
			qw422016.N().S(`</span>`)
//line lib/promrelabel/debug.qtpl:185
		} else {
//line lib/promrelabel/debug.qtpl:186
			qw422016.E().S(metricName)
//line lib/promrelabel/debug.qtpl:187
		}
//line lib/promrelabel/debug.qtpl:188
		if len(labelsList) == 0 {
//line lib/promrelabel/debug.qtpl:188
			return
//line lib/promrelabel/debug.qtpl:188
		}
//line lib/promrelabel/debug.qtpl:189
	}
//line lib/promrelabel/debug.qtpl:189
	qw422016.N().S(`{`)
//line lib/promrelabel/debug.qtpl:191
	for i, label := range labelsList {
//line lib/promrelabel/debug.qtpl:192
		if _, ok := highlight[label.Name]; ok {
//line lib/promrelabel/debug.qtpl:192
			qw422016.N().S(`<span style="font-weight:bold;color:`)
//line lib/promrelabel/debug.qtpl:193
			qw422016.E().S(color)
//line lib/promrelabel/debug.qtpl:193
			qw422016.N().S(`">`)
//line lib/promrelabel/debug.qtpl:193
			qw422016.E().S(label.Name)
//line lib/promrelabel/debug.qtpl:193
			qw422016.N().S(`=`)
//line lib/promrelabel/debug.qtpl:193
			qw422016.E().Q(label.Value)
//line lib/promrelabel/debug.qtpl:193
			qw422016.N().S(`</span>`)
//line lib/promrelabel/debug.qtpl:194
		} else {
//line lib/promrelabel/debug.qtpl:195
			qw422016.E().S(label.Name)
//line lib/promrelabel/debug.qtpl:195
			qw422016.N().S(`=`)
//line lib/promrelabel/debug.qtpl:195
			qw422016.E().Q(label.Value)
//line lib/promrelabel/debug.qtpl:196
		}
//line lib/promrelabel/debug.qtpl:197
		if i < len(labelsList)-1 {
//line lib/promrelabel/debug.qtpl:197
			qw422016.N().S(`,`)
//line lib/promrelabel/debug.qtpl:197
			qw422016.N().S(` `)
//line lib/promrelabel/debug.qtpl:197
		}
//line lib/promrelabel/debug.qtpl:198
	}
//line lib/promrelabel/debug.qtpl:198
	qw422016.N().S(`}`)
//line lib/promrelabel/debug.qtpl:200
}

//line lib/promrelabel/debug.qtpl:200
func writelabelsWithHighlight(qq422016 qtio422016.Writer, labels *promutils.Labels, highlight map[string]struct{}, color string) {
//line lib/promrelabel/debug.qtpl:200
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:200
	streamlabelsWithHighlight(qw422016, labels, highlight, color)
//line lib/promrelabel/debug.qtpl:200
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:200
}

//line lib/promrelabel/debug.qtpl:200
func labelsWithHighlight(labels *promutils.Labels, highlight map[string]struct{}, color string) string {
//line lib/promrelabel/debug.qtpl:200
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:200
	writelabelsWithHighlight(qb422016, labels, highlight, color)
//line lib/promrelabel/debug.qtpl:200
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:200
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:200
	return qs422016
//line lib/promrelabel/debug.qtpl:200
}

//line lib/promrelabel/debug.qtpl:202
func streammustFormatLabels(qw422016 *qt422016.Writer, s string) {
//line lib/promrelabel/debug.qtpl:203
	labels := promutils.MustNewLabelsFromString(s)

//line lib/promrelabel/debug.qtpl:204
	streamlabelsWithHighlight(qw422016, labels, nil, "")
//line lib/promrelabel/debug.qtpl:205
}

//line lib/promrelabel/debug.qtpl:205
func writemustFormatLabels(qq422016 qtio422016.Writer, s string) {
//line lib/promrelabel/debug.qtpl:205
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:205
	streammustFormatLabels(qw422016, s)
//line lib/promrelabel/debug.qtpl:205
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:205
}

//line lib/promrelabel/debug.qtpl:205
func mustFormatLabels(s string) string {
//line lib/promrelabel/debug.qtpl:205
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:205
	writemustFormatLabels(qb422016, s)
//line lib/promrelabel/debug.qtpl:205
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:205
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:205
	return qs422016
//line lib/promrelabel/debug.qtpl:205
}
