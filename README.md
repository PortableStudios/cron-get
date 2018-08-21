# cron-get

```
cron:
  image: portable/cron-get:5
  environment:
    SCHEDULE: "0 * * * * *" # Each Minute
    URL: http://example.com
```

## With Timezone use portable/cron-get:6 or later

```
cron:
  image: portable/cron-get:6
  environment:
    TZ: "Australia/Melbourne"
    SCHEDULE: "0 30 2 * * *" # 2:30 am Melbourne Time every day
    URL: http://example.com
```