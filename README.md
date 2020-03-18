# hostname-chainer

for experiment

# usage

host1

```aidl
HOSTNAME=host1 go run main.go
```

host2

```aidl
HOSTNAME=host2 NEXT_HOST_ADDRESS=host1-addr NEXT_HOST_PORT=8080 go run main.go
```

host3

```aidl
HOSTNAME=host3 NEXT_HOST_ADDRESS=host2-addr NEXT_HOST_PORT=8080 go run main.go
```

request to host3

```aidl
curl -i -X POST -H "Content-Type:application/json" -d '[{"host_name":"curl","to":"host3"}]' 'http://host3-addr:8080'
```

result

```aidl
[
    {"host_name":"curl","to":"host3"},
    {"host_name":"host3","to":"host2"},
    {"host_name":"host2","to":"host1"},
    {"host_name":"host1"}
]
```
