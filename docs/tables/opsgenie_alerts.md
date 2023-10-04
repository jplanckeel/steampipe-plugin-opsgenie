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
from opsgenie_alerts;
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
   COUNT(*) AS NumberOfAlerts 
from
   opsgenie_alerts
```

### Count alert by message priority

```sql
select
   message,
   priority,
   COUNT(*) AS NumberOfAlerts 
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
   COUNT(*) AS NumberOfAlerts 
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
   COUNT(*) AS NumberOfAlerts 
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

### Selectalert by message priority top 10 in last 30days

```sql
select
   message,
   created_at 
from
   opsgenie_alerts 
where
   created_at >= now() - '7 days' :: interval;
```
