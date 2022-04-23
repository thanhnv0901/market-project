Feature: I want to test product list-api feature


  Scenario: As a client I want to test for list a product not existing - Dirty case
    Given I truncate the "products" table
    And I setup table "products" table with data:
    | id | name                      | quantity  	| unit 	  | price | price_unit | user_id | company_id |
    | 1  | Ram DDR4 P2666 Gigabyte   | 10     	  | peace   | 50 	  | dollar     | 10      | 1          |
    Then I send "GET" request to "product/list/12"
    And the response code should be 200
    And the response success should be "true" and message should be "OK"
    And the response data should contain
    | id | name                      | quantity  	| unit 	  | price | price_unit | user_id | company_id |
    

  Scenario: As a client I want to test for list a product - Happy case
    Given I truncate the "products" table
    And I setup table "products" table with data:
    | id | name                      | quantity  	| unit 	  | price | price_unit | user_id | company_id |
    | 1  | Ram DDR4 P2666 Gigabyte   | 10     	  | peace   | 50 	  | dollar     | 10      | 1          |
    Then I send "GET" request to "product/list/1"
    And the response code should be 200
    And the response success should be "true" and message should be "OK"
    And the response data should contain
    | id | name                      | quantity  	| unit 	  | price | price_unit | user_id | company_id |
    | 1  | Ram DDR4 P2666 Gigabyte   | 10     	  | peace   | 50 	  | dollar     | 10      | 1          |
