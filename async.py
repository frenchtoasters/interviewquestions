from asyncio import Future
import asyncio


async def plan(dinner: Future):
    print('Planning my dinner...')
    await asyncio.sleep(1)
    dinner.set_result('Beef')


def pickmeal() -> Future:
    dinner = Future()
    asyncio.create_task(plan(dinner))
    return dinner


async def main():
    my_dinner = pickmeal()
    result = my_dinner

    print(result)


asyncio.run(main())
