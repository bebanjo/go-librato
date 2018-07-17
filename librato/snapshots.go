package librato

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// ChartSnapshot represents a snapshot of a Librato Chart.
type ChartSnapshot struct {
	Subject  *ChartSnapshotSubject `json:"subject,omitempty"`
	Duration *int                  `json:"duration,omitempty"`
	EndTime  *string               `json:"end_time,omitempty"`
	Image    *string               `json:"image_href,omitempty"`
	URL      *string               `json:"href,omitempty"`
}

// ChartSnapshotSubject represents the subject in a Librato Snapshot.
type ChartSnapshotSubject struct {
	Chart *ChartSnapshotInfo `json:"chart,omitempty"`
}

// ChartSnapshotInfo represents the information of the Chart in a Librato Snapshot.
type ChartSnapshotInfo struct {
	ID     *uint   `json:"id,omitempty"`
	Source *string `json:"source,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// CreateChartSnaphot creates a snapshot of chart given a Librato chart, duration, source and type.
//
// Librato API docs: http://dev.librato.com/v1/snapshots
func (s *SpacesService) CreateChartSnapshot(chartID uint, duration int, endTime *time.Time, source, chartType string) (*ChartSnapshot, *http.Response, error) {
	chartSnaphotInfo := &ChartSnapshotInfo{
		ID:     &chartID,
		Source: &source,
		Type:   &chartType,
	}

	chartSnapshotSubject := &ChartSnapshotSubject{
		Chart: chartSnaphotInfo,
	}

	var unixEndTime *string
	if endTime != nil {
		t := *endTime
		s := fmt.Sprintf("%d", t.Unix())
		unixEndTime = &s
	}

	chartSnapshot := &ChartSnapshot{
		Subject:  chartSnapshotSubject,
		Duration: &duration,
		EndTime:  unixEndTime,
	}

	req, err := s.client.NewRequest("POST", "snapshots", chartSnapshot)
	if err != nil {
		return nil, nil, err
	}

	cs := new(ChartSnapshot)
	resp, err := s.client.Do(req, cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, err
}

// GetChartSnapshot gets a chart snapshot with a given Snapshot URL.
//
// Librato API docs: http://dev.librato.com/v1/get/spaces/:id/charts
func (s *SpacesService) GetChartSnapshot(snapshotURL string) (*ChartSnapshot, *http.Response, error) {
	snapshotSplit := strings.Split(snapshotURL, "/")
	snapshotID := snapshotSplit[len(snapshotSplit)-1]
	u := fmt.Sprintf("snapshots/%s", snapshotID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	cs := new(ChartSnapshot)
	resp, err := s.client.Do(req, cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, err
}
