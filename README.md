# fevm-event-cron-job

### Regenerate swagger doc
swagger doc defined in router api comment.
if edited these comments, need to regenerate swagger doc.

```shell script
swag init -g cmd/busi/main.go
```
### Swagger doc
swagger doc please refer to
```
http://127.0.0.1:7001/busi/swagger/index.html
```
### How to make
```
make # make to see help
```
### Run
    bin/fevm-event --conf conf/service.conf
### Useage
    Call APIs: 
     - /api/v1/deal-proposal-create-event-tracing-cron[POST]
     - /wfil-event-tracing-cron[POST]
    without query paramters by dolphin scheduler every day.


    The gap APIs couldn't backfill the gap information deal to the DB schema has not a unique index.
     - /api/v1/deal-proposal-create-event-tracing
     - /api/v1/wfil-event-tracing