# cron-get

```
cron:
  image: portable/cron-get
  environment:
    SCHEDULE: "0 * * * * *" # Each Minute
    URL: http://example.com
```