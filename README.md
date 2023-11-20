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
    Call api: /deal-proposal-create-event-tracing-cron[POST] without query paramters by dolphin scheduler every day.

    If gap occurs, manually call /deal-proposal-create-event-tracing with query parameters: /deal-proposal-create-event-tracing?from=${min height}&to=${max height}