# Table: opsgenie_alerts

//TODO:

## Examples

### Basic group info

```sql
SELECT
  alert_id,
  message,
  owner_team_id,
  created_at,
  status
FROM opsgenie_alerts 
WHERE owner_team_id = '2bca1847-c59a-4e7d-9433-d29858d10cd6'
```


### Alert for a team

```sql
SELECT
  alert_id,
  message,
  owner_team_id,
  created_at,
  status
FROM opsgenie_alerts 
WHERE owner_team_id = '<<TEAM-ID>>'
```

### Count alert by message priority

```sql
SELECT
  message,
  priority,
  COUNT(*) AS NumberOfAlerts
FROM opsgenie_alerts
GROUP BY message,
  priority
ORDER BY NumberOfAlerts DESC;
```
