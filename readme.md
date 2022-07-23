How to run:

High Level Overview:

Description of approach:   
    * Read through problem and decided to get the HTTP API service objects created as well as the ability to access them
    * After creating corresponding API models, create ability to consume from the respective endpoints
    * With ability to interact with API done, created code in main.go to retrieve the needed data from the endpoints
    * Feeling it would be easier to interact with the various slices if they were maps added ability to convert them to maps

Relevant Logic:
    * All logic for interacting with the API is in internal/api/

Design Decisions or Assumptions Made:
    * Separated code to interact with API into it's own package so it could be used if needed else where as well. Also allows for better readability and understanding of the project
    * Made the assumption that if there are no debts, the script will exit as that means there can be no valid payment plans or payments