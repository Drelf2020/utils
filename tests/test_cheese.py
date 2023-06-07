# bilibili_api.cheese

from bilibili_api import cheese

course = cheese.CheeseList(256)
cheese_ep = cheese.CheeseVideo(5556)


async def test_a_CheeseList_get_meta():
    return await course.get_meta()


async def test_b_CheeseList_get_courses():
    return await course.get_list()


async def test_c_CheeseVideo_get_download_url():
    return await cheese_ep.get_download_url()
