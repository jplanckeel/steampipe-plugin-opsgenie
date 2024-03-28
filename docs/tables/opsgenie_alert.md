# Table: opsgenie_alert

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
  opsgenie_alert;
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
  opsgenie_alert
where
  owner_team_id = '<<TEAM-ID>>';
```

### Total no of alerts

```sql
select
  count(*) as number_of_alerts
from
  opsgenie_alert;
```

### Alert by message priority

```sql
select
  message,
  priority,
  count(*) as number_of_alerts
from
  opsgenie_alert
group by
  message,
  priority
order by
  number_of_alerts desc;
```

### Top 10 alerts by message priority
```sql
select
  message,
  priority,
  count(*) as number_of_alerts
from
  opsgenie_alert
group by
  message,
  priority
order by
  number_of_alerts desc limit 10;
```

### Top 10 alerts by message priority in last 30 days

```sql
select
  message,
  priority,
  count(*) as number_of_alerts
from
  opsgenie_alert
where
  created_at >= now() - '30 days' :: interval
group by
  message,
  priority
order by
  number_of_alerts desc limit 10;
```

### List alerts that are less than seven days old

```sql
select
  message,
  created_at
from
  opsgenie_alert
where
  created_at >= now() - '7 days' :: interval;
```

### List top 15 alerts with delta between 2 weeks

```sql
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
    opsgenie.opsgenie_alert
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
    select
      MAX(month) - 1
    from
      alert_by_month
  )
order by
  month desc,
  "Nb Alerts" desc limit 15;
```
