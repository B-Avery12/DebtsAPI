# How to run:
    1 - In the directory with the Makefile run "make build"
    2 - Run "./bin/DebtsAPI"
    3 - For Testing run "make test" from directory where makefile is located

# High Level Overview:
    Started by looking over the prompts to make sure I understood the requirements correctly. Then decided to work on the code in stages. First working on code to retrieve data from the provided endpoints. Second working on code to enrich data according to the requirements. For both sections of code I followed the approach of getting a working solution, then refactored to simplify and decouple code. Finished up with writing tests and verified code ran as expected.

# Description of approach:   
    * Read through problem and decided to get the HTTP API service objects created as well as the ability to access them
    * After creating corresponding API models, create ability to consume from the respective endpoints
    * Wrote code to enrich data from API
    * Wrote tests to verify data was enriched properlly and reworked code as needed
    * Finished up with main.go to actually ingest API data and then enrich.
    * Once code was all done, set up how to run script in a repeatable manner and then double checked work for any errors
    * If I had more time I would do better handling of various errors to make script more robust.

# Relevant Logic:
    * All logic for interacting with the API is in internal/api/
    * All logic for enriching debts is in internal/enricher/

# Design Decisions or Assumptions Made:
    * Separated code to interact with API into it's own package so it could be used if needed else where as well. Also allows for better readability and understanding of the project
    * Created code to enrich debts in it's own folder for the same reason as with the API folder
    * Reworked code to use functions that weren't bound to a specific type
    * Made the assumption that if there are no debts, the script will exit as that means there can be no valid payment plans or payments
    * Assumed incoming data was correct, and if errors were encountered ignored that data and moved on to next data