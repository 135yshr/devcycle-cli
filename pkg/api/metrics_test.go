package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_Metrics(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		metrics := []Metric{
			{
				ID:          "met-1",
				Key:         "click-rate",
				Name:        "Click Rate",
				Type:        "count",
				EventType:   "click",
				OptimizeFor: "increase",
				Description: "Measure click events",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          "met-2",
				Key:         "load-time",
				Name:        "Load Time",
				Type:        "average",
				EventType:   "pageLoad",
				OptimizeFor: "decrease",
				Description: "Measure page load time",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/metrics" {
				t.Errorf("expected /projects/my-project/metrics, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(metrics)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Metrics(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 metrics, got %d", len(result))
		}
		if result[0].Key != "click-rate" {
			t.Errorf("expected click-rate, got %s", result[0].Key)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Metric{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.Metrics(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 metrics, got %d", len(result))
		}
	})
}

func TestClient_Metric(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		metric := Metric{
			ID:          "met-1",
			Key:         "click-rate",
			Name:        "Click Rate",
			Type:        "count",
			EventType:   "click",
			OptimizeFor: "increase",
			Description: "Measure click events",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/metrics/click-rate" {
				t.Errorf("expected /projects/my-project/metrics/click-rate, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(metric)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Metric(context.Background(), "my-project", "click-rate")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "click-rate" {
			t.Errorf("expected click-rate, got %s", result.Key)
		}
		if result.Name != "Click Rate" {
			t.Errorf("expected Click Rate, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("metric not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.Metric(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateMetric(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		metric := Metric{
			ID:          "met-new",
			Key:         "new-metric",
			Name:        "New Metric",
			Type:        "count",
			EventType:   "purchase",
			OptimizeFor: "increase",
			Description: "A new metric",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/metrics" {
				t.Errorf("expected /projects/my-project/metrics, got %s", r.URL.Path)
			}

			var req CreateMetricRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "New Metric" {
				t.Errorf("expected name 'New Metric', got %s", req.Name)
			}
			if req.Key != "new-metric" {
				t.Errorf("expected key 'new-metric', got %s", req.Key)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(metric)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CreateMetric(context.Background(), "my-project", &CreateMetricRequest{
			Name:        "New Metric",
			Key:         "new-metric",
			Type:        "count",
			EventType:   "purchase",
			OptimizeFor: "increase",
			Description: "A new metric",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "new-metric" {
			t.Errorf("expected new-metric, got %s", result.Key)
		}
	})

	t.Run("conflict error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("metric already exists"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CreateMetric(context.Background(), "my-project", &CreateMetricRequest{
			Name:        "Duplicate Metric",
			Key:         "duplicate-metric",
			Type:        "count",
			EventType:   "click",
			OptimizeFor: "increase",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_UpdateMetric(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		metric := Metric{
			ID:          "met-1",
			Key:         "click-rate",
			Name:        "Updated Metric",
			Type:        "count",
			EventType:   "click",
			OptimizeFor: "increase",
			Description: "Updated description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/metrics/click-rate" {
				t.Errorf("expected /projects/my-project/metrics/click-rate, got %s", r.URL.Path)
			}

			var req UpdateMetricRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "Updated Metric" {
				t.Errorf("expected name 'Updated Metric', got %s", req.Name)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(metric)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateMetric(context.Background(), "my-project", "click-rate", &UpdateMetricRequest{
			Name: "Updated Metric",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Name != "Updated Metric" {
			t.Errorf("expected Updated Metric, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("metric not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateMetric(context.Background(), "my-project", "non-existent", &UpdateMetricRequest{
			Name: "Updated Metric",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_DeleteMetric(t *testing.T) {
	t.Run("successful delete", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/metrics/click-rate" {
				t.Errorf("expected /projects/my-project/metrics/click-rate, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		err := client.DeleteMetric(context.Background(), "my-project", "click-rate")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("metric not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		err := client.DeleteMetric(context.Background(), "my-project", "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_MetricResults(t *testing.T) {
	t.Run("successful results without options", func(t *testing.T) {
		results := MetricResults{
			Data: []MetricResultData{
				{VariationKey: "control", Count: 100, Value: 0.5},
				{VariationKey: "treatment", Count: 150, Value: 0.75},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project/metrics/click-rate/results" {
				t.Errorf("expected /projects/my-project/metrics/click-rate/results, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(results)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.MetricResults(context.Background(), "my-project", "click-rate", nil)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result.Data) != 2 {
			t.Errorf("expected 2 results, got %d", len(result.Data))
		}
		if result.Data[0].VariationKey != "control" {
			t.Errorf("expected control, got %s", result.Data[0].VariationKey)
		}
	})

	t.Run("successful results with options", func(t *testing.T) {
		results := MetricResults{
			Data: []MetricResultData{
				{VariationKey: "control", Count: 50, Value: 0.4},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Query().Get("environment") != "production" {
				t.Errorf("expected environment=production, got %s", r.URL.Query().Get("environment"))
			}
			if r.URL.Query().Get("feature") != "my-feature" {
				t.Errorf("expected feature=my-feature, got %s", r.URL.Query().Get("feature"))
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(results)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.MetricResults(context.Background(), "my-project", "click-rate", &MetricResultsOptions{
			Environment: "production",
			Feature:     "my-feature",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result.Data) != 1 {
			t.Errorf("expected 1 result, got %d", len(result.Data))
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("metric not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.MetricResults(context.Background(), "my-project", "non-existent", nil)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
