Feature: I want to test product upload-api feature

  Scenario: As a client I want to test for upload a product without name - Dirty case
    Given I truncate the "products" table
    And I prepare request with payload:
    ```
    {
      "name": "",
      "quantity": 10,
      "unit": "peace",
      "price": 50,
      "price_unit": "dollar",
      "user_id": 10,
      "company_id": 1
    }
    ```
    Then I send "POST" request to "product/upload"
    And the response code should be 200
    And the response success should be "false" and message should be "Name must be at least 1 character in length"

  Scenario: As a client I want to test for upload a product but quantity less then 1 - Dirty case
    Given I truncate the "products" table
    And I prepare request with payload:
    ```
    {
      "name": "Ram DDR4 P2666 Gigabyte",
      "quantity": -1,
      "unit": "peace",
      "price": 50,
      "price_unit": "dollar",
      "user_id": 10,
      "company_id": 1
    }
    ```
    Then I send "POST" request to "product/upload"
    And the response code should be 200
    And the response success should be "false" and message should be "Quantity must be 1 or greater"

  Scenario: As a client I want to test for upload a product but price less then 0 - Dirty case
    Given I truncate the "products" table
    And I prepare request with payload:
    ```
    {
      "name": "Ram DDR4 P2666 Gigabyte",
      "quantity": 10,
      "unit": "peace",
      "price": -1,
      "price_unit": "dollar",
      "user_id": 10,
      "company_id": 1
    }
    ```
    Then I send "POST" request to "product/upload"
    And the response code should be 200
    And the response success should be "false" and message should be "Price must be greater than 0"

  Scenario: As a client I want to test for upload a product - Happy case
    Given I truncate the "products" table
    And I prepare request with payload:
    ```
    {
      "name": "Ram DDR4 P2666 Gigabyte",
      "quantity": 10,
      "unit": "peace",
      "price": 50,
      "price_unit": "dollar",
      "user_id": 10,
      "company_id": 1
    }
    ```
    Then I send "POST" request to "product/upload"
    And the response code should be 200
    And the response success should be "true" and message should be "OK"
    And Data in "products" table shouble be:
    | name                      | quantity  	| unit 	  | price | price_unit | user_id | company_id |
    | Ram DDR4 P2666 Gigabyte   | 10     	    | peace 	| 50 	  | dollar     | 10      | 1          |