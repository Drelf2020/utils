import requests
import json

print(
    json.dumps(
        requests.get("https://api.bilibili.com/x/mv/tag").json(),
        indent=4,
        ensure_ascii=False,
    )
)
