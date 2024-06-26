import asyncio
import httpx

async def fetch(session, url):
    try:
        response = await session.get(url)
        return response.text
    except httpx.RequestError as exc:
        print(f"Request failed: {exc}")

async def fetch_ips(num_requests):
    url = 'https://ifconfig.me/ip'
    proxy = "http://terminaluwuea7ejicn6tve-package-isp:xofmck9gmnph@147.75.102.51:1999"
    async with httpx.AsyncClient(proxies=proxy) as session:
        tasks = [fetch(session, url) for _ in range(num_requests)]
        return await asyncio.gather(*tasks)

async def main():
    num_concurrent_requests = 500
    results = await fetch_ips(num_concurrent_requests)
    
    for idx, result in enumerate(results, start=1):
        print(f"Request {idx}: {result.strip()}")

if __name__ == '__main__':
    asyncio.run(main())
