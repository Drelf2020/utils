"""
bilibili_api.utils.initial_state

用于获取页码的初始化信息
"""
from .Credential import Credential
from ..exceptions import *
from .network_httpx import get_session
from .short import get_real_url
import httpx
import re
import json


async def get_initial_state(url: str, credential: Credential = Credential()) -> dict:
    '''
    异步获取初始化信息
    '''
    try:
        session = get_session()
        resp = await session.get(
            url,
            cookies=credential.get_cookies(),
            headers={"User-Agent": "Mozilla/5.0"},
            follow_redirects=True,
        )
    except Exception as e:
        raise e
    else:
        content = resp.text
        pattern = re.compile(r"window.__INITIAL_STATE__=(\{.*?\});")
        match = re.search(pattern, content)
        if match is None:
            raise ApiException("未找到相关信息")
        try:
            content = json.loads(match.group(1))
        except json.JSONDecodeError:
            raise ApiException("信息解析错误")

        return content

def get_initial_state_sync(url: str, credential: Credential = Credential()) -> dict:
    '''
    同步获取初始化信息
    '''
    try:
        resp = httpx.get(
            url,
            cookies=credential.get_cookies(),
            headers={"User-Agent": "Mozilla/5.0"},
            follow_redirects=True,
        )
    except Exception as e:
        raise e
    else:
        content = resp.text
        pattern = re.compile(r"window.__INITIAL_STATE__=(\{.*?\});")
        match = re.search(pattern, content)
        if match is None:
            raise ApiException("未找到相关信息")
        try:
            content = json.loads(match.group(1))
        except json.JSONDecodeError:
            raise ApiException("信息解析错误")

        return content