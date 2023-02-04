import logging
import json
from dataclasses import dataclass
from re import I

import requests
from bs4 import BeautifulSoup


URL = 'https://vuzoteka.ru/%D0%B2%D1%83%D0%B7%D1%8B'


@dataclass
class Institution:
    full_name: str
    short_name: str
    city: str
    logo_url: str


def main() -> None:
    resp = requests.get(URL)
    if resp.status_code != 200:
        raise Exception(f'Unsuccessful response: {resp.url}, {resp.text}')

    bs = BeautifulSoup(resp.text, 'html.parser')

    pagination_div = bs.select_one('#pagination-wrapper')
    if pagination_div == None:
        raise Exception('Last page not found')

    last_page_div = pagination_div.select_one(':nth-child(4)')
    if last_page_div == None:
        raise Exception('Last page not found')

    last_page = int(last_page_div.text)

    # institutions: list[Institution] = []

    with open('migration.sql', 'w') as outfile:
        outfile.write("""BEGIN
;

INSERT INTO university(long_name, short_name, city, logo_url)
VALUES""")

        for current_page in range(1, last_page+1):
            resp = requests.get(f'{URL}?page={current_page}')

            if current_page == 1:
                resp = requests.get(URL)

            bs = BeautifulSoup(resp.text, 'html.parser')

            if resp.status_code != 200:
                raise Exception(f'Unsuccessful response: {resp.url}, {resp.text}')

            rows_container = bs.select_one('.institute-rows')
            if rows_container == None:
                raise Exception('Not found university rows')

            rows = rows_container.select('.institute-row')

            i = 1
            for row in rows:
                if row == None:
                    print('Found None row')
                    continue

                logo_tag = row.select_one('.institute-search-logo a img')
                if logo_tag == None:
                    raise Exception('Logo container not found')

                university_tag = row.select_one('.institute-search-title')
                if university_tag == None:
                    raise Exception('University name not found')

                city_tag = row.select('.institute-search-value')
                if city_tag == None or len(city_tag) < 2:
                    raise Exception('City tag not found')

                parts = university_tag.text.split('–')
                if len(parts) < 1:
                    raise Exception(f'Invalid institution name found: {university_tag.text}')

                full_name = university_tag.text
                if len(parts) >= 2:
                    full_name = parts[1]

                short_name = parts[0]

                city = city_tag[1].text.strip()
                logo_url = f'https://{logo_tag["data-src"][2:]}'

                # institution = Institution(
                #     full_name=full_name.strip(), 
                #     short_name=short_name.strip(),
                #     city=city, 
                #     logo_url=logo_url
                # )

                # institutions.append(institution)

                outfile.write(f"\n\t('{full_name.strip()}', '{short_name.strip()}', '{city}', '{logo_url}')")

                if current_page == last_page and i == len(rows):
                    outfile.write(';\n\nCOMMIT;')
                    break
                    
                outfile.write(',')
                i += 1

    
    # json_string = json.dumps([institution.__dict__ for institution in institutions], ensure_ascii=False, indent=2)

    # with open('institutions.json', 'w') as f:
    #     f.write(json_string)

    # print(len(institutions))


if __name__ == '__main__':
    try:
        main()
    except Exception as e:
        logging.fatal(e)