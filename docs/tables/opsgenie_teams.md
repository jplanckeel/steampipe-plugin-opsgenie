# Table: opsgenie_teams

Many IT organizations are divided into operational units. This allows them to manage, design and increase the efficiency of their business operations to meet their internal requirements. Some organize by technical specialization, others by activity, some by services, by geography, or any combination of those (or others). Whichever structure the organization chooses, these operational units remain responsible for the problems which occur in their environment.

## Examples

### Basic group info

```sql
select 
  team_id,
  name,
 description
from 
  opsgenie_teams;
```
