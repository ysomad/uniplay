import logging
from dataclasses import dataclass

import requests
from bs4 import BeautifulSoup


HEADERS = {
    'Referer': 'https://vuzopedia.ru/captcha',
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/111.0',
    'Cookie': 'vzp_uid=eyJpdiI6Im1Bb1hCaWMzUU5jR1ZRdnZhemdCMkE9PSIsInZhbHVlIjoibGt1Y2xIeEVLYzdFd2JvZXprZ01YbzNqd0dlTnM1S0J4bThLcWZNVXFHOUlMdkFcLzk0YWtYREJ1WHNNZkRqOCtrWU1nWGNTYWZqWUw2MktXaEQwZVNkQ2paT2c5MWhrZWJId1AreDdBTGJRPSIsIm1hYyI6IjVmNWFlODFmNjkyZTVkOGFkMTNhODFjYjJjMDliOTFmY2ZjYzY3Y2RmMWU4NDExY2ZjODBkYzUxMDRiMmRkZGYifQ%3D%3D; _ym_uid=16755234021014955968; _ym_d=1679117173; tmr_lvid=748707f9ce14edc0f18f277d80556f9d; tmr_lvidTS=1675523401641; tmr_detect=1%7C1681018172771; _ym_isad=1; october_session=eyJpdiI6IjhFSCtEUlwvdmkrVGxveTR4cHNjblNBPT0iLCJ2YWx1ZSI6Ik9YcGJ1OHFYUkpvcG1SdDFwdFNEWUdkXC82V1hPeWlzOFA1YjlvQVk4RWV4MDFwQWRoMGVtTW1cL1AxN3d6dGZHbzR3bnlmNEhHSXF6WXh5YncyWW82QlJySGg3YUJlSzBKTG5YSE04VGs4eXlDMUEwZDRiakZoOUtDUCttM0FCdUUiLCJtYWMiOiIwZTQ3MTk0Yjk1NzYwMTZkYTA0OTIwYzQ3MzEyZTk2OGNjNzViNGFkNGZlNzY5NWI3Y2U1M2I5NzkwMGMwYmQyIn0%3D; _ym_visorc=b',
    'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8'
}
URL = 'https://vuzopedia.ru'


@dataclass
class Institution:
    city: str
    full_name: str
    short_name: str
    logo_url: str


def collect_institution(uni_path: str, logo_url: str) -> Institution:
    resp = requests.get(URL + uni_path, headers=HEADERS)
    if resp.status_code != 200:
        raise Exception(f'Got error on: {resp.url}, with resp: {resp.text}')

    bs = BeautifulSoup(resp.text, 'html.parser')

    city_tag = bs.select_one('#newChoose span')
    if city_tag == None:
        raise Exception(f'City not found on: {resp.url}')

    short_name_tag = bs.select_one('#newVuzchoos span')
    if short_name_tag == None:
        raise Exception(f'Short name not found on: {resp.url}')

    full_name_tag = bs.select_one('.mainTitle')
    if full_name_tag == None:
        raise Exception(f'Full name not found on: {resp.url}')

    return Institution(
        city=city_tag.text,
        full_name=full_name_tag.text.replace(u'\ue87e', ' ').strip(),
        short_name=short_name_tag.text.strip(),
        logo_url=logo_url
    )


def collect_institutions(universities: bool = True) -> list[Institution]:
    insts = []

    for page_num in range(1, 1000):
        print(f'Scraping page: {page_num}')

        resp = None

        if universities:
            resp = requests.get(f'{URL}/vuz?page={page_num}', headers=HEADERS)
        else:
            resp = requests.get(f'{URL}/ssuzy?page={page_num}', headers=HEADERS)

        if resp.status_code != 200:
            print(f'Got error on page: {page_num} with resp: {resp.text}')
            return insts


        bs = BeautifulSoup(resp.text, 'html.parser')

        uni_tags = bs.select('.vuzesfullnorm')
        if len(uni_tags) == 0:
            print(f'Collected all universities, total amount of pages: {page_num}')
            break

        for tag in uni_tags:
            try:
                uni_path_tag = tag.select_one('.blockAfter a')
                if uni_path_tag == None:
                    raise Exception(f'Link to university not found: {resp.url}')

                logo_url_tag = tag.select_one('.blockAfter a img')
                if logo_url_tag == None:
                    raise Exception(f'Logo of university not found: {resp.url}')

                inst_path = uni_path_tag['href']
                logo_url = logo_url_tag['data-src']

                inst = collect_institution(inst_path, logo_url)
                print(f'Scraped uni: {inst}')
                insts.append(inst)
            except Exception as e:
                print(e)
                continue

    return insts


def save_uni_urls_to_file(urls: list[str]) -> None:
    fp = open('urls.txt', 'w')
    for url in urls:
        fp.write("%s\n" % url)


def main() -> None:
    unis = collect_institutions(False)

    f = open('colleges.sql', 'w')    
    f.write("""INSERT INTO institution(name, short_name, city, logo_url)
VALUES""")

    for uni in unis:
        f.write(f"\n\t('{uni.full_name}', '{uni.short_name}', '{uni.city}', '{uni.logo_url}'),")

    # save_uni_urls_to_file(uni_urls)
    
        
if __name__ == '__main__':
    try:
        main()
    except Exception as e:
        logging.fatal(e)