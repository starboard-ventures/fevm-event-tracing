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
    Call api: /event-tracking[POST] without query paramters by dolphin scheduler per 60s.

    If gap occurs, manually call /event with query parameters: /event?from=${min height}&to=${max height}