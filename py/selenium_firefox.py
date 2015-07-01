from selenium.webdriver.firefox.webdriver import WebDriver

selenium = WebDriver()

def fill_out_form():
    selenium.get("https://login.mailchimp.com/")

    username_input = selenium.find_element_by_id("username")
    username_input.send_keys('sbhackerspace@gmail.com')

    password_input = selenium.find_element_by_id("password")
    password_input.send_keys('ouractualpassword')


fill_out_form()
