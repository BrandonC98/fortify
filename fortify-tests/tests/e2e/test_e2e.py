import pytest
import re
from playwright.sync_api import Page, sync_playwright, expect

@pytest.fixture(scope="session")
def browser():
    with sync_playwright() as playwright:
        browser = playwright.firefox.launch(headless=False)
        yield browser
        browser.close()

@pytest.fixture(scope="function")
def page(browser):
    page = browser.new_page()
    yield page
    page.close()

@pytest.mark.e2e
def test_create_secret(page: Page):
    page.goto("http://localhost:9002/")
    secret_name_field = page.locator('#nameField') 
    secret_value_field = page.locator('#valueField') 
    save_button = page.locator('#saveBtn') 
    show_button = page.locator('#showBtn') 

    secret_name_field.fill("E2E_Test_Name")
    secret_value_field.fill("E2E_Test_Value")

    save_button.click()
    show_button.click()

    expect(page.locator('#nameField')).to_have_value('E2E_Test_Name')
    expect(page.locator('#valueField')).to_have_value('E2E_Test_Value')
    expect(page.locator('#secretList')).to_have_value(re.compile('E2E_Test_Name: E2E_Test_Value'))

@pytest.mark.e2e
def test_create_secret_with_generated_value(page: Page):
    page.goto("http://localhost:9002/")
    secret_name_field = page.locator('#nameField') 
    save_button = page.locator('#saveBtn') 
    show_button = page.locator('#showBtn') 
    generate_button = page.locator('#generateBtn') 

    secret_name_field.fill("E2E_Test_Name")
    generate_button.click()

    save_button.click()
    show_button.click()

    expect(page.locator('#nameField')).to_have_value('E2E_Test_Name')
    expect(page.locator('#valueField')).to_have_value(re.compile(r'.+'))
    # Regular expression is used to match random string generated from the generator service
    expect(page.locator('#secretList')).to_have_value(re.compile(r'E2E_Test_Name: .+'))
