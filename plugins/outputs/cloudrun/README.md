### CloudRun


CloudRun Conf file syntax:
```
# A plugin that can transmit metrics over HTTP
[[outputs.cloudrun]]
  ## URL is the address to send metrics to
 url = 'https://metrics-proxy-gasdy23432653q-uc.a.run.app'


  ## Timeout for HTTP message
  timeout = "30s"

  cloudrun_email = "svc-wf-proxy-cr@io1-datalake-dev.iam.gserviceaccount.com"
  json_file_location = "C:/Scripts/Go/metrics-proxy-cr.json"
  data_format = "wavefront"
  convert_paths = false
```

Developed by Casey Flanigan, Zachary Mares, and John Farrington
