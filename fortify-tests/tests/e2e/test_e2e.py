import pytest
import re
from playwright.sync_api import Page, sync_playwright, expect

@pytest.fixture(scope="session")
def browser():
    with sync_playwright() as playwright:
        browser = playwright.firefox.launch(headless=True)
        yield browser
        browser.close()

@pytest.fixture(scope="function")
def page(browser):
    page = browser.new_page()
    yield page
    page.close()

@pytest.mark.e2e
def test_create_secret(page: Page, fortify_service):
    page.goto(fortify_service.base_url)
    secret_name_field = page.locator(fortify_service.nameFieldId)
    secret_value_field = page.locator(fortify_service.valueFieldId)
    save_button = page.locator(fortify_service.saveBtn)
    show_button = page.locator(fortify_service.showBtn)

    secret_name_field.fill("E2E_Test_Name")
    secret_value_field.fill("E2E_Test_Value")

    save_button.click()
    show_button.click()

    expect(page.locator(fortify_service.nameFieldId)).to_have_value('E2E_Test_Name')
    expect(page.locator(fortify_service.valueFieldId)).to_have_value('E2E_Test_Value')
    expect(page.locator(fortify_service.secretsFieldId)).to_have_value(re.compile('E2E_Test_Name: E2E_Test_Value'))

@pytest.mark.e2e
def test_create_secret_with_generated_value(page: Page, fortify_service):
    page.goto(fortify_service.base_url)
    secret_name_field = page.locator(fortify_service.nameFieldId)
    save_button = page.locator(fortify_service.saveBtn)
    show_button = page.locator(fortify_service.showBtn)
    generate_button = page.locator(fortify_service.generateBtn)

    secret_name_field.fill("E2E_Test_Name_Generated_secret")
    generate_button.click()

    save_button.click()
    show_button.click()

    expect(page.locator(fortify_service.nameFieldId)).to_have_value('E2E_Test_Name_Generated_secret')
    expect(page.locator(fortify_service.valueFieldId)).to_have_value(re.compile(r'.+'))
    # Regular expression is used to match random string generated from the generator service
    expect(page.locator(fortify_service.secretsFieldId)).to_have_value(re.compile(r'E2E_Test_Name_Generated_secret: .+'))
