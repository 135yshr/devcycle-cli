package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_Projects(t *testing.T) {
	t.Run("successful list", func(t *testing.T) {
		projects := []Project{
			{
				ID:          "proj-1",
				Key:         "my-project",
				Name:        "My Project",
				Description: "Test project",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          "proj-2",
				Key:         "another-project",
				Name:        "Another Project",
				Description: "",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects" {
				t.Errorf("expected /projects, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(projects)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Projects(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 projects, got %d", len(result))
		}
		if result[0].Key != "my-project" {
			t.Errorf("expected my-project, got %s", result[0].Key)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Project{})
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		result, err := client.Projects(context.Background())

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected 0 projects, got %d", len(result))
		}
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.Projects(context.Background())

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_Project(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		project := Project{
			ID:          "proj-1",
			Key:         "my-project",
			Name:        "My Project",
			Description: "Test project",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/projects/my-project" {
				t.Errorf("expected /projects/my-project, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(project)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.Project(context.Background(), "my-project")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "my-project" {
			t.Errorf("expected my-project, got %s", result.Key)
		}
		if result.Name != "My Project" {
			t.Errorf("expected My Project, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("project not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.Project(context.Background(), "non-existent")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}

func TestClient_CreateProject(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		project := Project{
			ID:        "proj-new",
			Key:       "new-project",
			Name:      "New Project",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/projects" {
				t.Errorf("expected /projects, got %s", r.URL.Path)
			}

			var req CreateProjectRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "New Project" {
				t.Errorf("expected name 'New Project', got %s", req.Name)
			}
			if req.Key != "new-project" {
				t.Errorf("expected key 'new-project', got %s", req.Key)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(project)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.CreateProject(context.Background(), &CreateProjectRequest{
			Name: "New Project",
			Key:  "new-project",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Key != "new-project" {
			t.Errorf("expected new-project, got %s", result.Key)
		}
	})

	t.Run("conflict error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("project already exists"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.CreateProject(context.Background(), &CreateProjectRequest{
			Name: "Duplicate Project",
			Key:  "duplicate-project",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_UpdateProject(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		project := Project{
			ID:        "proj-1",
			Key:       "my-project",
			Name:      "Updated Project",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				t.Errorf("expected PATCH, got %s", r.Method)
			}
			if r.URL.Path != "/projects/my-project" {
				t.Errorf("expected /projects/my-project, got %s", r.URL.Path)
			}

			var req UpdateProjectRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name != "Updated Project" {
				t.Errorf("expected name 'Updated Project', got %s", req.Name)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(project)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL), WithToken("test-token"))
		result, err := client.UpdateProject(context.Background(), "my-project", &UpdateProjectRequest{
			Name: "Updated Project",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Name != "Updated Project" {
			t.Errorf("expected Updated Project, got %s", result.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("project not found"))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL))
		_, err := client.UpdateProject(context.Background(), "non-existent", &UpdateProjectRequest{
			Name: "Updated Project",
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !IsNotFound(err) {
			t.Errorf("expected not found error, got %v", err)
		}
	})
}
