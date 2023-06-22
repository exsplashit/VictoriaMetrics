// Code generated by qtc from "metrics_find_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/graphite/metrics_find_response.qtpl:1
package graphite

//line app/vmselect/graphite/metrics_find_response.qtpl:1
import (
	"sort"
	"strings"

	"github.com/exsplashit/VictoriaMetrics/lib/logger"
)

// MetricsFindResponse generates response for /metrics/find .See https://graphite-api.readthedocs.io/en/latest/api.html#metrics-find

//line app/vmselect/graphite/metrics_find_response.qtpl:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/graphite/metrics_find_response.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/graphite/metrics_find_response.qtpl:12
func StreamMetricsFindResponse(qw422016 *qt422016.Writer, isPartial bool, paths []string, delimiter, format string, addWildcards bool, jsonp string) {
//line app/vmselect/graphite/metrics_find_response.qtpl:13
	if jsonp != "" {
//line app/vmselect/graphite/metrics_find_response.qtpl:13
		qw422016.N().S(jsonp)
//line app/vmselect/graphite/metrics_find_response.qtpl:13
		qw422016.N().S(`(`)
//line app/vmselect/graphite/metrics_find_response.qtpl:13
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:14
	switch format {
//line app/vmselect/graphite/metrics_find_response.qtpl:15
	case "completer":
//line app/vmselect/graphite/metrics_find_response.qtpl:16
		streammetricsFindResponseCompleter(qw422016, isPartial, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:17
	case "treejson":
//line app/vmselect/graphite/metrics_find_response.qtpl:18
		streammetricsFindResponseTreeJSON(qw422016, isPartial, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:19
	default:
//line app/vmselect/graphite/metrics_find_response.qtpl:20
		logger.Panicf("BUG: unexpected format=%q", format)

//line app/vmselect/graphite/metrics_find_response.qtpl:21
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	if jsonp != "" {
//line app/vmselect/graphite/metrics_find_response.qtpl:22
		qw422016.N().S(`)`)
//line app/vmselect/graphite/metrics_find_response.qtpl:22
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:23
}

//line app/vmselect/graphite/metrics_find_response.qtpl:23
func WriteMetricsFindResponse(qq422016 qtio422016.Writer, isPartial bool, paths []string, delimiter, format string, addWildcards bool, jsonp string) {
//line app/vmselect/graphite/metrics_find_response.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:23
	StreamMetricsFindResponse(qw422016, isPartial, paths, delimiter, format, addWildcards, jsonp)
//line app/vmselect/graphite/metrics_find_response.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:23
}

//line app/vmselect/graphite/metrics_find_response.qtpl:23
func MetricsFindResponse(isPartial bool, paths []string, delimiter, format string, addWildcards bool, jsonp string) string {
//line app/vmselect/graphite/metrics_find_response.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/metrics_find_response.qtpl:23
	WriteMetricsFindResponse(qb422016, isPartial, paths, delimiter, format, addWildcards, jsonp)
//line app/vmselect/graphite/metrics_find_response.qtpl:23
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/metrics_find_response.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:23
	return qs422016
//line app/vmselect/graphite/metrics_find_response.qtpl:23
}

//line app/vmselect/graphite/metrics_find_response.qtpl:25
func streammetricsFindResponseCompleter(qw422016 *qt422016.Writer, isPartial bool, paths []string, delimiter string, addWildcards bool) {
//line app/vmselect/graphite/metrics_find_response.qtpl:25
	qw422016.N().S(`{"metrics":[`)
//line app/vmselect/graphite/metrics_find_response.qtpl:28
	for i, path := range paths {
//line app/vmselect/graphite/metrics_find_response.qtpl:28
		qw422016.N().S(`{"path":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:30
		qw422016.N().Q(path)
//line app/vmselect/graphite/metrics_find_response.qtpl:30
		qw422016.N().S(`,"name":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:31
		streammetricPathName(qw422016, path, delimiter)
//line app/vmselect/graphite/metrics_find_response.qtpl:31
		qw422016.N().S(`,"is_leaf":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:32
		if strings.HasSuffix(path, delimiter) {
//line app/vmselect/graphite/metrics_find_response.qtpl:32
			qw422016.N().S(`0`)
//line app/vmselect/graphite/metrics_find_response.qtpl:32
		} else {
//line app/vmselect/graphite/metrics_find_response.qtpl:32
			qw422016.N().S(`1`)
//line app/vmselect/graphite/metrics_find_response.qtpl:32
		}
//line app/vmselect/graphite/metrics_find_response.qtpl:32
		qw422016.N().S(`}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:34
		if i+1 < len(paths) {
//line app/vmselect/graphite/metrics_find_response.qtpl:34
			qw422016.N().S(`,`)
//line app/vmselect/graphite/metrics_find_response.qtpl:34
		}
//line app/vmselect/graphite/metrics_find_response.qtpl:35
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:36
	if addWildcards && len(paths) > 1 {
//line app/vmselect/graphite/metrics_find_response.qtpl:36
		qw422016.N().S(`,{"name": "*"}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:40
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:40
	qw422016.N().S(`]}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:43
}

//line app/vmselect/graphite/metrics_find_response.qtpl:43
func writemetricsFindResponseCompleter(qq422016 qtio422016.Writer, isPartial bool, paths []string, delimiter string, addWildcards bool) {
//line app/vmselect/graphite/metrics_find_response.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:43
	streammetricsFindResponseCompleter(qw422016, isPartial, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:43
}

//line app/vmselect/graphite/metrics_find_response.qtpl:43
func metricsFindResponseCompleter(isPartial bool, paths []string, delimiter string, addWildcards bool) string {
//line app/vmselect/graphite/metrics_find_response.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/metrics_find_response.qtpl:43
	writemetricsFindResponseCompleter(qb422016, isPartial, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:43
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/metrics_find_response.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:43
	return qs422016
//line app/vmselect/graphite/metrics_find_response.qtpl:43
}

//line app/vmselect/graphite/metrics_find_response.qtpl:45
func streammetricsFindResponseTreeJSON(qw422016 *qt422016.Writer, isPartial bool, paths []string, delimiter string, addWildcards bool) {
//line app/vmselect/graphite/metrics_find_response.qtpl:45
	qw422016.N().S(`[`)
//line app/vmselect/graphite/metrics_find_response.qtpl:48
	if len(paths) > 1 {
		sort.Strings(paths)
		// Substitute `path` and `path<delimiter>` with `path<delimiter><delimiter>`.
		// Such path is treated specially during rendering - see code below for details.
		dst := paths[:1]
		for _, path := range paths[1:] {
			prevPath := dst[len(dst)-1]
			if len(path) == len(prevPath)+1 && strings.HasSuffix(path, delimiter) && strings.HasPrefix(path, prevPath) {
				// The path is equivalent to <prevPath> + <delimiter>
				// Overwrite the prevPath with <path> + <delimiter> as carbonapi does.
				// I.e. the resulting path ends with double delimiter.
				// Such path is treated specially during rendering - see metrics_find_response.qtpl for details.
				dst[len(dst)-1] = path + delimiter
				continue
			}
			dst = append(dst, path)
		}
		paths = dst
	}

//line app/vmselect/graphite/metrics_find_response.qtpl:68
	for i, path := range paths {
//line app/vmselect/graphite/metrics_find_response.qtpl:68
		qw422016.N().S(`{`)
//line app/vmselect/graphite/metrics_find_response.qtpl:71
		id := path
		allowChildren := "0"
		isLeaf := "1"
		if strings.HasSuffix(id, delimiter) {
			if strings.HasSuffix(id[:len(id)-1], delimiter) {
				// Special case when id ends with double delimiter.
				// See the code above for details.
				id = id[:len(id)-2]
			}
			allowChildren = "1"
			isLeaf = "0"
		}

//line app/vmselect/graphite/metrics_find_response.qtpl:83
		qw422016.N().S(`"id":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:84
		qw422016.N().Q(id)
//line app/vmselect/graphite/metrics_find_response.qtpl:84
		qw422016.N().S(`,"text":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:85
		streammetricPathName(qw422016, path, delimiter)
//line app/vmselect/graphite/metrics_find_response.qtpl:85
		qw422016.N().S(`,"allowChildren":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:86
		qw422016.N().S(allowChildren)
//line app/vmselect/graphite/metrics_find_response.qtpl:86
		qw422016.N().S(`,"expandable":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:87
		qw422016.N().S(allowChildren)
//line app/vmselect/graphite/metrics_find_response.qtpl:87
		qw422016.N().S(`,"leaf":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:88
		qw422016.N().S(isLeaf)
//line app/vmselect/graphite/metrics_find_response.qtpl:88
		qw422016.N().S(`}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:90
		if i+1 < len(paths) {
//line app/vmselect/graphite/metrics_find_response.qtpl:90
			qw422016.N().S(`,`)
//line app/vmselect/graphite/metrics_find_response.qtpl:90
		}
//line app/vmselect/graphite/metrics_find_response.qtpl:91
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:92
	if addWildcards && len(paths) > 1 {
//line app/vmselect/graphite/metrics_find_response.qtpl:92
		qw422016.N().S(`,{`)
//line app/vmselect/graphite/metrics_find_response.qtpl:95
		path := paths[0]
		for strings.HasSuffix(path, delimiter) {
			path = path[:len(path)-1]
		}
		id := ""
		if n := strings.LastIndexByte(path, delimiter[0]); n >= 0 {
			id = path[:n+1]
		}
		id += "*"

		allowChildren := "0"
		isLeaf := "1"
		for _, path := range paths {
			if strings.HasSuffix(path, delimiter) {
				allowChildren = "1"
				isLeaf = "0"
				break
			}
		}

//line app/vmselect/graphite/metrics_find_response.qtpl:114
		qw422016.N().S(`"id":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:115
		qw422016.N().Q(id)
//line app/vmselect/graphite/metrics_find_response.qtpl:115
		qw422016.N().S(`,"text": "*","allowChildren":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:117
		qw422016.N().S(allowChildren)
//line app/vmselect/graphite/metrics_find_response.qtpl:117
		qw422016.N().S(`,"expandable":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:118
		qw422016.N().S(allowChildren)
//line app/vmselect/graphite/metrics_find_response.qtpl:118
		qw422016.N().S(`,"leaf":`)
//line app/vmselect/graphite/metrics_find_response.qtpl:119
		qw422016.N().S(isLeaf)
//line app/vmselect/graphite/metrics_find_response.qtpl:119
		qw422016.N().S(`}`)
//line app/vmselect/graphite/metrics_find_response.qtpl:121
	}
//line app/vmselect/graphite/metrics_find_response.qtpl:121
	qw422016.N().S(`]`)
//line app/vmselect/graphite/metrics_find_response.qtpl:123
}

//line app/vmselect/graphite/metrics_find_response.qtpl:123
func writemetricsFindResponseTreeJSON(qq422016 qtio422016.Writer, isPartial bool, paths []string, delimiter string, addWildcards bool) {
//line app/vmselect/graphite/metrics_find_response.qtpl:123
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:123
	streammetricsFindResponseTreeJSON(qw422016, isPartial, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:123
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:123
}

//line app/vmselect/graphite/metrics_find_response.qtpl:123
func metricsFindResponseTreeJSON(isPartial bool, paths []string, delimiter string, addWildcards bool) string {
//line app/vmselect/graphite/metrics_find_response.qtpl:123
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/metrics_find_response.qtpl:123
	writemetricsFindResponseTreeJSON(qb422016, isPartial, paths, delimiter, addWildcards)
//line app/vmselect/graphite/metrics_find_response.qtpl:123
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/metrics_find_response.qtpl:123
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:123
	return qs422016
//line app/vmselect/graphite/metrics_find_response.qtpl:123
}

//line app/vmselect/graphite/metrics_find_response.qtpl:125
func streammetricPathName(qw422016 *qt422016.Writer, path, delimiter string) {
//line app/vmselect/graphite/metrics_find_response.qtpl:127
	name := path
	for strings.HasSuffix(name, delimiter) {
		name = name[:len(name)-1]
	}
	if n := strings.LastIndexByte(name, delimiter[0]); n >= 0 {
		name = name[n+1:]
	}

//line app/vmselect/graphite/metrics_find_response.qtpl:135
	qw422016.N().Q(name)
//line app/vmselect/graphite/metrics_find_response.qtpl:136
}

//line app/vmselect/graphite/metrics_find_response.qtpl:136
func writemetricPathName(qq422016 qtio422016.Writer, path, delimiter string) {
//line app/vmselect/graphite/metrics_find_response.qtpl:136
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:136
	streammetricPathName(qw422016, path, delimiter)
//line app/vmselect/graphite/metrics_find_response.qtpl:136
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:136
}

//line app/vmselect/graphite/metrics_find_response.qtpl:136
func metricPathName(path, delimiter string) string {
//line app/vmselect/graphite/metrics_find_response.qtpl:136
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/metrics_find_response.qtpl:136
	writemetricPathName(qb422016, path, delimiter)
//line app/vmselect/graphite/metrics_find_response.qtpl:136
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/metrics_find_response.qtpl:136
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/metrics_find_response.qtpl:136
	return qs422016
//line app/vmselect/graphite/metrics_find_response.qtpl:136
}
