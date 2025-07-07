## how to setup python project

create a `requirements.txt` file and add all dependency
for fast api
 
    fastapi
    uvicorn
    sqlalchemy
    alembic
    psycopg2-binary
    slowapi
    python-dotenv
    pyjwt
    passlib
    bcrypt==4.0.1
    python-multipart
    email-validator

### create virtual invironment 
    python3 -m venv .venv

### activete it
    source .venv/bin/activate

### Install the packages from requirements.txt
    pip install -r requirements.txt
