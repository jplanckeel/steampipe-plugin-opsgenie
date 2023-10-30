# Table: opsgenie_alerts

List all alerts in the Opsgenie account.

## Examples

### Basic group info

```sql
select
  alert_id,
  message,
  owner_team_id,
  created_at,
  status 
from
  opsgenie_alerts;
```

### Alert for a team

```sql
select
  alert_id,
  message,
  owner_team_id,
  created_at,
  status 
from
  opsgenie_alerts 
where
  owner_team_id = '<<TEAM-ID>>';
```

### Count alert

```sql
select
  count(*) AS NumberOfAlerts 
from
  opsgenie_alerts
```

### Count alert by message priority

```sql
select
  message,
  priority,
  count(*) AS NumberOfAlerts 
from
  opsgenie_alerts 
group by
  message,
  priority 
order by
  NumberOfAlerts desc;
```

### Count alert by message priority top 10
```sql
select
  message,
  priority,
  count(*) AS NumberOfAlerts 
from
  opsgenie_alerts 
group by
  message,
  priority 
order by
  NumberOfAlerts desc limit 10;
```

### Count alert by message priority top 10 in last 30days

```sql
select
  message,
  priority,
  count(*) AS NumberOfAlerts 
from
  opsgenie_alerts 
where
  created_at >= now() - '30 days' :: interval 
group by
  message,
  priority 
order by
  NumberOfAlerts desc limit 10;
```

### Select alert by message priority top 10 in last 30days

```sql
select
  message,
  created_at 
from
  opsgenie_alerts 
where
  created_at >= now() - '7 days' :: interval;
```

### Select Top 15 alerts with delta between 2 weeks 

```
with alert_by_month as 
(
  select
    message,
    count(*) as "Nb Alerts",
    date_part('month', created_at) as month,
    lag(count(*), 1) over (partition by message 
  order by
    date_part('month', created_at)) as "Nb Alerts Sprint - 1" 
  from
    opsgenie.opsgenie_alerts 
  where
    created_at >= now() - '5 months' :: interval 
  group by
    message,
    month 
  order by
    "Nb Alerts" desc 
)
select
  month as "Month",
  message as "Alert",
  "Nb Alerts",
  "Nb Alerts Sprint - 1",
  "Nb Alerts" - "Nb Alerts Sprint - 1" AS "Delta",
  ROUND(100.0 * (("Nb Alerts" - "Nb Alerts Sprint - 1") / "Nb Alerts Sprint - 1"::decimal), 2) AS "Delta % " 
from
  alert_by_month 
where
  month = 
  (
    SELECT
      MAX(month) - 1 
    FROM
      alert_by_month 
  )
order by
  month desc,
  "Nb Alerts" desc LIMIT 15;
```
