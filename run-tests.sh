docker-compose -f docker-compose-basic.yaml up --build -d 
newman run ./tests/Expense_Tracker_API_Tests.json -x -e ./tests/Development_Environment.json --color
docker-compose -f docker-compose-basic.yaml down