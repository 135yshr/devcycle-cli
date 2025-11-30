---
title: "metrics"
weight: 21
---

# metrics

Commands for managing DevCycle metrics. Metrics allow you to measure the impact of your feature flags.

## list

List all metrics in a project.

### Usage

```bash
dvcx metrics list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List all metrics
$ dvcx metrics list -p my-app
KEY              NAME              TYPE         OPTIMIZE FOR
conversion-rate  Conversion Rate   conversion   increase
page-views       Page Views        count        increase
revenue          Revenue           sum          increase

# List metrics in JSON format
$ dvcx metrics list -p my-app -o json
```

---

## get

Get details of a specific metric.

### Usage

```bash
dvcx metrics get <metric-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `metric-key` | The unique key of the metric |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get metric details
$ dvcx metrics get conversion-rate -p my-app

# Get metric details in JSON format
$ dvcx metrics get conversion-rate -p my-app -o json
{
  "_id": "metric-123",
  "key": "conversion-rate",
  "name": "Conversion Rate",
  "type": "conversion",
  "eventType": "purchase",
  "optimizeFor": "increase",
  "description": "Measures purchase conversion rate"
}
```

---

## create

Create a new metric.

### Usage

```bash
dvcx metrics create [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--key` | `-k` | Metric key | Yes |
| `--name` | `-n` | Metric name | Yes |
| `--type` | `-t` | Metric type (count, conversion, sum, average) | Yes |
| `--event-type` | | Event type to track | Yes |
| `--optimize-for` | | Optimization goal (increase, decrease) | Yes |
| `--description` | `-d` | Metric description | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Create a count metric
$ dvcx metrics create -p my-app \
  --key page-views \
  --name "Page Views" \
  --type count \
  --event-type pageview \
  --optimize-for increase \
  --description "Tracks page view events"

# Create a conversion metric
$ dvcx metrics create -p my-app \
  --key signup-rate \
  --name "Signup Rate" \
  --type conversion \
  --event-type signup \
  --optimize-for increase

# Create a sum metric for revenue
$ dvcx metrics create -p my-app \
  --key total-revenue \
  --name "Total Revenue" \
  --type sum \
  --event-type purchase \
  --optimize-for increase
```

### Metric Types

| Type | Description |
|------|-------------|
| `count` | Counts the number of events |
| `conversion` | Measures conversion rate (percentage of users who triggered the event) |
| `sum` | Sums event values |
| `average` | Calculates average event value |

### Optimization Goals

| Goal | Description |
|------|-------------|
| `increase` | Higher values are better |
| `decrease` | Lower values are better |

---

## update

Update an existing metric.

### Usage

```bash
dvcx metrics update <metric-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `metric-key` | The unique key of the metric to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--name` | `-n` | New metric name | No |
| `--description` | `-d` | New metric description | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update metric name
$ dvcx metrics update page-views -p my-app --name "Total Page Views"

# Update metric description
$ dvcx metrics update page-views -p my-app --description "Updated description"
```

### Notes

- Metric key, type, and event type cannot be changed after creation
- Only name and description can be updated

---

## delete

Delete a metric from a project.

### Usage

```bash
dvcx metrics delete <metric-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `metric-key` | The unique key of the metric to delete |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--force` | | Skip confirmation prompt | No |

### Example

```bash
# Delete a metric (with confirmation)
$ dvcx metrics delete page-views -p my-app
Are you sure you want to delete metric 'page-views'? [y/N]: y
Metric 'page-views' deleted successfully

# Delete without confirmation
$ dvcx metrics delete page-views -p my-app --force
```

---

## results

Get metric results for a specific metric.

### Usage

```bash
dvcx metrics results <metric-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `metric-key` | The unique key of the metric |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--environment` | `-e` | Filter by environment | No |
| `--feature` | `-f` | Filter by feature | No |
| `--start-date` | | Start date (YYYY-MM-DD) | No |
| `--end-date` | | End date (YYYY-MM-DD) | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get metric results
$ dvcx metrics results conversion-rate -p my-app
VARIATION    COUNT    VALUE
control      1500     0.032
treatment    1600     0.045

# Filter by environment
$ dvcx metrics results conversion-rate -p my-app --environment production

# Filter by feature
$ dvcx metrics results conversion-rate -p my-app --feature new-checkout

# Filter by date range
$ dvcx metrics results conversion-rate -p my-app \
  --start-date 2024-01-01 \
  --end-date 2024-01-31

# Combine all filters
$ dvcx metrics results conversion-rate -p my-app \
  --environment production \
  --feature new-checkout \
  --start-date 2024-01-01 \
  --end-date 2024-01-31

# Get results in JSON format
$ dvcx metrics results conversion-rate -p my-app -o json
{
  "data": [
    {
      "variationKey": "control",
      "count": 1500,
      "value": 0.032
    },
    {
      "variationKey": "treatment",
      "count": 1600,
      "value": 0.045
    }
  ]
}
```

### Notes

- Results show metric values per variation
- Use date filters to analyze specific time periods
- Environment and feature filters help narrow down results
