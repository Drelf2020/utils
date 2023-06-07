"""
ivitools.touch

获取 ivi 文件信息
"""
import zipfile
import json


def touch_ivi(path: str):
    ivi = zipfile.ZipFile(open(path, "rb"))
    info = ivi.open("bilivideo.json").read()
    print(json.loads(info))
    return json.loads(info)
