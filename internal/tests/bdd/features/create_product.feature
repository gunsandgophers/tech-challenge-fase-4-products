Feature: Create Product
  I need to be able to create new Products

  Scenario: Create a new SANDWICH
    Given That I need to create a new product via API
    And have the category as "SANDWICH"
    And name as "X-Burguer" 
    And price as "30.5" 
    And description as "Greate Burger"
    And image as "Image ..."
    When I send the data
    Then the product "X-Burguer" should be added to the list of products in the "SANDWICH" category
